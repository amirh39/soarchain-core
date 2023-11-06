// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/challenger_did.proto

package types

import (
	fmt "fmt"
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

type ChallengerDid struct {
	Id                  string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PubKey              string                     `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Keys                *Keys                      `protobuf:"bytes,3,opt,name=keys,proto3" json:"keys,omitempty"`
	VerificationMethods []*VerificationMethod      `protobuf:"bytes,4,rep,name=verificationMethods,json=verificationMethod,proto3" json:"verificationMethods,omitempty"`
	Authentications     []VerificationRelationship `protobuf:"bytes,5,rep,name=authentications,json=authentication,proto3,customtype=VerificationRelationship" json:"authentications,omitempty"`
	Services            []*Service                 `protobuf:"bytes,6,rep,name=services,json=service,proto3" json:"services,omitempty"`
	IpAddress           string                     `protobuf:"bytes,8,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	Created             string                     `protobuf:"bytes,9,opt,name=created,proto3" json:"created,omitempty"`
	Updated             string                     `protobuf:"bytes,10,opt,name=updated,proto3" json:"updated,omitempty"`
	Address             string                     `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ChallengerDid) Reset()         { *m = ChallengerDid{} }
func (m *ChallengerDid) String() string { return proto.CompactTextString(m) }
func (*ChallengerDid) ProtoMessage()    {}
func (*ChallengerDid) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0387ada292a4415, []int{0}
}
func (m *ChallengerDid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChallengerDid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChallengerDid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChallengerDid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengerDid.Merge(m, src)
}
func (m *ChallengerDid) XXX_Size() int {
	return m.Size()
}
func (m *ChallengerDid) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengerDid.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengerDid proto.InternalMessageInfo

func (m *ChallengerDid) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChallengerDid) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *ChallengerDid) GetKeys() *Keys {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ChallengerDid) GetVerificationMethods() []*VerificationMethod {
	if m != nil {
		return m.VerificationMethods
	}
	return nil
}

func (m *ChallengerDid) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ChallengerDid) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ChallengerDid) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *ChallengerDid) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *ChallengerDid) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// ChallengerDidWithSeq defines a message for Did with a sequence number for preventing replay attacks.
type ChallengerDidWithSeq struct {
	Document *ChallengerDid `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	Sequence uint64         `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *ChallengerDidWithSeq) Reset()         { *m = ChallengerDidWithSeq{} }
func (m *ChallengerDidWithSeq) String() string { return proto.CompactTextString(m) }
func (*ChallengerDidWithSeq) ProtoMessage()    {}
func (*ChallengerDidWithSeq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0387ada292a4415, []int{1}
}
func (m *ChallengerDidWithSeq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChallengerDidWithSeq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChallengerDidWithSeq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChallengerDidWithSeq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengerDidWithSeq.Merge(m, src)
}
func (m *ChallengerDidWithSeq) XXX_Size() int {
	return m.Size()
}
func (m *ChallengerDidWithSeq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengerDidWithSeq.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengerDidWithSeq proto.InternalMessageInfo

func (m *ChallengerDidWithSeq) GetDocument() *ChallengerDid {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *ChallengerDidWithSeq) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// ChallengerDataWithSeq defines a message for data with a sequence number for preventing replay attacks.
type ChallengerDataWithSeq struct {
	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Sequence uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *ChallengerDataWithSeq) Reset()         { *m = ChallengerDataWithSeq{} }
func (m *ChallengerDataWithSeq) String() string { return proto.CompactTextString(m) }
func (*ChallengerDataWithSeq) ProtoMessage()    {}
func (*ChallengerDataWithSeq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0387ada292a4415, []int{2}
}
func (m *ChallengerDataWithSeq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChallengerDataWithSeq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChallengerDataWithSeq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChallengerDataWithSeq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengerDataWithSeq.Merge(m, src)
}
func (m *ChallengerDataWithSeq) XXX_Size() int {
	return m.Size()
}
func (m *ChallengerDataWithSeq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengerDataWithSeq.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengerDataWithSeq proto.InternalMessageInfo

func (m *ChallengerDataWithSeq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ChallengerDataWithSeq) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*ChallengerDid)(nil), "soarchain.did.ChallengerDid")
	proto.RegisterType((*ChallengerDidWithSeq)(nil), "soarchain.did.ChallengerDidWithSeq")
	proto.RegisterType((*ChallengerDataWithSeq)(nil), "soarchain.did.ChallengerDataWithSeq")
}

func init() { proto.RegisterFile("did/challenger_did.proto", fileDescriptor_b0387ada292a4415) }

var fileDescriptor_b0387ada292a4415 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x93, 0x90, 0x26, 0x1b, 0x5a, 0xa4, 0x6d, 0x5a, 0xad, 0xa2, 0xc8, 0x84, 0x5c, 0x9a,
	0x93, 0x2d, 0xc2, 0x85, 0x2b, 0x05, 0x89, 0x43, 0xc5, 0x65, 0x2b, 0x81, 0xc4, 0x05, 0x6d, 0x3d,
	0x43, 0xbc, 0xc2, 0xec, 0xba, 0xde, 0x75, 0x85, 0xdf, 0x82, 0x77, 0xe2, 0xc2, 0xb1, 0x47, 0xc4,
	0x01, 0xa1, 0xe4, 0x45, 0x90, 0xd7, 0xae, 0x8b, 0x53, 0xb5, 0xb7, 0xfd, 0xe6, 0xfb, 0x99, 0xb1,
	0x67, 0x08, 0x03, 0x09, 0x61, 0x14, 0x8b, 0x24, 0x41, 0xb5, 0xc6, 0xec, 0x13, 0x48, 0x08, 0xd2,
	0x4c, 0x5b, 0x4d, 0xf7, 0x8d, 0x16, 0x59, 0x14, 0x0b, 0xa9, 0x02, 0x90, 0x30, 0x9d, 0xac, 0xf5,
	0x5a, 0x3b, 0x26, 0x2c, 0x5f, 0x95, 0x68, 0x3a, 0x71, 0xf6, 0x44, 0xa2, 0xb2, 0xb7, 0xd6, 0xc5,
	0x8f, 0x1e, 0xd9, 0x7f, 0xdd, 0x64, 0xbe, 0x91, 0x40, 0x0f, 0x48, 0x57, 0x02, 0xf3, 0xe6, 0xde,
	0x72, 0xc4, 0xbb, 0x12, 0xe8, 0x31, 0x19, 0xa4, 0xf9, 0xc5, 0x19, 0x16, 0xac, 0xeb, 0x6a, 0x35,
	0xa2, 0x27, 0xa4, 0xff, 0x05, 0x0b, 0xc3, 0x7a, 0x73, 0x6f, 0x39, 0x5e, 0x1d, 0x06, 0xad, 0x19,
	0x82, 0x33, 0x2c, 0x0c, 0x77, 0x02, 0xca, 0xc9, 0xe1, 0x15, 0x66, 0xf2, 0xb3, 0x8c, 0x84, 0x95,
	0x5a, 0xbd, 0x43, 0x1b, 0x6b, 0x30, 0xac, 0x3f, 0xef, 0x2d, 0xc7, 0xab, 0x67, 0x3b, 0xbe, 0xf7,
	0x77, 0x94, 0x9c, 0xde, 0x75, 0x53, 0x45, 0x9e, 0x88, 0xdc, 0xc6, 0xa8, 0x6c, 0x5d, 0x37, 0xec,
	0x91, 0xcb, 0x3b, 0x79, 0x20, 0x8f, 0x63, 0x52, 0x69, 0x63, 0x99, 0x9e, 0xce, 0x7e, 0xff, 0x79,
	0xca, 0xee, 0x63, 0xf9, 0x41, 0x3b, 0x9d, 0x3e, 0x27, 0x43, 0x83, 0xd9, 0x95, 0x8c, 0xd0, 0xb0,
	0x81, 0x6b, 0x74, 0xbc, 0xd3, 0xe8, 0xbc, 0xa2, 0xf9, 0x5e, 0xad, 0xa3, 0x33, 0x32, 0x92, 0xe9,
	0x2b, 0x80, 0x0c, 0x8d, 0x61, 0x43, 0xf7, 0xeb, 0x6e, 0x0b, 0x94, 0x91, 0xbd, 0x28, 0x43, 0x61,
	0x11, 0xd8, 0xc8, 0x71, 0x37, 0xb0, 0x64, 0xf2, 0x14, 0x1c, 0x43, 0x2a, 0xa6, 0x86, 0x25, 0x23,
	0xea, 0xbc, 0x71, 0xc5, 0xd4, 0x70, 0x91, 0x90, 0x49, 0x6b, 0x89, 0x1f, 0xa4, 0x8d, 0xcf, 0xf1,
	0x92, 0xbe, 0x24, 0x43, 0xd0, 0x51, 0xfe, 0x15, 0x95, 0x75, 0x1b, 0x1d, 0xaf, 0x66, 0x3b, 0x63,
	0xb7, 0x6c, 0xbc, 0x51, 0xd3, 0x69, 0xf9, 0xc1, 0x97, 0x39, 0xaa, 0x08, 0xdd, 0xde, 0xfb, 0xbc,
	0xc1, 0x8b, 0xb7, 0xe4, 0xe8, 0x3f, 0x9b, 0xb0, 0xe2, 0xa6, 0x1d, 0x25, 0x7d, 0x10, 0x56, 0xb8,
	0x56, 0x8f, 0xb9, 0x7b, 0x3f, 0x14, 0x74, 0x1a, 0xfe, 0xdc, 0xf8, 0xde, 0xf5, 0xc6, 0xf7, 0xfe,
	0x6e, 0x7c, 0xef, 0xfb, 0xd6, 0xef, 0x5c, 0x6f, 0xfd, 0xce, 0xaf, 0xad, 0xdf, 0xf9, 0x78, 0xd4,
	0x4c, 0x19, 0x7e, 0x0b, 0xcb, 0xc3, 0xb5, 0x45, 0x8a, 0xe6, 0x62, 0xe0, 0x8e, 0xf6, 0xc5, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0xf6, 0xeb, 0x07, 0x0b, 0x03, 0x00, 0x00,
}

func (m *ChallengerDid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChallengerDid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChallengerDid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintChallengerDid(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Updated) > 0 {
		i -= len(m.Updated)
		copy(dAtA[i:], m.Updated)
		i = encodeVarintChallengerDid(dAtA, i, uint64(len(m.Updated)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Created) > 0 {
		i -= len(m.Created)
		copy(dAtA[i:], m.Created)
		i = encodeVarintChallengerDid(dAtA, i, uint64(len(m.Created)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.IpAddress) > 0 {
		i -= len(m.IpAddress)
		copy(dAtA[i:], m.IpAddress)
		i = encodeVarintChallengerDid(dAtA, i, uint64(len(m.IpAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Services[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintChallengerDid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Authentications) > 0 {
		for iNdEx := len(m.Authentications) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Authentications[iNdEx].Size()
				i -= size
				if _, err := m.Authentications[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintChallengerDid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.VerificationMethods) > 0 {
		for iNdEx := len(m.VerificationMethods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerificationMethods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintChallengerDid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Keys != nil {
		{
			size, err := m.Keys.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintChallengerDid(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintChallengerDid(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintChallengerDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChallengerDidWithSeq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChallengerDidWithSeq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChallengerDidWithSeq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintChallengerDid(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x10
	}
	if m.Document != nil {
		{
			size, err := m.Document.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintChallengerDid(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChallengerDataWithSeq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChallengerDataWithSeq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChallengerDataWithSeq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintChallengerDid(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintChallengerDid(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChallengerDid(dAtA []byte, offset int, v uint64) int {
	offset -= sovChallengerDid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChallengerDid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	if m.Keys != nil {
		l = m.Keys.Size()
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	if len(m.VerificationMethods) > 0 {
		for _, e := range m.VerificationMethods {
			l = e.Size()
			n += 1 + l + sovChallengerDid(uint64(l))
		}
	}
	if len(m.Authentications) > 0 {
		for _, e := range m.Authentications {
			l = e.Size()
			n += 1 + l + sovChallengerDid(uint64(l))
		}
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovChallengerDid(uint64(l))
		}
	}
	l = len(m.IpAddress)
	if l > 0 {
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	l = len(m.Updated)
	if l > 0 {
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	return n
}

func (m *ChallengerDidWithSeq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Document != nil {
		l = m.Document.Size()
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovChallengerDid(uint64(m.Sequence))
	}
	return n
}

func (m *ChallengerDataWithSeq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovChallengerDid(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovChallengerDid(uint64(m.Sequence))
	}
	return n
}

func sovChallengerDid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChallengerDid(x uint64) (n int) {
	return sovChallengerDid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChallengerDid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChallengerDid
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
			return fmt.Errorf("proto: ChallengerDid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChallengerDid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Keys == nil {
				m.Keys = &Keys{}
			}
			if err := m.Keys.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethods = append(m.VerificationMethods, &VerificationMethod{})
			if err := m.VerificationMethods[len(m.VerificationMethods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentications", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authentications = append(m.Authentications, VerificationRelationship{})
			if err := m.Authentications[len(m.Authentications)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &Service{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Created = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Updated = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChallengerDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChallengerDid
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
func (m *ChallengerDidWithSeq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChallengerDid
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
			return fmt.Errorf("proto: ChallengerDidWithSeq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChallengerDidWithSeq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Document", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
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
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Document == nil {
				m.Document = &ChallengerDid{}
			}
			if err := m.Document.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChallengerDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChallengerDid
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
func (m *ChallengerDataWithSeq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChallengerDid
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
			return fmt.Errorf("proto: ChallengerDataWithSeq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChallengerDataWithSeq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChallengerDid
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChallengerDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallengerDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChallengerDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChallengerDid
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
func skipChallengerDid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChallengerDid
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
					return 0, ErrIntOverflowChallengerDid
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
					return 0, ErrIntOverflowChallengerDid
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
				return 0, ErrInvalidLengthChallengerDid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChallengerDid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChallengerDid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChallengerDid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChallengerDid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChallengerDid = fmt.Errorf("proto: unexpected end of group")
)
