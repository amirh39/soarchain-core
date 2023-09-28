// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/runner.proto

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

type Runner struct {
	PubKey             string `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Address            string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Score              string `protobuf:"bytes,3,opt,name=score,proto3" json:"score,omitempty"`
	RewardMultiplier   string `protobuf:"bytes,4,opt,name=rewardMultiplier,proto3" json:"rewardMultiplier,omitempty"`
	StakedAmount       string `protobuf:"bytes,5,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
	NetEarnings        string `protobuf:"bytes,6,opt,name=netEarnings,proto3" json:"netEarnings,omitempty"`
	IpAddress          string `protobuf:"bytes,7,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	LastTimeChallenged string `protobuf:"bytes,8,opt,name=lastTimeChallenged,proto3" json:"lastTimeChallenged,omitempty"`
	CoolDownTolerance  string `protobuf:"bytes,9,opt,name=coolDownTolerance,proto3" json:"coolDownTolerance,omitempty"`
	Type               string `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *Runner) Reset()         { *m = Runner{} }
func (m *Runner) String() string { return proto.CompactTextString(m) }
func (*Runner) ProtoMessage()    {}
func (*Runner) Descriptor() ([]byte, []int) {
	return fileDescriptor_54a904d1c202788c, []int{0}
}
func (m *Runner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Runner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Runner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Runner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runner.Merge(m, src)
}
func (m *Runner) XXX_Size() int {
	return m.Size()
}
func (m *Runner) XXX_DiscardUnknown() {
	xxx_messageInfo_Runner.DiscardUnknown(m)
}

var xxx_messageInfo_Runner proto.InternalMessageInfo

func (m *Runner) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Runner) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Runner) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *Runner) GetRewardMultiplier() string {
	if m != nil {
		return m.RewardMultiplier
	}
	return ""
}

func (m *Runner) GetStakedAmount() string {
	if m != nil {
		return m.StakedAmount
	}
	return ""
}

func (m *Runner) GetNetEarnings() string {
	if m != nil {
		return m.NetEarnings
	}
	return ""
}

func (m *Runner) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Runner) GetLastTimeChallenged() string {
	if m != nil {
		return m.LastTimeChallenged
	}
	return ""
}

func (m *Runner) GetCoolDownTolerance() string {
	if m != nil {
		return m.CoolDownTolerance
	}
	return ""
}

func (m *Runner) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Runner)(nil), "soarchain.poa.Runner")
}

func init() { proto.RegisterFile("poa/runner.proto", fileDescriptor_54a904d1c202788c) }

