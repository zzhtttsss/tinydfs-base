// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: Stat.proto

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

// MasterStatServiceClient is the client API for MasterStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterStatServiceClient interface {
	// CheckAndStat Called by client.
	// Check args and get the file info.
	CheckAndStat(ctx context.Context, in *CheckAndStatArgs, opts ...grpc.CallOption) (*CheckAndStatReply, error)
}

type masterStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterStatServiceClient(cc grpc.ClientConnInterface) MasterStatServiceClient {
	return &masterStatServiceClient{cc}
}

func (c *masterStatServiceClient) CheckAndStat(ctx context.Context, in *CheckAndStatArgs, opts ...grpc.CallOption) (*CheckAndStatReply, error) {
	out := new(CheckAndStatReply)
	err := c.cc.Invoke(ctx, "/pb.MasterStatService/CheckAndStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterStatServiceServer is the server API for MasterStatService service.
// All implementations must embed UnimplementedMasterStatServiceServer
// for forward compatibility
type MasterStatServiceServer interface {
	// CheckAndStat Called by client.
	// Check args and get the file info.
	CheckAndStat(context.Context, *CheckAndStatArgs) (*CheckAndStatReply, error)
	mustEmbedUnimplementedMasterStatServiceServer()
}

// UnimplementedMasterStatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMasterStatServiceServer struct {
}

func (UnimplementedMasterStatServiceServer) CheckAndStat(context.Context, *CheckAndStatArgs) (*CheckAndStatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndStat not implemented")
}
func (UnimplementedMasterStatServiceServer) mustEmbedUnimplementedMasterStatServiceServer() {}

// UnsafeMasterStatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterStatServiceServer will
// result in compilation errors.
type UnsafeMasterStatServiceServer interface {
	mustEmbedUnimplementedMasterStatServiceServer()
}

func RegisterMasterStatServiceServer(s grpc.ServiceRegistrar, srv MasterStatServiceServer) {
	s.RegisterService(&MasterStatService_ServiceDesc, srv)
}

func _MasterStatService_CheckAndStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAndStatArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterStatServiceServer).CheckAndStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterStatService/CheckAndStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterStatServiceServer).CheckAndStat(ctx, req.(*CheckAndStatArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterStatService_ServiceDesc is the grpc.ServiceDesc for MasterStatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterStatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MasterStatService",
	HandlerType: (*MasterStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAndStat",
			Handler:    _MasterStatService_CheckAndStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Stat.proto",
}
