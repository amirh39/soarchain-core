// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: soarmint/minter.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Minter struct {
	Inflation        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
	Phase            uint64                                 `protobuf:"varint,2,opt,name=phase,proto3" json:"phase,omitempty"`
	StartPhaseBlock  uint64                                 `protobuf:"varint,3,opt,name=startPhaseBlock,proto3" json:"startPhaseBlock,omitempty"`
	AnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=annualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"annualProvisions"`
	TargetSupply     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=targetSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"targetSupply"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0290e9e1a060a7f8, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetPhase() uint64 {
	if m != nil {
		return m.Phase
	}
	return 0
}

func (m *Minter) GetStartPhaseBlock() uint64 {
	if m != nil {
		return m.StartPhaseBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "soarchain.soarmint.Minter")
}

func init() { proto.RegisterFile("soarmint/minter.proto", fileDescriptor_0290e9e1a060a7f8) }

var fileDescriptor_0290e9e1a060a7f8 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0xce, 0x4f, 0x2c,
	0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x07, 0x11, 0xa9, 0x45, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x42, 0x20, 0xe1, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x3d, 0x98, 0x02, 0x29, 0x91, 0xf4, 0xfc, 0xf4,
	0x7c, 0xb0, 0xb4, 0x3e, 0x88, 0x05, 0x51, 0xa9, 0x74, 0x8c, 0x89, 0x8b, 0xcd, 0x17, 0xac, 0x55,
	0xc8, 0x87, 0x8b, 0x33, 0x33, 0x2f, 0x2d, 0x27, 0xb1, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xd3, 0x49, 0xef, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0xd5, 0xd2, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xa1, 0x94,
	0x6e, 0x71, 0x4a, 0xb6, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0xb1, 0x9e, 0x4b, 0x6a, 0x72, 0x10, 0xc2,
	0x00, 0x21, 0x11, 0x2e, 0xd6, 0x82, 0x8c, 0xc4, 0xe2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96,
	0x20, 0x08, 0x47, 0x48, 0x83, 0x8b, 0xbf, 0xb8, 0x24, 0xb1, 0xa8, 0x24, 0x00, 0xc4, 0x73, 0xca,
	0xc9, 0x4f, 0xce, 0x96, 0x60, 0x06, 0xcb, 0xa3, 0x0b, 0x0b, 0x45, 0x71, 0x09, 0x24, 0xe6, 0xe5,
	0x95, 0x26, 0xe6, 0x04, 0x14, 0xe5, 0x97, 0x65, 0x16, 0x67, 0xe6, 0xe7, 0x15, 0x4b, 0xb0, 0x90,
	0xe5, 0x28, 0x0c, 0x73, 0x84, 0x82, 0xb8, 0x78, 0x4a, 0x12, 0x8b, 0xd2, 0x53, 0x4b, 0x82, 0x4b,
	0x0b, 0x0a, 0x72, 0x2a, 0x25, 0x58, 0x49, 0x36, 0xd7, 0x33, 0xaf, 0x24, 0x08, 0xc5, 0x0c, 0x27,
	0x93, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x82, 0x47, 0x86, 0x7e,
	0x85, 0x3e, 0x3c, 0xbe, 0xc0, 0xe6, 0x24, 0xb1, 0x81, 0x63, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xce, 0x60, 0x82, 0x27, 0xc8, 0x01, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TargetSupply.Size()
		i -= size
		if _, err := m.TargetSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.StartPhaseBlock != 0 {
		i = encodeVarintMinter(dAtA, i, uint64(m.StartPhaseBlock))
		i--
		dAtA[i] = 0x18
	}
	if m.Phase != 0 {
		i = encodeVarintMinter(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMinter(dAtA []byte, offset int, v uint64) int {
	offset -= sovMinter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMinter(uint64(l))
	if m.Phase != 0 {
		n += 1 + sovMinter(uint64(m.Phase))
	}
	if m.StartPhaseBlock != 0 {
		n += 1 + sovMinter(uint64(m.StartPhaseBlock))
	}
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = m.TargetSupply.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func sovMinter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMinter(x uint64) (n int) {
	return sovMinter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			m.Phase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Phase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartPhaseBlock", wireType)
			}
			m.StartPhaseBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartPhaseBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TargetSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func skipMinter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMinter
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
					return 0, ErrIntOverflowMinter
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
					return 0, ErrIntOverflowMinter
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
				return 0, ErrInvalidLengthMinter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMinter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMinter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMinter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMinter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMinter = fmt.Errorf("proto: unexpected end of group")
)