var fileDescriptor_54a904d1c202788c = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4e, 0x02, 0x41,
	0x14, 0x46, 0x59, 0x84, 0x45, 0xae, 0x9a, 0xe0, 0x8d, 0x9a, 0x29, 0xcc, 0x86, 0x50, 0x19, 0x63,
	0xd8, 0xc2, 0x27, 0xc0, 0x9f, 0xca, 0xd8, 0x10, 0x2a, 0xbb, 0x61, 0xf7, 0x06, 0x36, 0x0e, 0x33,
	0x93, 0x99, 0xd9, 0x20, 0x6f, 0xe1, 0x63, 0x59, 0x12, 0x2b, 0x4b, 0xc3, 0xbe, 0x88, 0xe1, 0x8a,
	0x7f, 0xc1, 0x6e, 0xbf, 0x73, 0x4e, 0xb1, 0x99, 0x0b, 0x1d, 0x6b, 0x64, 0xea, 0x4a, 0xad, 0xc9,
	0xf5, 0xad, 0x33, 0xc1, 0xe0, 0x81, 0x37, 0xd2, 0x65, 0x53, 0x59, 0xe8, 0xbe, 0x35, 0xb2, 0xf7,
	0x5a, 0x87, 0x78, 0xc8, 0x1e, 0x4f, 0x20, 0xb6, 0xe5, 0xf8, 0x8e, 0x16, 0x22, 0xea, 0x46, 0x67,
	0xed, 0xe1, 0x66, 0xa1, 0x80, 0x96, 0xcc, 0x73, 0x47, 0xde, 0x8b, 0x3a, 0x8b, 0xaf, 0x89, 0x47,
	0xd0, 0xf4, 0x99, 0x71, 0x24, 0x76, 0x98, 0x7f, 0x0e, 0x3c, 0x87, 0x8e, 0xa3, 0xb9, 0x74, 0xf9,
	0x7d, 0xa9, 0x42, 0x61, 0x55, 0x41, 0x4e, 0x34, 0x38, 0xd8, 0xe2, 0xd8, 0x83, 0x7d, 0x1f, 0xe4,
	0x23, 0xe5, 0x83, 0x99, 0x29, 0x75, 0x10, 0x4d, 0xee, 0xfe, 0x30, 0xec, 0xc2, 0x9e, 0xa6, 0x70,
	0x2b, 0x9d, 0x2e, 0xf4, 0xc4, 0x8b, 0x98, 0x93, 0xdf, 0x08, 0x4f, 0xa1, 0x5d, 0xd8, 0xc1, 0xe6,
	0x1f, 0x5b, 0xec, 0x7f, 0x00, 0xf6, 0x01, 0x95, 0xf4, 0x61, 0x54, 0xcc, 0xe8, 0x7a, 0x2a, 0x95,
	0x22, 0x3d, 0xa1, 0x5c, 0xec, 0x72, 0xf6, 0x8f, 0xc1, 0x0b, 0x38, 0xcc, 0x8c, 0x51, 0x37, 0x66,
	0xae, 0x47, 0x46, 0x91, 0x93, 0x3a, 0x23, 0xd1, 0xe6, 0x7c, 0x5b, 0x20, 0x42, 0x23, 0x2c, 0x2c,
	0x09, 0xe0, 0x80, 0xbf, 0xaf, 0xd2, 0x97, 0x55, 0x12, 0x2d, 0x57, 0x49, 0xf4, 0xbe, 0x4a, 0xa2,
	0xe7, 0x2a, 0xa9, 0x2d, 0xab, 0xa4, 0xf6, 0x56, 0x25, 0xb5, 0x87, 0xe3, 0xef, 0xd7, 0x4f, 0x9f,
	0xd2, 0xf5, 0x6d, 0xd6, 0xbd, 0x1f, 0xc7, 0x7c, 0x9b, 0xcb, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x47, 0xdd, 0xcd, 0x60, 0xaf, 0x01, 0x00, 0x00,
}

func (m *Runner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Runner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Runner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.CoolDownTolerance) > 0 {
		i -= len(m.CoolDownTolerance)
		copy(dAtA[i:], m.CoolDownTolerance)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.CoolDownTolerance)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.LastTimeChallenged) > 0 {
		i -= len(m.LastTimeChallenged)
		copy(dAtA[i:], m.LastTimeChallenged)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.LastTimeChallenged)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.IpAddress) > 0 {
		i -= len(m.IpAddress)
		copy(dAtA[i:], m.IpAddress)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.IpAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.NetEarnings) > 0 {
		i -= len(m.NetEarnings)
		copy(dAtA[i:], m.NetEarnings)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.NetEarnings)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.StakedAmount) > 0 {
		i -= len(m.StakedAmount)
		copy(dAtA[i:], m.StakedAmount)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.StakedAmount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RewardMultiplier) > 0 {
		i -= len(m.RewardMultiplier)
		copy(dAtA[i:], m.RewardMultiplier)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.RewardMultiplier)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Score) > 0 {
		i -= len(m.Score)
		copy(dAtA[i:], m.Score)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.Score)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintRunner(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRunner(dAtA []byte, offset int, v uint64) int {
	offset -= sovRunner(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Runner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.Score)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.RewardMultiplier)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.StakedAmount)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.NetEarnings)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.IpAddress)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.LastTimeChallenged)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.CoolDownTolerance)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovRunner(uint64(l))
	}
	return n
}

func sovRunner(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRunner(x uint64) (n int) {
	return sovRunner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Runner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRunner
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
			return fmt.Errorf("proto: Runner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Runner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
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
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Score = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardMultiplier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
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
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetEarnings = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTimeChallenged", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastTimeChallenged = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoolDownTolerance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoolDownTolerance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRunner
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
				return ErrInvalidLengthRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRunner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRunner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRunner
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
func skipRunner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRunner
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
					return 0, ErrIntOverflowRunner
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
					return 0, ErrIntOverflowRunner
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
				return 0, ErrInvalidLengthRunner
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRunner
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRunner
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRunner        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRunner          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRunner = fmt.Errorf("proto: unexpected end of group")
)
