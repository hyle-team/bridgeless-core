// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/types/v1/indexer.proto

package types

import (
	fmt "fmt"
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

// TxResult is the value stored in eth tx indexer
type TxResult struct {
	// height of the blockchain
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// tx_index of the cosmos transaction
	TxIndex uint32 `protobuf:"varint,2,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	// msg_index in a batch transaction
	MsgIndex uint32 `protobuf:"varint,3,opt,name=msg_index,json=msgIndex,proto3" json:"msg_index,omitempty"`
	// eth_tx_index is the index in the list of valid eth tx in the block,
	// aka. the transaction list returned by eth_getBlock api.
	EthTxIndex int32 `protobuf:"varint,4,opt,name=eth_tx_index,json=ethTxIndex,proto3" json:"eth_tx_index,omitempty"`
	// failed is true if the eth transaction did not go succeed
	Failed bool `protobuf:"varint,5,opt,name=failed,proto3" json:"failed,omitempty"`
	// gas_used by the transaction. If it exceeds the block gas limit,
	// it's set to gas limit, which is what's actually deducted by ante handler.
	GasUsed uint64 `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// cumulative_gas_used specifies the cumulated amount of gas used for all
	// processed messages within the current batch transaction.
	CumulativeGasUsed uint64 `protobuf:"varint,7,opt,name=cumulative_gas_used,json=cumulativeGasUsed,proto3" json:"cumulative_gas_used,omitempty"`
}

func (m *TxResult) Reset()         { *m = TxResult{} }
func (m *TxResult) String() string { return proto.CompactTextString(m) }
func (*TxResult) ProtoMessage()    {}
func (*TxResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1197e10a8be8ed28, []int{0}
}
func (m *TxResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResult.Merge(m, src)
}
func (m *TxResult) XXX_Size() int {
	return m.Size()
}
func (m *TxResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResult.DiscardUnknown(m)
}

var xxx_messageInfo_TxResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TxResult)(nil), "ethermint.types.v1.TxResult")
}

func init() { proto.RegisterFile("ethermint/types/v1/indexer.proto", fileDescriptor_1197e10a8be8ed28) }

var fileDescriptor_1197e10a8be8ed28 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x33, 0x7f, 0xdb, 0x34, 0xff, 0xa0, 0x0b, 0xa3, 0x94, 0xa8, 0x10, 0x07, 0x57, 0x41,
	0x68, 0x42, 0x71, 0xe7, 0x52, 0x04, 0x71, 0x1b, 0xea, 0xc6, 0x4d, 0x48, 0x9b, 0xeb, 0xcc, 0x40,
	0xa6, 0x53, 0x32, 0x37, 0x25, 0x7d, 0x03, 0x97, 0x3e, 0x82, 0x8f, 0xe3, 0xb2, 0x4b, 0x97, 0xd2,
	0xe2, 0x7b, 0x48, 0x27, 0xa1, 0xee, 0xee, 0xe1, 0xfb, 0x2e, 0x07, 0x0e, 0x65, 0x80, 0x02, 0x2a,
	0x25, 0x17, 0x98, 0xe0, 0x7a, 0x09, 0x26, 0x59, 0x4d, 0x12, 0xb9, 0x28, 0xa0, 0x81, 0x2a, 0x5e,
	0x56, 0x1a, 0xb5, 0xef, 0x1f, 0x8c, 0xd8, 0x1a, 0xf1, 0x6a, 0x72, 0x71, 0xc6, 0x35, 0xd7, 0x16,
	0x27, 0xfb, 0xab, 0x35, 0xaf, 0x7f, 0x08, 0xf5, 0xa6, 0x4d, 0x0a, 0xa6, 0x2e, 0xd1, 0x1f, 0x51,
	0x57, 0x80, 0xe4, 0x02, 0x03, 0xc2, 0x48, 0xd4, 0x4b, 0xbb, 0xe4, 0x9f, 0x53, 0x0f, 0x9b, 0xcc,
	0x56, 0x04, 0xff, 0x18, 0x89, 0x8e, 0xd3, 0x21, 0x36, 0x4f, 0xfb, 0xe8, 0x5f, 0xd2, 0xff, 0xca,
	0xf0, 0x8e, 0xf5, 0x2c, 0xf3, 0x94, 0xe1, 0x2d, 0x64, 0xf4, 0x08, 0x50, 0x64, 0x87, 0xdf, 0x3e,
	0x23, 0xd1, 0x20, 0xa5, 0x80, 0x62, 0xda, 0xbd, 0x8f, 0xa8, 0xfb, 0x9a, 0xcb, 0x12, 0x8a, 0x60,
	0xc0, 0x48, 0xe4, 0xa5, 0x5d, 0xda, 0x37, 0xf2, 0xdc, 0x64, 0xb5, 0x81, 0x22, 0x70, 0x19, 0x89,
	0xfa, 0xe9, 0x90, 0xe7, 0xe6, 0xd9, 0x40, 0xe1, 0xc7, 0xf4, 0x74, 0x5e, 0xab, 0xba, 0xcc, 0x51,
	0xae, 0x20, 0x3b, 0x58, 0x43, 0x6b, 0x9d, 0xfc, 0xa1, 0xc7, 0xd6, 0xbf, 0xeb, 0xbf, 0x7d, 0x5c,
	0x39, 0xf7, 0x0f, 0x9f, 0xdb, 0x90, 0x6c, 0xb6, 0x21, 0xf9, 0xde, 0x86, 0xe4, 0x7d, 0x17, 0x3a,
	0x9b, 0x5d, 0xe8, 0x7c, 0xed, 0x42, 0xe7, 0xe5, 0x86, 0x4b, 0x14, 0xf5, 0x2c, 0x9e, 0x6b, 0x95,
	0x88, 0x75, 0x09, 0x63, 0x84, 0x5c, 0x25, 0xb3, 0x4a, 0x16, 0x1c, 0x4a, 0x30, 0x66, 0x3c, 0xd7,
	0x15, 0xb4, 0x43, 0xcf, 0x5c, 0x3b, 0xda, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x26,
	0x7e, 0xac, 0x82, 0x01, 0x00, 0x00,
}

