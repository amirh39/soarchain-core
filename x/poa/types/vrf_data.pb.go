// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/vrf_data.proto

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

type VrfData struct {
	Index         string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator       string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Count         string `protobuf:"bytes,3,opt,name=count,proto3" json:"count,omitempty"`
	Vrv           string `protobuf:"bytes,4,opt,name=vrv,proto3" json:"vrv,omitempty"`
	Multiplier    string `protobuf:"bytes,5,opt,name=Multiplier,proto3" json:"Multiplier,omitempty"`
	Proof         string `protobuf:"bytes,6,opt,name=proof,proto3" json:"proof,omitempty"`
	Pubkey        string `protobuf:"bytes,7,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Message       string `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
	ParsedVrv     string `protobuf:"bytes,9,opt,name=parsedVrv,proto3" json:"parsedVrv,omitempty"`
	FloatVrv      string `protobuf:"bytes,10,opt,name=floatVrv,proto3" json:"floatVrv,omitempty"`
	FinalVrv      string `protobuf:"bytes,11,opt,name=finalVrv,proto3" json:"finalVrv,omitempty"`
	FinalVrvFloat string `protobuf:"bytes,12,opt,name=finalVrvFloat,proto3" json:"finalVrvFloat,omitempty"`
}

func (m *VrfData) Reset()         { *m = VrfData{} }
func (m *VrfData) String() string { return proto.CompactTextString(m) }
func (*VrfData) ProtoMessage()    {}
func (*VrfData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb49cb41f6e8d202, []int{0}
}
func (m *VrfData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VrfData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VrfData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VrfData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrfData.Merge(m, src)
}
func (m *VrfData) XXX_Size() int {
	return m.Size()
}
func (m *VrfData) XXX_DiscardUnknown() {
	xxx_messageInfo_VrfData.DiscardUnknown(m)
}

var xxx_messageInfo_VrfData proto.InternalMessageInfo

func (m *VrfData) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *VrfData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *VrfData) GetCount() string {
	if m != nil {
		return m.Count
	}
	return ""
}

func (m *VrfData) GetVrv() string {
	if m != nil {
		return m.Vrv
	}
	return ""
}

func (m *VrfData) GetMultiplier() string {
	if m != nil {
		return m.Multiplier
	}
	return ""
}

func (m *VrfData) GetProof() string {
	if m != nil {
		return m.Proof
	}
	return ""
}

func (m *VrfData) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *VrfData) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *VrfData) GetParsedVrv() string {
	if m != nil {
		return m.ParsedVrv
	}
	return ""
}

func (m *VrfData) GetFloatVrv() string {
	if m != nil {
		return m.FloatVrv
	}
	return ""
}

func (m *VrfData) GetFinalVrv() string {
	if m != nil {
		return m.FinalVrv
	}
	return ""
}

func (m *VrfData) GetFinalVrvFloat() string {
	if m != nil {
		return m.FinalVrvFloat
	}
	return ""
}

func init() {
	proto.RegisterType((*VrfData)(nil), "soarchain.poa.VrfData")
}

func init() { proto.RegisterFile("poa/vrf_data.proto", fileDescriptor_fb49cb41f6e8d202) }

