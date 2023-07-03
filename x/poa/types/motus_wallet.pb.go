// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/motus_wallet.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type MotusWallet struct {
	Address string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Client  *Client `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
}

func (m *MotusWallet) Reset()         { *m = MotusWallet{} }
func (m *MotusWallet) String() string { return proto.CompactTextString(m) }
func (*MotusWallet) ProtoMessage()    {}
func (*MotusWallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_81288c7c373eeafd, []int{0}
}
func (m *MotusWallet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MotusWallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MotusWallet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MotusWallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MotusWallet.Merge(m, src)
}
func (m *MotusWallet) XXX_Size() int {
	return m.Size()
}
func (m *MotusWallet) XXX_DiscardUnknown() {
	xxx_messageInfo_MotusWallet.DiscardUnknown(m)
}

var xxx_messageInfo_MotusWallet proto.InternalMessageInfo

func (m *MotusWallet) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MotusWallet) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func init() {
	proto.RegisterType((*MotusWallet)(nil), "soarchain.poa.MotusWallet")
}

func init() { proto.RegisterFile("poa/motus_wallet.proto", fileDescriptor_81288c7c373eeafd) }

var fileDescriptor_81288c7c373eeafd = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xc8, 0x4f, 0xd4,
	0xcf, 0xcd, 0x2f, 0x29, 0x2d, 0x8e, 0x2f, 0x4f, 0xcc, 0xc9, 0x49, 0x2d, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0xce, 0x4f, 0x2c, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0xc8,
	0x4f, 0x94, 0x12, 0x00, 0x29, 0x4b, 0xce, 0xc9, 0x4c, 0xcd, 0x83, 0x2a, 0x50, 0x0a, 0xe3, 0xe2,
	0xf6, 0x05, 0x69, 0x0b, 0x07, 0xeb, 0x12, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d,
	0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x74, 0xb9, 0xd8, 0x20, 0x1a,
	0x25, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x44, 0xf5, 0x50, 0x8c, 0xd6, 0x73, 0x06, 0x4b, 0x06,
	0x41, 0x15, 0x39, 0xe9, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x28,
	0x5c, 0x9f, 0x7e, 0x85, 0x3e, 0xc8, 0x3d, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xf7,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0xaa, 0x1b, 0x98, 0xca, 0x00, 0x00, 0x00,
}

func (m *MotusWallet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MotusWallet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MotusWallet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Client != nil {
		{
			size, err := m.Client.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMotusWallet(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMotusWallet(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMotusWallet(dAtA []byte, offset int, v uint64) int {
	offset -= sovMotusWallet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MotusWallet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMotusWallet(uint64(l))
	}
	if m.Client != nil {
		l = m.Client.Size()
		n += 1 + l + sovMotusWallet(uint64(l))
	}
	return n
}

func sovMotusWallet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMotusWallet(x uint64) (n int) {
	return sovMotusWallet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MotusWallet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMotusWallet
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
			return fmt.Errorf("proto: MotusWallet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MotusWallet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMotusWallet
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
				return ErrInvalidLengthMotusWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMotusWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Client", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMotusWallet
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
				return ErrInvalidLengthMotusWallet
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMotusWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Client == nil {
				m.Client = &Client{}
			}
			if err := m.Client.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMotusWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMotusWallet
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
func skipMotusWallet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMotusWallet
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
					return 0, ErrIntOverflowMotusWallet
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
					return 0, ErrIntOverflowMotusWallet
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
				return 0, ErrInvalidLengthMotusWallet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMotusWallet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMotusWallet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMotusWallet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMotusWallet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMotusWallet = fmt.Errorf("proto: unexpected end of group")
)
