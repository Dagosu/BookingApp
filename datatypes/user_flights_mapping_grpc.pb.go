// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: user_flights_mapping.proto

package datatypes

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserFlightsMappingServiceClient is the client API for UserFlightsMappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserFlightsMappingServiceClient interface {
	PurchaseFlight(ctx context.Context, in *PurchaseFlightRequest, opts ...grpc.CallOption) (*PurchaseFlightResponse, error)
	FavoriteFlight(ctx context.Context, in *FavoriteFlightRequest, opts ...grpc.CallOption) (*FavoriteFlightResponse, error)
	GetPurchasedFlights(ctx context.Context, in *GetPurchasedFlightsRequest, opts ...grpc.CallOption) (*GetPurchasedFlightsResponse, error)
	GetFavoritedFlights(ctx context.Context, in *GetFavoritedFlightsRequest, opts ...grpc.CallOption) (*GetFavoritedFlightsResponse, error)
	RecommendFlight(ctx context.Context, in *RecommendFlightRequest, opts ...grpc.CallOption) (*RecommendFlightResponse, error)
}

type userFlightsMappingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserFlightsMappingServiceClient(cc grpc.ClientConnInterface) UserFlightsMappingServiceClient {
	return &userFlightsMappingServiceClient{cc}
}

func (c *userFlightsMappingServiceClient) PurchaseFlight(ctx context.Context, in *PurchaseFlightRequest, opts ...grpc.CallOption) (*PurchaseFlightResponse, error) {
	out := new(PurchaseFlightResponse)
	err := c.cc.Invoke(ctx, "/user_flights_mapping.UserFlightsMappingService/PurchaseFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFlightsMappingServiceClient) FavoriteFlight(ctx context.Context, in *FavoriteFlightRequest, opts ...grpc.CallOption) (*FavoriteFlightResponse, error) {
	out := new(FavoriteFlightResponse)
	err := c.cc.Invoke(ctx, "/user_flights_mapping.UserFlightsMappingService/FavoriteFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFlightsMappingServiceClient) GetPurchasedFlights(ctx context.Context, in *GetPurchasedFlightsRequest, opts ...grpc.CallOption) (*GetPurchasedFlightsResponse, error) {
	out := new(GetPurchasedFlightsResponse)
	err := c.cc.Invoke(ctx, "/user_flights_mapping.UserFlightsMappingService/GetPurchasedFlights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFlightsMappingServiceClient) GetFavoritedFlights(ctx context.Context, in *GetFavoritedFlightsRequest, opts ...grpc.CallOption) (*GetFavoritedFlightsResponse, error) {
	out := new(GetFavoritedFlightsResponse)
	err := c.cc.Invoke(ctx, "/user_flights_mapping.UserFlightsMappingService/GetFavoritedFlights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFlightsMappingServiceClient) RecommendFlight(ctx context.Context, in *RecommendFlightRequest, opts ...grpc.CallOption) (*RecommendFlightResponse, error) {
	out := new(RecommendFlightResponse)
	err := c.cc.Invoke(ctx, "/user_flights_mapping.UserFlightsMappingService/RecommendFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserFlightsMappingServiceServer is the server API for UserFlightsMappingService service.
// All implementations should embed UnimplementedUserFlightsMappingServiceServer
// for forward compatibility
type UserFlightsMappingServiceServer interface {
	PurchaseFlight(context.Context, *PurchaseFlightRequest) (*PurchaseFlightResponse, error)
	FavoriteFlight(context.Context, *FavoriteFlightRequest) (*FavoriteFlightResponse, error)
	GetPurchasedFlights(context.Context, *GetPurchasedFlightsRequest) (*GetPurchasedFlightsResponse, error)
	GetFavoritedFlights(context.Context, *GetFavoritedFlightsRequest) (*GetFavoritedFlightsResponse, error)
	RecommendFlight(context.Context, *RecommendFlightRequest) (*RecommendFlightResponse, error)
}

// UnimplementedUserFlightsMappingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserFlightsMappingServiceServer struct {
}

func (UnimplementedUserFlightsMappingServiceServer) PurchaseFlight(context.Context, *PurchaseFlightRequest) (*PurchaseFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseFlight not implemented")
}
func (UnimplementedUserFlightsMappingServiceServer) FavoriteFlight(context.Context, *FavoriteFlightRequest) (*FavoriteFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteFlight not implemented")
}
func (UnimplementedUserFlightsMappingServiceServer) GetPurchasedFlights(context.Context, *GetPurchasedFlightsRequest) (*GetPurchasedFlightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPurchasedFlights not implemented")
}
func (UnimplementedUserFlightsMappingServiceServer) GetFavoritedFlights(context.Context, *GetFavoritedFlightsRequest) (*GetFavoritedFlightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoritedFlights not implemented")
}
func (UnimplementedUserFlightsMappingServiceServer) RecommendFlight(context.Context, *RecommendFlightRequest) (*RecommendFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendFlight not implemented")
}

// UnsafeUserFlightsMappingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserFlightsMappingServiceServer will
// result in compilation errors.
type UnsafeUserFlightsMappingServiceServer interface {
	mustEmbedUnimplementedUserFlightsMappingServiceServer()
}

func RegisterUserFlightsMappingServiceServer(s grpc.ServiceRegistrar, srv UserFlightsMappingServiceServer) {
	s.RegisterService(&UserFlightsMappingService_ServiceDesc, srv)
}

func _UserFlightsMappingService_PurchaseFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFlightsMappingServiceServer).PurchaseFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_flights_mapping.UserFlightsMappingService/PurchaseFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFlightsMappingServiceServer).PurchaseFlight(ctx, req.(*PurchaseFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFlightsMappingService_FavoriteFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFlightsMappingServiceServer).FavoriteFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_flights_mapping.UserFlightsMappingService/FavoriteFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFlightsMappingServiceServer).FavoriteFlight(ctx, req.(*FavoriteFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFlightsMappingService_GetPurchasedFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPurchasedFlightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFlightsMappingServiceServer).GetPurchasedFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_flights_mapping.UserFlightsMappingService/GetPurchasedFlights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFlightsMappingServiceServer).GetPurchasedFlights(ctx, req.(*GetPurchasedFlightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFlightsMappingService_GetFavoritedFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoritedFlightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFlightsMappingServiceServer).GetFavoritedFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_flights_mapping.UserFlightsMappingService/GetFavoritedFlights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFlightsMappingServiceServer).GetFavoritedFlights(ctx, req.(*GetFavoritedFlightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFlightsMappingService_RecommendFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFlightsMappingServiceServer).RecommendFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_flights_mapping.UserFlightsMappingService/RecommendFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFlightsMappingServiceServer).RecommendFlight(ctx, req.(*RecommendFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserFlightsMappingService_ServiceDesc is the grpc.ServiceDesc for UserFlightsMappingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserFlightsMappingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_flights_mapping.UserFlightsMappingService",
	HandlerType: (*UserFlightsMappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseFlight",
			Handler:    _UserFlightsMappingService_PurchaseFlight_Handler,
		},
		{
			MethodName: "FavoriteFlight",
			Handler:    _UserFlightsMappingService_FavoriteFlight_Handler,
		},
		{
			MethodName: "GetPurchasedFlights",
			Handler:    _UserFlightsMappingService_GetPurchasedFlights_Handler,
		},
		{
			MethodName: "GetFavoritedFlights",
			Handler:    _UserFlightsMappingService_GetFavoritedFlights_Handler,
		},
		{
			MethodName: "RecommendFlight",
			Handler:    _UserFlightsMappingService_RecommendFlight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_flights_mapping.proto",
}