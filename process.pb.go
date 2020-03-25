// Code generated by protoc-gen-go. DO NOT EDIT.
// source: process.proto

package data

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Process struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string              `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Transaction          string              `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Definition           *Process_Definition `protobuf:"bytes,10,opt,name=definition,proto3" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Process) Reset()         { *m = Process{} }
func (m *Process) String() string { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()    {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0}
}

func (m *Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process.Unmarshal(m, b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process.Marshal(b, m, deterministic)
}
func (m *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(m, src)
}
func (m *Process) XXX_Size() int {
	return xxx_messageInfo_Process.Size(m)
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

func (m *Process) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Process) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Process) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Process) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

func (m *Process) GetDefinition() *Process_Definition {
	if m != nil {
		return m.Definition
	}
	return nil
}

type Process_Definition struct {
	Input                *Process_Data `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output               *Process_Data `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Error                *Process_Data `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Process_Definition) Reset()         { *m = Process_Definition{} }
func (m *Process_Definition) String() string { return proto.CompactTextString(m) }
func (*Process_Definition) ProtoMessage()    {}
func (*Process_Definition) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0}
}

func (m *Process_Definition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Definition.Unmarshal(m, b)
}
func (m *Process_Definition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Definition.Marshal(b, m, deterministic)
}
func (m *Process_Definition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Definition.Merge(m, src)
}
func (m *Process_Definition) XXX_Size() int {
	return xxx_messageInfo_Process_Definition.Size(m)
}
func (m *Process_Definition) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Definition.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Definition proto.InternalMessageInfo

func (m *Process_Definition) GetInput() *Process_Data {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Process_Definition) GetOutput() *Process_Data {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *Process_Definition) GetError() *Process_Data {
	if m != nil {
		return m.Error
	}
	return nil
}

type Process_Data struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Entity               *Entity  `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Data) Reset()         { *m = Process_Data{} }
func (m *Process_Data) String() string { return proto.CompactTextString(m) }
func (*Process_Data) ProtoMessage()    {}
func (*Process_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 1}
}

func (m *Process_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Data.Unmarshal(m, b)
}
func (m *Process_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Data.Marshal(b, m, deterministic)
}
func (m *Process_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Data.Merge(m, src)
}
func (m *Process_Data) XXX_Size() int {
	return xxx_messageInfo_Process_Data.Size(m)
}
func (m *Process_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Data proto.InternalMessageInfo

func (m *Process_Data) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Process_Data) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type Input struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string              `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Event                map[string]*any.Any `protobuf:"bytes,10,rep,name=event,proto3" json:"event,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Entity               map[string]*any.Any `protobuf:"bytes,11,rep,name=entity,proto3" json:"entity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{1}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Input) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Input) GetEvent() map[string]*any.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Input) GetEntity() map[string]*any.Any {
	if m != nil {
		return m.Entity
	}
	return nil
}

type Output struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event                map[string]*any.Any `protobuf:"bytes,10,rep,name=event,proto3" json:"event,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Entity               map[string]*any.Any `protobuf:"bytes,11,rep,name=entity,proto3" json:"entity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{2}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Output) GetEvent() map[string]*any.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Output) GetEntity() map[string]*any.Any {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*Process)(nil), "data.Process")
	proto.RegisterType((*Process_Definition)(nil), "data.Process.Definition")
	proto.RegisterType((*Process_Data)(nil), "data.Process.Data")
	proto.RegisterType((*Input)(nil), "data.Input")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.Input.EntityEntry")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.Input.EventEntry")
	proto.RegisterType((*Output)(nil), "data.Output")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.Output.EntityEntry")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.Output.EventEntry")
}

func init() {
	proto.RegisterFile("process.proto", fileDescriptor_54c4d0e8c0aaf5c3)
}

