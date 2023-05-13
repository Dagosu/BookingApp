package domain

import "go.mongodb.org/mongo-driver/mongo"

var _ MongoSingleResult = &singleResult{}

type MongoSingleResult interface {
	Decode(v interface{}) error
}

type singleResult struct {
	sr *mongo.SingleResult
}

func (msr *singleResult) Decode(val interface{}) error {
	return msr.sr.Decode(val)
}
