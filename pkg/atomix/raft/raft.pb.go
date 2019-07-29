// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomix/raft/raft.proto

package raft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ReadConsistency int32

const (
	ReadConsistency_SEQUENTIAL         ReadConsistency = 0
	ReadConsistency_LINEARIZABLE_LEASE ReadConsistency = 1
	ReadConsistency_LINEARIZABLE       ReadConsistency = 2
)

var ReadConsistency_name = map[int32]string{
	0: "SEQUENTIAL",
	1: "LINEARIZABLE_LEASE",
	2: "LINEARIZABLE",
}

var ReadConsistency_value = map[string]int32{
	"SEQUENTIAL":         0,
	"LINEARIZABLE_LEASE": 1,
	"LINEARIZABLE":       2,
}

func (x ReadConsistency) String() string {
	return proto.EnumName(ReadConsistency_name, int32(x))
}

func (ReadConsistency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_32de74386ddf389b, []int{0}
}

type StorageLevel int32

const (
	StorageLevel_DISK   StorageLevel = 0
	StorageLevel_MAPPED StorageLevel = 1
)

var StorageLevel_name = map[int32]string{
	0: "DISK",
	1: "MAPPED",
}

var StorageLevel_value = map[string]int32{
	"DISK":   0,
	"MAPPED": 1,
}

func (x StorageLevel) String() string {
	return proto.EnumName(StorageLevel_name, int32(x))
}

func (StorageLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_32de74386ddf389b, []int{1}
}

type RaftMember_Type int32

const (
	RaftMember_INACTIVE   RaftMember_Type = 0
	RaftMember_PASSIVE    RaftMember_Type = 1
	RaftMember_PROMOTABLE RaftMember_Type = 2
	RaftMember_ACTIVE     RaftMember_Type = 3
)

var RaftMember_Type_name = map[int32]string{
	0: "INACTIVE",
	1: "PASSIVE",
	2: "PROMOTABLE",
	3: "ACTIVE",
}

var RaftMember_Type_value = map[string]int32{
	"INACTIVE":   0,
	"PASSIVE":    1,
	"PROMOTABLE": 2,
	"ACTIVE":     3,
}

func (x RaftMember_Type) String() string {
	return proto.EnumName(RaftMember_Type_name, int32(x))
}

func (RaftMember_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_32de74386ddf389b, []int{0, 0}
}

type RaftMember struct {
	MemberId             string          `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Type                 RaftMember_Type `protobuf:"varint,2,opt,name=type,proto3,enum=atomix.raft.RaftMember_Type" json:"type,omitempty"`
	Updated              int64           `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RaftMember) Reset()         { *m = RaftMember{} }
func (m *RaftMember) String() string { return proto.CompactTextString(m) }
func (*RaftMember) ProtoMessage()    {}
func (*RaftMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_32de74386ddf389b, []int{0}
}

func (m *RaftMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaftMember.Unmarshal(m, b)
}
func (m *RaftMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaftMember.Marshal(b, m, deterministic)
}
func (m *RaftMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftMember.Merge(m, src)
}
func (m *RaftMember) XXX_Size() int {
	return xxx_messageInfo_RaftMember.Size(m)
}
func (m *RaftMember) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftMember.DiscardUnknown(m)
}

var xxx_messageInfo_RaftMember proto.InternalMessageInfo

func (m *RaftMember) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

func (m *RaftMember) GetType() RaftMember_Type {
	if m != nil {
		return m.Type
	}
	return RaftMember_INACTIVE
}

func (m *RaftMember) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type SessionMetadata struct {
	SessionId            int64    `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceType          string   `protobuf:"bytes,3,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionMetadata) Reset()         { *m = SessionMetadata{} }
func (m *SessionMetadata) String() string { return proto.CompactTextString(m) }
func (*SessionMetadata) ProtoMessage()    {}
func (*SessionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_32de74386ddf389b, []int{1}
}

func (m *SessionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionMetadata.Unmarshal(m, b)
}
func (m *SessionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionMetadata.Marshal(b, m, deterministic)
}
func (m *SessionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionMetadata.Merge(m, src)
}
func (m *SessionMetadata) XXX_Size() int {
	return xxx_messageInfo_SessionMetadata.Size(m)
}
func (m *SessionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SessionMetadata proto.InternalMessageInfo

func (m *SessionMetadata) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *SessionMetadata) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *SessionMetadata) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