var fileDescriptor_54c4d0e8c0aaf5c3 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe5, 0xf4, 0x9f, 0x78, 0xd3, 0x0d, 0x30, 0x7f, 0x16, 0x32, 0x86, 0x4a, 0xc6, 0xa1,
	0xaa, 0x20, 0xa9, 0xca, 0x65, 0xda, 0x6d, 0x83, 0x02, 0x93, 0xd0, 0x56, 0x05, 0x09, 0xb8, 0x7a,
	0x89, 0x37, 0x2c, 0x5a, 0x3b, 0x72, 0x9c, 0x8a, 0x0a, 0x71, 0x81, 0x2b, 0x37, 0x6e, 0x7c, 0x0f,
	0x6e, 0x7c, 0x0b, 0xbe, 0x02, 0x1f, 0x04, 0xc5, 0x76, 0xb7, 0xa6, 0xea, 0xe0, 0xb0, 0x13, 0xb7,
	0xd8, 0xef, 0xe3, 0x5f, 0x9e, 0xf7, 0xc9, 0x1b, 0xc3, 0x5a, 0x26, 0x45, 0x42, 0xf3, 0x3c, 0xcc,
	0xa4, 0x50, 0x02, 0xd7, 0x53, 0xa2, 0x88, 0x7f, 0xf7, 0x54, 0x88, 0xd3, 0x31, 0x8d, 0x48, 0xc6,
	0x22, 0xc2, 0xb9, 0x50, 0x44, 0x31, 0xc1, 0xad, 0xc6, 0xbf, 0x63, 0xab, 0x7a, 0x75, 0x5c, 0x9c,
	0x44, 0x84, 0xcf, 0x6c, 0xa9, 0x9d, 0x88, 0xc9, 0x44, 0x70, 0xbb, 0x72, 0xe9, 0x94, 0x72, 0x35,
	0x2f, 0x51, 0xae, 0x98, 0xb2, 0xc2, 0xe0, 0x4b, 0x0d, 0x5a, 0x23, 0xf3, 0x66, 0xbc, 0x0e, 0x0e,
	0x4b, 0x3d, 0xd4, 0x41, 0xdd, 0x2b, 0xb1, 0xc3, 0x52, 0x8c, 0xa1, 0xce, 0xc9, 0x84, 0x7a, 0x8e,
	0xde, 0xd1, 0xcf, 0xd8, 0x83, 0xd6, 0x94, 0xca, 0x9c, 0x09, 0xee, 0xd5, 0xf4, 0xf6, 0x7c, 0x89,
	0x3b, 0xe0, 0x2a, 0x49, 0x78, 0x4e, 0x92, 0xd2, 0xa3, 0x57, 0xd7, 0xd5, 0xc5, 0x2d, 0xbc, 0x03,
	0x90, 0xd2, 0x13, 0xc6, 0x99, 0x16, 0x40, 0x07, 0x75, 0xdd, 0x81, 0x17, 0x96, 0x8d, 0x86, 0xd6,
	0x42, 0xf8, 0xf4, 0xac, 0x1e, 0x2f, 0x68, 0xfd, 0xaf, 0x08, 0xe0, 0xbc, 0x84, 0xbb, 0xd0, 0x60,
	0x3c, 0x2b, 0x94, 0xf6, 0xea, 0x0e, 0xf0, 0x12, 0x83, 0x28, 0x12, 0x1b, 0x01, 0xee, 0x41, 0x53,
	0x14, 0xaa, 0x94, 0x3a, 0x17, 0x4a, 0xad, 0xa2, 0xa4, 0x52, 0x29, 0x85, 0xd4, 0x8d, 0x5d, 0x40,
	0xd5, 0x02, 0xff, 0x08, 0xea, 0xe5, 0x12, 0xdf, 0x87, 0x86, 0x4e, 0xd6, 0xfa, 0x70, 0xcd, 0x89,
	0x61, 0xb9, 0x15, 0x9b, 0x0a, 0x7e, 0x00, 0x4d, 0x93, 0xb7, 0x35, 0xd0, 0xb6, 0x1a, 0xbd, 0x17,
	0xdb, 0x5a, 0xf0, 0xc3, 0x81, 0xc6, 0x81, 0x36, 0x3c, 0xcf, 0x1c, 0xad, 0xce, 0xdc, 0xa9, 0x66,
	0xfe, 0x70, 0x6e, 0x00, 0x3a, 0xb5, 0xae, 0x3b, 0xb8, 0x6d, 0xe0, 0x9a, 0x64, 0x6c, 0x0c, 0xb9,
	0x92, 0xb3, 0xb9, 0x97, 0xe8, 0xcc, 0x8b, 0xab, 0xe5, 0x1b, 0x15, 0xb9, 0xae, 0x18, 0xbd, 0x95,
	0xf9, 0x87, 0x00, 0xe7, 0x14, 0x7c, 0x0d, 0x6a, 0xef, 0xe9, 0xcc, 0x3a, 0x2b, 0x1f, 0x71, 0x0f,
	0x1a, 0x53, 0x32, 0x2e, 0xa8, 0xed, 0xed, 0x66, 0x68, 0x06, 0x32, 0x9c, 0x0f, 0x64, 0xb8, 0xc7,
	0x67, 0xb1, 0x91, 0xec, 0x3a, 0x3b, 0xc8, 0x3f, 0x02, 0x77, 0xe1, 0x35, 0x97, 0x07, 0x06, 0xdf,
	0x1d, 0x68, 0x1e, 0x99, 0xaf, 0xb7, 0x3c, 0xbc, 0x8f, 0xaa, 0xd1, 0xd8, 0x5e, 0x8d, 0x78, 0x45,
	0x36, 0xfd, 0xa5, 0x6c, 0xbc, 0xaa, 0xfe, 0x3f, 0x0c, 0x67, 0xf0, 0x1a, 0xc0, 0x0e, 0xef, 0xde,
	0xe8, 0x00, 0xbf, 0x80, 0xd6, 0xf0, 0x03, 0x4d, 0x0a, 0x45, 0xb1, 0xbb, 0xf0, 0xdd, 0xfd, 0xf6,
	0x62, 0xa3, 0xc1, 0xbd, 0xcf, 0xbf, 0x7e, 0x7f, 0x73, 0xbc, 0xe0, 0x46, 0x64, 0xaf, 0xa3, 0x68,
	0xda, 0x8f, 0xa8, 0x39, 0xb7, 0x8b, 0x7a, 0x5d, 0xd4, 0x47, 0x83, 0x9f, 0x0e, 0xac, 0x5b, 0x70,
	0x4c, 0x49, 0x5a, 0xc2, 0x9f, 0x40, 0xfd, 0x25, 0xcb, 0x15, 0x5e, 0xb3, 0xb0, 0x4c, 0x5f, 0x53,
	0xfe, 0x5a, 0xe5, 0x17, 0x0a, 0x36, 0x35, 0xfc, 0x16, 0xae, 0xc0, 0x19, 0xcf, 0x33, 0x9a, 0xa8,
	0x3e, 0xc2, 0x87, 0xd0, 0x7a, 0x55, 0x24, 0xfa, 0x26, 0xfa, 0x3b, 0x67, 0x5b, 0x73, 0xb6, 0xf0,
	0xe6, 0x0a, 0x4e, 0x94, 0x1b, 0x84, 0xe1, 0x3d, 0x23, 0x6c, 0x5c, 0x48, 0x7a, 0x19, 0xde, 0x89,
	0x41, 0xf4, 0x11, 0x1e, 0x42, 0xed, 0x39, 0xfd, 0x57, 0x8f, 0x1d, 0xcd, 0xf2, 0xb1, 0xb7, 0x8a,
	0xf5, 0x91, 0xa5, 0x9f, 0x06, 0x6f, 0xe1, 0xaa, 0x15, 0xbf, 0x91, 0x4c, 0xd1, 0x32, 0xbe, 0x21,
	0x34, 0x63, 0x9a, 0x08, 0x99, 0xe2, 0x2a, 0x6d, 0x19, 0xbe, 0xa5, 0xe1, 0x1b, 0x01, 0x5e, 0x84,
	0x4b, 0x7d, 0x72, 0x17, 0xf5, 0xf6, 0xb7, 0xe1, 0x7a, 0x22, 0x26, 0x61, 0xf6, 0x8e, 0xe5, 0x05,
	0x53, 0x54, 0x9f, 0xdd, 0x6f, 0xdb, 0xc3, 0xa3, 0x72, 0x4a, 0x46, 0xe8, 0xb8, 0xa9, 0xc7, 0xe5,
	0xf1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xdd, 0x9e, 0x6f, 0x70, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProcessAPIClient is the client API for ProcessAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessAPIClient interface {
	Execute(ctx context.Context, opts ...grpc.CallOption) (ProcessAPI_ExecuteClient, error)
}

type processAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessAPIClient(cc grpc.ClientConnInterface) ProcessAPIClient {
	return &processAPIClient{cc}
}

func (c *processAPIClient) Execute(ctx context.Context, opts ...grpc.CallOption) (ProcessAPI_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessAPI_serviceDesc.Streams[0], "/data.ProcessAPI/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &processAPIExecuteClient{stream}
	return x, nil
}

type ProcessAPI_ExecuteClient interface {
	Send(*Input) error
	Recv() (*Output, error)
	grpc.ClientStream
}

type processAPIExecuteClient struct {
	grpc.ClientStream
}

func (x *processAPIExecuteClient) Send(m *Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processAPIExecuteClient) Recv() (*Output, error) {
	m := new(Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessAPIServer is the server API for ProcessAPI service.
type ProcessAPIServer interface {
	Execute(ProcessAPI_ExecuteServer) error
}

// UnimplementedProcessAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProcessAPIServer struct {
}

func (*UnimplementedProcessAPIServer) Execute(srv ProcessAPI_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterProcessAPIServer(s *grpc.Server, srv ProcessAPIServer) {
	s.RegisterService(&_ProcessAPI_serviceDesc, srv)
}

func _ProcessAPI_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessAPIServer).Execute(&processAPIExecuteServer{stream})
}

