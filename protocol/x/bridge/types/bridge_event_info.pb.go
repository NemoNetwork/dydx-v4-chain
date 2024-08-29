// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo-network/bridge/bridge_event_info.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// BridgeEventInfo stores information about the most recently processed bridge
// event.
type BridgeEventInfo struct {
	// The next event id (the last processed id plus one) of the logs from the
	// Ethereum contract.
	NextId uint32 `protobuf:"varint,1,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
	// The Ethereum block height of the most recently processed bridge event.
	EthBlockHeight uint64 `protobuf:"varint,2,opt,name=eth_block_height,json=ethBlockHeight,proto3" json:"eth_block_height,omitempty"`
}

func (m *BridgeEventInfo) Reset()         { *m = BridgeEventInfo{} }
func (m *BridgeEventInfo) String() string { return proto.CompactTextString(m) }
func (*BridgeEventInfo) ProtoMessage()    {}
func (*BridgeEventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca20c815789f7707, []int{0}
}
func (m *BridgeEventInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BridgeEventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BridgeEventInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BridgeEventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeEventInfo.Merge(m, src)
}
func (m *BridgeEventInfo) XXX_Size() int {
	return m.Size()
}
func (m *BridgeEventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeEventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeEventInfo proto.InternalMessageInfo

func (m *BridgeEventInfo) GetNextId() uint32 {
	if m != nil {
		return m.NextId
	}
	return 0
}

func (m *BridgeEventInfo) GetEthBlockHeight() uint64 {
	if m != nil {
		return m.EthBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*BridgeEventInfo)(nil), "nemo_network.bridge.BridgeEventInfo")
}

func init() {
	proto.RegisterFile("nemo_network/bridge/bridge_event_info.proto", fileDescriptor_ca20c815789f7707)
}

var fileDescriptor_ca20c815789f7707 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0x2a, 0xca, 0x4c, 0x49, 0x4f, 0x85,
	0x52, 0xf1, 0xa9, 0x65, 0xa9, 0x79, 0x25, 0xf1, 0x99, 0x79, 0x69, 0xf9, 0x7a, 0x60, 0x15, 0x42,
	0xc2, 0xc8, 0x8a, 0xf5, 0x20, 0xaa, 0x94, 0x42, 0xb8, 0xf8, 0x9d, 0xc0, 0x2c, 0x57, 0x90, 0x72,
	0xcf, 0xbc, 0xb4, 0x7c, 0x21, 0x71, 0x2e, 0xf6, 0xbc, 0xd4, 0x8a, 0x92, 0xf8, 0xcc, 0x14, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x36, 0x10, 0xd7, 0x33, 0x45, 0x48, 0x83, 0x4b, 0x20, 0xb5,
	0x24, 0x23, 0x3e, 0x29, 0x27, 0x3f, 0x39, 0x3b, 0x3e, 0x23, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x25, 0x88, 0x2f, 0xb5, 0x24, 0xc3, 0x09, 0x24, 0xec, 0x01, 0x16, 0x75,
	0x0a, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x8b, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x14, 0xc7, 0x97, 0x99, 0xe8, 0x26, 0x67, 0x24,
	0x66, 0xe6, 0xe9, 0xc3, 0x45, 0x2a, 0x60, 0x1e, 0x2a, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03,
	0x4b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xf4, 0xfc, 0xef, 0xf4, 0x00, 0x00, 0x00,
}

func (m *BridgeEventInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BridgeEventInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BridgeEventInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EthBlockHeight != 0 {
		i = encodeVarintBridgeEventInfo(dAtA, i, uint64(m.EthBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.NextId != 0 {
		i = encodeVarintBridgeEventInfo(dAtA, i, uint64(m.NextId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBridgeEventInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovBridgeEventInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BridgeEventInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextId != 0 {
		n += 1 + sovBridgeEventInfo(uint64(m.NextId))
	}
	if m.EthBlockHeight != 0 {
		n += 1 + sovBridgeEventInfo(uint64(m.EthBlockHeight))
	}
	return n
}

func sovBridgeEventInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBridgeEventInfo(x uint64) (n int) {
	return sovBridgeEventInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BridgeEventInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBridgeEventInfo
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
			return fmt.Errorf("proto: BridgeEventInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BridgeEventInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextId", wireType)
			}
			m.NextId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeEventInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthBlockHeight", wireType)
			}
			m.EthBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeEventInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBridgeEventInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBridgeEventInfo
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
func skipBridgeEventInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBridgeEventInfo
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
					return 0, ErrIntOverflowBridgeEventInfo
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
					return 0, ErrIntOverflowBridgeEventInfo
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
				return 0, ErrInvalidLengthBridgeEventInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBridgeEventInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBridgeEventInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBridgeEventInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBridgeEventInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBridgeEventInfo = fmt.Errorf("proto: unexpected end of group")
)
