// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

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

type Entity struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string              `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Transaction          string              `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Data                 map[string]*any.Any `protobuf:"bytes,10,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{0}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Entity) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

func (m *Entity) GetData() map[string]*any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Entity)(nil), "data.Entity")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.Entity.DataEntry")
}

func init() {
	proto.RegisterFile("entity.proto", fileDescriptor_cf50d946d740d100)
}

var fileDescriptor_cf50d946d740d100 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0xe5, 0xec, 0x12, 0xd4, 0xc9, 0xb6, 0xa2, 0xa3, 0x6a, 0x49, 0x43, 0x25, 0x96, 0x9c,
	0xaa, 0x1c, 0x9c, 0x2a, 0x5c, 0xd0, 0xde, 0x76, 0x4b, 0x55, 0x21, 0x81, 0x58, 0x45, 0x42, 0x9c,
	0xdd, 0xc4, 0x14, 0x8b, 0xc6, 0x8e, 0x12, 0x67, 0xa5, 0x08, 0x71, 0xe1, 0x15, 0x38, 0x70, 0xe2,
	0xa9, 0x78, 0x04, 0x78, 0x10, 0x94, 0xf1, 0x16, 0x05, 0x7a, 0xe8, 0xcd, 0xf3, 0x7b, 0xfe, 0xcf,
	0xff, 0x8c, 0x61, 0x26, 0xb5, 0x55, 0xb6, 0xe7, 0x75, 0x63, 0xac, 0xc1, 0x69, 0x29, 0xac, 0x88,
	0x4e, 0xae, 0x8d, 0xb9, 0xbe, 0x91, 0xa9, 0xa8, 0x55, 0x2a, 0xb4, 0x36, 0x56, 0x58, 0x65, 0x74,
	0xeb, 0x7a, 0xa2, 0xe3, 0xdd, 0x2d, 0x55, 0x57, 0xdd, 0x87, 0x54, 0xe8, 0x9d, 0x3d, 0x9a, 0x15,
	0xa6, 0xaa, 0x8c, 0x76, 0x55, 0xfc, 0x8b, 0x81, 0x7f, 0x41, 0x74, 0x3c, 0x00, 0x4f, 0x95, 0x21,
	0x5b, 0xb0, 0xd3, 0xbd, 0xdc, 0x53, 0x25, 0x22, 0x4c, 0xb5, 0xa8, 0x64, 0xe8, 0x91, 0x42, 0x67,
	0x0c, 0xe1, 0xe1, 0x56, 0x36, 0xad, 0x32, 0x3a, 0x9c, 0x90, 0x7c, 0x5b, 0xe2, 0x02, 0x02, 0xdb,
	0x08, 0xdd, 0x8a, 0x62, 0xc8, 0x11, 0x4e, 0xe9, 0x76, 0x2c, 0x61, 0x02, 0x94, 0x3c, 0x84, 0xc5,
	0xe4, 0x34, 0xc8, 0xe6, 0x7c, 0x28, 0xb8, 0x7b, 0x9b, 0xbf, 0x14, 0x56, 0x5c, 0x68, 0xdb, 0xf4,
	0xb9, 0x9b, 0xee, 0x0d, 0xec, 0xfd, 0x95, 0xf0, 0x11, 0x4c, 0x3e, 0xc9, 0x7e, 0x97, 0x6c, 0x38,
	0x62, 0x02, 0x0f, 0xb6, 0xe2, 0xa6, 0x73, 0xd9, 0x82, 0xec, 0x88, 0xbb, 0x71, 0xf9, 0xed, 0xb8,
	0x7c, 0xa5, 0xfb, 0xdc, 0xb5, 0x2c, 0xbd, 0x17, 0x2c, 0xfb, 0xce, 0x60, 0xdf, 0xbd, 0x94, 0x4b,
	0x51, 0xae, 0x36, 0xaf, 0x70, 0x05, 0xd3, 0xd7, 0xaa, 0xb5, 0xb8, 0xef, 0x62, 0xbc, 0xad, 0x69,
	0x7b, 0xd1, 0x6c, 0x9c, 0x2a, 0x8e, 0xbe, 0xfe, 0xfc, 0xfd, 0xcd, 0x3b, 0x42, 0x4c, 0xdd, 0x07,
	0xa4, 0xdb, 0xb3, 0x54, 0xe9, 0xb6, 0x96, 0x85, 0x3d, 0x63, 0x78, 0x0e, 0x93, 0x4b, 0x79, 0x0f,
	0xe1, 0x29, 0x11, 0x8e, 0xf1, 0xf1, 0x5d, 0x42, 0xfa, 0x59, 0x95, 0x5f, 0xb2, 0x1f, 0x0c, 0x0e,
	0x5c, 0xef, 0xfb, 0x46, 0x59, 0x39, 0x44, 0x5b, 0x83, 0x7f, 0xde, 0x48, 0x61, 0x25, 0xfe, 0xc3,
	0xfa, 0x8f, 0x7c, 0x42, 0xe4, 0x79, 0x7c, 0x38, 0x22, 0x17, 0x64, 0x5b, 0xb2, 0x04, 0x2f, 0xc1,
	0x7f, 0x57, 0x97, 0xf7, 0x31, 0x9e, 0x11, 0xe3, 0x49, 0x34, 0x1f, 0x31, 0x3a, 0xb2, 0x51, 0xb8,
	0x25, 0x4b, 0xd6, 0x31, 0x1c, 0x16, 0xa6, 0xe2, 0xf5, 0x47, 0xd5, 0x76, 0xca, 0x4a, 0xb2, 0xaf,
	0x03, 0xe7, 0xdf, 0x0c, 0xdb, 0xde, 0xb0, 0x2b, 0x9f, 0xd6, 0xfe, 0xfc, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x07, 0x51, 0xb9, 0x10, 0xa7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EntityReadAPIClient is the client API for EntityReadAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntityReadAPIClient interface {
	List(ctx context.Context, in *Options, opts ...grpc.CallOption) (EntityReadAPI_ListClient, error)
	Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Entity, error)
}

type entityReadAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityReadAPIClient(cc grpc.ClientConnInterface) EntityReadAPIClient {
	return &entityReadAPIClient{cc}
}

func (c *entityReadAPIClient) List(ctx context.Context, in *Options, opts ...grpc.CallOption) (EntityReadAPI_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EntityReadAPI_serviceDesc.Streams[0], "/data.EntityReadAPI/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &entityReadAPIListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EntityReadAPI_ListClient interface {
	Recv() (*Entity, error)
	grpc.ClientStream
}

type entityReadAPIListClient struct {
	grpc.ClientStream
}

func (x *entityReadAPIListClient) Recv() (*Entity, error) {
	m := new(Entity)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *entityReadAPIClient) Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/data.EntityReadAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityReadAPIServer is the server API for EntityReadAPI service.
type EntityReadAPIServer interface {
	List(*Options, EntityReadAPI_ListServer) error
	Get(context.Context, *Options) (*Entity, error)
}

// UnimplementedEntityReadAPIServer can be embedded to have forward compatible implementations.
type UnimplementedEntityReadAPIServer struct {
}

func (*UnimplementedEntityReadAPIServer) List(req *Options, srv EntityReadAPI_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEntityReadAPIServer) Get(ctx context.Context, req *Options) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterEntityReadAPIServer(s *grpc.Server, srv EntityReadAPIServer) {
	s.RegisterService(&_EntityReadAPI_serviceDesc, srv)
}

func _EntityReadAPI_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EntityReadAPIServer).List(m, &entityReadAPIListServer{stream})
}

type EntityReadAPI_ListServer interface {
	Send(*Entity) error
	grpc.ServerStream
}

type entityReadAPIListServer struct {
	grpc.ServerStream
}

func (x *entityReadAPIListServer) Send(m *Entity) error {
	return x.ServerStream.SendMsg(m)
}

func _EntityReadAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityReadAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.EntityReadAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityReadAPIServer).Get(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntityReadAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.EntityReadAPI",
	HandlerType: (*EntityReadAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _EntityReadAPI_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _EntityReadAPI_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entity.proto",
}

// EntityWriteAPIClient is the client API for EntityWriteAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntityWriteAPIClient interface {
	Create(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error)
	Update(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error)
}

type entityWriteAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityWriteAPIClient(cc grpc.ClientConnInterface) EntityWriteAPIClient {
	return &entityWriteAPIClient{cc}
}

func (c *entityWriteAPIClient) Create(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/data.EntityWriteAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityWriteAPIClient) Update(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/data.EntityWriteAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityWriteAPIServer is the server API for EntityWriteAPI service.
type EntityWriteAPIServer interface {
	Create(context.Context, *Entity) (*Entity, error)
	Update(context.Context, *Entity) (*Entity, error)
}

// UnimplementedEntityWriteAPIServer can be embedded to have forward compatible implementations.
type UnimplementedEntityWriteAPIServer struct {
}

func (*UnimplementedEntityWriteAPIServer) Create(ctx context.Context, req *Entity) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEntityWriteAPIServer) Update(ctx context.Context, req *Entity) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterEntityWriteAPIServer(s *grpc.Server, srv EntityWriteAPIServer) {
	s.RegisterService(&_EntityWriteAPI_serviceDesc, srv)
}

func _EntityWriteAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityWriteAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.EntityWriteAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityWriteAPIServer).Create(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityWriteAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityWriteAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.EntityWriteAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityWriteAPIServer).Update(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntityWriteAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.EntityWriteAPI",
	HandlerType: (*EntityWriteAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EntityWriteAPI_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EntityWriteAPI_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entity.proto",
}