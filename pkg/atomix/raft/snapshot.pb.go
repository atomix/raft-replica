// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomix/raft/snapshot.proto

package raft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Snapshot descriptor
type SnapshotDescriptor struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotDescriptor) Reset()         { *m = SnapshotDescriptor{} }
func (m *SnapshotDescriptor) String() string { return proto.CompactTextString(m) }
func (*SnapshotDescriptor) ProtoMessage()    {}
func (*SnapshotDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_58c3c314973b5e08, []int{0}
}

func (m *SnapshotDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotDescriptor.Unmarshal(m, b)
}
func (m *SnapshotDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotDescriptor.Marshal(b, m, deterministic)
}
func (m *SnapshotDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotDescriptor.Merge(m, src)
}
func (m *SnapshotDescriptor) XXX_Size() int {
	return xxx_messageInfo_SnapshotDescriptor.Size(m)
}
func (m *SnapshotDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotDescriptor proto.InternalMessageInfo

func (m *SnapshotDescriptor) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SnapshotDescriptor) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*SnapshotDescriptor)(nil), "atomix.raft.SnapshotDescriptor")
}

func init() { proto.RegisterFile("atomix/raft/snapshot.proto", fileDescriptor_58c3c314973b5e08) }

var fileDescriptor_58c3c314973b5e08 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0x2f, 0x4a, 0x4c, 0x2b, 0xd1, 0x2f, 0xce, 0x4b, 0x2c, 0x28, 0xce, 0xc8, 0x2f,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xc8, 0xe9, 0x81, 0xe4, 0x94, 0x3c, 0xb8,
	0x84, 0x82, 0xa1, 0xd2, 0x2e, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x42, 0x22,
	0x5c, 0xac, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e,
	0x90, 0x0c, 0x17, 0x67, 0x49, 0x66, 0x6e, 0x6a, 0x71, 0x49, 0x62, 0x6e, 0x81, 0x04, 0x13, 0x58,
	0x06, 0x21, 0x90, 0xc4, 0x06, 0x36, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x36, 0xce, 0x1d,
	0x62, 0x7b, 0x00, 0x00, 0x00,
}