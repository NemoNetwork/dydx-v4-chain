// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo-network/rewards/reward_share.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_nemo-network_v4_chain_protocol_dtypes "github.com/nemo-network/v4-chain/protocol/dtypes"
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

// RewardShare stores the relative weight of rewards that each address is
// entitled to.
type RewardShare struct {
	Address string                                                           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Weight  github_com_nemo-network_v4_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/nemo-network/v4-chain/protocol/dtypes.SerializableInt" json:"weight"`
}

func (m *RewardShare) Reset()         { *m = RewardShare{} }
func (m *RewardShare) String() string { return proto.CompactTextString(m) }
func (*RewardShare) ProtoMessage()    {}
func (*RewardShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d4fef3ba0acce6, []int{0}
}
func (m *RewardShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardShare.Merge(m, src)
}
func (m *RewardShare) XXX_Size() int {
	return m.Size()
}
func (m *RewardShare) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardShare.DiscardUnknown(m)
}

var xxx_messageInfo_RewardShare proto.InternalMessageInfo

func (m *RewardShare) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*RewardShare)(nil), "nemo-network.rewards.RewardShare")
}

func init() {
	proto.RegisterFile("nemo-network/rewards/reward_share.proto", fileDescriptor_a9d4fef3ba0acce6)
}

var fileDescriptor_a9d4fef3ba0acce6 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0x29,
	0x86, 0xd2, 0xf1, 0xc5, 0x19, 0x89, 0x45, 0xa9, 0x7a, 0x60, 0x59, 0x21, 0x11, 0x64, 0x85, 0x7a,
	0x50, 0x85, 0x52, 0x92, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xf1, 0x60, 0x09, 0x7d, 0x08, 0x07,
	0xa2, 0x41, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0x22, 0x0e, 0x62, 0x41, 0x44, 0x95, 0x16, 0x33,
	0x72, 0x71, 0x07, 0x81, 0x35, 0x07, 0x83, 0x0c, 0x17, 0x32, 0xe2, 0x62, 0x4f, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x92, 0xb8, 0xb4, 0x45, 0x57, 0x04,
	0x6a, 0x90, 0x23, 0x44, 0x26, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x3d, 0x08, 0xa6, 0x50, 0x28, 0x81,
	0x8b, 0xad, 0x3c, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0xc7, 0xc9, 0xe3,
	0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0x1d, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0x51, 0xbc, 0x55, 0x66, 0xa2, 0x9b, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x0f, 0x17,
	0x49, 0x29, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x0b, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0xc9, 0xac, 0x4a,
	0x4c, 0xca, 0x49, 0xf5, 0xcc, 0x2b, 0x09, 0x82, 0x9a, 0xeb, 0x14, 0x7c, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c,
	0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x96, 0xc4, 0xdb, 0x51, 0x01, 0x0f, 0x4e, 0xb0, 0x65, 0x49,
	0x6c, 0x60, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xfe, 0x98, 0xd3, 0x73, 0x01,
	0x00, 0x00,
}

func (m *RewardShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRewardShare(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRewardShare(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRewardShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovRewardShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRewardShare(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovRewardShare(uint64(l))
	return n
}

func sovRewardShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRewardShare(x uint64) (n int) {
	return sovRewardShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RewardShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewardShare
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
			return fmt.Errorf("proto: RewardShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardShare
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
				return ErrInvalidLengthRewardShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardShare
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
				return ErrInvalidLengthRewardShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewardShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewardShare
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
func skipRewardShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRewardShare
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
					return 0, ErrIntOverflowRewardShare
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
					return 0, ErrIntOverflowRewardShare
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
				return 0, ErrInvalidLengthRewardShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRewardShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRewardShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRewardShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRewardShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRewardShare = fmt.Errorf("proto: unexpected end of group")
)
