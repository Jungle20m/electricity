// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: internal/grpcserver/protoc/service.proto

package service

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

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerClient interface {
	GetCustomerID(ctx context.Context, in *GetCustomerIDRequest, opts ...grpc.CallOption) (*GetCustomerIDResponse, error)
	GetWorkerSiteID(ctx context.Context, in *GetWorkerSiteIDRequest, opts ...grpc.CallOption) (*GetWorkerSiteIDResponse, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetCustomerID(ctx context.Context, in *GetCustomerIDRequest, opts ...grpc.CallOption) (*GetCustomerIDResponse, error) {
	out := new(GetCustomerIDResponse)
	err := c.cc.Invoke(ctx, "/service.Customer/GetCustomerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetWorkerSiteID(ctx context.Context, in *GetWorkerSiteIDRequest, opts ...grpc.CallOption) (*GetWorkerSiteIDResponse, error) {
	out := new(GetWorkerSiteIDResponse)
	err := c.cc.Invoke(ctx, "/service.Customer/GetWorkerSiteID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
// All implementations must embed UnimplementedCustomerServer
// for forward compatibility
type CustomerServer interface {
	GetCustomerID(context.Context, *GetCustomerIDRequest) (*GetCustomerIDResponse, error)
	GetWorkerSiteID(context.Context, *GetWorkerSiteIDRequest) (*GetWorkerSiteIDResponse, error)
	mustEmbedUnimplementedCustomerServer()
}

// UnimplementedCustomerServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (UnimplementedCustomerServer) GetCustomerID(context.Context, *GetCustomerIDRequest) (*GetCustomerIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerID not implemented")
}
func (UnimplementedCustomerServer) GetWorkerSiteID(context.Context, *GetWorkerSiteIDRequest) (*GetWorkerSiteIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerSiteID not implemented")
}
func (UnimplementedCustomerServer) mustEmbedUnimplementedCustomerServer() {}

// UnsafeCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServer will
// result in compilation errors.
type UnsafeCustomerServer interface {
	mustEmbedUnimplementedCustomerServer()
}

func RegisterCustomerServer(s grpc.ServiceRegistrar, srv CustomerServer) {
	s.RegisterService(&Customer_ServiceDesc, srv)
}

func _Customer_GetCustomerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetCustomerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Customer/GetCustomerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetCustomerID(ctx, req.(*GetCustomerIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetWorkerSiteID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkerSiteIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetWorkerSiteID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Customer/GetWorkerSiteID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetWorkerSiteID(ctx, req.(*GetWorkerSiteIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_ServiceDesc is the grpc.ServiceDesc for Customer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerID",
			Handler:    _Customer_GetCustomerID_Handler,
		},
		{
			MethodName: "GetWorkerSiteID",
			Handler:    _Customer_GetWorkerSiteID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/grpcserver/protoc/service.proto",
}