type ProcessAPI_ExecuteServer interface {
	Send(*Output) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type processAPIExecuteServer struct {
	grpc.ServerStream
}

func (x *processAPIExecuteServer) Send(m *Output) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processAPIExecuteServer) Recv() (*Input, error) {
	m := new(Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ProcessAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.ProcessAPI",
	HandlerType: (*ProcessAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _ProcessAPI_Execute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "process.proto",
}

// ProcessReadAPIClient is the client API for ProcessReadAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessReadAPIClient interface {
	List(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessReadAPI_ListClient, error)
	Success(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessReadAPI_SuccessClient, error)
	Failure(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessReadAPI_FailureClient, error)
	Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Process, error)
}

type processReadAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessReadAPIClient(cc grpc.ClientConnInterface) ProcessReadAPIClient {
	return &processReadAPIClient{cc}
}

func (c *processReadAPIClient) List(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessReadAPI_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessReadAPI_serviceDesc.Streams[0], "/data.ProcessReadAPI/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &processReadAPIListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessReadAPI_ListClient interface {
	Recv() (*Process, error)
	grpc.ClientStream
}

type processReadAPIListClient struct {
	grpc.ClientStream
}

func (x *processReadAPIListClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processReadAPIClient) Success(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessReadAPI_SuccessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessReadAPI_serviceDesc.Streams[1], "/data.ProcessReadAPI/Success", opts...)
	if err != nil {
		return nil, err
	}
	x := &processReadAPISuccessClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessReadAPI_SuccessClient interface {
	Recv() (*Process, error)
	grpc.ClientStream
}

type processReadAPISuccessClient struct {
	grpc.ClientStream
}

func (x *processReadAPISuccessClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processReadAPIClient) Failure(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessReadAPI_FailureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessReadAPI_serviceDesc.Streams[2], "/data.ProcessReadAPI/Failure", opts...)
	if err != nil {
		return nil, err
	}
	x := &processReadAPIFailureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessReadAPI_FailureClient interface {
	Recv() (*Process, error)
	grpc.ClientStream
}

type processReadAPIFailureClient struct {
	grpc.ClientStream
}

