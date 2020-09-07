// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finance_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransferInRequest struct {
	PersonId             int64    `protobuf:"zigzag64,1,opt,name=personId,proto3" json:"personId,omitempty"`
	TransferWith         int32    `protobuf:"zigzag32,2,opt,name=transferWith,proto3" json:"transferWith,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferInRequest) Reset()         { *m = TransferInRequest{} }
func (m *TransferInRequest) String() string { return proto.CompactTextString(m) }
func (*TransferInRequest) ProtoMessage()    {}
func (*TransferInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_736f4f126335f71f, []int{0}
}
func (m *TransferInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferInRequest.Unmarshal(m, b)
}
func (m *TransferInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferInRequest.Marshal(b, m, deterministic)
}
func (dst *TransferInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferInRequest.Merge(dst, src)
}
func (m *TransferInRequest) XXX_Size() int {
	return xxx_messageInfo_TransferInRequest.Size(m)
}
func (m *TransferInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferInRequest proto.InternalMessageInfo

func (m *TransferInRequest) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func (m *TransferInRequest) GetTransferWith() int32 {
	if m != nil {
		return m.TransferWith
	}
	return 0
}

func (m *TransferInRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*TransferInRequest)(nil), "TransferInRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FinanceServiceClient is the client API for FinanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinanceServiceClient interface {
	// 转入(业务放在service,是为person_finance解耦)
	RiseTransferIn(ctx context.Context, in *TransferInRequest, opts ...grpc.CallOption) (*Result, error)
}

type financeServiceClient struct {
	cc *grpc.ClientConn
}

func NewFinanceServiceClient(cc *grpc.ClientConn) FinanceServiceClient {
	return &financeServiceClient{cc}
}

func (c *financeServiceClient) RiseTransferIn(ctx context.Context, in *TransferInRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/RiseTransferIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceServiceServer is the server API for FinanceService service.
type FinanceServiceServer interface {
	// 转入(业务放在service,是为person_finance解耦)
	RiseTransferIn(context.Context, *TransferInRequest) (*Result, error)
}

func RegisterFinanceServiceServer(s *grpc.Server, srv FinanceServiceServer) {
	s.RegisterService(&_FinanceService_serviceDesc, srv)
}

func _FinanceService_RiseTransferIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).RiseTransferIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/RiseTransferIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).RiseTransferIn(ctx, req.(*TransferInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FinanceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FinanceService",
	HandlerType: (*FinanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RiseTransferIn",
			Handler:    _FinanceService_RiseTransferIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance_service.proto",
}

func init() {
	proto.RegisterFile("finance_service.proto", fileDescriptor_finance_service_736f4f126335f71f)
}

var fileDescriptor_finance_service_736f4f126335f71f = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xcb, 0xcc, 0x4b,
	0xcc, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x97, 0xe2, 0x2e, 0x29, 0xa9, 0x2c, 0x80, 0x72, 0x94, 0xb2, 0xb9, 0x04, 0x43, 0x8a, 0x12,
	0xf3, 0x8a, 0xd3, 0x52, 0x8b, 0x3c, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4,
	0xb8, 0x38, 0x0a, 0x52, 0x8b, 0x8a, 0xf3, 0xf3, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x84,
	0x82, 0xe0, 0x7c, 0x21, 0x25, 0x2e, 0x9e, 0x12, 0xa8, 0x86, 0xf0, 0xcc, 0x92, 0x0c, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xc1, 0x20, 0x14, 0x31, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xdc, 0xfc, 0xd2, 0xbc,
	0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xc6, 0x20, 0x28, 0xcf, 0xc8, 0x91, 0x8b, 0xcf, 0x0d, 0xe2,
	0xa4, 0x60, 0x88, 0x8b, 0x84, 0xf4, 0xb9, 0xf8, 0x82, 0x32, 0x8b, 0x53, 0x11, 0x4e, 0x10, 0x12,
	0xd2, 0xc3, 0x70, 0x8f, 0x14, 0xbb, 0x5e, 0x50, 0x6a, 0x71, 0x69, 0x4e, 0x89, 0x12, 0x83, 0x13,
	0x67, 0x14, 0xbb, 0x9e, 0x35, 0xd8, 0xe9, 0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xc6, 0xb0, 0x5f, 0x1f, 0xe7, 0x00, 0x00, 0x00,
}
