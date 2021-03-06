// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.store.svc.proto

package clientstoresvc

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

type Client struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_639564e7882e1d32, []int{0}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Client) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Client) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UpsertRequest struct {
	Client               *Client  `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertRequest) Reset()         { *m = UpsertRequest{} }
func (m *UpsertRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertRequest) ProtoMessage()    {}
func (*UpsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_639564e7882e1d32, []int{1}
}

func (m *UpsertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertRequest.Unmarshal(m, b)
}
func (m *UpsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertRequest.Marshal(b, m, deterministic)
}
func (m *UpsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertRequest.Merge(m, src)
}
func (m *UpsertRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertRequest.Size(m)
}
func (m *UpsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertRequest proto.InternalMessageInfo

func (m *UpsertRequest) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

type UpsertResponse struct {
	Processed            int64    `protobuf:"varint,1,opt,name=processed,proto3" json:"processed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertResponse) Reset()         { *m = UpsertResponse{} }
func (m *UpsertResponse) String() string { return proto.CompactTextString(m) }
func (*UpsertResponse) ProtoMessage()    {}
func (*UpsertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_639564e7882e1d32, []int{2}
}

func (m *UpsertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertResponse.Unmarshal(m, b)
}
func (m *UpsertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertResponse.Marshal(b, m, deterministic)
}
func (m *UpsertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertResponse.Merge(m, src)
}
func (m *UpsertResponse) XXX_Size() int {
	return xxx_messageInfo_UpsertResponse.Size(m)
}
func (m *UpsertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertResponse proto.InternalMessageInfo

func (m *UpsertResponse) GetProcessed() int64 {
	if m != nil {
		return m.Processed
	}
	return 0
}

func init() {
	proto.RegisterType((*Client)(nil), "client.store.svc.Client")
	proto.RegisterType((*UpsertRequest)(nil), "client.store.svc.UpsertRequest")
	proto.RegisterType((*UpsertResponse)(nil), "client.store.svc.UpsertResponse")
}

func init() { proto.RegisterFile("client.store.svc.proto", fileDescriptor_639564e7882e1d32) }

var fileDescriptor_639564e7882e1d32 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x5a, 0x03, 0x9d, 0x6a, 0x91, 0x41, 0xca, 0x22, 0x82, 0x31, 0xa7, 0x9c, 0x52,
	0xa9, 0x4f, 0xa0, 0x9e, 0xf5, 0x90, 0xe2, 0xc5, 0x8b, 0xa4, 0xdb, 0x41, 0x17, 0x9b, 0xdd, 0x75,
	0x67, 0x9b, 0xe7, 0x17, 0x67, 0x2d, 0xa2, 0xc5, 0xdb, 0xee, 0x37, 0x3f, 0xff, 0x7c, 0x0c, 0xcc,
	0xf5, 0xd6, 0x90, 0x8d, 0x0d, 0x47, 0x17, 0xa8, 0xe1, 0x41, 0x37, 0x3e, 0xb8, 0xe8, 0xf0, 0xf4,
	0x2f, 0xaf, 0x08, 0x8a, 0x7b, 0x61, 0x38, 0x83, 0xdc, 0x6c, 0x54, 0x56, 0x66, 0xf5, 0xa8, 0xcd,
	0xcd, 0x06, 0x11, 0xc6, 0xb6, 0xeb, 0x49, 0xe5, 0x65, 0x56, 0x4f, 0x5a, 0x79, 0xe3, 0x15, 0x1c,
	0xfb, 0x37, 0x67, 0xe9, 0xc5, 0xee, 0xfa, 0x35, 0x05, 0x35, 0x92, 0xd9, 0x54, 0xd8, 0xa3, 0x20,
	0x3c, 0x83, 0x23, 0xea, 0x3b, 0xb3, 0x55, 0x63, 0x99, 0xa5, 0x4f, 0x75, 0x0b, 0x27, 0x4f, 0x9e,
	0x29, 0xc4, 0x96, 0x3e, 0x76, 0xc4, 0x11, 0xaf, 0xa1, 0x48, 0x2e, 0xb2, 0x71, 0xba, 0x54, 0xcd,
	0x81, 0x72, 0xf2, 0x6a, 0xbf, 0x73, 0x55, 0x03, 0xb3, 0x7d, 0x05, 0x7b, 0x67, 0x99, 0xf0, 0x02,
	0x26, 0x3e, 0x38, 0x4d, 0xcc, 0xb4, 0x17, 0xff, 0x01, 0x4b, 0x0d, 0x98, 0x1a, 0x56, 0x5f, 0x8d,
	0x2b, 0x0a, 0x83, 0xd1, 0x84, 0x0f, 0x50, 0xa4, 0x16, 0xbc, 0x3c, 0xdc, 0xf8, 0x4b, 0xf1, 0xbc,
	0xfc, 0x3f, 0x90, 0x04, 0xea, 0xec, 0x4e, 0x3d, 0xcf, 0xfd, 0xfb, 0xeb, 0xa2, 0xf3, 0x66, 0x91,
	0xc2, 0x92, 0xe5, 0x41, 0xaf, 0x0b, 0xb9, 0xf8, 0xcd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73,
	0x0c, 0xf6, 0x00, 0x8b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientStoreServiceClient is the client API for ClientStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientStoreServiceClient interface {
	Upsert(ctx context.Context, opts ...grpc.CallOption) (ClientStoreService_UpsertClient, error)
}

type clientStoreServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientStoreServiceClient(cc *grpc.ClientConn) ClientStoreServiceClient {
	return &clientStoreServiceClient{cc}
}

func (c *clientStoreServiceClient) Upsert(ctx context.Context, opts ...grpc.CallOption) (ClientStoreService_UpsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientStoreService_serviceDesc.Streams[0], "/client.store.svc.ClientStoreService/Upsert", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStoreServiceUpsertClient{stream}
	return x, nil
}

type ClientStoreService_UpsertClient interface {
	Send(*UpsertRequest) error
	CloseAndRecv() (*UpsertResponse, error)
	grpc.ClientStream
}

type clientStoreServiceUpsertClient struct {
	grpc.ClientStream
}

func (x *clientStoreServiceUpsertClient) Send(m *UpsertRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStoreServiceUpsertClient) CloseAndRecv() (*UpsertResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpsertResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStoreServiceServer is the server API for ClientStoreService service.
type ClientStoreServiceServer interface {
	Upsert(ClientStoreService_UpsertServer) error
}

func RegisterClientStoreServiceServer(s *grpc.Server, srv ClientStoreServiceServer) {
	s.RegisterService(&_ClientStoreService_serviceDesc, srv)
}

func _ClientStoreService_Upsert_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStoreServiceServer).Upsert(&clientStoreServiceUpsertServer{stream})
}

type ClientStoreService_UpsertServer interface {
	SendAndClose(*UpsertResponse) error
	Recv() (*UpsertRequest, error)
	grpc.ServerStream
}

type clientStoreServiceUpsertServer struct {
	grpc.ServerStream
}

func (x *clientStoreServiceUpsertServer) SendAndClose(m *UpsertResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStoreServiceUpsertServer) Recv() (*UpsertRequest, error) {
	m := new(UpsertRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClientStoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.store.svc.ClientStoreService",
	HandlerType: (*ClientStoreServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upsert",
			Handler:       _ClientStoreService_Upsert_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client.store.svc.proto",
}
