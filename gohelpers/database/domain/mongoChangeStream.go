package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// ensure MongoChangeStream struct implements MongoChangeStream interface
var _ MongoChangeStream = &mongoChangeStream{}

type MongoChangeStream interface {
	// createOperationDeployment(server driver.Server, connection driver.Connection) driver.Deployment
	// executeOperation(ctx context.Context, resuming bool) error
	// updatePbrtFromCommand()
	// storeResumeToken() error
	// buildPipelineSlice(pipeline interface) error
	// createPipelineOptionsDoc() bsoncore.Document
	// pipelineToBSON() (bsoncore.Document, error)
	// replaceOptions(ctx context.Context, wireVersion *description.VersionRange)
	// ID() int64
	Decode(val interface{}) error
	Err() error
	Close(ctx context.Context) error
	// ResumeToken() bson.Raw
	Next(ctx context.Context) bool
	// TryNext(ctx context.Context) bool
	// next(ctx context.Context, nonBlocking bool) bool
	// loopNext(ctx context.Context, nonBlocking bool)
	// isResumableError() bool
	// emptyBatch() bool
}

// newMongoChangeStreamToInternal used by the mongoCollection in order to wrap the returned cursors from the mongo.Collection functions.
func newMongoChangeStreamToInternal(cs *mongo.ChangeStream) MongoChangeStream {
	return &mongoChangeStream{cs: cs}
}

type mongoChangeStream struct {
	cs *mongo.ChangeStream
}

func (mcs *mongoChangeStream) Decode(val interface{}) error {
	return mcs.cs.Decode(val)
}

func (mcs *mongoChangeStream) Err() error {
	return mcs.cs.Err()
}

func (mcs *mongoChangeStream) Close(ctx context.Context) error {
	return mcs.cs.Close(ctx)
}

func (mcs *mongoChangeStream) Next(ctx context.Context) bool {
	return mcs.cs.Next(ctx)
}
