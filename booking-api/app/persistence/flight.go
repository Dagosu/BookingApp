package persistence

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Dagosu/BookingApp/booking-api/app/service"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/gohelpers/database"
	dbDomain "github.com/Dagosu/BookingApp/gohelpers/database/domain"
	"github.com/Dagosu/BookingApp/gohelpers/fielddescriptor"
	"github.com/Dagosu/BookingApp/gohelpers/isodate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type flightRepository struct {
	c  dbDomain.MongoCollection
	sm *database.SubscriptionsMux
}

type flightGenericType struct {
	stream dt.FlightService_FlightListServer
}

var (
	_ database.GenericType = &flightGenericType{}

	relativeTimeRegex = regexp.MustCompile(`(\S\d+)([a-zA-Z]{2})`)
	testNoQuery       = regexp.MustCompile(`^[0-9]?[a-zA-Z]{1,2}[0-9]?\s[0-9]{1,4}\s*[a-zA-Z]?$`)
)

func (ft *flightGenericType) New(id ...string) interface{ GetId() string } {
	f := &dt.Flight{}

	if len(id) > 0 {
		f.Id = id[0]
	}

	return f
}

func (ft *flightGenericType) SendResponse(data interface{}, operationType dt.OperationType) error {
	r := &dt.FlightListResponse{
		OperationType: operationType,
	}

	if data != nil {
		r.Flight = data.(*dt.Flight)
	}

	return ft.stream.Send(r)
}

// newFlightRepository instantiates a new flightRepository and returns it as a domain.FlightRepository
func newFlightRepository(d *database.Db) *flightRepository {
	const flightsCollectionName string = "flights"
	c := dbDomain.NewMongoCollection(d.Database, flightsCollectionName)

	return &flightRepository{
		c:  c,
		sm: database.NewSubscriptionsMux(c),
	}
}

func (fr *flightRepository) Get(ctx context.Context, id string) (*dt.Flight, error) {
	flight := &dt.Flight{}

	err := fr.c.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&flight)
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (fr *flightRepository) GetFutureFlights(ctx context.Context) ([]*dt.Flight, error) {
	var flights []*dt.Flight
	currentTime := isodate.TimeToGrpcTime(time.Now())

	pipeline := mongo.Pipeline{
		bson.D{{
			Key: "$match", Value: bson.D{{
				Key:   "departure_time",
				Value: bson.M{"$gt": currentTime},
			}},
		}},
	}
	qr, err := fr.c.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	err = qr.All(ctx, &flights)

	return flights, err
}

func (fr *flightRepository) FlightList(req *dt.FlightListRequest, stream dt.FlightService_FlightListServer) error {
	mogt := &flightGenericType{
		stream: stream,
	}

	return fr.sm.ServeSubscription(stream.Context(), mogt, fr.c, fr.getFilterPipeline(req, []string{}))
}

func (fr *flightRepository) getFilterPipeline(req *dt.FlightListRequest, fields []string) mongo.Pipeline {
	pipeline := mongo.Pipeline{
		bson.D{{
			Key: "$match", Value: fr.createMatchConditions(req),
		}},
	}

	if len(req.GetSorts()) > 0 {
		sort := bson.D{}
		for _, s := range req.GetSorts() {
			sort = append(sort, getSortStatement(s))
		}
		pipeline = append(pipeline, bson.D{{Key: "$sort", Value: sort}})
	}

	pipeline = append(pipeline, bson.D{{Key: "$skip", Value: req.Offset}})
	if req.GetLimit() > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$limit", Value: req.Limit}})
	}

	return pipeline
}

func (tr *flightRepository) createMatchConditions(req *dt.FlightListRequest) primitive.D {
	matchConditions := bson.D{}

	if req.GetQuery() != "" {
		var search bson.E
		query := req.GetQuery()
		regex := primitive.Regex{Pattern: query, Options: "im"}

		search = bson.E{Key: "$or", Value: bson.A{
			bson.D{{Key: "departure", Value: regex}},
			bson.D{{Key: "arrival", Value: regex}},
			bson.D{{Key: "bookable_seats", Value: regex}},
		}}

		matchConditions = append(matchConditions, search)
	}

	if len(req.GetFilter()) == 0 {
		return matchConditions
	}

	filterConds := getFilterConditions(req.GetFilter())
	matchConditions = append(matchConditions, filterConds...)

	return matchConditions
}

