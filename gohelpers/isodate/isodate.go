package isodate

import (
	"math"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimeToGrpcTime converts time.Time object into protobuf timestamp
func TimeToGrpcTime(t ...time.Time) *timestamppb.Timestamp {
	var originalT time.Time

	if len(t) == 0 {
		originalT = time.Now()
	} else if len(t) == 1 {
		originalT = t[0]
	}

	return timestamppb.New(originalT)
}

// FloatToGrpcTime converts float timestamp to proto timestamp
func FloatToGrpcTime(f float64) *timestamppb.Timestamp {
	sec, dec := math.Modf(f)
	t := time.Unix(int64(sec), int64(dec*(1e9)))

	return timestamppb.New(t)
}

// FloatToGrpcDuration converts float duration to proto duration
func FloatToGrpcDuration(f float64) *durationpb.Duration {
	sec, dec := math.Modf(f)
	t := time.Duration(int64(sec))*time.Second + time.Duration(int64(dec*(1e9)))

	return durationpb.New(t)
}
