// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bridge/models/chain.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ChainType int32

const (
	ChainType_EVM     ChainType = 0
	ChainType_BITCOIN ChainType = 1
	ChainType_COSMOS  ChainType = 2
	ChainType_OTHER   ChainType = 3
)

var ChainType_name = map[int32]string{
	0: "EVM",
	1: "BITCOIN",
	2: "COSMOS",
	3: "OTHER",
}

var ChainType_value = map[string]int32{
	"EVM":     0,
	"BITCOIN": 1,
	"COSMOS":  2,
	"OTHER":   3,
}

func (x ChainType) String() string {
	return proto.EnumName(ChainType_name, int32(x))
}

func (ChainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ec0e281693ee2200, []int{0}
}

type Chain struct {
	Id   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type ChainType `protobuf:"varint,2,opt,name=type,proto3,enum=core.bridge.ChainType" json:"type,omitempty"`
	// bridge_address is the address of the bridge contract on the chain
	BridgeAddress string `protobuf:"bytes,3,opt,name=bridge_address,json=bridgeAddress,proto3" json:"bridge_address,omitempty"`
	// operator is the address of the operator of the bridge contract
	Operator string `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *Chain) Reset()         { *m = Chain{} }
func (m *Chain) String() string { return proto.CompactTextString(m) }
func (*Chain) ProtoMessage()    {}
func (*Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec0e281693ee2200, []int{0}
}
func (m *Chain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Chain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chain.Merge(m, src)
}
func (m *Chain) XXX_Size() int {
	return m.Size()
}
func (m *Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_Chain proto.InternalMessageInfo

func (m *Chain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chain) GetType() ChainType {
	if m != nil {
		return m.Type
	}
	return ChainType_EVM
}

func (m *Chain) GetBridgeAddress() string {
	if m != nil {
		return m.BridgeAddress
	}
	return ""
}

func (m *Chain) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func init() {
	proto.RegisterEnum("core.bridge.ChainType", ChainType_name, ChainType_value)
	proto.RegisterType((*Chain)(nil), "core.bridge.Chain")
}

func init() { proto.RegisterFile("bridge/models/chain.proto", fileDescriptor_ec0e281693ee2200) }

var fileDescriptor_ec0e281693ee2200 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x6a, 0xf2, 0x40,
	0x14, 0xc5, 0x33, 0xf1, 0xdf, 0xe7, 0xc8, 0x27, 0x61, 0x28, 0x25, 0x86, 0x12, 0xa4, 0x50, 0x10,
	0xc1, 0x0c, 0x5a, 0x0a, 0xdd, 0x56, 0x11, 0xea, 0xc2, 0x06, 0x54, 0xba, 0xe8, 0xa6, 0x4c, 0x92,
	0x69, 0x1c, 0x48, 0x32, 0x69, 0x66, 0x14, 0xf2, 0x02, 0x5d, 0xf7, 0xb1, 0xba, 0x74, 0xd9, 0x65,
	0xd1, 0x17, 0x29, 0x99, 0x04, 0xe9, 0xee, 0xce, 0xf9, 0x9d, 0x39, 0x9c, 0x7b, 0x61, 0xcf, 0xcb,
	0x58, 0x10, 0x52, 0x1c, 0xf3, 0x80, 0x46, 0x02, 0xfb, 0x5b, 0xc2, 0x12, 0x27, 0xcd, 0xb8, 0xe4,
	0xa8, 0xe3, 0xf3, 0x8c, 0x3a, 0x25, 0xb7, 0x7a, 0x21, 0xe7, 0x61, 0x44, 0xb1, 0x42, 0xde, 0xee,
	0x0d, 0x93, 0x24, 0x2f, 0x7d, 0xd6, 0xd0, 0xe7, 0x22, 0xe6, 0x02, 0x7b, 0x44, 0x50, 0xfc, 0xbe,
	0xa3, 0x59, 0x8e, 0xf7, 0x63, 0x8f, 0x4a, 0x32, 0xc6, 0x29, 0x09, 0x59, 0x42, 0x24, 0xe3, 0x55,
	0xa6, 0x75, 0x55, 0xc5, 0x90, 0x94, 0x61, 0x92, 0x24, 0x5c, 0x2a, 0x28, 0x2a, 0x7a, 0x11, 0xf2,
	0x90, 0xab, 0x11, 0x17, 0x53, 0xa9, 0x5e, 0x7f, 0x00, 0xd8, 0x98, 0x15, 0xbd, 0x50, 0x17, 0xea,
	0x2c, 0x30, 0x41, 0x1f, 0x0c, 0xda, 0x2b, 0x9d, 0x05, 0x68, 0x08, 0xeb, 0x32, 0x4f, 0xa9, 0xa9,
	0xf7, 0xc1, 0xa0, 0x3b, 0xb9, 0x74, 0xfe, 0x14, 0x76, 0xd4, 0x8f, 0x4d, 0x9e, 0xd2, 0x95, 0xf2,
	0xa0, 0x1b, 0xd8, 0x2d, 0xc9, 0x2b, 0x09, 0x82, 0x8c, 0x0a, 0x61, 0xd6, 0x54, 0xce, 0xff, 0x52,
	0x7d, 0x28, 0x45, 0x64, 0xc1, 0x7f, 0x3c, 0xa5, 0x19, 0x91, 0x3c, 0x33, 0xeb, 0xca, 0x70, 0x7e,
	0x0f, 0xef, 0x61, 0xfb, 0x9c, 0x8a, 0x5a, 0xb0, 0x36, 0x7f, 0x5e, 0x1a, 0x1a, 0xea, 0xc0, 0xd6,
	0x74, 0xb1, 0x99, 0xb9, 0x8b, 0x27, 0x03, 0x20, 0x08, 0x9b, 0x33, 0x77, 0xbd, 0x74, 0xd7, 0x86,
	0x8e, 0xda, 0xb0, 0xe1, 0x6e, 0x1e, 0xe7, 0x2b, 0xa3, 0x36, 0x75, 0xbf, 0x8e, 0x36, 0x38, 0x1c,
	0x6d, 0xf0, 0x73, 0xb4, 0xc1, 0xe7, 0xc9, 0xd6, 0x0e, 0x27, 0x5b, 0xfb, 0x3e, 0xd9, 0xda, 0xcb,
	0x5d, 0xc8, 0xe4, 0x76, 0xe7, 0x39, 0x3e, 0x8f, 0xf1, 0x36, 0x8f, 0xe8, 0x48, 0x52, 0x12, 0xe3,
	0xb2, 0x53, 0x44, 0x85, 0x18, 0x15, 0x3b, 0xe1, 0xfd, 0x78, 0x52, 0x69, 0xb8, 0x58, 0x46, 0x78,
	0x4d, 0x75, 0x9a, 0xdb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xc4, 0xe8, 0xa8, 0xbf, 0x01,
	0x00, 0x00,
}

func (m *Chain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Chain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintChain(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BridgeAddress) > 0 {
		i -= len(m.BridgeAddress)
		copy(dAtA[i:], m.BridgeAddress)
		i = encodeVarintChain(dAtA, i, uint64(len(m.BridgeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintChain(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintChain(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChain(dAtA []byte, offset int, v uint64) int {
	offset -= sovChain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Chain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovChain(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovChain(uint64(m.Type))
	}
	l = len(m.BridgeAddress)
	if l > 0 {
		n += 1 + l + sovChain(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovChain(uint64(l))
	}
	return n
}

func sovChain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChain(x uint64) (n int) {
	return sovChain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Chain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChain
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
			return fmt.Errorf("proto: Chain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChain
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
				return ErrInvalidLengthChain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChain
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
				return ErrInvalidLengthChain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChain
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
				return ErrInvalidLengthChain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChain
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
func skipChain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChain
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
					return 0, ErrIntOverflowChain
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
					return 0, ErrIntOverflowChain
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
				return 0, ErrInvalidLengthChain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChain = fmt.Errorf("proto: unexpected end of group")
)
