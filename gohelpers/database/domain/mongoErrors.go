package domain

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// List of error codes: https://github.com/mongodb/mongo/blob/master/src/mongo/base/error_codes.yml
const duplicateKeyCode = 11000

func IsDuplicateKeyError(err error) bool {
	if e, ok := err.(mongo.WriteException); ok {
		for _, we := range e.WriteErrors {
			if we.Code == duplicateKeyCode {
				return true
			}
		}
	}

	return false
}
