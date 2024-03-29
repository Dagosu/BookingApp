// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: flight.proto

package datatypes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ViewSortOrder int32

const (
	ViewSortOrder_VIEW_SORT_ORDER_UNSPECIFIED ViewSortOrder = 0
	ViewSortOrder_ASC                         ViewSortOrder = 1
	ViewSortOrder_DESC                        ViewSortOrder = 2
)

// Enum value maps for ViewSortOrder.
var (
	ViewSortOrder_name = map[int32]string{
		0: "VIEW_SORT_ORDER_UNSPECIFIED",
		1: "ASC",
		2: "DESC",
	}
	ViewSortOrder_value = map[string]int32{
		"VIEW_SORT_ORDER_UNSPECIFIED": 0,
		"ASC":                         1,
		"DESC":                        2,
	}
)

func (x ViewSortOrder) Enum() *ViewSortOrder {
	p := new(ViewSortOrder)
	*p = x
	return p
}

func (x ViewSortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ViewSortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_flight_proto_enumTypes[0].Descriptor()
}

func (ViewSortOrder) Type() protoreflect.EnumType {
	return &file_flight_proto_enumTypes[0]
}

func (x ViewSortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ViewSortOrder.Descriptor instead.
func (ViewSortOrder) EnumDescriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{0}
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty" bson:"user_name,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty" bson:"text,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Review) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"_id,omitempty" bson:"_id,omitempty"`
	Departure     string                 `protobuf:"bytes,2,opt,name=departure,proto3" json:"departure,omitempty" bson:"departure,omitempty"`
	DepartureTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=departure_time,json=departureTime,proto3" json:"departure_time,omitempty" bson:"departure_time,omitempty"`
	Arrival       string                 `protobuf:"bytes,4,opt,name=arrival,proto3" json:"arrival,omitempty" bson:"arrival,omitempty"`
	ArrivalTime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=arrival_time,json=arrivalTime,proto3" json:"arrival_time,omitempty" bson:"arrival_time,omitempty"`
	TotalSeats    int32                  `protobuf:"varint,6,opt,name=total_seats,json=totalSeats,proto3" json:"total_seats,omitempty" bson:"total_seats,omitempty"`
	BookableSeats int32                  `protobuf:"varint,7,opt,name=bookable_seats,json=bookableSeats,proto3" json:"bookable_seats,omitempty" bson:"bookable_seats,omitempty"`
	Airline       string                 `protobuf:"bytes,8,opt,name=airline,proto3" json:"airline,omitempty" bson:"airline,omitempty"`
	Price         float32                `protobuf:"fixed32,9,opt,name=price,proto3" json:"price,omitempty" bson:"price,omitempty"`
	Status        string                 `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty" bson:"status,omitempty"`
	Reviews       []*Review              `protobuf:"bytes,11,rep,name=reviews,proto3" json:"reviews,omitempty" bson:"reviews,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{1}
}

func (x *Flight) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Flight) GetDeparture() string {
	if x != nil {
		return x.Departure
	}
	return ""
}

func (x *Flight) GetDepartureTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DepartureTime
	}
	return nil
}

func (x *Flight) GetArrival() string {
	if x != nil {
		return x.Arrival
	}
	return ""
}

func (x *Flight) GetArrivalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ArrivalTime
	}
	return nil
}

func (x *Flight) GetTotalSeats() int32 {
	if x != nil {
		return x.TotalSeats
	}
	return 0
}

func (x *Flight) GetBookableSeats() int32 {
	if x != nil {
		return x.BookableSeats
	}
	return 0
}

func (x *Flight) GetAirline() string {
	if x != nil {
		return x.Airline
	}
	return ""
}

func (x *Flight) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Flight) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Flight) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type FlightListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// limit is used to only show a specified number of documents
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty" bson:"limit,omitempty"`
	// offset is used to skip a specified number of documents
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty" bson:"offset,omitempty"`
	// query is the query parameter used for free text search
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty" bson:"query,omitempty"`
	// sorts suports sorting by multiple SortParam parameters
	Sorts []*SortParam `protobuf:"bytes,4,rep,name=sorts,proto3" json:"sorts,omitempty" bson:"sorts,omitempty"`
	// filter support filtering by multiple FilterParam parameters
	Filter []*FilterParam `protobuf:"bytes,5,rep,name=filter,proto3" json:"filter,omitempty" bson:"filter,omitempty"`
}

func (x *FlightListRequest) Reset() {
	*x = FlightListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightListRequest) ProtoMessage() {}

