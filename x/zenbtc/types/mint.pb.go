// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/zenbtc/mint.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// WalletType specifies the Layer 1 blockchain that this wallet will be used
// for.
type WalletType int32

const (
	// The wallet type is missing (all wallets will be derived)
	WalletType_WALLET_TYPE_UNSPECIFIED WalletType = 0
	// The wallet type for native cosmos accounts
	WalletType_WALLET_TYPE_NATIVE WalletType = 1
	// The wallet type for mainnet ETH and its ERC-20 tokens
	WalletType_WALLET_TYPE_EVM WalletType = 2
	// The Wallet type for Testnet BTC
	WalletType_WALLET_TYPE_BTC_TESTNET WalletType = 3
	// The Wallet type for Mainnet BTC
	WalletType_WALLET_TYPE_BTC_MAINNET WalletType = 4
	// The Wallet type for RegNet - Local Test Network
	WalletType_WALLET_TYPE_BTC_REGNET WalletType = 5
	// Wallet type for Solana
	WalletType_WALLET_TYPE_SOLANA WalletType = 6
)

var WalletType_name = map[int32]string{
	0: "WALLET_TYPE_UNSPECIFIED",
	1: "WALLET_TYPE_NATIVE",
	2: "WALLET_TYPE_EVM",
	3: "WALLET_TYPE_BTC_TESTNET",
	4: "WALLET_TYPE_BTC_MAINNET",
	5: "WALLET_TYPE_BTC_REGNET",
	6: "WALLET_TYPE_SOLANA",
}

var WalletType_value = map[string]int32{
	"WALLET_TYPE_UNSPECIFIED": 0,
	"WALLET_TYPE_NATIVE":      1,
	"WALLET_TYPE_EVM":         2,
	"WALLET_TYPE_BTC_TESTNET": 3,
	"WALLET_TYPE_BTC_MAINNET": 4,
	"WALLET_TYPE_BTC_REGNET":  5,
	"WALLET_TYPE_SOLANA":      6,
}

func (x WalletType) String() string {
	return proto.EnumName(WalletType_name, int32(x))
}

func (WalletType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8855381be8e1a27d, []int{0}
}

type NonceData struct {
	Nonce   uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Counter uint64 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	Skip    bool   `protobuf:"varint,3,opt,name=skip,proto3" json:"skip,omitempty"`
}

func (m *NonceData) Reset()         { *m = NonceData{} }
func (m *NonceData) String() string { return proto.CompactTextString(m) }
func (*NonceData) ProtoMessage()    {}
func (*NonceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8855381be8e1a27d, []int{0}
}
func (m *NonceData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NonceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NonceData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NonceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonceData.Merge(m, src)
}
func (m *NonceData) XXX_Size() int {
	return m.Size()
}
func (m *NonceData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonceData.DiscardUnknown(m)
}

var xxx_messageInfo_NonceData proto.InternalMessageInfo

func (m *NonceData) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *NonceData) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *NonceData) GetSkip() bool {
	if m != nil {
		return m.Skip
	}
	return false
}

type RequestedBitcoinHeaders struct {
	Heights []int64 `protobuf:"varint,1,rep,packed,name=heights,proto3" json:"heights,omitempty"`
}

