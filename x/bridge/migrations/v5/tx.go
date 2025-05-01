package v5

import (
	"fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	"io"
	"math"
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

type transaction struct {
	DepositChainId    string `protobuf:"bytes,1,opt,name=deposit_chain_id,json=depositChainId,proto3" json:"deposit_chain_id,omitempty"`
	DepositTxHash     string `protobuf:"bytes,2,opt,name=deposit_tx_hash,json=depositTxHash,proto3" json:"deposit_tx_hash,omitempty"`
	DepositTxIndex    uint64 `protobuf:"varint,3,opt,name=deposit_tx_index,json=depositTxIndex,proto3" json:"deposit_tx_index,omitempty"`
	DepositBlock      uint64 `protobuf:"varint,4,opt,name=deposit_block,json=depositBlock,proto3" json:"deposit_block,omitempty"`
	DepositToken      string `protobuf:"bytes,5,opt,name=deposit_token,json=depositToken,proto3" json:"deposit_token,omitempty"`
	DepositAmount     string `protobuf:"bytes,6,opt,name=deposit_amount,json=depositAmount,proto3" json:"deposit_amount,omitempty"`
	Depositor         string `protobuf:"bytes,7,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Receiver          string `protobuf:"bytes,8,opt,name=receiver,proto3" json:"receiver,omitempty"`
	WithdrawalChainId string `protobuf:"bytes,9,opt,name=withdrawal_chain_id,json=withdrawalChainId,proto3" json:"withdrawal_chain_id,omitempty"`
	WithdrawalTxHash  string `protobuf:"bytes,10,opt,name=withdrawal_tx_hash,json=withdrawalTxHash,proto3" json:"withdrawal_tx_hash,omitempty"`
	WithdrawalToken   string `protobuf:"bytes,11,opt,name=withdrawal_token,json=withdrawalToken,proto3" json:"withdrawal_token,omitempty"`
	Signature         string `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`
	IsWrapped         bool   `protobuf:"varint,13,opt,name=is_wrapped,json=isWrapped,proto3" json:"is_wrapped,omitempty"`
	WithdrawalAmount  string `protobuf:"bytes,14,opt,name=withdrawal_amount,json=withdrawalAmount,proto3" json:"withdrawal_amount,omitempty"`
}

func (m *transaction) Reset()         { *m = transaction{} }
func (m *transaction) String() string { return proto.CompactTextString(m) }
func (*transaction) ProtoMessage()    {}
func (*transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fcaf7a078652736, []int{0}
}
func (m *transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
func (m *transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *transaction) XXX_Size() int {
	return m.Size()
}
func (m *transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *transaction) GetDepositChainId() string {
	if m != nil {
		return m.DepositChainId
	}
	return ""
}

func (m *transaction) GetDepositTxHash() string {
	if m != nil {
		return m.DepositTxHash
	}
	return ""
}

func (m *transaction) GetDepositTxIndex() uint64 {
	if m != nil {
		return m.DepositTxIndex
	}
	return 0
}

func (m *transaction) GetDepositBlock() uint64 {
	if m != nil {
		return m.DepositBlock
	}
	return 0
}

func (m *transaction) GetDepositToken() string {
	if m != nil {
		return m.DepositToken
	}
	return ""
}

func (m *transaction) GetDepositAmount() string {
	if m != nil {
		return m.DepositAmount
	}
	return ""
}

func (m *transaction) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *transaction) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *transaction) GetWithdrawalChainId() string {
	if m != nil {
		return m.WithdrawalChainId
	}
	return ""
}

func (m *transaction) GetWithdrawalTxHash() string {
	if m != nil {
		return m.WithdrawalTxHash
	}
	return ""
}

func (m *transaction) GetWithdrawalToken() string {
	if m != nil {
		return m.WithdrawalToken
	}
	return ""
}

func (m *transaction) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *transaction) GetIsWrapped() bool {
	if m != nil {
		return m.IsWrapped
	}
	return false
}

func (m *transaction) GetWithdrawalAmount() string {
	if m != nil {
		return m.WithdrawalAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*transaction)(nil), "core.bridge.transaction")
}

func init() { proto.RegisterFile("bridge/models/transaction.proto", fileDescriptor_8fcaf7a078652736) }

var fileDescriptor_8fcaf7a078652736 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x1b, 0x28, 0x43, 0xeb, 0xce, 0xa5, 0x18, 0x16, 0xa6, 0x2a, 0xa1, 0x02, 0x81, 0xca,
	0x65, 0x6a, 0x15, 0x16, 0xac, 0x19, 0x36, 0xcc, 0x92, 0x51, 0x25, 0x24, 0x36, 0x91, 0x93, 0x1c,
	0x12, 0x6b, 0x12, 0x3b, 0xd8, 0x4e, 0x9b, 0xbe, 0x05, 0x6f, 0xc1, 0xab, 0xb0, 0x9c, 0x25, 0x4b,
	0xd4, 0xbe, 0x08, 0x8a, 0x93, 0x34, 0x61, 0x76, 0xee, 0xf7, 0x7f, 0x3d, 0x39, 0xfe, 0x65, 0xf4,
	0xd4, 0x57, 0x3c, 0x8c, 0x80, 0xa6, 0x32, 0x84, 0x44, 0x53, 0xa3, 0x98, 0xd0, 0x2c, 0x30, 0x5c,
	0x8a, 0x45, 0xa6, 0xa4, 0x91, 0x78, 0x14, 0x48, 0x05, 0x8b, 0xca, 0x9a, 0x3c, 0x8e, 0xa4, 0x8c,
	0x12, 0xa0, 0x36, 0xf2, 0xf3, 0xef, 0x94, 0x89, 0x6d, 0xe5, 0x4d, 0x5e, 0x07, 0x52, 0xa7, 0x52,
	0x53, 0x9f, 0x69, 0xa0, 0x3f, 0x72, 0x50, 0x5b, 0xba, 0x5e, 0xfa, 0x60, 0xd8, 0x92, 0x66, 0x2c,
	0xe2, 0x82, 0xb5, 0x33, 0x27, 0xd3, 0x7a, 0x0c, 0xcb, 0x38, 0x65, 0x42, 0x48, 0x63, 0x43, 0x5d,
	0xa7, 0x8f, 0x22, 0x19, 0x49, 0x7b, 0xa4, 0xe5, 0xa9, 0xa2, 0xcf, 0x7e, 0xf5, 0xd1, 0x68, 0xd5,
	0x6e, 0x87, 0xe7, 0x68, 0x1c, 0x42, 0x26, 0x35, 0x37, 0x5e, 0x10, 0x33, 0x2e, 0x3c, 0x1e, 0x12,
	0x67, 0xe6, 0xcc, 0x87, 0x57, 0xa7, 0x35, 0xff, 0x54, 0xe2, 0xcb, 0x10, 0xbf, 0x44, 0x67, 0x8d,
	0x69, 0x0a, 0x2f, 0x66, 0x3a, 0x26, 0x77, 0xac, 0x78, 0x52, 0xe3, 0x55, 0xf1, 0x99, 0xe9, 0xb8,
	0x3b, 0xd1, 0x14, 0x1e, 0x17, 0x21, 0x14, 0xe4, 0xee, 0xcc, 0x99, 0xf7, 0x0f, 0x13, 0x57, 0xc5,
	0x65, 0x49, 0xf1, 0x73, 0xd4, 0xfc, 0xd5, 0xf3, 0x13, 0x19, 0x5c, 0x93, 0xbe, 0xd5, 0x8e, 0x6b,
	0x78, 0x51, 0xb2, 0xae, 0x64, 0xe4, 0x35, 0x08, 0x72, 0xcf, 0x7e, 0xb4, 0x91, 0x56, 0x25, 0xc3,
	0x2f, 0x50, 0x33, 0xdb, 0x63, 0xa9, 0xcc, 0x85, 0x21, 0x47, 0xff, 0xad, 0xf6, 0xd1, 0x42, 0x3c,
	0x45, 0xc3, 0x1a, 0x48, 0x45, 0xee, 0x5b, 0xa3, 0x05, 0x78, 0x82, 0x06, 0x0a, 0x02, 0xe0, 0x6b,
	0x50, 0x64, 0x60, 0xc3, 0xc3, 0x6f, 0xbc, 0x40, 0x0f, 0x37, 0xdc, 0xc4, 0xa1, 0x62, 0x1b, 0x96,
	0xb4, 0x4d, 0x0d, 0xad, 0xf6, 0xa0, 0x8d, 0x9a, 0xb2, 0xde, 0x22, 0xdc, 0xf1, 0x9b, 0xbe, 0x90,
	0xd5, 0xc7, 0x6d, 0x52, 0x57, 0xf6, 0x0a, 0x8d, 0xbb, 0xb6, 0xbd, 0xe6, 0xc8, 0xba, 0x67, 0x1d,
	0xd7, 0xde, 0x74, 0x8a, 0x86, 0x9a, 0x47, 0x82, 0x99, 0x5c, 0x01, 0x39, 0xae, 0xae, 0x70, 0x00,
	0xf8, 0x09, 0x42, 0x5c, 0x7b, 0x1b, 0xc5, 0xb2, 0x0c, 0x42, 0x72, 0x32, 0x73, 0xe6, 0x83, 0xab,
	0x21, 0xd7, 0x5f, 0x2b, 0x80, 0xdf, 0xa0, 0xce, 0xaa, 0x4d, 0x53, 0xa7, 0xb7, 0x97, 0xaa, 0xca,
	0xba, 0xf8, 0xf2, 0x7b, 0xe7, 0x3a, 0x37, 0x3b, 0xd7, 0xf9, 0xbb, 0x73, 0x9d, 0x9f, 0x7b, 0xb7,
	0x77, 0xb3, 0x77, 0x7b, 0x7f, 0xf6, 0x6e, 0xef, 0xdb, 0x87, 0x88, 0x9b, 0x38, 0xf7, 0x17, 0x81,
	0x4c, 0x69, 0xbc, 0x4d, 0xe0, 0xdc, 0x00, 0x4b, 0x69, 0xf5, 0xb6, 0x13, 0xd0, 0xfa, 0xbc, 0x7c,
	0xeb, 0x74, 0xbd, 0x7c, 0x47, 0x8b, 0x9a, 0x52, 0xb3, 0xcd, 0x40, 0xfb, 0x47, 0xf6, 0x0d, 0xbe,
	0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xfb, 0xaa, 0x6b, 0x2e, 0x03, 0x00, 0x00,
}

func (m *transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *transaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *transaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *transaction) Size() (n int) {
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
	return n
}

func sovTransaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *transaction) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: transaction: illegal tag %d (wire type %d)", fieldNum, wire)
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
