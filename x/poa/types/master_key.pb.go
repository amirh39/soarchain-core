// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/master_key.proto

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

type MasterKey struct {
	MasterCertificate string `protobuf:"bytes,1,opt,name=masterCertificate,proto3" json:"masterCertificate,omitempty"`
	MasterAccount     string `protobuf:"bytes,2,opt,name=masterAccount,proto3" json:"masterAccount,omitempty"`
}

func (m *MasterKey) Reset()         { *m = MasterKey{} }
func (m *MasterKey) String() string { return proto.CompactTextString(m) }
func (*MasterKey) ProtoMessage()    {}
func (*MasterKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5191e4521fe648e, []int{0}
}
func (m *MasterKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MasterKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MasterKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MasterKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MasterKey.Merge(m, src)
}
func (m *MasterKey) XXX_Size() int {
	return m.Size()
}
func (m *MasterKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MasterKey.DiscardUnknown(m)
}

var xxx_messageInfo_MasterKey proto.InternalMessageInfo

func (m *MasterKey) GetMasterCertificate() string {
	if m != nil {
		return m.MasterCertificate
	}
	return ""
}

func (m *MasterKey) GetMasterAccount() string {
	if m != nil {
		return m.MasterAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*MasterKey)(nil), "soarchain.poa.MasterKey")
}

func init() { proto.RegisterFile("poa/master_key.proto", fileDescriptor_c5191e4521fe648e) }

var fileDescriptor_c5191e4521fe648e = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xc8, 0x4f, 0xd4,
	0xcf, 0x4d, 0x2c, 0x2e, 0x49, 0x2d, 0x8a, 0xcf, 0x4e, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x2d, 0xce, 0x4f, 0x2c, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0xc8, 0x4f, 0x54,
	0x8a, 0xe7, 0xe2, 0xf4, 0x05, 0x2b, 0xf1, 0x4e, 0xad, 0x14, 0xd2, 0xe1, 0x12, 0x84, 0xa8, 0x77,
	0x4e, 0x2d, 0x2a, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0xc2, 0x94, 0x10, 0x52, 0xe1, 0xe2, 0x85, 0x08, 0x3a, 0x26, 0x27, 0xe7, 0x97, 0xe6, 0x95,
	0x48, 0x30, 0x81, 0x55, 0xa2, 0x0a, 0x3a, 0xe9, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x94, 0x28, 0xdc, 0x25, 0xfa, 0x15, 0xfa, 0x20, 0xb7, 0x96, 0x54, 0x16, 0xa4, 0x16,
	0x27, 0xb1, 0x81, 0xdd, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x54, 0xb0, 0x98, 0xc4, 0xbf,
	0x00, 0x00, 0x00,
}

func (m *MasterKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MasterKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MasterKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MasterAccount) > 0 {
		i -= len(m.MasterAccount)
		copy(dAtA[i:], m.MasterAccount)
		i = encodeVarintMasterKey(dAtA, i, uint64(len(m.MasterAccount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MasterCertificate) > 0 {
		i -= len(m.MasterCertificate)
		copy(dAtA[i:], m.MasterCertificate)
		i = encodeVarintMasterKey(dAtA, i, uint64(len(m.MasterCertificate)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMasterKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovMasterKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MasterKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MasterCertificate)
	if l > 0 {
		n += 1 + l + sovMasterKey(uint64(l))
	}
	l = len(m.MasterAccount)
	if l > 0 {
		n += 1 + l + sovMasterKey(uint64(l))
	}
	return n
}

func sovMasterKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMasterKey(x uint64) (n int) {
	return sovMasterKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MasterKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMasterKey
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
			return fmt.Errorf("proto: MasterKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MasterKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterCertificate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMasterKey
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
				return ErrInvalidLengthMasterKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMasterKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MasterCertificate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMasterKey
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
				return ErrInvalidLengthMasterKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMasterKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MasterAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMasterKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMasterKey
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
func skipMasterKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMasterKey
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
					return 0, ErrIntOverflowMasterKey
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
					return 0, ErrIntOverflowMasterKey
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
				return 0, ErrInvalidLengthMasterKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMasterKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMasterKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMasterKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMasterKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMasterKey = fmt.Errorf("proto: unexpected end of group")
)
