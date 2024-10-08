// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo_network/clob/block_rate_limit_config.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Defines the block rate limits for CLOB specific operations.
type BlockRateLimitConfiguration struct {
	// How many short term order attempts (successful and failed) are allowed for
	// an account per N blocks. Note that the rate limits are applied
	// in an AND fashion such that an order placement must pass all rate limit
	// configurations.
	//
	// Specifying 0 values disables this rate limit.
	// Deprecated in favor of `max_short_term_orders_and_cancels_per_n_blocks`
	// for v5.x onwards.
	MaxShortTermOrdersPerNBlocks []MaxPerNBlocksRateLimit `protobuf:"bytes,1,rep,name=max_short_term_orders_per_n_blocks,json=maxShortTermOrdersPerNBlocks,proto3" json:"max_short_term_orders_per_n_blocks"` // Deprecated: Do not use.
	// How many stateful order attempts (successful and failed) are allowed for
	// an account per N blocks. Note that the rate limits are applied
	// in an AND fashion such that an order placement must pass all rate limit
	// configurations.
	//
	// Specifying 0 values disables this rate limit.
	MaxStatefulOrdersPerNBlocks []MaxPerNBlocksRateLimit `protobuf:"bytes,2,rep,name=max_stateful_orders_per_n_blocks,json=maxStatefulOrdersPerNBlocks,proto3" json:"max_stateful_orders_per_n_blocks"`
	// How many short term order cancellation attempts (successful and failed) are
	// allowed for an account per N blocks. Note that the rate limits are
	// applied in an AND fashion such that an order cancellation must pass all
	// rate limit configurations.
	//
	// Specifying 0 values disables this rate limit.
	// Deprecated in favor of `max_short_term_orders_and_cancels_per_n_blocks`
	// for v5.x onwards.
	MaxShortTermOrderCancellationsPerNBlocks []MaxPerNBlocksRateLimit `protobuf:"bytes,3,rep,name=max_short_term_order_cancellations_per_n_blocks,json=maxShortTermOrderCancellationsPerNBlocks,proto3" json:"max_short_term_order_cancellations_per_n_blocks"` // Deprecated: Do not use.
	// How many short term order place and cancel attempts (successful and failed)
	// are allowed for an account per N blocks. Note that the rate limits are
	// applied in an AND fashion such that an order placement must pass all rate
	// limit configurations.
	//
	// Specifying 0 values disables this rate limit.
	MaxShortTermOrdersAndCancelsPerNBlocks []MaxPerNBlocksRateLimit `protobuf:"bytes,4,rep,name=max_short_term_orders_and_cancels_per_n_blocks,json=maxShortTermOrdersAndCancelsPerNBlocks,proto3" json:"max_short_term_orders_and_cancels_per_n_blocks"`
}

func (m *BlockRateLimitConfiguration) Reset()         { *m = BlockRateLimitConfiguration{} }
func (m *BlockRateLimitConfiguration) String() string { return proto.CompactTextString(m) }
func (*BlockRateLimitConfiguration) ProtoMessage()    {}
func (*BlockRateLimitConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a78358196121352, []int{0}
}
func (m *BlockRateLimitConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockRateLimitConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockRateLimitConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockRateLimitConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRateLimitConfiguration.Merge(m, src)
}
func (m *BlockRateLimitConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *BlockRateLimitConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRateLimitConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRateLimitConfiguration proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *BlockRateLimitConfiguration) GetMaxShortTermOrdersPerNBlocks() []MaxPerNBlocksRateLimit {
	if m != nil {
		return m.MaxShortTermOrdersPerNBlocks
	}
	return nil
}

func (m *BlockRateLimitConfiguration) GetMaxStatefulOrdersPerNBlocks() []MaxPerNBlocksRateLimit {
	if m != nil {
		return m.MaxStatefulOrdersPerNBlocks
	}
	return nil
}

// Deprecated: Do not use.
func (m *BlockRateLimitConfiguration) GetMaxShortTermOrderCancellationsPerNBlocks() []MaxPerNBlocksRateLimit {
	if m != nil {
		return m.MaxShortTermOrderCancellationsPerNBlocks
	}
	return nil
}

