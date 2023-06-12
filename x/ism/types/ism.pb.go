// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyperlane/ism/v1/ism.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Hyperlane's default ISM
type Ism struct {
	Origin      uint32     `protobuf:"varint,1,opt,name=origin,proto3" json:"origin,omitempty"`
	AbstractIsm *types.Any `protobuf:"bytes,2,opt,name=abstract_ism,json=abstractIsm,proto3" json:"abstract_ism,omitempty"`
}

func (m *Ism) Reset()         { *m = Ism{} }
func (m *Ism) String() string { return proto.CompactTextString(m) }
func (*Ism) ProtoMessage()    {}
func (*Ism) Descriptor() ([]byte, []int) {
	return fileDescriptor_425ee6960ca43e9a, []int{0}
}

func (m *Ism) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Ism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ism.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Ism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ism.Merge(m, src)
}

func (m *Ism) XXX_Size() int {
	return m.Size()
}

func (m *Ism) XXX_DiscardUnknown() {
	xxx_messageInfo_Ism.DiscardUnknown(m)
}

var xxx_messageInfo_Ism proto.InternalMessageInfo

func (m *Ism) GetOrigin() uint32 {
	if m != nil {
		return m.Origin
	}
	return 0
}

func (m *Ism) GetAbstractIsm() *types.Any {
	if m != nil {
		return m.AbstractIsm
	}
	return nil
}

func init() {
	proto.RegisterType((*Ism)(nil), "hyperlane.ism.v1.Ism")
}

func init() { proto.RegisterFile("hyperlane/ism/v1/ism.proto", fileDescriptor_425ee6960ca43e9a) }

var fileDescriptor_425ee6960ca43e9a = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xa8, 0x2c, 0x48,
	0x2d, 0xca, 0x49, 0xcc, 0x4b, 0xd5, 0xcf, 0x2c, 0xce, 0xd5, 0x2f, 0x33, 0x04, 0x51, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x02, 0x70, 0x39, 0x3d, 0x90, 0x60, 0x99, 0xa1, 0x94, 0x64, 0x7a,
	0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x3e, 0xa9, 0x34, 0x4d, 0x3f, 0x31, 0xaf, 0x12, 0xa2,
	0x58, 0x29, 0x8c, 0x8b, 0xd9, 0xb3, 0x38, 0x57, 0x48, 0x8c, 0x8b, 0x2d, 0xbf, 0x28, 0x33, 0x3d,
	0x33, 0x4f, 0x82, 0x51, 0x81, 0x51, 0x83, 0x37, 0x08, 0xca, 0x13, 0x32, 0xe7, 0xe2, 0x49, 0x4c,
	0x2a, 0x2e, 0x29, 0x4a, 0x4c, 0x2e, 0x89, 0xcf, 0x2c, 0xce, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x36, 0x12, 0xd1, 0x83, 0x18, 0xa8, 0x07, 0x33, 0x50, 0xcf, 0x31, 0xaf, 0x32, 0x88, 0x1b, 0xa6,
	0xd2, 0xb3, 0x38, 0xd7, 0x29, 0xec, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0x6c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x41, 0xea, 0xf3, 0xd2,
	0x53, 0x73, 0xf2, 0xcb, 0x52, 0x75, 0xcb, 0x52, 0xf3, 0x4a, 0x4a, 0x8b, 0x52, 0x8b, 0xf5, 0xe1,
	0xce, 0xd7, 0x4d, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0xaf, 0x00, 0xfb, 0xb1, 0xa4, 0xb2, 0x20,
	0xb5, 0x38, 0x89, 0x0d, 0x6c, 0xa5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x98, 0xc2, 0x65,
	0x01, 0x01, 0x00, 0x00,
}

func (m *Ism) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ism) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ism) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AbstractIsm != nil {
		{
			size, err := m.AbstractIsm.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIsm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Origin != 0 {
		i = encodeVarintIsm(dAtA, i, uint64(m.Origin))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIsm(dAtA []byte, offset int, v uint64) int {
	offset -= sovIsm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Ism) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Origin != 0 {
		n += 1 + sovIsm(uint64(m.Origin))
	}
	if m.AbstractIsm != nil {
		l = m.AbstractIsm.Size()
		n += 1 + l + sovIsm(uint64(l))
	}
	return n
}

func sovIsm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozIsm(x uint64) (n int) {
	return sovIsm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Ism) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIsm
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
			return fmt.Errorf("proto: Ism: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ism: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			m.Origin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIsm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Origin |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbstractIsm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIsm
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
				return ErrInvalidLengthIsm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIsm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AbstractIsm == nil {
				m.AbstractIsm = &types.Any{}
			}
			if err := m.AbstractIsm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIsm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIsm
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

func skipIsm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIsm
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
					return 0, ErrIntOverflowIsm
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
					return 0, ErrIntOverflowIsm
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
				return 0, ErrInvalidLengthIsm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIsm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIsm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIsm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIsm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIsm = fmt.Errorf("proto: unexpected end of group")
)