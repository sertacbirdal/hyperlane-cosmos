// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyperlane/mailbox/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgDispatch defines the request type for the Dispatch rpc.
type MsgDispatch struct {
	Sender            string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	DestinationDomain uint32 `protobuf:"varint,2,opt,name=destination_domain,json=destinationDomain,proto3" json:"destination_domain,omitempty"`
	RecipientAddress  string `protobuf:"bytes,3,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	MessageBody       string `protobuf:"bytes,4,opt,name=message_body,json=messageBody,proto3" json:"message_body,omitempty"`
}

func (m *MsgDispatch) Reset()         { *m = MsgDispatch{} }
func (m *MsgDispatch) String() string { return proto.CompactTextString(m) }
func (*MsgDispatch) ProtoMessage()    {}
func (*MsgDispatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{0}
}

func (m *MsgDispatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgDispatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDispatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgDispatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDispatch.Merge(m, src)
}

func (m *MsgDispatch) XXX_Size() int {
	return m.Size()
}

func (m *MsgDispatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDispatch.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDispatch proto.InternalMessageInfo

// MsgDispatchResponse defines the Dispatch response type.
type MsgDispatchResponse struct {
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (m *MsgDispatchResponse) Reset()         { *m = MsgDispatchResponse{} }
func (m *MsgDispatchResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDispatchResponse) ProtoMessage()    {}
func (*MsgDispatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{1}
}

func (m *MsgDispatchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgDispatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDispatchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgDispatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDispatchResponse.Merge(m, src)
}

func (m *MsgDispatchResponse) XXX_Size() int {
	return m.Size()
}

func (m *MsgDispatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDispatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDispatchResponse proto.InternalMessageInfo

func (m *MsgDispatchResponse) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

// MsgProcess defines the request type for the Process rpc.
type MsgProcess struct {
	Sender   string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Metadata string `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *MsgProcess) Reset()         { *m = MsgProcess{} }
func (m *MsgProcess) String() string { return proto.CompactTextString(m) }
func (*MsgProcess) ProtoMessage()    {}
func (*MsgProcess) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{2}
}

func (m *MsgProcess) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgProcess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProcess.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgProcess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProcess.Merge(m, src)
}

func (m *MsgProcess) XXX_Size() int {
	return m.Size()
}

func (m *MsgProcess) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProcess.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProcess proto.InternalMessageInfo

// MsgProcessResponse defines the Process response type.
type MsgProcessResponse struct{}

func (m *MsgProcessResponse) Reset()         { *m = MsgProcessResponse{} }
func (m *MsgProcessResponse) String() string { return proto.CompactTextString(m) }
func (*MsgProcessResponse) ProtoMessage()    {}
func (*MsgProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{3}
}

func (m *MsgProcessResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProcessResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProcessResponse.Merge(m, src)
}

func (m *MsgProcessResponse) XXX_Size() int {
	return m.Size()
}

func (m *MsgProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProcessResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDispatch)(nil), "hyperlane.mailbox.v1.MsgDispatch")
	proto.RegisterType((*MsgDispatchResponse)(nil), "hyperlane.mailbox.v1.MsgDispatchResponse")
	proto.RegisterType((*MsgProcess)(nil), "hyperlane.mailbox.v1.MsgProcess")
	proto.RegisterType((*MsgProcessResponse)(nil), "hyperlane.mailbox.v1.MsgProcessResponse")
}

func init() { proto.RegisterFile("hyperlane/mailbox/v1/tx.proto", fileDescriptor_544026924916c03f) }

