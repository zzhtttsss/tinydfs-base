// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AddFile.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CheckArgs4AddArgs
// path: target path to add file
// fileName: name of the file to add
// size: size of the file. Use bytes as the unit of measurement which means 1kb will be 1024.
type CheckArgs4AddArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *CheckArgs4AddArgs) Reset() {
	*x = CheckArgs4AddArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddFile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckArgs4AddArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckArgs4AddArgs) ProtoMessage() {}

func (x *CheckArgs4AddArgs) ProtoReflect() protoreflect.Message {
	mi := &file_AddFile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckArgs4AddArgs.ProtoReflect.Descriptor instead.
func (*CheckArgs4AddArgs) Descriptor() ([]byte, []int) {
	return file_AddFile_proto_rawDescGZIP(), []int{0}
}

func (x *CheckArgs4AddArgs) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CheckArgs4AddArgs) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CheckArgs4AddArgs) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// CheckArgs4AddReply
// fileNodeId: file id stored in the directory tree
// chunkNum: the number of chunks the file will be cut into
type CheckArgs4AddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileNodeId string `protobuf:"bytes,1,opt,name=fileNodeId,proto3" json:"fileNodeId,omitempty"`
	ChunkNum   int32  `protobuf:"varint,2,opt,name=chunkNum,proto3" json:"chunkNum,omitempty"`
}

func (x *CheckArgs4AddReply) Reset() {
	*x = CheckArgs4AddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddFile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckArgs4AddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckArgs4AddReply) ProtoMessage() {}

func (x *CheckArgs4AddReply) ProtoReflect() protoreflect.Message {
	mi := &file_AddFile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckArgs4AddReply.ProtoReflect.Descriptor instead.
func (*CheckArgs4AddReply) Descriptor() ([]byte, []int) {
	return file_AddFile_proto_rawDescGZIP(), []int{1}
}

func (x *CheckArgs4AddReply) GetFileNodeId() string {
	if x != nil {
		return x.FileNodeId
	}
	return ""
}

func (x *CheckArgs4AddReply) GetChunkNum() int32 {
	if x != nil {
		return x.ChunkNum
	}
	return 0
}

// GetDataNodes4AddArgs
// fileNodeId: file id stored in the directory tree
// chunkIndex: the index of chunk
type GetDataNodes4AddArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileNodeId string `protobuf:"bytes,1,opt,name=fileNodeId,proto3" json:"fileNodeId,omitempty"`
	ChunkIndex int32  `protobuf:"varint,2,opt,name=chunkIndex,proto3" json:"chunkIndex,omitempty"`
}

func (x *GetDataNodes4AddArgs) Reset() {
	*x = GetDataNodes4AddArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddFile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataNodes4AddArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataNodes4AddArgs) ProtoMessage() {}

func (x *GetDataNodes4AddArgs) ProtoReflect() protoreflect.Message {
	mi := &file_AddFile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataNodes4AddArgs.ProtoReflect.Descriptor instead.
func (*GetDataNodes4AddArgs) Descriptor() ([]byte, []int) {
	return file_AddFile_proto_rawDescGZIP(), []int{2}
}

func (x *GetDataNodes4AddArgs) GetFileNodeId() string {
	if x != nil {
		return x.FileNodeId
	}
	return ""
}

func (x *GetDataNodes4AddArgs) GetChunkIndex() int32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

// GetDataNodes4AddReply
// dataNodes: ids of all DataNodes that will store this chunk
// primaryNode: id of the primary DataNode
type GetDataNodes4AddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataNodes   []string `protobuf:"bytes,1,rep,name=dataNodes,proto3" json:"dataNodes,omitempty"`
	PrimaryNode string   `protobuf:"bytes,2,opt,name=primaryNode,proto3" json:"primaryNode,omitempty"`
}

func (x *GetDataNodes4AddReply) Reset() {
	*x = GetDataNodes4AddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddFile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataNodes4AddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataNodes4AddReply) ProtoMessage() {}

func (x *GetDataNodes4AddReply) ProtoReflect() protoreflect.Message {
	mi := &file_AddFile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataNodes4AddReply.ProtoReflect.Descriptor instead.
func (*GetDataNodes4AddReply) Descriptor() ([]byte, []int) {
	return file_AddFile_proto_rawDescGZIP(), []int{3}
}

func (x *GetDataNodes4AddReply) GetDataNodes() []string {
	if x != nil {
		return x.DataNodes
	}
	return nil
}

func (x *GetDataNodes4AddReply) GetPrimaryNode() string {
	if x != nil {
		return x.PrimaryNode
	}
	return ""
}

// TransferFileArgs
// fileNodeId: file id stored in the directory tree
// chunkIndex: the index of chunk
type TransferFileArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId string `protobuf:"bytes,1,opt,name=chunkId,proto3" json:"chunkId,omitempty"`
	Chunk   []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *TransferFileArgs) Reset() {
	*x = TransferFileArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddFile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFileArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFileArgs) ProtoMessage() {}

