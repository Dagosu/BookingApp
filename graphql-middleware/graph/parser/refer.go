package parser

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// BoolDerefer safely converts a *bool to bool
func BoolDerefer(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}

// BoolRefer converts bool to *bool
func BoolRefer(b bool) *bool {
	return &b
}

// IntReferWrapped converts wrapped Int32Value to *int
func IntReferWrapped(i *wrapperspb.Int32Value) *int {
	if i == nil {
		return nil
	}
	v := int(i.Value)

	return &v
}

// IntDereferWrapped converts *int to wrapped Int32Value
func IntDereferWrapped(i *int) *wrapperspb.Int32Value {
	if i == nil {
		return nil
	}

	v := int32(*i)
	return &wrapperspb.Int32Value{
		Value: v,
	}
}

// IntDerefer safely converts *int to int32
func IntDerefer(i *int) int32 {
	if i == nil {
		return 0
	}

	return int32(*i)
}

// IntRefer converts int32 to *int
func IntRefer(i int32) *int {
	int32Nr := int(i)

	return &int32Nr
}

// StrRefer converts string to *string
func StrRefer(s string) *string {
	return &s
}

// StrDerefer safely converts *string to string
func StrDerefer(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

// StrReferWrapped converts wrapped value to *string
func StrReferWrapped(s *wrapperspb.StringValue) *string {
	if s == nil {
		return nil
	}
	return &s.Value
}

// StrDereferWrapped converts *string to wrapped string value
func StrDereferWrapped(s *string) *wrapperspb.StringValue {
	if s == nil {
		return nil
	}
	return &wrapperspb.StringValue{
		Value: *s,
	}
}

// boolReferWrapped converts wrapped value to *bool
func boolReferWrapped(b *wrapperspb.BoolValue) *bool {
	if b == nil {
		return nil
	}
	return &b.Value
}

// boolDereferWrapped converts *bool to *wrappers.BoolValue
func boolDereferWrapped(b *bool) *wrapperspb.BoolValue {
	if b == nil {
		return nil
	}
	return &wrapperspb.BoolValue{
		Value: *b,
	}
}

// BoolDereferWrapped converts *bool to *wrappers.BoolValue
func BoolDereferWrapped(b *bool) *wrapperspb.BoolValue {
	if b == nil {
		return nil
	}
	return &wrapperspb.BoolValue{
		Value: *b,
	}
}

// BoolReferWrapped converts wrapped value to *bool
func BoolReferWrapped(b *wrapperspb.BoolValue) *bool {
	if b == nil {
		return nil
	}
	return &b.Value
}

// FloatRefer safely converts float64 to *float64
func FloatRefer(i float64) *float64 {
	return &i
}

// FloatDerefer safely converts *float64 to float64
func FloatDerefer(i *float64) float64 {
	if i == nil {
		return 0
	}

	return *i
}
