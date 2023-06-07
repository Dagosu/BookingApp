package database

import (
	"context"
	"fmt"
	"log"
	sync "sync"
	"time"

	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/gohelpers/database/domain"
	op "github.com/Dagosu/BookingApp/gohelpers/operations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PollingInterval is the interval at which the subscription query is reevaluated in order to catch documents missed because of time-based queries.
// This is currently disabled, will be used for rolling window subscriptions.
var PollingInterval = 60 * time.Second

type GrpcModel interface {
	GetId() string
}

// GenericType is an interface that facilitates instantiating models and sending updates to streams
type GenericType interface {
	// New returns a genericType literal, the id is an optional parameter which can be used to set the returned document Id.
	// this is used for creating DELETE list responses in which we need to set the document Id to a known value.
	New(id ...string) interface {
		GetId() string
	}
	SendResponse(data interface{}, operationType dt.OperationType) error
}

type HookBeforeResponseType func(GrpcModel) GrpcModel

type GenericTypeWithHooks interface {
	HookBeforeResponse(GrpcModel) GrpcModel
}

type subscription struct {
	gt GenericType

	HookBeforeResponse HookBeforeResponseType

	// mutex represents the ugt mutex used to safely add and remove documentsIDs from the array and also ensure correct messages being sent out to subscripbers
	mutex sync.Mutex
	// documentIds tracks documents that the client currently holds
	documentIds []string
	// filter represents the query to be executed and watched
	filter []bson.D
	// index used for managing adding and removing subscription
	index int
	ctx   context.Context
}

// Checks if current subscription contains the correct documents against the subscription filter, fixes time based queries that are not
// Generates document insert and delete watch changes for inconsistencies
//
// example of inconsistent documents for subscription:
//
//	query = filter by `created_at` between T0 and T1
//	P     -  Polling execution
//	D     -  Documents matching the query window
//
// ------ T0 ------ D -- T1 ----
// - P0 ----- P1 ----- P2 ----- P3 -
// P0 no documents are being sent, no documents match the filter
// P1 documents being sent as inserts to subscriber
// P2 no documents being sent, all matching documents have already been sent to the subscriber subscriptions.documentIds
// P3 documents are being sent as `deletes` as they no longer match the filter.
//
// watched 1 2 3
// returned  2 3 4
//
// generates watch changes:
// -> update for 4
// -> delete for 1
func (s *subscription) fakeUpdates(c domain.MongoCollection, f []bson.D) error {
	p := append(
		f,
		bson.D{{
			"$group", bson.D{
				{"_id", "all"},
				{"ids", bson.D{
					{"$push", "$_id"},
				}},
			},
		}},
	)
	cursor, err := c.Aggregate(s.ctx, p)
	if err != nil {
		return err
	}

	defer cursor.Close(s.ctx)
	if !cursor.Next(s.ctx) {
		if len(s.documentIds) == 1 {
			s.handleDocumentChange(&watchChange{ID: s.documentIds[0], OperationType: dt.OperationType_DELETE}, c)
		}
		return nil
	}

	type aggregateResult struct {
		Ids []string `bson:"ids"`
	}
	r := &aggregateResult{}
	err = cursor.Decode(r)
	if err != nil {
		return err
	}

	// trigger INSERT changes
	for _, id := range r.Ids {
		if op.Contains(id, s.documentIds) {
			continue
		}

		s.handleDocumentChange(&watchChange{ID: id, OperationType: dt.OperationType_INSERT}, c)
	}

	// trigger DELETE changes
	for _, id := range s.documentIds {
		if op.Contains(id, r.Ids) {
			continue
		}

		s.handleDocumentChange(&watchChange{ID: id, OperationType: dt.OperationType_DELETE}, c)
	}

	return nil
}

func (s *subscription) fetchDocumentsAndServe(c domain.MongoCollection, f []bson.D) error {
	cursor, err := s.getQueryCursor(c, f)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		data := s.gt.New()
		err := cursor.Decode(data)
		if err != nil {
			// log.Println("Document error, cannot cursor.Decode, continuing:", fmt.Sprintf("%+v", err), cursor.Current())
			s.gt.SendResponse(nil, dt.OperationType_ERROR)

			continue
		}

		// process hooks (To be used when needed)
		// if s.HookBeforeResponse != nil {
		// 	data = s.HookBeforeResponse(data)
		// }

		s.documentIds = op.AppendElement(s.documentIds, data.GetId())
		s.gt.SendResponse(data, dt.OperationType_INSERT)
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown cursor error: %v", err))
	}

	// send initial completion flag to remove loading spinner
	s.gt.SendResponse(nil, dt.OperationType_READY)
	return nil
}

