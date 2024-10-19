// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/box/box.proto

package boxv1

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

const (
	BoxService_IsLock_FullMethodName = "/proto.box.v1.BoxService/isLock"
	BoxService_Lock_FullMethodName   = "/proto.box.v1.BoxService/lock"
	BoxService_Unlock_FullMethodName = "/proto.box.v1.BoxService/unlock"
)

// BoxServiceClient is the client API for BoxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoxServiceClient interface {
	IsLock(ctx context.Context, in *IsLockRequest, opts ...grpc.CallOption) (*IsLockResponse, error)
	Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error)
	Unlock(ctx context.Context, in *UnlockRequest, opts ...grpc.CallOption) (*UnlockResponse, error)
}

type boxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoxServiceClient(cc grpc.ClientConnInterface) BoxServiceClient {
	return &boxServiceClient{cc}
}

func (c *boxServiceClient) IsLock(ctx context.Context, in *IsLockRequest, opts ...grpc.CallOption) (*IsLockResponse, error) {
	out := new(IsLockResponse)
	err := c.cc.Invoke(ctx, BoxService_IsLock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error) {
	out := new(LockResponse)
	err := c.cc.Invoke(ctx, BoxService_Lock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) Unlock(ctx context.Context, in *UnlockRequest, opts ...grpc.CallOption) (*UnlockResponse, error) {
	out := new(UnlockResponse)
	err := c.cc.Invoke(ctx, BoxService_Unlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoxServiceServer is the server API for BoxService service.
// All implementations must embed UnimplementedBoxServiceServer
// for forward compatibility
type BoxServiceServer interface {
	IsLock(context.Context, *IsLockRequest) (*IsLockResponse, error)
	Lock(context.Context, *LockRequest) (*LockResponse, error)
	Unlock(context.Context, *UnlockRequest) (*UnlockResponse, error)
	mustEmbedUnimplementedBoxServiceServer()
}

// UnimplementedBoxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBoxServiceServer struct {
}

func (UnimplementedBoxServiceServer) IsLock(context.Context, *IsLockRequest) (*IsLockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLock not implemented")
}
func (UnimplementedBoxServiceServer) Lock(context.Context, *LockRequest) (*LockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (UnimplementedBoxServiceServer) Unlock(context.Context, *UnlockRequest) (*UnlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlock not implemented")
}
func (UnimplementedBoxServiceServer) mustEmbedUnimplementedBoxServiceServer() {}

// UnsafeBoxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoxServiceServer will
// result in compilation errors.
type UnsafeBoxServiceServer interface {
	mustEmbedUnimplementedBoxServiceServer()
}

func RegisterBoxServiceServer(s grpc.ServiceRegistrar, srv BoxServiceServer) {
	s.RegisterService(&BoxService_ServiceDesc, srv)
}

func _BoxService_IsLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).IsLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_IsLock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).IsLock(ctx, req.(*IsLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_Lock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).Lock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_Unlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).Unlock(ctx, req.(*UnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoxService_ServiceDesc is the grpc.ServiceDesc for BoxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.box.v1.BoxService",
	HandlerType: (*BoxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "isLock",
			Handler:    _BoxService_IsLock_Handler,
		},
		{
			MethodName: "lock",
			Handler:    _BoxService_Lock_Handler,
		},
		{
			MethodName: "unlock",
			Handler:    _BoxService_Unlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/box/box.proto",
}