func getFilterConditions(filters []*dt.FilterParam) primitive.D {
	conditions := bson.D{}
	orConditions := bson.A{}
	for _, f := range filters {
		f.Field = fielddescriptor.NormalizePath(f.GetField())
		if f.GetCondition() == "and" {
			if filter := getFilterStatement(f); len(filter.Key) > 0 {
				conditions = append(conditions, filter)
			}

			continue
		}
		if f.GetCondition() == "or" {
			if filter := getFilterStatement(f); len(filter.Key) > 0 {
				orConditions = append(orConditions, bson.D{filter})
			}
		}
	}
	past, next := addRelativeFilter(filters, database.PollingInterval)
	if len(past.Key) > 0 {
		conditions = append(conditions, past)
	}
	if len(next.Key) > 0 {
		conditions = append(conditions, next)
	}

	if len(orConditions) == 1 && len(conditions) == 1 {
		orConditions = append(orConditions, conditions)

		return bson.D{
			bson.E{Key: "$or", Value: orConditions},
		}
	}

	if len(orConditions) > 0 {
		conditions = append(conditions, bson.E{Key: "$or", Value: orConditions})
	}

	return conditions
}

func getFilterStatement(f *dt.FilterParam) primitive.E {
	fs := bson.E{}
	switch f.GetOperator() {
	case "eq", "ne", "lt", "gt", "lte", "gte":
		fs = getComparisonStatement(f)
	case "exists":
		fs = getExistsComparisonStatement(f)
	case "contains":
		fs = bson.E{Key: f.GetField(), Value: primitive.Regex{Pattern: `.*` + f.GetValue() + `.*`, Options: "is"}}
	case "begins":
		fs = bson.E{Key: f.GetField(), Value: primitive.Regex{Pattern: `^` + f.GetValue(), Options: "is"}}
	case "ends":
		fs = bson.E{Key: f.GetField(), Value: primitive.Regex{Pattern: f.GetValue() + `$`, Options: "is"}}
	case "pastX":
		ts, err := strconv.ParseInt(f.GetValue(), 10, 64)
		if err != nil {
			return fs
		}
		fs = bson.E{
			Key:   f.GetField(),
			Value: bson.D{{Key: "$gte", Value: isodate.TimeToGrpcTime(time.Unix(ts, 0))}, {Key: "$lte", Value: isodate.TimeToGrpcTime(time.Now())}},
		}
	case "nextX":
		ts, err := strconv.ParseInt(f.GetValue(), 10, 64)
		if err != nil {
			return fs
		}
		fs = bson.E{
			Key:   f.GetField(),
			Value: bson.D{{Key: "$gte", Value: isodate.TimeToGrpcTime(time.Now())}, {Key: "$lte", Value: isodate.TimeToGrpcTime(time.Unix(ts, 0))}},
		}
	case "in":
		inValues := strings.Split(f.GetValue(), ",")
		if len(inValues) == 0 {
			return fs
		}
		fs = bson.E{Key: f.GetField(), Value: bson.D{{Key: "$in", Value: inValues}}}
	}

	return fs
}

func getComparisonStatement(f *dt.FilterParam) primitive.E {
	field := f.GetField()
	operator := f.GetOperator()
	value := f.GetValue()
	if operator == "eq" && (strings.HasSuffix(field, "_id") || strings.HasSuffix(field, "alrn")) {
		return bson.E{Key: field, Value: value}
	}
	intVal, err := strconv.ParseInt(value, 10, 64)
	if service.IsTimeField(field) {
		return getTimeComparisonStatement(f, intVal)
	}
	if err == nil {
		return getIntComparisonStatement(f, intVal)
	}
	floatVal, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return bson.E{Key: field, Value: bson.D{{Key: "$" + operator, Value: floatVal}}}
	}

	if strings.EqualFold("true", value) || strings.EqualFold("false", value) { // "T" should not be intepreted as bool (see ParseBool doc)
		boolVal, err := strconv.ParseBool(value)
		if err == nil {
			return bson.E{Key: field, Value: bson.D{{Key: "$" + operator, Value: boolVal}}}
		}
	}

	if operator == "eq" {
		return bson.E{Key: field, Value: primitive.Regex{Pattern: `^` + value + `$`, Options: "is"}}
	}

	if operator == "ne" {
		return bson.E{Key: field, Value: bson.D{{Key: "$not", Value: primitive.Regex{Pattern: `.*` + value + `.*`, Options: "is"}}}}
	}

	return bson.E{Key: field, Value: bson.D{{Key: "$" + operator, Value: value}}}
}

