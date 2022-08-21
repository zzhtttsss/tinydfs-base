// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: Heartbeat.proto

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

// HeartbeatServiceClient is the client API for HeartbeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartbeatServiceClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatArgs, opts ...grpc.CallOption) (*HeartbeatReply, error)
}

type heartbeatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartbeatServiceClient(cc grpc.ClientConnInterface) HeartbeatServiceClient {
	return &heartbeatServiceClient{cc}
}

func (c *heartbeatServiceClient) Heartbeat(ctx context.Context, in *HeartbeatArgs, opts ...grpc.CallOption) (*HeartbeatReply, error) {
	out := new(HeartbeatReply)
	err := c.cc.Invoke(ctx, "/pb.HeartbeatService/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServiceServer is the server API for HeartbeatService service.
// All implementations must embed UnimplementedHeartbeatServiceServer
// for forward compatibility
type HeartbeatServiceServer interface {
	Heartbeat(context.Context, *HeartbeatArgs) (*HeartbeatReply, error)
	mustEmbedUnimplementedHeartbeatServiceServer()
}

// UnimplementedHeartbeatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeartbeatServiceServer struct {
}

func (UnimplementedHeartbeatServiceServer) Heartbeat(context.Context, *HeartbeatArgs) (*HeartbeatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedHeartbeatServiceServer) mustEmbedUnimplementedHeartbeatServiceServer() {}

// UnsafeHeartbeatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartbeatServiceServer will
// result in compilation errors.
type UnsafeHeartbeatServiceServer interface {
	mustEmbedUnimplementedHeartbeatServiceServer()
}

func RegisterHeartbeatServiceServer(s grpc.ServiceRegistrar, srv HeartbeatServiceServer) {
	s.RegisterService(&HeartbeatService_ServiceDesc, srv)
}

func _HeartbeatService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HeartbeatService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, req.(*HeartbeatArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartbeatService_ServiceDesc is the grpc.ServiceDesc for HeartbeatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartbeatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HeartbeatService",
	HandlerType: (*HeartbeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _HeartbeatService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Heartbeat.proto",
}
