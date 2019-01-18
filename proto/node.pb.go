// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type RemoteType int32

const (
	RemoteType_RemoteBasic RemoteType = 0
	RemoteType_RemoteRetry RemoteType = 1
	RemoteType_RemoteForce RemoteType = 2
)

var RemoteType_name = map[int32]string{
	0: "RemoteBasic",
	1: "RemoteRetry",
	2: "RemoteForce",
}
var RemoteType_value = map[string]int32{
	"RemoteBasic": 0,
	"RemoteRetry": 1,
	"RemoteForce": 2,
}

func (x RemoteType) String() string {
	return proto.EnumName(RemoteType_name, int32(x))
}
func (RemoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_22449f59138ae649, []int{0}
}

type BackType int32

const (
	BackType_BackHTTP BackType = 0
	BackType_BackGRPC BackType = 1
)

var BackType_name = map[int32]string{
	0: "BackHTTP",
	1: "BackGRPC",
}
var BackType_value = map[string]int32{
	"BackHTTP": 0,
	"BackGRPC": 1,
}

func (x BackType) String() string {
	return proto.EnumName(BackType_name, int32(x))
}
func (BackType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_22449f59138ae649, []int{1}
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_22449f59138ae649, []int{0}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type StatusRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_22449f59138ae649, []int{1}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (dst *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(dst, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

func (m *StatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoteDownloadRequest struct {
	ObjectKey            string     `protobuf:"bytes,1,opt,name=objectKey,proto3" json:"objectKey,omitempty"`
	RemoteType           RemoteType `protobuf:"varint,2,opt,name=remoteType,proto3,enum=proto.RemoteType" json:"remoteType,omitempty"`
	BackType             BackType   `protobuf:"varint,3,opt,name=backType,proto3,enum=proto.BackType" json:"backType,omitempty"`
	BackHost             string     `protobuf:"bytes,4,opt,name=backHost,proto3" json:"backHost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RemoteDownloadRequest) Reset()         { *m = RemoteDownloadRequest{} }
func (m *RemoteDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*RemoteDownloadRequest) ProtoMessage()    {}
func (*RemoteDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_22449f59138ae649, []int{2}
}
func (m *RemoteDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteDownloadRequest.Unmarshal(m, b)
}
func (m *RemoteDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteDownloadRequest.Marshal(b, m, deterministic)
}
func (dst *RemoteDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteDownloadRequest.Merge(dst, src)
}
func (m *RemoteDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_RemoteDownloadRequest.Size(m)
}
func (m *RemoteDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteDownloadRequest proto.InternalMessageInfo

func (m *RemoteDownloadRequest) GetObjectKey() string {
	if m != nil {
		return m.ObjectKey
	}
	return ""
}

func (m *RemoteDownloadRequest) GetRemoteType() RemoteType {
	if m != nil {
		return m.RemoteType
	}
	return RemoteType_RemoteBasic
}

func (m *RemoteDownloadRequest) GetBackType() BackType {
	if m != nil {
		return m.BackType
	}
	return BackType_BackHTTP
}

func (m *RemoteDownloadRequest) GetBackHost() string {
	if m != nil {
		return m.BackHost
	}
	return ""
}

type ServiceReply struct {
	Code                 int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Detail               *ReplyDetail `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ServiceReply) Reset()         { *m = ServiceReply{} }
func (m *ServiceReply) String() string { return proto.CompactTextString(m) }
func (*ServiceReply) ProtoMessage()    {}
func (*ServiceReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_22449f59138ae649, []int{3}
}
func (m *ServiceReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceReply.Unmarshal(m, b)
}
func (m *ServiceReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceReply.Marshal(b, m, deterministic)
}
func (dst *ServiceReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceReply.Merge(dst, src)
}
func (m *ServiceReply) XXX_Size() int {
	return xxx_messageInfo_ServiceReply.Size(m)
}
func (m *ServiceReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceReply.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceReply proto.InternalMessageInfo

func (m *ServiceReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ServiceReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ServiceReply) GetDetail() *ReplyDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

type ReplyDetail struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Json                 string   `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyDetail) Reset()         { *m = ReplyDetail{} }
func (m *ReplyDetail) String() string { return proto.CompactTextString(m) }
func (*ReplyDetail) ProtoMessage()    {}
func (*ReplyDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_22449f59138ae649, []int{4}
}
func (m *ReplyDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyDetail.Unmarshal(m, b)
}
func (m *ReplyDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyDetail.Marshal(b, m, deterministic)
}
func (dst *ReplyDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyDetail.Merge(dst, src)
}
func (m *ReplyDetail) XXX_Size() int {
	return xxx_messageInfo_ReplyDetail.Size(m)
}
func (m *ReplyDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyDetail proto.InternalMessageInfo

func (m *ReplyDetail) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ReplyDetail) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*StatusRequest)(nil), "proto.StatusRequest")
	proto.RegisterType((*RemoteDownloadRequest)(nil), "proto.RemoteDownloadRequest")
	proto.RegisterType((*ServiceReply)(nil), "proto.ServiceReply")
	proto.RegisterType((*ReplyDetail)(nil), "proto.ReplyDetail")
	proto.RegisterEnum("proto.RemoteType", RemoteType_name, RemoteType_value)
	proto.RegisterEnum("proto.BackType", BackType_name, BackType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	RemoteDownload(ctx context.Context, in *RemoteDownloadRequest, opts ...grpc.CallOption) (*ServiceReply, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*ServiceReply, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) RemoteDownload(ctx context.Context, in *RemoteDownloadRequest, opts ...grpc.CallOption) (*ServiceReply, error) {
	out := new(ServiceReply)
	err := c.cc.Invoke(ctx, "/proto.NodeService/RemoteDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*ServiceReply, error) {
	out := new(ServiceReply)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	RemoteDownload(context.Context, *RemoteDownloadRequest) (*ServiceReply, error)
	Status(context.Context, *StatusRequest) (*ServiceReply, error)
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_RemoteDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).RemoteDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/RemoteDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).RemoteDownload(ctx, req.(*RemoteDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoteDownload",
			Handler:    _NodeService_RemoteDownload_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _NodeService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_node_22449f59138ae649) }

var fileDescriptor_node_22449f59138ae649 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xed, 0x6e, 0xd3, 0x40,
	0x10, 0xcc, 0x85, 0x36, 0x24, 0x9b, 0x92, 0x9a, 0x05, 0x24, 0xab, 0xaa, 0x44, 0xe5, 0x1f, 0xa8,
	0x0a, 0x92, 0xa5, 0x06, 0xf1, 0x1b, 0x29, 0x04, 0x68, 0x85, 0x84, 0xa2, 0x6b, 0x5e, 0xc0, 0xb9,
	0x5b, 0xa5, 0x2e, 0x8e, 0xd7, 0xf8, 0x2e, 0x20, 0x3f, 0x03, 0x4f, 0xc3, 0x1b, 0x22, 0x9f, 0x3f,
	0x83, 0xc4, 0xaf, 0xec, 0xce, 0xce, 0x64, 0xc7, 0xb3, 0x07, 0x90, 0xb2, 0xa6, 0x30, 0xcb, 0xd9,
	0x32, 0x9e, 0xba, 0x9f, 0x00, 0x60, 0x2c, 0xc9, 0x64, 0x9c, 0x1a, 0x0a, 0x5e, 0xc3, 0xb3, 0x7b,
	0x1b, 0xd9, 0x83, 0x91, 0xf4, 0xe3, 0x40, 0xc6, 0xe2, 0x0c, 0x86, 0xb1, 0xf6, 0xc5, 0x95, 0xb8,
	0x9e, 0xc8, 0x61, 0xac, 0x83, 0x3f, 0x02, 0x5e, 0x49, 0xda, 0xb3, 0xa5, 0x15, 0xff, 0x4a, 0x13,
	0x8e, 0x74, 0xc3, 0xbc, 0x84, 0x09, 0x6f, 0x1f, 0x49, 0xd9, 0xaf, 0x54, 0xd4, 0x82, 0x0e, 0xc0,
	0x1b, 0x80, 0xdc, 0xc9, 0x36, 0x45, 0x46, 0xfe, 0xf0, 0x4a, 0x5c, 0xcf, 0x16, 0xcf, 0x2b, 0x1f,
	0xa1, 0x6c, 0x07, 0xb2, 0x47, 0xc2, 0xb7, 0x30, 0xde, 0x46, 0xea, 0xbb, 0x13, 0x3c, 0x71, 0x82,
	0xf3, 0x5a, 0xb0, 0xac, 0x61, 0xd9, 0x12, 0xf0, 0xa2, 0x22, 0xdf, 0xb2, 0xb1, 0xfe, 0x89, 0x5b,
	0xde, 0xf6, 0xc1, 0x03, 0x9c, 0xdd, 0x53, 0xfe, 0x33, 0x56, 0x24, 0x29, 0x4b, 0x0a, 0x44, 0x38,
	0x51, 0xac, 0xc9, 0x99, 0x3c, 0x95, 0xae, 0x46, 0x1f, 0x9e, 0xee, 0xc9, 0x98, 0x68, 0x57, 0x99,
	0x9b, 0xc8, 0xa6, 0xc5, 0x39, 0x8c, 0x34, 0xd9, 0x28, 0x4e, 0x9c, 0x89, 0xe9, 0x02, 0x5b, 0xd7,
	0x59, 0x52, 0xac, 0xdc, 0x44, 0xd6, 0x8c, 0xe0, 0x06, 0xa6, 0x3d, 0xb8, 0x0c, 0xef, 0x6e, 0xd5,
	0x84, 0x77, 0xb7, 0x2a, 0x17, 0x3f, 0x1a, 0x4e, 0xeb, 0x0d, 0xae, 0x9e, 0x7f, 0x00, 0xe8, 0xbe,
	0x1f, 0xcf, 0xcb, 0x3f, 0x28, 0xbb, 0x65, 0x64, 0x62, 0xe5, 0x0d, 0x3a, 0x40, 0x92, 0xcd, 0x0b,
	0x4f, 0x74, 0xc0, 0x67, 0xce, 0x15, 0x79, 0xc3, 0xf9, 0x1b, 0x18, 0x37, 0x79, 0xe0, 0x59, 0x55,
	0xdf, 0x6e, 0x36, 0x6b, 0x6f, 0xd0, 0x74, 0x5f, 0xe4, 0xfa, 0xa3, 0x27, 0x16, 0xbf, 0x05, 0x4c,
	0xbf, 0xb1, 0xa6, 0x3a, 0x0a, 0xfc, 0x04, 0xb3, 0xe3, 0x43, 0xe2, 0xe5, 0xd1, 0x3d, 0xfe, 0xb9,
	0xef, 0xc5, 0x8b, 0x7a, 0xda, 0x8f, 0x32, 0x18, 0xe0, 0x7b, 0x18, 0x55, 0x2f, 0x06, 0x5f, 0x36,
	0x84, 0xfe, 0x03, 0xfa, 0x8f, 0x6c, 0x19, 0x82, 0xaf, 0x78, 0x1f, 0xee, 0x62, 0xfb, 0x70, 0xd8,
	0x86, 0x3b, 0xd6, 0x8a, 0xd3, 0x5d, 0x45, 0x5d, 0x7a, 0x3d, 0x9b, 0xeb, 0x12, 0x59, 0x8b, 0xed,
	0xc8, 0x8d, 0xde, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x0f, 0xfe, 0xec, 0xc0, 0x02, 0x00,
	0x00,
}