func (m *TxResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CumulativeGasUsed != 0 {
		i = encodeVarintIndexer(dAtA, i, uint64(m.CumulativeGasUsed))
		i--
		dAtA[i] = 0x38
	}
	if m.GasUsed != 0 {
		i = encodeVarintIndexer(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.Failed {
		i--
		if m.Failed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.EthTxIndex != 0 {
		i = encodeVarintIndexer(dAtA, i, uint64(m.EthTxIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.MsgIndex != 0 {
		i = encodeVarintIndexer(dAtA, i, uint64(m.MsgIndex))
		i--
		dAtA[i] = 0x18
	}
	if m.TxIndex != 0 {
		i = encodeVarintIndexer(dAtA, i, uint64(m.TxIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintIndexer(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIndexer(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndexer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovIndexer(uint64(m.Height))
	}
	if m.TxIndex != 0 {
		n += 1 + sovIndexer(uint64(m.TxIndex))
	}
	if m.MsgIndex != 0 {
		n += 1 + sovIndexer(uint64(m.MsgIndex))
	}
	if m.EthTxIndex != 0 {
		n += 1 + sovIndexer(uint64(m.EthTxIndex))
	}
	if m.Failed {
		n += 2
	}
	if m.GasUsed != 0 {
		n += 1 + sovIndexer(uint64(m.GasUsed))
	}
	if m.CumulativeGasUsed != 0 {
		n += 1 + sovIndexer(uint64(m.CumulativeGasUsed))
	}
	return n
}

func sovIndexer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndexer(x uint64) (n int) {
	return sovIndexer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexer
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
			return fmt.Errorf("proto: TxResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgIndex", wireType)
			}
			m.MsgIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthTxIndex", wireType)
			}
			m.EthTxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthTxIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Failed = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeGasUsed", wireType)
			}
			m.CumulativeGasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CumulativeGasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIndexer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndexer
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
func skipIndexer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndexer
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
					return 0, ErrIntOverflowIndexer
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
					return 0, ErrIntOverflowIndexer
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
				return 0, ErrInvalidLengthIndexer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIndexer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIndexer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIndexer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndexer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIndexer = fmt.Errorf("proto: unexpected end of group")
)
