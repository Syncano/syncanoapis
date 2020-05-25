// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.1
// source: syncano/codebox/filerepo/v1/repo.proto

package filerepo

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ExistsRequest) Reset() {
	*x = ExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsRequest) ProtoMessage() {}

func (x *ExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsRequest.ProtoReflect.Descriptor instead.
func (*ExistsRequest) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_filerepo_v1_repo_proto_rawDescGZIP(), []int{0}
}

func (x *ExistsRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ExistsResponse) Reset() {
	*x = ExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsResponse) ProtoMessage() {}

func (x *ExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsResponse.ProtoReflect.Descriptor instead.
func (*ExistsResponse) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_filerepo_v1_repo_proto_rawDescGZIP(), []int{1}
}

func (x *ExistsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type UploadMetaMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UploadMetaMessage) Reset() {
	*x = UploadMetaMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadMetaMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMetaMessage) ProtoMessage() {}

func (x *UploadMetaMessage) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMetaMessage.ProtoReflect.Descriptor instead.
func (*UploadMetaMessage) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_filerepo_v1_repo_proto_rawDescGZIP(), []int{2}
}

func (x *UploadMetaMessage) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type UploadChunkMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadChunkMessage) Reset() {
	*x = UploadChunkMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadChunkMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadChunkMessage) ProtoMessage() {}

func (x *UploadChunkMessage) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadChunkMessage.ProtoReflect.Descriptor instead.
func (*UploadChunkMessage) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_filerepo_v1_repo_proto_rawDescGZIP(), []int{3}
}

