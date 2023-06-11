// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// Query
type CheckCredentialsInput struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type CheckCredentialsResponse struct {
	Authorized *bool `json:"authorized,omitempty"`
}

type FavoriteFlightInput struct {
	UserID   *string `json:"userId,omitempty"`
	FlightID *string `json:"flightId,omitempty"`
}

type FavoriteFlightResponse struct {
	FavoritedFlight *Flight `json:"favoritedFlight,omitempty"`
}

// Input
type FilterParamInput struct {
	Condition *string `json:"condition,omitempty"`
	Field     *string `json:"field,omitempty"`
	Operator  *string `json:"operator,omitempty"`
	Value     *string `json:"value,omitempty"`
}

type Flight struct {
	ID            string     `json:"id"`
	Departure     *string    `json:"departure,omitempty"`
	DepartureTime *Timestamp `json:"departureTime,omitempty"`
	Arrival       *string    `json:"arrival,omitempty"`
	ArrivalTime   *Timestamp `json:"arrivalTime,omitempty"`
	BookableSeats *int       `json:"bookableSeats,omitempty"`
}

// Subscription
type FlightListInput struct {
	Limit  *int                `json:"limit,omitempty"`
	Offset *int                `json:"offset,omitempty"`
	Query  *string             `json:"query,omitempty"`
	Sorts  []*SortParamInput   `json:"sorts,omitempty"`
	Filter []*FilterParamInput `json:"filter,omitempty"`
}

type FlightListResponse struct {
	OperationType *OperationType `json:"operationType,omitempty"`
	Flights       []*Flight      `json:"flights,omitempty"`
}

// Subscription
type PurchaseFlightInput struct {
	UserID   *string `json:"userId,omitempty"`
	FlightID *string `json:"flightId,omitempty"`
}

type PurchaseFlightResponse struct {
	PurchasedFlight *Flight `json:"purchasedFlight,omitempty"`
}

type SortParamInput struct {
	Field *string        `json:"field,omitempty"`
	Order *ViewSortOrder `json:"order,omitempty"`
}

// Type
type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds *int `json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos *int `json:"nanos,omitempty"`
}

// Objects
type OperationType string

const (
	OperationTypeUnknownUnspecified OperationType = "UNKNOWN_UNSPECIFIED"
	OperationTypeInsert             OperationType = "INSERT"
	OperationTypeReplace            OperationType = "REPLACE"
	OperationTypeUpdate             OperationType = "UPDATE"
	OperationTypeDelete             OperationType = "DELETE"
	// initial data send completed, but server remains connected
	// so it can send further updates
	OperationTypeReady OperationType = "READY"
	// on client display the `error` field
	OperationTypeError OperationType = "ERROR"
	// used to notify client about a long operation status
	OperationTypeProgress OperationType = "PROGRESS"
	// operation was finished, server should disconnect after sending this
	OperationTypeFinished OperationType = "FINISHED"
)

var AllOperationType = []OperationType{
	OperationTypeUnknownUnspecified,
	OperationTypeInsert,
	OperationTypeReplace,
	OperationTypeUpdate,
	OperationTypeDelete,
	OperationTypeReady,
	OperationTypeError,
	OperationTypeProgress,
	OperationTypeFinished,
}

func (e OperationType) IsValid() bool {
	switch e {
	case OperationTypeUnknownUnspecified, OperationTypeInsert, OperationTypeReplace, OperationTypeUpdate, OperationTypeDelete, OperationTypeReady, OperationTypeError, OperationTypeProgress, OperationTypeFinished:
		return true
	}
	return false
}

func (e OperationType) String() string {
	return string(e)
}

func (e *OperationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OperationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OperationType", str)
	}
	return nil
}

func (e OperationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ViewSortOrder string

const (
	ViewSortOrderViewSortOrderUnspecified ViewSortOrder = "VIEW_SORT_ORDER_UNSPECIFIED"
	ViewSortOrderAsc                      ViewSortOrder = "ASC"
	ViewSortOrderDesc                     ViewSortOrder = "DESC"
)

var AllViewSortOrder = []ViewSortOrder{
	ViewSortOrderViewSortOrderUnspecified,
	ViewSortOrderAsc,
	ViewSortOrderDesc,
}

func (e ViewSortOrder) IsValid() bool {
	switch e {
	case ViewSortOrderViewSortOrderUnspecified, ViewSortOrderAsc, ViewSortOrderDesc:
		return true
	}
	return false
}

func (e ViewSortOrder) String() string {
	return string(e)
}

func (e *ViewSortOrder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ViewSortOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ViewSortOrder", str)
	}
	return nil
}

func (e ViewSortOrder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
