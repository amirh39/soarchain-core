// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/epoch_data.proto

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

type EpochData struct {
	TotalEpochs                   uint64 `protobuf:"varint,1,opt,name=totalEpochs,proto3" json:"totalEpochs,omitempty"`
	EpochV2VRX                    string `protobuf:"bytes,2,opt,name=epochV2VRX,proto3" json:"epochV2VRX,omitempty"`
	EpochV2VBX                    string `protobuf:"bytes,3,opt,name=epochV2VBX,proto3" json:"epochV2VBX,omitempty"`
	EpochV2NBX                    string `protobuf:"bytes,4,opt,name=epochV2NBX,proto3" json:"epochV2NBX,omitempty"`
	EpochRunner                   string `protobuf:"bytes,5,opt,name=epochRunner,proto3" json:"epochRunner,omitempty"`
	EpochChallenger               string `protobuf:"bytes,6,opt,name=epochChallenger,proto3" json:"epochChallenger,omitempty"`
	V2VRXtotalChallenges          uint64 `protobuf:"varint,7,opt,name=V2VRXtotalChallenges,proto3" json:"V2VRXtotalChallenges,omitempty"`
	V2VBXtotalChallenges          uint64 `protobuf:"varint,8,opt,name=V2VBXtotalChallenges,proto3" json:"V2VBXtotalChallenges,omitempty"`
	V2NBXtotalChallenges          uint64 `protobuf:"varint,9,opt,name=V2NBXtotalChallenges,proto3" json:"V2NBXtotalChallenges,omitempty"`
	RunnerTotalChallenges         uint64 `protobuf:"varint,10,opt,name=RunnerTotalChallenges,proto3" json:"RunnerTotalChallenges,omitempty"`
	ChallengerTotalChallenges     uint64 `protobuf:"varint,11,opt,name=ChallengerTotalChallenges,proto3" json:"ChallengerTotalChallenges,omitempty"`
	V2VRXLastBlockChallenges      uint64 `protobuf:"varint,12,opt,name=V2VRXLastBlockChallenges,proto3" json:"V2VRXLastBlockChallenges,omitempty"`
	V2VBXLastBlockChallenges      uint64 `protobuf:"varint,13,opt,name=V2VBXLastBlockChallenges,proto3" json:"V2VBXLastBlockChallenges,omitempty"`
	V2NBXLastBlockChallenges      uint64 `protobuf:"varint,14,opt,name=V2NBXLastBlockChallenges,proto3" json:"V2NBXLastBlockChallenges,omitempty"`
	RunnerLastBlockChallenges     uint64 `protobuf:"varint,15,opt,name=RunnerLastBlockChallenges,proto3" json:"RunnerLastBlockChallenges,omitempty"`
	ChallengerLastBlockChallenges uint64 `protobuf:"varint,16,opt,name=ChallengerLastBlockChallenges,proto3" json:"ChallengerLastBlockChallenges,omitempty"`
	TotalChallengesPrevDay        uint64 `protobuf:"varint,17,opt,name=totalChallengesPrevDay,proto3" json:"totalChallengesPrevDay,omitempty"`
	InitialPerChallengeValue      uint64 `protobuf:"varint,18,opt,name=initialPerChallengeValue,proto3" json:"initialPerChallengeValue,omitempty"`
	V2NBXPerChallengeValue        uint64 `protobuf:"varint,19,opt,name=V2NBXPerChallengeValue,proto3" json:"V2NBXPerChallengeValue,omitempty"`
	RunnerPerChallengeValue       uint64 `protobuf:"varint,20,opt,name=RunnerPerChallengeValue,proto3" json:"RunnerPerChallengeValue,omitempty"`
	ChallengerPerChallengeValue   uint64 `protobuf:"varint,21,opt,name=ChallengerPerChallengeValue,proto3" json:"ChallengerPerChallengeValue,omitempty"`
	V2VBXPerChallengeValue        uint64 `protobuf:"varint,22,opt,name=V2VBXPerChallengeValue,proto3" json:"V2VBXPerChallengeValue,omitempty"`
	V2VRXPerChallengeValue        uint64 `protobuf:"varint,23,opt,name=V2VRXPerChallengeValue,proto3" json:"V2VRXPerChallengeValue,omitempty"`
}

