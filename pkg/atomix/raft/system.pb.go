// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/raft/system.proto

package raft

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Raft system metadata
type RaftMetadata struct {
	Term Term     `protobuf:"varint,1,opt,name=term,proto3,casttype=Term" json:"term,omitempty"`
	Vote MemberID `protobuf:"bytes,2,opt,name=vote,proto3,casttype=MemberID" json:"vote,omitempty"`
}

func (m *RaftMetadata) Reset()         { *m = RaftMetadata{} }
func (m *RaftMetadata) String() string { return proto.CompactTextString(m) }
func (*RaftMetadata) ProtoMessage()    {}
func (*RaftMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_95061fb6dacccbf8, []int{0}
}
func (m *RaftMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RaftMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RaftMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftMetadata.Merge(m, src)
}
func (m *RaftMetadata) XXX_Size() int {
	return m.Size()
}
func (m *RaftMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RaftMetadata proto.InternalMessageInfo

func (m *RaftMetadata) GetTerm() Term {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RaftMetadata) GetVote() MemberID {
	if m != nil {
		return m.Vote
	}
	return ""
}

// Raft system configuration
type RaftConfiguration struct {
	Index     Index         `protobuf:"varint,1,opt,name=index,proto3,casttype=Index" json:"index,omitempty"`
	Term      Term          `protobuf:"varint,2,opt,name=term,proto3,casttype=Term" json:"term,omitempty"`
	Timestamp *time.Time    `protobuf:"bytes,3,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
	Members   []*RaftMember `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
}

func (m *RaftConfiguration) Reset()         { *m = RaftConfiguration{} }
func (m *RaftConfiguration) String() string { return proto.CompactTextString(m) }
func (*RaftConfiguration) ProtoMessage()    {}
func (*RaftConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_95061fb6dacccbf8, []int{1}
}
func (m *RaftConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RaftConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RaftConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftConfiguration.Merge(m, src)
}
func (m *RaftConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *RaftConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RaftConfiguration proto.InternalMessageInfo

func (m *RaftConfiguration) GetIndex() Index {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RaftConfiguration) GetTerm() Term {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RaftConfiguration) GetTimestamp() *time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *RaftConfiguration) GetMembers() []*RaftMember {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*RaftMetadata)(nil), "atomix.raft.RaftMetadata")
	proto.RegisterType((*RaftConfiguration)(nil), "atomix.raft.RaftConfiguration")
}

func init() { proto.RegisterFile("atomix/raft/system.proto", fileDescriptor_95061fb6dacccbf8) }

