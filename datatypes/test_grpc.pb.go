// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: test.proto

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

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	TestEndpoint(ctx context.Context, in *TestEndpointRequest, opts ...grpc.CallOption) (*TestEndpointResponse, error)
	TestList(ctx context.Context, in *TestListRequest, opts ...grpc.CallOption) (TestService_TestListClient, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestEndpoint(ctx context.Context, in *TestEndpointRequest, opts ...grpc.CallOption) (*TestEndpointResponse, error) {
	out := new(TestEndpointResponse)
	err := c.cc.Invoke(ctx, "/test.TestService/TestEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestList(ctx context.Context, in *TestListRequest, opts ...grpc.CallOption) (TestService_TestListClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], "/test.TestService/TestList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTestListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_TestListClient interface {
	Recv() (*TestListResponse, error)
	grpc.ClientStream
}

type testServiceTestListClient struct {
	grpc.ClientStream
}

func (x *testServiceTestListClient) Recv() (*TestListResponse, error) {
	m := new(TestListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations should embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	TestEndpoint(context.Context, *TestEndpointRequest) (*TestEndpointResponse, error)
	TestList(*TestListRequest, TestService_TestListServer) error
}

// UnimplementedTestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) TestEndpoint(context.Context, *TestEndpointRequest) (*TestEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestEndpoint not implemented")
}
func (UnimplementedTestServiceServer) TestList(*TestListRequest, TestService_TestListServer) error {
	return status.Errorf(codes.Unimplemented, "method TestList not implemented")
}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_TestEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/TestEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestEndpoint(ctx, req.(*TestEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).TestList(m, &testServiceTestListServer{stream})
}

type TestService_TestListServer interface {
	Send(*TestListResponse) error
	grpc.ServerStream
}

type testServiceTestListServer struct {
	grpc.ServerStream
}

func (x *testServiceTestListServer) Send(m *TestListResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestEndpoint",
			Handler:    _TestService_TestEndpoint_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestList",
			Handler:       _TestService_TestList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}
