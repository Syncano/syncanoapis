// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: syncano/codebox/script/v1/script.proto

package script

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

type RunChunk_Type int32

const (
	RunChunk_GENERIC RunChunk_Type = 0
	RunChunk_ARGS    RunChunk_Type = 1
)

// Enum value maps for RunChunk_Type.
var (
	RunChunk_Type_name = map[int32]string{
		0: "GENERIC",
		1: "ARGS",
	}
	RunChunk_Type_value = map[string]int32{
		"GENERIC": 0,
		"ARGS":    1,
	}
)

func (x RunChunk_Type) Enum() *RunChunk_Type {
	p := new(RunChunk_Type)
	*p = x
	return p
}

func (x RunChunk_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunChunk_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_syncano_codebox_script_v1_script_proto_enumTypes[0].Descriptor()
}

func (RunChunk_Type) Type() protoreflect.EnumType {
	return &file_syncano_codebox_script_v1_script_proto_enumTypes[0]
}

func (x RunChunk_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunChunk_Type.Descriptor instead.
func (RunChunk_Type) EnumDescriptor() ([]byte, []int) {
	return file_syncano_codebox_script_v1_script_proto_rawDescGZIP(), []int{1, 0}
}

// Meta message specifies fields to describe what is being run.
type RunMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options     *RunMeta_Options `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	Runtime     string           `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`                         // Runtime name.
	SourceHash  string           `protobuf:"bytes,3,opt,name=source_hash,json=sourceHash,proto3" json:"source_hash,omitempty"` // Unique hash describing script sources.
	UserId      string           `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // Unique key relevant for user context.
	Environment string           `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`                 // Optional unique string describing script environment file.
}

func (x *RunMeta) Reset() {
	*x = RunMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunMeta) ProtoMessage() {}

func (x *RunMeta) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[0]
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
	return file_syncano_codebox_script_v1_script_proto_rawDescGZIP(), []int{0}
}

func (x *RunMeta) GetOptions() *RunMeta_Options {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *RunMeta) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *RunMeta) GetSourceHash() string {
	if x != nil {
		return x.SourceHash
	}
	return ""
}

func (x *RunMeta) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RunMeta) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

type RunChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Filename    string        `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	ContentType string        `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Data        []byte        `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Type        RunChunk_Type `protobuf:"varint,5,opt,name=type,proto3,enum=syncano.codebox.script.v1.RunChunk_Type" json:"type,omitempty"`
}

func (x *RunChunk) Reset() {
	*x = RunChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunChunk) ProtoMessage() {}

func (x *RunChunk) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunChunk.ProtoReflect.Descriptor instead.
func (*RunChunk) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_script_v1_script_proto_rawDescGZIP(), []int{1}
}

func (x *RunChunk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RunChunk) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *RunChunk) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *RunChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RunChunk) GetType() RunChunk_Type {
	if x != nil {
		return x.Type
	}
	return RunChunk_GENERIC
}

// RunRequest represents a Meta message and optionally repeated Chunk message.
type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*RunRequest_Meta
	//	*RunRequest_Chunk
	Value isRunRequest_Value `protobuf_oneof:"value"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[2]
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
	return file_syncano_codebox_script_v1_script_proto_rawDescGZIP(), []int{2}
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

