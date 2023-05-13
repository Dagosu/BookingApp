package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ensure mongoCursor struct implements MongoCursor interface
var _ MongoCursor = &mongoCursor{}

type MongoCursor interface {
	// ID() int64
	Next(ctx context.Context) bool
	// TryNext(ctx context.Context) bool
	Decode(val interface{}) error
	Err() error
	Close(ctx context.Context) error
	All(ctx context.Context, results interface{}) error
	// RemainingBatchLength() int
	//
	// New method -> used for accesing public Current field from mongo.Cursor.
	Current() bson.Raw
}

// newMongoCursorToInternal used by the mongoCollection in order to wrap the returned cursors from the mongo.Collection functions.
func newMongoCursorToInternal(c *mongo.Cursor) MongoCursor {
	return &mongoCursor{c: c}
}

var _ MongoCursor = &mongoCursor{}

type mongoCursor struct {
	c *mongo.Cursor
}

func (mc *mongoCursor) ID() int64 {
	return mc.c.ID()
}

func (mc *mongoCursor) Next(ctx context.Context) bool {
	return mc.c.Next(ctx)
}

func (mc *mongoCursor) TryNext(ctx context.Context) bool {
	return mc.c.TryNext(ctx)
}

func (mc *mongoCursor) Decode(val interface{}) error {
	return mc.c.Decode(val)
}

func (mc *mongoCursor) Err() error {
	return mc.c.Err()
}

func (mc *mongoCursor) Close(ctx context.Context) error {
	return mc.c.Close(ctx)
}

func (mc *mongoCursor) All(ctx context.Context, results interface{}) error {
	return mc.c.All(ctx, results)
}

func (mc *mongoCursor) RemainingBatchLength() int {
	return mc.c.RemainingBatchLength()
}

func (mc *mongoCursor) Current() bson.Raw {
	return mc.c.Current
}