var fileDescriptor_95061fb6dacccbf8 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4f, 0x02, 0x31,
	0x14, 0xc7, 0x29, 0x1c, 0x0a, 0x85, 0xc5, 0x8b, 0xd1, 0x0b, 0x31, 0xed, 0x85, 0xe9, 0xa6, 0x5e,
	0xc4, 0xdd, 0xe1, 0x74, 0x61, 0xc0, 0xa1, 0xe1, 0x0b, 0x94, 0xd0, 0xbb, 0x5c, 0x62, 0x29, 0xe9,
	0x3d, 0x0c, 0x7e, 0x0b, 0xbe, 0x94, 0x89, 0x23, 0xa3, 0x13, 0x1a, 0xf8, 0x16, 0x4c, 0xa6, 0x2d,
	0x20, 0x83, 0x4b, 0xf3, 0xfa, 0xff, 0xff, 0xfa, 0x5e, 0xff, 0x0f, 0x47, 0x02, 0xb4, 0x2a, 0x97,
	0xa9, 0x11, 0x39, 0xa4, 0xd5, 0x7b, 0x05, 0x52, 0xb1, 0xb9, 0xd1, 0xa0, 0xc3, 0x8e, 0x77, 0x98,
	0x75, 0x7a, 0x37, 0xe7, 0x98, 0x3d, 0x3c, 0xd4, 0xa3, 0x85, 0xd6, 0xc5, 0xab, 0x4c, 0xdd, 0x6d,
	0xb2, 0xc8, 0x53, 0x28, 0x95, 0xac, 0x40, 0xa8, 0xf9, 0x01, 0xb8, 0x2e, 0x74, 0xa1, 0x5d, 0x99,
	0xda, 0xca, 0xab, 0xfd, 0x17, 0xdc, 0xe5, 0x22, 0x87, 0x91, 0x04, 0x31, 0x15, 0x20, 0xc2, 0x3b,
	0x1c, 0x80, 0x34, 0x2a, 0x42, 0x31, 0x4a, 0x1a, 0x59, 0x6b, 0xbf, 0xa1, 0xc1, 0x58, 0x1a, 0xc5,
	0x9d, 0x1a, 0xc6, 0x38, 0x78, 0xd3, 0x20, 0xa3, 0x7a, 0x8c, 0x92, 0x76, 0xd6, 0xdd, 0x6f, 0x68,
	0x6b, 0x24, 0xd5, 0x44, 0x9a, 0xe1, 0x33, 0x77, 0x4e, 0xff, 0x03, 0xe1, 0x2b, 0xdb, 0xf0, 0x49,
	0xcf, 0xf2, 0xb2, 0x58, 0x18, 0x01, 0xa5, 0x9e, 0x85, 0x14, 0x37, 0xcb, 0xd9, 0x54, 0x2e, 0x0f,
	0x6d, 0xdb, 0xfb, 0x0d, 0x6d, 0x0e, 0xad, 0xc0, 0xbd, 0x7e, 0x1a, 0x5b, 0xff, 0x77, 0xec, 0x23,
	0x6e, 0x9f, 0xd2, 0x44, 0x8d, 0x18, 0x25, 0x9d, 0x41, 0x8f, 0xf9, 0xbc, 0xec, 0x98, 0x97, 0x8d,
	0x8f, 0x44, 0x16, 0xac, 0xbe, 0x29, 0xe2, 0x7f, 0x4f, 0xc2, 0x7b, 0x7c, 0xa9, 0xdc, 0x37, 0xab,
	0x28, 0x88, 0x1b, 0x49, 0x67, 0x70, 0xcb, 0xce, 0x56, 0xca, 0xfc, 0x02, 0xac, 0xcf, 0x8f, 0x5c,
	0x16, 0x7d, 0x6e, 0x09, 0x5a, 0x6f, 0x09, 0xfa, 0xd9, 0x12, 0xb4, 0xda, 0x91, 0xda, 0x7a, 0x47,
	0x6a, 0x5f, 0x3b, 0x52, 0x9b, 0x5c, 0xb8, 0x89, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0xb2, 0x98, 0xd7, 0xb0, 0x01, 0x00, 0x00,
}

func (m *RaftMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vote) > 0 {
		i -= len(m.Vote)
		copy(dAtA[i:], m.Vote)
		i = encodeVarintSystem(dAtA, i, uint64(len(m.Vote)))
		i--
		dAtA[i] = 0x12
	}
	if m.Term != 0 {
		i = encodeVarintSystem(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RaftConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSystem(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Timestamp != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.Timestamp):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintSystem(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x1a
	}
	if m.Term != 0 {
		i = encodeVarintSystem(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x10
	}
	if m.Index != 0 {
		i = encodeVarintSystem(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSystem(dAtA []byte, offset int, v uint64) int {
	offset -= sovSystem(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RaftMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovSystem(uint64(m.Term))
	}
	l = len(m.Vote)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	return n
}

func (m *RaftConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovSystem(uint64(m.Index))
	}
	if m.Term != 0 {
		n += 1 + sovSystem(uint64(m.Term))
	}
	if m.Timestamp != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Timestamp)
		n += 1 + l + sovSystem(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovSystem(uint64(l))
		}
	}
	return n
}

func sovSystem(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSystem(x uint64) (n int) {
	return sovSystem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSystem
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
			return fmt.Errorf("proto: RaftMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= Term(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
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
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSystem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vote = MemberID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSystem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSystem
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSystem
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
func (m *RaftConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSystem
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
			return fmt.Errorf("proto: RaftConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= Term(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
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
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSystem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
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
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSystem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &RaftMember{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSystem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSystem
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSystem
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
func skipSystem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSystem
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
					return 0, ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSystem
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
				return 0, ErrInvalidLengthSystem
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSystem
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSystem
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSystem(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSystem
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSystem = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSystem   = fmt.Errorf("proto: integer overflow")
)