func getExistsComparisonStatement(f *dt.FilterParam) bson.E {
	val, err := strconv.ParseBool(f.GetValue())
	if err != nil {
		return bson.E{}
	}

	return bson.E{Key: f.GetField(), Value: bson.D{{Key: "$exists", Value: val}}}
}

func getSortStatement(s *dt.SortParam) bson.E {
	order := 1
	if s.Order == dt.ViewSortOrder_DESC {
		order = -1
	}

	return bson.E{Key: s.Field, Value: order}
}

func addRelativeFilter(filters []*dt.FilterParam, pollingInterval time.Duration) (primitive.E, primitive.E) {
	var pastVal, nextVal, pastField, nextField string
	var foundPast, foundNext bool
	for _, f := range filters {
		if f.GetOperator() == "relativePastX" && !foundPast {
			pastVal = f.GetValue()
			pastField = f.GetField()
			foundPast = true

			continue
		}
		if f.GetOperator() == "relativeNextX" && !foundNext {
			nextVal = f.GetValue()
			nextField = f.GetField()
			foundNext = true
		}
	}

	if foundPast && foundNext {
		if pastField != nextField {
			return bson.E{
					Key: nextField,
					Value: bson.D{
						{Key: "$lte", Value: isodate.TimeToGrpcTime(getRelativeTime(nextVal, pollingInterval))},
					},
				},
				bson.E{
					Key: pastField,
					Value: bson.D{
						{Key: "$gte", Value: isodate.TimeToGrpcTime(getRelativeTime(pastVal, pollingInterval))},
					},
				}
		}

		return bson.E{
			Key: pastField,
			Value: bson.D{
				{Key: "$gte", Value: isodate.TimeToGrpcTime(getRelativeTime(pastVal, pollingInterval))},
				{Key: "$lte", Value: isodate.TimeToGrpcTime(getRelativeTime(nextVal, pollingInterval))},
			},
		}, bson.E{}
	}

	if foundPast {
		return bson.E{}, bson.E{
			Key: pastField,
			Value: bson.D{
				{Key: "$gte", Value: isodate.TimeToGrpcTime(getRelativeTime(pastVal, pollingInterval))},
				{Key: "$lte", Value: isodate.TimeToGrpcTime(time.Now())},
			},
		}
	}

	if foundNext {
		return bson.E{
			Key: nextField,
			Value: bson.D{
				{Key: "$gte", Value: isodate.TimeToGrpcTime(time.Now())},
				{Key: "$lte", Value: isodate.TimeToGrpcTime(getRelativeTime(nextVal, pollingInterval))},
			},
		}, bson.E{}
	}

	return bson.E{}, bson.E{}
}

func getRelativeTime(value string, pollingInterval time.Duration) time.Time {
	var (
		duration  = "0"
		timeframe = ""
	)
	tf := relativeTimeRegex.FindStringSubmatch(value)
	if len(tf) == 3 {
		duration = tf[1]
		timeframe = tf[2]
	}
	now := time.Now()

	switch timeframe {
	case "mm":
		minDuration, err := time.ParseDuration(duration + "m")
		if err != nil {
			return now
		}
		return now.Add(minDuration).Add(pollingInterval)
	case "HH":
		hourDuration, err := time.ParseDuration(duration + "h")
		if err != nil {
			return now
		}
		return now.Add(hourDuration).Add(pollingInterval)
	case "dd":
		dayDuration, err := time.ParseDuration(duration + "h")
		if err != nil {
			return now
		}
		return now.Add(dayDuration * 24).Add(pollingInterval)
	}

	return now
}

