// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: syncano/codebox/lb/v1/runner.proto

package lb

import (
	context "context"
	v1 "github.com/Syncano/syncanoapis/gen/go/syncano/codebox/script/v1"
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

	ConcurrencyKey   string `protobuf:"bytes,1,opt,name=concurrency_key,json=concurrencyKey,proto3" json:"concurrency_key,omitempty"`        // Key on which concurrency limiting will take place.
	ConcurrencyLimit int32  `protobuf:"varint,2,opt,name=concurrency_limit,json=concurrencyLimit,proto3" json:"concurrency_limit,omitempty"` // Max concurrency allowed for key.
}

func (x *RunMeta) Reset() {
	*x = RunMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_runner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunMeta) ProtoMessage() {}

func (x *RunMeta) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_runner_proto_msgTypes[0]
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
	return file_syncano_codebox_lb_v1_runner_proto_rawDescGZIP(), []int{0}
}

func (x *RunMeta) GetConcurrencyKey() string {
	if x != nil {
		return x.ConcurrencyKey
	}
	return ""
}

func (x *RunMeta) GetConcurrencyLimit() int32 {
	if x != nil {
		return x.ConcurrencyLimit
	}
	return 0
}

// RunRequest represents either a Meta message or a Chunk message.
// It should always consist of exactly 1 Meta, 1 Script Meta and optionally repeated Chunk messages.
type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*RunRequest_Meta
	//	*RunRequest_ScriptMeta
	//	*RunRequest_ScriptChunk
	Value isRunRequest_Value `protobuf_oneof:"value"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syncano_codebox_lb_v1_runner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_syncano_codebox_lb_v1_runner_proto_msgTypes[1]
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
	return file_syncano_codebox_lb_v1_runner_proto_rawDescGZIP(), []int{1}
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

func (x *RunRequest) GetScriptMeta() *v1.RunMeta {
	if x, ok := x.GetValue().(*RunRequest_ScriptMeta); ok {
		return x.ScriptMeta
	}
	return nil
}

func (x *RunRequest) GetScriptChunk() *v1.RunChunk {
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

type RunRequest_ScriptMeta struct {
	ScriptMeta *v1.RunMeta `protobuf:"bytes,2,opt,name=script_meta,json=scriptMeta,proto3,oneof"`
}

type RunRequest_ScriptChunk struct {
	ScriptChunk *v1.RunChunk `protobuf:"bytes,3,opt,name=script_chunk,json=scriptChunk,proto3,oneof"`
}

func (*RunRequest_Meta) isRunRequest_Value() {}

func (*RunRequest_ScriptMeta) isRunRequest_Value() {}

func (*RunRequest_ScriptChunk) isRunRequest_Value() {}

var File_syncano_codebox_lb_v1_runner_proto protoreflect.FileDescriptor

var file_syncano_codebox_lb_v1_runner_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f,
	0x78, 0x2f, 0x6c, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x73, 0x79, 0x6e,
	0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2f, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0xdc, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x0b, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78,
	0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x48, 0x0a, 0x0c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x0b, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x64, 0x0a, 0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x21, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x6c, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x6f, 0x78, 0x2e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62,
	0x6f, 0x78, 0x2f, 0x6c, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_syncano_codebox_lb_v1_runner_proto_rawDescOnce sync.Once
	file_syncano_codebox_lb_v1_runner_proto_rawDescData = file_syncano_codebox_lb_v1_runner_proto_rawDesc
)

func file_syncano_codebox_lb_v1_runner_proto_rawDescGZIP() []byte {
	file_syncano_codebox_lb_v1_runner_proto_rawDescOnce.Do(func() {
		file_syncano_codebox_lb_v1_runner_proto_rawDescData = protoimpl.X.CompressGZIP(file_syncano_codebox_lb_v1_runner_proto_rawDescData)
	})
	return file_syncano_codebox_lb_v1_runner_proto_rawDescData
}

var file_syncano_codebox_lb_v1_runner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_syncano_codebox_lb_v1_runner_proto_goTypes = []interface{}{
	(*RunMeta)(nil),        // 0: syncano.codebox.lb.v1.RunMeta
	(*RunRequest)(nil),     // 1: syncano.codebox.lb.v1.RunRequest
	(*v1.RunMeta)(nil),     // 2: syncano.codebox.script.v1.RunMeta
	(*v1.RunChunk)(nil),    // 3: syncano.codebox.script.v1.RunChunk
	(*v1.RunResponse)(nil), // 4: syncano.codebox.script.v1.RunResponse
}
var file_syncano_codebox_lb_v1_runner_proto_depIdxs = []int32{
	0, // 0: syncano.codebox.lb.v1.RunRequest.meta:type_name -> syncano.codebox.lb.v1.RunMeta
	2, // 1: syncano.codebox.lb.v1.RunRequest.script_meta:type_name -> syncano.codebox.script.v1.RunMeta
	3, // 2: syncano.codebox.lb.v1.RunRequest.script_chunk:type_name -> syncano.codebox.script.v1.RunChunk
	1, // 3: syncano.codebox.lb.v1.ScriptRunner.Run:input_type -> syncano.codebox.lb.v1.RunRequest
	4, // 4: syncano.codebox.lb.v1.ScriptRunner.Run:output_type -> syncano.codebox.script.v1.RunResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_syncano_codebox_lb_v1_runner_proto_init() }
func file_syncano_codebox_lb_v1_runner_proto_init() {
	if File_syncano_codebox_lb_v1_runner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_syncano_codebox_lb_v1_runner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_syncano_codebox_lb_v1_runner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_syncano_codebox_lb_v1_runner_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RunRequest_Meta)(nil),
		(*RunRequest_ScriptMeta)(nil),
		(*RunRequest_ScriptChunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_syncano_codebox_lb_v1_runner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_syncano_codebox_lb_v1_runner_proto_goTypes,
		DependencyIndexes: file_syncano_codebox_lb_v1_runner_proto_depIdxs,
		MessageInfos:      file_syncano_codebox_lb_v1_runner_proto_msgTypes,
	}.Build()
	File_syncano_codebox_lb_v1_runner_proto = out.File
	file_syncano_codebox_lb_v1_runner_proto_rawDesc = nil
	file_syncano_codebox_lb_v1_runner_proto_goTypes = nil
	file_syncano_codebox_lb_v1_runner_proto_depIdxs = nil
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
	stream, err := c.cc.NewStream(ctx, &_ScriptRunner_serviceDesc.Streams[0], "/syncano.codebox.lb.v1.ScriptRunner/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &scriptRunnerRunClient{stream}
	return x, nil
}

type ScriptRunner_RunClient interface {
	Send(*RunRequest) error
	Recv() (*v1.RunResponse, error)
	grpc.ClientStream
}

type scriptRunnerRunClient struct {
	grpc.ClientStream
}

func (x *scriptRunnerRunClient) Send(m *RunRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scriptRunnerRunClient) Recv() (*v1.RunResponse, error) {
	m := new(v1.RunResponse)
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
	Send(*v1.RunResponse) error
	Recv() (*RunRequest, error)
	grpc.ServerStream
}

type scriptRunnerRunServer struct {
	grpc.ServerStream
}

func (x *scriptRunnerRunServer) Send(m *v1.RunResponse) error {
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
	ServiceName: "syncano.codebox.lb.v1.ScriptRunner",
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
	Metadata: "syncano/codebox/lb/v1/runner.proto",
}