func (x *TransferFileArgs) ProtoReflect() protoreflect.Message {
	mi := &file_AddFile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFileArgs.ProtoReflect.Descriptor instead.
func (*TransferFileArgs) Descriptor() ([]byte, []int) {
	return file_AddFile_proto_rawDescGZIP(), []int{4}
}

func (x *TransferFileArgs) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *TransferFileArgs) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

// TransferFileReply
// dataNodes: ids of all DataNodes that will store this chunk
// primaryNode: id of the primary DataNode
type TransferFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransferFileReply) Reset() {
	*x = TransferFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddFile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFileReply) ProtoMessage() {}

func (x *TransferFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_AddFile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFileReply.ProtoReflect.Descriptor instead.
func (*TransferFileReply) Descriptor() ([]byte, []int) {
	return file_AddFile_proto_rawDescGZIP(), []int{5}
}

var File_AddFile_proto protoreflect.FileDescriptor

var file_AddFile_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x57, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73,
	0x34, 0x41, 0x64, 0x64, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x50, 0x0a, 0x12,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73, 0x34, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x22, 0x56,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x34, 0x41,
	0x64, 0x64, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x57, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x34, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x22,
	0x42, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x22, 0x13, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x9b, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73, 0x34, 0x41, 0x64, 0x64, 0x12, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73, 0x34, 0x41, 0x64,
	0x64, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x72, 0x67, 0x73, 0x34, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x47, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x34, 0x41, 0x64,
	0x64, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x34, 0x41, 0x64, 0x64, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x34, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x4d, 0x0a, 0x0e, 0x50, 0x69, 0x70, 0x4c, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AddFile_proto_rawDescOnce sync.Once
	file_AddFile_proto_rawDescData = file_AddFile_proto_rawDesc
)

func file_AddFile_proto_rawDescGZIP() []byte {
	file_AddFile_proto_rawDescOnce.Do(func() {
		file_AddFile_proto_rawDescData = protoimpl.X.CompressGZIP(file_AddFile_proto_rawDescData)
	})
	return file_AddFile_proto_rawDescData
}

var file_AddFile_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_AddFile_proto_goTypes = []interface{}{
	(*CheckArgs4AddArgs)(nil),     // 0: pb.CheckArgs4AddArgs
	(*CheckArgs4AddReply)(nil),    // 1: pb.CheckArgs4AddReply
	(*GetDataNodes4AddArgs)(nil),  // 2: pb.GetDataNodes4AddArgs
	(*GetDataNodes4AddReply)(nil), // 3: pb.GetDataNodes4AddReply
	(*TransferFileArgs)(nil),      // 4: pb.TransferFileArgs
	(*TransferFileReply)(nil),     // 5: pb.TransferFileReply
}
var file_AddFile_proto_depIdxs = []int32{
	0, // 0: pb.MasterAddService.CheckArgs4Add:input_type -> pb.CheckArgs4AddArgs
	2, // 1: pb.MasterAddService.GetDataNodes4Add:input_type -> pb.GetDataNodes4AddArgs
	4, // 2: pb.PipLineService.TransferFile:input_type -> pb.TransferFileArgs
	1, // 3: pb.MasterAddService.CheckArgs4Add:output_type -> pb.CheckArgs4AddReply
	3, // 4: pb.MasterAddService.GetDataNodes4Add:output_type -> pb.GetDataNodes4AddReply
	5, // 5: pb.PipLineService.TransferFile:output_type -> pb.TransferFileReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AddFile_proto_init() }
func file_AddFile_proto_init() {
	if File_AddFile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AddFile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckArgs4AddArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_AddFile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckArgs4AddReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_AddFile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataNodes4AddArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_AddFile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataNodes4AddReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_AddFile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFileArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_AddFile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFileReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AddFile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_AddFile_proto_goTypes,
		DependencyIndexes: file_AddFile_proto_depIdxs,
		MessageInfos:      file_AddFile_proto_msgTypes,
	}.Build()
	File_AddFile_proto = out.File
	file_AddFile_proto_rawDesc = nil
	file_AddFile_proto_goTypes = nil
	file_AddFile_proto_depIdxs = nil
}