func (m *RequestedBitcoinHeaders) Reset()         { *m = RequestedBitcoinHeaders{} }
func (m *RequestedBitcoinHeaders) String() string { return proto.CompactTextString(m) }
func (*RequestedBitcoinHeaders) ProtoMessage()    {}
func (*RequestedBitcoinHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_8855381be8e1a27d, []int{1}
}
func (m *RequestedBitcoinHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestedBitcoinHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestedBitcoinHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestedBitcoinHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestedBitcoinHeaders.Merge(m, src)
}
func (m *RequestedBitcoinHeaders) XXX_Size() int {
	return m.Size()
}
func (m *RequestedBitcoinHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestedBitcoinHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_RequestedBitcoinHeaders proto.InternalMessageInfo

func (m *RequestedBitcoinHeaders) GetHeights() []int64 {
	if m != nil {
		return m.Heights
	}
	return nil
}

type LockTransaction struct {
	RawTx         string `protobuf:"bytes,1,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	Vout          uint64 `protobuf:"varint,2,opt,name=vout,proto3" json:"vout,omitempty"`
	Sender        string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	MintRecipient string `protobuf:"bytes,4,opt,name=mint_recipient,json=mintRecipient,proto3" json:"mint_recipient,omitempty"`
	Amount        uint64 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	BlockHeight   int64  `protobuf:"varint,6,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *LockTransaction) Reset()         { *m = LockTransaction{} }
func (m *LockTransaction) String() string { return proto.CompactTextString(m) }
func (*LockTransaction) ProtoMessage()    {}
func (*LockTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8855381be8e1a27d, []int{2}
}
func (m *LockTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LockTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LockTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockTransaction.Merge(m, src)
}
func (m *LockTransaction) XXX_Size() int {
	return m.Size()
}
func (m *LockTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_LockTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_LockTransaction proto.InternalMessageInfo

func (m *LockTransaction) GetRawTx() string {
	if m != nil {
		return m.RawTx
	}
	return ""
}

func (m *LockTransaction) GetVout() uint64 {
	if m != nil {
		return m.Vout
	}
	return 0
}

func (m *LockTransaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *LockTransaction) GetMintRecipient() string {
	if m != nil {
		return m.MintRecipient
	}
	return ""
}

func (m *LockTransaction) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *LockTransaction) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

// PendingMintTransaction is the metadata for a pending zenBTC mint transaction.
type PendingMintTransaction struct {
	ChainId          uint64     `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ChainType        WalletType `protobuf:"varint,2,opt,name=chain_type,json=chainType,proto3,enum=zrchain.zenbtc.WalletType" json:"chain_type,omitempty"`
	RecipientAddress string     `protobuf:"bytes,3,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	Amount           uint64     `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Creator          string     `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	KeyId            uint64     `protobuf:"varint,6,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (m *PendingMintTransaction) Reset()         { *m = PendingMintTransaction{} }
func (m *PendingMintTransaction) String() string { return proto.CompactTextString(m) }
func (*PendingMintTransaction) ProtoMessage()    {}
func (*PendingMintTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8855381be8e1a27d, []int{3}
}
func (m *PendingMintTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingMintTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingMintTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingMintTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingMintTransaction.Merge(m, src)
}
func (m *PendingMintTransaction) XXX_Size() int {
	return m.Size()
}
func (m *PendingMintTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingMintTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_PendingMintTransaction proto.InternalMessageInfo

func (m *PendingMintTransaction) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *PendingMintTransaction) GetChainType() WalletType {
	if m != nil {
		return m.ChainType
	}
	return WalletType_WALLET_TYPE_UNSPECIFIED
}

func (m *PendingMintTransaction) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *PendingMintTransaction) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PendingMintTransaction) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PendingMintTransaction) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

