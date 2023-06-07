package parser

import (
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ParsePbTimestamp converts protobuf timestamp to model timestamp
func ParsePbTimestamp(ts *timestamppb.Timestamp) *model.Timestamp {
	if ts == nil {
		return nil
	}

	nanos := int(ts.Nanos)
	seconds := int(ts.Seconds)

	return &model.Timestamp{
		Nanos:   &nanos,
		Seconds: &seconds,
	}
}

// AssemblePbTimestamp converts model timestamp to protobuf timestamp
func AssemblePbTimestamp(ts *model.Timestamp) *timestamppb.Timestamp {
	if ts == nil {
		return nil
	}

	t := &timestamppb.Timestamp{}
	if ts.Seconds != nil {
		t.Seconds = int64(*ts.Seconds)
	}
	if ts.Nanos != nil {
		t.Nanos = int32(*ts.Nanos)
	}

	return t
}
