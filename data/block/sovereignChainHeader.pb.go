// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sovereignChainHeader.proto

package block

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// SovereignChainHeader extends the Header structure with extra fields needed by sovereign chain
type SovereignChainHeader struct {
	Header                    *Header             `protobuf:"bytes,1,opt,name=Header,proto3" json:"header"`
	ValidatorStatsRootHash    []byte              `protobuf:"bytes,2,opt,name=ValidatorStatsRootHash,proto3" json:"validatorStatsRootHash"`
	ExtendedShardHeaderHashes [][]byte            `protobuf:"bytes,3,rep,name=ExtendedShardHeaderHashes,proto3" json:"extendedShardHeaderHashes,omitempty"`
	OutGoingOperations        *OutGoingOperations `protobuf:"bytes,4,opt,name=OutGoingOperations,proto3" json:"outGoingOperations,omitempty"`
}

func (m *SovereignChainHeader) Reset()      { *m = SovereignChainHeader{} }
func (*SovereignChainHeader) ProtoMessage() {}
func (*SovereignChainHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b8ff297a820152, []int{0}
}
func (m *SovereignChainHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SovereignChainHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SovereignChainHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SovereignChainHeader.Merge(m, src)
}
func (m *SovereignChainHeader) XXX_Size() int {
	return m.Size()
}
func (m *SovereignChainHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SovereignChainHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SovereignChainHeader proto.InternalMessageInfo

func (m *SovereignChainHeader) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SovereignChainHeader) GetValidatorStatsRootHash() []byte {
	if m != nil {
		return m.ValidatorStatsRootHash
	}
	return nil
}

func (m *SovereignChainHeader) GetExtendedShardHeaderHashes() [][]byte {
	if m != nil {
		return m.ExtendedShardHeaderHashes
	}
	return nil
}

func (m *SovereignChainHeader) GetOutGoingOperations() *OutGoingOperations {
	if m != nil {
		return m.OutGoingOperations
	}
	return nil
}

type OutGoingOperations struct {
	OutGoingOperationHashes               [][]byte `protobuf:"bytes,1,rep,name=OutGoingOperationHashes,proto3" json:"outGoingOperationHashes,omitempty"`
	OutGoingOperationsHash                []byte   `protobuf:"bytes,2,opt,name=OutGoingOperationsHash,proto3" json:"outGoingOperationsHash,omitempty"`
	AggregatedSignatureOutGoingOperations []byte   `protobuf:"bytes,3,opt,name=AggregatedSignatureOutGoingOperations,proto3" json:"aggregatedSignatureOutGoingOperations,omitempty"`
	LeaderSignatureOutGoingOperations     []byte   `protobuf:"bytes,4,opt,name=LeaderSignatureOutGoingOperations,proto3" json:"leaderSignatureOutGoingOperations,omitempty"`
}

func (m *OutGoingOperations) Reset()      { *m = OutGoingOperations{} }
func (*OutGoingOperations) ProtoMessage() {}
func (*OutGoingOperations) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b8ff297a820152, []int{1}
}
func (m *OutGoingOperations) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutGoingOperations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *OutGoingOperations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutGoingOperations.Merge(m, src)
}
func (m *OutGoingOperations) XXX_Size() int {
	return m.Size()
}
func (m *OutGoingOperations) XXX_DiscardUnknown() {
	xxx_messageInfo_OutGoingOperations.DiscardUnknown(m)
}

var xxx_messageInfo_OutGoingOperations proto.InternalMessageInfo

func (m *OutGoingOperations) GetOutGoingOperationHashes() [][]byte {
	if m != nil {
		return m.OutGoingOperationHashes
	}
	return nil
}

func (m *OutGoingOperations) GetOutGoingOperationsHash() []byte {
	if m != nil {
		return m.OutGoingOperationsHash
	}
	return nil
}

func (m *OutGoingOperations) GetAggregatedSignatureOutGoingOperations() []byte {
	if m != nil {
		return m.AggregatedSignatureOutGoingOperations
	}
	return nil
}

func (m *OutGoingOperations) GetLeaderSignatureOutGoingOperations() []byte {
	if m != nil {
		return m.LeaderSignatureOutGoingOperations
	}
	return nil
}

func init() {
	proto.RegisterType((*SovereignChainHeader)(nil), "proto.SovereignChainHeader")
	proto.RegisterType((*OutGoingOperations)(nil), "proto.OutGoingOperations")
}

