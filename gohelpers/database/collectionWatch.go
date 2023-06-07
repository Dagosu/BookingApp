package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	dt "github.com/Dagosu/BookingApp/datatypes"
	dbDomain "github.com/Dagosu/BookingApp/gohelpers/database/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// watchChange holds information about the a document change on a given mongo collection
type watchChange struct {
	// ID represents the document id
	ID string
	// OperatioType represents the executed operation
	OperationType dt.OperationType
}

type docKey struct {
	ID string `bson:"_id,omitempty"`
}

type doc struct {
	OperationType string `bson:"operationType,omitempty"`
	DocumentKey   docKey `bson:"documentKey,omitempty" json:"document_key,omitempty"`
}

var (
	mongoUpdate = "update"
	mongoDelete = "delete"
	mongoInsert = "insert"
)

func getWatchChange(d doc) (*watchChange, error) {
	operationDict := map[string]dt.OperationType{
		mongoUpdate: dt.OperationType_UPDATE,
		mongoDelete: dt.OperationType_DELETE,
		mongoInsert: dt.OperationType_INSERT,
	}

	if _, ok := operationDict[d.OperationType]; !ok {
		return nil, fmt.Errorf("Operation not implemented %v ", d.OperationType)
	}
	if d.DocumentKey.ID == "" {
		return nil, errors.New("Missing id from mongo watch doc, ignoring for now")
	}

	return &watchChange{
		ID:            d.DocumentKey.ID,
		OperationType: operationDict[d.OperationType],
	}, nil
}

// collectionWatch watches provided collection and emits WatchChanges on the returned channel.
// logs and skips unknown mongo operations and decoding errors.
func collectionWatch(c dbDomain.MongoCollection) (<-chan *watchChange, error) {
	ctx := context.Background()
	mcs, err := c.Watch(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	mongoOperations := make(chan *watchChange)
	log.Printf("watching collection for updates\n")
	// log.Printf("watching %v collection for updates\n", collectionName)
	go func() {
		defer mcs.Close(ctx)
		defer close(mongoOperations)

		d := &doc{}
		for mcs.Next(ctx) {
			err := mcs.Decode(d)
			if err != nil {
				// fmt.Printf("Skipping change, decoding error: %v\n", err.Error())
				continue
			}
			wc, err := getWatchChange(*d)
			if err != nil {
				// fmt.Printf("Skipping change getWatchChange error, : %v\n", err.Error())
				continue
			}

			mongoOperations <- wc
		}

		log.Printf("Mongo change stream error :%v\n", mcs.Err())
	}()

	return mongoOperations, nil
}