type RaftProtocolConfig struct {
	ElectionTimeout      *duration.Duration    `protobuf:"bytes,1,opt,name=election_timeout,json=electionTimeout,proto3" json:"election_timeout,omitempty"`
	HeartbeatInterval    *duration.Duration    `protobuf:"bytes,2,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"`
	Storage              *RaftStorageConfig    `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Compaction           *RaftCompactionConfig `protobuf:"bytes,4,opt,name=compaction,proto3" json:"compaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RaftProtocolConfig) Reset()         { *m = RaftProtocolConfig{} }
func (m *RaftProtocolConfig) String() string { return proto.CompactTextString(m) }
func (*RaftProtocolConfig) ProtoMessage()    {}
func (*RaftProtocolConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_32de74386ddf389b, []int{2}
}

func (m *RaftProtocolConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaftProtocolConfig.Unmarshal(m, b)
}
func (m *RaftProtocolConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaftProtocolConfig.Marshal(b, m, deterministic)
}
func (m *RaftProtocolConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftProtocolConfig.Merge(m, src)
}
func (m *RaftProtocolConfig) XXX_Size() int {
	return xxx_messageInfo_RaftProtocolConfig.Size(m)
}
func (m *RaftProtocolConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftProtocolConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RaftProtocolConfig proto.InternalMessageInfo

func (m *RaftProtocolConfig) GetElectionTimeout() *duration.Duration {
	if m != nil {
		return m.ElectionTimeout
	}
	return nil
}

func (m *RaftProtocolConfig) GetHeartbeatInterval() *duration.Duration {
	if m != nil {
		return m.HeartbeatInterval
	}
	return nil
}

func (m *RaftProtocolConfig) GetStorage() *RaftStorageConfig {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *RaftProtocolConfig) GetCompaction() *RaftCompactionConfig {
	if m != nil {
		return m.Compaction
	}
	return nil
}

type RaftStorageConfig struct {
	Directory            string       `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	Level                StorageLevel `protobuf:"varint,2,opt,name=level,proto3,enum=atomix.raft.StorageLevel" json:"level,omitempty"`
	MaxEntrySize         uint32       `protobuf:"varint,3,opt,name=max_entry_size,json=maxEntrySize,proto3" json:"max_entry_size,omitempty"`
	SegmentSize          uint32       `protobuf:"varint,4,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	FlushOnCommit        bool         `protobuf:"varint,5,opt,name=flush_on_commit,json=flushOnCommit,proto3" json:"flush_on_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RaftStorageConfig) Reset()         { *m = RaftStorageConfig{} }
func (m *RaftStorageConfig) String() string { return proto.CompactTextString(m) }
func (*RaftStorageConfig) ProtoMessage()    {}
func (*RaftStorageConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_32de74386ddf389b, []int{3}
}

func (m *RaftStorageConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaftStorageConfig.Unmarshal(m, b)
}
func (m *RaftStorageConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaftStorageConfig.Marshal(b, m, deterministic)
}
func (m *RaftStorageConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftStorageConfig.Merge(m, src)
}
func (m *RaftStorageConfig) XXX_Size() int {
	return xxx_messageInfo_RaftStorageConfig.Size(m)
}
func (m *RaftStorageConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftStorageConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RaftStorageConfig proto.InternalMessageInfo

func (m *RaftStorageConfig) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *RaftStorageConfig) GetLevel() StorageLevel {
	if m != nil {
		return m.Level
	}
	return StorageLevel_DISK
}

func (m *RaftStorageConfig) GetMaxEntrySize() uint32 {
	if m != nil {
		return m.MaxEntrySize
	}
	return 0
}

func (m *RaftStorageConfig) GetSegmentSize() uint32 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func (m *RaftStorageConfig) GetFlushOnCommit() bool {
	if m != nil {
		return m.FlushOnCommit
	}
	return false
}

type RaftCompactionConfig struct {
	Dynamic              bool     `protobuf:"varint,1,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	FreeDiskBuffer       float32  `protobuf:"fixed32,2,opt,name=free_disk_buffer,json=freeDiskBuffer,proto3" json:"free_disk_buffer,omitempty"`
	FreeMemoryBuffer     float32  `protobuf:"fixed32,3,opt,name=free_memory_buffer,json=freeMemoryBuffer,proto3" json:"free_memory_buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RaftCompactionConfig) Reset()         { *m = RaftCompactionConfig{} }
func (m *RaftCompactionConfig) String() string { return proto.CompactTextString(m) }
func (*RaftCompactionConfig) ProtoMessage()    {}
func (*RaftCompactionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_32de74386ddf389b, []int{4}
}

func (m *RaftCompactionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaftCompactionConfig.Unmarshal(m, b)
}
func (m *RaftCompactionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaftCompactionConfig.Marshal(b, m, deterministic)
}
func (m *RaftCompactionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftCompactionConfig.Merge(m, src)
}
func (m *RaftCompactionConfig) XXX_Size() int {
	return xxx_messageInfo_RaftCompactionConfig.Size(m)
}
func (m *RaftCompactionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftCompactionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RaftCompactionConfig proto.InternalMessageInfo

func (m *RaftCompactionConfig) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *RaftCompactionConfig) GetFreeDiskBuffer() float32 {
	if m != nil {
		return m.FreeDiskBuffer
	}
	return 0
}

func (m *RaftCompactionConfig) GetFreeMemoryBuffer() float32 {
	if m != nil {
		return m.FreeMemoryBuffer
	}
	return 0
}

func init() {
	proto.RegisterEnum("atomix.raft.ReadConsistency", ReadConsistency_name, ReadConsistency_value)
	proto.RegisterEnum("atomix.raft.StorageLevel", StorageLevel_name, StorageLevel_value)
	proto.RegisterEnum("atomix.raft.RaftMember_Type", RaftMember_Type_name, RaftMember_Type_value)
	proto.RegisterType((*RaftMember)(nil), "atomix.raft.RaftMember")
	proto.RegisterType((*SessionMetadata)(nil), "atomix.raft.SessionMetadata")
	proto.RegisterType((*RaftProtocolConfig)(nil), "atomix.raft.RaftProtocolConfig")
	proto.RegisterType((*RaftStorageConfig)(nil), "atomix.raft.RaftStorageConfig")
	proto.RegisterType((*RaftCompactionConfig)(nil), "atomix.raft.RaftCompactionConfig")
}

func init() { proto.RegisterFile("atomix/raft/raft.proto", fileDescriptor_32de74386ddf389b) }

var fileDescriptor_32de74386ddf389b = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xd1, 0x6e, 0xda, 0x30,
	0x14, 0x86, 0x1b, 0xa0, 0x2d, 0x1c, 0x28, 0xa4, 0xd6, 0x54, 0xd1, 0xad, 0xab, 0x28, 0xaa, 0x26,
	0x54, 0x4d, 0x61, 0xea, 0x6e, 0x76, 0xb3, 0x8b, 0x14, 0x22, 0x2d, 0x2a, 0x50, 0xe6, 0xb0, 0x5d,
	0xec, 0x26, 0x32, 0x89, 0x43, 0xad, 0xe2, 0x18, 0x25, 0x06, 0x95, 0x3e, 0xc1, 0x5e, 0x63, 0xef,
	0xb0, 0x37, 0xd9, 0x0b, 0x4d, 0x76, 0x42, 0x4b, 0xd5, 0x69, 0x37, 0x51, 0xfc, 0xfb, 0xfb, 0xed,
	0xf3, 0xfb, 0xd8, 0x70, 0x44, 0xa4, 0xe0, 0xec, 0xbe, 0x9b, 0x90, 0x48, 0xea, 0x8f, 0xb5, 0x48,
	0x84, 0x14, 0xa8, 0x9a, 0xe9, 0x96, 0x92, 0x5e, 0x9f, 0xce, 0x84, 0x98, 0xcd, 0x69, 0x57, 0x4f,
	0x4d, 0x97, 0x51, 0x37, 0x5c, 0x26, 0x44, 0x32, 0x11, 0x67, 0x70, 0xfb, 0xb7, 0x01, 0x80, 0x49,
	0x24, 0x87, 0x94, 0x4f, 0x69, 0x82, 0xde, 0x40, 0x85, 0xeb, 0x3f, 0x9f, 0x85, 0x4d, 0xa3, 0x65,
	0x74, 0x2a, 0xb8, 0x9c, 0x09, 0x6e, 0x88, 0x3e, 0x40, 0x49, 0xae, 0x17, 0xb4, 0x59, 0x68, 0x19,
	0x9d, 0xfa, 0xe5, 0x89, 0xb5, 0xb5, 0x8f, 0xf5, 0xb4, 0x86, 0x35, 0x59, 0x2f, 0x28, 0xd6, 0x24,
	0x6a, 0xc2, 0xfe, 0x72, 0x11, 0x12, 0x49, 0xc3, 0x66, 0xb1, 0x65, 0x74, 0x8a, 0x78, 0x33, 0x6c,
	0x7f, 0x86, 0x92, 0xe2, 0x50, 0x0d, 0xca, 0xee, 0xc8, 0xee, 0x4d, 0xdc, 0xef, 0x8e, 0xb9, 0x83,
	0xaa, 0xb0, 0x3f, 0xb6, 0x3d, 0x4f, 0x0d, 0x0c, 0x54, 0x07, 0x18, 0xe3, 0x9b, 0xe1, 0xcd, 0xc4,
	0xbe, 0x1a, 0x38, 0x66, 0x01, 0x01, 0xec, 0xe5, 0x60, 0xb1, 0xbd, 0x82, 0x86, 0x47, 0xd3, 0x94,
	0x89, 0x78, 0x48, 0x25, 0x09, 0x89, 0x24, 0xe8, 0x2d, 0x40, 0x9a, 0x49, 0x9b, 0xda, 0x8b, 0xb8,
	0x92, 0x2b, 0x6e, 0x88, 0xce, 0xa0, 0x96, 0xd2, 0x64, 0xc5, 0x02, 0xea, 0xc7, 0x84, 0x67, 0x21,
	0x2a, 0xb8, 0x9a, 0x6b, 0x23, 0xc2, 0xe9, 0x36, 0xa2, 0x73, 0x16, 0x9f, 0x21, 0xaa, 0xdc, 0xf6,
	0xaf, 0x02, 0x20, 0x15, 0x75, 0xac, 0x0e, 0x2f, 0x10, 0xf3, 0x9e, 0x88, 0x23, 0x36, 0x43, 0x7d,
	0x30, 0xe9, 0x9c, 0x06, 0xea, 0x5c, 0x7d, 0xc9, 0x38, 0x15, 0x4b, 0xa9, 0x2b, 0xa8, 0x5e, 0x1e,
	0x5b, 0x59, 0x03, 0xac, 0x4d, 0x03, 0xac, 0x7e, 0xde, 0x00, 0xdc, 0xd8, 0x58, 0x26, 0x99, 0x03,
	0x7d, 0x01, 0x74, 0x4b, 0x49, 0x22, 0xa7, 0x94, 0x48, 0x9f, 0xc5, 0x92, 0x26, 0x2b, 0x32, 0xd7,
	0x85, 0xfe, 0x77, 0x9d, 0xc3, 0x47, 0x93, 0x9b, 0x7b, 0xd0, 0x27, 0xd8, 0x4f, 0xa5, 0x48, 0xc8,
	0x2c, 0x0b, 0x51, 0xbd, 0x3c, 0x7d, 0xd1, 0x2c, 0x2f, 0x9b, 0xcf, 0x02, 0xe0, 0x0d, 0x8e, 0x6c,
	0x80, 0x40, 0xf0, 0x05, 0xd1, 0x85, 0x35, 0x4b, 0xda, 0x7c, 0xf6, 0xc2, 0xdc, 0x7b, 0x44, 0x72,
	0xff, 0x96, 0xa9, 0xfd, 0xc7, 0x80, 0xc3, 0x17, 0x3b, 0xa0, 0x13, 0xa8, 0x84, 0x2c, 0xa1, 0x81,
	0x14, 0xc9, 0x3a, 0xbf, 0x59, 0x4f, 0x02, 0xea, 0xc2, 0xee, 0x9c, 0xae, 0xe8, 0x3c, 0xbf, 0x5b,
	0xc7, 0xcf, 0x76, 0xcc, 0x17, 0x1a, 0x28, 0x00, 0x67, 0x1c, 0x3a, 0x87, 0x3a, 0x27, 0xf7, 0x3e,
	0x8d, 0x65, 0xb2, 0xf6, 0x53, 0xf6, 0x90, 0x05, 0x3d, 0xc0, 0x35, 0x4e, 0xee, 0x1d, 0x25, 0x7a,
	0xec, 0x21, 0xef, 0xe8, 0x8c, 0xd3, 0x58, 0x66, 0x4c, 0x49, 0x33, 0xd5, 0x5c, 0xd3, 0xc8, 0x3b,
	0x68, 0x44, 0xf3, 0x65, 0x7a, 0xeb, 0x8b, 0xd8, 0x0f, 0x04, 0xe7, 0x4c, 0x36, 0x77, 0x5b, 0x46,
	0xa7, 0x8c, 0x0f, 0xb4, 0x7c, 0x13, 0xf7, 0xb4, 0xd8, 0xfe, 0x69, 0xc0, 0xab, 0x7f, 0x45, 0x57,
	0x77, 0x3c, 0x5c, 0xc7, 0x84, 0xb3, 0x40, 0xc7, 0x2a, 0xe3, 0xcd, 0x10, 0x75, 0xc0, 0x8c, 0x12,
	0x4a, 0xfd, 0x90, 0xa5, 0x77, 0xfe, 0x74, 0x19, 0x45, 0x34, 0xd1, 0xf9, 0x0a, 0xb8, 0xae, 0xf4,
	0x3e, 0x4b, 0xef, 0xae, 0xb4, 0x8a, 0xde, 0x03, 0xd2, 0x24, 0xa7, 0x5c, 0x24, 0xeb, 0x0d, 0x5b,
	0xd4, 0xac, 0x5e, 0x63, 0xa8, 0x27, 0x32, 0xfa, 0xe2, 0x1a, 0x1a, 0x98, 0x92, 0xb0, 0x27, 0xe2,
	0x94, 0xa5, 0x92, 0xc6, 0xc1, 0x5a, 0xbd, 0x15, 0xcf, 0xf9, 0xfa, 0xcd, 0x19, 0x4d, 0x5c, 0x7b,
	0x60, 0xee, 0xa0, 0x23, 0x40, 0x03, 0x77, 0xe4, 0xd8, 0xd8, 0xfd, 0xa1, 0x5e, 0x8f, 0x3f, 0x70,
	0x6c, 0x4f, 0xbd, 0x29, 0x13, 0x6a, 0xdb, 0xba, 0x59, 0xb8, 0x38, 0x87, 0xda, 0xf6, 0xf9, 0xa2,
	0x32, 0x94, 0xfa, 0xae, 0x77, 0x6d, 0xee, 0xa8, 0xf7, 0x36, 0xb4, 0xc7, 0x63, 0xa7, 0x6f, 0x1a,
	0xd3, 0x3d, 0x7d, 0xed, 0x3e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xd6, 0xb4, 0x5f, 0x74,
	0x04, 0x00, 0x00,
}