func init() { proto.RegisterFile("sovereignChainHeader.proto", fileDescriptor_b9b8ff297a820152) }

var fileDescriptor_b9b8ff297a820152 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0x75, 0xeb, 0xc1, 0x1b, 0x17, 0x0b, 0x95, 0xae, 0x42, 0x4e, 0x36, 0x98, 0xa8,
	0x04, 0xb4, 0x82, 0x7d, 0x00, 0x44, 0x10, 0x62, 0x07, 0xa4, 0x49, 0xa9, 0xc4, 0x01, 0x21, 0x21,
	0x77, 0x31, 0x4e, 0xb4, 0x36, 0xaf, 0x72, 0xdc, 0x09, 0x0e, 0x48, 0x5c, 0xb9, 0xf1, 0x31, 0xf8,
	0x28, 0x1c, 0x7b, 0xec, 0xc9, 0xa2, 0xee, 0x05, 0xf9, 0x34, 0xbe, 0x01, 0xc2, 0xe9, 0x21, 0x52,
	0x92, 0xad, 0xa7, 0xbc, 0xf7, 0xf2, 0x7f, 0xff, 0xdf, 0x7b, 0x71, 0x8c, 0x7b, 0x39, 0x5c, 0x71,
	0xc9, 0x53, 0x91, 0xbd, 0x4a, 0x58, 0x9a, 0x9d, 0x71, 0x16, 0x73, 0x39, 0x98, 0x49, 0x50, 0x40,
	0xf6, 0xdc, 0xa3, 0xf7, 0x54, 0xa4, 0x2a, 0x99, 0x8f, 0x07, 0x17, 0x30, 0x1d, 0x0a, 0x10, 0x30,
	0x74, 0xe5, 0xf1, 0xfc, 0x93, 0xcb, 0x5c, 0xe2, 0xa2, 0xa2, 0xab, 0xb7, 0x3f, 0x9e, 0xc0, 0xc5,
	0x65, 0x91, 0x1c, 0xff, 0xdd, 0xc1, 0x77, 0x47, 0x35, 0x04, 0xf2, 0x0c, 0xb7, 0x8b, 0xa8, 0x8b,
	0x02, 0xd4, 0xdf, 0x7f, 0x7e, 0xa7, 0x68, 0x18, 0x14, 0xc5, 0x10, 0x5b, 0xed, 0xb7, 0x13, 0x17,
	0x47, 0x1b, 0x21, 0x89, 0x70, 0xe7, 0x1d, 0x9b, 0xa4, 0x31, 0x53, 0x20, 0x47, 0x8a, 0xa9, 0x3c,
	0x02, 0x50, 0x67, 0x2c, 0x4f, 0xba, 0x3b, 0x01, 0xea, 0x1f, 0x84, 0x3d, 0xab, 0xfd, 0xce, 0x55,
	0xad, 0x22, 0x6a, 0xe8, 0x24, 0x1c, 0x1f, 0xbe, 0xfe, 0xac, 0x78, 0x16, 0xf3, 0x78, 0x94, 0x30,
	0x19, 0x17, 0xa8, 0xff, 0xaf, 0x78, 0xde, 0x6d, 0x05, 0xad, 0xfe, 0x41, 0xf8, 0xc8, 0x6a, 0xff,
	0x01, 0x6f, 0x12, 0x3d, 0x81, 0x69, 0xaa, 0xf8, 0x74, 0xa6, 0xbe, 0x44, 0xcd, 0x4e, 0xe4, 0x12,
	0x93, 0xf3, 0xb9, 0x7a, 0x03, 0x69, 0x26, 0xce, 0x67, 0x5c, 0x32, 0x95, 0x42, 0x96, 0x77, 0x77,
	0xdd, 0xe6, 0x87, 0x9b, 0xcd, 0xab, 0x82, 0x30, 0xb0, 0xda, 0xbf, 0x0f, 0x95, 0x7a, 0x89, 0x59,
	0x63, 0x7b, 0xac, 0x5b, 0x75, 0x34, 0xf2, 0x11, 0xdf, 0xab, 0x54, 0x37, 0x8b, 0x22, 0xb7, 0xe8,
	0x89, 0xd5, 0xfe, 0x11, 0xd4, 0x4b, 0x4a, 0xc8, 0x26, 0x17, 0xf2, 0x01, 0x77, 0xaa, 0xd8, 0xd2,
	0xf9, 0x3c, 0xb4, 0xda, 0x0f, 0xa0, 0x56, 0x51, 0xb2, 0x6f, 0xf0, 0x20, 0xdf, 0x11, 0x3e, 0x79,
	0x29, 0x84, 0xe4, 0x82, 0x29, 0x1e, 0x8f, 0x52, 0x91, 0x31, 0x35, 0x97, 0xbc, 0xe6, 0xb3, 0xb6,
	0x1c, 0xed, 0xd4, 0x6a, 0x7f, 0xc8, 0xb6, 0x69, 0x28, 0xc1, 0xb7, 0x23, 0x90, 0xaf, 0xf8, 0xe8,
	0xad, 0x3b, 0xde, 0x9b, 0xc6, 0xd8, 0x75, 0x63, 0x0c, 0xad, 0xf6, 0x1f, 0x4f, 0x6e, 0x13, 0x97,
	0x46, 0xb8, 0xdd, 0x39, 0x7c, 0xb1, 0x58, 0x51, 0x6f, 0xb9, 0xa2, 0xde, 0xf5, 0x8a, 0xa2, 0x6f,
	0x86, 0xa2, 0x9f, 0x86, 0xa2, 0x5f, 0x86, 0xa2, 0x85, 0xa1, 0x68, 0x69, 0x28, 0xfa, 0x6d, 0x28,
	0xfa, 0x63, 0xa8, 0x77, 0x6d, 0x28, 0xfa, 0xb1, 0xa6, 0xde, 0x62, 0x4d, 0xbd, 0xe5, 0x9a, 0x7a,
	0xef, 0xf7, 0xdc, 0xdd, 0x1c, 0xb7, 0xdd, 0x1f, 0x77, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x74,
	0x9d, 0xd1, 0xa5, 0xfd, 0x03, 0x00, 0x00,
}