func (m *EpochData) Reset()         { *m = EpochData{} }
func (m *EpochData) String() string { return proto.CompactTextString(m) }
func (*EpochData) ProtoMessage()    {}
func (*EpochData) Descriptor() ([]byte, []int) {
	return fileDescriptor_911419e87680592b, []int{0}
}
func (m *EpochData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochData.Merge(m, src)
}
func (m *EpochData) XXX_Size() int {
	return m.Size()
}
func (m *EpochData) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochData.DiscardUnknown(m)
}

var xxx_messageInfo_EpochData proto.InternalMessageInfo

func (m *EpochData) GetTotalEpochs() uint64 {
	if m != nil {
		return m.TotalEpochs
	}
	return 0
}

func (m *EpochData) GetEpochV2VRX() string {
	if m != nil {
		return m.EpochV2VRX
	}
	return ""
}

func (m *EpochData) GetEpochV2VBX() string {
	if m != nil {
		return m.EpochV2VBX
	}
	return ""
}

func (m *EpochData) GetEpochV2NBX() string {
	if m != nil {
		return m.EpochV2NBX
	}
	return ""
}

func (m *EpochData) GetEpochRunner() string {
	if m != nil {
		return m.EpochRunner
	}
	return ""
}

func (m *EpochData) GetEpochChallenger() string {
	if m != nil {
		return m.EpochChallenger
	}
	return ""
}

func (m *EpochData) GetV2VRXtotalChallenges() uint64 {
	if m != nil {
		return m.V2VRXtotalChallenges
	}
	return 0
}

func (m *EpochData) GetV2VBXtotalChallenges() uint64 {
	if m != nil {
		return m.V2VBXtotalChallenges
	}
	return 0
}

func (m *EpochData) GetV2NBXtotalChallenges() uint64 {
	if m != nil {
		return m.V2NBXtotalChallenges
	}
	return 0
}

func (m *EpochData) GetRunnerTotalChallenges() uint64 {
	if m != nil {
		return m.RunnerTotalChallenges
	}
	return 0
}

func (m *EpochData) GetChallengerTotalChallenges() uint64 {
	if m != nil {
		return m.ChallengerTotalChallenges
	}
	return 0
}

func (m *EpochData) GetV2VRXLastBlockChallenges() uint64 {
	if m != nil {
		return m.V2VRXLastBlockChallenges
	}
	return 0
}

func (m *EpochData) GetV2VBXLastBlockChallenges() uint64 {
	if m != nil {
		return m.V2VBXLastBlockChallenges
	}
	return 0
}

func (m *EpochData) GetV2NBXLastBlockChallenges() uint64 {
	if m != nil {
		return m.V2NBXLastBlockChallenges
	}
	return 0
}

func (m *EpochData) GetRunnerLastBlockChallenges() uint64 {
	if m != nil {
		return m.RunnerLastBlockChallenges
	}
	return 0
}

func (m *EpochData) GetChallengerLastBlockChallenges() uint64 {
	if m != nil {
		return m.ChallengerLastBlockChallenges
	}
	return 0
}

func (m *EpochData) GetTotalChallengesPrevDay() uint64 {
	if m != nil {
		return m.TotalChallengesPrevDay
	}
	return 0
}

func (m *EpochData) GetInitialPerChallengeValue() uint64 {
	if m != nil {
		return m.InitialPerChallengeValue
	}
	return 0
}

func (m *EpochData) GetV2NBXPerChallengeValue() uint64 {
	if m != nil {
		return m.V2NBXPerChallengeValue
	}
	return 0
}

func (m *EpochData) GetRunnerPerChallengeValue() uint64 {
	if m != nil {
		return m.RunnerPerChallengeValue
	}
	return 0
}

