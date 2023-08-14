// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/perpetuals/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// FundingPremium represents a funding premium value for a perpetual
// market. Can be used to represent a premium vote or a premium sample.
type FundingPremium struct {
	// The id of the perpetual market.
	PerpetualId uint32 `protobuf:"varint,1,opt,name=perpetual_id,json=perpetualId,proto3" json:"perpetual_id,omitempty"`
	// The sampled premium rate. In parts-per-million.
	PremiumPpm int32 `protobuf:"varint,2,opt,name=premium_ppm,json=premiumPpm,proto3" json:"premium_ppm,omitempty"`
}

func (m *FundingPremium) Reset()         { *m = FundingPremium{} }
func (m *FundingPremium) String() string { return proto.CompactTextString(m) }
func (*FundingPremium) ProtoMessage()    {}
func (*FundingPremium) Descriptor() ([]byte, []int) {
	return fileDescriptor_daed24c15760c356, []int{0}
}
func (m *FundingPremium) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundingPremium) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundingPremium.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundingPremium) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingPremium.Merge(m, src)
}
func (m *FundingPremium) XXX_Size() int {
	return m.Size()
}
func (m *FundingPremium) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingPremium.DiscardUnknown(m)
}

var xxx_messageInfo_FundingPremium proto.InternalMessageInfo

func (m *FundingPremium) GetPerpetualId() uint32 {
	if m != nil {
		return m.PerpetualId
	}
	return 0
}

func (m *FundingPremium) GetPremiumPpm() int32 {
	if m != nil {
		return m.PremiumPpm
	}
	return 0
}

// MsgAddPremiumVotes is a request type for the AddPremiumVotes method.
type MsgAddPremiumVotes struct {
	Votes []FundingPremium `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes"`
}

func (m *MsgAddPremiumVotes) Reset()         { *m = MsgAddPremiumVotes{} }
func (m *MsgAddPremiumVotes) String() string { return proto.CompactTextString(m) }
func (*MsgAddPremiumVotes) ProtoMessage()    {}
func (*MsgAddPremiumVotes) Descriptor() ([]byte, []int) {
	return fileDescriptor_daed24c15760c356, []int{1}
}
func (m *MsgAddPremiumVotes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddPremiumVotes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddPremiumVotes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddPremiumVotes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddPremiumVotes.Merge(m, src)
}
func (m *MsgAddPremiumVotes) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddPremiumVotes) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddPremiumVotes.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddPremiumVotes proto.InternalMessageInfo

func (m *MsgAddPremiumVotes) GetVotes() []FundingPremium {
	if m != nil {
		return m.Votes
	}
	return nil
}

// MsgAddPremiumVotesResponse defines the AddPremiumVotes
// response type.
type MsgAddPremiumVotesResponse struct {
}

func (m *MsgAddPremiumVotesResponse) Reset()         { *m = MsgAddPremiumVotesResponse{} }
func (m *MsgAddPremiumVotesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddPremiumVotesResponse) ProtoMessage()    {}
func (*MsgAddPremiumVotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_daed24c15760c356, []int{2}
}
func (m *MsgAddPremiumVotesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddPremiumVotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddPremiumVotesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddPremiumVotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddPremiumVotesResponse.Merge(m, src)
}
func (m *MsgAddPremiumVotesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddPremiumVotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddPremiumVotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddPremiumVotesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FundingPremium)(nil), "dydxprotocol.perpetuals.FundingPremium")
	proto.RegisterType((*MsgAddPremiumVotes)(nil), "dydxprotocol.perpetuals.MsgAddPremiumVotes")
	proto.RegisterType((*MsgAddPremiumVotesResponse)(nil), "dydxprotocol.perpetuals.MsgAddPremiumVotesResponse")
}

func init() { proto.RegisterFile("dydxprotocol/perpetuals/tx.proto", fileDescriptor_daed24c15760c356) }

