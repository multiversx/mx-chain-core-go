// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rewardTx.proto

package rewardTx

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_multiversx_mx_chain_core_go_data "github.com/multiversx/mx-chain-core-go/data"
	io "io"
	math "math"
	math_big "math/big"
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

// RewardTx holds the data for a reward transaction
type RewardTx struct {
	Round   uint64        `protobuf:"varint,1,opt,name=Round,proto3" json:"round"`
	Value   *math_big.Int `protobuf:"bytes,3,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/multiversx/mx-chain-core-go/data.BigIntCaster" json:"value"`
	RcvAddr []byte        `protobuf:"bytes,4,opt,name=RcvAddr,proto3" json:"receiver"`
	Epoch   uint32        `protobuf:"varint,2,opt,name=Epoch,proto3" json:"epoch"`
}

func (m *RewardTx) Reset()      { *m = RewardTx{} }
func (*RewardTx) ProtoMessage() {}
func (*RewardTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_25dbfb608d6baddf, []int{0}
}
func (m *RewardTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RewardTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardTx.Merge(m, src)
}
func (m *RewardTx) XXX_Size() int {
	return m.Size()
}
func (m *RewardTx) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardTx.DiscardUnknown(m)
}

var xxx_messageInfo_RewardTx proto.InternalMessageInfo

func (m *RewardTx) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *RewardTx) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RewardTx) GetRcvAddr() []byte {
	if m != nil {
		return m.RcvAddr
	}
	return nil
}

func (m *RewardTx) GetEpoch() uint32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func init() {
	proto.RegisterType((*RewardTx)(nil), "proto.RewardTx")
}

func init() { proto.RegisterFile("rewardTx.proto", fileDescriptor_25dbfb608d6baddf) }

var fileDescriptor_25dbfb608d6baddf = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x73, 0xd0, 0x40, 0x89, 0x0a, 0x43, 0xa6, 0x88, 0xe1, 0x52, 0x31, 0xa0, 0x2e, 0x49,
	0x06, 0xc6, 0x8a, 0x81, 0xa0, 0x0e, 0x5d, 0x2d, 0xc4, 0xc0, 0xe6, 0x24, 0x26, 0x89, 0xd4, 0xc4,
	0x95, 0xeb, 0xb4, 0x1d, 0x79, 0x04, 0x1e, 0x03, 0xf1, 0x24, 0x8c, 0x1d, 0x3b, 0x15, 0xea, 0x2e,
	0xa8, 0x53, 0x1f, 0x01, 0xc5, 0x21, 0x12, 0x2b, 0x93, 0xef, 0x3e, 0xdb, 0xff, 0xff, 0xdb, 0x67,
	0x5d, 0x08, 0xb6, 0xa0, 0x22, 0x79, 0x58, 0xfa, 0x53, 0xc1, 0x25, 0xb7, 0x4d, 0xbd, 0x5c, 0x7a,
	0x69, 0x2e, 0xb3, 0x2a, 0xf2, 0x63, 0x5e, 0x04, 0x29, 0x4f, 0x79, 0xa0, 0x71, 0x54, 0x3d, 0xeb,
	0x4e, 0x37, 0xba, 0x6a, 0x6e, 0x5d, 0x6d, 0xc0, 0xea, 0x92, 0x5f, 0x21, 0xdb, 0xb5, 0x4c, 0xc2,
	0xab, 0x32, 0x71, 0xa0, 0x0f, 0x83, 0x4e, 0x78, 0xb6, 0xdf, 0xb8, 0xa6, 0xa8, 0x01, 0x69, 0xb8,
	0x9d, 0x59, 0xe6, 0x23, 0x9d, 0x54, 0xcc, 0x39, 0xee, 0xc3, 0xa0, 0x17, 0x92, 0xfa, 0xc0, 0xbc,
	0x06, 0xef, 0x9f, 0xee, 0xa8, 0xa0, 0x32, 0x0b, 0xa2, 0x3c, 0xf5, 0xc7, 0xa5, 0x1c, 0xfe, 0x49,
	0x51, 0x54, 0x13, 0x99, 0xcf, 0x99, 0x98, 0x2d, 0x83, 0x62, 0xe9, 0xc5, 0x19, 0xcd, 0x4b, 0x2f,
	0xe6, 0x82, 0x79, 0x29, 0x0f, 0x12, 0x2a, 0xa9, 0x1f, 0xe6, 0xe9, 0xb8, 0x94, 0xf7, 0x74, 0x26,
	0x99, 0x20, 0x8d, 0x81, 0x7d, 0x6d, 0x9d, 0x92, 0x78, 0x7e, 0x97, 0x24, 0xc2, 0xe9, 0x68, 0xaf,
	0xde, 0x7e, 0xe3, 0x76, 0x05, 0x8b, 0x59, 0x2d, 0x45, 0xda, 0xcd, 0x3a, 0xf2, 0x68, 0xca, 0xe3,
	0xcc, 0x39, 0xea, 0xc3, 0xe0, 0xbc, 0x89, 0xcc, 0x6a, 0x40, 0x1a, 0x1e, 0x2e, 0x56, 0x5b, 0x34,
	0xd6, 0x5b, 0x34, 0x0e, 0x5b, 0x84, 0x17, 0x85, 0xf0, 0xa6, 0x10, 0x3e, 0x14, 0xc2, 0x4a, 0x21,
	0xac, 0x15, 0xc2, 0x97, 0x42, 0xf8, 0x56, 0x68, 0x1c, 0x14, 0xc2, 0xeb, 0x0e, 0x8d, 0xd5, 0x0e,
	0x8d, 0xf5, 0x0e, 0x8d, 0xa7, 0xdb, 0x7f, 0xbc, 0x21, 0x68, 0x87, 0x31, 0x6c, 0x8b, 0xe8, 0x44,
	0x7f, 0xf0, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xd6, 0xe2, 0x86, 0xa8, 0x01, 0x00,
	0x00,
}