func (m *EpochData) GetChallengerPerChallengeValue() uint64 {
	if m != nil {
		return m.ChallengerPerChallengeValue
	}
	return 0
}

func (m *EpochData) GetV2VBXPerChallengeValue() uint64 {
	if m != nil {
		return m.V2VBXPerChallengeValue
	}
	return 0
}

func (m *EpochData) GetV2VRXPerChallengeValue() uint64 {
	if m != nil {
		return m.V2VRXPerChallengeValue
	}
	return 0
}

func init() {
	proto.RegisterType((*EpochData)(nil), "soarchain.poa.EpochData")
}

func init() { proto.RegisterFile("poa/epoch_data.proto", fileDescriptor_911419e87680592b) }

var fileDescriptor_911419e87680592b = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0x87, 0xad, 0xd6, 0x75, 0xab, 0x75, 0x5d, 0xb7, 0x5b, 0xff, 0xd9, 0x52, 0x2a, 0x4c, 0x4f,
	0x3e, 0x59, 0xe0, 0x96, 0x52, 0x4a, 0x0f, 0x65, 0xeb, 0xdc, 0x82, 0x30, 0x22, 0x08, 0x91, 0x4b,
	0xd8, 0x38, 0x4b, 0x2c, 0x22, 0xb4, 0x42, 0x5a, 0x87, 0xf8, 0x2d, 0xf2, 0x3a, 0x79, 0x83, 0x1c,
	0x7d, 0xcc, 0x31, 0xd8, 0x2f, 0x12, 0x34, 0xb2, 0x65, 0x59, 0x5a, 0xf9, 0xa8, 0xf9, 0xe6, 0xdb,
	0x9d, 0xdf, 0x08, 0x16, 0x75, 0x42, 0xc1, 0x4c, 0x1e, 0x8a, 0xd9, 0xfc, 0xe2, 0x8a, 0x49, 0x36,
	0x0a, 0x23, 0x21, 0x05, 0x6e, 0xc5, 0x82, 0x45, 0xb3, 0x39, 0xf3, 0x82, 0x51, 0x28, 0xd8, 0xf7,
	0x07, 0x1d, 0xe9, 0x27, 0x49, 0xcf, 0x84, 0x49, 0x86, 0x07, 0xa8, 0x29, 0x85, 0x64, 0x3e, 0x54,
	0x62, 0xa2, 0x0d, 0xb4, 0x61, 0xdd, 0xce, 0x97, 0xb0, 0x81, 0x10, 0x1c, 0xe9, 0x8c, 0x1d, 0xdb,
	0x25, 0xaf, 0x06, 0xda, 0x50, 0xb7, 0x73, 0x95, 0x3c, 0xa7, 0x2e, 0x79, 0x7d, 0xc8, 0x69, 0x9e,
	0x5b, 0xd4, 0x25, 0xf5, 0x03, 0x6e, 0x51, 0x37, 0x99, 0x00, 0xbe, 0xec, 0x45, 0x10, 0xf0, 0x88,
	0xbc, 0x81, 0x86, 0x7c, 0x09, 0x0f, 0x51, 0x1b, 0x3e, 0xff, 0xcf, 0x99, 0xef, 0xf3, 0xe0, 0x9a,
	0x47, 0xa4, 0x01, 0x5d, 0xc5, 0x32, 0x1e, 0xa3, 0x0e, 0x0c, 0x05, 0xf3, 0x67, 0xf5, 0x98, 0xbc,
	0x85, 0x58, 0x4a, 0xb6, 0x75, 0x68, 0xc9, 0x79, 0x97, 0x39, 0x54, 0xed, 0x58, 0x65, 0x47, 0xdf,
	0x39, 0x65, 0x86, 0x7f, 0xa2, 0x6e, 0x9a, 0xe7, 0xac, 0x20, 0x21, 0x90, 0xd4, 0x10, 0xff, 0x45,
	0x5f, 0xf6, 0xf9, 0x8a, 0x66, 0x13, 0xcc, 0xea, 0x06, 0xfc, 0x07, 0x11, 0xc8, 0x7c, 0xca, 0x62,
	0x49, 0x7d, 0x31, 0xbb, 0xc9, 0xc9, 0xef, 0x41, 0xae, 0xe4, 0x5b, 0x97, 0x2a, 0xdd, 0x56, 0xe6,
	0xd2, 0x6a, 0xd7, 0x52, 0xbb, 0x1f, 0x76, 0xae, 0x9a, 0x27, 0x89, 0xd3, 0x55, 0xa8, 0xe4, 0x76,
	0x9a, 0xb8, 0xb2, 0x01, 0x4f, 0xd0, 0xb7, 0xfd, 0x3a, 0x54, 0x27, 0x7c, 0x84, 0x13, 0x8e, 0x37,
	0xe1, 0x5f, 0xa8, 0x57, 0xf8, 0x7d, 0xd3, 0x88, 0xdf, 0x4e, 0xd8, 0x92, 0x7c, 0x02, 0xbd, 0x82,
	0x26, 0xb9, 0xbd, 0xc0, 0x93, 0x1e, 0xf3, 0xa7, 0x3c, 0xca, 0xb0, 0xc3, 0xfc, 0x05, 0x27, 0x38,
	0xcd, 0x5d, 0xc5, 0x93, 0x3b, 0x61, 0x27, 0x65, 0xf3, 0x73, 0x7a, 0xa7, 0x9a, 0xe2, 0xdf, 0xa8,
	0x9f, 0xae, 0xa3, 0x2c, 0x76, 0x40, 0xac, 0xc2, 0xf8, 0x1f, 0xfa, 0xba, 0x5f, 0x43, 0xd9, 0xee,
	0x82, 0x7d, 0xac, 0x25, 0x9d, 0xd9, 0x51, 0xcd, 0xdc, 0xdb, 0xcd, 0xac, 0xa2, 0x5b, 0xcf, 0x56,
	0x78, 0xfd, 0xcc, 0x53, 0x50, 0x6a, 0x3e, 0xae, 0x0d, 0x6d, 0xb5, 0x36, 0xb4, 0xe7, 0xb5, 0xa1,
	0xdd, 0x6f, 0x8c, 0xda, 0x6a, 0x63, 0xd4, 0x9e, 0x36, 0x46, 0xed, 0xbc, 0x9b, 0x3d, 0x72, 0xe6,
	0x9d, 0x99, 0x3c, 0x83, 0x72, 0x19, 0xf2, 0xf8, 0xb2, 0x01, 0x4f, 0xe0, 0x8f, 0x97, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0xe2, 0x2d, 0xaf, 0x1a, 0x05, 0x00, 0x00,
}

