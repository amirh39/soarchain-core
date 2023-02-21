// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/genesis.proto

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

// GenesisState defines the poa module's genesis state.
type GenesisState struct {
	Params          Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ClientList      []Client      `protobuf:"bytes,2,rep,name=clientList,proto3" json:"clientList"`
	ChallengerList  []Challenger  `protobuf:"bytes,3,rep,name=challengerList,proto3" json:"challengerList"`
	RunnerList      []Runner      `protobuf:"bytes,4,rep,name=runnerList,proto3" json:"runnerList"`
	GuardList       []Guard       `protobuf:"bytes,5,rep,name=guardList,proto3" json:"guardList"`
	VrfDataList     []VrfData     `protobuf:"bytes,6,rep,name=vrfDataList,proto3" json:"vrfDataList"`
	VrfUserList     []VrfUser     `protobuf:"bytes,7,rep,name=vrfUserList,proto3" json:"vrfUserList"`
	EpochData       EpochData     `protobuf:"bytes,8,opt,name=epochData,proto3" json:"epochData"`
	MotusWalletList []MotusWallet `protobuf:"bytes,9,rep,name=motusWalletList,proto3" json:"motusWalletList"`
	MasterKey       MasterKey     `protobuf:"bytes,10,opt,name=masterKey,proto3" json:"masterKey"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c9d5a9eb5268c7b, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClientList() []Client {
	if m != nil {
		return m.ClientList
	}
	return nil
}

func (m *GenesisState) GetChallengerList() []Challenger {
	if m != nil {
		return m.ChallengerList
	}
	return nil
}

func (m *GenesisState) GetRunnerList() []Runner {
	if m != nil {
		return m.RunnerList
	}
	return nil
}

func (m *GenesisState) GetGuardList() []Guard {
	if m != nil {
		return m.GuardList
	}
	return nil
}

func (m *GenesisState) GetVrfDataList() []VrfData {
	if m != nil {
		return m.VrfDataList
	}
	return nil
}

func (m *GenesisState) GetVrfUserList() []VrfUser {
	if m != nil {
		return m.VrfUserList
	}
	return nil
}

func (m *GenesisState) GetEpochData() EpochData {
	if m != nil {
		return m.EpochData
	}
	return EpochData{}
}

func (m *GenesisState) GetMotusWalletList() []MotusWallet {
	if m != nil {
		return m.MotusWalletList
	}
	return nil
}

func (m *GenesisState) GetMasterKey() MasterKey {
	if m != nil {
		return m.MasterKey
	}
	return MasterKey{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "soarchain.poa.GenesisState")
}

func init() { proto.RegisterFile("poa/genesis.proto", fileDescriptor_5c9d5a9eb5268c7b) }

var fileDescriptor_5c9d5a9eb5268c7b = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x0f, 0xd2, 0x30,
	0x18, 0xc6, 0x37, 0xf9, 0xa3, 0x14, 0x15, 0x6d, 0x80, 0x4c, 0x0e, 0x93, 0x78, 0xf2, 0xb4, 0x25,
	0x72, 0x31, 0xd1, 0x78, 0x40, 0x0d, 0x89, 0x4a, 0x62, 0x30, 0x6a, 0xe2, 0x85, 0x54, 0x28, 0x63,
	0x11, 0xd6, 0xa5, 0xed, 0x50, 0xbe, 0x81, 0x47, 0x3f, 0x16, 0x47, 0x8e, 0x9e, 0x8c, 0x81, 0x2f,
	0x62, 0xfa, 0xb6, 0xfb, 0x43, 0x83, 0xb7, 0xe5, 0x79, 0x9f, 0xdf, 0xb3, 0xa7, 0xed, 0x8b, 0xee,
	0xa7, 0x8c, 0x84, 0x11, 0x4d, 0xa8, 0x88, 0x45, 0x90, 0x72, 0x26, 0x19, 0xbe, 0x23, 0x18, 0xe1,
	0x8b, 0x35, 0x89, 0x93, 0x20, 0x65, 0x64, 0xd0, 0x8d, 0x58, 0xc4, 0x60, 0x12, 0xaa, 0x2f, 0x6d,
	0x1a, 0xdc, 0x53, 0x5c, 0x4a, 0x38, 0xd9, 0x8a, 0xaa, 0xb2, 0xd8, 0xc4, 0x34, 0x91, 0x46, 0xe9,
	0x82, 0xb2, 0x26, 0x9b, 0x0d, 0x4d, 0x22, 0xca, 0xab, 0x3e, 0x9e, 0x25, 0x49, 0xa1, 0x74, 0xa0,
	0x43, 0x46, 0xf8, 0xd2, 0x08, 0x58, 0x09, 0x3b, 0xbe, 0x9a, 0x2f, 0x89, 0x24, 0xb6, 0x96, 0x89,
	0x02, 0x84, 0x1f, 0xd0, 0x94, 0x2d, 0xd6, 0x55, 0x67, 0x5f, 0xa9, 0x5b, 0x26, 0x33, 0x31, 0xff,
	0xae, 0xfe, 0x7d, 0x51, 0x67, 0x4b, 0x84, 0xa4, 0x7c, 0xfe, 0x8d, 0xee, 0xb5, 0xfa, 0xe8, 0x67,
	0x03, 0xdd, 0x9e, 0xe8, 0xf3, 0x7f, 0x90, 0x44, 0x52, 0x3c, 0x42, 0x4d, 0x7d, 0x2e, 0xcf, 0x1d,
	0xba, 0x8f, 0xdb, 0x4f, 0x7a, 0xc1, 0xc5, 0x7d, 0x04, 0xef, 0x61, 0x38, 0xae, 0x1f, 0xfe, 0x3c,
	0x74, 0x66, 0xc6, 0x8a, 0x9f, 0x21, 0xa4, 0x8f, 0xfe, 0x2e, 0x16, 0xd2, 0xbb, 0x31, 0xac, 0x5d,
	0x01, 0x5f, 0x82, 0xc1, 0x80, 0x15, 0x3b, 0x9e, 0xa0, 0xbb, 0xe5, 0x2d, 0x41, 0x40, 0x0d, 0x02,
	0x1e, 0xd8, 0x01, 0x85, 0xc9, 0x84, 0x58, 0x98, 0x6a, 0xa1, 0x2f, 0x16, 0x42, 0xea, 0x57, 0x5b,
	0xcc, 0xc0, 0x90, 0xb7, 0x28, 0xed, 0xf8, 0x29, 0x6a, 0xc1, 0x1b, 0x00, 0xdb, 0x00, 0xb6, 0x6b,
	0xb1, 0x13, 0x35, 0x37, 0x68, 0x69, 0xc6, 0x2f, 0x50, 0x7b, 0xc7, 0x57, 0xaf, 0x88, 0x24, 0xc0,
	0x36, 0x81, 0xed, 0x5b, 0xec, 0x27, 0xed, 0x30, 0x74, 0x15, 0x30, 0xfc, 0x47, 0x61, 0x7a, 0xdf,
	0xfc, 0x1f, 0xaf, 0x1c, 0x15, 0x3e, 0x07, 0xf0, 0x73, 0xd4, 0x82, 0x25, 0x50, 0x81, 0xde, 0x2d,
	0x78, 0x34, 0xcf, 0xa2, 0x5f, 0xe7, 0xf3, 0xbc, 0x7d, 0x01, 0xe0, 0x37, 0xa8, 0x03, 0xcb, 0xf2,
	0x19, 0x76, 0x05, 0x1a, 0xb4, 0xa0, 0xc1, 0xc0, 0xca, 0x98, 0x96, 0x2e, 0x93, 0x62, 0x83, 0xaa,
	0x89, 0x5e, 0xb0, 0xb7, 0x74, 0xef, 0xa1, 0xab, 0x4d, 0xa6, 0xf9, 0x3c, 0x6f, 0x52, 0x00, 0xe3,
	0xf0, 0x70, 0xf2, 0xdd, 0xe3, 0xc9, 0x77, 0xff, 0x9e, 0x7c, 0xf7, 0xd7, 0xd9, 0x77, 0x8e, 0x67,
	0xdf, 0xf9, 0x7d, 0xf6, 0x9d, 0x2f, 0xbd, 0x22, 0x23, 0xfc, 0x11, 0xaa, 0x35, 0x96, 0xfb, 0x94,
	0x8a, 0xaf, 0x4d, 0x58, 0xe1, 0xd1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xea, 0x09, 0x72,
	0xc5, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MasterKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.MotusWalletList) > 0 {
		for iNdEx := len(m.MotusWalletList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MotusWalletList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	{
		size, err := m.EpochData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.VrfUserList) > 0 {
		for iNdEx := len(m.VrfUserList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VrfUserList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.VrfDataList) > 0 {
		for iNdEx := len(m.VrfDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VrfDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.GuardList) > 0 {
		for iNdEx := len(m.GuardList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GuardList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.RunnerList) > 0 {
		for iNdEx := len(m.RunnerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RunnerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ChallengerList) > 0 {
		for iNdEx := len(m.ChallengerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChallengerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClientList) > 0 {
		for iNdEx := len(m.ClientList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClientList) > 0 {
		for _, e := range m.ClientList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChallengerList) > 0 {
		for _, e := range m.ChallengerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RunnerList) > 0 {
		for _, e := range m.RunnerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GuardList) > 0 {
		for _, e := range m.GuardList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VrfDataList) > 0 {
		for _, e := range m.VrfDataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VrfUserList) > 0 {
		for _, e := range m.VrfUserList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.EpochData.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.MotusWalletList) > 0 {
		for _, e := range m.MotusWalletList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.MasterKey.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientList = append(m.ClientList, Client{})
			if err := m.ClientList[len(m.ClientList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengerList = append(m.ChallengerList, Challenger{})
			if err := m.ChallengerList[len(m.ChallengerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunnerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunnerList = append(m.RunnerList, Runner{})
			if err := m.RunnerList[len(m.RunnerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardList = append(m.GuardList, Guard{})
			if err := m.GuardList[len(m.GuardList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VrfDataList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VrfDataList = append(m.VrfDataList, VrfData{})
			if err := m.VrfDataList[len(m.VrfDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VrfUserList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VrfUserList = append(m.VrfUserList, VrfUser{})
			if err := m.VrfUserList[len(m.VrfUserList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EpochData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MotusWalletList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MotusWalletList = append(m.MotusWalletList, MotusWallet{})
			if err := m.MotusWalletList[len(m.MotusWalletList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MasterKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
