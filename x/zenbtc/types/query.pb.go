// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/zenbtc/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CompletionFilter int32

const (
	CompletionFilter_ALL       CompletionFilter = 0
	CompletionFilter_PENDING   CompletionFilter = 1
	CompletionFilter_COMPLETED CompletionFilter = 2
)

var CompletionFilter_name = map[int32]string{
	0: "ALL",
	1: "PENDING",
	2: "COMPLETED",
}

var CompletionFilter_value = map[string]int32{
	"ALL":       0,
	"PENDING":   1,
	"COMPLETED": 2,
}

func (x CompletionFilter) String() string {
	return proto.EnumName(CompletionFilter_name, int32(x))
}

func (CompletionFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8cea4d05ef3f869, []int{0}
}

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cea4d05ef3f869, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cea4d05ef3f869, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryLockTransactionsRequest struct {
}

func (m *QueryLockTransactionsRequest) Reset()         { *m = QueryLockTransactionsRequest{} }
func (m *QueryLockTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLockTransactionsRequest) ProtoMessage()    {}
func (*QueryLockTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cea4d05ef3f869, []int{2}
}
func (m *QueryLockTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLockTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLockTransactionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLockTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLockTransactionsRequest.Merge(m, src)
}
func (m *QueryLockTransactionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLockTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLockTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLockTransactionsRequest proto.InternalMessageInfo

type QueryLockTransactionsResponse struct {
	LockTransactions []string `protobuf:"bytes,1,rep,name=lock_transactions,json=lockTransactions,proto3" json:"lock_transactions,omitempty"`
}

func (m *QueryLockTransactionsResponse) Reset()         { *m = QueryLockTransactionsResponse{} }
func (m *QueryLockTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLockTransactionsResponse) ProtoMessage()    {}
func (*QueryLockTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cea4d05ef3f869, []int{3}
}
func (m *QueryLockTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLockTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLockTransactionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLockTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLockTransactionsResponse.Merge(m, src)
}
func (m *QueryLockTransactionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLockTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLockTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLockTransactionsResponse proto.InternalMessageInfo

func (m *QueryLockTransactionsResponse) GetLockTransactions() []string {
	if m != nil {
		return m.LockTransactions
	}
	return nil
}

type QueryRedemptionsRequest struct {
	StartIndex       uint64           `protobuf:"varint,1,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
	CompletionFilter CompletionFilter `protobuf:"varint,2,opt,name=completion_filter,json=completionFilter,proto3,enum=zrchain.zenbtc.CompletionFilter" json:"completion_filter,omitempty"`
}

func (m *QueryRedemptionsRequest) Reset()         { *m = QueryRedemptionsRequest{} }
func (m *QueryRedemptionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRedemptionsRequest) ProtoMessage()    {}
func (*QueryRedemptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cea4d05ef3f869, []int{4}
}
func (m *QueryRedemptionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRedemptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRedemptionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRedemptionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRedemptionsRequest.Merge(m, src)
}
func (m *QueryRedemptionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRedemptionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRedemptionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRedemptionsRequest proto.InternalMessageInfo

func (m *QueryRedemptionsRequest) GetStartIndex() uint64 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

func (m *QueryRedemptionsRequest) GetCompletionFilter() CompletionFilter {
	if m != nil {
		return m.CompletionFilter
	}
	return CompletionFilter_ALL
}

type QueryRedemptionsResponse struct {
	Redemptions []Redemption `protobuf:"bytes,1,rep,name=redemptions,proto3" json:"redemptions"`
}

func (m *QueryRedemptionsResponse) Reset()         { *m = QueryRedemptionsResponse{} }
func (m *QueryRedemptionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRedemptionsResponse) ProtoMessage()    {}
func (*QueryRedemptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cea4d05ef3f869, []int{5}
}
func (m *QueryRedemptionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRedemptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRedemptionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRedemptionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRedemptionsResponse.Merge(m, src)
}
func (m *QueryRedemptionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRedemptionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRedemptionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRedemptionsResponse proto.InternalMessageInfo

func (m *QueryRedemptionsResponse) GetRedemptions() []Redemption {
	if m != nil {
		return m.Redemptions
	}
	return nil
}

func init() {
	proto.RegisterEnum("zrchain.zenbtc.CompletionFilter", CompletionFilter_name, CompletionFilter_value)
	proto.RegisterType((*QueryParamsRequest)(nil), "zrchain.zenbtc.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "zrchain.zenbtc.QueryParamsResponse")
	proto.RegisterType((*QueryLockTransactionsRequest)(nil), "zrchain.zenbtc.QueryLockTransactionsRequest")
	proto.RegisterType((*QueryLockTransactionsResponse)(nil), "zrchain.zenbtc.QueryLockTransactionsResponse")
	proto.RegisterType((*QueryRedemptionsRequest)(nil), "zrchain.zenbtc.QueryRedemptionsRequest")
	proto.RegisterType((*QueryRedemptionsResponse)(nil), "zrchain.zenbtc.QueryRedemptionsResponse")
}

func init() { proto.RegisterFile("zrchain/zenbtc/query.proto", fileDescriptor_d8cea4d05ef3f869) }

var fileDescriptor_d8cea4d05ef3f869 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0x35, 0x25, 0x55, 0xce, 0x22, 0x72, 0xae, 0xa5, 0x04, 0xa7, 0xb8, 0xc1, 0x0c, 0x44,
	0xa1, 0xc4, 0x6a, 0x18, 0x10, 0x23, 0x69, 0x53, 0x54, 0x29, 0x2d, 0xc1, 0xea, 0xc4, 0x40, 0x74,
	0x76, 0x0f, 0xd7, 0x8a, 0xed, 0x73, 0x7d, 0x17, 0xd4, 0x56, 0x62, 0x61, 0x40, 0x62, 0x43, 0xf0,
	0x27, 0x18, 0xf9, 0x19, 0x1d, 0x2b, 0xb1, 0x30, 0x21, 0x94, 0x20, 0xf1, 0x37, 0x90, 0xcf, 0x0e,
	0x75, 0x9c, 0x14, 0x58, 0xa2, 0xd3, 0xfb, 0xbe, 0xef, 0xbd, 0xef, 0xcb, 0x7b, 0x86, 0xca, 0x59,
	0x68, 0x1d, 0x61, 0xc7, 0xd7, 0xcf, 0x88, 0x6f, 0x72, 0x4b, 0x3f, 0x1e, 0x92, 0xf0, 0xb4, 0x19,
	0x84, 0x94, 0x53, 0x54, 0x4a, 0xb0, 0x66, 0x8c, 0x29, 0x65, 0xec, 0x39, 0x3e, 0xd5, 0xc5, 0x6f,
	0x4c, 0x51, 0x1a, 0x16, 0x65, 0x1e, 0x65, 0xba, 0x89, 0x19, 0x89, 0xb5, 0xfa, 0xeb, 0x4d, 0x93,
	0x70, 0xbc, 0xa9, 0x07, 0xd8, 0x76, 0x7c, 0xcc, 0x1d, 0xea, 0x27, 0xdc, 0x15, 0x9b, 0xda, 0x54,
	0x3c, 0xf5, 0xe8, 0x95, 0x54, 0xd7, 0x6c, 0x4a, 0x6d, 0x97, 0xe8, 0x38, 0x70, 0x74, 0xec, 0xfb,
	0x94, 0x0b, 0x09, 0x4b, 0xd0, 0x6a, 0xc6, 0x5e, 0x80, 0x43, 0xec, 0x4d, 0xc0, 0x5a, 0x06, 0x0c,
	0xc9, 0x21, 0xf1, 0x82, 0x94, 0x5c, 0x5b, 0x81, 0xe8, 0x79, 0x64, 0xaa, 0x27, 0x64, 0x06, 0x39,
	0x1e, 0x12, 0xc6, 0xb5, 0x1e, 0x5c, 0x9e, 0xaa, 0xb2, 0x80, 0xfa, 0x8c, 0xa0, 0xc7, 0xb0, 0x10,
	0xb7, 0xaf, 0x80, 0x1a, 0xa8, 0x4b, 0xad, 0xd5, 0xe6, 0x74, 0xfe, 0x66, 0xcc, 0x6f, 0x17, 0xcf,
	0xbf, 0xaf, 0xe7, 0x3e, 0xff, 0xfa, 0xd2, 0x00, 0x46, 0x22, 0xd0, 0x54, 0xb8, 0x26, 0x3a, 0x76,
	0xa9, 0x35, 0x38, 0x08, 0xb1, 0xcf, 0xb0, 0x25, 0x6c, 0x4c, 0x26, 0x76, 0xe1, 0xed, 0x2b, 0xf0,
	0x64, 0xf6, 0x7d, 0x58, 0x76, 0xa9, 0x35, 0xe8, 0xf3, 0x14, 0x58, 0x01, 0xb5, 0x7c, 0xbd, 0x68,
	0xc8, 0x6e, 0x46, 0xa4, 0xbd, 0x07, 0xf0, 0xa6, 0x68, 0x67, 0x5c, 0x06, 0x4e, 0x26, 0xa1, 0x75,
	0x28, 0x31, 0x8e, 0x43, 0xde, 0x77, 0xfc, 0x43, 0x72, 0x22, 0x92, 0x2c, 0x1a, 0x50, 0x94, 0x76,
	0xa3, 0x0a, 0xda, 0x83, 0x65, 0x8b, 0x7a, 0x81, 0x4b, 0x22, 0x59, 0xff, 0x95, 0xe3, 0x72, 0x12,
	0x56, 0x16, 0x6a, 0xa0, 0x5e, 0x6a, 0xd5, 0xb2, 0x81, 0xb7, 0xfe, 0x10, 0x77, 0x04, 0xcf, 0x90,
	0xad, 0x4c, 0x45, 0x7b, 0x09, 0x2b, 0xb3, 0x56, 0x92, 0x50, 0x6d, 0x28, 0xa5, 0x56, 0x22, 0xe2,
	0x48, 0x2d, 0x25, 0x3b, 0xe4, 0x52, 0xd9, 0x5e, 0x8c, 0xfe, 0x59, 0x23, 0x2d, 0x6a, 0x3c, 0x82,
	0x72, 0xd6, 0x05, 0x5a, 0x82, 0xf9, 0x27, 0xdd, 0xae, 0x9c, 0x43, 0x12, 0x5c, 0xea, 0x75, 0xf6,
	0xb7, 0x77, 0xf7, 0x9f, 0xca, 0x00, 0x5d, 0x87, 0xc5, 0xad, 0x67, 0x7b, 0xbd, 0x6e, 0xe7, 0xa0,
	0xb3, 0x2d, 0x2f, 0xb4, 0xde, 0xe5, 0xe1, 0x35, 0xe1, 0x0c, 0x39, 0xb0, 0x10, 0x6f, 0x0e, 0x69,
	0xd9, 0xd9, 0xb3, 0xc7, 0xa1, 0xdc, 0xfd, 0x2b, 0x27, 0x4e, 0xa6, 0xad, 0xbe, 0xfd, 0xfa, 0xf3,
	0xd3, 0x82, 0x8c, 0x4a, 0xd3, 0x77, 0x89, 0x3e, 0x02, 0x28, 0x67, 0x77, 0x8c, 0x36, 0xe6, 0x76,
	0xbc, 0xe2, 0x54, 0x94, 0x07, 0xff, 0xc9, 0x4e, 0x9c, 0xdc, 0x11, 0x4e, 0xaa, 0xe8, 0xd6, 0xc4,
	0xc9, 0xcc, 0x19, 0xa1, 0x37, 0x50, 0x4a, 0x6d, 0x07, 0xdd, 0x9b, 0x3b, 0x60, 0xf6, 0x94, 0x94,
	0xfa, 0xbf, 0x89, 0x89, 0x89, 0xaa, 0x30, 0x71, 0x03, 0x2d, 0xcf, 0xf9, 0x12, 0xdb, 0x3b, 0xe7,
	0x23, 0x15, 0x5c, 0x8c, 0x54, 0xf0, 0x63, 0xa4, 0x82, 0x0f, 0x63, 0x35, 0x77, 0x31, 0x56, 0x73,
	0xdf, 0xc6, 0x6a, 0xee, 0xc5, 0x86, 0xed, 0xf0, 0xa3, 0xa1, 0xd9, 0xb4, 0xa8, 0x17, 0x09, 0x43,
	0x6a, 0x0d, 0x5c, 0x6c, 0xb2, 0x49, 0x93, 0x93, 0xc9, 0x83, 0x9f, 0x06, 0x84, 0x99, 0x05, 0xf1,
	0x49, 0x3f, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xed, 0x39, 0x74, 0xa2, 0xb2, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of LockTransactions items.
	LockTransactions(ctx context.Context, in *QueryLockTransactionsRequest, opts ...grpc.CallOption) (*QueryLockTransactionsResponse, error)
	// Queries a list of Redemptions items.
	Redemptions(ctx context.Context, in *QueryRedemptionsRequest, opts ...grpc.CallOption) (*QueryRedemptionsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/zrchain.zenbtc.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LockTransactions(ctx context.Context, in *QueryLockTransactionsRequest, opts ...grpc.CallOption) (*QueryLockTransactionsResponse, error) {
	out := new(QueryLockTransactionsResponse)
	err := c.cc.Invoke(ctx, "/zrchain.zenbtc.Query/LockTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Redemptions(ctx context.Context, in *QueryRedemptionsRequest, opts ...grpc.CallOption) (*QueryRedemptionsResponse, error) {
	out := new(QueryRedemptionsResponse)
	err := c.cc.Invoke(ctx, "/zrchain.zenbtc.Query/Redemptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of LockTransactions items.
	LockTransactions(context.Context, *QueryLockTransactionsRequest) (*QueryLockTransactionsResponse, error)
	// Queries a list of Redemptions items.
	Redemptions(context.Context, *QueryRedemptionsRequest) (*QueryRedemptionsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) LockTransactions(ctx context.Context, req *QueryLockTransactionsRequest) (*QueryLockTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockTransactions not implemented")
}
func (*UnimplementedQueryServer) Redemptions(ctx context.Context, req *QueryRedemptionsRequest) (*QueryRedemptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redemptions not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.zenbtc.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LockTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLockTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LockTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.zenbtc.Query/LockTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LockTransactions(ctx, req.(*QueryLockTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Redemptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRedemptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Redemptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zrchain.zenbtc.Query/Redemptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Redemptions(ctx, req.(*QueryRedemptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zrchain.zenbtc.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "LockTransactions",
			Handler:    _Query_LockTransactions_Handler,
		},
		{
			MethodName: "Redemptions",
			Handler:    _Query_Redemptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zrchain/zenbtc/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryLockTransactionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLockTransactionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLockTransactionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryLockTransactionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLockTransactionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLockTransactionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LockTransactions) > 0 {
		for iNdEx := len(m.LockTransactions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.LockTransactions[iNdEx])
			copy(dAtA[i:], m.LockTransactions[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.LockTransactions[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryRedemptionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRedemptionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRedemptionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CompletionFilter != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CompletionFilter))
		i--
		dAtA[i] = 0x10
	}
	if m.StartIndex != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.StartIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryRedemptionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRedemptionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRedemptionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Redemptions) > 0 {
		for iNdEx := len(m.Redemptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Redemptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryLockTransactionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryLockTransactionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LockTransactions) > 0 {
		for _, s := range m.LockTransactions {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryRedemptionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartIndex != 0 {
		n += 1 + sovQuery(uint64(m.StartIndex))
	}
	if m.CompletionFilter != 0 {
		n += 1 + sovQuery(uint64(m.CompletionFilter))
	}
	return n
}

func (m *QueryRedemptionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Redemptions) > 0 {
		for _, e := range m.Redemptions {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryLockTransactionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryLockTransactionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLockTransactionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryLockTransactionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryLockTransactionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLockTransactionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockTransactions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockTransactions = append(m.LockTransactions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRedemptionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRedemptionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRedemptionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartIndex", wireType)
			}
			m.StartIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionFilter", wireType)
			}
			m.CompletionFilter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompletionFilter |= CompletionFilter(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRedemptionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRedemptionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRedemptionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redemptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Redemptions = append(m.Redemptions, Redemption{})
			if err := m.Redemptions[len(m.Redemptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
