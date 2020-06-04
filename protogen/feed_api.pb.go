// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feed_api.proto

package protov1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PostMessageRequest struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostMessageRequest) Reset()         { *m = PostMessageRequest{} }
func (m *PostMessageRequest) String() string { return proto.CompactTextString(m) }
func (*PostMessageRequest) ProtoMessage()    {}
func (*PostMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_690fbc31cc97fd09, []int{0}
}

func (m *PostMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostMessageRequest.Unmarshal(m, b)
}
func (m *PostMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostMessageRequest.Marshal(b, m, deterministic)
}
func (m *PostMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostMessageRequest.Merge(m, src)
}
func (m *PostMessageRequest) XXX_Size() int {
	return xxx_messageInfo_PostMessageRequest.Size(m)
}
func (m *PostMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostMessageRequest proto.InternalMessageInfo

func (m *PostMessageRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PostMessageRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type PostMessageResponse struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostMessageResponse) Reset()         { *m = PostMessageResponse{} }
func (m *PostMessageResponse) String() string { return proto.CompactTextString(m) }
func (*PostMessageResponse) ProtoMessage()    {}
func (*PostMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_690fbc31cc97fd09, []int{1}
}

func (m *PostMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostMessageResponse.Unmarshal(m, b)
}
func (m *PostMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostMessageResponse.Marshal(b, m, deterministic)
}
func (m *PostMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostMessageResponse.Merge(m, src)
}
func (m *PostMessageResponse) XXX_Size() int {
	return xxx_messageInfo_PostMessageResponse.Size(m)
}
func (m *PostMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostMessageResponse proto.InternalMessageInfo

func (m *PostMessageResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ListMessagesRequest struct {
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMessagesRequest) Reset()         { *m = ListMessagesRequest{} }
func (m *ListMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMessagesRequest) ProtoMessage()    {}
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_690fbc31cc97fd09, []int{2}
}

func (m *ListMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessagesRequest.Unmarshal(m, b)
}
func (m *ListMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessagesRequest.Marshal(b, m, deterministic)
}
func (m *ListMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessagesRequest.Merge(m, src)
}
func (m *ListMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMessagesRequest.Size(m)
}
func (m *ListMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessagesRequest proto.InternalMessageInfo

func (m *ListMessagesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListMessagesResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListMessagesResponse) Reset()         { *m = ListMessagesResponse{} }
func (m *ListMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMessagesResponse) ProtoMessage()    {}
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_690fbc31cc97fd09, []int{3}
}

func (m *ListMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessagesResponse.Unmarshal(m, b)
}
func (m *ListMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessagesResponse.Marshal(b, m, deterministic)
}
func (m *ListMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessagesResponse.Merge(m, src)
}
func (m *ListMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMessagesResponse.Size(m)
}
func (m *ListMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessagesResponse proto.InternalMessageInfo

func (m *ListMessagesResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*PostMessageRequest)(nil), "proto.v1.PostMessageRequest")
	proto.RegisterType((*PostMessageResponse)(nil), "proto.v1.PostMessageResponse")
	proto.RegisterType((*ListMessagesRequest)(nil), "proto.v1.ListMessagesRequest")
	proto.RegisterType((*ListMessagesResponse)(nil), "proto.v1.ListMessagesResponse")
}

func init() { proto.RegisterFile("feed_api.proto", fileDescriptor_690fbc31cc97fd09) }

var fileDescriptor_690fbc31cc97fd09 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4b, 0x4d, 0x4d,
	0x89, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x7a, 0x65,
	0x86, 0x52, 0xbc, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x10, 0x09, 0xa5, 0x28, 0x2e, 0xa1,
	0x80, 0xfc, 0xe2, 0x12, 0x5f, 0x88, 0x60, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x36,
	0x17, 0x3b, 0x54, 0x99, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xa0, 0x1e, 0xcc, 0x00, 0x3d,
	0x98, 0x52, 0x98, 0x0a, 0x21, 0x71, 0x2e, 0xf6, 0xd2, 0xe2, 0xd4, 0xa2, 0xf8, 0xcc, 0x14, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x36, 0x10, 0xd7, 0x33, 0x45, 0xc9, 0x89, 0x4b, 0x18, 0xc5,
	0xec, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x92, 0x0c, 0x57, 0xd2, 0xe3, 0x12, 0xf6, 0xc9, 0x84,
	0x9b, 0x51, 0x0c, 0x73, 0x20, 0x4e, 0x3b, 0x5d, 0xb9, 0x44, 0x50, 0xd5, 0x43, 0x2d, 0xd5, 0xe5,
	0xe2, 0x80, 0x1a, 0x59, 0x2c, 0xc1, 0xa8, 0xc0, 0x8c, 0xdd, 0x56, 0xb8, 0x12, 0xa3, 0x15, 0x8c,
	0x5c, 0xec, 0x6e, 0xa9, 0xa9, 0x29, 0x8e, 0x01, 0x9e, 0x42, 0x3e, 0x5c, 0xdc, 0x48, 0xde, 0x10,
	0x92, 0x41, 0xe8, 0xc3, 0x0c, 0x39, 0x29, 0x59, 0x1c, 0xb2, 0x10, 0x67, 0x28, 0x31, 0x08, 0xf9,
	0x73, 0xf1, 0x20, 0x3b, 0x50, 0x08, 0x49, 0x03, 0x16, 0x8f, 0x4a, 0xc9, 0xe1, 0x92, 0x86, 0x19,
	0xe8, 0xe4, 0xcc, 0xc5, 0x93, 0x9c, 0x9f, 0x0b, 0x57, 0xe6, 0xc4, 0x03, 0x76, 0x77, 0x41, 0x66,
	0x00, 0x48, 0x20, 0x80, 0x31, 0x8a, 0x1d, 0x2c, 0x53, 0x66, 0xb8, 0x88, 0x89, 0x39, 0x20, 0x22,
	0x62, 0x15, 0x13, 0x07, 0x58, 0x42, 0x2f, 0xcc, 0xf0, 0x14, 0x94, 0x19, 0x13, 0x66, 0x98, 0xc4,
	0x06, 0x56, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x5a, 0x35, 0x60, 0x38, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedAPIClient is the client API for FeedAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedAPIClient interface {
	PostMessage(ctx context.Context, in *PostMessageRequest, opts ...grpc.CallOption) (*PostMessageResponse, error)
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error)
}

