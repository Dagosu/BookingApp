package persistence

import (
	"context"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/gohelpers/database"
	dbDomain "github.com/Dagosu/BookingApp/gohelpers/database/domain"
	"gitlab.com/airportlabs/gohelpers/random"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userFlightsMappingRepository struct {
	c dbDomain.MongoCollection
}

var _ domain.UserFlightsMappingRepository = &userFlightsMappingRepository{}

func newUserFlightsMappingRepository(d *database.Db) *userFlightsMappingRepository {
	const usersFlightsMappingCollectionName string = "userFlightsMapping"
	c := dbDomain.NewMongoCollection(d.Database, usersFlightsMappingCollectionName)

	return &userFlightsMappingRepository{
		c: c,
	}
}

func (ufmr *userFlightsMappingRepository) Create(ctx context.Context, user *dt.User, flight *dt.Flight) (*dt.UserFlightsMapping, error) {
	purchasedFlights := []*dt.Flight{flight}
	userFlights := &dt.UserFlightsMapping{
		Id:               random.RandStringBytes(12),
		User:             user,
		PurchasedFlights: purchasedFlights,
	}

	_, err := ufmr.c.InsertOne(ctx, &userFlights)
	if err != nil {
		return nil, err
	}

	return userFlights, nil
}

func (ufmr *userFlightsMappingRepository) GetByUser(ctx context.Context, userId string) (*dt.UserFlightsMapping, error) {
	userFlights := &dt.UserFlightsMapping{}

	err := ufmr.c.FindOne(ctx, bson.D{{Key: "user._id", Value: userId}}).Decode(&userFlights)
	if err != nil {
		return nil, err
	}

	return userFlights, nil
}

func (ufmr *userFlightsMappingRepository) AddPurchasedFlight(ctx context.Context, userFlightsId string, flight *dt.Flight) error {
	filter := bson.M{"_id": userFlightsId}
	update := bson.M{"$push": bson.M{"purchased_flights": flight}}

	_, err := ufmr.c.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (ufmr *userFlightsMappingRepository) AddFavoritedFlight(ctx context.Context, userFlightsId string, flight *dt.Flight) error {
	filter := bson.M{"_id": userFlightsId}
	update := bson.M{"$push": bson.M{"favorited_flights": flight}}

	_, err := ufmr.c.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (ufmr *userFlightsMappingRepository) GetPurchasedFlights(ctx context.Context, userId string) ([]*dt.Flight, error) {
	var flights []*dt.Flight

	pipeline := mongo.Pipeline{
		bson.D{{
			Key: "$match", Value: bson.D{{"user._id", userId}},
		}},
	}
	qr, err := ufmr.c.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	err = qr.All(ctx, &flights)

	return flights, err
}