var fileDescriptor_fb49cb41f6e8d202 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x93, 0xd6, 0x26, 0xcd, 0x68, 0x41, 0x16, 0x95, 0x45, 0x64, 0x11, 0xf1, 0xe0, 0xa9,
	0x39, 0xf8, 0x06, 0x22, 0xde, 0xbc, 0x78, 0xe8, 0xc1, 0x8b, 0x4c, 0x93, 0x8d, 0x06, 0x63, 0x76,
	0xd9, 0x6c, 0x42, 0xfb, 0x16, 0xbe, 0x94, 0xe0, 0xb1, 0x47, 0x8f, 0x92, 0xbc, 0x88, 0xec, 0x24,
	0xad, 0x7a, 0xdb, 0xef, 0xff, 0x7e, 0x86, 0x61, 0x07, 0x98, 0x56, 0x18, 0x37, 0x26, 0x7b, 0x4a,
	0xd1, 0xe2, 0x5c, 0x1b, 0x65, 0x15, 0x9b, 0x55, 0x0a, 0x4d, 0xf2, 0x82, 0x79, 0x39, 0xd7, 0x0a,
	0x2f, 0x3e, 0x46, 0x10, 0x2e, 0x4c, 0x76, 0x8b, 0x16, 0xd9, 0x11, 0x4c, 0xf2, 0x32, 0x95, 0x2b,
	0xee, 0x9f, 0xfb, 0x57, 0xd1, 0x43, 0x0f, 0x8c, 0x43, 0x98, 0x18, 0x89, 0x56, 0x19, 0x3e, 0xa2,
	0x7c, 0x8b, 0xae, 0x9f, 0xa8, 0xba, 0xb4, 0x7c, 0xdc, 0xf7, 0x09, 0xd8, 0x21, 0x8c, 0x1b, 0xd3,
	0xf0, 0x3d, 0xca, 0xdc, 0x93, 0x09, 0x80, 0xfb, 0xba, 0xb0, 0xb9, 0x2e, 0x72, 0x69, 0xf8, 0x84,
	0xc4, 0x9f, 0xc4, 0xcd, 0xd1, 0x46, 0xa9, 0x8c, 0x07, 0xfd, 0x1c, 0x02, 0x76, 0x02, 0x81, 0xae,
	0x97, 0xaf, 0x72, 0xcd, 0x43, 0x8a, 0x07, 0x72, 0xfb, 0xbc, 0xc9, 0xaa, 0xc2, 0x67, 0xc9, 0xa7,
	0xfd, 0x3e, 0x03, 0xb2, 0x33, 0x88, 0x34, 0x9a, 0x4a, 0xa6, 0x0b, 0xd3, 0xf0, 0x88, 0xdc, 0x6f,
	0xc0, 0x4e, 0x61, 0x9a, 0x15, 0x0a, 0xad, 0x93, 0x40, 0x72, 0xc7, 0xe4, 0xf2, 0x12, 0x0b, 0xe7,
	0xf6, 0x07, 0x37, 0x30, 0xbb, 0x84, 0xd9, 0xf6, 0x7d, 0xe7, 0xfa, 0xfc, 0x80, 0x0a, 0xff, 0xc3,
	0x9b, 0xf8, 0xb3, 0x15, 0xfe, 0xa6, 0x15, 0xfe, 0x77, 0x2b, 0xfc, 0xf7, 0x4e, 0x78, 0x9b, 0x4e,
	0x78, 0x5f, 0x9d, 0xf0, 0x1e, 0x8f, 0x77, 0x1f, 0x1e, 0xaf, 0x62, 0x77, 0x10, 0xbb, 0xd6, 0xb2,
	0x5a, 0x06, 0x74, 0x8e, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x94, 0xf0, 0x77, 0xa4,
	0x01, 0x00, 0x00,
}

func (m *VrfData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VrfData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VrfData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FinalVrvFloat) > 0 {
		i -= len(m.FinalVrvFloat)
		copy(dAtA[i:], m.FinalVrvFloat)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.FinalVrvFloat)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.FinalVrv) > 0 {
		i -= len(m.FinalVrv)
		copy(dAtA[i:], m.FinalVrv)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.FinalVrv)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.FloatVrv) > 0 {
		i -= len(m.FloatVrv)
		copy(dAtA[i:], m.FloatVrv)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.FloatVrv)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.ParsedVrv) > 0 {
		i -= len(m.ParsedVrv)
		copy(dAtA[i:], m.ParsedVrv)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.ParsedVrv)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Multiplier) > 0 {
		i -= len(m.Multiplier)
		copy(dAtA[i:], m.Multiplier)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.Multiplier)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Vrv) > 0 {
		i -= len(m.Vrv)
		copy(dAtA[i:], m.Vrv)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.Vrv)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Count) > 0 {
		i -= len(m.Count)
		copy(dAtA[i:], m.Count)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.Count)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintVrfData(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVrfData(dAtA []byte, offset int, v uint64) int {
	offset -= sovVrfData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VrfData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.Count)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.Vrv)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.Multiplier)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.ParsedVrv)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.FloatVrv)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.FinalVrv)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	l = len(m.FinalVrvFloat)
	if l > 0 {
		n += 1 + l + sovVrfData(uint64(l))
	}
	return n
}

func sovVrfData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVrfData(x uint64) (n int) {
	return sovVrfData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VrfData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVrfData
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
			return fmt.Errorf("proto: VrfData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VrfData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Count = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vrv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vrv = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Multiplier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParsedVrv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParsedVrv = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatVrv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FloatVrv = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalVrv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalVrv = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalVrvFloat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfData
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
				return ErrInvalidLengthVrfData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalVrvFloat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVrfData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVrfData
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
func skipVrfData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVrfData
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
					return 0, ErrIntOverflowVrfData
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
					return 0, ErrIntOverflowVrfData
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
				return 0, ErrInvalidLengthVrfData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVrfData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVrfData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVrfData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVrfData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVrfData = fmt.Errorf("proto: unexpected end of group")
)
