package persistence

import (
	"context"

	"github.com/Dagosu/BookingApp/gohelpers/database"
	dbDomain "github.com/Dagosu/BookingApp/gohelpers/database/domain"
	"github.com/k0kubun/pp"
)

type testRepository struct {
	c dbDomain.MongoCollection
}

// newTestRepository instantiates a new testRepository and returns it as a domain.TestRepository
func newTestRepository(d *database.Db) *testRepository {
	const testCollectionName string = "test"
	c := dbDomain.NewMongoCollection(d.Database, testCollectionName)

	return &testRepository{
		c: c,
	}
}

func (tr *testRepository) TestEndpoint(ctx context.Context, request string) (string, error) {
	pp.Println("testtest")

	request = request + " YES "

	return request, nil
}
