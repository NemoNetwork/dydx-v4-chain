// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo-network/ratelimit/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

// MsgSetLimitParams is the Msg/SetLimitParams request type.
type MsgSetLimitParams struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// Defines the parameters to set. All parameters must be supplied.
	LimitParams LimitParams `protobuf:"bytes,2,opt,name=limit_params,json=limitParams,proto3" json:"limit_params"`
}

func (m *MsgSetLimitParams) Reset()         { *m = MsgSetLimitParams{} }
func (m *MsgSetLimitParams) String() string { return proto.CompactTextString(m) }
func (*MsgSetLimitParams) ProtoMessage()    {}
func (*MsgSetLimitParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c12b4609ad9be85, []int{0}
}
func (m *MsgSetLimitParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetLimitParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetLimitParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetLimitParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetLimitParams.Merge(m, src)
}
func (m *MsgSetLimitParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetLimitParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetLimitParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetLimitParams proto.InternalMessageInfo

func (m *MsgSetLimitParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSetLimitParams) GetLimitParams() LimitParams {
	if m != nil {
		return m.LimitParams
	}
	return LimitParams{}
}

// MsgSetLimitParamsResponse is the Msg/SetLimitParams response type.
type MsgSetLimitParamsResponse struct {
}

func (m *MsgSetLimitParamsResponse) Reset()         { *m = MsgSetLimitParamsResponse{} }
func (m *MsgSetLimitParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetLimitParamsResponse) ProtoMessage()    {}
func (*MsgSetLimitParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c12b4609ad9be85, []int{1}
}
func (m *MsgSetLimitParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetLimitParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetLimitParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetLimitParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetLimitParamsResponse.Merge(m, src)
}
func (m *MsgSetLimitParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetLimitParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetLimitParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetLimitParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetLimitParams)(nil), "nemo_network.ratelimit.MsgSetLimitParams")
	proto.RegisterType((*MsgSetLimitParamsResponse)(nil), "nemo_network.ratelimit.MsgSetLimitParamsResponse")
}

func init() { proto.RegisterFile("nemo_network/ratelimit/tx.proto", fileDescriptor_3c12b4609ad9be85) }

var fileDescriptor_3c12b4609ad9be85 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4a, 0x2c, 0x49, 0xcd, 0xc9, 0xcc,
	0xcd, 0x2c, 0xd1, 0x2f, 0xa9, 0xd0, 0x03, 0x8b, 0x0a, 0x89, 0x21, 0x2b, 0xd0, 0x83, 0x2b, 0x90,
	0x12, 0x4f, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x33, 0x04, 0x51,
	0x10, 0x0d, 0x52, 0x9a, 0x38, 0x4c, 0x04, 0x93, 0xf1, 0x05, 0x89, 0x45, 0x89, 0xb9, 0xc5, 0x50,
	0xa5, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6, 0x3e, 0x88, 0x05, 0x11, 0x55, 0xea, 0x67, 0xe4,
	0x12, 0xf4, 0x2d, 0x4e, 0x0f, 0x4e, 0x2d, 0xf1, 0x01, 0x69, 0x09, 0x00, 0xeb, 0x10, 0x92, 0xe1,
	0xe2, 0x4c, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca, 0x2c, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x42, 0x08, 0x08, 0xf9, 0x70, 0xf1, 0x20, 0x9b, 0x2f, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d,
	0xa4, 0xac, 0x87, 0xdd, 0xf1, 0x7a, 0x48, 0x06, 0x3b, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0xc4,
	0x9d, 0x83, 0x10, 0xb2, 0xe2, 0x6b, 0x7a, 0xbe, 0x41, 0x0b, 0x61, 0xba, 0x92, 0x34, 0x97, 0x24,
	0x86, 0x83, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x8d, 0x4a, 0xb9, 0x98, 0x7d, 0x8b,
	0xd3, 0x85, 0xf2, 0xb8, 0xf8, 0xd0, 0x5c, 0xac, 0x89, 0xcb, 0x76, 0x0c, 0xb3, 0xa4, 0x0c, 0x89,
	0x56, 0x0a, 0xb3, 0xd6, 0x29, 0xf4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0xac, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x51, 0xe2, 0xa2, 0xcc,
	0x44, 0x37, 0x39, 0x23, 0x31, 0x33, 0x4f, 0x1f, 0x2e, 0x52, 0x81, 0x1c, 0xe3, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0x60, 0x39, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x6a, 0x27, 0x56,
	0x18, 0x02, 0x00, 0x00,
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
	// SetLimitParams sets a `LimitParams` object in state.
	SetLimitParams(ctx context.Context, in *MsgSetLimitParams, opts ...grpc.CallOption) (*MsgSetLimitParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetLimitParams(ctx context.Context, in *MsgSetLimitParams, opts ...grpc.CallOption) (*MsgSetLimitParamsResponse, error) {
	out := new(MsgSetLimitParamsResponse)
	err := c.cc.Invoke(ctx, "/nemo_network.ratelimit.Msg/SetLimitParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SetLimitParams sets a `LimitParams` object in state.
	SetLimitParams(context.Context, *MsgSetLimitParams) (*MsgSetLimitParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetLimitParams(ctx context.Context, req *MsgSetLimitParams) (*MsgSetLimitParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLimitParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetLimitParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetLimitParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetLimitParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo_network.ratelimit.Msg/SetLimitParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetLimitParams(ctx, req.(*MsgSetLimitParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nemo_network.ratelimit.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLimitParams",
			Handler:    _Msg_SetLimitParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nemo_network/ratelimit/tx.proto",
}

func (m *MsgSetLimitParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetLimitParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetLimitParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.LimitParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetLimitParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetLimitParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetLimitParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSetLimitParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.LimitParams.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSetLimitParamsResponse) Size() (n int) {
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
func (m *MsgSetLimitParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetLimitParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetLimitParams: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitParams", wireType)
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
			if err := m.LimitParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgSetLimitParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetLimitParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetLimitParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
