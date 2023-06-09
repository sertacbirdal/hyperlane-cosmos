// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyperlane/ism/v1/merkle_root_multisig.proto

package merkle_root_multisig

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
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

// MerkleRootMultiSig ISM for a specific origin
type MerkleRootMultiSig struct {
	// Validator pub keys
	ValidatorPubKeys []string `protobuf:"bytes,1,rep,name=validator_pub_keys,json=validatorPubKeys,proto3" json:"validator_pub_keys,omitempty"`
	// number of validators required
	Threshold uint32 `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (m *MerkleRootMultiSig) Reset()         { *m = MerkleRootMultiSig{} }
func (m *MerkleRootMultiSig) String() string { return proto.CompactTextString(m) }
func (*MerkleRootMultiSig) ProtoMessage()    {}
func (*MerkleRootMultiSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_91aeb6c6ec826c59, []int{0}
}

func (m *MerkleRootMultiSig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MerkleRootMultiSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MerkleRootMultiSig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MerkleRootMultiSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerkleRootMultiSig.Merge(m, src)
}

func (m *MerkleRootMultiSig) XXX_Size() int {
	return m.Size()
}

func (m *MerkleRootMultiSig) XXX_DiscardUnknown() {
	xxx_messageInfo_MerkleRootMultiSig.DiscardUnknown(m)
}

var xxx_messageInfo_MerkleRootMultiSig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MerkleRootMultiSig)(nil), "hyperlane.ism.v1.MerkleRootMultiSig")
}

func init() {
	proto.RegisterFile("hyperlane/ism/v1/merkle_root_multisig.proto", fileDescriptor_91aeb6c6ec826c59)
}

var fileDescriptor_91aeb6c6ec826c59 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x13, 0x15, 0xe1, 0x02, 0xc2, 0x11, 0x2c, 0x0e, 0x91, 0xf5, 0xb0, 0x3a, 0xd0, 0xcb,
	0x72, 0xd8, 0x59, 0xda, 0xca, 0x81, 0xc6, 0xce, 0x26, 0x24, 0x77, 0xe3, 0x66, 0xc9, 0x6e, 0x26,
	0xec, 0x4c, 0x82, 0x79, 0x03, 0x4b, 0x1f, 0xc1, 0xc7, 0xb1, 0xbc, 0xd2, 0x52, 0x92, 0x17, 0x91,
	0x8b, 0x10, 0x1b, 0xbb, 0x61, 0xfe, 0x9f, 0x0f, 0xbe, 0x3f, 0xb8, 0xca, 0xdb, 0x0a, 0x9c, 0x49,
	0x4b, 0x90, 0x9a, 0xac, 0x6c, 0x56, 0xd2, 0x82, 0x2b, 0x0c, 0x24, 0x0e, 0x91, 0x13, 0x5b, 0x1b,
	0xd6, 0xa4, 0x55, 0x54, 0x39, 0x64, 0x0c, 0xa7, 0x63, 0x39, 0xd2, 0x64, 0xa3, 0x66, 0x75, 0x76,
	0xaa, 0x50, 0xe1, 0x10, 0xca, 0xfd, 0xf5, 0xdb, 0xbb, 0x7c, 0x09, 0xc2, 0xf5, 0x40, 0x89, 0x11,
	0x79, 0xbd, 0x67, 0x3c, 0x69, 0x15, 0x5e, 0x07, 0x61, 0x93, 0x1a, 0xbd, 0x4d, 0x19, 0x5d, 0x52,
	0xd5, 0x59, 0x52, 0x40, 0x4b, 0x33, 0x7f, 0x7e, 0xb8, 0x98, 0xc4, 0xd3, 0x31, 0x79, 0xa8, 0xb3,
	0x7b, 0x68, 0x29, 0x3c, 0x0f, 0x26, 0x9c, 0x3b, 0xa0, 0x1c, 0xcd, 0x76, 0x76, 0x30, 0xf7, 0x17,
	0x27, 0xf1, 0xdf, 0xe3, 0xf6, 0xe8, 0xed, 0xe3, 0xc2, 0xbb, 0x2b, 0x3e, 0x3b, 0xe1, 0xef, 0x3a,
	0xe1, 0x7f, 0x77, 0xc2, 0x7f, 0xef, 0x85, 0xb7, 0xeb, 0x85, 0xf7, 0xd5, 0x0b, 0xef, 0xf9, 0x51,
	0x69, 0xce, 0xeb, 0x2c, 0xda, 0xa0, 0x95, 0xc4, 0x2e, 0x2d, 0x15, 0x18, 0x6c, 0x60, 0xd9, 0x40,
	0xc9, 0xb5, 0x03, 0x92, 0xa3, 0xc9, 0x72, 0x83, 0x64, 0x91, 0xe4, 0xeb, 0xe0, 0xcf, 0x6d, 0x05,
	0xf4, 0xef, 0x04, 0xd9, 0xf1, 0xe0, 0x76, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x97, 0xfb, 0xa1,
	0x02, 0x32, 0x01, 0x00, 0x00,
}

func (m *MerkleRootMultiSig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MerkleRootMultiSig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MerkleRootMultiSig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Threshold != 0 {
		i = encodeVarintMerkleRootMultisig(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorPubKeys) > 0 {
		for iNdEx := len(m.ValidatorPubKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ValidatorPubKeys[iNdEx])
			copy(dAtA[i:], m.ValidatorPubKeys[iNdEx])
			i = encodeVarintMerkleRootMultisig(dAtA, i, uint64(len(m.ValidatorPubKeys[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMerkleRootMultisig(dAtA []byte, offset int, v uint64) int {
	offset -= sovMerkleRootMultisig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *MerkleRootMultiSig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorPubKeys) > 0 {
		for _, s := range m.ValidatorPubKeys {
			l = len(s)
			n += 1 + l + sovMerkleRootMultisig(uint64(l))
		}
	}
	if m.Threshold != 0 {
		n += 1 + sovMerkleRootMultisig(uint64(m.Threshold))
	}
	return n
}

func sovMerkleRootMultisig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozMerkleRootMultisig(x uint64) (n int) {
	return sovMerkleRootMultisig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *MerkleRootMultiSig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMerkleRootMultisig
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
			return fmt.Errorf("proto: MerkleRootMultiSig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MerkleRootMultiSig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorPubKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkleRootMultisig
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
				return ErrInvalidLengthMerkleRootMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMerkleRootMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorPubKeys = append(m.ValidatorPubKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkleRootMultisig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMerkleRootMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMerkleRootMultisig
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

func skipMerkleRootMultisig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMerkleRootMultisig
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
					return 0, ErrIntOverflowMerkleRootMultisig
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
					return 0, ErrIntOverflowMerkleRootMultisig
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
				return 0, ErrInvalidLengthMerkleRootMultisig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMerkleRootMultisig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMerkleRootMultisig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMerkleRootMultisig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMerkleRootMultisig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMerkleRootMultisig = fmt.Errorf("proto: unexpected end of group")
)
