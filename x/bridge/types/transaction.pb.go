// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bridge/models/transaction.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Transaction struct {
	DepositChainId    string  `protobuf:"bytes,1,opt,name=deposit_chain_id,json=depositChainId,proto3" json:"deposit_chain_id,omitempty"`
	DepositTxHash     string  `protobuf:"bytes,2,opt,name=deposit_tx_hash,json=depositTxHash,proto3" json:"deposit_tx_hash,omitempty"`
	DepositTxIndex    uint64  `protobuf:"varint,3,opt,name=deposit_tx_index,json=depositTxIndex,proto3" json:"deposit_tx_index,omitempty"`
	DepositBlock      uint64  `protobuf:"varint,4,opt,name=deposit_block,json=depositBlock,proto3" json:"deposit_block,omitempty"`
	DepositToken      string  `protobuf:"bytes,5,opt,name=deposit_token,json=depositToken,proto3" json:"deposit_token,omitempty"`
	DepositAmount     string  `protobuf:"bytes,6,opt,name=deposit_amount,json=depositAmount,proto3" json:"deposit_amount,omitempty"`
	Depositor         string  `protobuf:"bytes,7,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Receiver          string  `protobuf:"bytes,8,opt,name=receiver,proto3" json:"receiver,omitempty"`
	WithdrawalChainId string  `protobuf:"bytes,9,opt,name=withdrawal_chain_id,json=withdrawalChainId,proto3" json:"withdrawal_chain_id,omitempty"`
	WithdrawalTxHash  string  `protobuf:"bytes,10,opt,name=withdrawal_tx_hash,json=withdrawalTxHash,proto3" json:"withdrawal_tx_hash,omitempty"`
	WithdrawalToken   string  `protobuf:"bytes,11,opt,name=withdrawal_token,json=withdrawalToken,proto3" json:"withdrawal_token,omitempty"`
	Signature         string  `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`
	IsWrapped         bool    `protobuf:"varint,13,opt,name=is_wrapped,json=isWrapped,proto3" json:"is_wrapped,omitempty"`
	WithdrawalAmount  string  `protobuf:"bytes,14,opt,name=withdrawal_amount,json=withdrawalAmount,proto3" json:"withdrawal_amount,omitempty"`
	CommissionAmount  float32 `protobuf:"fixed32,15,opt,name=commission_amount,json=commissionAmount,proto3" json:"commission_amount,omitempty"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fcaf7a078652736, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetDepositChainId() string {
	if m != nil {
		return m.DepositChainId
	}
	return ""
}

func (m *Transaction) GetDepositTxHash() string {
	if m != nil {
		return m.DepositTxHash
	}
	return ""
}

func (m *Transaction) GetDepositTxIndex() uint64 {
	if m != nil {
		return m.DepositTxIndex
	}
	return 0
}

func (m *Transaction) GetDepositBlock() uint64 {
	if m != nil {
		return m.DepositBlock
	}
	return 0
}

func (m *Transaction) GetDepositToken() string {
	if m != nil {
		return m.DepositToken
	}
	return ""
}

func (m *Transaction) GetDepositAmount() string {
	if m != nil {
		return m.DepositAmount
	}
	return ""
}

func (m *Transaction) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *Transaction) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Transaction) GetWithdrawalChainId() string {
	if m != nil {
		return m.WithdrawalChainId
	}
	return ""
}

func (m *Transaction) GetWithdrawalTxHash() string {
	if m != nil {
		return m.WithdrawalTxHash
	}
	return ""
}

func (m *Transaction) GetWithdrawalToken() string {
	if m != nil {
		return m.WithdrawalToken
	}
	return ""
}

func (m *Transaction) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Transaction) GetIsWrapped() bool {
	if m != nil {
		return m.IsWrapped
	}
	return false
}

func (m *Transaction) GetWithdrawalAmount() string {
	if m != nil {
		return m.WithdrawalAmount
	}
	return ""
}

func (m *Transaction) GetCommissionAmount() float32 {
	if m != nil {
		return m.CommissionAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Transaction)(nil), "core.bridge.Transaction")
}

func init() { proto.RegisterFile("bridge/models/transaction.proto", fileDescriptor_8fcaf7a078652736) }