func (x *RunRequest) GetChunk() *RunChunk {
	if x, ok := x.GetValue().(*RunRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isRunRequest_Value interface {
	isRunRequest_Value()
}

type RunRequest_Meta struct {
	Meta *RunMeta `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type RunRequest_Chunk struct {
	Chunk *RunChunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*RunRequest_Meta) isRunRequest_Value() {}

func (*RunRequest_Chunk) isRunRequest_Value() {}

// HTTPResponseMessage describes custom response from script.
type HTTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32             `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	ContentType string            `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Content     []byte            `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Headers     map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HTTPResponse) Reset() {
	*x = HTTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPResponse) ProtoMessage() {}

func (x *HTTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPResponse.ProtoReflect.Descriptor instead.
func (*HTTPResponse) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_script_v1_script_proto_rawDescGZIP(), []int{3}
}

func (x *HTTPResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *HTTPResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *HTTPResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *HTTPResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId string        `protobuf:"bytes,9,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Code        int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Stdout      []byte        `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr      []byte        `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Response    *HTTPResponse `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	Took        int64         `protobuf:"varint,5,opt,name=took,proto3" json:"took,omitempty"`
	Cached      bool          `protobuf:"varint,6,opt,name=cached,proto3" json:"cached,omitempty"`
	Time        int64         `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`
	Weight      uint32        `protobuf:"varint,8,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_script_v1_script_proto_rawDescGZIP(), []int{4}
}

func (x *RunResponse) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *RunResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RunResponse) GetStdout() []byte {
	if x != nil {
		return x.Stdout
	}
	return nil
}

func (x *RunResponse) GetStderr() []byte {
	if x != nil {
		return x.Stderr
	}
	return nil
}

func (x *RunResponse) GetResponse() *HTTPResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *RunResponse) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

func (x *RunResponse) GetCached() bool {
	if x != nil {
		return x.Cached
	}
	return false
}

func (x *RunResponse) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RunResponse) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type RunMeta_Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entrypoint  string `protobuf:"bytes,1,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`                       // If empty, uses default entrypoint for specified runtime.
	OutputLimit uint32 `protobuf:"varint,2,opt,name=output_limit,json=outputLimit,proto3" json:"output_limit,omitempty"` // If 0, will not hijack any output.
	Timeout     int64  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`                            // Timeout is meant to be in milliseconds. If 0 takes a default 30s value.
	Mcpu        uint32 `protobuf:"varint,4,opt,name=mcpu,proto3" json:"mcpu,omitempty"`                                  // MCPU is CPU constaints in milliseconds. If 0 takes a default 250 value.
	Async       uint32 `protobuf:"varint,5,opt,name=async,proto3" json:"async,omitempty"`                                // Async defines if async connections should be used. Disabled by default.
	// Empty args, config, meta are acceptable.
	Args   []byte `protobuf:"bytes,6,opt,name=args,proto3" json:"args,omitempty"`
	Config []byte `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	Meta   []byte `protobuf:"bytes,8,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *RunMeta_Options) Reset() {
	*x = RunMeta_Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunMeta_Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunMeta_Options) ProtoMessage() {}

func (x *RunMeta_Options) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_script_v1_script_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunMeta_Options.ProtoReflect.Descriptor instead.
func (*RunMeta_Options) Descriptor() ([]byte, []int) {
	return file_syncano_codebox_script_v1_script_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RunMeta_Options) GetEntrypoint() string {
	if x != nil {
		return x.Entrypoint
	}
	return ""
}

func (x *RunMeta_Options) GetOutputLimit() uint32 {
	if x != nil {
		return x.OutputLimit
	}
	return 0
}

func (x *RunMeta_Options) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *RunMeta_Options) GetMcpu() uint32 {
	if x != nil {
		return x.Mcpu
	}
	return 0
}

func (x *RunMeta_Options) GetAsync() uint32 {
	if x != nil {
		return x.Async
	}
	return 0
}

func (x *RunMeta_Options) GetArgs() []byte {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *RunMeta_Options) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RunMeta_Options) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_syncano_codebox_script_v1_script_proto protoreflect.FileDescriptor

var file_syncano_codebox_script_v1_script_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f,
	0x78, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e,
	0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x2e, 0x76, 0x31, 0x22, 0x98, 0x03, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x44, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62,
	0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0xd0, 0x01, 0x0a, 0x07,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x6d, 0x63, 0x70, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0xce,
	0x01, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62,
	0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x1d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x45, 0x4e, 0x45,
	0x52, 0x49, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x52, 0x47, 0x53, 0x10, 0x01, 0x22,
	0x8c, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e,
	0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf8,
	0x01, 0x0a, 0x0c, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78,
	0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x91, 0x02, 0x0a, 0x0b, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65,
	0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72,
	0x12, 0x43, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0x68, 0x0a,
	0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x58, 0x0a,
	0x03, 0x52, 0x75, 0x6e, 0x12, 0x25, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x73, 0x79,
	0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78,
	0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_syncano_codebox_script_v1_script_proto_rawDescOnce sync.Once
	file_syncano_codebox_script_v1_script_proto_rawDescData = file_syncano_codebox_script_v1_script_proto_rawDesc
)

func file_syncano_codebox_script_v1_script_proto_rawDescGZIP() []byte {
	file_syncano_codebox_script_v1_script_proto_rawDescOnce.Do(func() {
		file_syncano_codebox_script_v1_script_proto_rawDescData = protoimpl.X.CompressGZIP(file_syncano_codebox_script_v1_script_proto_rawDescData)
	})
	return file_syncano_codebox_script_v1_script_proto_rawDescData
}

var file_syncano_codebox_script_v1_script_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_syncano_codebox_script_v1_script_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_syncano_codebox_script_v1_script_proto_goTypes = []interface{}{
	(RunChunk_Type)(0),      // 0: syncano.codebox.script.v1.RunChunk.Type
	(*RunMeta)(nil),         // 1: syncano.codebox.script.v1.RunMeta
	(*RunChunk)(nil),        // 2: syncano.codebox.script.v1.RunChunk
	(*RunRequest)(nil),      // 3: syncano.codebox.script.v1.RunRequest
	(*HTTPResponse)(nil),    // 4: syncano.codebox.script.v1.HTTPResponse
	(*RunResponse)(nil),     // 5: syncano.codebox.script.v1.RunResponse
	(*RunMeta_Options)(nil), // 6: syncano.codebox.script.v1.RunMeta.Options
	nil,                     // 7: syncano.codebox.script.v1.HTTPResponse.HeadersEntry
}
var file_syncano_codebox_script_v1_script_proto_depIdxs = []int32{
	6, // 0: syncano.codebox.script.v1.RunMeta.options:type_name -> syncano.codebox.script.v1.RunMeta.Options
	0, // 1: syncano.codebox.script.v1.RunChunk.type:type_name -> syncano.codebox.script.v1.RunChunk.Type
	1, // 2: syncano.codebox.script.v1.RunRequest.meta:type_name -> syncano.codebox.script.v1.RunMeta
	2, // 3: syncano.codebox.script.v1.RunRequest.chunk:type_name -> syncano.codebox.script.v1.RunChunk
	7, // 4: syncano.codebox.script.v1.HTTPResponse.headers:type_name -> syncano.codebox.script.v1.HTTPResponse.HeadersEntry
	4, // 5: syncano.codebox.script.v1.RunResponse.response:type_name -> syncano.codebox.script.v1.HTTPResponse
	3, // 6: syncano.codebox.script.v1.ScriptRunner.Run:input_type -> syncano.codebox.script.v1.RunRequest
	5, // 7: syncano.codebox.script.v1.ScriptRunner.Run:output_type -> syncano.codebox.script.v1.RunResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_syncano_codebox_script_v1_script_proto_init() }
func file_syncano_codebox_script_v1_script_proto_init() {
	if File_syncano_codebox_script_v1_script_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_syncano_codebox_script_v1_script_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_syncano_codebox_script_v1_script_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunChunk); i {
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
		file_syncano_codebox_script_v1_script_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_syncano_codebox_script_v1_script_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPResponse); i {
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
		file_syncano_codebox_script_v1_script_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
		file_syncano_codebox_script_v1_script_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunMeta_Options); i {
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
	file_syncano_codebox_script_v1_script_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*RunRequest_Meta)(nil),
		(*RunRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_syncano_codebox_script_v1_script_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_syncano_codebox_script_v1_script_proto_goTypes,
		DependencyIndexes: file_syncano_codebox_script_v1_script_proto_depIdxs,
		EnumInfos:         file_syncano_codebox_script_v1_script_proto_enumTypes,
		MessageInfos:      file_syncano_codebox_script_v1_script_proto_msgTypes,
	}.Build()
	File_syncano_codebox_script_v1_script_proto = out.File
	file_syncano_codebox_script_v1_script_proto_rawDesc = nil
	file_syncano_codebox_script_v1_script_proto_goTypes = nil
	file_syncano_codebox_script_v1_script_proto_depIdxs = nil
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
	// Run runs script in secure environment of worker.
	Run(ctx context.Context, opts ...grpc.CallOption) (ScriptRunner_RunClient, error)
}

type scriptRunnerClient struct {
	cc grpc.ClientConnInterface
}

func NewScriptRunnerClient(cc grpc.ClientConnInterface) ScriptRunnerClient {
	return &scriptRunnerClient{cc}
}

func (c *scriptRunnerClient) Run(ctx context.Context, opts ...grpc.CallOption) (ScriptRunner_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScriptRunner_serviceDesc.Streams[0], "/syncano.codebox.script.v1.ScriptRunner/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &scriptRunnerRunClient{stream}
	return x, nil
}

type ScriptRunner_RunClient interface {
	Send(*RunRequest) error
	Recv() (*RunResponse, error)
	grpc.ClientStream
}

type scriptRunnerRunClient struct {
	grpc.ClientStream
}

func (x *scriptRunnerRunClient) Send(m *RunRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scriptRunnerRunClient) Recv() (*RunResponse, error) {
	m := new(RunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ScriptRunnerServer is the server API for ScriptRunner service.
type ScriptRunnerServer interface {
	// Run runs script in secure environment of worker.
	Run(ScriptRunner_RunServer) error
}

// UnimplementedScriptRunnerServer can be embedded to have forward compatible implementations.
type UnimplementedScriptRunnerServer struct {
}

func (*UnimplementedScriptRunnerServer) Run(ScriptRunner_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterScriptRunnerServer(s *grpc.Server, srv ScriptRunnerServer) {
	s.RegisterService(&_ScriptRunner_serviceDesc, srv)
}

func _ScriptRunner_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScriptRunnerServer).Run(&scriptRunnerRunServer{stream})
}

type ScriptRunner_RunServer interface {
	Send(*RunResponse) error
	Recv() (*RunRequest, error)
	grpc.ServerStream
}

type scriptRunnerRunServer struct {
	grpc.ServerStream
}

func (x *scriptRunnerRunServer) Send(m *RunResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scriptRunnerRunServer) Recv() (*RunRequest, error) {
	m := new(RunRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ScriptRunner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "syncano.codebox.script.v1.ScriptRunner",
	HandlerType: (*ScriptRunnerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _ScriptRunner_Run_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "syncano/codebox/script/v1/script.proto",
}
