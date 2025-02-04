// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/zenbtc/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	EthBatcherAddr      string   `protobuf:"bytes,1,opt,name=ethBatcherAddr,proto3" json:"ethBatcherAddr,omitempty"`
	DepositKeyringAddr  string   `protobuf:"bytes,2,opt,name=depositKeyringAddr,proto3" json:"depositKeyringAddr,omitempty"`
	EthMinterKeyID      uint64   `protobuf:"varint,3,opt,name=ethMinterKeyID,proto3" json:"ethMinterKeyID,omitempty"`
	WithdrawerKeyID     uint64   `protobuf:"varint,4,opt,name=withdrawerKeyID,proto3" json:"withdrawerKeyID,omitempty"` // Deprecated: Do not use.
	UnstakerKeyID       uint64   `protobuf:"varint,5,opt,name=unstakerKeyID,proto3" json:"unstakerKeyID,omitempty"`
	RewardsDepositKeyID uint64   `protobuf:"varint,6,opt,name=rewardsDepositKeyID,proto3" json:"rewardsDepositKeyID,omitempty"`
	ChangeAddressKeyIDs []uint64 `protobuf:"varint,7,rep,packed,name=changeAddressKeyIDs,proto3" json:"changeAddressKeyIDs,omitempty"`
	BitcoinProxyAddress string   `protobuf:"bytes,8,opt,name=bitcoinProxyAddress,proto3" json:"bitcoinProxyAddress,omitempty"`
	Authority           string   `protobuf:"bytes,9,opt,name=authority,proto3" json:"authority,omitempty"`
	StakerKeyID         uint64   `protobuf:"varint,10,opt,name=stakerKeyID,proto3" json:"stakerKeyID,omitempty"`
	CompleterKeyID      uint64   `protobuf:"varint,11,opt,name=completerKeyID,proto3" json:"completerKeyID,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_53cfc222fdb324be, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEthBatcherAddr() string {
	if m != nil {
		return m.EthBatcherAddr
	}
	return ""
}

func (m *Params) GetDepositKeyringAddr() string {
	if m != nil {
		return m.DepositKeyringAddr
	}
	return ""
}

func (m *Params) GetEthMinterKeyID() uint64 {
	if m != nil {
		return m.EthMinterKeyID
	}
	return 0
}

// Deprecated: Do not use.
func (m *Params) GetWithdrawerKeyID() uint64 {
	if m != nil {
		return m.WithdrawerKeyID
	}
	return 0
}

func (m *Params) GetUnstakerKeyID() uint64 {
	if m != nil {
		return m.UnstakerKeyID
	}
	return 0
}

func (m *Params) GetRewardsDepositKeyID() uint64 {
	if m != nil {
		return m.RewardsDepositKeyID
	}
	return 0
}

func (m *Params) GetChangeAddressKeyIDs() []uint64 {
	if m != nil {
		return m.ChangeAddressKeyIDs
	}
	return nil
}

func (m *Params) GetBitcoinProxyAddress() string {
	if m != nil {
		return m.BitcoinProxyAddress
	}
	return ""
}

func (m *Params) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *Params) GetStakerKeyID() uint64 {
	if m != nil {
		return m.StakerKeyID
	}
	return 0
}

func (m *Params) GetCompleterKeyID() uint64 {
	if m != nil {
		return m.CompleterKeyID
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "zrchain.zenbtc.Params")
}

func init() { proto.RegisterFile("zrchain/zenbtc/params.proto", fileDescriptor_53cfc222fdb324be) }

var fileDescriptor_53cfc222fdb324be = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8a, 0xda, 0x40,
	0x18, 0xc7, 0x1d, 0xb5, 0xb6, 0x8e, 0xd4, 0xd2, 0x69, 0xa1, 0xc1, 0x96, 0x34, 0x94, 0x52, 0xa4,
	0x48, 0x52, 0xe8, 0xad, 0xb7, 0x8a, 0x14, 0x44, 0x0a, 0xe2, 0xb1, 0xb7, 0xc9, 0x64, 0x48, 0x06,
	0xcd, 0x4c, 0x98, 0x19, 0xd1, 0xf8, 0x08, 0x3d, 0xf5, 0x11, 0xfa, 0x08, 0xfb, 0x18, 0x7b, 0xf4,
	0xb8, 0xc7, 0x45, 0x0f, 0xbb, 0x0f, 0xb1, 0x87, 0x65, 0x26, 0xc6, 0x55, 0xc9, 0x25, 0x7c, 0xfc,
	0xff, 0xbf, 0xff, 0xc7, 0x97, 0x6f, 0x3e, 0xf8, 0x7e, 0x23, 0x49, 0x82, 0x19, 0x0f, 0x36, 0x94,
	0x87, 0x9a, 0x04, 0x19, 0x96, 0x38, 0x55, 0x7e, 0x26, 0x85, 0x16, 0xa8, 0x7b, 0x30, 0xfd, 0xc2,
	0xec, 0xbd, 0xc6, 0x29, 0xe3, 0x22, 0xb0, 0xdf, 0x02, 0xe9, 0xbd, 0x8d, 0x45, 0x2c, 0x6c, 0x19,
	0x98, 0xaa, 0x50, 0x3f, 0x3d, 0x34, 0x60, 0x6b, 0x6a, 0x3b, 0xa1, 0x2f, 0xb0, 0x4b, 0x75, 0x32,
	0xc4, 0x9a, 0x24, 0x54, 0xfe, 0x8c, 0x22, 0xe9, 0x00, 0x0f, 0xf4, 0xdb, 0xb3, 0x0b, 0x15, 0xf9,
	0x10, 0x45, 0x34, 0x13, 0x8a, 0xe9, 0x09, 0xcd, 0x25, 0xe3, 0xb1, 0x65, 0xeb, 0x96, 0xad, 0x70,
	0x0e, 0x7d, 0x7f, 0x33, 0xae, 0xa9, 0x9c, 0xd0, 0x7c, 0x3c, 0x72, 0x1a, 0x1e, 0xe8, 0x37, 0x67,
	0x17, 0x2a, 0x1a, 0xc0, 0x57, 0x2b, 0xa6, 0x93, 0x48, 0xe2, 0x55, 0x09, 0x36, 0x0d, 0x38, 0xac,
	0x3b, 0x60, 0x76, 0x69, 0xa1, 0xcf, 0xf0, 0xe5, 0x92, 0x2b, 0x8d, 0xe7, 0x25, 0xfb, 0xcc, 0x36,
	0x3d, 0x17, 0xd1, 0x37, 0xf8, 0x46, 0xd2, 0x15, 0x96, 0x91, 0x1a, 0x1d, 0x07, 0x1b, 0x8f, 0x9c,
	0x96, 0x65, 0xab, 0x2c, 0x93, 0x20, 0x09, 0xe6, 0x31, 0x35, 0xb3, 0x53, 0xa5, 0xac, 0xaa, 0x9c,
	0xe7, 0x5e, 0xc3, 0x24, 0x2a, 0x2c, 0x93, 0x08, 0x99, 0x26, 0x82, 0xf1, 0xa9, 0x14, 0xeb, 0xfc,
	0x60, 0x3a, 0x2f, 0xec, 0x42, 0xaa, 0x2c, 0xf4, 0x01, 0xb6, 0xf1, 0x52, 0x27, 0x42, 0x32, 0x9d,
	0x3b, 0x6d, 0xcb, 0x3d, 0x09, 0xc8, 0x83, 0x9d, 0xd3, 0xff, 0x82, 0x76, 0xd6, 0x53, 0xc9, 0x6c,
	0x94, 0x88, 0x34, 0x5b, 0xd0, 0xe3, 0x46, 0x3b, 0xc5, 0x46, 0xcf, 0xd5, 0x1f, 0xde, 0xfd, 0xff,
	0x8f, 0xe0, 0xef, 0xdd, 0xd5, 0xd7, 0x77, 0xe5, 0xed, 0xac, 0xcb, 0xeb, 0x29, 0xde, 0x7c, 0xf8,
	0xeb, 0x7a, 0xe7, 0x82, 0xed, 0xce, 0x05, 0xb7, 0x3b, 0x17, 0xfc, 0xdb, 0xbb, 0xb5, 0xed, 0xde,
	0xad, 0xdd, 0xec, 0xdd, 0xda, 0x9f, 0x41, 0xcc, 0x74, 0xb2, 0x0c, 0x7d, 0x22, 0x52, 0x93, 0x91,
	0x82, 0xcc, 0x17, 0x38, 0x54, 0x65, 0xfe, 0xd8, 0x48, 0xe7, 0x19, 0x55, 0x61, 0xcb, 0x5e, 0xd3,
	0xf7, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x8a, 0x63, 0x02, 0xa5, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.EthBatcherAddr != that1.EthBatcherAddr {
		return false
	}
	if this.DepositKeyringAddr != that1.DepositKeyringAddr {
		return false
	}
	if this.EthMinterKeyID != that1.EthMinterKeyID {
		return false
	}
	if this.WithdrawerKeyID != that1.WithdrawerKeyID {
		return false
	}
	if this.UnstakerKeyID != that1.UnstakerKeyID {
		return false
	}
	if this.RewardsDepositKeyID != that1.RewardsDepositKeyID {
		return false
	}
	if len(this.ChangeAddressKeyIDs) != len(that1.ChangeAddressKeyIDs) {
		return false
	}
	for i := range this.ChangeAddressKeyIDs {
		if this.ChangeAddressKeyIDs[i] != that1.ChangeAddressKeyIDs[i] {
			return false
		}
	}
	if this.BitcoinProxyAddress != that1.BitcoinProxyAddress {
		return false
	}
	if this.Authority != that1.Authority {
		return false
	}
	if this.StakerKeyID != that1.StakerKeyID {
		return false
	}
	if this.CompleterKeyID != that1.CompleterKeyID {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CompleterKeyID != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CompleterKeyID))
		i--
		dAtA[i] = 0x58
	}
	if m.StakerKeyID != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.StakerKeyID))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.BitcoinProxyAddress) > 0 {
		i -= len(m.BitcoinProxyAddress)
		copy(dAtA[i:], m.BitcoinProxyAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BitcoinProxyAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ChangeAddressKeyIDs) > 0 {
		dAtA2 := make([]byte, len(m.ChangeAddressKeyIDs)*10)
		var j1 int
		for _, num := range m.ChangeAddressKeyIDs {
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
		i = encodeVarintParams(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x3a
	}
	if m.RewardsDepositKeyID != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RewardsDepositKeyID))
		i--
		dAtA[i] = 0x30
	}
	if m.UnstakerKeyID != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnstakerKeyID))
		i--
		dAtA[i] = 0x28
	}
	if m.WithdrawerKeyID != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.WithdrawerKeyID))
		i--
		dAtA[i] = 0x20
	}
	if m.EthMinterKeyID != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EthMinterKeyID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DepositKeyringAddr) > 0 {
		i -= len(m.DepositKeyringAddr)
		copy(dAtA[i:], m.DepositKeyringAddr)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DepositKeyringAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EthBatcherAddr) > 0 {
		i -= len(m.EthBatcherAddr)
		copy(dAtA[i:], m.EthBatcherAddr)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EthBatcherAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EthBatcherAddr)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DepositKeyringAddr)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.EthMinterKeyID != 0 {
		n += 1 + sovParams(uint64(m.EthMinterKeyID))
	}
	if m.WithdrawerKeyID != 0 {
		n += 1 + sovParams(uint64(m.WithdrawerKeyID))
	}
	if m.UnstakerKeyID != 0 {
		n += 1 + sovParams(uint64(m.UnstakerKeyID))
	}
	if m.RewardsDepositKeyID != 0 {
		n += 1 + sovParams(uint64(m.RewardsDepositKeyID))
	}
	if len(m.ChangeAddressKeyIDs) > 0 {
		l = 0
		for _, e := range m.ChangeAddressKeyIDs {
			l += sovParams(uint64(e))
		}
		n += 1 + sovParams(uint64(l)) + l
	}
	l = len(m.BitcoinProxyAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.StakerKeyID != 0 {
		n += 1 + sovParams(uint64(m.StakerKeyID))
	}
	if m.CompleterKeyID != 0 {
		n += 1 + sovParams(uint64(m.CompleterKeyID))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthBatcherAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthBatcherAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositKeyringAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositKeyringAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthMinterKeyID", wireType)
			}
			m.EthMinterKeyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthMinterKeyID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawerKeyID", wireType)
			}
			m.WithdrawerKeyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawerKeyID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakerKeyID", wireType)
			}
			m.UnstakerKeyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnstakerKeyID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsDepositKeyID", wireType)
			}
			m.RewardsDepositKeyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardsDepositKeyID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ChangeAddressKeyIDs = append(m.ChangeAddressKeyIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
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
					return ErrInvalidLengthParams
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthParams
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
				if elementCount != 0 && len(m.ChangeAddressKeyIDs) == 0 {
					m.ChangeAddressKeyIDs = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ChangeAddressKeyIDs = append(m.ChangeAddressKeyIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangeAddressKeyIDs", wireType)
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitcoinProxyAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BitcoinProxyAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerKeyID", wireType)
			}
			m.StakerKeyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakerKeyID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompleterKeyID", wireType)
			}
			m.CompleterKeyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompleterKeyID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
