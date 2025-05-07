package v6

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

type tranasction6 struct {
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
	CommissionAmount  string `protobuf:"bytes,15,opt,name=commission_amount,json=commissionAmount,proto3" json:"commission_amount,omitempty"`
	TxData            string `protobuf:"bytes,16,opt,name=tx_data,json=txData,proto3" json:"tx_data,omitempty"`
}

func (m *tranasction6) Reset()         { *m = tranasction6{} }
func (m *tranasction6) String() string { return proto.CompactTextString(m) }
func (*tranasction6) ProtoMessage()    {}
func (*tranasction6) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fcaf7a078652736, []int{0}
}
func (m *tranasction6) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *tranasction6) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
func (m *tranasction6) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *tranasction6) XXX_Size() int {
	return m.Size()
}
func (m *tranasction6) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *tranasction6) GetDepositChainId() string {
	if m != nil {
		return m.DepositChainId
	}
	return ""
}

func (m *tranasction6) GetDepositTxHash() string {
	if m != nil {
		return m.DepositTxHash
	}
	return ""
}

func (m *tranasction6) GetDepositTxIndex() uint64 {
	if m != nil {
		return m.DepositTxIndex
	}
	return 0
}

func (m *tranasction6) GetDepositBlock() uint64 {
	if m != nil {
		return m.DepositBlock
	}
	return 0
}

func (m *tranasction6) GetDepositToken() string {
	if m != nil {
		return m.DepositToken
	}
	return ""
}

func (m *tranasction6) GetDepositAmount() string {
	if m != nil {
		return m.DepositAmount
	}
	return ""
}

func (m *tranasction6) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *tranasction6) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *tranasction6) GetWithdrawalChainId() string {
	if m != nil {
		return m.WithdrawalChainId
	}
	return ""
}

func (m *tranasction6) GetWithdrawalTxHash() string {
	if m != nil {
		return m.WithdrawalTxHash
	}
	return ""
}

func (m *tranasction6) GetWithdrawalToken() string {
	if m != nil {
		return m.WithdrawalToken
	}
	return ""
}

func (m *tranasction6) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *tranasction6) GetIsWrapped() bool {
	if m != nil {
		return m.IsWrapped
	}
	return false
}

func (m *tranasction6) GetWithdrawalAmount() string {
	if m != nil {
		return m.WithdrawalAmount
	}
	return ""
}

func (m *tranasction6) GetCommissionAmount() string {
	if m != nil {
		return m.CommissionAmount
	}
	return ""
}

func (m *tranasction6) GetTxData() string {
	if m != nil {
		return m.TxData
	}
	return ""
}

