// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: AddFile.proto

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

// MasterAddServiceClient is the client API for MasterAddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterAddServiceClient interface {
	// CheckArgs4Add Called by client.
	// Check whether the path and file name entered by the user in the Add operation are legal.
	CheckArgs4Add(ctx context.Context, in *CheckArgs4AddArgs, opts ...grpc.CallOption) (*CheckArgs4AddReply, error)
	// GetDataNodes4Add Called by client.
	// Allocate some DataNode to store a Chunk and select the primary DataNode
	GetDataNodes4Add(ctx context.Context, in *GetDataNodes4AddArgs, opts ...grpc.CallOption) (*GetDataNodes4AddReply, error)
}

type masterAddServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterAddServiceClient(cc grpc.ClientConnInterface) MasterAddServiceClient {
	return &masterAddServiceClient{cc}
}

func (c *masterAddServiceClient) CheckArgs4Add(ctx context.Context, in *CheckArgs4AddArgs, opts ...grpc.CallOption) (*CheckArgs4AddReply, error) {
	out := new(CheckArgs4AddReply)
	err := c.cc.Invoke(ctx, "/pb.MasterAddService/CheckArgs4Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAddServiceClient) GetDataNodes4Add(ctx context.Context, in *GetDataNodes4AddArgs, opts ...grpc.CallOption) (*GetDataNodes4AddReply, error) {
	out := new(GetDataNodes4AddReply)
	err := c.cc.Invoke(ctx, "/pb.MasterAddService/GetDataNodes4Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterAddServiceServer is the server API for MasterAddService service.
// All implementations must embed UnimplementedMasterAddServiceServer
// for forward compatibility
type MasterAddServiceServer interface {
	// CheckArgs4Add Called by client.
	// Check whether the path and file name entered by the user in the Add operation are legal.
	CheckArgs4Add(context.Context, *CheckArgs4AddArgs) (*CheckArgs4AddReply, error)
	// GetDataNodes4Add Called by client.
	// Allocate some DataNode to store a Chunk and select the primary DataNode
	GetDataNodes4Add(context.Context, *GetDataNodes4AddArgs) (*GetDataNodes4AddReply, error)
	mustEmbedUnimplementedMasterAddServiceServer()
}

// UnimplementedMasterAddServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMasterAddServiceServer struct {
}

func (UnimplementedMasterAddServiceServer) CheckArgs4Add(context.Context, *CheckArgs4AddArgs) (*CheckArgs4AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckArgs4Add not implemented")
}
func (UnimplementedMasterAddServiceServer) GetDataNodes4Add(context.Context, *GetDataNodes4AddArgs) (*GetDataNodes4AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataNodes4Add not implemented")
}
func (UnimplementedMasterAddServiceServer) mustEmbedUnimplementedMasterAddServiceServer() {}

// UnsafeMasterAddServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterAddServiceServer will
// result in compilation errors.
type UnsafeMasterAddServiceServer interface {
	mustEmbedUnimplementedMasterAddServiceServer()
}

func RegisterMasterAddServiceServer(s grpc.ServiceRegistrar, srv MasterAddServiceServer) {
	s.RegisterService(&MasterAddService_ServiceDesc, srv)
}

func _MasterAddService_CheckArgs4Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckArgs4AddArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAddServiceServer).CheckArgs4Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterAddService/CheckArgs4Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAddServiceServer).CheckArgs4Add(ctx, req.(*CheckArgs4AddArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAddService_GetDataNodes4Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataNodes4AddArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAddServiceServer).GetDataNodes4Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterAddService/GetDataNodes4Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAddServiceServer).GetDataNodes4Add(ctx, req.(*GetDataNodes4AddArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterAddService_ServiceDesc is the grpc.ServiceDesc for MasterAddService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterAddService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MasterAddService",
	HandlerType: (*MasterAddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckArgs4Add",
			Handler:    _MasterAddService_CheckArgs4Add_Handler,
		},
		{
			MethodName: "GetDataNodes4Add",
			Handler:    _MasterAddService_GetDataNodes4Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AddFile.proto",
}

// PipLineServiceClient is the client API for PipLineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipLineServiceClient interface {
	// TransferFile Called by client or chunkserver.
	// Transfer a chunk of the file to a chunkserver using stream and let that chunkserver transfer
	// this chunk to another chunkserver if needed.
	TransferChunk(ctx context.Context, opts ...grpc.CallOption) (PipLineService_TransferChunkClient, error)
}

type pipLineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipLineServiceClient(cc grpc.ClientConnInterface) PipLineServiceClient {
	return &pipLineServiceClient{cc}
}

func (c *pipLineServiceClient) TransferChunk(ctx context.Context, opts ...grpc.CallOption) (PipLineService_TransferChunkClient, error) {
	stream, err := c.cc.NewStream(ctx, &PipLineService_ServiceDesc.Streams[0], "/pb.PipLineService/TransferChunk", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipLineServiceTransferChunkClient{stream}
	return x, nil
}

type PipLineService_TransferChunkClient interface {
	Send(*PieceOfChunk) error
	CloseAndRecv() (*TransferChunkReply, error)
	grpc.ClientStream
}

type pipLineServiceTransferChunkClient struct {
	grpc.ClientStream
}

func (x *pipLineServiceTransferChunkClient) Send(m *PieceOfChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pipLineServiceTransferChunkClient) CloseAndRecv() (*TransferChunkReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TransferChunkReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PipLineServiceServer is the server API for PipLineService service.
// All implementations must embed UnimplementedPipLineServiceServer
// for forward compatibility
type PipLineServiceServer interface {
	// TransferFile Called by client or chunkserver.
	// Transfer a chunk of the file to a chunkserver using stream and let that chunkserver transfer
	// this chunk to another chunkserver if needed.
	TransferChunk(PipLineService_TransferChunkServer) error
	mustEmbedUnimplementedPipLineServiceServer()
}

// UnimplementedPipLineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPipLineServiceServer struct {
}

func (UnimplementedPipLineServiceServer) TransferChunk(PipLineService_TransferChunkServer) error {
	return status.Errorf(codes.Unimplemented, "method TransferChunk not implemented")
}
func (UnimplementedPipLineServiceServer) mustEmbedUnimplementedPipLineServiceServer() {}

// UnsafePipLineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipLineServiceServer will
// result in compilation errors.
type UnsafePipLineServiceServer interface {
	mustEmbedUnimplementedPipLineServiceServer()
}

func RegisterPipLineServiceServer(s grpc.ServiceRegistrar, srv PipLineServiceServer) {
	s.RegisterService(&PipLineService_ServiceDesc, srv)
}

func _PipLineService_TransferChunk_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PipLineServiceServer).TransferChunk(&pipLineServiceTransferChunkServer{stream})
}

type PipLineService_TransferChunkServer interface {
	SendAndClose(*TransferChunkReply) error
	Recv() (*PieceOfChunk, error)
	grpc.ServerStream
}

type pipLineServiceTransferChunkServer struct {
	grpc.ServerStream
}

func (x *pipLineServiceTransferChunkServer) SendAndClose(m *TransferChunkReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pipLineServiceTransferChunkServer) Recv() (*PieceOfChunk, error) {
	m := new(PieceOfChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PipLineService_ServiceDesc is the grpc.ServiceDesc for PipLineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipLineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PipLineService",
	HandlerType: (*PipLineServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferChunk",
			Handler:       _PipLineService_TransferChunk_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "AddFile.proto",
}