type feedAPIClient struct {
	cc *grpc.ClientConn
}

func NewFeedAPIClient(cc *grpc.ClientConn) FeedAPIClient {
	return &feedAPIClient{cc}
}

func (c *feedAPIClient) PostMessage(ctx context.Context, in *PostMessageRequest, opts ...grpc.CallOption) (*PostMessageResponse, error) {
	out := new(PostMessageResponse)
	err := c.cc.Invoke(ctx, "/proto.v1.FeedAPI/PostMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedAPIClient) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error) {
	out := new(ListMessagesResponse)
	err := c.cc.Invoke(ctx, "/proto.v1.FeedAPI/ListMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedAPIServer is the server API for FeedAPI service.
type FeedAPIServer interface {
	PostMessage(context.Context, *PostMessageRequest) (*PostMessageResponse, error)
	ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
}

// UnimplementedFeedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedFeedAPIServer struct {
}

func (*UnimplementedFeedAPIServer) PostMessage(ctx context.Context, req *PostMessageRequest) (*PostMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMessage not implemented")
}
func (*UnimplementedFeedAPIServer) ListMessages(ctx context.Context, req *ListMessagesRequest) (*ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}

func RegisterFeedAPIServer(s *grpc.Server, srv FeedAPIServer) {
	s.RegisterService(&_FeedAPI_serviceDesc, srv)
}

func _FeedAPI_PostMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedAPIServer).PostMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.FeedAPI/PostMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedAPIServer).PostMessage(ctx, req.(*PostMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedAPI_ListMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedAPIServer).ListMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.FeedAPI/ListMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedAPIServer).ListMessages(ctx, req.(*ListMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.FeedAPI",
	HandlerType: (*FeedAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostMessage",
			Handler:    _FeedAPI_PostMessage_Handler,
		},
		{
			MethodName: "ListMessages",
			Handler:    _FeedAPI_ListMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feed_api.proto",
}