// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/listing/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_dydxprotocol_v4_chain_protocol_dtypes "github.com/dydxprotocol/v4-chain/protocol/dtypes"
	types "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"
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

// MsgSetMarketsHardCap is used to set a hard cap on the number of markets
// listed
type MsgSetMarketsHardCap struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// Hard cap for the total number of markets listed
	HardCapForMarkets uint32 `protobuf:"varint,2,opt,name=hard_cap_for_markets,json=hardCapForMarkets,proto3" json:"hard_cap_for_markets,omitempty"`
}

func (m *MsgSetMarketsHardCap) Reset()         { *m = MsgSetMarketsHardCap{} }
func (m *MsgSetMarketsHardCap) String() string { return proto.CompactTextString(m) }
func (*MsgSetMarketsHardCap) ProtoMessage()    {}
func (*MsgSetMarketsHardCap) Descriptor() ([]byte, []int) {
	return fileDescriptor_144a579c1e2dcb94, []int{0}
}
func (m *MsgSetMarketsHardCap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetMarketsHardCap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetMarketsHardCap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetMarketsHardCap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetMarketsHardCap.Merge(m, src)
}
func (m *MsgSetMarketsHardCap) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetMarketsHardCap) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetMarketsHardCap.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetMarketsHardCap proto.InternalMessageInfo

func (m *MsgSetMarketsHardCap) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSetMarketsHardCap) GetHardCapForMarkets() uint32 {
	if m != nil {
		return m.HardCapForMarkets
	}
	return 0
}

// MsgSetMarketsHardCapResponse defines the MsgSetMarketsHardCap response
type MsgSetMarketsHardCapResponse struct {
}

func (m *MsgSetMarketsHardCapResponse) Reset()         { *m = MsgSetMarketsHardCapResponse{} }
func (m *MsgSetMarketsHardCapResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetMarketsHardCapResponse) ProtoMessage()    {}
func (*MsgSetMarketsHardCapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_144a579c1e2dcb94, []int{1}
}
func (m *MsgSetMarketsHardCapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetMarketsHardCapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetMarketsHardCapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetMarketsHardCapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetMarketsHardCapResponse.Merge(m, src)
}
func (m *MsgSetMarketsHardCapResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetMarketsHardCapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetMarketsHardCapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetMarketsHardCapResponse proto.InternalMessageInfo

// MsgCreateMarketPermissionless is a message used to create new markets without
// // going through x/gov
type MsgCreateMarketPermissionless struct {
	// The name of the `Perpetual` (e.g. `BTC-USD`).
	Ticker string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	// The subaccount to deposit from.
	SubaccountId *types.SubaccountId `protobuf:"bytes,2,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id,omitempty"`
	// Number of quote quantums to deposit.
	QuoteQuantums github_com_dydxprotocol_v4_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,3,opt,name=quote_quantums,json=quoteQuantums,proto3,customtype=github.com/dydxprotocol/v4-chain/protocol/dtypes.SerializableInt" json:"quote_quantums"`
}

func (m *MsgCreateMarketPermissionless) Reset()         { *m = MsgCreateMarketPermissionless{} }
func (m *MsgCreateMarketPermissionless) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMarketPermissionless) ProtoMessage()    {}
func (*MsgCreateMarketPermissionless) Descriptor() ([]byte, []int) {
	return fileDescriptor_144a579c1e2dcb94, []int{2}
}
func (m *MsgCreateMarketPermissionless) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMarketPermissionless) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMarketPermissionless.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMarketPermissionless) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMarketPermissionless.Merge(m, src)
}
func (m *MsgCreateMarketPermissionless) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMarketPermissionless) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMarketPermissionless.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMarketPermissionless proto.InternalMessageInfo

func (m *MsgCreateMarketPermissionless) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *MsgCreateMarketPermissionless) GetSubaccountId() *types.SubaccountId {
	if m != nil {
		return m.SubaccountId
	}
	return nil
}

// MsgCreateMarketPermissionlessResponse defines the
// MsgCreateMarketPermissionless response
type MsgCreateMarketPermissionlessResponse struct {
}

