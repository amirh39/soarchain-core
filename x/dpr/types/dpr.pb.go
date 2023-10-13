// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dpr/dpr.proto

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

type Dpr struct {
	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator       string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	SupportedPIDs string   `protobuf:"bytes,3,opt,name=SupportedPIDs,proto3" json:"SupportedPIDs,omitempty"`
	IsActive      bool     `protobuf:"varint,6,opt,name=isActive,proto3" json:"isActive,omitempty"`
	Vin           []string `protobuf:"bytes,7,rep,name=vin,proto3" json:"vin,omitempty"`
	ClientPubkeys []string `protobuf:"bytes,8,rep,name=clientPubkeys,proto3" json:"clientPubkeys,omitempty"`
	Duration      uint64   `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	DPRendTime    string   `protobuf:"bytes,10,opt,name=DPRendTime,proto3" json:"DPRendTime,omitempty"`
}

func (m *Dpr) Reset()         { *m = Dpr{} }
func (m *Dpr) String() string { return proto.CompactTextString(m) }
func (*Dpr) ProtoMessage()    {}
func (*Dpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_221c37c1dd019f8d, []int{0}
}
func (m *Dpr) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Dpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Dpr.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Dpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dpr.Merge(m, src)
}
func (m *Dpr) XXX_Size() int {
	return m.Size()
}
func (m *Dpr) XXX_DiscardUnknown() {
	xxx_messageInfo_Dpr.DiscardUnknown(m)
}

var xxx_messageInfo_Dpr proto.InternalMessageInfo

func (m *Dpr) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Dpr) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Dpr) GetSupportedPIDs() string {
	if m != nil {
		return m.SupportedPIDs
	}
	return ""
}

func (m *Dpr) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Dpr) GetVin() []string {
	if m != nil {
		return m.Vin
	}
	return nil
}

func (m *Dpr) GetClientPubkeys() []string {
	if m != nil {
		return m.ClientPubkeys
	}
	return nil
}

func (m *Dpr) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Dpr) GetDPRendTime() string {
	if m != nil {
		return m.DPRendTime
	}
	return ""
}

func init() {
	proto.RegisterType((*Dpr)(nil), "soarchain.dpr.Dpr")
}

func init() { proto.RegisterFile("dpr/dpr.proto", fileDescriptor_221c37c1dd019f8d) }

var fileDescriptor_221c37c1dd019f8d = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x9b, 0x56, 0x66, 0xda, 0x07, 0x15, 0x09, 0x08, 0xc1, 0x45, 0x28, 0xe2, 0xa2, 0xab,
	0xe9, 0xc2, 0x2f, 0x50, 0xba, 0x71, 0x57, 0xaa, 0x2b, 0x77, 0x9d, 0x26, 0x60, 0x50, 0x93, 0xf0,
	0x92, 0x0e, 0xce, 0x5f, 0xf8, 0x59, 0x2e, 0x67, 0xe9, 0x52, 0x5a, 0xfc, 0x0f, 0x69, 0xc0, 0x32,
	0xb3, 0x7b, 0xf7, 0x5c, 0xb8, 0x3c, 0x0e, 0xe4, 0xc2, 0x62, 0x25, 0x2c, 0x6e, 0x2c, 0x1a, 0x6f,
	0x68, 0xee, 0x4c, 0x87, 0xfd, 0x4b, 0xa7, 0xf4, 0x46, 0x58, 0xbc, 0xfe, 0x25, 0x90, 0xd4, 0x16,
	0xe9, 0x39, 0xc4, 0x4a, 0x30, 0x52, 0x90, 0x32, 0x6b, 0x63, 0x25, 0x28, 0x83, 0x75, 0x8f, 0xb2,
	0xf3, 0x06, 0x59, 0x1c, 0xe0, 0x7f, 0xa4, 0x37, 0x90, 0x3f, 0x0e, 0xd6, 0x1a, 0xf4, 0x52, 0x34,
	0x0f, 0xb5, 0x63, 0x49, 0xe8, 0x4f, 0x21, 0xbd, 0x82, 0x54, 0xb9, 0xbb, 0xde, 0xab, 0x9d, 0x64,
	0xab, 0x82, 0x94, 0x69, 0xbb, 0x64, 0x7a, 0x01, 0xc9, 0x4e, 0x69, 0xb6, 0x2e, 0x92, 0x32, 0x6b,
	0xe7, 0x73, 0xde, 0xec, 0xdf, 0x94, 0xd4, 0xbe, 0x19, 0xb6, 0xaf, 0x72, 0xef, 0x58, 0x1a, 0xba,
	0x53, 0x38, 0x6f, 0x8a, 0x01, 0x3b, 0xaf, 0x8c, 0x66, 0x59, 0x41, 0xca, 0xb3, 0x76, 0xc9, 0x94,
	0x03, 0xd4, 0x4d, 0x2b, 0xb5, 0x78, 0x52, 0xef, 0x92, 0x41, 0x78, 0xe9, 0x88, 0xdc, 0x57, 0x5f,
	0x23, 0x27, 0x87, 0x91, 0x93, 0x9f, 0x91, 0x93, 0xcf, 0x89, 0x47, 0x87, 0x89, 0x47, 0xdf, 0x13,
	0x8f, 0x9e, 0x2f, 0x17, 0x21, 0xd5, 0xc7, 0xec, 0xa9, 0xf2, 0x7b, 0x2b, 0xdd, 0x76, 0x15, 0x74,
	0xdd, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x63, 0xb6, 0xea, 0x3f, 0x01, 0x00, 0x00,
}

func (m *Dpr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dpr) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Dpr) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DPRendTime) > 0 {
		i -= len(m.DPRendTime)
		copy(dAtA[i:], m.DPRendTime)
		i = encodeVarintDpr(dAtA, i, uint64(len(m.DPRendTime)))
		i--
		dAtA[i] = 0x52
	}
	if m.Duration != 0 {
		i = encodeVarintDpr(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x48
	}
	if len(m.ClientPubkeys) > 0 {
		for iNdEx := len(m.ClientPubkeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClientPubkeys[iNdEx])
			copy(dAtA[i:], m.ClientPubkeys[iNdEx])
			i = encodeVarintDpr(dAtA, i, uint64(len(m.ClientPubkeys[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Vin) > 0 {
		for iNdEx := len(m.Vin) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Vin[iNdEx])
			copy(dAtA[i:], m.Vin[iNdEx])
			i = encodeVarintDpr(dAtA, i, uint64(len(m.Vin[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.SupportedPIDs) > 0 {
		i -= len(m.SupportedPIDs)
		copy(dAtA[i:], m.SupportedPIDs)
		i = encodeVarintDpr(dAtA, i, uint64(len(m.SupportedPIDs)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDpr(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDpr(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDpr(dAtA []byte, offset int, v uint64) int {
	offset -= sovDpr(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Dpr) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDpr(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDpr(uint64(l))
	}
	l = len(m.SupportedPIDs)
	if l > 0 {
		n += 1 + l + sovDpr(uint64(l))
	}
	if m.IsActive {
		n += 2
	}
	if len(m.Vin) > 0 {
		for _, s := range m.Vin {
			l = len(s)
			n += 1 + l + sovDpr(uint64(l))
		}
	}
	if len(m.ClientPubkeys) > 0 {
		for _, s := range m.ClientPubkeys {
			l = len(s)
			n += 1 + l + sovDpr(uint64(l))
		}
	}
	if m.Duration != 0 {
		n += 1 + sovDpr(uint64(m.Duration))
	}
	l = len(m.DPRendTime)
	if l > 0 {
		n += 1 + l + sovDpr(uint64(l))
	}
	return n
}

func sovDpr(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDpr(x uint64) (n int) {
	return sovDpr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Dpr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpr
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
			return fmt.Errorf("proto: Dpr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dpr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpr
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
				return ErrInvalidLengthDpr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDpr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpr
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
				return ErrInvalidLengthDpr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDpr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedPIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpr
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
				return ErrInvalidLengthDpr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDpr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportedPIDs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpr
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
			m.IsActive = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpr
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
				return ErrInvalidLengthDpr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDpr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vin = append(m.Vin, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientPubkeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpr
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
				return ErrInvalidLengthDpr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDpr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientPubkeys = append(m.ClientPubkeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DPRendTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpr
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
				return ErrInvalidLengthDpr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDpr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DPRendTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDpr
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
func skipDpr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDpr
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
					return 0, ErrIntOverflowDpr
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
					return 0, ErrIntOverflowDpr
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
				return 0, ErrInvalidLengthDpr
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDpr
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDpr
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDpr        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDpr          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDpr = fmt.Errorf("proto: unexpected end of group")
)