func (m *EpochData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.V2VRXPerChallengeValue != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2VRXPerChallengeValue))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb8
	}
	if m.V2VBXPerChallengeValue != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2VBXPerChallengeValue))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb0
	}
	if m.ChallengerPerChallengeValue != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.ChallengerPerChallengeValue))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa8
	}
	if m.RunnerPerChallengeValue != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.RunnerPerChallengeValue))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.V2NBXPerChallengeValue != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2NBXPerChallengeValue))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if m.InitialPerChallengeValue != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.InitialPerChallengeValue))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.TotalChallengesPrevDay != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.TotalChallengesPrevDay))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.ChallengerLastBlockChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.ChallengerLastBlockChallenges))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.RunnerLastBlockChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.RunnerLastBlockChallenges))
		i--
		dAtA[i] = 0x78
	}
	if m.V2NBXLastBlockChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2NBXLastBlockChallenges))
		i--
		dAtA[i] = 0x70
	}
	if m.V2VBXLastBlockChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2VBXLastBlockChallenges))
		i--
		dAtA[i] = 0x68
	}
	if m.V2VRXLastBlockChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2VRXLastBlockChallenges))
		i--
		dAtA[i] = 0x60
	}
	if m.ChallengerTotalChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.ChallengerTotalChallenges))
		i--
		dAtA[i] = 0x58
	}
	if m.RunnerTotalChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.RunnerTotalChallenges))
		i--
		dAtA[i] = 0x50
	}
	if m.V2NBXtotalChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2NBXtotalChallenges))
		i--
		dAtA[i] = 0x48
	}
	if m.V2VBXtotalChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2VBXtotalChallenges))
		i--
		dAtA[i] = 0x40
	}
	if m.V2VRXtotalChallenges != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.V2VRXtotalChallenges))
		i--
		dAtA[i] = 0x38
	}
	if len(m.EpochChallenger) > 0 {
		i -= len(m.EpochChallenger)
		copy(dAtA[i:], m.EpochChallenger)
		i = encodeVarintEpochData(dAtA, i, uint64(len(m.EpochChallenger)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.EpochRunner) > 0 {
		i -= len(m.EpochRunner)
		copy(dAtA[i:], m.EpochRunner)
		i = encodeVarintEpochData(dAtA, i, uint64(len(m.EpochRunner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EpochV2NBX) > 0 {
		i -= len(m.EpochV2NBX)
		copy(dAtA[i:], m.EpochV2NBX)
		i = encodeVarintEpochData(dAtA, i, uint64(len(m.EpochV2NBX)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EpochV2VBX) > 0 {
		i -= len(m.EpochV2VBX)
		copy(dAtA[i:], m.EpochV2VBX)
		i = encodeVarintEpochData(dAtA, i, uint64(len(m.EpochV2VBX)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EpochV2VRX) > 0 {
		i -= len(m.EpochV2VRX)
		copy(dAtA[i:], m.EpochV2VRX)
		i = encodeVarintEpochData(dAtA, i, uint64(len(m.EpochV2VRX)))
		i--
		dAtA[i] = 0x12
	}
	if m.TotalEpochs != 0 {
		i = encodeVarintEpochData(dAtA, i, uint64(m.TotalEpochs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEpochData(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpochData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalEpochs != 0 {
		n += 1 + sovEpochData(uint64(m.TotalEpochs))
	}
	l = len(m.EpochV2VRX)
	if l > 0 {
		n += 1 + l + sovEpochData(uint64(l))
	}
	l = len(m.EpochV2VBX)
	if l > 0 {
		n += 1 + l + sovEpochData(uint64(l))
	}
	l = len(m.EpochV2NBX)
	if l > 0 {
		n += 1 + l + sovEpochData(uint64(l))
	}
	l = len(m.EpochRunner)
	if l > 0 {
		n += 1 + l + sovEpochData(uint64(l))
	}
	l = len(m.EpochChallenger)
	if l > 0 {
		n += 1 + l + sovEpochData(uint64(l))
	}
	if m.V2VRXtotalChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.V2VRXtotalChallenges))
	}
	if m.V2VBXtotalChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.V2VBXtotalChallenges))
	}
	if m.V2NBXtotalChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.V2NBXtotalChallenges))
	}
	if m.RunnerTotalChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.RunnerTotalChallenges))
	}
	if m.ChallengerTotalChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.ChallengerTotalChallenges))
	}
	if m.V2VRXLastBlockChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.V2VRXLastBlockChallenges))
	}
	if m.V2VBXLastBlockChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.V2VBXLastBlockChallenges))
	}
	if m.V2NBXLastBlockChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.V2NBXLastBlockChallenges))
	}
	if m.RunnerLastBlockChallenges != 0 {
		n += 1 + sovEpochData(uint64(m.RunnerLastBlockChallenges))
	}
	if m.ChallengerLastBlockChallenges != 0 {
		n += 2 + sovEpochData(uint64(m.ChallengerLastBlockChallenges))
	}
	if m.TotalChallengesPrevDay != 0 {
		n += 2 + sovEpochData(uint64(m.TotalChallengesPrevDay))
	}
	if m.InitialPerChallengeValue != 0 {
		n += 2 + sovEpochData(uint64(m.InitialPerChallengeValue))
	}
	if m.V2NBXPerChallengeValue != 0 {
		n += 2 + sovEpochData(uint64(m.V2NBXPerChallengeValue))
	}
	if m.RunnerPerChallengeValue != 0 {
		n += 2 + sovEpochData(uint64(m.RunnerPerChallengeValue))
	}
	if m.ChallengerPerChallengeValue != 0 {
		n += 2 + sovEpochData(uint64(m.ChallengerPerChallengeValue))
	}
	if m.V2VBXPerChallengeValue != 0 {
		n += 2 + sovEpochData(uint64(m.V2VBXPerChallengeValue))
	}
	if m.V2VRXPerChallengeValue != 0 {
		n += 2 + sovEpochData(uint64(m.V2VRXPerChallengeValue))
	}
	return n
}