func (m *BlockRateLimitConfiguration) GetMaxShortTermOrdersAndCancelsPerNBlocks() []MaxPerNBlocksRateLimit {
	if m != nil {
		return m.MaxShortTermOrdersAndCancelsPerNBlocks
	}
	return nil
}

// Defines a rate limit over a specific number of blocks.
type MaxPerNBlocksRateLimit struct {
	// How many blocks the rate limit is over.
	// Specifying 0 is invalid.
	NumBlocks uint32 `protobuf:"varint,1,opt,name=num_blocks,json=numBlocks,proto3" json:"num_blocks,omitempty"`
	// What the limit is for `num_blocks`.
	// Specifying 0 is invalid.
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *MaxPerNBlocksRateLimit) Reset()         { *m = MaxPerNBlocksRateLimit{} }
func (m *MaxPerNBlocksRateLimit) String() string { return proto.CompactTextString(m) }
func (*MaxPerNBlocksRateLimit) ProtoMessage()    {}
func (*MaxPerNBlocksRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a78358196121352, []int{1}
}
func (m *MaxPerNBlocksRateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MaxPerNBlocksRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MaxPerNBlocksRateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MaxPerNBlocksRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxPerNBlocksRateLimit.Merge(m, src)
}
func (m *MaxPerNBlocksRateLimit) XXX_Size() int {
	return m.Size()
}
func (m *MaxPerNBlocksRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxPerNBlocksRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_MaxPerNBlocksRateLimit proto.InternalMessageInfo

func (m *MaxPerNBlocksRateLimit) GetNumBlocks() uint32 {
	if m != nil {
		return m.NumBlocks
	}
	return 0
}

func (m *MaxPerNBlocksRateLimit) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockRateLimitConfiguration)(nil), "nemo_network.clob.BlockRateLimitConfiguration")
	proto.RegisterType((*MaxPerNBlocksRateLimit)(nil), "nemo_network.clob.MaxPerNBlocksRateLimit")
}

func init() {
	proto.RegisterFile("nemo_network/clob/block_rate_limit_config.proto", fileDescriptor_9a78358196121352)
}

