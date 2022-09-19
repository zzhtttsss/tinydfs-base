// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: Remove.proto

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

// MasterRemoveServiceClient is the client API for MasterRemoveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterRemoveServiceClient interface {
	// CheckAndMove Called by client.
	// Check args and remove directory or file at target path.
	CheckAndMove(ctx context.Context, in *CheckAndRemoveArgs, opts ...grpc.CallOption) (*CheckAndRemoveReply, error)
}

type masterRemoveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterRemoveServiceClient(cc grpc.ClientConnInterface) MasterRemoveServiceClient {
	return &masterRemoveServiceClient{cc}
}

func (c *masterRemoveServiceClient) CheckAndMove(ctx context.Context, in *CheckAndRemoveArgs, opts ...grpc.CallOption) (*CheckAndRemoveReply, error) {
	out := new(CheckAndRemoveReply)
	err := c.cc.Invoke(ctx, "/pb.MasterRemoveService/CheckAndMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterRemoveServiceServer is the server API for MasterRemoveService service.
// All implementations must embed UnimplementedMasterRemoveServiceServer
// for forward compatibility
type MasterRemoveServiceServer interface {
	// CheckAndMove Called by client.
	// Check args and remove directory or file at target path.
	CheckAndMove(context.Context, *CheckAndRemoveArgs) (*CheckAndRemoveReply, error)
	mustEmbedUnimplementedMasterRemoveServiceServer()
}

// UnimplementedMasterRemoveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMasterRemoveServiceServer struct {
}

func (UnimplementedMasterRemoveServiceServer) CheckAndMove(context.Context, *CheckAndRemoveArgs) (*CheckAndRemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndMove not implemented")
}
func (UnimplementedMasterRemoveServiceServer) mustEmbedUnimplementedMasterRemoveServiceServer() {}

// UnsafeMasterRemoveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterRemoveServiceServer will
// result in compilation errors.
type UnsafeMasterRemoveServiceServer interface {
	mustEmbedUnimplementedMasterRemoveServiceServer()
}

func RegisterMasterRemoveServiceServer(s grpc.ServiceRegistrar, srv MasterRemoveServiceServer) {
	s.RegisterService(&MasterRemoveService_ServiceDesc, srv)
}

func _MasterRemoveService_CheckAndMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAndRemoveArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterRemoveServiceServer).CheckAndMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterRemoveService/CheckAndMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterRemoveServiceServer).CheckAndMove(ctx, req.(*CheckAndRemoveArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterRemoveService_ServiceDesc is the grpc.ServiceDesc for MasterRemoveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterRemoveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MasterRemoveService",
	HandlerType: (*MasterRemoveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAndMove",
			Handler:    _MasterRemoveService_CheckAndMove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Remove.proto",
}
