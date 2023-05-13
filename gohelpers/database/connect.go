package database

// convert go timestamps to mongo isodate
// todo: watch otelmongo. At the moment (v0.16.0) the latest module doesn't work. using 0.14.0
import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	// https://medium.com/@amsokol.com/new-official-mongodb-go-driver-and-google-protobuf-making-them-work-together-6357b0118f3f
	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// Db structure facilitates access to the database
// and opened connection to the passed uri.
type Db struct {
	con          *mongo.Client
	conOpLog     *mongo.Client
	conSecondary *mongo.Client

	// Database is the currently connected database with collections
	Database          *mongo.Database
	DatabaseSubs      *mongo.Database
	DatabaseSecondary *mongo.Database
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
	clientOpts.SetMonitor(otelmongo.NewMonitor("mongodb"))

	d.con, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to Mongo %#v", err)
	}

	// init another connection without otel
	clientOptsOpLog := options.Client().ApplyURI(uri).SetRegistry(reg.Build())
	d.conOpLog, err = mongo.Connect(ctx, clientOptsOpLog)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to Mongo for subscription watch %#v", err)
	}

	// init another connection for secondary mongos (for heavy queries)
	// https://www.mongodb.com/docs/manual/core/read-preference/#mongodb-readmode-secondaryPreferred
	clientOptsSecondary := options.Client().ApplyURI(uri).SetRegistry(reg.Build())
	clientOptsSecondary.SetReadPreference(readpref.SecondaryPreferred())
	d.conSecondary, err = mongo.Connect(ctx, clientOptsSecondary)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to Mongo for secondary read: %#v", err)
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
	d.DatabaseSubs = d.conOpLog.Database(name)
	d.DatabaseSecondary = d.conSecondary.Database(name)

	return d, nil
}

// Close closes the database connection
func (d *Db) Close() {
	log.Println("Disconnecting from MongoDB")
	d.con.Disconnect(context.Background())
	d.conOpLog.Disconnect(context.Background())
	d.conSecondary.Disconnect(context.Background())
}

// HealthCheck verifies that the connection to the mongo database is active (responds to mongo ping).
func (d *Db) HealthCheck(ctx context.Context) error {
	return d.con.Ping(ctx, nil)
}
