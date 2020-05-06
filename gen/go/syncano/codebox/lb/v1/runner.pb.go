// Code generated by protoc-gen-go. DO NOT EDIT.
// source: syncano/codebox/lb/v1/runner.proto

package lb

import (
	context "context"
	fmt "fmt"
	v1 "github.com/Syncano/syncanoapis/gen/go/syncano/codebox/script/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// RunRequest represents either a Meta message or a Chunk message.
// It should always consist of exactly 1 Meta and optionally repeated Chunk messages.
type RunRequest struct {
	// Types that are valid to be assigned to Value:
	//	*RunRequest_Meta
	//	*RunRequest_Request
	Value                isRunRequest_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead677061985e507, []int{0}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

type isRunRequest_Value interface {
	isRunRequest_Value()
}

type RunRequest_Meta struct {
	Meta *RunRequest_MetaMessage `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type RunRequest_Request struct {
	Request *v1.RunRequest `protobuf:"bytes,2,opt,name=request,proto3,oneof"`
}

func (*RunRequest_Meta) isRunRequest_Value() {}

func (*RunRequest_Request) isRunRequest_Value() {}

func (m *RunRequest) GetValue() isRunRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RunRequest) GetMeta() *RunRequest_MetaMessage {
	if x, ok := m.GetValue().(*RunRequest_Meta); ok {
		return x.Meta
	}
	return nil
}

func (m *RunRequest) GetRequest() *v1.RunRequest {
	if x, ok := m.GetValue().(*RunRequest_Request); ok {
		return x.Request
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RunRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RunRequest_Meta)(nil),
		(*RunRequest_Request)(nil),
	}
}

// Meta message specifies fields to describe what is being run.
type RunRequest_MetaMessage struct {
	RequestId            string   `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ConcurrencyKey       string   `protobuf:"bytes,1,opt,name=concurrency_key,json=concurrencyKey,proto3" json:"concurrency_key,omitempty"`
	ConcurrencyLimit     int32    `protobuf:"varint,2,opt,name=concurrency_limit,json=concurrencyLimit,proto3" json:"concurrency_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunRequest_MetaMessage) Reset()         { *m = RunRequest_MetaMessage{} }
func (m *RunRequest_MetaMessage) String() string { return proto.CompactTextString(m) }
func (*RunRequest_MetaMessage) ProtoMessage()    {}
func (*RunRequest_MetaMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead677061985e507, []int{0, 0}
}

func (m *RunRequest_MetaMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest_MetaMessage.Unmarshal(m, b)
}
func (m *RunRequest_MetaMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest_MetaMessage.Marshal(b, m, deterministic)
}
func (m *RunRequest_MetaMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest_MetaMessage.Merge(m, src)
}
func (m *RunRequest_MetaMessage) XXX_Size() int {
	return xxx_messageInfo_RunRequest_MetaMessage.Size(m)
}
func (m *RunRequest_MetaMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest_MetaMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest_MetaMessage proto.InternalMessageInfo

func (m *RunRequest_MetaMessage) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *RunRequest_MetaMessage) GetConcurrencyKey() string {
	if m != nil {
		return m.ConcurrencyKey
	}
	return ""
}

func (m *RunRequest_MetaMessage) GetConcurrencyLimit() int32 {
	if m != nil {
		return m.ConcurrencyLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*RunRequest)(nil), "syncano.codebox.lb.v1.RunRequest")
	proto.RegisterType((*RunRequest_MetaMessage)(nil), "syncano.codebox.lb.v1.RunRequest.MetaMessage")
}

func init() { proto.RegisterFile("syncano/codebox/lb/v1/runner.proto", fileDescriptor_ead677061985e507) }

var fileDescriptor_ead677061985e507 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x9b, 0xf6, 0xed, 0x5b, 0x3a, 0x15, 0xff, 0x2c, 0x08, 0xa5, 0x20, 0xd4, 0x82, 0xb5,
	0x20, 0x6e, 0x6c, 0x3d, 0x0a, 0xa2, 0xf5, 0x52, 0xd1, 0x5e, 0xb6, 0x9e, 0xbc, 0x94, 0x64, 0x33,
	0xc4, 0x60, 0xba, 0x1b, 0x77, 0xb3, 0xc5, 0x5c, 0xfd, 0x28, 0x7e, 0x52, 0x69, 0xb2, 0x6a, 0xaa,
	0x42, 0x6f, 0xcb, 0x33, 0xf3, 0x7b, 0x78, 0x66, 0x66, 0xa1, 0xa7, 0x33, 0xc1, 0x3d, 0x21, 0x5d,
	0x2e, 0x03, 0xf4, 0xe5, 0xab, 0x1b, 0xfb, 0xee, 0x72, 0xe8, 0x2a, 0x23, 0x04, 0x2a, 0x9a, 0x28,
	0x99, 0x4a, 0xb2, 0x6f, 0x7b, 0xa8, 0xed, 0xa1, 0xb1, 0x4f, 0x97, 0xc3, 0x4e, 0xff, 0x27, 0xaa,
	0xb9, 0x8a, 0x92, 0x74, 0x85, 0x17, 0xaf, 0x02, 0xef, 0xbd, 0x57, 0x01, 0x98, 0x11, 0x0c, 0x5f,
	0x0c, 0xea, 0x94, 0xdc, 0xc0, 0xbf, 0x05, 0xa6, 0x5e, 0xdb, 0xe9, 0x3a, 0x83, 0xd6, 0xe8, 0x94,
	0xfe, 0x69, 0x4e, 0xbf, 0x01, 0x3a, 0xc5, 0xd4, 0x9b, 0xa2, 0xd6, 0x5e, 0x88, 0x93, 0x0a, 0xcb,
	0x61, 0x72, 0x0d, 0x0d, 0x55, 0x94, 0xdb, 0xd5, 0xdc, 0xe7, 0xe8, 0x97, 0x8f, 0xcd, 0xb0, 0xe6,
	0x35, 0xa9, 0xb0, 0x4f, 0xae, 0xf3, 0xe6, 0x40, 0xab, 0x64, 0x4d, 0x0e, 0x00, 0x6c, 0x69, 0x1e,
	0x05, 0xed, 0x5a, 0xd7, 0x19, 0x34, 0x59, 0xd3, 0x2a, 0xb7, 0x01, 0x39, 0x86, 0x1d, 0x2e, 0x05,
	0x37, 0x4a, 0xa1, 0xe0, 0xd9, 0xfc, 0x19, 0xb3, 0x7c, 0x82, 0x26, 0xdb, 0x2e, 0xc9, 0x77, 0x98,
	0x91, 0x13, 0xd8, 0x2b, 0x37, 0xc6, 0xd1, 0x22, 0x2a, 0x42, 0xd6, 0xd9, 0x6e, 0xa9, 0x70, 0xbf,
	0xd2, 0xc7, 0x0d, 0xa8, 0x2f, 0xbd, 0xd8, 0xe0, 0x28, 0x80, 0xad, 0x59, 0x1e, 0x98, 0xe5, 0x9b,
	0x27, 0x0f, 0x50, 0x63, 0x46, 0x90, 0xc3, 0x8d, 0xeb, 0xe9, 0xf4, 0x37, 0x4d, 0xae, 0x13, 0x29,
	0x34, 0x0e, 0x9c, 0x33, 0x67, 0x7c, 0xf5, 0x78, 0x19, 0x46, 0xe9, 0x93, 0xf1, 0x29, 0x97, 0x0b,
	0x77, 0x66, 0xef, 0x67, 0x79, 0x2f, 0x89, 0xb4, 0x1b, 0xa2, 0x70, 0xc3, 0x2f, 0x69, 0xfd, 0x57,
	0x5c, 0xc4, 0xbe, 0xff, 0x3f, 0xbf, 0xe9, 0xf9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xd4,
	0xc7, 0xf6, 0x38, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScriptRunnerClient is the client API for ScriptRunner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScriptRunnerClient interface {
	// Run runs script in secure environment of worker.
	Run(ctx context.Context, opts ...grpc.CallOption) (ScriptRunner_RunClient, error)
}

type scriptRunnerClient struct {
	cc *grpc.ClientConn
}

func NewScriptRunnerClient(cc *grpc.ClientConn) ScriptRunnerClient {
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