var fileDescriptor_8fcaf7a078652736 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x51, 0x46, 0xeb, 0x6e, 0x6b, 0x67, 0x38, 0x98, 0xaa, 0x84, 0x0a, 0x04, 0x2a,
	0x8c, 0xd5, 0x2a, 0x1c, 0x38, 0x33, 0x2e, 0xec, 0x48, 0x55, 0x09, 0x89, 0x4b, 0xe4, 0x24, 0x26,
	0xb1, 0x96, 0xf8, 0x05, 0xdb, 0x6d, 0xd3, 0x6f, 0xc1, 0xc7, 0xe2, 0xb8, 0x03, 0x07, 0x8e, 0xa8,
	0xfd, 0x22, 0x28, 0x4e, 0xd2, 0x04, 0x6e, 0xce, 0xef, 0xff, 0xcb, 0xcb, 0x7b, 0x4f, 0x31, 0x7a,
	0xea, 0x2b, 0x11, 0x46, 0x9c, 0xa6, 0x10, 0xf2, 0x44, 0x53, 0xa3, 0x98, 0xd4, 0x2c, 0x30, 0x02,
	0xe4, 0x2c, 0x53, 0x60, 0x00, 0xf7, 0x03, 0x50, 0x7c, 0x56, 0x5a, 0xa3, 0xc7, 0x11, 0x40, 0x94,
	0x70, 0x6a, 0x23, 0x7f, 0xf5, 0x8d, 0x32, 0xb9, 0x2d, 0xbd, 0xd1, 0xeb, 0x00, 0x74, 0x0a, 0x9a,
	0xfa, 0x4c, 0x73, 0xfa, 0x7d, 0xc5, 0xd5, 0x96, 0xae, 0xe7, 0x3e, 0x37, 0x6c, 0x4e, 0x33, 0x16,
	0x09, 0xc9, 0x9a, 0x9a, 0xa3, 0x71, 0x55, 0x86, 0x65, 0x82, 0x32, 0x29, 0xc1, 0xd8, 0x50, 0x57,
	0xe9, 0xa3, 0x08, 0x22, 0xb0, 0x47, 0x5a, 0x9c, 0x4a, 0xfa, 0xec, 0x57, 0x07, 0xf5, 0x97, 0x4d,
	0x77, 0x78, 0x8a, 0x86, 0x21, 0xcf, 0x40, 0x0b, 0xe3, 0x05, 0x31, 0x13, 0xd2, 0x13, 0x21, 0x71,
	0x26, 0xce, 0xb4, 0xb7, 0x38, 0xaf, 0xf8, 0xc7, 0x02, 0xdf, 0x84, 0xf8, 0x25, 0x1a, 0xd4, 0xa6,
	0xc9, 0xbd, 0x98, 0xe9, 0x98, 0x1c, 0x5b, 0xf1, 0xac, 0xc2, 0xcb, 0xfc, 0x13, 0xd3, 0x71, 0xbb,
	0xa2, 0xc9, 0x3d, 0x21, 0x43, 0x9e, 0x93, 0x7b, 0x13, 0x67, 0xda, 0x39, 0x54, 0x5c, 0xe6, 0x37,
	0x05, 0xc5, 0xcf, 0x51, 0xfd, 0xaa, 0xe7, 0x27, 0x10, 0xdc, 0x92, 0x8e, 0xd5, 0x4e, 0x2b, 0x78,
	0x5d, 0xb0, 0xb6, 0x64, 0xe0, 0x96, 0x4b, 0x72, 0xdf, 0x7e, 0xb4, 0x96, 0x96, 0x05, 0xc3, 0x2f,
	0x50, 0x5d, 0xdb, 0x63, 0x29, 0xac, 0xa4, 0x21, 0x27, 0xff, 0xb4, 0xf6, 0xc1, 0x42, 0x3c, 0x46,
	0xbd, 0x0a, 0x80, 0x22, 0x0f, 0xac, 0xd1, 0x00, 0x3c, 0x42, 0x5d, 0xc5, 0x03, 0x2e, 0xd6, 0x5c,
	0x91, 0xae, 0x0d, 0x0f, 0xcf, 0x78, 0x86, 0x1e, 0x6e, 0x84, 0x89, 0x43, 0xc5, 0x36, 0x2c, 0x69,
	0x36, 0xd5, 0xb3, 0xda, 0x45, 0x13, 0xd5, 0xcb, 0x7a, 0x83, 0x70, 0xcb, 0xaf, 0xf7, 0x85, 0xac,
	0x3e, 0x6c, 0x92, 0x6a, 0x65, 0xaf, 0xd0, 0xb0, 0x6d, 0xdb, 0x31, 0xfb, 0xd6, 0x1d, 0xb4, 0x5c,
	0x3b, 0xe9, 0x18, 0xf5, 0xb4, 0x88, 0x24, 0x33, 0x2b, 0xc5, 0xc9, 0x69, 0x39, 0xc2, 0x01, 0xe0,
	0x27, 0x08, 0x09, 0xed, 0x6d, 0x14, 0xcb, 0x32, 0x1e, 0x92, 0xb3, 0x89, 0x33, 0xed, 0x2e, 0x7a,
	0x42, 0x7f, 0x29, 0x01, 0xbe, 0x44, 0xad, 0x56, 0xeb, 0x4d, 0x9d, 0xff, 0xdf, 0x54, 0xb5, 0xac,
	0x4b, 0x74, 0x11, 0x40, 0x9a, 0x0a, 0xad, 0x05, 0xc8, 0x5a, 0x1e, 0x4c, 0x9c, 0xe9, 0xf1, 0x62,
	0xd8, 0x04, 0xa5, 0x7c, 0xfd, 0xf9, 0xe7, 0xce, 0x75, 0xee, 0x76, 0xae, 0xf3, 0x67, 0xe7, 0x3a,
	0x3f, 0xf6, 0xee, 0xd1, 0xdd, 0xde, 0x3d, 0xfa, 0xbd, 0x77, 0x8f, 0xbe, 0xbe, 0x8f, 0x84, 0x89,
	0x57, 0xfe, 0x2c, 0x80, 0x94, 0xc6, 0xdb, 0x84, 0x5f, 0x19, 0xce, 0x52, 0x5a, 0x5e, 0x84, 0x84,
	0x6b, 0x7d, 0x55, 0x5c, 0x0c, 0xba, 0x9e, 0xbf, 0xa5, 0x79, 0x45, 0xa9, 0xd9, 0x66, 0x5c, 0xfb,
	0x27, 0xf6, 0x87, 0x7d, 0xf7, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x25, 0xf8, 0xc0, 0x5b, 0x03,
	0x00, 0x00,
}

