package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ensure mongoCollection struct implements MongoCollection interface
var _ MongoCollection = &mongoCollection{}

type MongoCollection interface {
	// Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error)
	Name() string
	// Database() *mongo.Database
	// BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	// ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	// altered mongo.Cursor -> MongoCursor
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (MongoCursor, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	// EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)
	// Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error)
	// altered mongo.Cursor -> MongoCursor
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (MongoCursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) MongoSingleResult
	// FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult
	// FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	// altered mongo.ChangeStream -> MongoChangeStream
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (MongoChangeStream, error)
	// Indexes() mongo.IndexView
	Drop(ctx context.Context) error
}

type mongoCollection struct {
	c *mongo.Collection
}

func NewMongoCollection(d *mongo.Database, name string) MongoCollection {
	c := d.Collection(name)

	return &mongoCollection{c: c}
}

func (mc *mongoCollection) Name() string {
	return mc.c.Name()
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return mc.c.InsertOne(ctx, document, opts...)
}

func (mc *mongoCollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return mc.c.InsertMany(ctx, documents, opts...)
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return mc.c.DeleteOne(ctx, filter, opts...)
}

func (mc *mongoCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return mc.c.DeleteMany(ctx, filter, opts...)
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.c.UpdateOne(ctx, filter, update, opts...)
}

func (mc *mongoCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.c.UpdateMany(ctx, filter, update, opts...)
}

func (mc *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (MongoCursor, error) {
	c, err := mc.c.Aggregate(ctx, pipeline, opts...)

	return newMongoCursorToInternal(c), err
}

func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return mc.c.CountDocuments(ctx, filter, opts...)
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (MongoCursor, error) {
	c, err := mc.c.Find(ctx, filter, opts...)

	return newMongoCursorToInternal(c), err
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) MongoSingleResult {
	return mc.c.FindOne(ctx, filter, opts...)
}

func (mc *mongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return mc.c.FindOneAndUpdate(ctx, filter, update, opts...)
}

func (mc *mongoCollection) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (MongoChangeStream, error) {
	mcs, err := mc.c.Watch(ctx, pipeline, opts...)

	return newMongoChangeStreamToInternal(mcs), err
}

func (mc *mongoCollection) Drop(ctx context.Context) error {
	return mc.c.Drop(ctx)
}