// PendingMintTransactions is a collection of pending mint transactions.
type PendingMintTransactions struct {
	Txs []*PendingMintTransaction `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *PendingMintTransactions) Reset()         { *m = PendingMintTransactions{} }
func (m *PendingMintTransactions) String() string { return proto.CompactTextString(m) }
func (*PendingMintTransactions) ProtoMessage()    {}
func (*PendingMintTransactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8855381be8e1a27d, []int{4}
}
func (m *PendingMintTransactions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingMintTransactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingMintTransactions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingMintTransactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingMintTransactions.Merge(m, src)
}
func (m *PendingMintTransactions) XXX_Size() int {
	return m.Size()
}
func (m *PendingMintTransactions) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingMintTransactions.DiscardUnknown(m)
}

var xxx_messageInfo_PendingMintTransactions proto.InternalMessageInfo

func (m *PendingMintTransactions) GetTxs() []*PendingMintTransaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

func init() {
	proto.RegisterEnum("zrchain.zenbtc.WalletType", WalletType_name, WalletType_value)
	proto.RegisterType((*NonceData)(nil), "zrchain.zenbtc.NonceData")
	proto.RegisterType((*RequestedBitcoinHeaders)(nil), "zrchain.zenbtc.RequestedBitcoinHeaders")
	proto.RegisterType((*LockTransaction)(nil), "zrchain.zenbtc.LockTransaction")
	proto.RegisterType((*PendingMintTransaction)(nil), "zrchain.zenbtc.PendingMintTransaction")
	proto.RegisterType((*PendingMintTransactions)(nil), "zrchain.zenbtc.PendingMintTransactions")
}

func init() { proto.RegisterFile("zrchain/zenbtc/mint.proto", fileDescriptor_8855381be8e1a27d) }

var fileDescriptor_8855381be8e1a27d = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x5d, 0x4f, 0x13, 0x41,
	0x14, 0xed, 0xd0, 0x0f, 0xe8, 0x45, 0xa1, 0x8e, 0x5a, 0x16, 0x4c, 0x9a, 0xda, 0x44, 0xd3, 0xa8,
	0x29, 0x09, 0xbc, 0xe8, 0x63, 0x81, 0x45, 0x9a, 0x94, 0x42, 0xa6, 0x2b, 0x44, 0x5f, 0x36, 0xd3,
	0xdd, 0x09, 0x9d, 0xb4, 0xcc, 0xd4, 0xd9, 0xa9, 0xb4, 0xfc, 0x0a, 0x7f, 0x8d, 0xbf, 0xc0, 0x07,
	0x1f, 0x79, 0xf4, 0xd1, 0x80, 0x3f, 0xc4, 0xcc, 0xec, 0x2e, 0x96, 0xa6, 0x6f, 0xf7, 0x9c, 0x93,
	0xb9, 0x7b, 0xce, 0xbd, 0x7b, 0x61, 0xf3, 0x5a, 0x05, 0x7d, 0xca, 0xc5, 0xf6, 0x35, 0x13, 0x3d,
	0x1d, 0x6c, 0x5f, 0x72, 0xa1, 0x1b, 0x23, 0x25, 0xb5, 0xc4, 0x6b, 0x89, 0xd4, 0x88, 0xa5, 0xda,
	0x09, 0x14, 0x3b, 0x52, 0x04, 0xec, 0x80, 0x6a, 0x8a, 0x9f, 0x41, 0x5e, 0x18, 0xe0, 0xa0, 0x2a,
	0xaa, 0xe7, 0x48, 0x0c, 0xb0, 0x03, 0xcb, 0x81, 0x1c, 0x0b, 0xcd, 0x94, 0xb3, 0x64, 0xf9, 0x14,
	0x62, 0x0c, 0xb9, 0x68, 0xc0, 0x47, 0x4e, 0xb6, 0x8a, 0xea, 0x2b, 0xc4, 0xd6, 0xb5, 0x5d, 0xd8,
	0x20, 0xec, 0xeb, 0x98, 0x45, 0x9a, 0x85, 0x7b, 0x5c, 0x07, 0x92, 0x8b, 0x23, 0x46, 0x43, 0xa6,
	0x22, 0xd3, 0xa8, 0xcf, 0xf8, 0x45, 0x5f, 0x47, 0x0e, 0xaa, 0x66, 0xeb, 0x59, 0x92, 0xc2, 0xda,
	0x0f, 0x04, 0xeb, 0x6d, 0x19, 0x0c, 0x3c, 0x45, 0x45, 0x44, 0x03, 0xcd, 0xa5, 0xc0, 0xcf, 0xa1,
	0xa0, 0xe8, 0x95, 0xaf, 0x27, 0xd6, 0x4d, 0x91, 0xe4, 0x15, 0xbd, 0xf2, 0x26, 0xe6, 0x9b, 0xdf,
	0xe4, 0x58, 0x27, 0x56, 0x6c, 0x8d, 0xcb, 0x50, 0x88, 0x98, 0x08, 0x99, 0xb2, 0x4e, 0x8a, 0x24,
	0x41, 0xf8, 0x15, 0xac, 0x99, 0xe8, 0xbe, 0x62, 0x01, 0x1f, 0x71, 0x26, 0xb4, 0x93, 0xb3, 0xfa,
	0x63, 0xc3, 0x92, 0x94, 0x34, 0xcf, 0xe9, 0xa5, 0x89, 0xe4, 0xe4, 0x6d, 0xd3, 0x04, 0xe1, 0x97,
	0xf0, 0xa8, 0x37, 0x94, 0xc1, 0xc0, 0x8f, 0x6d, 0x3a, 0x85, 0x2a, 0xaa, 0x67, 0xc9, 0xaa, 0xe5,
	0x8e, 0x2c, 0x55, 0xfb, 0x8b, 0xa0, 0x7c, 0xca, 0x44, 0xc8, 0xc5, 0xc5, 0x31, 0x17, 0x7a, 0xd6,
	0xff, 0x26, 0xac, 0xd8, 0x49, 0xfb, 0x3c, 0x4c, 0xe6, 0xb9, 0x6c, 0x71, 0x2b, 0xc4, 0x1f, 0x00,
	0x62, 0x49, 0x4f, 0x47, 0xcc, 0x26, 0x59, 0xdb, 0xd9, 0x6a, 0x3c, 0xdc, 0x4c, 0xe3, 0x9c, 0x0e,
	0x87, 0x4c, 0x7b, 0xd3, 0x11, 0x23, 0x45, 0x2b, 0x98, 0x12, 0xbf, 0x85, 0x27, 0xf7, 0x69, 0x7c,
	0x1a, 0x86, 0x8a, 0x45, 0x51, 0x92, 0xba, 0x74, 0x2f, 0x34, 0x63, 0x7e, 0x26, 0x58, 0xee, 0x41,
	0x30, 0xb3, 0x51, 0xc5, 0xa8, 0x96, 0xca, 0x26, 0x2e, 0x92, 0x14, 0x9a, 0xa1, 0x0f, 0xd8, 0xd4,
	0x58, 0x2e, 0xc4, 0xbf, 0xc0, 0x80, 0x4d, 0x5b, 0x61, 0xad, 0x0b, 0x1b, 0x8b, 0x53, 0x46, 0xf8,
	0x3d, 0x64, 0xf5, 0x24, 0x5e, 0xe8, 0xea, 0xce, 0xeb, 0xf9, 0x10, 0x8b, 0x5f, 0x11, 0xf3, 0xe4,
	0xcd, 0x4f, 0x04, 0xf0, 0x3f, 0x24, 0x7e, 0x01, 0x1b, 0xe7, 0xcd, 0x76, 0xdb, 0xf5, 0x7c, 0xef,
	0xf3, 0xa9, 0xeb, 0x7f, 0xea, 0x74, 0x4f, 0xdd, 0xfd, 0xd6, 0x61, 0xcb, 0x3d, 0x28, 0x65, 0x70,
	0x19, 0xf0, 0xac, 0xd8, 0x69, 0x7a, 0xad, 0x33, 0xb7, 0x84, 0xf0, 0x53, 0x58, 0x9f, 0xe5, 0xdd,
	0xb3, 0xe3, 0xd2, 0xd2, 0x7c, 0xa7, 0x3d, 0x6f, 0xdf, 0xf7, 0xdc, 0xae, 0xd7, 0x71, 0xbd, 0x52,
	0x76, 0x91, 0x78, 0xdc, 0x6c, 0x75, 0x8c, 0x98, 0xc3, 0x5b, 0x50, 0x9e, 0x17, 0x89, 0xfb, 0xd1,
	0x68, 0xf9, 0x79, 0x0b, 0xdd, 0x93, 0x76, 0xb3, 0xd3, 0x2c, 0x15, 0xf6, 0x0e, 0x7f, 0xdd, 0x56,
	0xd0, 0xcd, 0x6d, 0x05, 0xfd, 0xb9, 0xad, 0xa0, 0xef, 0x77, 0x95, 0xcc, 0xcd, 0x5d, 0x25, 0xf3,
	0xfb, 0xae, 0x92, 0xf9, 0xf2, 0xee, 0x82, 0xeb, 0xfe, 0xb8, 0xd7, 0x08, 0xe4, 0xa5, 0xb9, 0x44,
	0x25, 0x83, 0xc1, 0x90, 0xf6, 0xa2, 0xf4, 0x2a, 0x27, 0x69, 0x61, 0xfe, 0x82, 0xa8, 0x57, 0xb0,
	0x07, 0xba, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x66, 0x80, 0x78, 0x95, 0xbd, 0x03, 0x00, 0x00,
}

func (m *NonceData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonceData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NonceData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Skip {
		i--
		if m.Skip {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Counter != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x10
	}
	if m.Nonce != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RequestedBitcoinHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestedBitcoinHeaders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestedBitcoinHeaders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Heights) > 0 {
		dAtA2 := make([]byte, len(m.Heights)*10)
		var j1 int
		for _, num1 := range m.Heights {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintMint(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LockTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.Amount != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MintRecipient) > 0 {
		i -= len(m.MintRecipient)
		copy(dAtA[i:], m.MintRecipient)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintRecipient)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Vout != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Vout))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RawTx) > 0 {
		i -= len(m.RawTx)
		copy(dAtA[i:], m.RawTx)
		i = encodeVarintMint(dAtA, i, uint64(len(m.RawTx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PendingMintTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingMintTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingMintTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KeyId != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.KeyId))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Amount != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintMint(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChainType != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainId != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PendingMintTransactions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingMintTransactions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingMintTransactions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NonceData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovMint(uint64(m.Nonce))
	}
	if m.Counter != 0 {
		n += 1 + sovMint(uint64(m.Counter))
	}
	if m.Skip {
		n += 2
	}
	return n
}

func (m *RequestedBitcoinHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Heights) > 0 {
		l = 0
		for _, e := range m.Heights {
			l += sovMint(uint64(e))
		}
		n += 1 + sovMint(uint64(l)) + l
	}
	return n
}

func (m *LockTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawTx)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.Vout != 0 {
		n += 1 + sovMint(uint64(m.Vout))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.MintRecipient)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovMint(uint64(m.Amount))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovMint(uint64(m.BlockHeight))
	}
	return n
}

func (m *PendingMintTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovMint(uint64(m.ChainId))
	}
	if m.ChainType != 0 {
		n += 1 + sovMint(uint64(m.ChainType))
	}
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovMint(uint64(m.Amount))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.KeyId != 0 {
		n += 1 + sovMint(uint64(m.KeyId))
	}
	return n
}

func (m *PendingMintTransactions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NonceData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: NonceData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonceData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skip", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
			m.Skip = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *RequestedBitcoinHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: RequestedBitcoinHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestedBitcoinHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMint
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Heights = append(m.Heights, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMint
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMint
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthMint
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Heights) == 0 {
					m.Heights = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMint
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Heights = append(m.Heights, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Heights", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *LockTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: LockTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawTx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawTx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vout", wireType)
			}
			m.Vout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRecipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintRecipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *PendingMintTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: PendingMintTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingMintTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= WalletType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *PendingMintTransactions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: PendingMintTransactions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingMintTransactions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, &PendingMintTransaction{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