func (this *SovereignChainHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SovereignChainHeader)
	if !ok {
		that2, ok := that.(SovereignChainHeader)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Header.Equal(that1.Header) {
		return false
	}
	if !bytes.Equal(this.ValidatorStatsRootHash, that1.ValidatorStatsRootHash) {
		return false
	}
	if len(this.ExtendedShardHeaderHashes) != len(that1.ExtendedShardHeaderHashes) {
		return false
	}
	for i := range this.ExtendedShardHeaderHashes {
		if !bytes.Equal(this.ExtendedShardHeaderHashes[i], that1.ExtendedShardHeaderHashes[i]) {
			return false
		}
	}
	if !this.OutGoingOperations.Equal(that1.OutGoingOperations) {
		return false
	}
	return true
}
func (this *OutGoingOperations) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OutGoingOperations)
	if !ok {
		that2, ok := that.(OutGoingOperations)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.OutGoingOperationHashes) != len(that1.OutGoingOperationHashes) {
		return false
	}
	for i := range this.OutGoingOperationHashes {
		if !bytes.Equal(this.OutGoingOperationHashes[i], that1.OutGoingOperationHashes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.OutGoingOperationsHash, that1.OutGoingOperationsHash) {
		return false
	}
	if !bytes.Equal(this.AggregatedSignatureOutGoingOperations, that1.AggregatedSignatureOutGoingOperations) {
		return false
	}
	if !bytes.Equal(this.LeaderSignatureOutGoingOperations, that1.LeaderSignatureOutGoingOperations) {
		return false
	}
	return true
}
func (this *SovereignChainHeader) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&block.SovereignChainHeader{")
	if this.Header != nil {
		s = append(s, "Header: "+fmt.Sprintf("%#v", this.Header)+",\n")
	}
	s = append(s, "ValidatorStatsRootHash: "+fmt.Sprintf("%#v", this.ValidatorStatsRootHash)+",\n")
	s = append(s, "ExtendedShardHeaderHashes: "+fmt.Sprintf("%#v", this.ExtendedShardHeaderHashes)+",\n")
	if this.OutGoingOperations != nil {
		s = append(s, "OutGoingOperations: "+fmt.Sprintf("%#v", this.OutGoingOperations)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *OutGoingOperations) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&block.OutGoingOperations{")
	s = append(s, "OutGoingOperationHashes: "+fmt.Sprintf("%#v", this.OutGoingOperationHashes)+",\n")
	s = append(s, "OutGoingOperationsHash: "+fmt.Sprintf("%#v", this.OutGoingOperationsHash)+",\n")
	s = append(s, "AggregatedSignatureOutGoingOperations: "+fmt.Sprintf("%#v", this.AggregatedSignatureOutGoingOperations)+",\n")
	s = append(s, "LeaderSignatureOutGoingOperations: "+fmt.Sprintf("%#v", this.LeaderSignatureOutGoingOperations)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSovereignChainHeader(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SovereignChainHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SovereignChainHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SovereignChainHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OutGoingOperations != nil {
		{
			size, err := m.OutGoingOperations.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSovereignChainHeader(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.ExtendedShardHeaderHashes) > 0 {
		for iNdEx := len(m.ExtendedShardHeaderHashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ExtendedShardHeaderHashes[iNdEx])
			copy(dAtA[i:], m.ExtendedShardHeaderHashes[iNdEx])
			i = encodeVarintSovereignChainHeader(dAtA, i, uint64(len(m.ExtendedShardHeaderHashes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ValidatorStatsRootHash) > 0 {
		i -= len(m.ValidatorStatsRootHash)
		copy(dAtA[i:], m.ValidatorStatsRootHash)
		i = encodeVarintSovereignChainHeader(dAtA, i, uint64(len(m.ValidatorStatsRootHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSovereignChainHeader(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OutGoingOperations) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutGoingOperations) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutGoingOperations) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LeaderSignatureOutGoingOperations) > 0 {
		i -= len(m.LeaderSignatureOutGoingOperations)
		copy(dAtA[i:], m.LeaderSignatureOutGoingOperations)
		i = encodeVarintSovereignChainHeader(dAtA, i, uint64(len(m.LeaderSignatureOutGoingOperations)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AggregatedSignatureOutGoingOperations) > 0 {
		i -= len(m.AggregatedSignatureOutGoingOperations)
		copy(dAtA[i:], m.AggregatedSignatureOutGoingOperations)
		i = encodeVarintSovereignChainHeader(dAtA, i, uint64(len(m.AggregatedSignatureOutGoingOperations)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OutGoingOperationsHash) > 0 {
		i -= len(m.OutGoingOperationsHash)
		copy(dAtA[i:], m.OutGoingOperationsHash)
		i = encodeVarintSovereignChainHeader(dAtA, i, uint64(len(m.OutGoingOperationsHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OutGoingOperationHashes) > 0 {
		for iNdEx := len(m.OutGoingOperationHashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OutGoingOperationHashes[iNdEx])
			copy(dAtA[i:], m.OutGoingOperationHashes[iNdEx])
			i = encodeVarintSovereignChainHeader(dAtA, i, uint64(len(m.OutGoingOperationHashes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSovereignChainHeader(dAtA []byte, offset int, v uint64) int {
	offset -= sovSovereignChainHeader(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SovereignChainHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovSovereignChainHeader(uint64(l))
	}
	l = len(m.ValidatorStatsRootHash)
	if l > 0 {
		n += 1 + l + sovSovereignChainHeader(uint64(l))
	}
	if len(m.ExtendedShardHeaderHashes) > 0 {
		for _, b := range m.ExtendedShardHeaderHashes {
			l = len(b)
			n += 1 + l + sovSovereignChainHeader(uint64(l))
		}
	}
	if m.OutGoingOperations != nil {
		l = m.OutGoingOperations.Size()
		n += 1 + l + sovSovereignChainHeader(uint64(l))
	}
	return n
}

func (m *OutGoingOperations) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.OutGoingOperationHashes) > 0 {
		for _, b := range m.OutGoingOperationHashes {
			l = len(b)
			n += 1 + l + sovSovereignChainHeader(uint64(l))
		}
	}
	l = len(m.OutGoingOperationsHash)
	if l > 0 {
		n += 1 + l + sovSovereignChainHeader(uint64(l))
	}
	l = len(m.AggregatedSignatureOutGoingOperations)
	if l > 0 {
		n += 1 + l + sovSovereignChainHeader(uint64(l))
	}
	l = len(m.LeaderSignatureOutGoingOperations)
	if l > 0 {
		n += 1 + l + sovSovereignChainHeader(uint64(l))
	}
	return n
}

func sovSovereignChainHeader(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSovereignChainHeader(x uint64) (n int) {
	return sovSovereignChainHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SovereignChainHeader) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SovereignChainHeader{`,
		`Header:` + strings.Replace(fmt.Sprintf("%v", this.Header), "Header", "Header", 1) + `,`,
		`ValidatorStatsRootHash:` + fmt.Sprintf("%v", this.ValidatorStatsRootHash) + `,`,
		`ExtendedShardHeaderHashes:` + fmt.Sprintf("%v", this.ExtendedShardHeaderHashes) + `,`,
		`OutGoingOperations:` + strings.Replace(this.OutGoingOperations.String(), "OutGoingOperations", "OutGoingOperations", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *OutGoingOperations) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&OutGoingOperations{`,
		`OutGoingOperationHashes:` + fmt.Sprintf("%v", this.OutGoingOperationHashes) + `,`,
		`OutGoingOperationsHash:` + fmt.Sprintf("%v", this.OutGoingOperationsHash) + `,`,
		`AggregatedSignatureOutGoingOperations:` + fmt.Sprintf("%v", this.AggregatedSignatureOutGoingOperations) + `,`,
		`LeaderSignatureOutGoingOperations:` + fmt.Sprintf("%v", this.LeaderSignatureOutGoingOperations) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSovereignChainHeader(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SovereignChainHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSovereignChainHeader
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
			return fmt.Errorf("proto: SovereignChainHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SovereignChainHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSovereignChainHeader
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
				return ErrInvalidLengthSovereignChainHeader
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorStatsRootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSovereignChainHeader
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
				return ErrInvalidLengthSovereignChainHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorStatsRootHash = append(m.ValidatorStatsRootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorStatsRootHash == nil {
				m.ValidatorStatsRootHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtendedShardHeaderHashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSovereignChainHeader
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
				return ErrInvalidLengthSovereignChainHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtendedShardHeaderHashes = append(m.ExtendedShardHeaderHashes, make([]byte, postIndex-iNdEx))
			copy(m.ExtendedShardHeaderHashes[len(m.ExtendedShardHeaderHashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutGoingOperations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSovereignChainHeader
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
				return ErrInvalidLengthSovereignChainHeader
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OutGoingOperations == nil {
				m.OutGoingOperations = &OutGoingOperations{}
			}
			if err := m.OutGoingOperations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSovereignChainHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSovereignChainHeader
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
func (m *OutGoingOperations) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSovereignChainHeader
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
			return fmt.Errorf("proto: OutGoingOperations: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutGoingOperations: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutGoingOperationHashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSovereignChainHeader
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
				return ErrInvalidLengthSovereignChainHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutGoingOperationHashes = append(m.OutGoingOperationHashes, make([]byte, postIndex-iNdEx))
			copy(m.OutGoingOperationHashes[len(m.OutGoingOperationHashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutGoingOperationsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSovereignChainHeader
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
				return ErrInvalidLengthSovereignChainHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutGoingOperationsHash = append(m.OutGoingOperationsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.OutGoingOperationsHash == nil {
				m.OutGoingOperationsHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregatedSignatureOutGoingOperations", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSovereignChainHeader
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
				return ErrInvalidLengthSovereignChainHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregatedSignatureOutGoingOperations = append(m.AggregatedSignatureOutGoingOperations[:0], dAtA[iNdEx:postIndex]...)
			if m.AggregatedSignatureOutGoingOperations == nil {
				m.AggregatedSignatureOutGoingOperations = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderSignatureOutGoingOperations", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSovereignChainHeader
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
				return ErrInvalidLengthSovereignChainHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeaderSignatureOutGoingOperations = append(m.LeaderSignatureOutGoingOperations[:0], dAtA[iNdEx:postIndex]...)
			if m.LeaderSignatureOutGoingOperations == nil {
				m.LeaderSignatureOutGoingOperations = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSovereignChainHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSovereignChainHeader
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSovereignChainHeader
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
func skipSovereignChainHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSovereignChainHeader
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
					return 0, ErrIntOverflowSovereignChainHeader
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
					return 0, ErrIntOverflowSovereignChainHeader
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
				return 0, ErrInvalidLengthSovereignChainHeader
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSovereignChainHeader
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSovereignChainHeader
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSovereignChainHeader        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSovereignChainHeader          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSovereignChainHeader = fmt.Errorf("proto: unexpected end of group")
)
