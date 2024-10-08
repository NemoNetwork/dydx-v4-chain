// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo_network/subaccounts/perpetual_position.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_nemo_network_v4_chain_protocol_dtypes "github.com/nemo-network/v4-chain/protocol/dtypes"
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

// PerpetualPositions are an account’s positions of a `Perpetual`.
// Therefore they hold any information needed to trade perpetuals.
type PerpetualPosition struct {
	// The `Id` of the `Perpetual`.
	PerpetualId uint32 `protobuf:"varint,1,opt,name=perpetual_id,json=perpetualId,proto3" json:"perpetual_id,omitempty"`
	// The size of the position in base quantums.
	Quantums github_com_nemo_network_v4_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,2,opt,name=quantums,proto3,customtype=github.com/nemo-network/v4-chain/protocol/dtypes.SerializableInt" json:"quantums"`
	// The funding_index of the `Perpetual` the last time this position was
	// settled.
	FundingIndex github_com_nemo_network_v4_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,3,opt,name=funding_index,json=fundingIndex,proto3,customtype=github.com/nemo-network/v4-chain/protocol/dtypes.SerializableInt" json:"funding_index"`
	// The quote_balance of the `Perpetual`.
	QuoteBalance github_com_nemo_network_v4_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,4,opt,name=quote_balance,json=quoteBalance,proto3,customtype=github.com/nemo-network/v4-chain/protocol/dtypes.SerializableInt" json:"quote_balance"`
}

func (m *PerpetualPosition) Reset()         { *m = PerpetualPosition{} }
func (m *PerpetualPosition) String() string { return proto.CompactTextString(m) }
func (*PerpetualPosition) ProtoMessage()    {}
func (*PerpetualPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_211b9880e1a5e0d8, []int{0}
}
func (m *PerpetualPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerpetualPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerpetualPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerpetualPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerpetualPosition.Merge(m, src)
}
func (m *PerpetualPosition) XXX_Size() int {
	return m.Size()
}
func (m *PerpetualPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_PerpetualPosition.DiscardUnknown(m)
}

var xxx_messageInfo_PerpetualPosition proto.InternalMessageInfo

func (m *PerpetualPosition) GetPerpetualId() uint32 {
	if m != nil {
		return m.PerpetualId
	}
	return 0
}

func init() {
	proto.RegisterType((*PerpetualPosition)(nil), "nemo_network.subaccounts.PerpetualPosition")
}

func init() {
	proto.RegisterFile("nemo_network/subaccounts/perpetual_position.proto", fileDescriptor_211b9880e1a5e0d8)
}

