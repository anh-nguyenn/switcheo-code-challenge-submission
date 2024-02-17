// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rental/rental/rental.proto

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

type Rental struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount  string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee     string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Asset   string `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	DueDate string `protobuf:"bytes,5,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
	State   string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Renter  string `protobuf:"bytes,7,opt,name=renter,proto3" json:"renter,omitempty"`
	Owner   string `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Rental) Reset()         { *m = Rental{} }
func (m *Rental) String() string { return proto.CompactTextString(m) }
func (*Rental) ProtoMessage()    {}
func (*Rental) Descriptor() ([]byte, []int) {
	return fileDescriptor_98599b8a9fb167e8, []int{0}
}
func (m *Rental) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rental) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rental.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rental) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rental.Merge(m, src)
}
func (m *Rental) XXX_Size() int {
	return m.Size()
}
func (m *Rental) XXX_DiscardUnknown() {
	xxx_messageInfo_Rental.DiscardUnknown(m)
}

var xxx_messageInfo_Rental proto.InternalMessageInfo

func (m *Rental) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Rental) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Rental) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *Rental) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *Rental) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *Rental) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Rental) GetRenter() string {
	if m != nil {
		return m.Renter
	}
	return ""
}

func (m *Rental) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*Rental)(nil), "rental.rental.Rental")
}

func init() { proto.RegisterFile("rental/rental/rental.proto", fileDescriptor_98599b8a9fb167e8) }

var fileDescriptor_98599b8a9fb167e8 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4a, 0xcd, 0x2b,
	0x49, 0xcc, 0xd1, 0x47, 0xa1, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0xa1, 0x3c, 0x08,
	0xa5, 0xb4, 0x8d, 0x91, 0x8b, 0x2d, 0x08, 0xcc, 0x14, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe3, 0x62, 0x4b, 0xcc, 0xcd, 0x2f,
	0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x04, 0xb8, 0x98, 0xd3,
	0x52, 0x53, 0x25, 0x98, 0xc1, 0x82, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x62, 0x71, 0x71, 0x6a,
	0x89, 0x04, 0x0b, 0x58, 0x0c, 0xc2, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x29, 0x4d, 0x75, 0x49, 0x2c,
	0x49, 0x95, 0x60, 0x05, 0x8b, 0xc3, 0xb8, 0x20, 0xf5, 0xc5, 0x25, 0x20, 0x71, 0x36, 0x88, 0x7a,
	0x30, 0x07, 0x64, 0x1f, 0xc8, 0x51, 0xa9, 0x45, 0x12, 0xec, 0x10, 0xfb, 0x20, 0x3c, 0x90, 0xea,
	0xfc, 0xf2, 0xbc, 0xd4, 0x22, 0x09, 0x0e, 0x88, 0x6a, 0x30, 0xc7, 0x49, 0xff, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x44, 0xa1, 0xde, 0xae, 0x80, 0xf9, 0xbf, 0xa4, 0xb2,
	0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x7f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xcf,
	0x79, 0x37, 0x1d, 0x01, 0x00, 0x00,
}

func (m *Rental) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rental) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rental) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintRental(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Renter) > 0 {
		i -= len(m.Renter)
		copy(dAtA[i:], m.Renter)
		i = encodeVarintRental(dAtA, i, uint64(len(m.Renter)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintRental(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DueDate) > 0 {
		i -= len(m.DueDate)
		copy(dAtA[i:], m.DueDate)
		i = encodeVarintRental(dAtA, i, uint64(len(m.DueDate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintRental(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintRental(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintRental(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintRental(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRental(dAtA []byte, offset int, v uint64) int {
	offset -= sovRental(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Rental) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRental(uint64(m.Id))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	l = len(m.DueDate)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	l = len(m.Renter)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	return n
}

func sovRental(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRental(x uint64) (n int) {
	return sovRental(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Rental) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRental
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
			return fmt.Errorf("proto: Rental: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rental: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DueDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DueDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Renter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Renter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRental(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRental
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
func skipRental(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRental
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
					return 0, ErrIntOverflowRental
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
					return 0, ErrIntOverflowRental
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
				return 0, ErrInvalidLengthRental
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRental
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRental
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRental        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRental          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRental = fmt.Errorf("proto: unexpected end of group")
)
