// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo-network/indexer/protocol/v1/perpetual.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
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

// Market type of perpetual.
// Defined in perpetual.
type PerpetualMarketType int32

const (
	// Unspecified market type.
	PerpetualMarketType_PERPETUAL_MARKET_TYPE_UNSPECIFIED PerpetualMarketType = 0
	// Market type for cross margin perpetual markets.
	PerpetualMarketType_PERPETUAL_MARKET_TYPE_CROSS PerpetualMarketType = 1
	// Market type for isolated margin perpetual markets.
	PerpetualMarketType_PERPETUAL_MARKET_TYPE_ISOLATED PerpetualMarketType = 2
)

var PerpetualMarketType_name = map[int32]string{
	0: "PERPETUAL_MARKET_TYPE_UNSPECIFIED",
	1: "PERPETUAL_MARKET_TYPE_CROSS",
	2: "PERPETUAL_MARKET_TYPE_ISOLATED",
}

var PerpetualMarketType_value = map[string]int32{
	"PERPETUAL_MARKET_TYPE_UNSPECIFIED": 0,
	"PERPETUAL_MARKET_TYPE_CROSS":       1,
	"PERPETUAL_MARKET_TYPE_ISOLATED":    2,
}

func (x PerpetualMarketType) String() string {
	return proto.EnumName(PerpetualMarketType_name, int32(x))
}

func (PerpetualMarketType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ebf0c6c8fa38d8c8, []int{0}
}

func init() {
	proto.RegisterEnum("nemo-network.indexer.protocol.v1.PerpetualMarketType", PerpetualMarketType_name, PerpetualMarketType_value)
}

func init() {
	proto.RegisterFile("nemo-network/indexer/protocol/v1/perpetual.proto", fileDescriptor_ebf0c6c8fa38d8c8)
}

var fileDescriptor_ebf0c6c8fa38d8c8 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0xcf, 0xcc, 0x4b, 0x49, 0xad, 0x48, 0x2d,
	0xd2, 0x87, 0x0b, 0x94, 0x19, 0xea, 0x17, 0xa4, 0x16, 0x15, 0xa4, 0x96, 0x94, 0x26, 0xe6, 0xe8,
	0x81, 0x45, 0x85, 0x14, 0x90, 0x75, 0xe8, 0x41, 0x75, 0xe8, 0xc1, 0x05, 0xca, 0x0c, 0xb5, 0x1a,
	0x19, 0xb9, 0x84, 0x03, 0x60, 0xba, 0x7c, 0x13, 0x8b, 0xb2, 0x53, 0x4b, 0x42, 0x2a, 0x0b, 0x52,
	0x85, 0x54, 0xb9, 0x14, 0x03, 0x5c, 0x83, 0x02, 0x5c, 0x43, 0x42, 0x1d, 0x7d, 0xe2, 0x7d, 0x1d,
	0x83, 0xbc, 0x5d, 0x43, 0xe2, 0x43, 0x22, 0x03, 0x5c, 0xe3, 0x43, 0xfd, 0x82, 0x03, 0x5c, 0x9d,
	0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x84, 0xe4, 0xb9, 0xa4, 0xb1, 0x2b, 0x73, 0x0e, 0xf2,
	0x0f, 0x0e, 0x16, 0x60, 0x14, 0x52, 0xe2, 0x92, 0xc3, 0xae, 0xc0, 0x33, 0xd8, 0xdf, 0xc7, 0x31,
	0xc4, 0xd5, 0x45, 0x80, 0xc9, 0x29, 0xf6, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0x9c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x51, 0x3c, 0x5f,
	0x66, 0xa2, 0x9b, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x8f, 0x37, 0x38, 0x4a, 0x2a, 0x0b, 0x52, 0x8b,
	0x93, 0xd8, 0xc0, 0x42, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xc4, 0x7a, 0xc4, 0x3f,
	0x01, 0x00, 0x00,
}
