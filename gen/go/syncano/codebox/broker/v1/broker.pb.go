// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.13.0
// source: syncano/codebox/broker/v1/broker.proto

package broker

import (
	context "context"
	v1 "github.com/Syncano/syncanoapis/gen/go/syncano/codebox/lb/v1"
	v11 "github.com/Syncano/syncanoapis/gen/go/syncano/codebox/script/v1"
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

// Meta message specifies fields to describe what is being run.
type RunMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files          map[string]string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of script file URL->Name.
	EnvironmentUrl string            `protobuf:"bytes,2,opt,name=environment_url,json=environmentUrl,proto3" json:"environment_url,omitempty"`                                                 // Environment file URL.
	Trace          []byte            `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
	TraceId        uint64            `protobuf:"varint,4,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Sync           bool              `protobuf:"varint,5,opt,name=sync,proto3" json:"sync,omitempty"`
}

func (x *RunMeta) Reset() {
	*x = RunMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_broker_v1_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunMeta) ProtoMessage() {}

func (x *RunMeta) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_broker_v1_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunMeta.ProtoReflect.Descriptor instead.
func (*RunMeta) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_broker_v1_broker_proto_rawDescGZIP(), []int{0}
}

func (x *RunMeta) GetFiles() map[string]string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *RunMeta) GetEnvironmentUrl() string {
	if x != nil {
		return x.EnvironmentUrl
	}
	return ""
}

func (x *RunMeta) GetTrace() []byte {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *RunMeta) GetTraceId() uint64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (x *RunMeta) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

// RunRequest represents either a Meta message, LB Meta or a Chunk message.
// It should always consist of exactly 1 Meta, 1 LB Meta, 1 Script Meta and optionally repeated Script Chunk messages.
type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*RunRequest_Meta
	//	*RunRequest_LbMeta
	//	*RunRequest_ScriptMeta
	//	*RunRequest_ScriptChunk
	Value isRunRequest_Value `protobuf_oneof:"value"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_broker_v1_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_broker_v1_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_broker_v1_broker_proto_rawDescGZIP(), []int{1}
}

func (m *RunRequest) GetValue() isRunRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *RunRequest) GetMeta() *RunMeta {
	if x, ok := x.GetValue().(*RunRequest_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *RunRequest) GetLbMeta() *v1.RunMeta {
	if x, ok := x.GetValue().(*RunRequest_LbMeta); ok {
		return x.LbMeta
	}
	return nil
}

func (x *RunRequest) GetScriptMeta() *v11.RunMeta {
	if x, ok := x.GetValue().(*RunRequest_ScriptMeta); ok {
		return x.ScriptMeta
	}
	return nil
}

func (x *RunRequest) GetScriptChunk() *v11.RunChunk {
	if x, ok := x.GetValue().(*RunRequest_ScriptChunk); ok {
		return x.ScriptChunk
	}
	return nil
}

type isRunRequest_Value interface {
	isRunRequest_Value()
}

type RunRequest_Meta struct {
	Meta *RunMeta `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type RunRequest_LbMeta struct {
	LbMeta *v1.RunMeta `protobuf:"bytes,2,opt,name=lb_meta,json=lbMeta,proto3,oneof"`
}

type RunRequest_ScriptMeta struct {
	ScriptMeta *v11.RunMeta `protobuf:"bytes,3,opt,name=script_meta,json=scriptMeta,proto3,oneof"`
}

type RunRequest_ScriptChunk struct {
	ScriptChunk *v11.RunChunk `protobuf:"bytes,4,opt,name=script_chunk,json=scriptChunk,proto3,oneof"`
}

func (*RunRequest_Meta) isRunRequest_Value() {}

func (*RunRequest_LbMeta) isRunRequest_Value() {}

func (*RunRequest_ScriptMeta) isRunRequest_Value() {}

func (*RunRequest_ScriptChunk) isRunRequest_Value() {}

type SimpleRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *RunMeta     `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	LbMeta     *v1.RunMeta  `protobuf:"bytes,2,opt,name=lb_meta,json=lbMeta,proto3" json:"lb_meta,omitempty"`
	ScriptMeta *v11.RunMeta `protobuf:"bytes,3,opt,name=script_meta,json=scriptMeta,proto3" json:"script_meta,omitempty"`
}

func (x *SimpleRunRequest) Reset() {
	*x = SimpleRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_broker_v1_broker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRunRequest) ProtoMessage() {}

func (x *SimpleRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_broker_v1_broker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRunRequest.ProtoReflect.Descriptor instead.
func (*SimpleRunRequest) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_broker_v1_broker_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleRunRequest) GetMeta() *RunMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SimpleRunRequest) GetLbMeta() *v1.RunMeta {
	if x != nil {
		return x.LbMeta
	}
	return nil
}

func (x *SimpleRunRequest) GetScriptMeta() *v11.RunMeta {
	if x != nil {
		return x.ScriptMeta
	}
	return nil
}

var File_syncano_codebox_broker_v1_broker_proto protoreflect.FileDescriptor

var file_syncano_codebox_broker_v1_broker_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f,
	0x78, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e,
	0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x22, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x62, 0x6f, 0x78, 0x2f, 0x6c, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf6, 0x01, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x79,
	0x6e, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x1a, 0x38,
	0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b, 0x02, 0x0a, 0x0a, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x39, 0x0a, 0x07, 0x6c, 0x62, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x0b,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x48, 0x0a, 0x0c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00,
	0x52, 0x0b, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x07, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x07, 0x6c, 0x62, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x06, 0x6c, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x0b,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x32, 0xab, 0x02, 0x0a, 0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x58, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x25, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62,
	0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x62, 0x0a, 0x09,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x75, 0x6e, 0x12, 0x2b, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x5d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x79,
	0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_syncano_codebox_broker_v1_broker_proto_rawDescOnce sync.Once
	file_syncano_codebox_broker_v1_broker_proto_rawDescData = file_syncano_codebox_broker_v1_broker_proto_rawDesc
)

func file_syncano_codebox_broker_v1_broker_proto_rawDescGZIP() []byte {
	file_syncano_codebox_broker_v1_broker_proto_rawDescOnce.Do(func() {
		file_syncano_codebox_broker_v1_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_syncano_codebox_broker_v1_broker_proto_rawDescData)
	})
	return file_syncano_codebox_broker_v1_broker_proto_rawDescData
}

var file_syncano_codebox_broker_v1_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_syncano_codebox_broker_v1_broker_proto_goTypes = []interface{}{
	(*RunMeta)(nil),            // 0: syncano.codebox.broker.v1.RunMeta
	(*RunRequest)(nil),         // 1: syncano.codebox.broker.v1.RunRequest
	(*SimpleRunRequest)(nil),   // 2: syncano.codebox.broker.v1.SimpleRunRequest
	nil,                        // 3: syncano.codebox.broker.v1.RunMeta.FilesEntry
	(*v1.RunMeta)(nil),         // 4: syncano.codebox.lb.v1.RunMeta
	(*v11.RunMeta)(nil),        // 5: syncano.codebox.script.v1.RunMeta
	(*v11.RunChunk)(nil),       // 6: syncano.codebox.script.v1.RunChunk
	(*v11.DeleteRequest)(nil),  // 7: syncano.codebox.script.v1.DeleteRequest
	(*v11.RunResponse)(nil),    // 8: syncano.codebox.script.v1.RunResponse
	(*v11.DeleteResponse)(nil), // 9: syncano.codebox.script.v1.DeleteResponse
}
var file_syncano_codebox_broker_v1_broker_proto_depIdxs = []int32{
	3,  // 0: syncano.codebox.broker.v1.RunMeta.files:type_name -> syncano.codebox.broker.v1.RunMeta.FilesEntry
	0,  // 1: syncano.codebox.broker.v1.RunRequest.meta:type_name -> syncano.codebox.broker.v1.RunMeta
	4,  // 2: syncano.codebox.broker.v1.RunRequest.lb_meta:type_name -> syncano.codebox.lb.v1.RunMeta
	5,  // 3: syncano.codebox.broker.v1.RunRequest.script_meta:type_name -> syncano.codebox.script.v1.RunMeta
	6,  // 4: syncano.codebox.broker.v1.RunRequest.script_chunk:type_name -> syncano.codebox.script.v1.RunChunk
	0,  // 5: syncano.codebox.broker.v1.SimpleRunRequest.meta:type_name -> syncano.codebox.broker.v1.RunMeta
	4,  // 6: syncano.codebox.broker.v1.SimpleRunRequest.lb_meta:type_name -> syncano.codebox.lb.v1.RunMeta
	5,  // 7: syncano.codebox.broker.v1.SimpleRunRequest.script_meta:type_name -> syncano.codebox.script.v1.RunMeta
	1,  // 8: syncano.codebox.broker.v1.ScriptRunner.Run:input_type -> syncano.codebox.broker.v1.RunRequest
	2,  // 9: syncano.codebox.broker.v1.ScriptRunner.SimpleRun:input_type -> syncano.codebox.broker.v1.SimpleRunRequest
	7,  // 10: syncano.codebox.broker.v1.ScriptRunner.Delete:input_type -> syncano.codebox.script.v1.DeleteRequest
	8,  // 11: syncano.codebox.broker.v1.ScriptRunner.Run:output_type -> syncano.codebox.script.v1.RunResponse
	8,  // 12: syncano.codebox.broker.v1.ScriptRunner.SimpleRun:output_type -> syncano.codebox.script.v1.RunResponse
	9,  // 13: syncano.codebox.broker.v1.ScriptRunner.Delete:output_type -> syncano.codebox.script.v1.DeleteResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_syncano_codebox_broker_v1_broker_proto_init() }
func file_syncano_codebox_broker_v1_broker_proto_init() {
	if File_syncano_codebox_broker_v1_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_syncano_codebox_broker_v1_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunMeta); i {
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
		file_syncano_codebox_broker_v1_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
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
		file_syncano_codebox_broker_v1_broker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRunRequest); i {
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
	file_syncano_codebox_broker_v1_broker_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RunRequest_Meta)(nil),
		(*RunRequest_LbMeta)(nil),
		(*RunRequest_ScriptMeta)(nil),
		(*RunRequest_ScriptChunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_syncano_codebox_broker_v1_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_syncano_codebox_broker_v1_broker_proto_goTypes,
		DependencyIndexes: file_syncano_codebox_broker_v1_broker_proto_depIdxs,
		MessageInfos:      file_syncano_codebox_broker_v1_broker_proto_msgTypes,
	}.Build()
	File_syncano_codebox_broker_v1_broker_proto = out.File
	file_syncano_codebox_broker_v1_broker_proto_rawDesc = nil
	file_syncano_codebox_broker_v1_broker_proto_goTypes = nil
	file_syncano_codebox_broker_v1_broker_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ScriptRunnerClient is the client API for ScriptRunner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScriptRunnerClient interface {
	// Run runs script in secure environment.
	Run(ctx context.Context, opts ...grpc.CallOption) (ScriptRunner_RunClient, error)
	// SimpleRun is a simpler alternative to Run that does not require streaming request.
	// As such, it is only usable for small payloads and does not support chunks.
	SimpleRun(ctx context.Context, in *SimpleRunRequest, opts ...grpc.CallOption) (ScriptRunner_SimpleRunClient, error)
	// Delete deletes all containers for specified script index.
	Delete(ctx context.Context, in *v11.DeleteRequest, opts ...grpc.CallOption) (*v11.DeleteResponse, error)
}

type scriptRunnerClient struct {
	cc grpc.ClientConnInterface
}

func NewScriptRunnerClient(cc grpc.ClientConnInterface) ScriptRunnerClient {
	return &scriptRunnerClient{cc}
}

func (c *scriptRunnerClient) Run(ctx context.Context, opts ...grpc.CallOption) (ScriptRunner_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScriptRunner_serviceDesc.Streams[0], "/syncano.codebox.broker.v1.ScriptRunner/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &scriptRunnerRunClient{stream}
	return x, nil
}

type ScriptRunner_RunClient interface {
	Send(*RunRequest) error
	Recv() (*v11.RunResponse, error)
	grpc.ClientStream
}

type scriptRunnerRunClient struct {
	grpc.ClientStream
}

func (x *scriptRunnerRunClient) Send(m *RunRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scriptRunnerRunClient) Recv() (*v11.RunResponse, error) {
	m := new(v11.RunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scriptRunnerClient) SimpleRun(ctx context.Context, in *SimpleRunRequest, opts ...grpc.CallOption) (ScriptRunner_SimpleRunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScriptRunner_serviceDesc.Streams[1], "/syncano.codebox.broker.v1.ScriptRunner/SimpleRun", opts...)
	if err != nil {
		return nil, err
	}
	x := &scriptRunnerSimpleRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ScriptRunner_SimpleRunClient interface {
	Recv() (*v11.RunResponse, error)
	grpc.ClientStream
}

type scriptRunnerSimpleRunClient struct {
	grpc.ClientStream
}

func (x *scriptRunnerSimpleRunClient) Recv() (*v11.RunResponse, error) {
	m := new(v11.RunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scriptRunnerClient) Delete(ctx context.Context, in *v11.DeleteRequest, opts ...grpc.CallOption) (*v11.DeleteResponse, error) {
	out := new(v11.DeleteResponse)
	err := c.cc.Invoke(ctx, "/syncano.codebox.broker.v1.ScriptRunner/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScriptRunnerServer is the server API for ScriptRunner service.
type ScriptRunnerServer interface {
	// Run runs script in secure environment.
	Run(ScriptRunner_RunServer) error
	// SimpleRun is a simpler alternative to Run that does not require streaming request.
	// As such, it is only usable for small payloads and does not support chunks.
	SimpleRun(*SimpleRunRequest, ScriptRunner_SimpleRunServer) error
	// Delete deletes all containers for specified script index.
	Delete(context.Context, *v11.DeleteRequest) (*v11.DeleteResponse, error)
}

// UnimplementedScriptRunnerServer can be embedded to have forward compatible implementations.
type UnimplementedScriptRunnerServer struct {
}

func (*UnimplementedScriptRunnerServer) Run(ScriptRunner_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedScriptRunnerServer) SimpleRun(*SimpleRunRequest, ScriptRunner_SimpleRunServer) error {
	return status.Errorf(codes.Unimplemented, "method SimpleRun not implemented")
}
func (*UnimplementedScriptRunnerServer) Delete(context.Context, *v11.DeleteRequest) (*v11.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterScriptRunnerServer(s *grpc.Server, srv ScriptRunnerServer) {
	s.RegisterService(&_ScriptRunner_serviceDesc, srv)
}

func _ScriptRunner_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScriptRunnerServer).Run(&scriptRunnerRunServer{stream})
}

type ScriptRunner_RunServer interface {
	Send(*v11.RunResponse) error
	Recv() (*RunRequest, error)
	grpc.ServerStream
}

type scriptRunnerRunServer struct {
	grpc.ServerStream
}

func (x *scriptRunnerRunServer) Send(m *v11.RunResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scriptRunnerRunServer) Recv() (*RunRequest, error) {
	m := new(RunRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScriptRunner_SimpleRun_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRunRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ScriptRunnerServer).SimpleRun(m, &scriptRunnerSimpleRunServer{stream})
}

type ScriptRunner_SimpleRunServer interface {
	Send(*v11.RunResponse) error
	grpc.ServerStream
}

type scriptRunnerSimpleRunServer struct {
	grpc.ServerStream
}

func (x *scriptRunnerSimpleRunServer) Send(m *v11.RunResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ScriptRunner_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScriptRunnerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/syncano.codebox.broker.v1.ScriptRunner/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScriptRunnerServer).Delete(ctx, req.(*v11.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScriptRunner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "syncano.codebox.broker.v1.ScriptRunner",
	HandlerType: (*ScriptRunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _ScriptRunner_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _ScriptRunner_Run_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SimpleRun",
			Handler:       _ScriptRunner_SimpleRun_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "syncano/codebox/broker/v1/broker.proto",
}
