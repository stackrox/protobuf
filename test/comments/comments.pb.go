// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test/comments/comments.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TagTest struct {
	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"` // @gotags: tag:"on-field-a"
	// Types that are valid to be assigned to OneOf:
	//
	//	*TagTest_OneOfA
	//	*TagTest_OneOfB
	//	*TagTest_OneOfC
	//	*TagTest_OneOfD
	OneOf                isTagTest_OneOf `protobuf_oneof:"OneOf"`
	C                    string          `protobuf:"bytes,6,opt,name=c,proto3" json:"c,omitempty"` // @gotags: tag:"on-field-c"
	D                    string          `protobuf:"bytes,7,opt,name=d,proto3" json:"d,omitempty"` // Not relevant comment
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TagTest) Reset()         { *m = TagTest{} }
func (m *TagTest) String() string { return proto.CompactTextString(m) }
func (*TagTest) ProtoMessage()    {}
func (*TagTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2de1ce75404ef10, []int{0}
}
func (m *TagTest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TagTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TagTest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TagTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagTest.Merge(m, src)
}
func (m *TagTest) XXX_Size() int {
	return m.Size()
}
func (m *TagTest) XXX_DiscardUnknown() {
	xxx_messageInfo_TagTest.DiscardUnknown(m)
}

var xxx_messageInfo_TagTest proto.InternalMessageInfo

type isTagTest_OneOf interface {
	isTagTest_OneOf()
	MarshalTo([]byte) (int, error)
	Size() int
	Clone() isTagTest_OneOf
}

type TagTest_OneOfA struct {
	OneOfA string `protobuf:"bytes,2,opt,name=oneOfA,proto3,oneof" json:"oneOfA,omitempty"` // @gotags: tag:"on-field-one-of-a"
}
type TagTest_OneOfB struct {
	OneOfB string `protobuf:"bytes,3,opt,name=oneOfB,proto3,oneof" json:"oneOfB,omitempty"`
}
type TagTest_OneOfC struct {
	OneOfC string `protobuf:"bytes,4,opt,name=oneOfC,proto3,oneof" json:"oneOfC,omitempty"` // @gotags: search:"-" sql:"ignore-fks,ignore-index"
}
type TagTest_OneOfD struct {
	OneOfD string `protobuf:"bytes,5,opt,name=oneOfD,proto3,oneof" json:"oneOfD,omitempty"` // Not relevant comment
}

func (*TagTest_OneOfA) isTagTest_OneOf() {}
func (m *TagTest_OneOfA) Clone() isTagTest_OneOf {
	if m == nil {
		return nil
	}
	cloned := new(TagTest_OneOfA)
	*cloned = *m

	return cloned
}
func (*TagTest_OneOfB) isTagTest_OneOf() {}
func (m *TagTest_OneOfB) Clone() isTagTest_OneOf {
	if m == nil {
		return nil
	}
	cloned := new(TagTest_OneOfB)
	*cloned = *m

	return cloned
}
func (*TagTest_OneOfC) isTagTest_OneOf() {}
func (m *TagTest_OneOfC) Clone() isTagTest_OneOf {
	if m == nil {
		return nil
	}
	cloned := new(TagTest_OneOfC)
	*cloned = *m

	return cloned
}
func (*TagTest_OneOfD) isTagTest_OneOf() {}
func (m *TagTest_OneOfD) Clone() isTagTest_OneOf {
	if m == nil {
		return nil
	}
	cloned := new(TagTest_OneOfD)
	*cloned = *m

	return cloned
}

func (m *TagTest) GetOneOf() isTagTest_OneOf {
	if m != nil {
		return m.OneOf
	}
	return nil
}

func (m *TagTest) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *TagTest) GetOneOfA() string {
	if x, ok := m.GetOneOf().(*TagTest_OneOfA); ok {
		return x.OneOfA
	}
	return ""
}

func (m *TagTest) GetOneOfB() string {
	if x, ok := m.GetOneOf().(*TagTest_OneOfB); ok {
		return x.OneOfB
	}
	return ""
}

func (m *TagTest) GetOneOfC() string {
	if x, ok := m.GetOneOf().(*TagTest_OneOfC); ok {
		return x.OneOfC
	}
	return ""
}

func (m *TagTest) GetOneOfD() string {
	if x, ok := m.GetOneOf().(*TagTest_OneOfD); ok {
		return x.OneOfD
	}
	return ""
}

func (m *TagTest) GetC() string {
	if m != nil {
		return m.C
	}
	return ""
}

func (m *TagTest) GetD() string {
	if m != nil {
		return m.D
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagTest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagTest_OneOfA)(nil),
		(*TagTest_OneOfB)(nil),
		(*TagTest_OneOfC)(nil),
		(*TagTest_OneOfD)(nil),
	}
}

func (m *TagTest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *TagTest) Clone() *TagTest {
	if m == nil {
		return nil
	}
	cloned := new(TagTest)
	*cloned = *m

	if m.OneOf != nil {
		cloned.OneOf = m.OneOf.Clone()
	}
	return cloned
}