func (x *processReadAPIFailureClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processReadAPIClient) Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/data.ProcessReadAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessReadAPIServer is the server API for ProcessReadAPI service.
type ProcessReadAPIServer interface {
	List(*Options, ProcessReadAPI_ListServer) error
	Success(*Options, ProcessReadAPI_SuccessServer) error
	Failure(*Options, ProcessReadAPI_FailureServer) error
	Get(context.Context, *Options) (*Process, error)
}

// UnimplementedProcessReadAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProcessReadAPIServer struct {
}

func (*UnimplementedProcessReadAPIServer) List(req *Options, srv ProcessReadAPI_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedProcessReadAPIServer) Success(req *Options, srv ProcessReadAPI_SuccessServer) error {
	return status.Errorf(codes.Unimplemented, "method Success not implemented")
}
func (*UnimplementedProcessReadAPIServer) Failure(req *Options, srv ProcessReadAPI_FailureServer) error {
	return status.Errorf(codes.Unimplemented, "method Failure not implemented")
}
func (*UnimplementedProcessReadAPIServer) Get(ctx context.Context, req *Options) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterProcessReadAPIServer(s *grpc.Server, srv ProcessReadAPIServer) {
	s.RegisterService(&_ProcessReadAPI_serviceDesc, srv)
}

func _ProcessReadAPI_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessReadAPIServer).List(m, &processReadAPIListServer{stream})
}

type ProcessReadAPI_ListServer interface {
	Send(*Process) error
	grpc.ServerStream
}

type processReadAPIListServer struct {
	grpc.ServerStream
}

func (x *processReadAPIListServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessReadAPI_Success_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessReadAPIServer).Success(m, &processReadAPISuccessServer{stream})
}

type ProcessReadAPI_SuccessServer interface {
	Send(*Process) error
	grpc.ServerStream
}

type processReadAPISuccessServer struct {
	grpc.ServerStream
}

func (x *processReadAPISuccessServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessReadAPI_Failure_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessReadAPIServer).Failure(m, &processReadAPIFailureServer{stream})
}

type ProcessReadAPI_FailureServer interface {
	Send(*Process) error
	grpc.ServerStream
}

type processReadAPIFailureServer struct {
	grpc.ServerStream
}

func (x *processReadAPIFailureServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessReadAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessReadAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.ProcessReadAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessReadAPIServer).Get(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessReadAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.ProcessReadAPI",
	HandlerType: (*ProcessReadAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ProcessReadAPI_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ProcessReadAPI_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Success",
			Handler:       _ProcessReadAPI_Success_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Failure",
			Handler:       _ProcessReadAPI_Failure_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "process.proto",
}

// ProcessWriteAPIClient is the client API for ProcessWriteAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessWriteAPIClient interface {
	Record(ctx context.Context, in *Process, opts ...grpc.CallOption) (*Process, error)
}

type processWriteAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessWriteAPIClient(cc grpc.ClientConnInterface) ProcessWriteAPIClient {
	return &processWriteAPIClient{cc}
}

func (c *processWriteAPIClient) Record(ctx context.Context, in *Process, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/data.ProcessWriteAPI/Record", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessWriteAPIServer is the server API for ProcessWriteAPI service.
type ProcessWriteAPIServer interface {
	Record(context.Context, *Process) (*Process, error)
}

// UnimplementedProcessWriteAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProcessWriteAPIServer struct {
}

func (*UnimplementedProcessWriteAPIServer) Record(ctx context.Context, req *Process) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Record not implemented")
}

func RegisterProcessWriteAPIServer(s *grpc.Server, srv ProcessWriteAPIServer) {
	s.RegisterService(&_ProcessWriteAPI_serviceDesc, srv)
}

func _ProcessWriteAPI_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Process)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessWriteAPIServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.ProcessWriteAPI/Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessWriteAPIServer).Record(ctx, req.(*Process))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessWriteAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.ProcessWriteAPI",
	HandlerType: (*ProcessWriteAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Record",
			Handler:    _ProcessWriteAPI_Record_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "process.proto",
}