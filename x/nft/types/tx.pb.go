// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nft/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

func init() { proto.RegisterFile("nft/tx.proto", fileDescriptor_09d30374d974e015) }

var fileDescriptor_09d30374d974e015 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0x31, 0x4e, 0x85, 0x40,
	0x10, 0x86, 0x21, 0x46, 0x63, 0x88, 0x95, 0xb1, 0xa2, 0x58, 0x4f, 0x00, 0x23, 0x7a, 0x03, 0x4b,
	0x13, 0x2f, 0x60, 0xb7, 0xbb, 0x0e, 0xcb, 0x26, 0x2c, 0x43, 0x98, 0xc1, 0xc0, 0x2d, 0x3c, 0xd6,
	0x2b, 0x29, 0x5f, 0xf9, 0x02, 0x17, 0x79, 0x61, 0x79, 0xdd, 0x37, 0x93, 0x2f, 0x7f, 0xbe, 0xec,
	0xa9, 0xab, 0x05, 0x64, 0x2a, 0xfb, 0x81, 0x84, 0x9e, 0x1f, 0x2d, 0x0d, 0x58, 0x76, 0xb5, 0xe4,
	0xca, 0x12, 0x07, 0x62, 0x30, 0x9a, 0x11, 0xfe, 0x2a, 0x83, 0xa2, 0x2b, 0xb0, 0xe4, 0xbb, 0xc3,
	0xcc, 0x5f, 0x1d, 0x91, 0x6b, 0x11, 0xe2, 0x65, 0xc6, 0x1a, 0xc4, 0x07, 0x64, 0xd1, 0xa1, 0xbf,
	0x09, 0x2f, 0x8e, 0x1c, 0x45, 0x84, 0x9d, 0x8e, 0xef, 0xfb, 0x7d, 0x76, 0xf7, 0xcd, 0xee, 0xf3,
	0xeb, 0xb4, 0xaa, 0x74, 0x59, 0x55, 0x7a, 0x59, 0x55, 0xfa, 0xbf, 0xa9, 0x64, 0xd9, 0x54, 0x72,
	0xde, 0x54, 0xf2, 0xf3, 0xe6, 0xbc, 0x34, 0xa3, 0x29, 0x2d, 0x05, 0x68, 0xe6, 0x16, 0x0b, 0x41,
	0x1d, 0xc0, 0x0c, 0xfe, 0xd7, 0x61, 0x8b, 0xcc, 0xc5, 0x5e, 0x08, 0x13, 0xc4, 0xec, 0xb9, 0x47,
	0x36, 0x0f, 0x71, 0xf9, 0xe3, 0x1a, 0x00, 0x00, 0xff, 0xff, 0x00, 0xdb, 0xe0, 0xbb, 0xca, 0x00,
	0x00, 0x00,
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
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.nft.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "nft/tx.proto",
}
