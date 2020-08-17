// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: miniblockMetadata.proto

package fullHistory

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

// MiniblockMetadata is used to store information about a history transaction
type MiniblockMetadata struct {
	RcvShardID  uint32 `protobuf:"varint,1,opt,name=RcvShardID,proto3" json:"RcvShardID,omitempty"`
	SndShardID  uint32 `protobuf:"varint,2,opt,name=SndShardID,proto3" json:"SndShardID,omitempty"`
	Round       uint64 `protobuf:"varint,3,opt,name=Round,proto3" json:"Round,omitempty"`
	HeaderNonce uint64 `protobuf:"varint,4,opt,name=HeaderNonce,proto3" json:"HeaderNonce,omitempty"`
	HeaderHash  []byte `protobuf:"bytes,5,opt,name=HeaderHash,proto3" json:"HeaderHash,omitempty"`
	MbHash      []byte `protobuf:"bytes,6,opt,name=MbHash,proto3" json:"MbHash,omitempty"`
	Status      []byte `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
	Epoch       uint32 `protobuf:"varint,8,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
}

func (m *MiniblockMetadata) Reset()      { *m = MiniblockMetadata{} }
func (*MiniblockMetadata) ProtoMessage() {}
func (*MiniblockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd82f29831cbb1fe, []int{0}
}
func (m *MiniblockMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MiniblockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MiniblockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiniblockMetadata.Merge(m, src)
}
func (m *MiniblockMetadata) XXX_Size() int {
	return m.Size()
}
func (m *MiniblockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MiniblockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MiniblockMetadata proto.InternalMessageInfo

func (m *MiniblockMetadata) GetRcvShardID() uint32 {
	if m != nil {
		return m.RcvShardID
	}
	return 0
}

func (m *MiniblockMetadata) GetSndShardID() uint32 {
	if m != nil {
		return m.SndShardID
	}
	return 0
}

func (m *MiniblockMetadata) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *MiniblockMetadata) GetHeaderNonce() uint64 {
	if m != nil {
		return m.HeaderNonce
	}
	return 0
}

func (m *MiniblockMetadata) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *MiniblockMetadata) GetMbHash() []byte {
	if m != nil {
		return m.MbHash
	}
	return nil
}

func (m *MiniblockMetadata) GetStatus() []byte {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MiniblockMetadata) GetEpoch() uint32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func init() {
	proto.RegisterType((*MiniblockMetadata)(nil), "proto.MiniblockMetadata")
}

func init() { proto.RegisterFile("miniblockMetadata.proto", fileDescriptor_cd82f29831cbb1fe) }

var fileDescriptor_cd82f29831cbb1fe = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0x80, 0xfd, 0xa0, 0x2d, 0xc8, 0x85, 0x81, 0x08, 0x81, 0xc5, 0xf0, 0x14, 0x31, 0x75, 0xa1,
	0x1d, 0xb8, 0x01, 0xa2, 0x52, 0x19, 0xc2, 0x90, 0x6c, 0x6c, 0xce, 0x4f, 0x93, 0x88, 0x34, 0xae,
	0x12, 0x07, 0x89, 0x8d, 0x23, 0x70, 0x0c, 0x8e, 0xc2, 0x98, 0x31, 0x23, 0x71, 0x16, 0xc6, 0x72,
	0x03, 0x14, 0x1b, 0x44, 0xc4, 0x64, 0x7f, 0xdf, 0x27, 0x3f, 0x3d, 0x99, 0x9e, 0x6f, 0xd2, 0x3c,
	0xf5, 0x33, 0x11, 0x3c, 0x3a, 0x91, 0xe4, 0x21, 0x97, 0x7c, 0xbe, 0x2d, 0x84, 0x14, 0xd6, 0x58,
	0x1f, 0x17, 0x57, 0x71, 0x2a, 0x93, 0xca, 0x9f, 0x07, 0x62, 0xb3, 0x88, 0x45, 0x2c, 0x16, 0x5a,
	0xfb, 0xd5, 0x5a, 0x93, 0x06, 0x7d, 0x33, 0xaf, 0x2e, 0xbf, 0x80, 0x9e, 0x38, 0xff, 0x27, 0x5a,
	0x48, 0xa9, 0x1b, 0x3c, 0x79, 0x09, 0x2f, 0xc2, 0xbb, 0x5b, 0x06, 0x36, 0xcc, 0x8e, 0xdd, 0x81,
	0xe9, 0xbb, 0x97, 0x87, 0xbf, 0x7d, 0xcf, 0xf4, 0x3f, 0x63, 0x9d, 0xd2, 0xb1, 0x2b, 0xaa, 0x3c,
	0x64, 0xfb, 0x36, 0xcc, 0x46, 0xae, 0x01, 0xcb, 0xa6, 0xd3, 0x55, 0xc4, 0xc3, 0xa8, 0xb8, 0x17,
	0x79, 0x10, 0xb1, 0x91, 0x6e, 0x43, 0xd5, 0xcf, 0x35, 0xb8, 0xe2, 0x65, 0xc2, 0xc6, 0x36, 0xcc,
	0x8e, 0xdc, 0x81, 0xb1, 0xce, 0xe8, 0xc4, 0xf1, 0x75, 0x9b, 0xe8, 0xf6, 0x43, 0xbd, 0xf7, 0x24,
	0x97, 0x55, 0xc9, 0x0e, 0x8c, 0x37, 0xd4, 0xef, 0xb1, 0xdc, 0x8a, 0x20, 0x61, 0x87, 0x7a, 0x45,
	0x03, 0x37, 0xcb, 0xba, 0x45, 0xd2, 0xb4, 0x48, 0x76, 0x2d, 0xc2, 0x8b, 0x42, 0x78, 0x53, 0x08,
	0xef, 0x0a, 0xa1, 0x56, 0x08, 0x8d, 0x42, 0xf8, 0x50, 0x08, 0x9f, 0x0a, 0xc9, 0x4e, 0x21, 0xbc,
	0x76, 0x48, 0xea, 0x0e, 0x49, 0xd3, 0x21, 0x79, 0x98, 0xae, 0xab, 0x2c, 0x5b, 0xa5, 0xa5, 0x14,
	0xc5, 0xb3, 0x3f, 0xd1, 0x3f, 0x78, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xb7, 0x4f, 0xea,
	0x92, 0x01, 0x00, 0x00,
}

