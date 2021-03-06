// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package useradd

import (
	context "context"
	fmt "fmt"
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

type AddDataRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDataRequest) Reset()         { *m = AddDataRequest{} }
func (m *AddDataRequest) String() string { return proto.CompactTextString(m) }
func (*AddDataRequest) ProtoMessage()    {}
func (*AddDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *AddDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDataRequest.Unmarshal(m, b)
}
func (m *AddDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDataRequest.Marshal(b, m, deterministic)
}
func (m *AddDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDataRequest.Merge(m, src)
}
func (m *AddDataRequest) XXX_Size() int {
	return xxx_messageInfo_AddDataRequest.Size(m)
}
func (m *AddDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddDataRequest proto.InternalMessageInfo

func (m *AddDataRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AddDataResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDataResponse) Reset()         { *m = AddDataResponse{} }
func (m *AddDataResponse) String() string { return proto.CompactTextString(m) }
func (*AddDataResponse) ProtoMessage()    {}
func (*AddDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *AddDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDataResponse.Unmarshal(m, b)
}
func (m *AddDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDataResponse.Marshal(b, m, deterministic)
}
func (m *AddDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDataResponse.Merge(m, src)
}
func (m *AddDataResponse) XXX_Size() int {
	return xxx_messageInfo_AddDataResponse.Size(m)
}
func (m *AddDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddDataResponse proto.InternalMessageInfo

func (m *AddDataResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*AddDataRequest)(nil), "useradd.AddDataRequest")
	proto.RegisterType((*AddDataResponse)(nil), "useradd.AddDataResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x2d, 0x4e, 0x2d, 0x4a,
	0x4c, 0x49, 0x51, 0x52, 0xe3, 0xe2, 0x73, 0x4c, 0x49, 0x71, 0x49, 0x2c, 0x49, 0x0c, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x34, 0xb9, 0xf8, 0xe1, 0xea, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xa0, 0x2a, 0xa1, 0x3c,
	0x23, 0x4f, 0x2e, 0xf6, 0xd0, 0xe2, 0xd4, 0x22, 0xc7, 0x94, 0x14, 0x21, 0x3b, 0x2e, 0x76, 0xc7,
	0x94, 0x14, 0x10, 0x4f, 0x48, 0x5c, 0x0f, 0x6a, 0xa5, 0x1e, 0xaa, 0x7d, 0x52, 0x12, 0x98, 0x12,
	0x10, 0x0b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0xae, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa2,
	0x51, 0x39, 0xb2, 0xbe, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserAddClient is the client API for UserAdd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserAddClient interface {
	AddUser(ctx context.Context, in *AddDataRequest, opts ...grpc.CallOption) (*AddDataResponse, error)
}

type userAddClient struct {
	cc *grpc.ClientConn
}

func NewUserAddClient(cc *grpc.ClientConn) UserAddClient {
	return &userAddClient{cc}
}

func (c *userAddClient) AddUser(ctx context.Context, in *AddDataRequest, opts ...grpc.CallOption) (*AddDataResponse, error) {
	out := new(AddDataResponse)
	err := c.cc.Invoke(ctx, "/useradd.UserAdd/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAddServer is the server API for UserAdd service.
type UserAddServer interface {
	AddUser(context.Context, *AddDataRequest) (*AddDataResponse, error)
}

func RegisterUserAddServer(s *grpc.Server, srv UserAddServer) {
	s.RegisterService(&_UserAdd_serviceDesc, srv)
}

func _UserAdd_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAddServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/useradd.UserAdd/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAddServer).AddUser(ctx, req.(*AddDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAdd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "useradd.UserAdd",
	HandlerType: (*UserAddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserAdd_AddUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