type txSubmissions struct {
	TxHash           string   `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Submitters       []string `protobuf:"bytes,2,rep,name=submitters,proto3" json:"submitters,omitempty"`
	CommissionAmount string   `protobuf:"bytes,15,opt,name=commission_amount,json=commissionAmount,proto3" json:"commission_amount,omitempty"`
}

func (m *txSubmissions) Reset()         { *m = txSubmissions{} }
func (m *txSubmissions) String() string { return proto.CompactTextString(m) }
func (*txSubmissions) ProtoMessage()    {}
func (*txSubmissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fcaf7a078652736, []int{1}
}
func (m *txSubmissions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *txSubmissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionSubmissions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *txSubmissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionSubmissions.Merge(m, src)
}
func (m *txSubmissions) XXX_Size() int {
	return m.Size()
}
func (m *txSubmissions) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionSubmissions.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionSubmissions proto.InternalMessageInfo

func (m *txSubmissions) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *txSubmissions) GetSubmitters() []string {
	if m != nil {
		return m.Submitters
	}
	return nil
}

func (m *txSubmissions) GetCommissionAmount() string {
	if m != nil {
		return m.CommissionAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*tranasction6)(nil), "core.bridge.tranasction6")
	proto.RegisterType((*txSubmissions)(nil), "core.bridge.txSubmissions")
}

func init() { proto.RegisterFile("bridge/models/tranasction6.proto", fileDescriptor_8fcaf7a078652736) }

var fileDescriptor_8fcaf7a078652736 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0xc7, 0x97, 0x6d, 0x6f, 0xd7, 0xb8, 0xdb, 0xda, 0xe5, 0x45, 0x60, 0xaa, 0x12, 0xaa, 0x21,
	0x50, 0x61, 0xac, 0x56, 0xe1, 0xc0, 0x99, 0xc1, 0x81, 0x1d, 0x29, 0x95, 0x90, 0xb8, 0x44, 0x4e,
	0x62, 0x12, 0x6b, 0x89, 0x1d, 0x6c, 0xa7, 0x4d, 0x2f, 0x7c, 0x06, 0x3e, 0x16, 0xc7, 0x1d, 0x39,
	0xa2, 0xf6, 0xcc, 0x77, 0x40, 0x71, 0x9c, 0x26, 0x70, 0xe3, 0xe6, 0xfe, 0x9e, 0xdf, 0xf3, 0xf4,
	0xf1, 0x5f, 0x31, 0x78, 0xe8, 0x0b, 0x1a, 0x46, 0x04, 0xa5, 0x3c, 0x24, 0x89, 0x44, 0x4a, 0x60,
	0x26, 0x71, 0xa0, 0x28, 0x67, 0xd3, 0x4c, 0x70, 0xc5, 0x9d, 0x5e, 0xc0, 0x05, 0x99, 0x56, 0xd6,
	0xf0, 0x7e, 0xc4, 0x79, 0x94, 0x10, 0xa4, 0x4b, 0x7e, 0xfe, 0x19, 0x61, 0xb6, 0xae, 0xbc, 0xe1,
	0xb3, 0x80, 0xcb, 0x94, 0x4b, 0xe4, 0x63, 0x49, 0xd0, 0x97, 0x9c, 0x88, 0x35, 0x5a, 0xce, 0x7c,
	0xa2, 0xf0, 0x0c, 0x65, 0x38, 0xa2, 0x0c, 0x37, 0x33, 0x87, 0x23, 0x33, 0x06, 0x67, 0x14, 0x61,
	0xc6, 0xb8, 0xd2, 0x45, 0x69, 0xaa, 0x77, 0x22, 0x1e, 0x71, 0x7d, 0x44, 0xe5, 0xa9, 0xa2, 0xe7,
	0xbf, 0x0e, 0x41, 0x6f, 0xd1, 0x6c, 0xe7, 0x4c, 0xc0, 0x20, 0x24, 0x19, 0x97, 0x54, 0x79, 0x41,
	0x8c, 0x29, 0xf3, 0x68, 0x08, 0xad, 0xb1, 0x35, 0xb1, 0xe7, 0xa7, 0x86, 0xbf, 0x29, 0xf1, 0x75,
	0xe8, 0x3c, 0x01, 0xfd, 0xda, 0x54, 0x85, 0x17, 0x63, 0x19, 0xc3, 0x7d, 0x2d, 0x9e, 0x18, 0xbc,
	0x28, 0xde, 0x61, 0x19, 0xb7, 0x27, 0xaa, 0xc2, 0xa3, 0x2c, 0x24, 0x05, 0x3c, 0x18, 0x5b, 0x93,
	0xc3, 0xdd, 0xc4, 0x45, 0x71, 0x5d, 0x52, 0xe7, 0x11, 0xa8, 0x5b, 0x3d, 0x3f, 0xe1, 0xc1, 0x0d,
	0x3c, 0xd4, 0xda, 0xb1, 0x81, 0x57, 0x25, 0x6b, 0x4b, 0x8a, 0xdf, 0x10, 0x06, 0xff, 0xd3, 0x7f,
	0x5a, 0x4b, 0x8b, 0x92, 0x39, 0x8f, 0x41, 0x3d, 0xdb, 0xc3, 0x29, 0xcf, 0x99, 0x82, 0x9d, 0x3f,
	0x56, 0x7b, 0xad, 0xa1, 0x33, 0x02, 0xb6, 0x01, 0x5c, 0xc0, 0x23, 0x6d, 0x34, 0xc0, 0x19, 0x82,
	0xae, 0x20, 0x01, 0xa1, 0x4b, 0x22, 0x60, 0x57, 0x17, 0x77, 0xbf, 0x9d, 0x29, 0xf8, 0x7f, 0x45,
	0x55, 0x1c, 0x0a, 0xbc, 0xc2, 0x49, 0x93, 0x94, 0xad, 0xb5, 0xb3, 0xa6, 0x54, 0x87, 0xf5, 0x1c,
	0x38, 0x2d, 0xbf, 0xce, 0x0b, 0x68, 0x7d, 0xd0, 0x54, 0x4c, 0x64, 0x4f, 0xc1, 0xa0, 0x6d, 0xeb,
	0x6b, 0xf6, 0xb4, 0xdb, 0x6f, 0xb9, 0xfa, 0xa6, 0x23, 0x60, 0x4b, 0x1a, 0x31, 0xac, 0x72, 0x41,
	0xe0, 0x71, 0x75, 0x85, 0x1d, 0x70, 0x1e, 0x00, 0x40, 0xa5, 0xb7, 0x12, 0x38, 0xcb, 0x48, 0x08,
	0x4f, 0xc6, 0xd6, 0xa4, 0x3b, 0xb7, 0xa9, 0xfc, 0x58, 0x01, 0xe7, 0x02, 0xb4, 0x56, 0xad, 0x93,
	0x3a, 0xfd, 0x7b, 0x29, 0x13, 0xd6, 0x05, 0x38, 0x0b, 0x78, 0x9a, 0x52, 0x29, 0x29, 0x67, 0xb5,
	0xdc, 0xaf, 0xe4, 0xa6, 0x60, 0xe4, 0x7b, 0xe0, 0x48, 0x15, 0x5e, 0x88, 0x15, 0x86, 0x03, 0xad,
	0x74, 0x54, 0xf1, 0x16, 0x2b, 0x7c, 0xfe, 0x15, 0xdc, 0x6d, 0x7d, 0x6e, 0x1f, 0x72, 0xdf, 0xf4,
	0x49, 0xd3, 0xa2, 0x73, 0xb1, 0xea, 0x16, 0x9d, 0x86, 0x0b, 0x80, 0x2c, 0x3d, 0xa5, 0x88, 0x90,
	0x70, 0x7f, 0x7c, 0x30, 0xb1, 0xe7, 0x2d, 0xf2, 0x4f, 0x8b, 0x5d, 0xbd, 0xff, 0xbe, 0x71, 0xad,
	0xdb, 0x8d, 0x6b, 0xfd, 0xdc, 0xb8, 0xd6, 0xb7, 0xad, 0xbb, 0x77, 0xbb, 0x75, 0xf7, 0x7e, 0x6c,
	0xdd, 0xbd, 0x4f, 0xaf, 0x22, 0xaa, 0xe2, 0xdc, 0x9f, 0x06, 0x3c, 0x45, 0xf1, 0x3a, 0x21, 0x97,
	0x8a, 0xe0, 0x14, 0x55, 0x2f, 0x34, 0x21, 0x52, 0x5e, 0x96, 0x2f, 0x16, 0x2d, 0x67, 0x2f, 0x50,
	0x61, 0x28, 0x52, 0xeb, 0x8c, 0x48, 0xbf, 0xa3, 0x5f, 0xd2, 0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x84, 0xd3, 0x7d, 0x1a, 0xf4, 0x03, 0x00, 0x00,
}

func (m *tranasction6) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *tranasction6) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *tranasction6) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxData) > 0 {
		i -= len(m.TxData)
		copy(dAtA[i:], m.TxData)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.TxData)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.CommissionAmount) > 0 {
		i -= len(m.CommissionAmount)
		copy(dAtA[i:], m.CommissionAmount)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.CommissionAmount)))
		i--
		dAtA[i] = 0x7a
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

func (m *txSubmissions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *txSubmissions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *txSubmissions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CommissionAmount) > 0 {
		i -= len(m.CommissionAmount)
		copy(dAtA[i:], m.CommissionAmount)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.CommissionAmount)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.Submitters) > 0 {
		for iNdEx := len(m.Submitters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Submitters[iNdEx])
			copy(dAtA[i:], m.Submitters[iNdEx])
			i = encodeVarintTransaction(dAtA, i, uint64(len(m.Submitters[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.TxHash)))
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
func (m *tranasction6) Size() (n int) {
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
	l = len(m.CommissionAmount)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.TxData)
	if l > 0 {
		n += 2 + l + sovTransaction(uint64(l))
	}
	return n
}

func (m *txSubmissions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if len(m.Submitters) > 0 {
		for _, s := range m.Submitters {
			l = len(s)
			n += 1 + l + sovTransaction(uint64(l))
		}
	}
	l = len(m.CommissionAmount)
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
func (m *tranasction6) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: tranasction6: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: tranasction6: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionAmount", wireType)
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
			m.CommissionAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxData", wireType)
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
			m.TxData = string(dAtA[iNdEx:postIndex])
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
func (m *txSubmissions) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: txSubmissions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: txSubmissions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
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
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submitters", wireType)
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
			m.Submitters = append(m.Submitters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionAmount", wireType)
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
			m.CommissionAmount = string(dAtA[iNdEx:postIndex])
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
