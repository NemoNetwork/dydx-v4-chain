// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo_network/clob/liquidations.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/nemo-network/v4-chain/protocol/x/subaccounts/types"
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

// PerpetualLiquidationInfo holds information about a liquidation that occurred
// for a position held by a subaccount.
// Note this proto is defined to make it easier to hash
// the metadata of a liquidation, and is never written to state.
type PerpetualLiquidationInfo struct {
	// The id of the subaccount that got liquidated/deleveraged or was deleveraged
	// onto.
	SubaccountId types.SubaccountId `protobuf:"bytes,1,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id"`
	// The id of the perpetual involved.
	PerpetualId uint32 `protobuf:"varint,2,opt,name=perpetual_id,json=perpetualId,proto3" json:"perpetual_id,omitempty"`
}

func (m *PerpetualLiquidationInfo) Reset()         { *m = PerpetualLiquidationInfo{} }
func (m *PerpetualLiquidationInfo) String() string { return proto.CompactTextString(m) }
func (*PerpetualLiquidationInfo) ProtoMessage()    {}
func (*PerpetualLiquidationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ae8eafa85a9846, []int{0}
}
func (m *PerpetualLiquidationInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerpetualLiquidationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerpetualLiquidationInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerpetualLiquidationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerpetualLiquidationInfo.Merge(m, src)
}
func (m *PerpetualLiquidationInfo) XXX_Size() int {
	return m.Size()
}
func (m *PerpetualLiquidationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PerpetualLiquidationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PerpetualLiquidationInfo proto.InternalMessageInfo

func (m *PerpetualLiquidationInfo) GetSubaccountId() types.SubaccountId {
	if m != nil {
		return m.SubaccountId
	}
	return types.SubaccountId{}
}

func (m *PerpetualLiquidationInfo) GetPerpetualId() uint32 {
	if m != nil {
		return m.PerpetualId
	}
	return 0
}

// SubaccountLiquidationInfo holds liquidation information per-subaccount in the
// current block.
type SubaccountLiquidationInfo struct {
	// An unsorted list of unique perpetual IDs that the subaccount has previously
	// liquidated.
	PerpetualsLiquidated []uint32 `protobuf:"varint,1,rep,packed,name=perpetuals_liquidated,json=perpetualsLiquidated,proto3" json:"perpetuals_liquidated,omitempty"`
	// The notional value (in quote quantums, determined by the oracle price) of
	// all positions liquidated for this subaccount.
	NotionalLiquidated uint64 `protobuf:"varint,2,opt,name=notional_liquidated,json=notionalLiquidated,proto3" json:"notional_liquidated,omitempty"`
	// The amount of funds that the insurance fund has lost
	// covering this subaccount.
	QuantumsInsuranceLost uint64 `protobuf:"varint,3,opt,name=quantums_insurance_lost,json=quantumsInsuranceLost,proto3" json:"quantums_insurance_lost,omitempty"`
}

func (m *SubaccountLiquidationInfo) Reset()         { *m = SubaccountLiquidationInfo{} }
func (m *SubaccountLiquidationInfo) String() string { return proto.CompactTextString(m) }
func (*SubaccountLiquidationInfo) ProtoMessage()    {}
func (*SubaccountLiquidationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ae8eafa85a9846, []int{1}
}
func (m *SubaccountLiquidationInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubaccountLiquidationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubaccountLiquidationInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubaccountLiquidationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubaccountLiquidationInfo.Merge(m, src)
}
func (m *SubaccountLiquidationInfo) XXX_Size() int {
	return m.Size()
}
func (m *SubaccountLiquidationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SubaccountLiquidationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SubaccountLiquidationInfo proto.InternalMessageInfo

func (m *SubaccountLiquidationInfo) GetPerpetualsLiquidated() []uint32 {
	if m != nil {
		return m.PerpetualsLiquidated
	}
	return nil
}

func (m *SubaccountLiquidationInfo) GetNotionalLiquidated() uint64 {
	if m != nil {
		return m.NotionalLiquidated
	}
	return 0
}

func (m *SubaccountLiquidationInfo) GetQuantumsInsuranceLost() uint64 {
	if m != nil {
		return m.QuantumsInsuranceLost
	}
	return 0
}

// SubaccountOpenPositionInfo holds information about open positions for a
// perpetual.
type SubaccountOpenPositionInfo struct {
	// The id of the perpetual.
	PerpetualId uint32 `protobuf:"varint,1,opt,name=perpetual_id,json=perpetualId,proto3" json:"perpetual_id,omitempty"`
	// The ids of the subaccounts with long positions.
	SubaccountsWithLongPosition []types.SubaccountId `protobuf:"bytes,2,rep,name=subaccounts_with_long_position,json=subaccountsWithLongPosition,proto3" json:"subaccounts_with_long_position"`
	// The ids of the subaccounts with short positions.
	SubaccountsWithShortPosition []types.SubaccountId `protobuf:"bytes,3,rep,name=subaccounts_with_short_position,json=subaccountsWithShortPosition,proto3" json:"subaccounts_with_short_position"`
}

func (m *SubaccountOpenPositionInfo) Reset()         { *m = SubaccountOpenPositionInfo{} }
func (m *SubaccountOpenPositionInfo) String() string { return proto.CompactTextString(m) }
func (*SubaccountOpenPositionInfo) ProtoMessage()    {}
func (*SubaccountOpenPositionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ae8eafa85a9846, []int{2}
}
func (m *SubaccountOpenPositionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubaccountOpenPositionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubaccountOpenPositionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubaccountOpenPositionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubaccountOpenPositionInfo.Merge(m, src)
}
func (m *SubaccountOpenPositionInfo) XXX_Size() int {
	return m.Size()
}
func (m *SubaccountOpenPositionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SubaccountOpenPositionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SubaccountOpenPositionInfo proto.InternalMessageInfo

func (m *SubaccountOpenPositionInfo) GetPerpetualId() uint32 {
	if m != nil {
		return m.PerpetualId
	}
	return 0
}

func (m *SubaccountOpenPositionInfo) GetSubaccountsWithLongPosition() []types.SubaccountId {
	if m != nil {
		return m.SubaccountsWithLongPosition
	}
	return nil
}

func (m *SubaccountOpenPositionInfo) GetSubaccountsWithShortPosition() []types.SubaccountId {
	if m != nil {
		return m.SubaccountsWithShortPosition
	}
	return nil
}

func init() {
	proto.RegisterType((*PerpetualLiquidationInfo)(nil), "nemo_network.clob.PerpetualLiquidationInfo")
	proto.RegisterType((*SubaccountLiquidationInfo)(nil), "nemo_network.clob.SubaccountLiquidationInfo")
	proto.RegisterType((*SubaccountOpenPositionInfo)(nil), "nemo_network.clob.SubaccountOpenPositionInfo")
}

func init() {
	proto.RegisterFile("nemo_network/clob/liquidations.proto", fileDescriptor_a3ae8eafa85a9846)
}

var fileDescriptor_a3ae8eafa85a9846 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0x33, 0x49, 0xc5, 0x62, 0xda, 0x2c, 0x30, 0xad, 0x30, 0x01, 0xb9, 0x21, 0x42, 0x28,
	0x2c, 0x6a, 0x4b, 0x14, 0xf5, 0x00, 0xdd, 0x59, 0xb2, 0x44, 0x70, 0x17, 0x48, 0x6c, 0x46, 0xfe,
	0x33, 0xd8, 0x23, 0x9c, 0x79, 0x8e, 0x67, 0x86, 0xc2, 0x2d, 0x7a, 0x0b, 0xce, 0xc0, 0x0d, 0xba,
	0xec, 0x92, 0x15, 0x42, 0xc9, 0x45, 0xd0, 0x38, 0x8c, 0x3d, 0x4d, 0x56, 0x74, 0xf7, 0xf4, 0xbe,
	0xf7, 0xbe, 0xdf, 0x37, 0xe3, 0x31, 0x7e, 0xc5, 0xe9, 0x12, 0x08, 0xa7, 0xf2, 0x1a, 0x9a, 0x2f,
	0x41, 0x56, 0x41, 0x1a, 0x54, 0x6c, 0xa5, 0x58, 0x9e, 0x48, 0x06, 0x5c, 0xf8, 0x75, 0x03, 0x12,
	0x9c, 0xc7, 0xf6, 0x94, 0xaf, 0xa7, 0x26, 0xc7, 0x05, 0x14, 0xd0, 0xaa, 0x81, 0xae, 0xb6, 0x83,
	0x93, 0x37, 0xf7, 0xec, 0x84, 0x4a, 0x93, 0x2c, 0x03, 0xc5, 0xa5, 0xb0, 0xea, 0xed, 0xe8, 0xec,
	0x06, 0x61, 0x77, 0x41, 0x9b, 0x9a, 0x4a, 0x95, 0x54, 0x51, 0xcf, 0x0c, 0xf9, 0x67, 0x70, 0x3e,
	0xe0, 0x71, 0xbf, 0x40, 0x58, 0xee, 0xa2, 0x29, 0x9a, 0x1f, 0xbe, 0x7d, 0xed, 0xdf, 0x0b, 0x62,
	0xf9, 0xfb, 0x57, 0x5d, 0x1d, 0xe6, 0x97, 0x07, 0xb7, 0xbf, 0x4f, 0x07, 0xf1, 0x91, 0xb0, 0x7a,
	0xce, 0x4b, 0x7c, 0x54, 0x1b, 0x9c, 0x76, 0x1c, 0x4e, 0xd1, 0x7c, 0x1c, 0x1f, 0x76, 0xbd, 0x30,
	0x9f, 0xfd, 0x44, 0xf8, 0x59, 0xef, 0xb3, 0x9b, 0xe9, 0x1c, 0x9f, 0x74, 0xc3, 0x82, 0x98, 0x5b,
	0xa2, 0x3a, 0xdb, 0x68, 0x3e, 0x8e, 0x8f, 0x7b, 0x31, 0xea, 0x34, 0x27, 0xc0, 0x4f, 0x38, 0x68,
	0x8b, 0xa4, 0xb2, 0x57, 0x34, 0xfc, 0x20, 0x76, 0x8c, 0x64, 0x2d, 0x5c, 0xe0, 0xa7, 0x2b, 0x95,
	0x70, 0xa9, 0x96, 0x82, 0x30, 0x2e, 0x54, 0x93, 0xf0, 0x8c, 0x92, 0x0a, 0x84, 0x74, 0x47, 0xed,
	0xd2, 0x89, 0x91, 0x43, 0xa3, 0x46, 0x20, 0xe4, 0xec, 0xc7, 0x10, 0x4f, 0xfa, 0xec, 0xef, 0x6b,
	0xca, 0x17, 0x20, 0x58, 0x17, 0x7e, 0xf7, 0xf4, 0x68, 0xef, 0xf4, 0xce, 0x0a, 0x7b, 0xd6, 0x85,
	0x92, 0x6b, 0x26, 0x4b, 0x52, 0x01, 0x2f, 0x48, 0xfd, 0xcf, 0xc8, 0x1d, 0x4e, 0x47, 0xff, 0xfd,
	0x11, 0x9e, 0x5b, 0xfa, 0x47, 0x26, 0xcb, 0x08, 0x78, 0x61, 0x92, 0x39, 0x02, 0x9f, 0xee, 0x21,
	0x45, 0x09, 0x8d, 0xec, 0x99, 0xa3, 0x07, 0x30, 0x5f, 0xec, 0x30, 0xaf, 0xb4, 0xa5, 0x81, 0x5e,
	0x2e, 0x6e, 0xd7, 0x1e, 0xba, 0x5b, 0x7b, 0xe8, 0xcf, 0xda, 0x43, 0x37, 0x1b, 0x6f, 0x70, 0xb7,
	0xf1, 0x06, 0xbf, 0x36, 0xde, 0xe0, 0xd3, 0x45, 0xc1, 0x64, 0xa9, 0x52, 0x3f, 0x83, 0x65, 0xa0,
	0x79, 0x67, 0xe6, 0x21, 0x7f, 0x7d, 0x77, 0x96, 0x95, 0x09, 0xe3, 0x41, 0xfb, 0x74, 0x33, 0xa8,
	0x82, 0x6f, 0xdb, 0x7f, 0x45, 0x7e, 0xaf, 0xa9, 0x48, 0x1f, 0xb5, 0xed, 0xf3, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x9b, 0xfe, 0xdd, 0x79, 0x4d, 0x03, 0x00, 0x00,
}

func (m *PerpetualLiquidationInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerpetualLiquidationInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerpetualLiquidationInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PerpetualId != 0 {
		i = encodeVarintLiquidations(dAtA, i, uint64(m.PerpetualId))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.SubaccountId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidations(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *SubaccountLiquidationInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubaccountLiquidationInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubaccountLiquidationInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QuantumsInsuranceLost != 0 {
		i = encodeVarintLiquidations(dAtA, i, uint64(m.QuantumsInsuranceLost))
		i--
		dAtA[i] = 0x18
	}
	if m.NotionalLiquidated != 0 {
		i = encodeVarintLiquidations(dAtA, i, uint64(m.NotionalLiquidated))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PerpetualsLiquidated) > 0 {
		dAtA3 := make([]byte, len(m.PerpetualsLiquidated)*10)
		var j2 int
		for _, num := range m.PerpetualsLiquidated {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintLiquidations(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubaccountOpenPositionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubaccountOpenPositionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubaccountOpenPositionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubaccountsWithShortPosition) > 0 {
		for iNdEx := len(m.SubaccountsWithShortPosition) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubaccountsWithShortPosition[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidations(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SubaccountsWithLongPosition) > 0 {
		for iNdEx := len(m.SubaccountsWithLongPosition) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubaccountsWithLongPosition[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidations(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PerpetualId != 0 {
		i = encodeVarintLiquidations(dAtA, i, uint64(m.PerpetualId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidations(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PerpetualLiquidationInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SubaccountId.Size()
	n += 1 + l + sovLiquidations(uint64(l))
	if m.PerpetualId != 0 {
		n += 1 + sovLiquidations(uint64(m.PerpetualId))
	}
	return n
}

func (m *SubaccountLiquidationInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PerpetualsLiquidated) > 0 {
		l = 0
		for _, e := range m.PerpetualsLiquidated {
			l += sovLiquidations(uint64(e))
		}
		n += 1 + sovLiquidations(uint64(l)) + l
	}
	if m.NotionalLiquidated != 0 {
		n += 1 + sovLiquidations(uint64(m.NotionalLiquidated))
	}
	if m.QuantumsInsuranceLost != 0 {
		n += 1 + sovLiquidations(uint64(m.QuantumsInsuranceLost))
	}
	return n
}

func (m *SubaccountOpenPositionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PerpetualId != 0 {
		n += 1 + sovLiquidations(uint64(m.PerpetualId))
	}
	if len(m.SubaccountsWithLongPosition) > 0 {
		for _, e := range m.SubaccountsWithLongPosition {
			l = e.Size()
			n += 1 + l + sovLiquidations(uint64(l))
		}
	}
	if len(m.SubaccountsWithShortPosition) > 0 {
		for _, e := range m.SubaccountsWithShortPosition {
			l = e.Size()
			n += 1 + l + sovLiquidations(uint64(l))
		}
	}
	return n
}

func sovLiquidations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidations(x uint64) (n int) {
	return sovLiquidations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PerpetualLiquidationInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidations
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
			return fmt.Errorf("proto: PerpetualLiquidationInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerpetualLiquidationInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidations
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
				return ErrInvalidLengthLiquidations
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubaccountId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualId", wireType)
			}
			m.PerpetualId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidations
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
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidations
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
func (m *SubaccountLiquidationInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidations
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
			return fmt.Errorf("proto: SubaccountLiquidationInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubaccountLiquidationInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLiquidations
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PerpetualsLiquidated = append(m.PerpetualsLiquidated, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLiquidations
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLiquidations
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLiquidations
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PerpetualsLiquidated) == 0 {
					m.PerpetualsLiquidated = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLiquidations
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PerpetualsLiquidated = append(m.PerpetualsLiquidated, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualsLiquidated", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotionalLiquidated", wireType)
			}
			m.NotionalLiquidated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotionalLiquidated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuantumsInsuranceLost", wireType)
			}
			m.QuantumsInsuranceLost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuantumsInsuranceLost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidations
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
func (m *SubaccountOpenPositionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidations
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
			return fmt.Errorf("proto: SubaccountOpenPositionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubaccountOpenPositionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualId", wireType)
			}
			m.PerpetualId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidations
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountsWithLongPosition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidations
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
				return ErrInvalidLengthLiquidations
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubaccountsWithLongPosition = append(m.SubaccountsWithLongPosition, types.SubaccountId{})
			if err := m.SubaccountsWithLongPosition[len(m.SubaccountsWithLongPosition)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountsWithShortPosition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidations
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
				return ErrInvalidLengthLiquidations
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubaccountsWithShortPosition = append(m.SubaccountsWithShortPosition, types.SubaccountId{})
			if err := m.SubaccountsWithShortPosition[len(m.SubaccountsWithShortPosition)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidations
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
func skipLiquidations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidations
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
					return 0, ErrIntOverflowLiquidations
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
					return 0, ErrIntOverflowLiquidations
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
				return 0, ErrInvalidLengthLiquidations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidations = fmt.Errorf("proto: unexpected end of group")
)
