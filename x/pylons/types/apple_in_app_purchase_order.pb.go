// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/apple_in_app_purchase_order.proto

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

type AppleInAppPurchaseOrder struct {
	Quantity     string `protobuf:"bytes,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ProductID    string `protobuf:"bytes,2,opt,name=productID,proto3" json:"productID,omitempty"`
	PurchaseID   string `protobuf:"bytes,3,opt,name=purchaseID,proto3" json:"purchaseID,omitempty"`
	PurchaseDate string `protobuf:"bytes,4,opt,name=purchaseDate,proto3" json:"purchaseDate,omitempty"`
	Creator      string `protobuf:"bytes,5,opt,name=Creator,proto3" json:"Creator,omitempty"`
}

func (m *AppleInAppPurchaseOrder) Reset()         { *m = AppleInAppPurchaseOrder{} }
func (m *AppleInAppPurchaseOrder) String() string { return proto.CompactTextString(m) }
func (*AppleInAppPurchaseOrder) ProtoMessage()    {}
func (*AppleInAppPurchaseOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac0a95fd515feea7, []int{0}
}
func (m *AppleInAppPurchaseOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppleInAppPurchaseOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppleInAppPurchaseOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppleInAppPurchaseOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppleInAppPurchaseOrder.Merge(m, src)
}
func (m *AppleInAppPurchaseOrder) XXX_Size() int {
	return m.Size()
}
func (m *AppleInAppPurchaseOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_AppleInAppPurchaseOrder.DiscardUnknown(m)
}

var xxx_messageInfo_AppleInAppPurchaseOrder proto.InternalMessageInfo

func (m *AppleInAppPurchaseOrder) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *AppleInAppPurchaseOrder) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *AppleInAppPurchaseOrder) GetPurchaseID() string {
	if m != nil {
		return m.PurchaseID
	}
	return ""
}

func (m *AppleInAppPurchaseOrder) GetPurchaseDate() string {
	if m != nil {
		return m.PurchaseDate
	}
	return ""
}

func (m *AppleInAppPurchaseOrder) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*AppleInAppPurchaseOrder)(nil), "pylonstech.pylons.pylons.AppleInAppPurchaseOrder")
}

func init() {
	proto.RegisterFile("pylons/apple_in_app_purchase_order.proto", fileDescriptor_ac0a95fd515feea7)
}

var fileDescriptor_ac0a95fd515feea7 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0xa8, 0xcc, 0xc9,
	0xcf, 0x2b, 0xd6, 0x4f, 0x2c, 0x28, 0xc8, 0x49, 0x8d, 0xcf, 0xcc, 0x8b, 0x4f, 0x2c, 0x28, 0x88,
	0x2f, 0x28, 0x2d, 0x4a, 0xce, 0x48, 0x2c, 0x4e, 0x8d, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x80, 0xa8, 0x2c, 0x49, 0x4d, 0xce, 0xd0, 0x83, 0x30, 0xa1,
	0x94, 0xd2, 0x46, 0x46, 0x2e, 0x71, 0x47, 0x90, 0x7e, 0xcf, 0x3c, 0xc7, 0x82, 0x82, 0x00, 0xa8,
	0x66, 0x7f, 0x90, 0x5e, 0x21, 0x29, 0x2e, 0x8e, 0xc2, 0xd2, 0xc4, 0xbc, 0x92, 0xcc, 0x92, 0x4a,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x48, 0x86, 0x8b, 0xb3, 0xa0, 0x28, 0x3f,
	0xa5, 0x34, 0xb9, 0xc4, 0xd3, 0x45, 0x82, 0x09, 0x2c, 0x89, 0x10, 0x10, 0x92, 0xe3, 0xe2, 0x82,
	0xb9, 0xc3, 0xd3, 0x45, 0x82, 0x19, 0x2c, 0x8d, 0x24, 0x22, 0xa4, 0xc4, 0xc5, 0x03, 0xe3, 0xb9,
	0x24, 0x96, 0xa4, 0x4a, 0xb0, 0x80, 0x55, 0xa0, 0x88, 0x09, 0x49, 0x70, 0xb1, 0x3b, 0x17, 0xa5,
	0x26, 0x96, 0xe4, 0x17, 0x49, 0xb0, 0x82, 0xa5, 0x61, 0x5c, 0x27, 0xb7, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x0f, 0x00, 0x7b, 0x50, 0x17, 0xe4, 0x67, 0x7d, 0x68, 0x40, 0x55, 0xc0, 0x18, 0x25,
	0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xc0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3a,
	0xef, 0x33, 0x5e, 0x48, 0x01, 0x00, 0x00,
}

func (m *AppleInAppPurchaseOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppleInAppPurchaseOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppleInAppPurchaseOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAppleInAppPurchaseOrder(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PurchaseDate) > 0 {
		i -= len(m.PurchaseDate)
		copy(dAtA[i:], m.PurchaseDate)
		i = encodeVarintAppleInAppPurchaseOrder(dAtA, i, uint64(len(m.PurchaseDate)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PurchaseID) > 0 {
		i -= len(m.PurchaseID)
		copy(dAtA[i:], m.PurchaseID)
		i = encodeVarintAppleInAppPurchaseOrder(dAtA, i, uint64(len(m.PurchaseID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProductID) > 0 {
		i -= len(m.ProductID)
		copy(dAtA[i:], m.ProductID)
		i = encodeVarintAppleInAppPurchaseOrder(dAtA, i, uint64(len(m.ProductID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Quantity) > 0 {
		i -= len(m.Quantity)
		copy(dAtA[i:], m.Quantity)
		i = encodeVarintAppleInAppPurchaseOrder(dAtA, i, uint64(len(m.Quantity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAppleInAppPurchaseOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovAppleInAppPurchaseOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AppleInAppPurchaseOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Quantity)
	if l > 0 {
		n += 1 + l + sovAppleInAppPurchaseOrder(uint64(l))
	}
	l = len(m.ProductID)
	if l > 0 {
		n += 1 + l + sovAppleInAppPurchaseOrder(uint64(l))
	}
	l = len(m.PurchaseID)
	if l > 0 {
		n += 1 + l + sovAppleInAppPurchaseOrder(uint64(l))
	}
	l = len(m.PurchaseDate)
	if l > 0 {
		n += 1 + l + sovAppleInAppPurchaseOrder(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAppleInAppPurchaseOrder(uint64(l))
	}
	return n
}

func sovAppleInAppPurchaseOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAppleInAppPurchaseOrder(x uint64) (n int) {
	return sovAppleInAppPurchaseOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppleInAppPurchaseOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppleInAppPurchaseOrder
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
			return fmt.Errorf("proto: AppleInAppPurchaseOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppleInAppPurchaseOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppleInAppPurchaseOrder
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
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quantity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppleInAppPurchaseOrder
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
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProductID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchaseID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppleInAppPurchaseOrder
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
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PurchaseID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchaseDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppleInAppPurchaseOrder
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
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PurchaseDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppleInAppPurchaseOrder
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
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppleInAppPurchaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAppleInAppPurchaseOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAppleInAppPurchaseOrder
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
func skipAppleInAppPurchaseOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAppleInAppPurchaseOrder
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
					return 0, ErrIntOverflowAppleInAppPurchaseOrder
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
					return 0, ErrIntOverflowAppleInAppPurchaseOrder
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
				return 0, ErrInvalidLengthAppleInAppPurchaseOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAppleInAppPurchaseOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAppleInAppPurchaseOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAppleInAppPurchaseOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAppleInAppPurchaseOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAppleInAppPurchaseOrder = fmt.Errorf("proto: unexpected end of group")
)