var fileDescriptor_544026924916c03f = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x14, 0xda, 0xe6, 0x0a, 0x12, 0x39, 0x22, 0x61, 0x0c, 0x75, 0xd3, 0x4c, 0x21,
	0x28, 0x36, 0xa5, 0x9d, 0x3a, 0x41, 0xd4, 0x85, 0x21, 0x12, 0x32, 0x42, 0x42, 0x5d, 0xa2, 0x8b,
	0xef, 0x74, 0x39, 0x29, 0xbe, 0xb3, 0xfc, 0xae, 0x51, 0xb2, 0x21, 0x26, 0xc4, 0xc4, 0x47, 0xe8,
	0xca, 0x96, 0x81, 0x95, 0x9d, 0xb1, 0x62, 0x62, 0x44, 0xc9, 0x10, 0x3e, 0x00, 0x1f, 0x00, 0xc5,
	0x3e, 0x9b, 0x48, 0xb4, 0xaa, 0x58, 0xa2, 0xfc, 0xdf, 0xff, 0xef, 0xf7, 0xee, 0xfd, 0xec, 0xc3,
	0xbb, 0xc3, 0x69, 0xc2, 0xd3, 0x11, 0x55, 0x3c, 0x88, 0xa9, 0x1c, 0x0d, 0xf4, 0x24, 0x18, 0x1f,
	0x04, 0x66, 0xe2, 0x27, 0xa9, 0x36, 0x9a, 0xd4, 0x4b, 0xdb, 0xb7, 0xb6, 0x3f, 0x3e, 0x70, 0x6b,
	0x34, 0x96, 0x4a, 0x07, 0xd9, 0x6f, 0x1e, 0x74, 0xef, 0x47, 0x1a, 0x62, 0x0d, 0x41, 0x0c, 0x62,
	0xd5, 0x20, 0x06, 0x61, 0x8d, 0x07, 0xb9, 0xd1, 0xcf, 0x54, 0x90, 0x0b, 0x6b, 0xd5, 0x85, 0x16,
	0x3a, 0xaf, 0xaf, 0xfe, 0xe5, 0xd5, 0xe6, 0x6f, 0x84, 0x77, 0x7a, 0x20, 0x4e, 0x24, 0x24, 0xd4,
	0x44, 0x43, 0xf2, 0x14, 0x6f, 0x02, 0x57, 0x8c, 0xa7, 0x0e, 0x6a, 0xa0, 0x56, 0xb5, 0xeb, 0x7c,
	0xff, 0xd2, 0xa9, 0xdb, 0x3e, 0x2f, 0x18, 0x4b, 0x39, 0xc0, 0x6b, 0x93, 0x4a, 0x25, 0x42, 0x9b,
	0x23, 0x1d, 0x4c, 0x18, 0x07, 0x23, 0x15, 0x35, 0x52, 0xab, 0x3e, 0xd3, 0x31, 0x95, 0xca, 0xb9,
	0xd1, 0x40, 0xad, 0x3b, 0x61, 0x6d, 0xcd, 0x39, 0xc9, 0x0c, 0xf2, 0x04, 0xd7, 0x52, 0x1e, 0xc9,
	0x44, 0x72, 0x65, 0xfa, 0x34, 0xef, 0xe8, 0x6c, 0xac, 0x66, 0x85, 0x77, 0x4b, 0xc3, 0x4e, 0x22,
	0xfb, 0xf8, 0x76, 0xcc, 0x01, 0xa8, 0xe0, 0xfd, 0x81, 0x66, 0x53, 0xe7, 0x66, 0x96, 0xdb, 0xb1,
	0xb5, 0xae, 0x66, 0xd3, 0xe3, 0xa3, 0x0f, 0xe7, 0x7b, 0x95, 0x5f, 0xe7, 0x7b, 0x95, 0xf7, 0xcb,
	0x59, 0xdb, 0x9e, 0xe9, 0xe3, 0x72, 0xd6, 0x7e, 0x54, 0xb2, 0xec, 0x58, 0x58, 0x6b, 0x6b, 0x36,
	0x8f, 0xf0, 0xbd, 0x35, 0x19, 0x72, 0x48, 0xb4, 0x02, 0x4e, 0x76, 0x31, 0x2e, 0xe6, 0x49, 0x96,
	0x13, 0x08, 0xab, 0xb6, 0xf2, 0x92, 0x35, 0x3f, 0x23, 0x8c, 0x7b, 0x20, 0x5e, 0xa5, 0x3a, 0x5a,
	0x9d, 0xee, 0xff, 0x59, 0xb9, 0x78, 0x3b, 0xe6, 0x86, 0x32, 0x6a, 0x68, 0x46, 0xa8, 0x1a, 0x96,
	0x9a, 0x38, 0x78, 0xcb, 0x4e, 0xb2, 0x38, 0x0a, 0x79, 0x7c, 0x78, 0xc5, 0x8a, 0x0f, 0x2f, 0x5b,
	0xd1, 0x1e, 0xae, 0x59, 0xc7, 0xe4, 0xaf, 0x2a, 0x16, 0x7c, 0xf6, 0x15, 0xe1, 0x8d, 0x1e, 0x08,
	0xf2, 0x16, 0x6f, 0x97, 0xaf, 0x7c, 0xdf, 0xbf, 0xec, 0xb3, 0xf3, 0xd7, 0xf8, 0xb8, 0x8f, 0xaf,
	0x8d, 0x94, 0x08, 0xdf, 0xe0, 0xad, 0x82, 0x4f, 0xe3, 0xca, 0xa7, 0x6c, 0xc2, 0x6d, 0x5d, 0x97,
	0x28, 0xda, 0xba, 0xb7, 0xde, 0x2d, 0x67, 0x6d, 0xd4, 0x3d, 0xfd, 0x36, 0xf7, 0xd0, 0xc5, 0xdc,
	0x43, 0x3f, 0xe7, 0x1e, 0xfa, 0xb4, 0xf0, 0x2a, 0x17, 0x0b, 0xaf, 0xf2, 0x63, 0xe1, 0x55, 0x4e,
	0x9f, 0x0b, 0x69, 0x86, 0x67, 0x03, 0x3f, 0xd2, 0x71, 0x00, 0x26, 0xa5, 0x4a, 0xf0, 0x91, 0x1e,
	0xf3, 0xce, 0x98, 0x2b, 0x73, 0x96, 0x72, 0x08, 0xfe, 0x81, 0x35, 0x29, 0xef, 0xa0, 0x99, 0x26,
	0x1c, 0x06, 0x9b, 0xd9, 0x8d, 0x38, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xe6, 0xae, 0xd4,
	0xa5, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Dispatch sends interchain messages
	Dispatch(ctx context.Context, in *MsgDispatch, opts ...grpc.CallOption) (*MsgDispatchResponse, error)
	// Process delivers interchain messages
	Process(ctx context.Context, in *MsgProcess, opts ...grpc.CallOption) (*MsgProcessResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Dispatch(ctx context.Context, in *MsgDispatch, opts ...grpc.CallOption) (*MsgDispatchResponse, error) {
	out := new(MsgDispatchResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.mailbox.v1.Msg/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Process(ctx context.Context, in *MsgProcess, opts ...grpc.CallOption) (*MsgProcessResponse, error) {
	out := new(MsgProcessResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.mailbox.v1.Msg/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Dispatch sends interchain messages
	Dispatch(context.Context, *MsgDispatch) (*MsgDispatchResponse, error)
	// Process delivers interchain messages
	Process(context.Context, *MsgProcess) (*MsgProcessResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct{}

func (*UnimplementedMsgServer) Dispatch(ctx context.Context, req *MsgDispatch) (*MsgDispatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}

func (*UnimplementedMsgServer) Process(ctx context.Context, req *MsgProcess) (*MsgProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDispatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyperlane.mailbox.v1.Msg/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Dispatch(ctx, req.(*MsgDispatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProcess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyperlane.mailbox.v1.Msg/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Process(ctx, req.(*MsgProcess))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hyperlane.mailbox.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dispatch",
			Handler:    _Msg_Dispatch_Handler,
		},
		{
			MethodName: "Process",
			Handler:    _Msg_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyperlane/mailbox/v1/tx.proto",
}

func (m *MsgDispatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDispatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDispatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageBody) > 0 {
		i -= len(m.MessageBody)
		copy(dAtA[i:], m.MessageBody)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MessageBody)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DestinationDomain != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.DestinationDomain))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDispatchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDispatchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDispatchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageId) > 0 {
		i -= len(m.MessageId)
		copy(dAtA[i:], m.MessageId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MessageId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgProcess) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProcess) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProcess) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgProcessResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProcessResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProcessResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *MsgDispatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.DestinationDomain != 0 {
		n += 1 + sovTx(uint64(m.DestinationDomain))
	}
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MessageBody)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDispatchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgProcess) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgProcessResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *MsgDispatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDispatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDispatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationDomain", wireType)
			}
			m.DestinationDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestinationDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageBody", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageBody = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgDispatchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDispatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDispatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgProcess) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgProcess: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProcess: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgProcessResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgProcessResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProcessResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
