package persistence

import (
	"context"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/gohelpers/database"
	dbDomain "github.com/Dagosu/BookingApp/gohelpers/database/domain"
)

type userFlightMappingRepository struct {
	c dbDomain.MongoCollection
}

var _ domain.UserRepository = &userRepository{}

func newUserFlightsMappingRepository(d *database.Db) *userRepository {
	const usersFlightsMappingCollectionName string = "usersFlightsMapping"
	c := dbDomain.NewMongoCollection(d.Database, usersFlightsMappingCollectionName)

	return &userRepository{
		c: c,
	}
}

func (ur *userRepository) Create(ctx context.Context, email, password string) (*dt.User, error) {
	// user := &dt.User{}
	// filter := bson.D{{"email", email}, {"password", password}}
	// err := ur.c.FindOne(ctx, filter).Decode(&user)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