func (s *subscription) getDocument(c domain.MongoCollection, f []bson.D) (interface{}, error) {
	cursor, err := s.getQueryCursor(c, f)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)
	if !cursor.Next(s.ctx) {
		return nil, nil
	}

	data := s.gt.New()
	err = cursor.Decode(data)
	if err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, cursor.Err()
	}

	return data, nil
}

func (s *subscription) getQueryCursor(c domain.MongoCollection, f []bson.D) (domain.MongoCursor, error) {
	// todo aggregation vs find should be supported later on.
	// https://www.mongodb.com/docs/manual/reference/method/cursor.allowDiskUse/
	return c.Aggregate(s.ctx, f, options.Aggregate().SetAllowDiskUse(true))
}

// handleDocumentChange keeps subscriber in sync with initial subscription query (s.filter) and ongoing document changes:
//
// It call the s.gt.sendResponse function which send updates on the stream along with the corresponding operation (e.g.: INSERT DELETE UPDATE see operation.proto OperationType)
//
// s.documentIds represents the documents the subscriber currently got from the active subscription
//
// for a wc *WatchChange
//
//	DELETE operation ->
//		if the s.documentIds contains the received wc.ID:
//			a DELETE operation is send to the subscriber.
//		else
//			the change is ignored.
//	INSERT operation ->
//		if the document matches the s.filter:
//			the document is returned along with a INSERT operation
//		else
//			the change is ignored.
//	UPDATE operation ->
//		if the document matches the s.filter and document hasn't already been returned
//			the document is returned along with a UPDATE operation
//		if the document matches the s.filter and the document has no already been returned:
//			the document is returned along with a INSERT operation
//		else
//			the change is ignored.
func (s *subscription) handleDocumentChange(wc *watchChange, c domain.MongoCollection) {
	// we're using a mutex here in order to prevent inconsistent data being delivered to the client.
	//
	// e.g. 2 updates occur on the same document: both deliver an "insert" even if the second one should be update
	s.mutex.Lock()
	defer s.mutex.Unlock()

	wcID := wc.ID

	if wc.OperationType == dt.OperationType_DELETE {
		s.operationDelete(wcID)
		return
	}

	subFilter := make([]primitive.D, len(s.filter)+1)
	subFilter[0] = primitive.D{
		{Key: "$match", Value: bson.D{
			{Key: "_id", Value: wcID},
		}},
	}
	for i, v := range s.filter {
		subFilter[i+1] = v
	}

	d, err := s.getDocument(c, subFilter)
	if err != nil {
		return
	}

	switch wc.OperationType {
	case dt.OperationType_INSERT:
		s.operationInsert(wcID, d)
	case dt.OperationType_UPDATE:
		s.operationUpdate(wcID, d)
	}
}

// operationUpdate handles an update watch change for a document d, d may be nil meaning the document with _id: id doesn't match the subscription filter
func (s *subscription) operationUpdate(id string, d interface{}) {
	contains := op.Contains(id, s.documentIds)
	if d == nil {
		if contains {
			d := s.gt.New(id)

			s.documentIds = op.UnsetElement(s.documentIds, id)
			s.gt.SendResponse(d, dt.OperationType_DELETE)
			return
		}

		return
	}

	if contains {
		s.gt.SendResponse(d, dt.OperationType_UPDATE)
		return
	}

	s.documentIds = op.AppendElement(s.documentIds, id)
	s.gt.SendResponse(d, dt.OperationType_INSERT)
}

// operationInsert handles an insert watch change for a document d, d may be nil meaning the document with _id: id doesn't match the subscription filter
func (s *subscription) operationInsert(id string, d interface{}) {
	if d == nil {
		return
	}

	s.documentIds = op.AppendElement(s.documentIds, id)
	s.gt.SendResponse(d, dt.OperationType_INSERT)
}

// operationDelete send a delete operation if the document with _id: id has been served to the subscription
func (s *subscription) operationDelete(id string) {
	// ignore deleted documents that weren't returned on subscription
	if !op.Contains(id, s.documentIds) {
		return
	}

	d := s.gt.New(id)
	s.documentIds = op.UnsetElement(s.documentIds, id)

	s.gt.SendResponse(d, dt.OperationType_DELETE)
}

func (sm *SubscriptionsMux) logSubscriptionsData() {
	for {
		time.Sleep(PollingInterval)

		// only log stats for flights collection for now
		if sm.c.Name() == "flights" {
			sm.ops <- func(subs *[]*subscription) {
				var subDocIDS int
				for _, s := range *subs {
					subDocIDS += len(s.documentIds)
				}

				log.Printf("subscription data for collection: %s \n total subscriptions:%d \n total documentIDS:%d", sm.c.Name(), len(*subs), subDocIDS)
			}
		}
	}
}

func (sm *SubscriptionsMux) loop() {
	var subscriptions []*subscription

	fmt.Println("Starting subscription muxer")
	for op := range sm.ops {
		op(&subscriptions)
	}
}

