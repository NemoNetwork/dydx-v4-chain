// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo-network/delaymsg/delayed_message.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// DelayedMessage is a message that is delayed until a certain block height.
type DelayedMessage struct {
	// The ID of the delayed message.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The message to be executed.
	Msg *types.Any `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// The block height at which the message should be executed.
	BlockHeight uint32 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *DelayedMessage) Reset()         { *m = DelayedMessage{} }
func (m *DelayedMessage) String() string { return proto.CompactTextString(m) }
func (*DelayedMessage) ProtoMessage()    {}
func (*DelayedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff78478f6237d0fe, []int{0}
}
func (m *DelayedMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelayedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelayedMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelayedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelayedMessage.Merge(m, src)
}
func (m *DelayedMessage) XXX_Size() int {
	return m.Size()
}
func (m *DelayedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DelayedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DelayedMessage proto.InternalMessageInfo

func (m *DelayedMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DelayedMessage) GetMsg() *types.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *DelayedMessage) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*DelayedMessage)(nil), "nemo_network.delaymsg.DelayedMessage")
}

func init() {
	proto.RegisterFile("nemo_network/delaymsg/delayed_message.proto", fileDescriptor_ff78478f6237d0fe)
}

var fileDescriptor_ff78478f6237d0fe = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0x49, 0xcd, 0x49, 0xac, 0xcc, 0x2d,
	0x4e, 0x87, 0x30, 0x52, 0x53, 0xe2, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xf5, 0xc0, 0x2a,
	0x84, 0x44, 0x91, 0x15, 0xeb, 0xc1, 0x14, 0x4b, 0x49, 0xa6, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea,
	0x83, 0x65, 0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x3a, 0x94, 0xb2, 0xb9, 0xf8, 0x5c,
	0x20, 0x46, 0xf9, 0x42, 0x4c, 0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x52, 0xe3, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x52, 0x60,
	0xd4, 0xe0, 0x36, 0x12, 0xd1, 0x83, 0x18, 0xa5, 0x07, 0x33, 0x4a, 0xcf, 0x31, 0xaf, 0x32, 0x08,
	0xa4, 0x40, 0x48, 0x91, 0x8b, 0x27, 0x29, 0x27, 0x3f, 0x39, 0x3b, 0x3e, 0x23, 0x35, 0x33, 0x3d,
	0xa3, 0x44, 0x82, 0x19, 0x6c, 0x02, 0x37, 0x58, 0xcc, 0x03, 0x2c, 0xe4, 0x14, 0x72, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1,
	0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x56, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a,
	0xc9, 0xf9, 0xb9, 0xfa, 0x28, 0x1e, 0x2e, 0x33, 0xd1, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87,
	0x8b, 0x54, 0x20, 0x02, 0xa1, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x2c, 0x65, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0xfd, 0xd8, 0x2f, 0x2a, 0x01, 0x00, 0x00,
}

func (m *DelayedMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelayedMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelayedMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintDelayedMessage(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDelayedMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintDelayedMessage(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelayedMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelayedMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DelayedMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovDelayedMessage(uint64(m.Id))
	}
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovDelayedMessage(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovDelayedMessage(uint64(m.BlockHeight))
	}
	return n
}

func sovDelayedMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelayedMessage(x uint64) (n int) {
	return sovDelayedMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DelayedMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelayedMessage
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
			return fmt.Errorf("proto: DelayedMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelayedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelayedMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelayedMessage
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
				return ErrInvalidLengthDelayedMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelayedMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelayedMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDelayedMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelayedMessage
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
func skipDelayedMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelayedMessage
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
					return 0, ErrIntOverflowDelayedMessage
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
					return 0, ErrIntOverflowDelayedMessage
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
				return 0, ErrInvalidLengthDelayedMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelayedMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelayedMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelayedMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelayedMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelayedMessage = fmt.Errorf("proto: unexpected end of group")
)
