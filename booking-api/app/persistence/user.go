package persistence

import (
	"context"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/gohelpers/database"
	dbDomain "github.com/Dagosu/BookingApp/gohelpers/database/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// userRepository is the struct that implements the domain user repository
type userRepository struct {
	c dbDomain.MongoCollection
}

// UserRepository implements the domain.UserRepository interface
var _ domain.UserRepository = &userRepository{}

// newUserRepository instantiates a new userRepository and returns it as a domain.UserRepository
func newUserRepository(d *database.Db) *userRepository {
	const usersCollectionName string = "users"
	c := dbDomain.NewMongoCollection(d.Database, usersCollectionName)

	return &userRepository{
		c: c,
	}
}

func (ur *userRepository) CheckCredentials(ctx context.Context, email, password string) (*dt.User, error) {
	user := &dt.User{}
	filter := bson.D{{"email", email}, {"password", password}}
	err := ur.c.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
