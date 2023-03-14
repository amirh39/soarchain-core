// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/factory_keys.proto

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

type FactoryKeys struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FactoryCert string `protobuf:"bytes,2,opt,name=factoryCert,proto3" json:"factoryCert,omitempty"`
}

func (m *FactoryKeys) Reset()         { *m = FactoryKeys{} }
func (m *FactoryKeys) String() string { return proto.CompactTextString(m) }
func (*FactoryKeys) ProtoMessage()    {}
func (*FactoryKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa6dcd5ed0512afa, []int{0}
}
func (m *FactoryKeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FactoryKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FactoryKeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FactoryKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FactoryKeys.Merge(m, src)
}
func (m *FactoryKeys) XXX_Size() int {
	return m.Size()
}
func (m *FactoryKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_FactoryKeys.DiscardUnknown(m)
}

var xxx_messageInfo_FactoryKeys proto.InternalMessageInfo

func (m *FactoryKeys) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FactoryKeys) GetFactoryCert() string {
	if m != nil {
		return m.FactoryCert
	}
	return ""
}

func init() {
	proto.RegisterType((*FactoryKeys)(nil), "soarchain.poa.FactoryKeys")
}

func init() { proto.RegisterFile("poa/factory_keys.proto", fileDescriptor_aa6dcd5ed0512afa) }

var fileDescriptor_aa6dcd5ed0512afa = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xc8, 0x4f, 0xd4,
	0x4f, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0xaa, 0x8c, 0xcf, 0x4e, 0xad, 0x2c, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0xce, 0x4f, 0x2c, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0xc8,
	0x4f, 0x54, 0xb2, 0xe7, 0xe2, 0x76, 0x83, 0x28, 0xf2, 0x4e, 0xad, 0x2c, 0x16, 0xe2, 0xe3, 0x62,
	0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x52, 0xe0, 0xe2,
	0x86, 0x9a, 0xe1, 0x9c, 0x5a, 0x54, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2c, 0xe4,
	0xa4, 0x7f, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xa2, 0x70, 0x9b, 0xf4,
	0x2b, 0xf4, 0x41, 0xae, 0x29, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xc3, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xea, 0x03, 0x56, 0x20, 0xa1, 0x00, 0x00, 0x00,
}

func (m *FactoryKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FactoryKeys) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FactoryKeys) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FactoryCert) > 0 {
		i -= len(m.FactoryCert)
		copy(dAtA[i:], m.FactoryCert)
		i = encodeVarintFactoryKeys(dAtA, i, uint64(len(m.FactoryCert)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintFactoryKeys(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFactoryKeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovFactoryKeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FactoryKeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovFactoryKeys(uint64(m.Id))
	}
	l = len(m.FactoryCert)
	if l > 0 {
		n += 1 + l + sovFactoryKeys(uint64(l))
	}
	return n
}

func sovFactoryKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFactoryKeys(x uint64) (n int) {
	return sovFactoryKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FactoryKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFactoryKeys
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
			return fmt.Errorf("proto: FactoryKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FactoryKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFactoryKeys
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
				return fmt.Errorf("proto: wrong wireType = %d for field FactoryCert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFactoryKeys
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
				return ErrInvalidLengthFactoryKeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFactoryKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FactoryCert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFactoryKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFactoryKeys
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
func skipFactoryKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFactoryKeys
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
					return 0, ErrIntOverflowFactoryKeys
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
					return 0, ErrIntOverflowFactoryKeys
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
				return 0, ErrInvalidLengthFactoryKeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFactoryKeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFactoryKeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFactoryKeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFactoryKeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFactoryKeys = fmt.Errorf("proto: unexpected end of group")
)
