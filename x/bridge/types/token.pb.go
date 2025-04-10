// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bridge/models/token.proto

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

type TokenInfo struct {
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Decimals  uint64 `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
	ChainId   string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	TokenId   uint64 `protobuf:"varint,4,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	IsWrapped bool   `protobuf:"varint,5,opt,name=is_wrapped,json=isWrapped,proto3" json:"is_wrapped,omitempty"`
}

func (m *TokenInfo) Reset()         { *m = TokenInfo{} }
func (m *TokenInfo) String() string { return proto.CompactTextString(m) }
func (*TokenInfo) ProtoMessage()    {}
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_68574e97ac2633fb, []int{0}
}
func (m *TokenInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenInfo.Merge(m, src)
}
func (m *TokenInfo) XXX_Size() int {
	return m.Size()
}
func (m *TokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TokenInfo proto.InternalMessageInfo

func (m *TokenInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TokenInfo) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *TokenInfo) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *TokenInfo) GetTokenId() uint64 {
	if m != nil {
		return m.TokenId
	}
	return 0
}

func (m *TokenInfo) GetIsWrapped() bool {
	if m != nil {
		return m.IsWrapped
	}
	return false
}

type TokenMetadata struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Uri    string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (m *TokenMetadata) Reset()         { *m = TokenMetadata{} }
func (m *TokenMetadata) String() string { return proto.CompactTextString(m) }
func (*TokenMetadata) ProtoMessage()    {}
func (*TokenMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_68574e97ac2633fb, []int{1}
}
func (m *TokenMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenMetadata.Merge(m, src)
}
func (m *TokenMetadata) XXX_Size() int {
	return m.Size()
}
func (m *TokenMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TokenMetadata proto.InternalMessageInfo

func (m *TokenMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TokenMetadata) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TokenMetadata) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type Token struct {
	Id       uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Metadata TokenMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
	// info is the token information on different chains
	Info           []TokenInfo `protobuf:"bytes,3,rep,name=info,proto3" json:"info"`
	CommissionRate float32     `protobuf:"fixed32,4,opt,name=commission_rate,json=commissionRate,proto3" json:"commission_rate,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_68574e97ac2633fb, []int{2}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Token) GetMetadata() TokenMetadata {
	if m != nil {
		return m.Metadata
	}
	return TokenMetadata{}
}

func (m *Token) GetInfo() []TokenInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Token) GetCommissionRate() float32 {
	if m != nil {
		return m.CommissionRate
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenInfo)(nil), "core.bridge.TokenInfo")
	proto.RegisterType((*TokenMetadata)(nil), "core.bridge.TokenMetadata")
	proto.RegisterType((*Token)(nil), "core.bridge.Token")
}

func init() { proto.RegisterFile("bridge/models/token.proto", fileDescriptor_68574e97ac2633fb) }

var fileDescriptor_68574e97ac2633fb = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xc1, 0x8e, 0xd3, 0x30,
	0x14, 0x6c, 0xda, 0xec, 0x6e, 0xeb, 0x8a, 0x05, 0x59, 0x68, 0x95, 0x46, 0x10, 0xaa, 0x5e, 0xa8,
	0x90, 0x36, 0xa6, 0xe5, 0xc0, 0x85, 0xd3, 0xde, 0x7a, 0xd8, 0x03, 0x11, 0x12, 0x12, 0x97, 0xca,
	0x89, 0x5f, 0x53, 0x8b, 0xd8, 0x2f, 0xc4, 0xee, 0x42, 0xfe, 0x82, 0x13, 0xdf, 0xc1, 0x67, 0xec,
	0x71, 0x8f, 0x9c, 0x10, 0x6a, 0x7f, 0x04, 0xc5, 0x09, 0x05, 0xb4, 0xb7, 0x37, 0x33, 0x9e, 0xf1,
	0xd8, 0x8f, 0x4c, 0xd2, 0x4a, 0x8a, 0x1c, 0x98, 0x42, 0x01, 0x85, 0x61, 0x16, 0x3f, 0x82, 0x8e,
	0xcb, 0x0a, 0x2d, 0xd2, 0x71, 0x86, 0x15, 0xc4, 0xad, 0x1e, 0x4e, 0x72, 0xc4, 0xbc, 0x00, 0xe6,
	0xa4, 0x74, 0xb7, 0x61, 0x5c, 0xd7, 0xed, 0xb9, 0xf0, 0x45, 0x86, 0x46, 0xa1, 0x61, 0x29, 0x37,
	0xc0, 0x3e, 0xed, 0xa0, 0xaa, 0xd9, 0xcd, 0x22, 0x05, 0xcb, 0x17, 0xac, 0xe4, 0xb9, 0xd4, 0xdc,
	0x4a, 0xec, 0x32, 0xc3, 0x27, 0x5d, 0x0c, 0x2f, 0x25, 0xe3, 0x5a, 0xa3, 0x75, 0xa2, 0xe9, 0xd4,
	0xc7, 0x39, 0xe6, 0xe8, 0x46, 0xd6, 0x4c, 0x2d, 0x3b, 0xfb, 0xe6, 0x91, 0xd1, 0xbb, 0xa6, 0xd7,
	0x4a, 0x6f, 0x90, 0x06, 0xe4, 0x8c, 0x0b, 0x51, 0x81, 0x31, 0x81, 0x37, 0xf5, 0xe6, 0xa3, 0xe4,
	0x0f, 0xa4, 0x21, 0x19, 0x0a, 0xc8, 0xa4, 0xe2, 0x85, 0x09, 0xfa, 0x53, 0x6f, 0xee, 0x27, 0x47,
	0x4c, 0x27, 0x64, 0x98, 0x6d, 0xb9, 0xd4, 0x6b, 0x29, 0x82, 0x41, 0x6b, 0x73, 0x78, 0x25, 0x1a,
	0xc9, 0xbd, 0xba, 0x91, 0x7c, 0x67, 0x3b, 0x73, 0x78, 0x25, 0xe8, 0x53, 0x42, 0xa4, 0x59, 0x7f,
	0xae, 0x78, 0x59, 0x82, 0x08, 0x4e, 0xa6, 0xde, 0x7c, 0x98, 0x8c, 0xa4, 0x79, 0xdf, 0x12, 0xb3,
	0x6b, 0xf2, 0xc0, 0xf5, 0xba, 0x06, 0xcb, 0x05, 0xb7, 0x9c, 0x52, 0xe2, 0x6b, 0xae, 0xa0, 0x2b,
	0xe6, 0x66, 0x7a, 0x41, 0x4e, 0x4d, 0xad, 0x52, 0x2c, 0x5c, 0xa7, 0x51, 0xd2, 0x21, 0xfa, 0x88,
	0x0c, 0x76, 0x95, 0xec, 0xca, 0x34, 0xe3, 0xec, 0xbb, 0x47, 0x4e, 0x5c, 0x1e, 0x3d, 0x27, 0x7d,
	0x29, 0x5c, 0x8a, 0x9f, 0xf4, 0xa5, 0xa0, 0x6f, 0xc8, 0x50, 0x75, 0x77, 0xb8, 0x94, 0xf1, 0x32,
	0x8c, 0xff, 0x59, 0x4e, 0xfc, 0x5f, 0x8b, 0x2b, 0xff, 0xf6, 0xe7, 0xb3, 0x5e, 0x72, 0x74, 0xd0,
	0x97, 0xc4, 0x97, 0x7a, 0x83, 0xc1, 0x60, 0x3a, 0x98, 0x8f, 0x97, 0x17, 0xf7, 0x9d, 0xcd, 0xbf,
	0x76, 0x2e, 0x77, 0x92, 0x3e, 0x27, 0x0f, 0x33, 0x54, 0x4a, 0x1a, 0x23, 0x51, 0xaf, 0x2b, 0x6e,
	0xc1, 0xfd, 0x4c, 0x3f, 0x39, 0xff, 0x4b, 0x27, 0xdc, 0xc2, 0xd5, 0xdb, 0xdb, 0x7d, 0xe4, 0xdd,
	0xed, 0x23, 0xef, 0xd7, 0x3e, 0xf2, 0xbe, 0x1e, 0xa2, 0xde, 0xdd, 0x21, 0xea, 0xfd, 0x38, 0x44,
	0xbd, 0x0f, 0xaf, 0x73, 0x69, 0xb7, 0xbb, 0x34, 0xce, 0x50, 0xb1, 0x6d, 0x5d, 0xc0, 0xa5, 0x05,
	0xae, 0x58, 0x7b, 0x6b, 0x01, 0xc6, 0x5c, 0x36, 0x2d, 0xd8, 0xcd, 0x62, 0xc9, 0xbe, 0x74, 0x2c,
	0xb3, 0x75, 0x09, 0x26, 0x3d, 0x75, 0x4b, 0x7f, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x74,
	0x28, 0x64, 0x99, 0x02, 0x00, 0x00,
}

func (m *TokenInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsWrapped {
		i--
		if m.IsWrapped {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.TokenId != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintToken(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Decimals != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CommissionRate != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.CommissionRate))))
		i--
		dAtA[i] = 0x25
	}
	if len(m.Info) > 0 {
		for iNdEx := len(m.Info) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Info[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintToken(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovToken(uint64(m.Decimals))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.TokenId != 0 {
		n += 1 + sovToken(uint64(m.TokenId))
	}
	if m.IsWrapped {
		n += 2
	}
	return n
}

func (m *TokenMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovToken(uint64(m.Id))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovToken(uint64(l))
	if len(m.Info) > 0 {
		for _, e := range m.Info {
			l = e.Size()
			n += 1 + l + sovToken(uint64(l))
		}
	}
	if m.CommissionRate != 0 {
		n += 5
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: TokenInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsWrapped", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func (m *TokenMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: TokenMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = append(m.Info, TokenInfo{})
			if err := m.Info[len(m.Info)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.CommissionRate = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
