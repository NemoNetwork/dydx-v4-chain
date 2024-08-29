// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo_network/sending/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgCreateTransfer is a request type used for initiating new transfers.
type MsgCreateTransfer struct {
	Transfer *Transfer `protobuf:"bytes,1,opt,name=transfer,proto3" json:"transfer,omitempty"`
}

func (m *MsgCreateTransfer) Reset()         { *m = MsgCreateTransfer{} }
func (m *MsgCreateTransfer) String() string { return proto.CompactTextString(m) }
func (*MsgCreateTransfer) ProtoMessage()    {}
func (*MsgCreateTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ca7f6311e9962d, []int{0}
}
func (m *MsgCreateTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateTransfer.Merge(m, src)
}
func (m *MsgCreateTransfer) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateTransfer proto.InternalMessageInfo

func (m *MsgCreateTransfer) GetTransfer() *Transfer {
	if m != nil {
		return m.Transfer
	}
	return nil
}

// MsgCreateTransferResponse is a response type used for new transfers.
type MsgCreateTransferResponse struct {
}

func (m *MsgCreateTransferResponse) Reset()         { *m = MsgCreateTransferResponse{} }
func (m *MsgCreateTransferResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateTransferResponse) ProtoMessage()    {}
func (*MsgCreateTransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ca7f6311e9962d, []int{1}
}
func (m *MsgCreateTransferResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateTransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateTransferResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateTransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateTransferResponse.Merge(m, src)
}
func (m *MsgCreateTransferResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateTransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateTransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateTransferResponse proto.InternalMessageInfo

// MsgDepositToSubaccountResponse is a response type used for new
// account-to-subaccount transfers.
type MsgDepositToSubaccountResponse struct {
}

func (m *MsgDepositToSubaccountResponse) Reset()         { *m = MsgDepositToSubaccountResponse{} }
func (m *MsgDepositToSubaccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDepositToSubaccountResponse) ProtoMessage()    {}
func (*MsgDepositToSubaccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ca7f6311e9962d, []int{2}
}
func (m *MsgDepositToSubaccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositToSubaccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositToSubaccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositToSubaccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositToSubaccountResponse.Merge(m, src)
}
func (m *MsgDepositToSubaccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositToSubaccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositToSubaccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositToSubaccountResponse proto.InternalMessageInfo

// MsgWithdrawFromSubaccountResponse is a response type used for new
// subaccount-to-account transfers.
type MsgWithdrawFromSubaccountResponse struct {
}

func (m *MsgWithdrawFromSubaccountResponse) Reset()         { *m = MsgWithdrawFromSubaccountResponse{} }
func (m *MsgWithdrawFromSubaccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawFromSubaccountResponse) ProtoMessage()    {}
func (*MsgWithdrawFromSubaccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ca7f6311e9962d, []int{3}
}
func (m *MsgWithdrawFromSubaccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawFromSubaccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawFromSubaccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawFromSubaccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawFromSubaccountResponse.Merge(m, src)
}
func (m *MsgWithdrawFromSubaccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawFromSubaccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawFromSubaccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawFromSubaccountResponse proto.InternalMessageInfo

// MsgSendFromModuleToAccountResponse is a response type used for new
// module-to-account transfers.
type MsgSendFromModuleToAccountResponse struct {
}

func (m *MsgSendFromModuleToAccountResponse) Reset()         { *m = MsgSendFromModuleToAccountResponse{} }
func (m *MsgSendFromModuleToAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendFromModuleToAccountResponse) ProtoMessage()    {}
func (*MsgSendFromModuleToAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ca7f6311e9962d, []int{4}
}
func (m *MsgSendFromModuleToAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendFromModuleToAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendFromModuleToAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendFromModuleToAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendFromModuleToAccountResponse.Merge(m, src)
}
func (m *MsgSendFromModuleToAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendFromModuleToAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendFromModuleToAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendFromModuleToAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateTransfer)(nil), "nemo_network.sending.MsgCreateTransfer")
	proto.RegisterType((*MsgCreateTransferResponse)(nil), "nemo_network.sending.MsgCreateTransferResponse")
	proto.RegisterType((*MsgDepositToSubaccountResponse)(nil), "nemo_network.sending.MsgDepositToSubaccountResponse")
	proto.RegisterType((*MsgWithdrawFromSubaccountResponse)(nil), "nemo_network.sending.MsgWithdrawFromSubaccountResponse")
	proto.RegisterType((*MsgSendFromModuleToAccountResponse)(nil), "nemo_network.sending.MsgSendFromModuleToAccountResponse")
}

func init() { proto.RegisterFile("nemo_network/sending/tx.proto", fileDescriptor_85ca7f6311e9962d) }

var fileDescriptor_85ca7f6311e9962d = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4b, 0x4b, 0xfb, 0x40,
	0x14, 0xc5, 0x1b, 0x0a, 0x7f, 0xfe, 0x8c, 0x20, 0x38, 0x8a, 0x8f, 0x88, 0x43, 0x4d, 0x05, 0x5d,
	0xd8, 0x44, 0x6a, 0xc1, 0xc7, 0xce, 0x07, 0xee, 0x82, 0xd0, 0x16, 0x04, 0x37, 0x92, 0x26, 0xe3,
	0x34, 0xda, 0xce, 0x0d, 0x33, 0x13, 0xdb, 0x6e, 0x05, 0xf7, 0x7e, 0x2c, 0x97, 0x5d, 0xba, 0x94,
	0xf6, 0x53, 0xb8, 0x93, 0xd6, 0x26, 0x3e, 0x3a, 0x01, 0xbb, 0x3e, 0xbf, 0x73, 0xee, 0x3d, 0x33,
	0x5c, 0xb4, 0xc1, 0x69, 0x1b, 0x6e, 0x38, 0x55, 0x1d, 0x10, 0xf7, 0x8e, 0xa4, 0x3c, 0x08, 0x39,
	0x73, 0x54, 0xd7, 0x8e, 0x04, 0x28, 0xc0, 0x4b, 0xdf, 0x65, 0x7b, 0x22, 0x9b, 0x45, 0xbd, 0x49,
	0x78, 0x5c, 0xde, 0x52, 0xf1, 0x69, 0xb5, 0x2e, 0xd1, 0x82, 0x2b, 0xd9, 0x99, 0xa0, 0x9e, 0xa2,
	0xf5, 0x89, 0x84, 0x8f, 0xd1, 0xff, 0x04, 0x5b, 0x35, 0x0a, 0xc6, 0xce, 0x5c, 0x99, 0xd8, 0xba,
	0x11, 0x76, 0xe2, 0xa8, 0xa6, 0xbc, 0xb5, 0x8e, 0xd6, 0xa6, 0x02, 0xab, 0x54, 0x46, 0xc0, 0x25,
	0xb5, 0x0a, 0x88, 0xb8, 0x92, 0x9d, 0xd3, 0x08, 0x64, 0xa8, 0xea, 0x50, 0x8b, 0x1b, 0x9e, 0xef,
	0x43, 0xcc, 0x55, 0x4a, 0x14, 0xd1, 0xa6, 0x2b, 0xd9, 0x55, 0xa8, 0x9a, 0x81, 0xf0, 0x3a, 0x17,
	0x02, 0xda, 0x1a, 0x68, 0x0b, 0x59, 0xae, 0x64, 0x35, 0xca, 0x83, 0x11, 0xe0, 0x42, 0x10, 0xb7,
	0x68, 0x1d, 0x4e, 0x7e, 0x52, 0xe5, 0xf7, 0x3c, 0xca, 0xbb, 0x92, 0xe1, 0x3b, 0x34, 0xff, 0xab,
	0xdf, 0xb6, 0xbe, 0xcd, 0xd4, 0xde, 0xa6, 0xf3, 0x47, 0x30, 0x99, 0x89, 0x7b, 0x68, 0x51, 0xd3,
	0x0e, 0xef, 0x66, 0xe6, 0x68, 0x68, 0xb3, 0x32, 0x0b, 0x9d, 0x8e, 0x7e, 0x34, 0xd0, 0xb2, 0xfe,
	0xdd, 0x70, 0x76, 0x0d, 0xbd, 0xc1, 0x3c, 0x98, 0xd1, 0x90, 0x2e, 0xf1, 0x64, 0xa0, 0x95, 0x8c,
	0x7f, 0xc1, 0x7b, 0x99, 0xa1, 0x19, 0x0e, 0xf3, 0x70, 0x56, 0x47, 0xb2, 0xc7, 0x69, 0xed, 0x65,
	0x40, 0x8c, 0xfe, 0x80, 0x18, 0x6f, 0x03, 0x62, 0x3c, 0x0f, 0x49, 0xae, 0x3f, 0x24, 0xb9, 0xd7,
	0x21, 0xc9, 0x5d, 0x1f, 0xb1, 0x50, 0x35, 0xe3, 0x86, 0xed, 0x43, 0xdb, 0x19, 0xa5, 0x97, 0x92,
	0x03, 0x79, 0xa8, 0x94, 0xfc, 0xa6, 0x17, 0x72, 0x67, 0x7c, 0x18, 0x3e, 0xb4, 0x9c, 0xee, 0xd7,
	0xd1, 0xf4, 0x22, 0x2a, 0x1b, 0xff, 0xc6, 0xca, 0xfe, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31,
	0xe9, 0x12, 0xd3, 0x8e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateTransfer initiates a new transfer between subaccounts.
	CreateTransfer(ctx context.Context, in *MsgCreateTransfer, opts ...grpc.CallOption) (*MsgCreateTransferResponse, error)
	// DepositToSubaccount initiates a new transfer from an `x/bank` account
	// to an `x/subaccounts` subaccount.
	DepositToSubaccount(ctx context.Context, in *MsgDepositToSubaccount, opts ...grpc.CallOption) (*MsgDepositToSubaccountResponse, error)
	// WithdrawFromSubaccount initiates a new transfer from an `x/subaccounts`
	// subaccount to an `x/bank` account.
	WithdrawFromSubaccount(ctx context.Context, in *MsgWithdrawFromSubaccount, opts ...grpc.CallOption) (*MsgWithdrawFromSubaccountResponse, error)
	// SendFromModuleToAccount initiates a new transfer from a module to an
	// `x/bank` account (should only be executed by governance).
	SendFromModuleToAccount(ctx context.Context, in *MsgSendFromModuleToAccount, opts ...grpc.CallOption) (*MsgSendFromModuleToAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateTransfer(ctx context.Context, in *MsgCreateTransfer, opts ...grpc.CallOption) (*MsgCreateTransferResponse, error) {
	out := new(MsgCreateTransferResponse)
	err := c.cc.Invoke(ctx, "/nemo_network.sending.Msg/CreateTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DepositToSubaccount(ctx context.Context, in *MsgDepositToSubaccount, opts ...grpc.CallOption) (*MsgDepositToSubaccountResponse, error) {
	out := new(MsgDepositToSubaccountResponse)
	err := c.cc.Invoke(ctx, "/nemo_network.sending.Msg/DepositToSubaccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawFromSubaccount(ctx context.Context, in *MsgWithdrawFromSubaccount, opts ...grpc.CallOption) (*MsgWithdrawFromSubaccountResponse, error) {
	out := new(MsgWithdrawFromSubaccountResponse)
	err := c.cc.Invoke(ctx, "/nemo_network.sending.Msg/WithdrawFromSubaccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendFromModuleToAccount(ctx context.Context, in *MsgSendFromModuleToAccount, opts ...grpc.CallOption) (*MsgSendFromModuleToAccountResponse, error) {
	out := new(MsgSendFromModuleToAccountResponse)
	err := c.cc.Invoke(ctx, "/nemo_network.sending.Msg/SendFromModuleToAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateTransfer initiates a new transfer between subaccounts.
	CreateTransfer(context.Context, *MsgCreateTransfer) (*MsgCreateTransferResponse, error)
	// DepositToSubaccount initiates a new transfer from an `x/bank` account
	// to an `x/subaccounts` subaccount.
	DepositToSubaccount(context.Context, *MsgDepositToSubaccount) (*MsgDepositToSubaccountResponse, error)
	// WithdrawFromSubaccount initiates a new transfer from an `x/subaccounts`
	// subaccount to an `x/bank` account.
	WithdrawFromSubaccount(context.Context, *MsgWithdrawFromSubaccount) (*MsgWithdrawFromSubaccountResponse, error)
	// SendFromModuleToAccount initiates a new transfer from a module to an
	// `x/bank` account (should only be executed by governance).
	SendFromModuleToAccount(context.Context, *MsgSendFromModuleToAccount) (*MsgSendFromModuleToAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateTransfer(ctx context.Context, req *MsgCreateTransfer) (*MsgCreateTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (*UnimplementedMsgServer) DepositToSubaccount(ctx context.Context, req *MsgDepositToSubaccount) (*MsgDepositToSubaccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositToSubaccount not implemented")
}
func (*UnimplementedMsgServer) WithdrawFromSubaccount(ctx context.Context, req *MsgWithdrawFromSubaccount) (*MsgWithdrawFromSubaccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawFromSubaccount not implemented")
}
func (*UnimplementedMsgServer) SendFromModuleToAccount(ctx context.Context, req *MsgSendFromModuleToAccount) (*MsgSendFromModuleToAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFromModuleToAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo_network.sending.Msg/CreateTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateTransfer(ctx, req.(*MsgCreateTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DepositToSubaccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDepositToSubaccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositToSubaccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo_network.sending.Msg/DepositToSubaccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositToSubaccount(ctx, req.(*MsgDepositToSubaccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawFromSubaccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawFromSubaccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawFromSubaccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo_network.sending.Msg/WithdrawFromSubaccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawFromSubaccount(ctx, req.(*MsgWithdrawFromSubaccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendFromModuleToAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendFromModuleToAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendFromModuleToAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo_network.sending.Msg/SendFromModuleToAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendFromModuleToAccount(ctx, req.(*MsgSendFromModuleToAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nemo_network.sending.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransfer",
			Handler:    _Msg_CreateTransfer_Handler,
		},
		{
			MethodName: "DepositToSubaccount",
			Handler:    _Msg_DepositToSubaccount_Handler,
		},
		{
			MethodName: "WithdrawFromSubaccount",
			Handler:    _Msg_WithdrawFromSubaccount_Handler,
		},
		{
			MethodName: "SendFromModuleToAccount",
			Handler:    _Msg_SendFromModuleToAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nemo_network/sending/tx.proto",
}

func (m *MsgCreateTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Transfer != nil {
		{
			size, err := m.Transfer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateTransferResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateTransferResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateTransferResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDepositToSubaccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositToSubaccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositToSubaccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawFromSubaccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawFromSubaccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawFromSubaccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSendFromModuleToAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendFromModuleToAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendFromModuleToAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Transfer != nil {
		l = m.Transfer.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateTransferResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDepositToSubaccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawFromSubaccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSendFromModuleToAccountResponse) Size() (n int) {
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
func (m *MsgCreateTransfer) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transfer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Transfer == nil {
				m.Transfer = &Transfer{}
			}
			if err := m.Transfer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgCreateTransferResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateTransferResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateTransferResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDepositToSubaccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDepositToSubaccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositToSubaccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgWithdrawFromSubaccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawFromSubaccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawFromSubaccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgSendFromModuleToAccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendFromModuleToAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendFromModuleToAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