func (m *MsgCreateMarketPermissionlessResponse) Reset()         { *m = MsgCreateMarketPermissionlessResponse{} }
func (m *MsgCreateMarketPermissionlessResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMarketPermissionlessResponse) ProtoMessage()    {}
func (*MsgCreateMarketPermissionlessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_144a579c1e2dcb94, []int{3}
}
func (m *MsgCreateMarketPermissionlessResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMarketPermissionlessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMarketPermissionlessResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMarketPermissionlessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMarketPermissionlessResponse.Merge(m, src)
}
func (m *MsgCreateMarketPermissionlessResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMarketPermissionlessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMarketPermissionlessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMarketPermissionlessResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetMarketsHardCap)(nil), "dydxprotocol.listing.MsgSetMarketsHardCap")
	proto.RegisterType((*MsgSetMarketsHardCapResponse)(nil), "dydxprotocol.listing.MsgSetMarketsHardCapResponse")
	proto.RegisterType((*MsgCreateMarketPermissionless)(nil), "dydxprotocol.listing.MsgCreateMarketPermissionless")
	proto.RegisterType((*MsgCreateMarketPermissionlessResponse)(nil), "dydxprotocol.listing.MsgCreateMarketPermissionlessResponse")
}

func init() { proto.RegisterFile("dydxprotocol/listing/tx.proto", fileDescriptor_144a579c1e2dcb94) }

