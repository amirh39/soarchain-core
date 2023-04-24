// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/challenger.proto

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

type Challenger struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PubKey       string `protobuf:"bytes,2,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	Address      string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Score        string `protobuf:"bytes,4,opt,name=score,proto3" json:"score,omitempty"`
	StakedAmount string `protobuf:"bytes,5,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
	NetEarnings  string `protobuf:"bytes,6,opt,name=netEarnings,proto3" json:"netEarnings,omitempty"`
	Type         string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	IpAddr       string `protobuf:"bytes,8,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
}

func (m *Challenger) Reset()         { *m = Challenger{} }
func (m *Challenger) String() string { return proto.CompactTextString(m) }
func (*Challenger) ProtoMessage()    {}
func (*Challenger) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2723ec770faa160, []int{0}
}
func (m *Challenger) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Challenger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Challenger.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Challenger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Challenger.Merge(m, src)
}
func (m *Challenger) XXX_Size() int {
	return m.Size()
}
func (m *Challenger) XXX_DiscardUnknown() {
	xxx_messageInfo_Challenger.DiscardUnknown(m)
}

var xxx_messageInfo_Challenger proto.InternalMessageInfo

func (m *Challenger) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Challenger) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Challenger) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Challenger) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *Challenger) GetStakedAmount() string {
	if m != nil {
		return m.StakedAmount
	}
	return ""
}

func (m *Challenger) GetNetEarnings() string {
	if m != nil {
		return m.NetEarnings
	}
	return ""
}

func (m *Challenger) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Challenger) GetIpAddr() string {
	if m != nil {
		return m.IpAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*Challenger)(nil), "soarchain.poa.Challenger")
}

func init() { proto.RegisterFile("poa/challenger.proto", fileDescriptor_d2723ec770faa160) }

var fileDescriptor_d2723ec770faa160 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x63, 0x68, 0x53, 0x38, 0x60, 0xb1, 0x0a, 0xba, 0xc9, 0xaa, 0x3a, 0x31, 0x35, 0x03,
	0xbf, 0xa0, 0x20, 0x26, 0x16, 0xc4, 0xc8, 0x76, 0x4d, 0x4e, 0x6d, 0x44, 0xb1, 0xad, 0xb3, 0x2b,
	0xd1, 0x7f, 0xc1, 0xcf, 0x62, 0xec, 0xc8, 0x58, 0x25, 0x7f, 0x04, 0xc5, 0x29, 0x15, 0x6c, 0xfe,
	0xde, 0xbb, 0x67, 0x3d, 0x3d, 0x18, 0x7b, 0x47, 0x45, 0xb9, 0xa2, 0xf5, 0x9a, 0xed, 0x92, 0x65,
	0xe6, 0xc5, 0x45, 0xa7, 0xaf, 0x82, 0x23, 0x29, 0x57, 0x54, 0xdb, 0x99, 0x77, 0x34, 0xdd, 0x2b,
	0x80, 0x87, 0xe3, 0x8d, 0x46, 0x18, 0x95, 0xc2, 0x14, 0x9d, 0xa0, 0x9a, 0xa8, 0xdb, 0xf3, 0x97,
	0x5f, 0xd4, 0x37, 0x90, 0x3f, 0x6f, 0x16, 0x4f, 0xbc, 0xc5, 0x93, 0x64, 0x1c, 0xa8, 0x4b, 0x50,
	0x55, 0x09, 0x87, 0x80, 0xa7, 0x7d, 0xe2, 0x80, 0x7a, 0x0c, 0xc3, 0x50, 0x3a, 0x61, 0x1c, 0x24,
	0xbd, 0x07, 0x3d, 0x85, 0xcb, 0x10, 0xe9, 0x8d, 0xab, 0xf9, 0xbb, 0xdb, 0xd8, 0x88, 0xc3, 0x64,
	0xfe, 0xd3, 0xf4, 0x04, 0x2e, 0x2c, 0xc7, 0x47, 0x12, 0x5b, 0xdb, 0x65, 0xc0, 0x3c, 0x9d, 0xfc,
	0x95, 0xb4, 0x86, 0x41, 0xdc, 0x7a, 0xc6, 0x51, 0xb2, 0xd2, 0xbb, 0x6b, 0x58, 0xfb, 0x79, 0x55,
	0x09, 0x9e, 0xf5, 0x0d, 0x7b, 0xba, 0x2f, 0xbe, 0x1a, 0xa3, 0x76, 0x8d, 0x51, 0xfb, 0xc6, 0xa8,
	0xcf, 0xd6, 0x64, 0xbb, 0xd6, 0x64, 0xdf, 0xad, 0xc9, 0x5e, 0xaf, 0x8f, 0x5b, 0x14, 0x1f, 0x45,
	0xb7, 0x56, 0xf7, 0x4f, 0x58, 0xe4, 0x69, 0xa9, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0xdd, 0xef, 0xb1, 0x41, 0x01, 0x00, 0x00,
}

func (m *Challenger) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Challenger) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Challenger) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IpAddr) > 0 {
		i -= len(m.IpAddr)
		copy(dAtA[i:], m.IpAddr)
		i = encodeVarintChallenger(dAtA, i, uint64(len(m.IpAddr)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintChallenger(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.NetEarnings) > 0 {
		i -= len(m.NetEarnings)
		copy(dAtA[i:], m.NetEarnings)
		i = encodeVarintChallenger(dAtA, i, uint64(len(m.NetEarnings)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.StakedAmount) > 0 {
		i -= len(m.StakedAmount)
		copy(dAtA[i:], m.StakedAmount)
		i = encodeVarintChallenger(dAtA, i, uint64(len(m.StakedAmount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Score) > 0 {
		i -= len(m.Score)
		copy(dAtA[i:], m.Score)
		i = encodeVarintChallenger(dAtA, i, uint64(len(m.Score)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintChallenger(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintChallenger(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintChallenger(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChallenger(dAtA []byte, offset int, v uint64) int {
	offset -= sovChallenger(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Challenger) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovChallenger(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovChallenger(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovChallenger(uint64(l))
	}
	l = len(m.Score)
	if l > 0 {
		n += 1 + l + sovChallenger(uint64(l))
	}
	l = len(m.StakedAmount)
	if l > 0 {
		n += 1 + l + sovChallenger(uint64(l))
	}
	l = len(m.NetEarnings)
	if l > 0 {
		n += 1 + l + sovChallenger(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovChallenger(uint64(l))
	}
	l = len(m.IpAddr)
	if l > 0 {
		n += 1 + l + sovChallenger(uint64(l))
	}
	return n
}

func sovChallenger(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChallenger(x uint64) (n int) {
	return sovChallenger(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Challenger) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChallenger
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
			return fmt.Errorf("proto: Challenger: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Challenger: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenger
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
				return ErrInvalidLengthChallenger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenger
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
				return ErrInvalidLengthChallenger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenger
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
				return ErrInvalidLengthChallenger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenger
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
				return ErrInvalidLengthChallenger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Score = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenger
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
				return ErrInvalidLengthChallenger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakedAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetEarnings", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenger
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
				return ErrInvalidLengthChallenger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetEarnings = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenger
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
				return ErrInvalidLengthChallenger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenger
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
				return ErrInvalidLengthChallenger
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenger
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChallenger(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChallenger
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
func skipChallenger(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChallenger
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
					return 0, ErrIntOverflowChallenger
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
					return 0, ErrIntOverflowChallenger
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
				return 0, ErrInvalidLengthChallenger
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChallenger
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChallenger
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChallenger        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChallenger          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChallenger = fmt.Errorf("proto: unexpected end of group")
)
