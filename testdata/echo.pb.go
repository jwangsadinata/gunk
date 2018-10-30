// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdata.tld/util/echo.gunk

package util

/*
package util contains a simple Echo service.
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/gunk/opt/http"
import imported "testdata.tld/util/imported"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// CheckStatusResponse is the response for a check status.
type CheckStatusResponse struct {
	Status               Status   `protobuf:"varint,1,,name=Status,proto3,enum=testdata.tld/util.Status" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStatusResponse) Reset()         { *m = CheckStatusResponse{} }
func (m *CheckStatusResponse) String() string { return proto.CompactTextString(m) }
func (*CheckStatusResponse) ProtoMessage()    {}
func (*CheckStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_c1091e0a2bee8438, []int{0}
}
func (m *CheckStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckStatusResponse.Unmarshal(m, b)
}
func (m *CheckStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckStatusResponse.Marshal(b, m, deterministic)
}
func (dst *CheckStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStatusResponse.Merge(dst, src)
}
func (m *CheckStatusResponse) XXX_Size() int {
	return xxx_messageInfo_CheckStatusResponse.Size(m)
}
func (m *CheckStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStatusResponse proto.InternalMessageInfo

func (m *CheckStatusResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Unknown
}

func init() {
	proto.RegisterType((*CheckStatusResponse)(nil), "testdata.tld/util.CheckStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UtilClient is the client API for Util service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UtilClient interface {
	// Echo echoes a message.
	Echo(ctx context.Context, in *imported.Message, opts ...grpc.CallOption) (*imported.Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error)
}

type utilClient struct {
	cc *grpc.ClientConn
}

func NewUtilClient(cc *grpc.ClientConn) UtilClient {
	return &utilClient{cc}
}

func (c *utilClient) Echo(ctx context.Context, in *imported.Message, opts ...grpc.CallOption) (*imported.Message, error) {
	out := new(imported.Message)
	err := c.cc.Invoke(ctx, "/testdata.tld/util.Util/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilClient) CheckStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := c.cc.Invoke(ctx, "/testdata.tld/util.Util/CheckStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilServer is the server API for Util service.
type UtilServer interface {
	// Echo echoes a message.
	Echo(context.Context, *imported.Message) (*imported.Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(context.Context, *empty.Empty) (*CheckStatusResponse, error)
}

func RegisterUtilServer(s *grpc.Server, srv UtilServer) {
	s.RegisterService(&_Util_serviceDesc, srv)
}

func _Util_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(imported.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.tld/util.Util/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).Echo(ctx, req.(*imported.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Util_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.tld/util.Util/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).CheckStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Util_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.tld/util.Util",
	HandlerType: (*UtilServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Util_Echo_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _Util_CheckStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdata.tld/util/echo.gunk",
}

func init() { proto.RegisterFile("testdata.tld/util/echo.gunk", fileDescriptor_echo_c1091e0a2bee8438) }

var fileDescriptor_echo_c1091e0a2bee8438 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4a, 0xc4, 0x30,
	0x14, 0xc6, 0xad, 0x96, 0x82, 0x51, 0x04, 0x23, 0x88, 0x76, 0x5c, 0x48, 0x07, 0x44, 0x5d, 0xbc,
	0xc0, 0xb8, 0x73, 0x25, 0xca, 0x2c, 0xdd, 0x28, 0x82, 0xb8, 0x10, 0xfa, 0xe7, 0x99, 0x94, 0x69,
	0xe7, 0x85, 0xc9, 0xcb, 0x62, 0xb6, 0x5e, 0xc1, 0x0b, 0x79, 0x07, 0xaf, 0xe0, 0x41, 0xa4, 0x29,
	0x82, 0xd0, 0x11, 0x66, 0x13, 0x92, 0x97, 0x1f, 0xbf, 0x7c, 0xf9, 0xc4, 0x88, 0xd1, 0x71, 0x95,
	0x73, 0x0e, 0xdc, 0x54, 0xca, 0x73, 0xdd, 0x28, 0x2c, 0x0d, 0x81, 0xf6, 0xf3, 0x99, 0xdc, 0x1f,
	0x5c, 0xa6, 0x23, 0x4d, 0xa4, 0x1b, 0x54, 0x76, 0x41, 0x4c, 0x85, 0x7f, 0x53, 0xd8, 0x5a, 0x5e,
	0x42, 0x38, 0xa6, 0x27, 0x43, 0x19, 0x2f, 0x2d, 0xba, 0x60, 0x4b, 0x33, 0x5d, 0xb3, 0xf1, 0x05,
	0x94, 0xd4, 0xaa, 0x6e, 0xa0, 0xc8, 0xb2, 0x32, 0xcc, 0x36, 0x2c, 0x3d, 0x33, 0x1e, 0x1a, 0xea,
	0xd6, 0xd2, 0x82, 0xb1, 0xea, 0x36, 0x01, 0xca, 0x6e, 0xc4, 0xc1, 0x9d, 0xc1, 0x72, 0xf6, 0xc8,
	0x39, 0x7b, 0xf7, 0x80, 0xce, 0xd2, 0xdc, 0xa1, 0xbc, 0x10, 0x49, 0x3f, 0x39, 0x8a, 0x4e, 0x37,
	0xce, 0xf7, 0x26, 0xc7, 0x30, 0x90, 0x41, 0x0f, 0x4c, 0x3e, 0x23, 0x11, 0x3f, 0x71, 0xdd, 0xc8,
	0x57, 0x11, 0x4f, 0x4b, 0x43, 0x72, 0x0c, 0xff, 0x3f, 0x0c, 0xf7, 0xe8, 0x5c, 0xae, 0x31, 0x5d,
	0x07, 0xca, 0x76, 0xdf, 0xbf, 0xbe, 0x3f, 0x36, 0x93, 0xeb, 0xe8, 0x52, 0x46, 0x4a, 0x3e, 0x8b,
	0x9d, 0x3f, 0x51, 0xe5, 0x21, 0xf4, 0xf5, 0xc1, 0x6f, 0x7d, 0x30, 0xed, 0xea, 0x4b, 0xcf, 0x56,
	0x44, 0x5d, 0xf1, 0xc5, 0x6c, 0x3b, 0xc8, 0xb7, 0x64, 0xa4, 0x6e, 0x93, 0x97, 0xb8, 0xc3, 0x8a,
	0x24, 0xa8, 0xae, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x8c, 0xa9, 0xae, 0xc9, 0x01, 0x00,
	0x00,
}