var fileDescriptor_211b9880e1a5e0d8 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0x82, 0x10, 0x0a, 0xed, 0x40, 0xc4, 0x10, 0x31, 0xb8, 0x85, 0xa9, 0x4b, 0x63,
	0x21, 0x58, 0x91, 0x50, 0x26, 0xb2, 0x55, 0x65, 0x40, 0x62, 0x89, 0x1c, 0xdb, 0xa4, 0x16, 0x89,
	0x6f, 0x9a, 0xd8, 0x50, 0x78, 0x0a, 0x1e, 0xab, 0x63, 0x47, 0xc4, 0x50, 0xa1, 0xe4, 0x2d, 0x98,
	0x50, 0xd2, 0x1f, 0xca, 0xc6, 0xd0, 0xcd, 0xb2, 0xcf, 0xf9, 0x3e, 0x5f, 0xe9, 0xda, 0x17, 0x4a,
	0xa4, 0x10, 0x2a, 0xa1, 0x5f, 0x20, 0x7f, 0x22, 0x85, 0x89, 0x28, 0x63, 0x60, 0x94, 0x2e, 0x48,
	0x26, 0xf2, 0x4c, 0x68, 0x43, 0x93, 0x30, 0x83, 0x42, 0x6a, 0x09, 0xca, 0xcb, 0x72, 0xd0, 0xe0,
	0xb8, 0xdb, 0x15, 0x6f, 0xab, 0x72, 0x7a, 0x12, 0x43, 0x0c, 0x4d, 0x88, 0xd4, 0xa7, 0x65, 0xfe,
	0xfc, 0xbb, 0x65, 0x1f, 0x0f, 0xd7, 0xb0, 0xe1, 0x8a, 0xe5, 0x9c, 0xd9, 0xed, 0x5f, 0x83, 0xe4,
	0x2e, 0xea, 0xa1, 0x7e, 0x67, 0x74, 0xb4, 0xb9, 0x0b, 0xb8, 0xc3, 0xed, 0xc3, 0x89, 0xa1, 0x4a,
	0x9b, 0xb4, 0x70, 0x5b, 0x3d, 0xd4, 0x6f, 0xfb, 0xb7, 0xb3, 0x45, 0xd7, 0xfa, 0x5c, 0x74, 0x6f,
	0x62, 0xa9, 0xc7, 0x26, 0xf2, 0x18, 0xa4, 0xa4, 0xfe, 0xcd, 0x60, 0x3d, 0xc0, 0xf3, 0xd5, 0x80,
	0x8d, 0xa9, 0x54, 0xa4, 0x51, 0x33, 0x48, 0x08, 0xd7, 0xaf, 0x99, 0x28, 0xbc, 0x3b, 0x91, 0x4b,
	0x9a, 0xc8, 0x37, 0x1a, 0x25, 0x22, 0x50, 0x7a, 0xb4, 0x21, 0x3b, 0xa9, 0xdd, 0x79, 0x34, 0x8a,
	0x4b, 0x15, 0x87, 0x52, 0x71, 0x31, 0x75, 0xf7, 0x76, 0xac, 0x6a, 0xaf, 0xf0, 0x41, 0x4d, 0xaf,
	0x75, 0x13, 0x03, 0x5a, 0x84, 0x11, 0x4d, 0xa8, 0x62, 0xc2, 0xdd, 0xdf, 0xb5, 0xae, 0xc1, 0xfb,
	0x4b, 0xba, 0x7f, 0x3f, 0x2b, 0x31, 0x9a, 0x97, 0x18, 0x7d, 0x95, 0x18, 0xbd, 0x57, 0xd8, 0x9a,
	0x57, 0xd8, 0xfa, 0xa8, 0xb0, 0xf5, 0x70, 0xfd, 0x7f, 0xd3, 0xf4, 0xcf, 0x62, 0x34, 0xda, 0xe8,
	0xa0, 0x79, 0xbd, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xea, 0x5f, 0x26, 0xaf, 0x41, 0x02, 0x00,
	0x00,
}

func (m *PerpetualPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerpetualPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerpetualPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.QuoteBalance.Size()
		i -= size
		if _, err := m.QuoteBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualPosition(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.FundingIndex.Size()
		i -= size
		if _, err := m.FundingIndex.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualPosition(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Quantums.Size()
		i -= size
		if _, err := m.Quantums.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualPosition(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PerpetualId != 0 {
		i = encodeVarintPerpetualPosition(dAtA, i, uint64(m.PerpetualId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPerpetualPosition(dAtA []byte, offset int, v uint64) int {
	offset -= sovPerpetualPosition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PerpetualPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PerpetualId != 0 {
		n += 1 + sovPerpetualPosition(uint64(m.PerpetualId))
	}
	l = m.Quantums.Size()
	n += 1 + l + sovPerpetualPosition(uint64(l))
	l = m.FundingIndex.Size()
	n += 1 + l + sovPerpetualPosition(uint64(l))
	l = m.QuoteBalance.Size()
	n += 1 + l + sovPerpetualPosition(uint64(l))
	return n
}

func sovPerpetualPosition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPerpetualPosition(x uint64) (n int) {
	return sovPerpetualPosition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PerpetualPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerpetualPosition
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
			return fmt.Errorf("proto: PerpetualPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerpetualPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualId", wireType)
			}
			m.PerpetualId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualPosition
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
				return fmt.Errorf("proto: wrong wireType = %d for field Quantums", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualPosition
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
				return ErrInvalidLengthPerpetualPosition
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantums.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundingIndex", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualPosition
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
				return ErrInvalidLengthPerpetualPosition
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FundingIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteBalance", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualPosition
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
				return ErrInvalidLengthPerpetualPosition
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerpetualPosition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPerpetualPosition
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
func skipPerpetualPosition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPerpetualPosition
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
					return 0, ErrIntOverflowPerpetualPosition
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
					return 0, ErrIntOverflowPerpetualPosition
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
				return 0, ErrInvalidLengthPerpetualPosition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPerpetualPosition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPerpetualPosition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPerpetualPosition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPerpetualPosition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPerpetualPosition = fmt.Errorf("proto: unexpected end of group")
)