func (sm *SubscriptionsMux) remove(sub *subscription) {
	sm.ops <- func(subs *[]*subscription) {
		i := sub.index
		l := len(*subs)

		(*subs)[i] = (*subs)[l-1]
		// update moved subscription index.
		(*subs)[i].index = i
		(*subs)[l-1] = nil
		*subs = (*subs)[:l-1]
	}
}

func (sm *SubscriptionsMux) add(sub *subscription) {
	sm.ops <- func(subs *[]*subscription) {
		// set index for subscription, makes deletion easier
		i := len(*subs)
		sub.index = i

		*subs = append(*subs, sub)
	}
}

// ServeSubscription sends initial matched documents and registers the new subscription in order to receive future updates
func (sm *SubscriptionsMux) ServeSubscription(ctx context.Context, gt GenericType, c domain.MongoCollection, f []bson.D) error {
	s := &subscription{
		gt:          gt,
		mutex:       sync.Mutex{},
		documentIds: []string{},
		filter:      f,
		ctx:         ctx,
	}

	if gtWithHooks, ok := gt.(GenericTypeWithHooks); ok {
		// forward hooks to subscription
		s.HookBeforeResponse = gtWithHooks.HookBeforeResponse
	}

	// add subscription to SubscriptionMux
	sm.add(s)

	// todo make sure already received documents don't get sent with an INSERT operation -> use UPDATE
	err := s.fetchDocumentsAndServe(sm.c, f)
	if err != nil {
		return err
	}

	// enable polling queries.
	ticker := time.NewTicker(PollingInterval)
	quit := make(chan struct{})

	defer close(quit)
	go func() {
		for {
			select {
			case <-ticker.C:
				err = s.fakeUpdates(sm.c, f)
			case <-quit:
				log.Println("stopped polling.")
				ticker.Stop()
				return
			}
		}
	}()

	// todo
	// between sendingInitialDocuments and registering the new subscriptions updates might occur
	<-ctx.Done()
	sm.remove(s)

	return nil
}

type streamFilter func() []bson.D

func (sm *SubscriptionsMux) ServeSubscriptionRollingWindow(ctx context.Context, gt GenericType, c domain.MongoCollection, filter streamFilter) error {
	f := filter()
	s := &subscription{
		gt:          gt,
		mutex:       sync.Mutex{},
		documentIds: []string{},
		filter:      f,
		ctx:         ctx,
	}

	if gtWithHooks, ok := gt.(GenericTypeWithHooks); ok {
		// forward hooks to subscription
		s.HookBeforeResponse = gtWithHooks.HookBeforeResponse
	}

	// add subscription to SubscriptionMux
	sm.add(s)

	// todo make sure already received documents don't get sent with an INSERT operation -> use UPDATE
	err := s.fetchDocumentsAndServe(sm.c, f)
	if err != nil {
		return err
	}

	// enable polling queries.
	ticker := time.NewTicker(PollingInterval)
	quit := make(chan struct{})

	defer close(quit)
	go func() {
		for {
			select {
			case <-ticker.C:
				err = s.fakeUpdates(sm.c, filter())
			case <-quit:
				log.Println("stopped polling.")
				ticker.Stop()
				return
			}
		}
	}()

	// todo
	// between sendingInitialDocuments and registering the new subscriptions updates might occur
	<-ctx.Done()
	sm.remove(s)

	return nil
}

func (sm *SubscriptionsMux) notifySubscribers(wc *watchChange) {
	sm.ops <- func(subs *[]*subscription) {
		// fmt.Printf("notifying %v subscriptions %v\n", len(*subs), wc)

		for _, s := range *subs {
			// todo sync iteration.
			// all routines should finish before we return from this function
			//
			// might make sense to make current loop serialized (remove next line routine) until we implement sync.
			go s.handleDocumentChange(wc, sm.c)
		}
	}
}

// SubscriptionsMux handles subscription management and sending updates
type SubscriptionsMux struct {
	ops chan func(*[]*subscription)
	c   domain.MongoCollection
}

// NewSubscriptionsMux starts a collection watch, instantiate a new subscription muxer which handles watchChanges, subscription management and sending updates
func NewSubscriptionsMux(c domain.MongoCollection) *SubscriptionsMux {
	mongoOperations, err := collectionWatch(c)
	if err != nil {
		log.Fatal(err)
	}

	m := &SubscriptionsMux{
		ops: make(chan func(*[]*subscription)),
		c:   c,
	}
	go m.loop()
	go func() {
		for wc := range mongoOperations {
			m.notifySubscribers(wc)
		}
		// todo maybe send an informative message to the subscribers.
	}()
	go m.logSubscriptionsData()

	return m
}
