package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// Db structure facilitates access to the database
// and opened connection to the passed uri.
type Db struct {
	con          *mongo.Client
	conOpLog     *mongo.Client
	conSecondary *mongo.Client

	// Database is the currently connected database with collections
	Database *mongo.Database
}

// Init opens and verifies connection to database
// returns struct that holds connection and database.
func Init(uri string, cdx ...func(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder) (*Db, error) {
	ctx := context.Background()
	log.Println("Connecting to MongoDB")

	d := &Db{}
	var err error

	reg := bson.NewRegistryBuilder()

	// Register custom codecs for protobuf Timestamp and wrapper types
	reg = codecs.Register(reg)

	// register any custom codecs sent by the client
	for _, register := range cdx {
		reg = register(reg)
	}

	clientOpts := options.Client().ApplyURI(uri).SetRegistry(reg.Build())
	clientOpts.SetMonitor(otelmongo.NewMonitor())

	d.con, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to Mongo %#v", err)
	}

	err = d.con.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot verify MongoDB connection %#v", err)
	}

	// extract database name from uri
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse MongoDB URI '%s': %#v", uri, err)
	}

	name := strings.TrimPrefix(u.Path, "/")
	if name == "" {
		return nil, fmt.Errorf("Missing database name from MongoDB URI '%s'", uri)
	}

	log.Println("Connected to MongoDB, selecting database", name)
	d.Database = d.con.Database(name)

	return d, nil
}

// Close closes the database connection
func (d *Db) Close() {
	log.Println("Disconnecting from MongoDB")
	d.con.Disconnect(context.Background())
}
