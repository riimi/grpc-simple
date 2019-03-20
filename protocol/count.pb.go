// Code generated by protoc-gen-go. DO NOT EDIT.
// source: count.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type IncrRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Sid                  string   `protobuf:"bytes,2,opt,name=sid,proto3" json:"sid,omitempty"`
	Uid                  string   `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Key                  string   `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrRequest) Reset()         { *m = IncrRequest{} }
func (m *IncrRequest) String() string { return proto.CompactTextString(m) }
func (*IncrRequest) ProtoMessage()    {}
func (*IncrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc4a427207cbe48, []int{0}
}

func (m *IncrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrRequest.Unmarshal(m, b)
}
func (m *IncrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrRequest.Marshal(b, m, deterministic)
}
func (m *IncrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrRequest.Merge(m, src)
}
func (m *IncrRequest) XXX_Size() int {
	return xxx_messageInfo_IncrRequest.Size(m)
}
func (m *IncrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncrRequest proto.InternalMessageInfo

func (m *IncrRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *IncrRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *IncrRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *IncrRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type IncrResponse struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Api                  string               `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	Code                 int32                `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Error                string               `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Count                int32                `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IncrResponse) Reset()         { *m = IncrResponse{} }
func (m *IncrResponse) String() string { return proto.CompactTextString(m) }
func (*IncrResponse) ProtoMessage()    {}
func (*IncrResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc4a427207cbe48, []int{1}
}

func (m *IncrResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrResponse.Unmarshal(m, b)
}
func (m *IncrResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrResponse.Marshal(b, m, deterministic)
}
func (m *IncrResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrResponse.Merge(m, src)
}
func (m *IncrResponse) XXX_Size() int {
	return xxx_messageInfo_IncrResponse.Size(m)
}
func (m *IncrResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncrResponse proto.InternalMessageInfo

func (m *IncrResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *IncrResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *IncrResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *IncrResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *IncrResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*IncrRequest)(nil), "protocol.IncrRequest")
	proto.RegisterType((*IncrResponse)(nil), "protocol.IncrResponse")
}

func init() { proto.RegisterFile("count.proto", fileDescriptor_bdc4a427207cbe48) }

var fileDescriptor_bdc4a427207cbe48 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x49, 0x9b, 0x20, 0x7a, 0xe9, 0x80, 0x2c, 0x40, 0x56, 0x16, 0x50, 0x26, 0x26, 0x57,
	0x2a, 0x03, 0xec, 0x0c, 0x88, 0x35, 0xc0, 0x03, 0xb4, 0xce, 0x51, 0x59, 0x6d, 0x73, 0xc1, 0x7f,
	0x90, 0xfa, 0x2a, 0x3c, 0x2d, 0xb2, 0x0f, 0x53, 0xc4, 0xe4, 0xbb, 0xcf, 0x3f, 0x9d, 0x3f, 0x1f,
	0xd4, 0x9a, 0xc2, 0xe0, 0xd5, 0x68, 0xc9, 0x93, 0x38, 0x4b, 0x87, 0xa6, 0x5d, 0x73, 0xbd, 0x21,
	0xda, 0xec, 0x70, 0x91, 0xc0, 0x3a, 0xbc, 0x2f, 0xbc, 0xd9, 0xa3, 0xf3, 0xab, 0xfd, 0xc8, 0xd1,
	0xf6, 0x0d, 0xea, 0xe7, 0x41, 0xdb, 0x0e, 0x3f, 0x02, 0x3a, 0x2f, 0xce, 0x61, 0xba, 0x1a, 0x8d,
	0x2c, 0x6e, 0x8a, 0xdb, 0x59, 0x17, 0xcb, 0x48, 0x9c, 0xe9, 0xe5, 0x84, 0x89, 0x33, 0x7d, 0x24,
	0xc1, 0xf4, 0x72, 0xca, 0x24, 0x30, 0xd9, 0xe2, 0x41, 0x96, 0x4c, 0xb6, 0x78, 0x68, 0xbf, 0x0a,
	0x98, 0xf3, 0x5c, 0x37, 0xd2, 0xe0, 0x50, 0x3c, 0xc0, 0xec, 0xf7, 0xe9, 0x34, 0xbe, 0x5e, 0x36,
	0x8a, 0xe5, 0x54, 0x96, 0x53, 0xaf, 0x39, 0xd1, 0x1d, 0xc3, 0x59, 0x69, 0x72, 0x54, 0x12, 0x50,
	0x6a, 0xea, 0x31, 0x19, 0x54, 0x5d, 0xaa, 0xc5, 0x05, 0x54, 0x68, 0x2d, 0xd9, 0x1f, 0x09, 0x6e,
	0x22, 0x4d, 0x7b, 0x91, 0x55, 0x8a, 0x72, 0xb3, 0x7c, 0x82, 0xf9, 0x63, 0x2c, 0x5e, 0xd0, 0x7e,
	0x1a, 0x8d, 0xe2, 0x1e, 0xca, 0xe8, 0x2a, 0x2e, 0x55, 0xde, 0x9b, 0xfa, 0xb3, 0x93, 0xe6, 0xea,
	0x3f, 0xe6, 0x2f, 0xb5, 0x27, 0xeb, 0xd3, 0x74, 0x71, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x88,
	0xd3, 0x54, 0x88, 0x7d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CountServiceClient is the client API for CountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountServiceClient interface {
	Incr(ctx context.Context, in *IncrRequest, opts ...grpc.CallOption) (*IncrResponse, error)
}

type countServiceClient struct {
	cc *grpc.ClientConn
}

func NewCountServiceClient(cc *grpc.ClientConn) CountServiceClient {
	return &countServiceClient{cc}
}

func (c *countServiceClient) Incr(ctx context.Context, in *IncrRequest, opts ...grpc.CallOption) (*IncrResponse, error) {
	out := new(IncrResponse)
	err := c.cc.Invoke(ctx, "/protocol.CountService/Incr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountServiceServer is the server API for CountService service.
type CountServiceServer interface {
	Incr(context.Context, *IncrRequest) (*IncrResponse, error)
}

// UnimplementedCountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCountServiceServer struct {
}

func (*UnimplementedCountServiceServer) Incr(ctx context.Context, req *IncrRequest) (*IncrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Incr not implemented")
}

func RegisterCountServiceServer(s *grpc.Server, srv CountServiceServer) {
	s.RegisterService(&_CountService_serviceDesc, srv)
}

func _CountService_Incr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountServiceServer).Incr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.CountService/Incr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountServiceServer).Incr(ctx, req.(*IncrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.CountService",
	HandlerType: (*CountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Incr",
			Handler:    _CountService_Incr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "count.proto",
}