func (x *UploadChunkMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadChunkMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*UploadRequest_Meta
	//	*UploadRequest_Chunk
	//	*UploadRequest_Done
	Value isUploadRequest_Value `protobuf_oneof:"value"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_filerepo_v1_repo_proto_rawDescGZIP(), []int{4}
}

func (m *UploadRequest) GetValue() isUploadRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *UploadRequest) GetMeta() *UploadMetaMessage {
	if x, ok := x.GetValue().(*UploadRequest_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *UploadRequest) GetChunk() *UploadChunkMessage {
	if x, ok := x.GetValue().(*UploadRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (x *UploadRequest) GetDone() bool {
	if x, ok := x.GetValue().(*UploadRequest_Done); ok {
		return x.Done
	}
	return false
}

type isUploadRequest_Value interface {
	isUploadRequest_Value()
}

type UploadRequest_Meta struct {
	Meta *UploadMetaMessage `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type UploadRequest_Chunk struct {
	Chunk *UploadChunkMessage `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

type UploadRequest_Done struct {
	Done bool `protobuf:"varint,3,opt,name=done,proto3,oneof"`
}

func (*UploadRequest_Meta) isUploadRequest_Value() {}

func (*UploadRequest_Chunk) isUploadRequest_Value() {}

func (*UploadRequest_Done) isUploadRequest_Value() {}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_filerepo_v1_repo_proto_rawDescGZIP(), []int{5}
}

func (x *UploadResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

var File_syncano_codebox_filerepo_v1_repo_proto protoreflect.FileDescriptor

var file_syncano_codebox_filerepo_v1_repo_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f,
	0x78, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e,
	0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65,
	0x70, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x21, 0x0a, 0x0d, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x25, 0x0a, 0x11, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x3c, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0xbd, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x44, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f,
	0x78, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65, 0x70,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x14, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x32, 0xd0, 0x01,
	0x0a, 0x04, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x61, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x2a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62,
	0x6f, 0x78, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x06, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x2a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f,
	0x78, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e,
	0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65,
	0x70, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_syncano_codebox_filerepo_v1_repo_proto_rawDescOnce sync.Once
	file_syncano_codebox_filerepo_v1_repo_proto_rawDescData = file_syncano_codebox_filerepo_v1_repo_proto_rawDesc
)

func file_syncano_codebox_filerepo_v1_repo_proto_rawDescGZIP() []byte {
	file_syncano_codebox_filerepo_v1_repo_proto_rawDescOnce.Do(func() {
		file_syncano_codebox_filerepo_v1_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_syncano_codebox_filerepo_v1_repo_proto_rawDescData)
	})
	return file_syncano_codebox_filerepo_v1_repo_proto_rawDescData
}

var file_syncano_codebox_filerepo_v1_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_syncano_codebox_filerepo_v1_repo_proto_goTypes = []interface{}{
	(*ExistsRequest)(nil),      // 0: syncano.codebox.filerepo.v1.ExistsRequest
	(*ExistsResponse)(nil),     // 1: syncano.codebox.filerepo.v1.ExistsResponse
	(*UploadMetaMessage)(nil),  // 2: syncano.codebox.filerepo.v1.UploadMetaMessage
	(*UploadChunkMessage)(nil), // 3: syncano.codebox.filerepo.v1.UploadChunkMessage
	(*UploadRequest)(nil),      // 4: syncano.codebox.filerepo.v1.UploadRequest
	(*UploadResponse)(nil),     // 5: syncano.codebox.filerepo.v1.UploadResponse
}
var file_syncano_codebox_filerepo_v1_repo_proto_depIdxs = []int32{
	2, // 0: syncano.codebox.filerepo.v1.UploadRequest.meta:type_name -> syncano.codebox.filerepo.v1.UploadMetaMessage
	3, // 1: syncano.codebox.filerepo.v1.UploadRequest.chunk:type_name -> syncano.codebox.filerepo.v1.UploadChunkMessage
	0, // 2: syncano.codebox.filerepo.v1.Repo.Exists:input_type -> syncano.codebox.filerepo.v1.ExistsRequest
	4, // 3: syncano.codebox.filerepo.v1.Repo.Upload:input_type -> syncano.codebox.filerepo.v1.UploadRequest
	1, // 4: syncano.codebox.filerepo.v1.Repo.Exists:output_type -> syncano.codebox.filerepo.v1.ExistsResponse
	5, // 5: syncano.codebox.filerepo.v1.Repo.Upload:output_type -> syncano.codebox.filerepo.v1.UploadResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_syncano_codebox_filerepo_v1_repo_proto_init() }
func file_syncano_codebox_filerepo_v1_repo_proto_init() {
	if File_syncano_codebox_filerepo_v1_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsRequest); i {
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
		file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsResponse); i {
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
		file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadMetaMessage); i {
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
		file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadChunkMessage); i {
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
		file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
	file_syncano_codebox_filerepo_v1_repo_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UploadRequest_Meta)(nil),
		(*UploadRequest_Chunk)(nil),
		(*UploadRequest_Done)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_syncano_codebox_filerepo_v1_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_syncano_codebox_filerepo_v1_repo_proto_goTypes,
		DependencyIndexes: file_syncano_codebox_filerepo_v1_repo_proto_depIdxs,
		MessageInfos:      file_syncano_codebox_filerepo_v1_repo_proto_msgTypes,
	}.Build()
	File_syncano_codebox_filerepo_v1_repo_proto = out.File
	file_syncano_codebox_filerepo_v1_repo_proto_rawDesc = nil
	file_syncano_codebox_filerepo_v1_repo_proto_goTypes = nil
	file_syncano_codebox_filerepo_v1_repo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RepoClient is the client API for Repo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepoClient interface {
	// Exists checks if file was defined in file repo.
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	// Upload streams file(s) to server.
	Upload(ctx context.Context, opts ...grpc.CallOption) (Repo_UploadClient, error)
}

type repoClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoClient(cc grpc.ClientConnInterface) RepoClient {
	return &repoClient{cc}
}

func (c *repoClient) Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error) {
	out := new(ExistsResponse)
	err := c.cc.Invoke(ctx, "/syncano.codebox.filerepo.v1.Repo/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Repo_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Repo_serviceDesc.Streams[0], "/syncano.codebox.filerepo.v1.Repo/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &repoUploadClient{stream}
	return x, nil
}

type Repo_UploadClient interface {
	Send(*UploadRequest) error
	Recv() (*UploadResponse, error)
	grpc.ClientStream
}

type repoUploadClient struct {
	grpc.ClientStream
}

func (x *repoUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *repoUploadClient) Recv() (*UploadResponse, error) {
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RepoServer is the server API for Repo service.
type RepoServer interface {
	// Exists checks if file was defined in file repo.
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	// Upload streams file(s) to server.
	Upload(Repo_UploadServer) error
}

// UnimplementedRepoServer can be embedded to have forward compatible implementations.
type UnimplementedRepoServer struct {
}

func (*UnimplementedRepoServer) Exists(context.Context, *ExistsRequest) (*ExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (*UnimplementedRepoServer) Upload(Repo_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterRepoServer(s *grpc.Server, srv RepoServer) {
	s.RegisterService(&_Repo_serviceDesc, srv)
}

func _Repo_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncano.codebox.filerepo.v1.Repo/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServer).Exists(ctx, req.(*ExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repo_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RepoServer).Upload(&repoUploadServer{stream})
}

type Repo_UploadServer interface {
	Send(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type repoUploadServer struct {
	grpc.ServerStream
}

func (x *repoUploadServer) Send(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *repoUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Repo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "syncano.codebox.filerepo.v1.Repo",
	HandlerType: (*RepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exists",
			Handler:    _Repo_Exists_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Repo_Upload_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "syncano/codebox/filerepo/v1/repo.proto",
}