func getIntComparisonStatement(f *dt.FilterParam, intVal int64) bson.E {
	if intVal == 0 {
		query := getPseudoIntZeroValueQuery(f, intVal)
		if query != nil {
			return *query
		}
	}

	switch f.GetOperator() {
	case "eq", "ne":
		return getPseudoIntQuery(f, intVal)
	case "gte", "gt":
		// consider zero value when looking for a number greater than a negative number
		if intVal < 0 {
			return getNegativeIntQuery(f, intVal)
		}
	case "lte", "lt":
		// consider zero value when looking for a number less then a positive number
		if intVal > 0 {
			return getPositiveIntQuery(f, intVal)
		}
	}

	return bson.E{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: intVal}}}
}

func getPseudoIntQuery(f *dt.FilterParam, intVal int64) bson.E {
	// when querying using $eq and $ne and the filter value can be intepreted as an int
	// there are situations when it is indeed an int, but there are situation when it is a string
	// so, we call that "pseudo-int" and we query in db both for int and for string values
	// as we are not sure how they are saved in db
	if f.GetOperator() == "eq" {
		return getEqPseudoIntQuery(f, intVal)
	}

	return getNePsudoIntQuery(f, intVal)
}

func getEqPseudoIntQuery(f *dt.FilterParam, intVal int64) bson.E {
	// if the value is a string in db, however its form is numeric and was interpreted as a number
	return bson.E{Key: "$or", Value: bson.A{
		bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: intVal}}}},
		bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: f.GetValue()}}}},
	}}
}

func getNePsudoIntQuery(f *dt.FilterParam, intVal int64) bson.E {
	// if the value is a string in db, check if one of the following conditions is true
	// the number is different than 0
	andCondition := bson.D{{
		Key: "$and",
		Value: bson.A{
			bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: intVal}}}},
			bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: f.GetValue()}}}},
		},
	}}

	return bson.E{
		Key: "$or",
		Value: bson.A{
			andCondition,
			bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$exists", Value: false}}}},
		},
	}
}

func getPseudoIntZeroValueQuery(f *dt.FilterParam, intVal int64) *bson.E {
	// consider 0 value when querying for 0 value
	// also, verify both for string and int
	operator := f.GetOperator()
	stringQuery := bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: f.GetValue()}}}}
	intQuery := bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: intVal}}}}
	existQuery := bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$exists", Value: true}}}}
	notExistQuery := bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$exists", Value: false}}}}
	switch operator {
	case "eq":
		// Is string and we should use simple $eq operator
		// Is int and we should verify for zero value (exists:true)
		return &bson.E{
			Key: "$or",
			Value: bson.A{
				intQuery,      // equal to int
				stringQuery,   // equal to string
				notExistQuery, // not exists
			},
		}
	case "gte", "lte":
		return &bson.E{Key: "$or", Value: bson.A{
			intQuery,      // greater or equal or lesser than int
			notExistQuery, // not exists
		}}
	case "ne":
		return &bson.E{
			Key: "$and",
			Value: bson.A{
				intQuery,    // not equal to int
				stringQuery, // not equal to string
				existQuery,  // exists
			},
		}
	}

	return nil
}

func getTimeComparisonStatement(f *dt.FilterParam, intVal int64) bson.E {
	t := time.Unix(intVal, 0)

	if f.GetOperator() == "eq" {
		tStart := t.Add(-time.Duration(t.Second()) * time.Second)
		tEnd := tStart.Add(59 * time.Second)

		return bson.E{Key: "$and", Value: bson.A{
			bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$gte", Value: tStart}}}},
			bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$lte", Value: tEnd}}}},
		}}
	}

	return bson.E{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: isodate.TimeToGrpcTime(t)}}}
}

func getNegativeIntQuery(f *dt.FilterParam, intVal int64) bson.E {
	return bson.E{Key: "$or", Value: bson.A{
		bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: intVal}}}},
		bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$exists", Value: false}}}},
	}}
}

func getPositiveIntQuery(f *dt.FilterParam, intVal int64) bson.E {
	return bson.E{Key: "$or", Value: bson.A{
		bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$" + f.GetOperator(), Value: intVal}}}},
		bson.D{{Key: f.GetField(), Value: bson.D{{Key: "$exists", Value: false}}}},
	}}
}
