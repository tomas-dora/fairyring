// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/keyshare/authorized_address.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type AuthorizedAddress struct {
	Target       string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	IsAuthorized bool   `protobuf:"varint,2,opt,name=isAuthorized,proto3" json:"isAuthorized,omitempty"`
	AuthorizedBy string `protobuf:"bytes,3,opt,name=authorizedBy,proto3" json:"authorizedBy,omitempty"`
}

func (m *AuthorizedAddress) Reset()         { *m = AuthorizedAddress{} }
func (m *AuthorizedAddress) String() string { return proto.CompactTextString(m) }
func (*AuthorizedAddress) ProtoMessage()    {}
func (*AuthorizedAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b09dee94c56d60e, []int{0}
}
func (m *AuthorizedAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthorizedAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthorizedAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthorizedAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizedAddress.Merge(m, src)
}
func (m *AuthorizedAddress) XXX_Size() int {
	return m.Size()
}
func (m *AuthorizedAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizedAddress.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizedAddress proto.InternalMessageInfo

func (m *AuthorizedAddress) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *AuthorizedAddress) GetIsAuthorized() bool {
	if m != nil {
		return m.IsAuthorized
	}
	return false
}

func (m *AuthorizedAddress) GetAuthorizedBy() string {
	if m != nil {
		return m.AuthorizedBy
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthorizedAddress)(nil), "fairyring.keyshare.AuthorizedAddress")
}

func init() {
	proto.RegisterFile("fairyring/keyshare/authorized_address.proto", fileDescriptor_7b09dee94c56d60e)
}

var fileDescriptor_7b09dee94c56d60e = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0x4b, 0xcc, 0x2c,
	0xaa, 0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0xcf, 0x4e, 0xad, 0x2c, 0xce, 0x48, 0x2c, 0x4a, 0xd5, 0x4f,
	0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca, 0xac, 0x4a, 0x4d, 0x89, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d,
	0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0x2b, 0xd6, 0x83, 0x29, 0x56, 0x2a,
	0xe6, 0x12, 0x74, 0x84, 0xab, 0x77, 0x84, 0x28, 0x17, 0x12, 0xe3, 0x62, 0x2b, 0x49, 0x2c, 0x4a,
	0x4f, 0x2d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x94, 0xb8, 0x78, 0x32,
	0x8b, 0x11, 0xca, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x82, 0x50, 0xc4, 0x40, 0x6a, 0x10, 0x0e,
	0x70, 0xaa, 0x94, 0x60, 0x06, 0x9b, 0x80, 0x22, 0xe6, 0x64, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3,
	0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x52, 0x08, 0xff, 0x54, 0x20, 0x7c, 0x54, 0x52, 0x59, 0x90, 0x5a,
	0x9c, 0xc4, 0x06, 0xf6, 0x85, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x25, 0xfc, 0x33, 0x00, 0xf4,
	0x00, 0x00, 0x00,
}

func (m *AuthorizedAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizedAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthorizedAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthorizedBy) > 0 {
		i -= len(m.AuthorizedBy)
		copy(dAtA[i:], m.AuthorizedBy)
		i = encodeVarintAuthorizedAddress(dAtA, i, uint64(len(m.AuthorizedBy)))
		i--
		dAtA[i] = 0x1a
	}
	if m.IsAuthorized {
		i--
		if m.IsAuthorized {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintAuthorizedAddress(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthorizedAddress(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthorizedAddress(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuthorizedAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovAuthorizedAddress(uint64(l))
	}
	if m.IsAuthorized {
		n += 2
	}
	l = len(m.AuthorizedBy)
	if l > 0 {
		n += 1 + l + sovAuthorizedAddress(uint64(l))
	}
	return n
}

func sovAuthorizedAddress(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthorizedAddress(x uint64) (n int) {
	return sovAuthorizedAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthorizedAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthorizedAddress
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
			return fmt.Errorf("proto: AuthorizedAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizedAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizedAddress
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
				return ErrInvalidLengthAuthorizedAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthorizedAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsAuthorized", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizedAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsAuthorized = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizedAddress
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
				return ErrInvalidLengthAuthorizedAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthorizedAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthorizedAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthorizedAddress
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
func skipAuthorizedAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthorizedAddress
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
					return 0, ErrIntOverflowAuthorizedAddress
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
					return 0, ErrIntOverflowAuthorizedAddress
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
				return 0, ErrInvalidLengthAuthorizedAddress
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthorizedAddress
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthorizedAddress
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthorizedAddress        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthorizedAddress          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthorizedAddress = fmt.Errorf("proto: unexpected end of group")
)
