// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ozon_keyword_api

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

// OzonKeywordApiServiceClient is the client API for OzonKeywordApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OzonKeywordApiServiceClient interface {
	// DescribeKeywordV1 - Describe a keyword
	DescribeKeywordV1(ctx context.Context, in *DescribeKeywordV1Request, opts ...grpc.CallOption) (*DescribeKeywordV1Response, error)
	CreateKeywordV1(ctx context.Context, in *CreateKeywordV1Request, opts ...grpc.CallOption) (*CreateKeywordV1Response, error)
	ListKeywordV1(ctx context.Context, in *ListKeywordV1Request, opts ...grpc.CallOption) (*ListKeywordV1Response, error)
	RemoveKeywordV1(ctx context.Context, in *RemoveKeywordV1Request, opts ...grpc.CallOption) (*RemoveKeywordV1Response, error)
}

type ozonKeywordApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOzonKeywordApiServiceClient(cc grpc.ClientConnInterface) OzonKeywordApiServiceClient {
	return &ozonKeywordApiServiceClient{cc}
}

func (c *ozonKeywordApiServiceClient) DescribeKeywordV1(ctx context.Context, in *DescribeKeywordV1Request, opts ...grpc.CallOption) (*DescribeKeywordV1Response, error) {
	out := new(DescribeKeywordV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.ozon_keyword_api.v1.OzonKeywordApiService/DescribeKeywordV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonKeywordApiServiceClient) CreateKeywordV1(ctx context.Context, in *CreateKeywordV1Request, opts ...grpc.CallOption) (*CreateKeywordV1Response, error) {
	out := new(CreateKeywordV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.ozon_keyword_api.v1.OzonKeywordApiService/CreateKeywordV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonKeywordApiServiceClient) ListKeywordV1(ctx context.Context, in *ListKeywordV1Request, opts ...grpc.CallOption) (*ListKeywordV1Response, error) {
	out := new(ListKeywordV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.ozon_keyword_api.v1.OzonKeywordApiService/ListKeywordV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonKeywordApiServiceClient) RemoveKeywordV1(ctx context.Context, in *RemoveKeywordV1Request, opts ...grpc.CallOption) (*RemoveKeywordV1Response, error) {
	out := new(RemoveKeywordV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.ozon_keyword_api.v1.OzonKeywordApiService/RemoveKeywordV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OzonKeywordApiServiceServer is the server API for OzonKeywordApiService service.
// All implementations must embed UnimplementedOzonKeywordApiServiceServer
// for forward compatibility
type OzonKeywordApiServiceServer interface {
	// DescribeKeywordV1 - Describe a keyword
	DescribeKeywordV1(context.Context, *DescribeKeywordV1Request) (*DescribeKeywordV1Response, error)
	CreateKeywordV1(context.Context, *CreateKeywordV1Request) (*CreateKeywordV1Response, error)
	ListKeywordV1(context.Context, *ListKeywordV1Request) (*ListKeywordV1Response, error)
	RemoveKeywordV1(context.Context, *RemoveKeywordV1Request) (*RemoveKeywordV1Response, error)
	mustEmbedUnimplementedOzonKeywordApiServiceServer()
}

// UnimplementedOzonKeywordApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOzonKeywordApiServiceServer struct {
}

func (UnimplementedOzonKeywordApiServiceServer) DescribeKeywordV1(context.Context, *DescribeKeywordV1Request) (*DescribeKeywordV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeKeywordV1 not implemented")
}
func (UnimplementedOzonKeywordApiServiceServer) CreateKeywordV1(context.Context, *CreateKeywordV1Request) (*CreateKeywordV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeywordV1 not implemented")
}
func (UnimplementedOzonKeywordApiServiceServer) ListKeywordV1(context.Context, *ListKeywordV1Request) (*ListKeywordV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeywordV1 not implemented")
}
func (UnimplementedOzonKeywordApiServiceServer) RemoveKeywordV1(context.Context, *RemoveKeywordV1Request) (*RemoveKeywordV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveKeywordV1 not implemented")
}
func (UnimplementedOzonKeywordApiServiceServer) mustEmbedUnimplementedOzonKeywordApiServiceServer() {}

// UnsafeOzonKeywordApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OzonKeywordApiServiceServer will
// result in compilation errors.
type UnsafeOzonKeywordApiServiceServer interface {
	mustEmbedUnimplementedOzonKeywordApiServiceServer()
}

func RegisterOzonKeywordApiServiceServer(s grpc.ServiceRegistrar, srv OzonKeywordApiServiceServer) {
	s.RegisterService(&OzonKeywordApiService_ServiceDesc, srv)
}

func _OzonKeywordApiService_DescribeKeywordV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeKeywordV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonKeywordApiServiceServer).DescribeKeywordV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.ozon_keyword_api.v1.OzonKeywordApiService/DescribeKeywordV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonKeywordApiServiceServer).DescribeKeywordV1(ctx, req.(*DescribeKeywordV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OzonKeywordApiService_CreateKeywordV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKeywordV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonKeywordApiServiceServer).CreateKeywordV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.ozon_keyword_api.v1.OzonKeywordApiService/CreateKeywordV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonKeywordApiServiceServer).CreateKeywordV1(ctx, req.(*CreateKeywordV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OzonKeywordApiService_ListKeywordV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeywordV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonKeywordApiServiceServer).ListKeywordV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.ozon_keyword_api.v1.OzonKeywordApiService/ListKeywordV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonKeywordApiServiceServer).ListKeywordV1(ctx, req.(*ListKeywordV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OzonKeywordApiService_RemoveKeywordV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveKeywordV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonKeywordApiServiceServer).RemoveKeywordV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.ozon_keyword_api.v1.OzonKeywordApiService/RemoveKeywordV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonKeywordApiServiceServer).RemoveKeywordV1(ctx, req.(*RemoveKeywordV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OzonKeywordApiService_ServiceDesc is the grpc.ServiceDesc for OzonKeywordApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OzonKeywordApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.ozon_keyword_api.v1.OzonKeywordApiService",
	HandlerType: (*OzonKeywordApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeKeywordV1",
			Handler:    _OzonKeywordApiService_DescribeKeywordV1_Handler,
		},
		{
			MethodName: "CreateKeywordV1",
			Handler:    _OzonKeywordApiService_CreateKeywordV1_Handler,
		},
		{
			MethodName: "ListKeywordV1",
			Handler:    _OzonKeywordApiService_ListKeywordV1_Handler,
		},
		{
			MethodName: "RemoveKeywordV1",
			Handler:    _OzonKeywordApiService_RemoveKeywordV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/ozon_keyword_api/v1/ozon_keyword_api.proto",
}
