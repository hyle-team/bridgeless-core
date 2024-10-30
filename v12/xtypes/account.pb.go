// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/types/v1/account.proto

package xtypes

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// EthAccount implements the authtypes.AccountI interface and embeds an
// authtypes.BaseAccount type. It is compatible with the auth AccountKeeper.
type EthAccount struct {
	// base_account is an authtypes.BaseAccount
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	// code_hash is the hash calculated from the code contents
	CodeHash string `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty" yaml:"code_hash"`
}

func (m *EthAccount) Reset()      { *m = EthAccount{} }
func (*EthAccount) ProtoMessage() {}
func (*EthAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_4edc057d42a619ef, []int{0}
}
func (m *EthAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthAccount.Merge(m, src)
}
func (m *EthAccount) XXX_Size() int {
	return m.Size()
}
func (m *EthAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EthAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EthAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EthAccount)(nil), "ethermint.types.v1.EthAccount")
}

func init() { proto.RegisterFile("ethermint/types/v1/account.proto", fileDescriptor_4edc057d42a619ef) }

var fileDescriptor_4edc057d42a619ef = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xbd, 0x6e, 0xea, 0x30,
	0x14, 0x8e, 0xef, 0x70, 0x55, 0x42, 0x87, 0x2a, 0x65, 0xa0, 0x54, 0x72, 0xa2, 0x4c, 0x2c, 0x89,
	0x15, 0xba, 0xb1, 0x35, 0x52, 0xa5, 0xb2, 0x32, 0x76, 0xa1, 0x8e, 0x39, 0x8a, 0x51, 0x09, 0x46,
	0xb1, 0x89, 0xe0, 0x0d, 0x3a, 0x76, 0xec, 0xc8, 0x43, 0xf4, 0x21, 0xaa, 0x4e, 0x8c, 0x9d, 0x50,
	0x45, 0x54, 0xa9, 0x33, 0x4f, 0x50, 0x61, 0x5b, 0x94, 0x29, 0xe7, 0x9c, 0xef, 0x2f, 0xfe, 0xdc,
	0x00, 0x14, 0x87, 0xb2, 0x98, 0xcc, 0x14, 0x51, 0xab, 0x39, 0x48, 0x52, 0x25, 0x84, 0x32, 0x26,
	0x16, 0x33, 0x15, 0xcf, 0x4b, 0xa1, 0x84, 0xe7, 0x1d, 0x19, 0xb1, 0x66, 0xc4, 0x55, 0xd2, 0xc1,
	0x4c, 0xc8, 0x42, 0x48, 0x42, 0x17, 0x8a, 0x93, 0x2a, 0xc9, 0x40, 0xd1, 0x44, 0x2f, 0x46, 0xd3,
	0xb9, 0x32, 0xf8, 0x48, 0x6f, 0xc4, 0x2c, 0x16, 0x6a, 0xe5, 0x22, 0x17, 0xe6, 0x7e, 0x98, 0xcc,
	0x35, 0xfc, 0x46, 0xae, 0x7b, 0xa7, 0xf8, 0xad, 0x49, 0xf6, 0x1e, 0xdd, 0xf3, 0x8c, 0x4a, 0x18,
	0xd9, 0x3f, 0x69, 0xa3, 0x00, 0x75, 0x9b, 0xbd, 0x20, 0xb6, 0x4e, 0x3a, 0xc9, 0xc6, 0xc6, 0x29,
	0x95, 0x60, 0x75, 0xe9, 0xf5, 0x66, 0xeb, 0xa3, 0xfd, 0xd6, 0xbf, 0x5c, 0xd1, 0x62, 0xda, 0x0f,
	0x4f, 0x3d, 0xc2, 0x61, 0x33, 0xfb, 0x63, 0x7a, 0x89, 0xdb, 0x60, 0x62, 0x0c, 0x23, 0x4e, 0x25,
	0x6f, 0xff, 0x0b, 0x50, 0xb7, 0x91, 0xb6, 0xf6, 0x5b, 0xff, 0xc2, 0x08, 0x8f, 0x50, 0x38, 0x3c,
	0x3b, 0xcc, 0xf7, 0x54, 0xf2, 0x7e, 0xfa, 0xbc, 0xf6, 0x9d, 0xd7, 0xb5, 0xef, 0xfc, 0xac, 0x7d,
	0xe7, 0xe3, 0x2d, 0xea, 0xe5, 0x13, 0xc5, 0x17, 0x59, 0xcc, 0x44, 0x61, 0x9f, 0x68, 0x3f, 0x91,
	0x1c, 0x3f, 0x91, 0xa5, 0x29, 0xc7, 0x54, 0x66, 0x53, 0x07, 0xe9, 0xe0, 0x7d, 0x87, 0xd1, 0x66,
	0x87, 0xd1, 0xd7, 0x0e, 0xa3, 0x97, 0x1a, 0x3b, 0x9b, 0x1a, 0x3b, 0x9f, 0x35, 0x76, 0x1e, 0xc8,
	0x89, 0x1b, 0x5f, 0x4d, 0x21, 0x52, 0x40, 0x0b, 0x92, 0x95, 0x93, 0x71, 0x0e, 0x53, 0x90, 0x32,
	0x62, 0xa2, 0x04, 0x52, 0x25, 0x3d, 0xb2, 0xd4, 0x9e, 0xd9, 0x7f, 0xdd, 0xdc, 0xcd, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x04, 0xe1, 0x18, 0x7e, 0xc2, 0x01, 0x00, 0x00,
}

func (m *EthAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CodeHash) > 0 {
		i -= len(m.CodeHash)
		copy(dAtA[i:], m.CodeHash)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.CodeHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.CodeHash)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: EthAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)