var fileDescriptor_9a78358196121352 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xda, 0x30,
	0x18, 0xc6, 0x63, 0x60, 0x48, 0xf3, 0xb4, 0xc3, 0x22, 0x34, 0xa1, 0xb1, 0x65, 0x88, 0xc3, 0xc4,
	0x0e, 0x24, 0xd2, 0x56, 0xf5, 0x5e, 0xb8, 0x96, 0x16, 0xa5, 0x3d, 0xf5, 0x62, 0x39, 0xc6, 0x84,
	0x88, 0xd8, 0x46, 0x8e, 0x43, 0x53, 0xf5, 0x43, 0xb4, 0x87, 0x7e, 0x9b, 0x7e, 0x01, 0x8e, 0x1c,
	0x7b, 0xaa, 0x2a, 0xf8, 0x22, 0x55, 0x1c, 0x5a, 0x22, 0x92, 0x4b, 0xb9, 0x25, 0xfe, 0xf3, 0xfc,
	0x1e, 0x3f, 0xef, 0xfb, 0x42, 0x87, 0x53, 0x26, 0x10, 0xa7, 0xea, 0x5a, 0xc8, 0x99, 0x43, 0x42,
	0xe1, 0x39, 0x5e, 0x28, 0xc8, 0x0c, 0x49, 0xac, 0x28, 0x0a, 0x03, 0x16, 0x28, 0x44, 0x04, 0x9f,
	0x04, 0xbe, 0x3d, 0x97, 0x42, 0x09, 0xf3, 0x5b, 0xfe, 0x82, 0x9d, 0x5e, 0xf8, 0xd1, 0xf0, 0x85,
	0x2f, 0xf4, 0xae, 0x93, 0x7e, 0x65, 0x07, 0x3b, 0x8f, 0x35, 0xd8, 0xea, 0xa7, 0x52, 0x2e, 0x56,
	0xf4, 0x34, 0x15, 0x1a, 0x68, 0x9d, 0x58, 0x62, 0x15, 0x08, 0x6e, 0xde, 0xc2, 0x0e, 0xc3, 0x09,
	0x8a, 0xa6, 0x42, 0x2a, 0xa4, 0xa8, 0x64, 0x48, 0xc8, 0x31, 0x95, 0x11, 0x9a, 0x53, 0x89, 0x38,
	0xd2, 0x2e, 0xa2, 0x26, 0x68, 0x57, 0xbb, 0x5f, 0xfe, 0xfd, 0xb5, 0x0b, 0x54, 0x7b, 0x88, 0x93,
	0x11, 0x95, 0x67, 0x1a, 0x11, 0xbd, 0x33, 0xfa, 0xf5, 0xe5, 0xf3, 0x6f, 0xa3, 0x09, 0xdc, 0x9f,
	0x0c, 0x27, 0x17, 0xa9, 0xf6, 0x25, 0x95, 0xec, 0x5c, 0x2b, 0xef, 0x8e, 0x9b, 0x0b, 0xd8, 0xd6,
	0x70, 0x85, 0x15, 0x9d, 0xc4, 0x61, 0x29, 0xba, 0xf2, 0x51, 0x74, 0x2d, 0x45, 0xbb, 0xad, 0x14,
	0xbc, 0xd5, 0x2d, 0x70, 0x1f, 0x00, 0x74, 0xca, 0x5e, 0x8d, 0x08, 0xe6, 0x84, 0x86, 0xa1, 0x8e,
	0x66, 0xcf, 0x47, 0xf5, 0xd0, 0x08, 0xba, 0x85, 0x08, 0x06, 0x79, 0x4a, 0xce, 0xd6, 0x1d, 0x80,
	0x76, 0x79, 0x31, 0x30, 0x1f, 0x6f, 0xbd, 0xed, 0xb9, 0xaa, 0x1d, 0x96, 0xce, 0x9f, 0x62, 0x59,
	0x4e, 0xf8, 0x38, 0xf3, 0x95, 0x73, 0xd4, 0x19, 0xc2, 0xef, 0xe5, 0x3a, 0xe6, 0x2f, 0x08, 0x79,
	0xcc, 0x76, 0xfd, 0x01, 0xba, 0x5f, 0xdd, 0xcf, 0x3c, 0x66, 0xdb, 0xa7, 0x34, 0xe0, 0x27, 0xdd,
	0xb5, 0xcd, 0x8a, 0xde, 0xc9, 0x7e, 0xfa, 0xa3, 0xe5, 0xda, 0x02, 0xab, 0xb5, 0x05, 0x5e, 0xd6,
	0x16, 0xb8, 0xdf, 0x58, 0xc6, 0x6a, 0x63, 0x19, 0x4f, 0x1b, 0xcb, 0xb8, 0x3a, 0xf6, 0x03, 0x35,
	0x8d, 0x3d, 0x9b, 0x08, 0xa6, 0x67, 0xa1, 0xf7, 0x36, 0x0b, 0x8b, 0xa3, 0x1e, 0x99, 0xe2, 0x80,
	0x3b, 0xba, 0x9d, 0x89, 0x08, 0x9d, 0x24, 0x9b, 0x0f, 0x75, 0x33, 0xa7, 0x91, 0x57, 0xd7, 0xcb,
	0xff, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x56, 0xc8, 0xfc, 0xe8, 0x41, 0x03, 0x00, 0x00,
}

