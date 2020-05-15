// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: syncano/codebox/lb/v1/workerplug.proto

package lb

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

type ContainerRemovedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContainerId string `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	SourceHash  string `protobuf:"bytes,3,opt,name=source_hash,json=sourceHash,proto3" json:"source_hash,omitempty"`
	Environment string `protobuf:"bytes,4,opt,name=environment,proto3" json:"environment,omitempty"`
	UserId      string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ContainerRemovedRequest) Reset() {
	*x = ContainerRemovedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerRemovedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerRemovedRequest) ProtoMessage() {}

func (x *ContainerRemovedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerRemovedRequest.ProtoReflect.Descriptor instead.
func (*ContainerRemovedRequest) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerRemovedRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContainerRemovedRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ContainerRemovedRequest) GetSourceHash() string {
	if x != nil {
		return x.SourceHash
	}
	return ""
}

func (x *ContainerRemovedRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *ContainerRemovedRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ContainerRemovedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContainerRemovedResponse) Reset() {
	*x = ContainerRemovedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerRemovedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerRemovedResponse) ProtoMessage() {}

func (x *ContainerRemovedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerRemovedResponse.ProtoReflect.Descriptor instead.
func (*ContainerRemovedResponse) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP(), []int{1}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Port        uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Mcpu        uint32 `protobuf:"varint,3,opt,name=mcpu,proto3" json:"mcpu,omitempty"`
	Memory      uint64 `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	DefaultMcpu uint32 `protobuf:"varint,5,opt,name=default_mcpu,json=defaultMcpu,proto3" json:"default_mcpu,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RegisterRequest) GetMcpu() uint32 {
	if x != nil {
		return x.Mcpu
	}
	return 0
}

func (x *RegisterRequest) GetMemory() uint64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *RegisterRequest) GetDefaultMcpu() uint32 {
	if x != nil {
		return x.DefaultMcpu
	}
	return 0
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP(), []int{3}
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Memory uint64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP(), []int{4}
}

func (x *HeartbeatRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HeartbeatRequest) GetMemory() uint64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP(), []int{5}
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP(), []int{6}
}

func (x *DisconnectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP(), []int{7}
}

var File_syncano_codebox_lb_v1_workerplug_proto protoreflect.FileDescriptor

var file_syncano_codebox_lb_v1_workerplug_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f,
	0x78, 0x2f, 0x6c, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x70, 0x6c,
	0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e,
	0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x22,
	0xa8, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x63, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x63,
	0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x63, 0x70, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x63, 0x70, 0x75, 0x22, 0x12, 0x0a,
	0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3a, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x13, 0x0a,
	0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa1, 0x03,
	0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x12, 0x73, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x2e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62,
	0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62,
	0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x26, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e,
	0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e,
	0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x27, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61,
	0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x28, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x61,
	0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2f, 0x6c, 0x62, 0x2f, 0x76, 0x31,
	0x3b, 0x6c, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_syncano_codebox_lb_v1_workerplug_proto_rawDescOnce sync.Once
	file_syncano_codebox_lb_v1_workerplug_proto_rawDescData = file_syncano_codebox_lb_v1_workerplug_proto_rawDesc
)

func file_syncano_codebox_lb_v1_workerplug_proto_rawDescGZIP() []byte {
	file_syncano_codebox_lb_v1_workerplug_proto_rawDescOnce.Do(func() {
		file_syncano_codebox_lb_v1_workerplug_proto_rawDescData = protoimpl.X.CompressGZIP(file_syncano_codebox_lb_v1_workerplug_proto_rawDescData)
	})
	return file_syncano_codebox_lb_v1_workerplug_proto_rawDescData
}

var file_syncano_codebox_lb_v1_workerplug_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_syncano_codebox_lb_v1_workerplug_proto_goTypes = []interface{}{
	(*ContainerRemovedRequest)(nil),  // 0: syncano.codebox.lb.v1.ContainerRemovedRequest
	(*ContainerRemovedResponse)(nil), // 1: syncano.codebox.lb.v1.ContainerRemovedResponse
	(*RegisterRequest)(nil),          // 2: syncano.codebox.lb.v1.RegisterRequest
	(*RegisterResponse)(nil),         // 3: syncano.codebox.lb.v1.RegisterResponse
	(*HeartbeatRequest)(nil),         // 4: syncano.codebox.lb.v1.HeartbeatRequest
	(*HeartbeatResponse)(nil),        // 5: syncano.codebox.lb.v1.HeartbeatResponse
	(*DisconnectRequest)(nil),        // 6: syncano.codebox.lb.v1.DisconnectRequest
	(*DisconnectResponse)(nil),       // 7: syncano.codebox.lb.v1.DisconnectResponse
}
var file_syncano_codebox_lb_v1_workerplug_proto_depIdxs = []int32{
	0, // 0: syncano.codebox.lb.v1.WorkerPlug.ContainerRemoved:input_type -> syncano.codebox.lb.v1.ContainerRemovedRequest
	2, // 1: syncano.codebox.lb.v1.WorkerPlug.Register:input_type -> syncano.codebox.lb.v1.RegisterRequest
	4, // 2: syncano.codebox.lb.v1.WorkerPlug.Heartbeat:input_type -> syncano.codebox.lb.v1.HeartbeatRequest
	6, // 3: syncano.codebox.lb.v1.WorkerPlug.Disconnect:input_type -> syncano.codebox.lb.v1.DisconnectRequest
	1, // 4: syncano.codebox.lb.v1.WorkerPlug.ContainerRemoved:output_type -> syncano.codebox.lb.v1.ContainerRemovedResponse
	3, // 5: syncano.codebox.lb.v1.WorkerPlug.Register:output_type -> syncano.codebox.lb.v1.RegisterResponse
	5, // 6: syncano.codebox.lb.v1.WorkerPlug.Heartbeat:output_type -> syncano.codebox.lb.v1.HeartbeatResponse
	7, // 7: syncano.codebox.lb.v1.WorkerPlug.Disconnect:output_type -> syncano.codebox.lb.v1.DisconnectResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_syncano_codebox_lb_v1_workerplug_proto_init() }
func file_syncano_codebox_lb_v1_workerplug_proto_init() {
	if File_syncano_codebox_lb_v1_workerplug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerRemovedRequest); i {
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
		file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerRemovedResponse); i {
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
		file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_syncano_codebox_lb_v1_workerplug_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectResponse); i {
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
			RawDescriptor: file_syncano_codebox_lb_v1_workerplug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_syncano_codebox_lb_v1_workerplug_proto_goTypes,
		DependencyIndexes: file_syncano_codebox_lb_v1_workerplug_proto_depIdxs,
		MessageInfos:      file_syncano_codebox_lb_v1_workerplug_proto_msgTypes,
	}.Build()
	File_syncano_codebox_lb_v1_workerplug_proto = out.File
	file_syncano_codebox_lb_v1_workerplug_proto_rawDesc = nil
	file_syncano_codebox_lb_v1_workerplug_proto_goTypes = nil
	file_syncano_codebox_lb_v1_workerplug_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkerPlugClient is the client API for WorkerPlug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerPlugClient interface {
	// ContainerRemoved handles notifications sent by client whenever a container gets removed from cache.
	ContainerRemoved(ctx context.Context, in *ContainerRemovedRequest, opts ...grpc.CallOption) (*ContainerRemovedResponse, error)
	// Register is sent at the beginning by the worker.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Heartbeat is meant to be called periodically by worker. If it's not sent for some time, worker will be removed.
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	// Disconnect gracefully removes worker.
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
}

type workerPlugClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerPlugClient(cc grpc.ClientConnInterface) WorkerPlugClient {
	return &workerPlugClient{cc}
}

func (c *workerPlugClient) ContainerRemoved(ctx context.Context, in *ContainerRemovedRequest, opts ...grpc.CallOption) (*ContainerRemovedResponse, error) {
	out := new(ContainerRemovedResponse)
	err := c.cc.Invoke(ctx, "/syncano.codebox.lb.v1.WorkerPlug/ContainerRemoved", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerPlugClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/syncano.codebox.lb.v1.WorkerPlug/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerPlugClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/syncano.codebox.lb.v1.WorkerPlug/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerPlugClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/syncano.codebox.lb.v1.WorkerPlug/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerPlugServer is the server API for WorkerPlug service.
type WorkerPlugServer interface {
	// ContainerRemoved handles notifications sent by client whenever a container gets removed from cache.
	ContainerRemoved(context.Context, *ContainerRemovedRequest) (*ContainerRemovedResponse, error)
	// Register is sent at the beginning by the worker.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Heartbeat is meant to be called periodically by worker. If it's not sent for some time, worker will be removed.
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	// Disconnect gracefully removes worker.
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
}

// UnimplementedWorkerPlugServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerPlugServer struct {
}

func (*UnimplementedWorkerPlugServer) ContainerRemoved(context.Context, *ContainerRemovedRequest) (*ContainerRemovedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerRemoved not implemented")
}
func (*UnimplementedWorkerPlugServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedWorkerPlugServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (*UnimplementedWorkerPlugServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

func RegisterWorkerPlugServer(s *grpc.Server, srv WorkerPlugServer) {
	s.RegisterService(&_WorkerPlug_serviceDesc, srv)
}

func _WorkerPlug_ContainerRemoved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerRemovedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerPlugServer).ContainerRemoved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncano.codebox.lb.v1.WorkerPlug/ContainerRemoved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerPlugServer).ContainerRemoved(ctx, req.(*ContainerRemovedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerPlug_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerPlugServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncano.codebox.lb.v1.WorkerPlug/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerPlugServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerPlug_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerPlugServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncano.codebox.lb.v1.WorkerPlug/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerPlugServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerPlug_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerPlugServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncano.codebox.lb.v1.WorkerPlug/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerPlugServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkerPlug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "syncano.codebox.lb.v1.WorkerPlug",
	HandlerType: (*WorkerPlugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContainerRemoved",
			Handler:    _WorkerPlug_ContainerRemoved_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _WorkerPlug_Register_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _WorkerPlug_Heartbeat_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _WorkerPlug_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncano/codebox/lb/v1/workerplug.proto",
}