func (this *MiniblockMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MiniblockMetadata)
	if !ok {
		that2, ok := that.(MiniblockMetadata)
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
	if this.RcvShardID != that1.RcvShardID {
		return false
	}
	if this.SndShardID != that1.SndShardID {
		return false
	}
	if this.Round != that1.Round {
		return false
	}
	if this.HeaderNonce != that1.HeaderNonce {
		return false
	}
	if !bytes.Equal(this.HeaderHash, that1.HeaderHash) {
		return false
	}
	if !bytes.Equal(this.MbHash, that1.MbHash) {
		return false
	}
	if !bytes.Equal(this.Status, that1.Status) {
		return false
	}
	if this.Epoch != that1.Epoch {
		return false
	}
	return true
}
func (this *MiniblockMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&fullHistory.MiniblockMetadata{")
	s = append(s, "RcvShardID: "+fmt.Sprintf("%#v", this.RcvShardID)+",\n")
	s = append(s, "SndShardID: "+fmt.Sprintf("%#v", this.SndShardID)+",\n")
	s = append(s, "Round: "+fmt.Sprintf("%#v", this.Round)+",\n")
	s = append(s, "HeaderNonce: "+fmt.Sprintf("%#v", this.HeaderNonce)+",\n")
	s = append(s, "HeaderHash: "+fmt.Sprintf("%#v", this.HeaderHash)+",\n")
	s = append(s, "MbHash: "+fmt.Sprintf("%#v", this.MbHash)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Epoch: "+fmt.Sprintf("%#v", this.Epoch)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMiniblockMetadata(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *MiniblockMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MiniblockMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MiniblockMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Epoch != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.MbHash) > 0 {
		i -= len(m.MbHash)
		copy(dAtA[i:], m.MbHash)
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(len(m.MbHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.HeaderHash) > 0 {
		i -= len(m.HeaderHash)
		copy(dAtA[i:], m.HeaderHash)
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(len(m.HeaderHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.HeaderNonce != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.HeaderNonce))
		i--
		dAtA[i] = 0x20
	}
	if m.Round != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x18
	}
	if m.SndShardID != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.SndShardID))
		i--
		dAtA[i] = 0x10
	}
	if m.RcvShardID != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.RcvShardID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMiniblockMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMiniblockMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MiniblockMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RcvShardID != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.RcvShardID))
	}
	if m.SndShardID != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.SndShardID))
	}
	if m.Round != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.Round))
	}
	if m.HeaderNonce != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.HeaderNonce))
	}
	l = len(m.HeaderHash)
	if l > 0 {
		n += 1 + l + sovMiniblockMetadata(uint64(l))
	}
	l = len(m.MbHash)
	if l > 0 {
		n += 1 + l + sovMiniblockMetadata(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovMiniblockMetadata(uint64(l))
	}
	if m.Epoch != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.Epoch))
	}
	return n
}

func sovMiniblockMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMiniblockMetadata(x uint64) (n int) {
	return sovMiniblockMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MiniblockMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MiniblockMetadata{`,
		`RcvShardID:` + fmt.Sprintf("%v", this.RcvShardID) + `,`,
		`SndShardID:` + fmt.Sprintf("%v", this.SndShardID) + `,`,
		`Round:` + fmt.Sprintf("%v", this.Round) + `,`,
		`HeaderNonce:` + fmt.Sprintf("%v", this.HeaderNonce) + `,`,
		`HeaderHash:` + fmt.Sprintf("%v", this.HeaderHash) + `,`,
		`MbHash:` + fmt.Sprintf("%v", this.MbHash) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Epoch:` + fmt.Sprintf("%v", this.Epoch) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMiniblockMetadata(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MiniblockMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMiniblockMetadata
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
			return fmt.Errorf("proto: MiniblockMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MiniblockMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RcvShardID", wireType)
			}
			m.RcvShardID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RcvShardID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SndShardID", wireType)
			}
			m.SndShardID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SndShardID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderNonce", wireType)
			}
			m.HeaderNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeaderNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
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
				return ErrInvalidLengthMiniblockMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderHash = append(m.HeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderHash == nil {
				m.HeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MbHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
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
				return ErrInvalidLengthMiniblockMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MbHash = append(m.MbHash[:0], dAtA[iNdEx:postIndex]...)
			if m.MbHash == nil {
				m.MbHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
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
				return ErrInvalidLengthMiniblockMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = append(m.Status[:0], dAtA[iNdEx:postIndex]...)
			if m.Status == nil {
				m.Status = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMiniblockMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMiniblockMetadata
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
func skipMiniblockMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMiniblockMetadata
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
					return 0, ErrIntOverflowMiniblockMetadata
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
					return 0, ErrIntOverflowMiniblockMetadata
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
				return 0, ErrInvalidLengthMiniblockMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMiniblockMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMiniblockMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMiniblockMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMiniblockMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMiniblockMetadata = fmt.Errorf("proto: unexpected end of group")
)