var fileDescriptor_daed24c15760c356 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x48, 0x2d, 0x2a, 0x48, 0x2d, 0x29,
	0x4d, 0xcc, 0x29, 0xd6, 0x2f, 0xa9, 0xd0, 0x03, 0x0b, 0x0b, 0x89, 0x23, 0xab, 0xd0, 0x43, 0xa8,
	0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xe8, 0x83, 0x58, 0x10, 0xe5, 0x4a, 0x21, 0x5c,
	0x7c, 0x6e, 0xa5, 0x79, 0x29, 0x99, 0x79, 0xe9, 0x01, 0x45, 0xa9, 0xb9, 0x99, 0xa5, 0xb9, 0x42,
	0x8a, 0x5c, 0x3c, 0x70, 0x5d, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0xdc,
	0x70, 0x31, 0xcf, 0x14, 0x21, 0x79, 0x2e, 0xee, 0x02, 0x88, 0xea, 0xf8, 0x82, 0x82, 0x5c, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x2e, 0xa8, 0x50, 0x40, 0x41, 0xae, 0x52, 0x24, 0x97, 0x90,
	0x6f, 0x71, 0xba, 0x63, 0x4a, 0x0a, 0xd4, 0xd0, 0xb0, 0xfc, 0x92, 0xd4, 0x62, 0x21, 0x67, 0x2e,
	0xd6, 0x32, 0x10, 0x43, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5d, 0x0f, 0x87, 0x53, 0xf5,
	0x50, 0x5d, 0xe4, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x44, 0xaf, 0x92, 0x0c, 0x97, 0x14,
	0xa6, 0xd1, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x46, 0x55, 0x5c, 0xcc, 0xbe, 0xc5,
	0xe9, 0x42, 0xc5, 0x5c, 0xfc, 0xe8, 0x96, 0x6b, 0xe3, 0xb4, 0x0d, 0xd3, 0x38, 0x29, 0x63, 0x12,
	0x14, 0xc3, 0xec, 0x76, 0x0a, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0x9b, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x94, 0x08, 0x2c, 0x33,
	0xd1, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87, 0x8b, 0x54, 0xa0, 0x44, 0x6a, 0x65, 0x41, 0x6a,
	0x71, 0x12, 0x1b, 0x58, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xe3, 0x25, 0x33, 0xfc,
	0x01, 0x00, 0x00,
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
	// AddPremiumVotes add new samples of the funding premiums to the
	// application.
	AddPremiumVotes(ctx context.Context, in *MsgAddPremiumVotes, opts ...grpc.CallOption) (*MsgAddPremiumVotesResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddPremiumVotes(ctx context.Context, in *MsgAddPremiumVotes, opts ...grpc.CallOption) (*MsgAddPremiumVotesResponse, error) {
	out := new(MsgAddPremiumVotesResponse)
	err := c.cc.Invoke(ctx, "/dydxprotocol.perpetuals.Msg/AddPremiumVotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// AddPremiumVotes add new samples of the funding premiums to the
	// application.
	AddPremiumVotes(context.Context, *MsgAddPremiumVotes) (*MsgAddPremiumVotesResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddPremiumVotes(ctx context.Context, req *MsgAddPremiumVotes) (*MsgAddPremiumVotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPremiumVotes not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddPremiumVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddPremiumVotes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddPremiumVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dydxprotocol.perpetuals.Msg/AddPremiumVotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddPremiumVotes(ctx, req.(*MsgAddPremiumVotes))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dydxprotocol.perpetuals.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPremiumVotes",
			Handler:    _Msg_AddPremiumVotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dydxprotocol/perpetuals/tx.proto",
}

func (m *FundingPremium) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundingPremium) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundingPremium) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PremiumPpm != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PremiumPpm))
		i--
		dAtA[i] = 0x10
	}
	if m.PerpetualId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PerpetualId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddPremiumVotes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddPremiumVotes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddPremiumVotes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddPremiumVotesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddPremiumVotesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddPremiumVotesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *FundingPremium) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PerpetualId != 0 {
		n += 1 + sovTx(uint64(m.PerpetualId))
	}
	if m.PremiumPpm != 0 {
		n += 1 + sovTx(uint64(m.PremiumPpm))
	}
	return n
}

func (m *MsgAddPremiumVotes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgAddPremiumVotesResponse) Size() (n int) {
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
func (m *FundingPremium) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: FundingPremium: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundingPremium: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualId", wireType)
			}
			m.PerpetualId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PerpetualId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PremiumPpm", wireType)
			}
			m.PremiumPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PremiumPpm |= int32(b&0x7F) << shift
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
func (m *MsgAddPremiumVotes) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddPremiumVotes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddPremiumVotes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
			m.Votes = append(m.Votes, FundingPremium{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgAddPremiumVotesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddPremiumVotesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddPremiumVotesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
