// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tracking/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type Params struct {
	BorrowEventName      string `protobuf:"bytes,1,opt,name=borrow_event_name,json=borrowEventName,proto3" json:"borrow_event_name,omitempty"`
	ContractAddress      string `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	SenderAddress        string `protobuf:"bytes,3,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	LiquidatorAddress    string `protobuf:"bytes,4,opt,name=liquidator_address,json=liquidatorAddress,proto3" json:"liquidator_address,omitempty"`
	LiquidationEventName string `protobuf:"bytes,5,opt,name=liquidation_event_name,json=liquidationEventName,proto3" json:"liquidation_event_name,omitempty"`
	OracleAddress        string `protobuf:"bytes,6,opt,name=oracle_address,json=oracleAddress,proto3" json:"oracle_address,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2968bee4da122ea2, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetBorrowEventName() string {
	if m != nil {
		return m.BorrowEventName
	}
	return ""
}

func (m *Params) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Params) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *Params) GetLiquidatorAddress() string {
	if m != nil {
		return m.LiquidatorAddress
	}
	return ""
}

func (m *Params) GetLiquidationEventName() string {
	if m != nil {
		return m.LiquidationEventName
	}
	return ""
}

func (m *Params) GetOracleAddress() string {
	if m != nil {
		return m.OracleAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "core.tracking.Params")
}

func init() { proto.RegisterFile("tracking/params.proto", fileDescriptor_2968bee4da122ea2) }

var fileDescriptor_2968bee4da122ea2 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x9b, 0xaa, 0x05, 0x17, 0x5a, 0x6d, 0xa8, 0x52, 0x3c, 0x2c, 0x22, 0x14, 0x54, 0x68,
	0x17, 0xff, 0x5c, 0x3c, 0x2a, 0x78, 0x15, 0xd1, 0x9b, 0x97, 0xb0, 0x49, 0x86, 0x74, 0x31, 0xd9,
	0x89, 0xbb, 0xdb, 0x6a, 0xdf, 0x42, 0x7c, 0x2a, 0x8f, 0x3d, 0x7a, 0x94, 0xe4, 0x45, 0x24, 0xbb,
	0x26, 0xe9, 0x2d, 0xfc, 0xbe, 0xdf, 0x4c, 0x3e, 0x66, 0xc9, 0x81, 0x51, 0x3c, 0x7a, 0x15, 0x32,
	0x61, 0x39, 0x57, 0x3c, 0xd3, 0xb3, 0x5c, 0xa1, 0x41, 0xbf, 0x1f, 0xa1, 0x82, 0x59, 0x9d, 0x1d,
	0x8d, 0x12, 0x4c, 0xd0, 0x26, 0xac, 0xfa, 0x72, 0xd2, 0xc9, 0x57, 0x97, 0xf4, 0x1e, 0xed, 0x94,
	0x7f, 0x4e, 0x86, 0x21, 0x2a, 0x85, 0xef, 0x01, 0x2c, 0x41, 0x9a, 0x40, 0xf2, 0x0c, 0xc6, 0xde,
	0xb1, 0x77, 0xba, 0xfb, 0xb4, 0xe7, 0x82, 0xfb, 0x8a, 0x3f, 0xf0, 0x0c, 0xfc, 0x33, 0xb2, 0x1f,
	0xa1, 0xac, 0x76, 0x9b, 0x80, 0xc7, 0xb1, 0x02, 0xad, 0xc7, 0x5d, 0xa7, 0xd6, 0xfc, 0xd6, 0x61,
	0x7f, 0x42, 0x06, 0x1a, 0x64, 0x0c, 0xaa, 0x11, 0xb7, 0xac, 0xd8, 0x77, 0xb4, 0xd6, 0xa6, 0xc4,
	0x4f, 0xc5, 0xdb, 0x42, 0xc4, 0xdc, 0x60, 0xab, 0x6e, 0x5b, 0x75, 0xd8, 0x26, 0xb5, 0x7e, 0x4d,
	0x0e, 0x6b, 0x28, 0x50, 0x6e, 0x36, 0xde, 0xb1, 0x23, 0xa3, 0x8d, 0xb4, 0xad, 0x3d, 0x21, 0x03,
	0x54, 0x3c, 0x4a, 0xa1, 0xf9, 0x41, 0xcf, 0x75, 0x71, 0xf4, 0x7f, 0xf9, 0xdd, 0xf3, 0x77, 0x41,
	0xbd, 0x75, 0x41, 0xbd, 0xdf, 0x82, 0x7a, 0x9f, 0x25, 0xed, 0xac, 0x4b, 0xda, 0xf9, 0x29, 0x69,
	0xe7, 0xe5, 0x26, 0x11, 0x66, 0xbe, 0x08, 0x67, 0x11, 0x66, 0x6c, 0xbe, 0x4a, 0x61, 0x6a, 0x80,
	0x67, 0x2c, 0x54, 0x22, 0x4e, 0x20, 0x05, 0xad, 0xa7, 0xd5, 0xcd, 0xd9, 0xf2, 0xe2, 0x92, 0x7d,
	0xb0, 0xe6, 0x55, 0xcc, 0x2a, 0x07, 0x1d, 0xf6, 0xec, 0xc1, 0xaf, 0xfe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x89, 0x0c, 0x00, 0xca, 0xae, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OracleAddress) > 0 {
		i -= len(m.OracleAddress)
		copy(dAtA[i:], m.OracleAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.OracleAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.LiquidationEventName) > 0 {
		i -= len(m.LiquidationEventName)
		copy(dAtA[i:], m.LiquidationEventName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.LiquidationEventName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.LiquidatorAddress) > 0 {
		i -= len(m.LiquidatorAddress)
		copy(dAtA[i:], m.LiquidatorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.LiquidatorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SenderAddress) > 0 {
		i -= len(m.SenderAddress)
		copy(dAtA[i:], m.SenderAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SenderAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BorrowEventName) > 0 {
		i -= len(m.BorrowEventName)
		copy(dAtA[i:], m.BorrowEventName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BorrowEventName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BorrowEventName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.SenderAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.LiquidatorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.LiquidationEventName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.OracleAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowEventName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BorrowEventName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationEventName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidationEventName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