func (this *RewardTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RewardTx)
	if !ok {
		that2, ok := that.(RewardTx)
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
	if this.Round != that1.Round {
		return false
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.RcvAddr, that1.RcvAddr) {
		return false
	}
	if this.Epoch != that1.Epoch {
		return false
	}
	return true
}
func (this *RewardTx) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&rewardTx.RewardTx{")
	s = append(s, "Round: "+fmt.Sprintf("%#v", this.Round)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "RcvAddr: "+fmt.Sprintf("%#v", this.RcvAddr)+",\n")
	s = append(s, "Epoch: "+fmt.Sprintf("%#v", this.Epoch)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRewardTx(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *RewardTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RcvAddr) > 0 {
		i -= len(m.RcvAddr)
		copy(dAtA[i:], m.RcvAddr)
		i = encodeVarintRewardTx(dAtA, i, uint64(len(m.RcvAddr)))
		i--
		dAtA[i] = 0x22
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRewardTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Epoch != 0 {
		i = encodeVarintRewardTx(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x10
	}
	if m.Round != 0 {
		i = encodeVarintRewardTx(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRewardTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovRewardTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovRewardTx(uint64(m.Round))
	}
	if m.Epoch != 0 {
		n += 1 + sovRewardTx(uint64(m.Epoch))
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovRewardTx(uint64(l))
	}
	l = len(m.RcvAddr)
	if l > 0 {
		n += 1 + l + sovRewardTx(uint64(l))
	}
	return n
}

func sovRewardTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRewardTx(x uint64) (n int) {
	return sovRewardTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RewardTx) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RewardTx{`,
		`Round:` + fmt.Sprintf("%v", this.Round) + `,`,
		`Epoch:` + fmt.Sprintf("%v", this.Epoch) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`RcvAddr:` + fmt.Sprintf("%v", this.RcvAddr) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRewardTx(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RewardTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewardTx
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
			return fmt.Errorf("proto: RewardTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardTx
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardTx
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardTx
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
				return ErrInvalidLengthRewardTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RcvAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardTx
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
				return ErrInvalidLengthRewardTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RcvAddr = append(m.RcvAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.RcvAddr == nil {
				m.RcvAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewardTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRewardTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRewardTx
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
func skipRewardTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRewardTx
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
					return 0, ErrIntOverflowRewardTx
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
					return 0, ErrIntOverflowRewardTx
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
				return 0, ErrInvalidLengthRewardTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRewardTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRewardTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRewardTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRewardTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRewardTx = fmt.Errorf("proto: unexpected end of group")
)