func init() {
	proto.RegisterType((*TagTest)(nil), "storage.TagTest")
}

func init() { proto.RegisterFile("test/comments/comments.proto", fileDescriptor_d2de1ce75404ef10) }

var fileDescriptor_d2de1ce75404ef10 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x49, 0x2d, 0x2e,
	0xd1, 0x4f, 0xce, 0xcf, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0x86, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0xd8, 0x8b, 0x4b, 0xf2, 0x8b, 0x12, 0xd3, 0x53, 0x95, 0x96, 0x30, 0x72, 0xb1, 0x87,
	0x24, 0xa6, 0x87, 0xa4, 0x16, 0x97, 0x08, 0xf1, 0x70, 0x31, 0x26, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x31, 0x26, 0x0a, 0x49, 0x70, 0xb1, 0xe5, 0xe7, 0xa5, 0xfa, 0xa7, 0x39, 0x4a, 0x30,
	0x81, 0x84, 0x3c, 0x18, 0x82, 0xa0, 0x7c, 0xb8, 0x8c, 0x93, 0x04, 0x33, 0x8a, 0x8c, 0x13, 0x5c,
	0xc6, 0x59, 0x82, 0x05, 0x45, 0xc6, 0x19, 0x2e, 0xe3, 0x22, 0xc1, 0x8a, 0x22, 0xe3, 0x02, 0xb2,
	0x35, 0x59, 0x82, 0x0d, 0x62, 0x6b, 0x32, 0x88, 0x97, 0x22, 0xc1, 0x0e, 0xe1, 0xa5, 0x38, 0xb1,
	0x73, 0xb1, 0xfa, 0x83, 0x4d, 0x96, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x88, 0x82, 0xf9, 0x20, 0x89, 0x0d, 0xec, 0x23, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xdb, 0xf5, 0x5f, 0xf1, 0x00, 0x00, 0x00,
}

func (m *TagTest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TagTest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TagTest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.D) > 0 {
		i -= len(m.D)
		copy(dAtA[i:], m.D)
		i = encodeVarintComments(dAtA, i, uint64(len(m.D)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.C) > 0 {
		i -= len(m.C)
		copy(dAtA[i:], m.C)
		i = encodeVarintComments(dAtA, i, uint64(len(m.C)))
		i--
		dAtA[i] = 0x32
	}
	if m.OneOf != nil {
		{
			size := m.OneOf.Size()
			i -= size
			if _, err := m.OneOf.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.A) > 0 {
		i -= len(m.A)
		copy(dAtA[i:], m.A)
		i = encodeVarintComments(dAtA, i, uint64(len(m.A)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TagTest_OneOfA) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TagTest_OneOfA) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.OneOfA)
	copy(dAtA[i:], m.OneOfA)
	i = encodeVarintComments(dAtA, i, uint64(len(m.OneOfA)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *TagTest_OneOfB) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TagTest_OneOfB) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.OneOfB)
	copy(dAtA[i:], m.OneOfB)
	i = encodeVarintComments(dAtA, i, uint64(len(m.OneOfB)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *TagTest_OneOfC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TagTest_OneOfC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.OneOfC)
	copy(dAtA[i:], m.OneOfC)
	i = encodeVarintComments(dAtA, i, uint64(len(m.OneOfC)))
	i--
	dAtA[i] = 0x22
	return len(dAtA) - i, nil
}
func (m *TagTest_OneOfD) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TagTest_OneOfD) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.OneOfD)
	copy(dAtA[i:], m.OneOfD)
	i = encodeVarintComments(dAtA, i, uint64(len(m.OneOfD)))
	i--
	dAtA[i] = 0x2a
	return len(dAtA) - i, nil
}
func encodeVarintComments(dAtA []byte, offset int, v uint64) int {
	offset -= sovComments(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TagTest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.A)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	if m.OneOf != nil {
		n += m.OneOf.Size()
	}
	l = len(m.C)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	l = len(m.D)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TagTest_OneOfA) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OneOfA)
	n += 1 + l + sovComments(uint64(l))
	return n
}
func (m *TagTest_OneOfB) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OneOfB)
	n += 1 + l + sovComments(uint64(l))
	return n
}
func (m *TagTest_OneOfC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OneOfC)
	n += 1 + l + sovComments(uint64(l))
	return n
}
func (m *TagTest_OneOfD) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OneOfD)
	n += 1 + l + sovComments(uint64(l))
	return n
}

func sovComments(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComments(x uint64) (n int) {
	return sovComments(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TagTest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComments
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
			return fmt.Errorf("proto: TagTest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TagTest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneOfA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OneOf = &TagTest_OneOfA{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneOfB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OneOf = &TagTest_OneOfB{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneOfC", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OneOf = &TagTest_OneOfC{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneOfD", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OneOf = &TagTest_OneOfD{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComments(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComments
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipComments(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComments
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
					return 0, ErrIntOverflowComments
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
					return 0, ErrIntOverflowComments
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
				return 0, ErrInvalidLengthComments
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComments
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComments
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComments        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComments          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComments = fmt.Errorf("proto: unexpected end of group")
)