func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CommissionAmount != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.CommissionAmount))))
		i--
		dAtA[i] = 0x7d
	}
	if len(m.WithdrawalAmount) > 0 {
		i -= len(m.WithdrawalAmount)
		copy(dAtA[i:], m.WithdrawalAmount)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.WithdrawalAmount)))
		i--
		dAtA[i] = 0x72
	}
	if m.IsWrapped {
		i--
		if m.IsWrapped {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.WithdrawalToken) > 0 {
		i -= len(m.WithdrawalToken)
		copy(dAtA[i:], m.WithdrawalToken)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.WithdrawalToken)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.WithdrawalTxHash) > 0 {
		i -= len(m.WithdrawalTxHash)
		copy(dAtA[i:], m.WithdrawalTxHash)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.WithdrawalTxHash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.WithdrawalChainId) > 0 {
		i -= len(m.WithdrawalChainId)
		copy(dAtA[i:], m.WithdrawalChainId)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.WithdrawalChainId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.DepositAmount) > 0 {
		i -= len(m.DepositAmount)
		copy(dAtA[i:], m.DepositAmount)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.DepositAmount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DepositToken) > 0 {
		i -= len(m.DepositToken)
		copy(dAtA[i:], m.DepositToken)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.DepositToken)))
		i--
		dAtA[i] = 0x2a
	}
	if m.DepositBlock != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.DepositBlock))
		i--
		dAtA[i] = 0x20
	}
	if m.DepositTxIndex != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.DepositTxIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DepositTxHash) > 0 {
		i -= len(m.DepositTxHash)
		copy(dAtA[i:], m.DepositTxHash)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.DepositTxHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DepositChainId) > 0 {
		i -= len(m.DepositChainId)
		copy(dAtA[i:], m.DepositChainId)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.DepositChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepositChainId)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.DepositTxHash)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.DepositTxIndex != 0 {
		n += 1 + sovTransaction(uint64(m.DepositTxIndex))
	}
	if m.DepositBlock != 0 {
		n += 1 + sovTransaction(uint64(m.DepositBlock))
	}
	l = len(m.DepositToken)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.DepositAmount)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.WithdrawalChainId)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.WithdrawalTxHash)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.WithdrawalToken)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.IsWrapped {
		n += 2
	}
	l = len(m.WithdrawalAmount)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.CommissionAmount != 0 {
		n += 5
	}
	return n
}

func sovTransaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositTxIndex", wireType)
			}
			m.DepositTxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositTxIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositBlock", wireType)
			}
			m.DepositBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawalChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawalTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawalToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsWrapped", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
			m.IsWrapped = bool(v != 0)
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawalAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionAmount", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.CommissionAmount = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransaction
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
func skipTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
				return 0, ErrInvalidLengthTransaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransaction = fmt.Errorf("proto: unexpected end of group")
)