func sovEpochData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpochData(x uint64) (n int) {
	return sovEpochData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpochData
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
			return fmt.Errorf("proto: EpochData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalEpochs", wireType)
			}
			m.TotalEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochV2VRX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
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
				return ErrInvalidLengthEpochData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochV2VRX = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochV2VBX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
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
				return ErrInvalidLengthEpochData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochV2VBX = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochV2NBX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
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
				return ErrInvalidLengthEpochData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochV2NBX = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochRunner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
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
				return ErrInvalidLengthEpochData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochRunner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochChallenger", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
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
				return ErrInvalidLengthEpochData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochChallenger = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2VRXtotalChallenges", wireType)
			}
			m.V2VRXtotalChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2VRXtotalChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2VBXtotalChallenges", wireType)
			}
			m.V2VBXtotalChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2VBXtotalChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2NBXtotalChallenges", wireType)
			}
			m.V2NBXtotalChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2NBXtotalChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunnerTotalChallenges", wireType)
			}
			m.RunnerTotalChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RunnerTotalChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerTotalChallenges", wireType)
			}
			m.ChallengerTotalChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengerTotalChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2VRXLastBlockChallenges", wireType)
			}
			m.V2VRXLastBlockChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2VRXLastBlockChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2VBXLastBlockChallenges", wireType)
			}
			m.V2VBXLastBlockChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2VBXLastBlockChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2NBXLastBlockChallenges", wireType)
			}
			m.V2NBXLastBlockChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2NBXLastBlockChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunnerLastBlockChallenges", wireType)
			}
			m.RunnerLastBlockChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RunnerLastBlockChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerLastBlockChallenges", wireType)
			}
			m.ChallengerLastBlockChallenges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengerLastBlockChallenges |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalChallengesPrevDay", wireType)
			}
			m.TotalChallengesPrevDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalChallengesPrevDay |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPerChallengeValue", wireType)
			}
			m.InitialPerChallengeValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialPerChallengeValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2NBXPerChallengeValue", wireType)
			}
			m.V2NBXPerChallengeValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2NBXPerChallengeValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunnerPerChallengeValue", wireType)
			}
			m.RunnerPerChallengeValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RunnerPerChallengeValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerPerChallengeValue", wireType)
			}
			m.ChallengerPerChallengeValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengerPerChallengeValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2VBXPerChallengeValue", wireType)
			}
			m.V2VBXPerChallengeValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2VBXPerChallengeValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2VRXPerChallengeValue", wireType)
			}
			m.V2VRXPerChallengeValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.V2VRXPerChallengeValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpochData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpochData
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
func skipEpochData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpochData
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
					return 0, ErrIntOverflowEpochData
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
					return 0, ErrIntOverflowEpochData
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
				return 0, ErrInvalidLengthEpochData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpochData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpochData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpochData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpochData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpochData = fmt.Errorf("proto: unexpected end of group")
)
