// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pb/pb.proto

package pb

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

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	// Sends a login request.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/pb.Login/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
// All implementations must embed UnimplementedLoginServer
// for forward compatibility
type LoginServer interface {
	// Sends a login request.
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	mustEmbedUnimplementedLoginServer()
}

// UnimplementedLoginServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (UnimplementedLoginServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServer) mustEmbedUnimplementedLoginServer() {}

// UnsafeLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServer will
// result in compilation errors.
type UnsafeLoginServer interface {
	mustEmbedUnimplementedLoginServer()
}

func RegisterLoginServer(s grpc.ServiceRegistrar, srv LoginServer) {
	s.RegisterService(&Login_ServiceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Login_ServiceDesc is the grpc.ServiceDesc for Login service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Login_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Login_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}

// RegisterAreaClient is the client API for RegisterArea service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterAreaClient interface {
	// Sends a RegisterArea request.
	RegisterArea(ctx context.Context, in *RegisterAreaRequest, opts ...grpc.CallOption) (*RegisterAreaReply, error)
}

type registerAreaClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterAreaClient(cc grpc.ClientConnInterface) RegisterAreaClient {
	return &registerAreaClient{cc}
}

func (c *registerAreaClient) RegisterArea(ctx context.Context, in *RegisterAreaRequest, opts ...grpc.CallOption) (*RegisterAreaReply, error) {
	out := new(RegisterAreaReply)
	err := c.cc.Invoke(ctx, "/pb.RegisterArea/RegisterArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterAreaServer is the server API for RegisterArea service.
// All implementations must embed UnimplementedRegisterAreaServer
// for forward compatibility
type RegisterAreaServer interface {
	// Sends a RegisterArea request.
	RegisterArea(context.Context, *RegisterAreaRequest) (*RegisterAreaReply, error)
	mustEmbedUnimplementedRegisterAreaServer()
}

// UnimplementedRegisterAreaServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterAreaServer struct {
}

func (UnimplementedRegisterAreaServer) RegisterArea(context.Context, *RegisterAreaRequest) (*RegisterAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterArea not implemented")
}
func (UnimplementedRegisterAreaServer) mustEmbedUnimplementedRegisterAreaServer() {}

// UnsafeRegisterAreaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterAreaServer will
// result in compilation errors.
type UnsafeRegisterAreaServer interface {
	mustEmbedUnimplementedRegisterAreaServer()
}

func RegisterRegisterAreaServer(s grpc.ServiceRegistrar, srv RegisterAreaServer) {
	s.RegisterService(&RegisterArea_ServiceDesc, srv)
}

func _RegisterArea_RegisterArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterAreaServer).RegisterArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegisterArea/RegisterArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterAreaServer).RegisterArea(ctx, req.(*RegisterAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterArea_ServiceDesc is the grpc.ServiceDesc for RegisterArea service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegisterArea_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RegisterArea",
	HandlerType: (*RegisterAreaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterArea",
			Handler:    _RegisterArea_RegisterArea_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterServiceClient interface {
	// Sends a RegisterService request.
	RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceReply, error)
}

type registerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterServiceClient(cc grpc.ClientConnInterface) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceReply, error) {
	out := new(RegisterServiceReply)
	err := c.cc.Invoke(ctx, "/pb.RegisterService/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServiceServer is the server API for RegisterService service.
// All implementations must embed UnimplementedRegisterServiceServer
// for forward compatibility
type RegisterServiceServer interface {
	// Sends a RegisterService request.
	RegisterService(context.Context, *RegisterServiceRequest) (*RegisterServiceReply, error)
	mustEmbedUnimplementedRegisterServiceServer()
}

// UnimplementedRegisterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServiceServer struct {
}

func (UnimplementedRegisterServiceServer) RegisterService(context.Context, *RegisterServiceRequest) (*RegisterServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedRegisterServiceServer) mustEmbedUnimplementedRegisterServiceServer() {}

// UnsafeRegisterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServiceServer will
// result in compilation errors.
type UnsafeRegisterServiceServer interface {
	mustEmbedUnimplementedRegisterServiceServer()
}

func RegisterRegisterServiceServer(s grpc.ServiceRegistrar, srv RegisterServiceServer) {
	s.RegisterService(&RegisterService_ServiceDesc, srv)
}

func _RegisterService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegisterService/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).RegisterService(ctx, req.(*RegisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterService_ServiceDesc is the grpc.ServiceDesc for RegisterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegisterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _RegisterService_RegisterService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}

// DeleteAreaClient is the client API for DeleteArea service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeleteAreaClient interface {
	// Sends a DeleteArea request.
	DeleteArea(ctx context.Context, in *DeleteAreaRequest, opts ...grpc.CallOption) (*DeleteAreaReply, error)
}

type deleteAreaClient struct {
	cc grpc.ClientConnInterface
}

func NewDeleteAreaClient(cc grpc.ClientConnInterface) DeleteAreaClient {
	return &deleteAreaClient{cc}
}

func (c *deleteAreaClient) DeleteArea(ctx context.Context, in *DeleteAreaRequest, opts ...grpc.CallOption) (*DeleteAreaReply, error) {
	out := new(DeleteAreaReply)
	err := c.cc.Invoke(ctx, "/pb.DeleteArea/DeleteArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteAreaServer is the server API for DeleteArea service.
// All implementations must embed UnimplementedDeleteAreaServer
// for forward compatibility
type DeleteAreaServer interface {
	// Sends a DeleteArea request.
	DeleteArea(context.Context, *DeleteAreaRequest) (*DeleteAreaReply, error)
	mustEmbedUnimplementedDeleteAreaServer()
}

// UnimplementedDeleteAreaServer must be embedded to have forward compatible implementations.
type UnimplementedDeleteAreaServer struct {
}

func (UnimplementedDeleteAreaServer) DeleteArea(context.Context, *DeleteAreaRequest) (*DeleteAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArea not implemented")
}
func (UnimplementedDeleteAreaServer) mustEmbedUnimplementedDeleteAreaServer() {}

// UnsafeDeleteAreaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeleteAreaServer will
// result in compilation errors.
type UnsafeDeleteAreaServer interface {
	mustEmbedUnimplementedDeleteAreaServer()
}

func RegisterDeleteAreaServer(s grpc.ServiceRegistrar, srv DeleteAreaServer) {
	s.RegisterService(&DeleteArea_ServiceDesc, srv)
}

func _DeleteArea_DeleteArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteAreaServer).DeleteArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeleteArea/DeleteArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteAreaServer).DeleteArea(ctx, req.(*DeleteAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeleteArea_ServiceDesc is the grpc.ServiceDesc for DeleteArea service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeleteArea_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DeleteArea",
	HandlerType: (*DeleteAreaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteArea",
			Handler:    _DeleteArea_DeleteArea_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}

// DeleteServiceClient is the client API for DeleteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeleteServiceClient interface {
	// Sends a DeleteService request.
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceReply, error)
}

type deleteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeleteServiceClient(cc grpc.ClientConnInterface) DeleteServiceClient {
	return &deleteServiceClient{cc}
}

func (c *deleteServiceClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceReply, error) {
	out := new(DeleteServiceReply)
	err := c.cc.Invoke(ctx, "/pb.DeleteService/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteServiceServer is the server API for DeleteService service.
// All implementations must embed UnimplementedDeleteServiceServer
// for forward compatibility
type DeleteServiceServer interface {
	// Sends a DeleteService request.
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceReply, error)
	mustEmbedUnimplementedDeleteServiceServer()
}

// UnimplementedDeleteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeleteServiceServer struct {
}

func (UnimplementedDeleteServiceServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedDeleteServiceServer) mustEmbedUnimplementedDeleteServiceServer() {}

// UnsafeDeleteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeleteServiceServer will
// result in compilation errors.
type UnsafeDeleteServiceServer interface {
	mustEmbedUnimplementedDeleteServiceServer()
}

func RegisterDeleteServiceServer(s grpc.ServiceRegistrar, srv DeleteServiceServer) {
	s.RegisterService(&DeleteService_ServiceDesc, srv)
}

func _DeleteService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeleteService/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteServiceServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeleteService_ServiceDesc is the grpc.ServiceDesc for DeleteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeleteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DeleteService",
	HandlerType: (*DeleteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteService",
			Handler:    _DeleteService_DeleteService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pb.proto",
}
