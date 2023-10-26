// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/reputation.proto

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

type Reputation struct {
	PubKey             string `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Address            string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Score              string `protobuf:"bytes,4,opt,name=score,proto3" json:"score,omitempty"`
	RewardMultiplier   string `protobuf:"bytes,5,opt,name=rewardMultiplier,proto3" json:"rewardMultiplier,omitempty"`
	NetEarnings        string `protobuf:"bytes,6,opt,name=netEarnings,proto3" json:"netEarnings,omitempty"`
	LastTimeChallenged string `protobuf:"bytes,7,opt,name=lastTimeChallenged,proto3" json:"lastTimeChallenged,omitempty"`
	CoolDownTolerance  string `protobuf:"bytes,8,opt,name=coolDownTolerance,proto3" json:"coolDownTolerance,omitempty"`
	Type               string `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	StakedAmount       string `protobuf:"bytes,10,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
}

func (m *Reputation) Reset()         { *m = Reputation{} }
func (m *Reputation) String() string { return proto.CompactTextString(m) }
func (*Reputation) ProtoMessage()    {}
func (*Reputation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2e8fbfe6aa20f7, []int{0}
}
func (m *Reputation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reputation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reputation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reputation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reputation.Merge(m, src)
}
func (m *Reputation) XXX_Size() int {
	return m.Size()
}
func (m *Reputation) XXX_DiscardUnknown() {
	xxx_messageInfo_Reputation.DiscardUnknown(m)
}

var xxx_messageInfo_Reputation proto.InternalMessageInfo

func (m *Reputation) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Reputation) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Reputation) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *Reputation) GetRewardMultiplier() string {
	if m != nil {
		return m.RewardMultiplier
	}
	return ""
}

func (m *Reputation) GetNetEarnings() string {
	if m != nil {
		return m.NetEarnings
	}
	return ""
}

func (m *Reputation) GetLastTimeChallenged() string {
	if m != nil {
		return m.LastTimeChallenged
	}
	return ""
}

func (m *Reputation) GetCoolDownTolerance() string {
	if m != nil {
		return m.CoolDownTolerance
	}
	return ""
}

func (m *Reputation) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Reputation) GetStakedAmount() string {
	if m != nil {
		return m.StakedAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*Reputation)(nil), "soarchain.poa.Reputation")
}

func init() { proto.RegisterFile("poa/reputation.proto", fileDescriptor_2a2e8fbfe6aa20f7) }

var fileDescriptor_2a2e8fbfe6aa20f7 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4e, 0x02, 0x41,
	0x14, 0x85, 0x59, 0xc2, 0x8f, 0x5c, 0x35, 0xd1, 0x1b, 0x34, 0x53, 0x6d, 0x08, 0x95, 0x31, 0x86,
	0x2d, 0x7c, 0x02, 0xff, 0x2a, 0x63, 0x43, 0xa8, 0xec, 0x2e, 0xbb, 0x37, 0xb0, 0x71, 0x98, 0x3b,
	0x99, 0x99, 0x0d, 0xf2, 0x16, 0x3e, 0x8e, 0x8f, 0x60, 0x49, 0x69, 0x69, 0xe0, 0x45, 0x0c, 0xa3,
	0x12, 0x0d, 0x76, 0x73, 0xce, 0xf7, 0x65, 0x8a, 0x7b, 0xa0, 0x6b, 0x85, 0x32, 0xc7, 0xb6, 0x0a,
	0x14, 0x4a, 0x31, 0x03, 0xeb, 0x24, 0x08, 0x1e, 0x7a, 0x21, 0x97, 0x4f, 0xa9, 0x34, 0x03, 0x2b,
	0xd4, 0x7f, 0xad, 0x03, 0x0c, 0xb7, 0x0e, 0x9e, 0x42, 0xcb, 0x56, 0xe3, 0x7b, 0x5e, 0xa8, 0xa4,
	0x97, 0x9c, 0x75, 0x86, 0xdf, 0x09, 0x15, 0xb4, 0xa9, 0x28, 0x1c, 0x7b, 0xaf, 0xea, 0x11, 0xfc,
	0x44, 0xec, 0x42, 0xd3, 0xe7, 0xe2, 0x58, 0x35, 0x62, 0xff, 0x15, 0xf0, 0x1c, 0x8e, 0x1c, 0xcf,
	0xc9, 0x15, 0x0f, 0x95, 0x0e, 0xa5, 0xd5, 0x25, 0x3b, 0xd5, 0x8c, 0xc2, 0x4e, 0x8f, 0x3d, 0xd8,
	0x37, 0x1c, 0xee, 0xc8, 0x99, 0xd2, 0x4c, 0xbc, 0x6a, 0x45, 0xed, 0x77, 0x85, 0x03, 0x40, 0x4d,
	0x3e, 0x8c, 0xca, 0x19, 0xdf, 0x4c, 0x49, 0x6b, 0x36, 0x13, 0x2e, 0x54, 0x3b, 0x8a, 0xff, 0x10,
	0xbc, 0x80, 0xe3, 0x5c, 0x44, 0xdf, 0xca, 0xdc, 0x8c, 0x44, 0xb3, 0x23, 0x93, 0xb3, 0xda, 0x8b,
	0xfa, 0x2e, 0x40, 0x84, 0x46, 0x58, 0x58, 0x56, 0x9d, 0x28, 0xc4, 0x37, 0xf6, 0xe1, 0xc0, 0x07,
	0x7a, 0xe2, 0xe2, 0x6a, 0x26, 0x95, 0x09, 0x0a, 0x22, 0xfb, 0xd3, 0x5d, 0x67, 0x6f, 0xab, 0x34,
	0x59, 0xae, 0xd2, 0xe4, 0x63, 0x95, 0x26, 0x2f, 0xeb, 0xb4, 0xb6, 0x5c, 0xa7, 0xb5, 0xf7, 0x75,
	0x5a, 0x7b, 0x3c, 0xd9, 0xde, 0x38, 0x7b, 0xce, 0x36, 0x2b, 0x6c, 0xfe, 0xf4, 0xe3, 0x56, 0x5c,
	0xe0, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xce, 0xf2, 0xe3, 0xcb, 0x99, 0x01, 0x00, 0x00,
}

func (m *Reputation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reputation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reputation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StakedAmount) > 0 {
		i -= len(m.StakedAmount)
		copy(dAtA[i:], m.StakedAmount)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.StakedAmount)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.CoolDownTolerance) > 0 {
		i -= len(m.CoolDownTolerance)
		copy(dAtA[i:], m.CoolDownTolerance)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.CoolDownTolerance)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.LastTimeChallenged) > 0 {
		i -= len(m.LastTimeChallenged)
		copy(dAtA[i:], m.LastTimeChallenged)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.LastTimeChallenged)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.NetEarnings) > 0 {
		i -= len(m.NetEarnings)
		copy(dAtA[i:], m.NetEarnings)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.NetEarnings)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RewardMultiplier) > 0 {
		i -= len(m.RewardMultiplier)
		copy(dAtA[i:], m.RewardMultiplier)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.RewardMultiplier)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Score) > 0 {
		i -= len(m.Score)
		copy(dAtA[i:], m.Score)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.Score)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintReputation(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReputation(dAtA []byte, offset int, v uint64) int {
	offset -= sovReputation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Reputation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	l = len(m.Score)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	l = len(m.RewardMultiplier)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	l = len(m.NetEarnings)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	l = len(m.LastTimeChallenged)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	l = len(m.CoolDownTolerance)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	l = len(m.StakedAmount)
	if l > 0 {
		n += 1 + l + sovReputation(uint64(l))
	}
	return n
}

func sovReputation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReputation(x uint64) (n int) {
	return sovReputation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Reputation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReputation
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
			return fmt.Errorf("proto: Reputation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reputation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
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
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Score = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardMultiplier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetEarnings", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetEarnings = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTimeChallenged", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastTimeChallenged = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoolDownTolerance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoolDownTolerance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReputation
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
				return ErrInvalidLengthReputation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReputation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakedAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReputation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReputation
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
func skipReputation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReputation
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
					return 0, ErrIntOverflowReputation
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
					return 0, ErrIntOverflowReputation
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
				return 0, ErrInvalidLengthReputation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReputation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReputation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReputation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReputation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReputation = fmt.Errorf("proto: unexpected end of group")
)
