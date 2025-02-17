// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zrchain/zenbtc/supply.proto

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

type Supply struct {
	CustodiedBTC  uint64 `protobuf:"varint,1,opt,name=custodiedBTC,proto3" json:"custodiedBTC,omitempty"`
	MintedZenBTC  uint64 `protobuf:"varint,2,opt,name=mintedZenBTC,proto3" json:"mintedZenBTC,omitempty"`
	PendingZenBTC uint64 `protobuf:"varint,3,opt,name=pendingZenBTC,proto3" json:"pendingZenBTC,omitempty"`
}

func (m *Supply) Reset()         { *m = Supply{} }
func (m *Supply) String() string { return proto.CompactTextString(m) }
func (*Supply) ProtoMessage()    {}
func (*Supply) Descriptor() ([]byte, []int) {
	return fileDescriptor_98c0cec278312a34, []int{0}
}
func (m *Supply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Supply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Supply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supply.Merge(m, src)
}
func (m *Supply) XXX_Size() int {
	return m.Size()
}
func (m *Supply) XXX_DiscardUnknown() {
	xxx_messageInfo_Supply.DiscardUnknown(m)
}

var xxx_messageInfo_Supply proto.InternalMessageInfo

func (m *Supply) GetCustodiedBTC() uint64 {
	if m != nil {
		return m.CustodiedBTC
	}
	return 0
}

func (m *Supply) GetMintedZenBTC() uint64 {
	if m != nil {
		return m.MintedZenBTC
	}
	return 0
}

func (m *Supply) GetPendingZenBTC() uint64 {
	if m != nil {
		return m.PendingZenBTC
	}
	return 0
}

func init() {
	proto.RegisterType((*Supply)(nil), "zrchain.zenbtc.Supply")
}

func init() { proto.RegisterFile("zrchain/zenbtc/supply.proto", fileDescriptor_98c0cec278312a34) }

var fileDescriptor_98c0cec278312a34 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xae, 0x2a, 0x4a, 0xce,
	0x48, 0xcc, 0xcc, 0xd3, 0xaf, 0x4a, 0xcd, 0x4b, 0x2a, 0x49, 0xd6, 0x2f, 0x2e, 0x2d, 0x28, 0xc8,
	0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x4a, 0xea, 0x41, 0x24, 0x95, 0xca,
	0xb8, 0xd8, 0x82, 0xc1, 0xf2, 0x42, 0x4a, 0x5c, 0x3c, 0xc9, 0xa5, 0xc5, 0x25, 0xf9, 0x29, 0x99,
	0xa9, 0x29, 0x4e, 0x21, 0xce, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x28, 0x62, 0x20, 0x35,
	0xb9, 0x99, 0x79, 0x25, 0xa9, 0x29, 0x51, 0xa9, 0x79, 0x20, 0x35, 0x4c, 0x10, 0x35, 0xc8, 0x62,
	0x42, 0x2a, 0x5c, 0xbc, 0x05, 0xa9, 0x79, 0x29, 0x99, 0x79, 0xe9, 0x50, 0x45, 0xcc, 0x60, 0x45,
	0xa8, 0x82, 0x4e, 0x6e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x93,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0x0b, 0xf2, 0x41, 0x51, 0x7e, 0x72, 0x76,
	0x4e, 0x62, 0x52, 0x31, 0xcc, 0x37, 0x15, 0x30, 0x46, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b,
	0xd8, 0x5b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x2a, 0xac, 0x28, 0xf5, 0x00, 0x00,
	0x00,
}

func (m *Supply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PendingZenBTC != 0 {
		i = encodeVarintSupply(dAtA, i, uint64(m.PendingZenBTC))
		i--
		dAtA[i] = 0x18
	}
	if m.MintedZenBTC != 0 {
		i = encodeVarintSupply(dAtA, i, uint64(m.MintedZenBTC))
		i--
		dAtA[i] = 0x10
	}
	if m.CustodiedBTC != 0 {
		i = encodeVarintSupply(dAtA, i, uint64(m.CustodiedBTC))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSupply(dAtA []byte, offset int, v uint64) int {
	offset -= sovSupply(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Supply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CustodiedBTC != 0 {
		n += 1 + sovSupply(uint64(m.CustodiedBTC))
	}
	if m.MintedZenBTC != 0 {
		n += 1 + sovSupply(uint64(m.MintedZenBTC))
	}
	if m.PendingZenBTC != 0 {
		n += 1 + sovSupply(uint64(m.PendingZenBTC))
	}
	return n
}

func sovSupply(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSupply(x uint64) (n int) {
	return sovSupply(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Supply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupply
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
			return fmt.Errorf("proto: Supply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodiedBTC", wireType)
			}
			m.CustodiedBTC = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CustodiedBTC |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintedZenBTC", wireType)
			}
			m.MintedZenBTC = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintedZenBTC |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingZenBTC", wireType)
			}
			m.PendingZenBTC = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupply
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PendingZenBTC |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSupply(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSupply
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
func skipSupply(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSupply
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
					return 0, ErrIntOverflowSupply
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
					return 0, ErrIntOverflowSupply
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
				return 0, ErrInvalidLengthSupply
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSupply
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSupply
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSupply        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSupply          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSupply = fmt.Errorf("proto: unexpected end of group")
)
