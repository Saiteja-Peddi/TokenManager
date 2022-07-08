// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tokenmanagerpb

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

// TokenManagerClient is the client API for TokenManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenManagerClient interface {
	Create(ctx context.Context, in *NormalRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	Read(ctx context.Context, in *NormalRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	Drop(ctx context.Context, in *NormalRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	CopyToken(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*ServerResponse, error)
}

type tokenManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenManagerClient(cc grpc.ClientConnInterface) TokenManagerClient {
	return &tokenManagerClient{cc}
}

func (c *tokenManagerClient) Create(ctx context.Context, in *NormalRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/tokenmanager.TokenManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/tokenmanager.TokenManager/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerClient) Read(ctx context.Context, in *NormalRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/tokenmanager.TokenManager/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerClient) Drop(ctx context.Context, in *NormalRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/tokenmanager.TokenManager/Drop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenManagerClient) CopyToken(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/tokenmanager.TokenManager/CopyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenManagerServer is the server API for TokenManager service.
// All implementations must embed UnimplementedTokenManagerServer
// for forward compatibility
type TokenManagerServer interface {
	Create(context.Context, *NormalRequest) (*ServerResponse, error)
	Write(context.Context, *WriteRequest) (*ServerResponse, error)
	Read(context.Context, *NormalRequest) (*ServerResponse, error)
	Drop(context.Context, *NormalRequest) (*ServerResponse, error)
	CopyToken(context.Context, *CopyRequest) (*ServerResponse, error)
	mustEmbedUnimplementedTokenManagerServer()
}

// UnimplementedTokenManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTokenManagerServer struct {
}

func (UnimplementedTokenManagerServer) Create(context.Context, *NormalRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTokenManagerServer) Write(context.Context, *WriteRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedTokenManagerServer) Read(context.Context, *NormalRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedTokenManagerServer) Drop(context.Context, *NormalRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Drop not implemented")
}
func (UnimplementedTokenManagerServer) CopyToken(context.Context, *CopyRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyToken not implemented")
}
func (UnimplementedTokenManagerServer) mustEmbedUnimplementedTokenManagerServer() {}

// UnsafeTokenManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenManagerServer will
// result in compilation errors.
type UnsafeTokenManagerServer interface {
	mustEmbedUnimplementedTokenManagerServer()
}

func RegisterTokenManagerServer(s grpc.ServiceRegistrar, srv TokenManagerServer) {
	s.RegisterService(&TokenManager_ServiceDesc, srv)
}

func _TokenManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NormalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmanager.TokenManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServer).Create(ctx, req.(*NormalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManager_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmanager.TokenManager/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManager_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NormalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmanager.TokenManager/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServer).Read(ctx, req.(*NormalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManager_Drop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NormalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServer).Drop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmanager.TokenManager/Drop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServer).Drop(ctx, req.(*NormalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenManager_CopyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenManagerServer).CopyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmanager.TokenManager/CopyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenManagerServer).CopyToken(ctx, req.(*CopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenManager_ServiceDesc is the grpc.ServiceDesc for TokenManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenmanager.TokenManager",
	HandlerType: (*TokenManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TokenManager_Create_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _TokenManager_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _TokenManager_Read_Handler,
		},
		{
			MethodName: "Drop",
			Handler:    _TokenManager_Drop_Handler,
		},
		{
			MethodName: "CopyToken",
			Handler:    _TokenManager_CopyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}