func (m *BlockRateLimitConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockRateLimitConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockRateLimitConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MaxShortTermOrdersAndCancelsPerNBlocks) > 0 {
		for iNdEx := len(m.MaxShortTermOrdersAndCancelsPerNBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxShortTermOrdersAndCancelsPerNBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlockRateLimitConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MaxShortTermOrderCancellationsPerNBlocks) > 0 {
		for iNdEx := len(m.MaxShortTermOrderCancellationsPerNBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxShortTermOrderCancellationsPerNBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlockRateLimitConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.MaxStatefulOrdersPerNBlocks) > 0 {
		for iNdEx := len(m.MaxStatefulOrdersPerNBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxStatefulOrdersPerNBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlockRateLimitConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.MaxShortTermOrdersPerNBlocks) > 0 {
		for iNdEx := len(m.MaxShortTermOrdersPerNBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxShortTermOrdersPerNBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlockRateLimitConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MaxPerNBlocksRateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MaxPerNBlocksRateLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MaxPerNBlocksRateLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintBlockRateLimitConfig(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x10
	}
	if m.NumBlocks != 0 {
		i = encodeVarintBlockRateLimitConfig(dAtA, i, uint64(m.NumBlocks))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlockRateLimitConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlockRateLimitConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockRateLimitConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MaxShortTermOrdersPerNBlocks) > 0 {
		for _, e := range m.MaxShortTermOrdersPerNBlocks {
			l = e.Size()
			n += 1 + l + sovBlockRateLimitConfig(uint64(l))
		}
	}
	if len(m.MaxStatefulOrdersPerNBlocks) > 0 {
		for _, e := range m.MaxStatefulOrdersPerNBlocks {
			l = e.Size()
			n += 1 + l + sovBlockRateLimitConfig(uint64(l))
		}
	}
	if len(m.MaxShortTermOrderCancellationsPerNBlocks) > 0 {
		for _, e := range m.MaxShortTermOrderCancellationsPerNBlocks {
			l = e.Size()
			n += 1 + l + sovBlockRateLimitConfig(uint64(l))
		}
	}
	if len(m.MaxShortTermOrdersAndCancelsPerNBlocks) > 0 {
		for _, e := range m.MaxShortTermOrdersAndCancelsPerNBlocks {
			l = e.Size()
			n += 1 + l + sovBlockRateLimitConfig(uint64(l))
		}
	}
	return n
}

func (m *MaxPerNBlocksRateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumBlocks != 0 {
		n += 1 + sovBlockRateLimitConfig(uint64(m.NumBlocks))
	}
	if m.Limit != 0 {
		n += 1 + sovBlockRateLimitConfig(uint64(m.Limit))
	}
	return n
}

func sovBlockRateLimitConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockRateLimitConfig(x uint64) (n int) {
	return sovBlockRateLimitConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockRateLimitConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockRateLimitConfig
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
			return fmt.Errorf("proto: BlockRateLimitConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockRateLimitConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxShortTermOrdersPerNBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockRateLimitConfig
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
				return ErrInvalidLengthBlockRateLimitConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockRateLimitConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxShortTermOrdersPerNBlocks = append(m.MaxShortTermOrdersPerNBlocks, MaxPerNBlocksRateLimit{})
			if err := m.MaxShortTermOrdersPerNBlocks[len(m.MaxShortTermOrdersPerNBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStatefulOrdersPerNBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockRateLimitConfig
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
				return ErrInvalidLengthBlockRateLimitConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockRateLimitConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxStatefulOrdersPerNBlocks = append(m.MaxStatefulOrdersPerNBlocks, MaxPerNBlocksRateLimit{})
			if err := m.MaxStatefulOrdersPerNBlocks[len(m.MaxStatefulOrdersPerNBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxShortTermOrderCancellationsPerNBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockRateLimitConfig
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
				return ErrInvalidLengthBlockRateLimitConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockRateLimitConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxShortTermOrderCancellationsPerNBlocks = append(m.MaxShortTermOrderCancellationsPerNBlocks, MaxPerNBlocksRateLimit{})
			if err := m.MaxShortTermOrderCancellationsPerNBlocks[len(m.MaxShortTermOrderCancellationsPerNBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxShortTermOrdersAndCancelsPerNBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockRateLimitConfig
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
				return ErrInvalidLengthBlockRateLimitConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockRateLimitConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxShortTermOrdersAndCancelsPerNBlocks = append(m.MaxShortTermOrdersAndCancelsPerNBlocks, MaxPerNBlocksRateLimit{})
			if err := m.MaxShortTermOrdersAndCancelsPerNBlocks[len(m.MaxShortTermOrdersAndCancelsPerNBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockRateLimitConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockRateLimitConfig
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
func (m *MaxPerNBlocksRateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockRateLimitConfig
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
			return fmt.Errorf("proto: MaxPerNBlocksRateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MaxPerNBlocksRateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocks", wireType)
			}
			m.NumBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockRateLimitConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockRateLimitConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlockRateLimitConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockRateLimitConfig
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
func skipBlockRateLimitConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockRateLimitConfig
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
					return 0, ErrIntOverflowBlockRateLimitConfig
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
					return 0, ErrIntOverflowBlockRateLimitConfig
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
				return 0, ErrInvalidLengthBlockRateLimitConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlockRateLimitConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlockRateLimitConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlockRateLimitConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockRateLimitConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlockRateLimitConfig = fmt.Errorf("proto: unexpected end of group")
)