func (x *FlightListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightListRequest.ProtoReflect.Descriptor instead.
func (*FlightListRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{2}
}

func (x *FlightListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FlightListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FlightListRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *FlightListRequest) GetSorts() []*SortParam {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *FlightListRequest) GetFilter() []*FilterParam {
	if x != nil {
		return x.Filter
	}
	return nil
}

type FlightListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flight  *Flight   `protobuf:"bytes,1,opt,name=flight,proto3" json:"flight,omitempty" bson:"flight,omitempty"`
	Flights []*Flight `protobuf:"bytes,2,rep,name=flights,proto3" json:"flights,omitempty" bson:"flights,omitempty"`
	// operation_type for current message
	OperationType OperationType `protobuf:"varint,3,opt,name=operation_type,json=operationType,proto3,enum=operation.OperationType" json:"operation_type,omitempty" bson:"operation_type,omitempty"`
}

func (x *FlightListResponse) Reset() {
	*x = FlightListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightListResponse) ProtoMessage() {}

func (x *FlightListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightListResponse.ProtoReflect.Descriptor instead.
func (*FlightListResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{3}
}

func (x *FlightListResponse) GetFlight() *Flight {
	if x != nil {
		return x.Flight
	}
	return nil
}

func (x *FlightListResponse) GetFlights() []*Flight {
	if x != nil {
		return x.Flights
	}
	return nil
}

func (x *FlightListResponse) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_UNKNOWN_UNSPECIFIED
}

// SortParam represents the sort query
type SortParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// field is the field path we want to sort
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty" bson:"field,omitempty"`
	// order is the order direction we want to sort the field
	Order ViewSortOrder `protobuf:"varint,2,opt,name=order,proto3,enum=flight.ViewSortOrder" json:"order,omitempty" bson:"order,omitempty"`
}

func (x *SortParam) Reset() {
	*x = SortParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortParam) ProtoMessage() {}

func (x *SortParam) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortParam.ProtoReflect.Descriptor instead.
func (*SortParam) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{4}
}

func (x *SortParam) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SortParam) GetOrder() ViewSortOrder {
	if x != nil {
		return x.Order
	}
	return ViewSortOrder_VIEW_SORT_ORDER_UNSPECIFIED
}

// FilterParam represents the filter query
type FilterParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// condition is the type of filter (and, or)
	Condition string `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty" bson:"condition,omitempty"`
	// field is the field path we want to filter
	Field string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty" bson:"field,omitempty"`
	// operator is the type of filter (accepted values: eq, ne, lt, gt, lte,
	// gte, contains, begins, ends, pastX, nextX, relativePastX, relativeNextX)
	Operator string `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty" bson:"operator,omitempty"`
	// value is the value of the field we want to search against
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty" bson:"value,omitempty"`
}

func (x *FilterParam) Reset() {
	*x = FilterParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterParam) ProtoMessage() {}

func (x *FilterParam) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterParam.ProtoReflect.Descriptor instead.
func (*FilterParam) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{5}
}

func (x *FilterParam) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *FilterParam) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *FilterParam) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *FilterParam) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightId string `protobuf:"bytes,1,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty" bson:"flight_id,omitempty"`
}

func (x *GetFlightRequest) Reset() {
	*x = GetFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightRequest) ProtoMessage() {}

func (x *GetFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightRequest.ProtoReflect.Descriptor instead.
func (*GetFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{6}
}

func (x *GetFlightRequest) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

type GetFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flight *Flight `protobuf:"bytes,1,opt,name=flight,proto3" json:"flight,omitempty" bson:"flight,omitempty"`
}

func (x *GetFlightResponse) Reset() {
	*x = GetFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightResponse) ProtoMessage() {}

func (x *GetFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightResponse.ProtoReflect.Descriptor instead.
func (*GetFlightResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{7}
}

func (x *GetFlightResponse) GetFlight() *Flight {
	if x != nil {
		return x.Flight
	}
	return nil
}

type WriteReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightId string `protobuf:"bytes,1,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty" bson:"flight_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id,omitempty"`
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty" bson:"text,omitempty"`
}

func (x *WriteReviewRequest) Reset() {
	*x = WriteReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteReviewRequest) ProtoMessage() {}