var fileDescriptor_144a579c1e2dcb94 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x51, 0x6b, 0xd3, 0x50,
	0x14, 0x6e, 0x3a, 0x18, 0xec, 0xba, 0x0e, 0x16, 0x82, 0xd6, 0xe0, 0xb2, 0x52, 0x50, 0xeb, 0x60,
	0x09, 0x76, 0x22, 0x38, 0x5f, 0xb4, 0x03, 0xd9, 0x90, 0x82, 0x26, 0x6f, 0xbe, 0x84, 0xdb, 0xe4,
	0x7a, 0x73, 0x59, 0x92, 0x9b, 0xdd, 0x73, 0x33, 0x5a, 0x1f, 0x7d, 0x13, 0x41, 0xfd, 0x29, 0x3e,
	0xf8, 0x23, 0xf6, 0x38, 0x7c, 0x12, 0x1f, 0x86, 0xb4, 0x0f, 0xfe, 0x0d, 0x69, 0x6e, 0xba, 0xb6,
	0xd8, 0x8d, 0xb2, 0xa7, 0x9e, 0x73, 0xbe, 0x73, 0xbe, 0xf3, 0xdd, 0xf3, 0x35, 0x68, 0x2b, 0x1c,
	0x84, 0xfd, 0x4c, 0x70, 0xc9, 0x03, 0x1e, 0x3b, 0x31, 0x03, 0xc9, 0x52, 0xea, 0xc8, 0xbe, 0x5d,
	0xd4, 0x74, 0x63, 0x16, 0xb6, 0x4b, 0xd8, 0xbc, 0x1b, 0x70, 0x48, 0x38, 0xf8, 0x05, 0xe0, 0xa8,
	0x44, 0x0d, 0x98, 0x77, 0x54, 0xe6, 0x24, 0x40, 0x9d, 0xd3, 0xc7, 0xe3, 0x9f, 0x12, 0x30, 0x28,
	0xa7, 0x5c, 0x0d, 0x8c, 0xa3, 0xb2, 0xfa, 0x68, 0x6e, 0x3d, 0xe4, 0x3d, 0x1c, 0x04, 0x3c, 0x4f,
	0x25, 0xcc, 0xc4, 0xaa, 0xb5, 0xf9, 0x55, 0x43, 0x46, 0x17, 0xa8, 0x47, 0x64, 0x17, 0x8b, 0x63,
	0x22, 0xe1, 0x10, 0x8b, 0xf0, 0x00, 0x67, 0xfa, 0x53, 0xb4, 0x86, 0x73, 0x19, 0x71, 0xc1, 0xe4,
	0xa0, 0xae, 0x35, 0xb4, 0xd6, 0x5a, 0xa7, 0xfe, 0xf3, 0xc7, 0xae, 0x51, 0xea, 0x7a, 0x19, 0x86,
	0x82, 0x00, 0x78, 0x52, 0xb0, 0x94, 0xba, 0xd3, 0x56, 0xdd, 0x41, 0x46, 0x84, 0x45, 0xe8, 0x07,
	0x38, 0xf3, 0xdf, 0x73, 0xe1, 0x27, 0x8a, 0xb6, 0x5e, 0x6d, 0x68, 0xad, 0x9a, 0xbb, 0x19, 0x29,
	0xfa, 0x57, 0x5c, 0x94, 0xfb, 0xf6, 0x37, 0x3e, 0xfe, 0xfd, 0xbe, 0x33, 0x25, 0x68, 0x5a, 0xe8,
	0xde, 0x22, 0x41, 0x2e, 0x81, 0x8c, 0xa7, 0x40, 0x9a, 0x9f, 0xaa, 0x68, 0xab, 0x0b, 0xf4, 0x40,
	0x10, 0x2c, 0x89, 0xea, 0x79, 0x43, 0x44, 0xc2, 0x00, 0x18, 0x4f, 0x63, 0x02, 0xa0, 0xdf, 0x46,
	0xab, 0x92, 0x05, 0xc7, 0x44, 0x28, 0xdd, 0x6e, 0x99, 0xe9, 0xaf, 0x51, 0x6d, 0xfa, 0x7e, 0x9f,
	0x85, 0x85, 0xa6, 0x5b, 0xed, 0x07, 0xf6, 0x9c, 0x1d, 0x33, 0xe7, 0xb2, 0xbd, 0xcb, 0xf8, 0x28,
	0x74, 0xd7, 0x61, 0x26, 0xd3, 0x39, 0xda, 0x38, 0xc9, 0xb9, 0x24, 0xfe, 0x49, 0x8e, 0x53, 0x99,
	0x27, 0x50, 0x5f, 0x69, 0x68, 0xad, 0xf5, 0xce, 0xe1, 0xd9, 0xc5, 0x76, 0xe5, 0xf7, 0xc5, 0xf6,
	0x0b, 0xca, 0x64, 0x94, 0xf7, 0xec, 0x80, 0x27, 0xce, 0x9c, 0x1d, 0xa7, 0x4f, 0x76, 0x83, 0x08,
	0xb3, 0xd4, 0xb9, 0xac, 0x84, 0x72, 0x90, 0x11, 0xb0, 0x3d, 0x22, 0x18, 0x8e, 0xd9, 0x07, 0xdc,
	0x8b, 0xc9, 0x51, 0x2a, 0xdd, 0x5a, 0xc1, 0xff, 0xb6, 0xa4, 0xdf, 0xd7, 0xc7, 0x77, 0x9a, 0x7f,
	0x40, 0xf3, 0x21, 0xba, 0x7f, 0xed, 0x29, 0x26, 0x47, 0x6b, 0x7f, 0xae, 0xa2, 0x95, 0x2e, 0x50,
	0x1d, 0xd0, 0xe6, 0xff, 0x56, 0xef, 0xd8, 0x8b, 0xfe, 0x8f, 0xf6, 0x22, 0x17, 0xcc, 0xf6, 0xf2,
	0xbd, 0x93, 0xe5, 0xfa, 0x17, 0x0d, 0x99, 0xd7, 0xd8, 0xb5, 0x77, 0x25, 0xe5, 0xd5, 0x43, 0xe6,
	0xf3, 0x1b, 0x0c, 0x4d, 0x04, 0x75, 0xbc, 0xb3, 0xa1, 0xa5, 0x9d, 0x0f, 0x2d, 0xed, 0xcf, 0xd0,
	0xd2, 0xbe, 0x8d, 0xac, 0xca, 0xf9, 0xc8, 0xaa, 0xfc, 0x1a, 0x59, 0x95, 0x77, 0xcf, 0x96, 0x77,
	0xad, 0x3f, 0xfd, 0xae, 0xc7, 0xf6, 0xf5, 0x56, 0x0b, 0x64, 0xef, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0xf1, 0x01, 0x5e, 0xfc, 0x03, 0x00, 0x00,
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
	// SetMarketsHardCap sets a hard cap on the number of markets listed
	SetMarketsHardCap(ctx context.Context, in *MsgSetMarketsHardCap, opts ...grpc.CallOption) (*MsgSetMarketsHardCapResponse, error)
	// CreateMarketPermissionless creates a new market without going through x/gov
	CreateMarketPermissionless(ctx context.Context, in *MsgCreateMarketPermissionless, opts ...grpc.CallOption) (*MsgCreateMarketPermissionlessResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetMarketsHardCap(ctx context.Context, in *MsgSetMarketsHardCap, opts ...grpc.CallOption) (*MsgSetMarketsHardCapResponse, error) {
	out := new(MsgSetMarketsHardCapResponse)
	err := c.cc.Invoke(ctx, "/dydxprotocol.listing.Msg/SetMarketsHardCap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateMarketPermissionless(ctx context.Context, in *MsgCreateMarketPermissionless, opts ...grpc.CallOption) (*MsgCreateMarketPermissionlessResponse, error) {
	out := new(MsgCreateMarketPermissionlessResponse)
	err := c.cc.Invoke(ctx, "/dydxprotocol.listing.Msg/CreateMarketPermissionless", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SetMarketsHardCap sets a hard cap on the number of markets listed
	SetMarketsHardCap(context.Context, *MsgSetMarketsHardCap) (*MsgSetMarketsHardCapResponse, error)
	// CreateMarketPermissionless creates a new market without going through x/gov
	CreateMarketPermissionless(context.Context, *MsgCreateMarketPermissionless) (*MsgCreateMarketPermissionlessResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetMarketsHardCap(ctx context.Context, req *MsgSetMarketsHardCap) (*MsgSetMarketsHardCapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMarketsHardCap not implemented")
}
func (*UnimplementedMsgServer) CreateMarketPermissionless(ctx context.Context, req *MsgCreateMarketPermissionless) (*MsgCreateMarketPermissionlessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMarketPermissionless not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetMarketsHardCap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetMarketsHardCap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetMarketsHardCap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dydxprotocol.listing.Msg/SetMarketsHardCap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetMarketsHardCap(ctx, req.(*MsgSetMarketsHardCap))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateMarketPermissionless_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMarketPermissionless)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMarketPermissionless(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dydxprotocol.listing.Msg/CreateMarketPermissionless",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMarketPermissionless(ctx, req.(*MsgCreateMarketPermissionless))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dydxprotocol.listing.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMarketsHardCap",
			Handler:    _Msg_SetMarketsHardCap_Handler,
		},
		{
			MethodName: "CreateMarketPermissionless",
			Handler:    _Msg_CreateMarketPermissionless_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dydxprotocol/listing/tx.proto",
}

func (m *MsgSetMarketsHardCap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetMarketsHardCap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetMarketsHardCap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HardCapForMarkets != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.HardCapForMarkets))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetMarketsHardCapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetMarketsHardCapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetMarketsHardCapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreateMarketPermissionless) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMarketPermissionless) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMarketPermissionless) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.QuoteQuantums.Size()
		i -= size
		if _, err := m.QuoteQuantums.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.SubaccountId != nil {
		{
			size, err := m.SubaccountId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Ticker) > 0 {
		i -= len(m.Ticker)
		copy(dAtA[i:], m.Ticker)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ticker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateMarketPermissionlessResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMarketPermissionlessResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMarketPermissionlessResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSetMarketsHardCap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.HardCapForMarkets != 0 {
		n += 1 + sovTx(uint64(m.HardCapForMarkets))
	}
	return n
}

func (m *MsgSetMarketsHardCapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreateMarketPermissionless) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.SubaccountId != nil {
		l = m.SubaccountId.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.QuoteQuantums.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCreateMarketPermissionlessResponse) Size() (n int) {
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
func (m *MsgSetMarketsHardCap) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetMarketsHardCap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetMarketsHardCap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HardCapForMarkets", wireType)
			}
			m.HardCapForMarkets = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HardCapForMarkets |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgSetMarketsHardCapResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetMarketsHardCapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetMarketsHardCapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgCreateMarketPermissionless) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateMarketPermissionless: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMarketPermissionless: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
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
			m.Ticker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountId", wireType)
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
			if m.SubaccountId == nil {
				m.SubaccountId = &types.SubaccountId{}
			}
			if err := m.SubaccountId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteQuantums", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteQuantums.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreateMarketPermissionlessResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateMarketPermissionlessResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMarketPermissionlessResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