func (x *WriteReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteReviewRequest.ProtoReflect.Descriptor instead.
func (*WriteReviewRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{8}
}

func (x *WriteReviewRequest) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

func (x *WriteReviewRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *WriteReviewRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type WriteReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flight *Flight `protobuf:"bytes,1,opt,name=flight,proto3" json:"flight,omitempty" bson:"flight,omitempty"`
}

func (x *WriteReviewResponse) Reset() {
	*x = WriteReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteReviewResponse) ProtoMessage() {}

func (x *WriteReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteReviewResponse.ProtoReflect.Descriptor instead.
func (*WriteReviewResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{9}
}

func (x *WriteReviewResponse) GetFlight() *Flight {
	if x != nil {
		return x.Flight
	}
	return nil
}

var File_flight_proto protoreflect.FileDescriptor

var file_flight_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x8c, 0x03, 0x0a, 0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x12, 0x41, 0x0a, 0x0e,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x72, 0x72,
	0x69, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x61, 0x72, 0x72,
	0x69, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6f, 0x6f,
	0x6b, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x74, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x11, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x05,
	0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05,
	0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0xa7, 0x01, 0x0a, 0x12, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x28, 0x0a, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4e, 0x0a, 0x09,
	0x53, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x2b, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x6f, 0x72, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x73, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x49, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x5e, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x3d, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2a, 0x43,
	0x0a, 0x0d, 0x56, 0x69, 0x65, 0x77, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x1b, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53,
	0x43, 0x10, 0x02, 0x32, 0xe0, 0x01, 0x0a, 0x0d, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x2e, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1a, 0x2e,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x67, 0x6f, 0x73, 0x75, 0x2f, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flight_proto_rawDescOnce sync.Once
	file_flight_proto_rawDescData = file_flight_proto_rawDesc
)

func file_flight_proto_rawDescGZIP() []byte {
	file_flight_proto_rawDescOnce.Do(func() {
		file_flight_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_proto_rawDescData)
	})
	return file_flight_proto_rawDescData
}

var file_flight_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flight_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_flight_proto_goTypes = []interface{}{
	(ViewSortOrder)(0),            // 0: flight.ViewSortOrder
	(*Review)(nil),                // 1: flight.Review
	(*Flight)(nil),                // 2: flight.Flight
	(*FlightListRequest)(nil),     // 3: flight.FlightListRequest
	(*FlightListResponse)(nil),    // 4: flight.FlightListResponse
	(*SortParam)(nil),             // 5: flight.SortParam
	(*FilterParam)(nil),           // 6: flight.FilterParam
	(*GetFlightRequest)(nil),      // 7: flight.GetFlightRequest
	(*GetFlightResponse)(nil),     // 8: flight.GetFlightResponse
	(*WriteReviewRequest)(nil),    // 9: flight.WriteReviewRequest
	(*WriteReviewResponse)(nil),   // 10: flight.WriteReviewResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(OperationType)(0),            // 12: operation.OperationType
}
var file_flight_proto_depIdxs = []int32{
	11, // 0: flight.Flight.departure_time:type_name -> google.protobuf.Timestamp
	11, // 1: flight.Flight.arrival_time:type_name -> google.protobuf.Timestamp
	1,  // 2: flight.Flight.reviews:type_name -> flight.Review
	5,  // 3: flight.FlightListRequest.sorts:type_name -> flight.SortParam
	6,  // 4: flight.FlightListRequest.filter:type_name -> flight.FilterParam
	2,  // 5: flight.FlightListResponse.flight:type_name -> flight.Flight
	2,  // 6: flight.FlightListResponse.flights:type_name -> flight.Flight
	12, // 7: flight.FlightListResponse.operation_type:type_name -> operation.OperationType
	0,  // 8: flight.SortParam.order:type_name -> flight.ViewSortOrder
	2,  // 9: flight.GetFlightResponse.flight:type_name -> flight.Flight
	2,  // 10: flight.WriteReviewResponse.flight:type_name -> flight.Flight
	3,  // 11: flight.FlightService.FlightList:input_type -> flight.FlightListRequest
	7,  // 12: flight.FlightService.GetFlight:input_type -> flight.GetFlightRequest
	9,  // 13: flight.FlightService.WriteReview:input_type -> flight.WriteReviewRequest
	4,  // 14: flight.FlightService.FlightList:output_type -> flight.FlightListResponse
	8,  // 15: flight.FlightService.GetFlight:output_type -> flight.GetFlightResponse
	10, // 16: flight.FlightService.WriteReview:output_type -> flight.WriteReviewResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_flight_proto_init() }
func file_flight_proto_init() {
	if File_flight_proto != nil {
		return
	}
	file_operation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteReviewRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flight_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteReviewResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flight_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flight_proto_goTypes,
		DependencyIndexes: file_flight_proto_depIdxs,
		EnumInfos:         file_flight_proto_enumTypes,
		MessageInfos:      file_flight_proto_msgTypes,
	}.Build()
	File_flight_proto = out.File
	file_flight_proto_rawDesc = nil
	file_flight_proto_goTypes = nil
	file_flight_proto_depIdxs = nil
}
