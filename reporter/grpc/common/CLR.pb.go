// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/CLR.proto

package common

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

type CLRMetric struct {
	Time                 int64      `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu                  *CPU       `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gc                   *ClrGC     `protobuf:"bytes,3,opt,name=gc,proto3" json:"gc,omitempty"`
	Thread               *ClrThread `protobuf:"bytes,4,opt,name=thread,proto3" json:"thread,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CLRMetric) Reset()         { *m = CLRMetric{} }
func (m *CLRMetric) String() string { return proto.CompactTextString(m) }
func (*CLRMetric) ProtoMessage()    {}
func (*CLRMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{0}
}

func (m *CLRMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CLRMetric.Unmarshal(m, b)
}
func (m *CLRMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CLRMetric.Marshal(b, m, deterministic)
}
func (m *CLRMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CLRMetric.Merge(m, src)
}
func (m *CLRMetric) XXX_Size() int {
	return xxx_messageInfo_CLRMetric.Size(m)
}
func (m *CLRMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_CLRMetric.DiscardUnknown(m)
}

var xxx_messageInfo_CLRMetric proto.InternalMessageInfo

func (m *CLRMetric) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *CLRMetric) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *CLRMetric) GetGc() *ClrGC {
	if m != nil {
		return m.Gc
	}
	return nil
}

func (m *CLRMetric) GetThread() *ClrThread {
	if m != nil {
		return m.Thread
	}
	return nil
}

type ClrGC struct {
	Gen0CollectCount     int64    `protobuf:"varint,1,opt,name=Gen0CollectCount,proto3" json:"Gen0CollectCount,omitempty"`
	Gen1CollectCount     int64    `protobuf:"varint,2,opt,name=Gen1CollectCount,proto3" json:"Gen1CollectCount,omitempty"`
	Gen2CollectCount     int64    `protobuf:"varint,3,opt,name=Gen2CollectCount,proto3" json:"Gen2CollectCount,omitempty"`
	HeapMemory           int64    `protobuf:"varint,4,opt,name=HeapMemory,proto3" json:"HeapMemory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClrGC) Reset()         { *m = ClrGC{} }
func (m *ClrGC) String() string { return proto.CompactTextString(m) }
func (*ClrGC) ProtoMessage()    {}
func (*ClrGC) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{1}
}

func (m *ClrGC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClrGC.Unmarshal(m, b)
}
func (m *ClrGC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClrGC.Marshal(b, m, deterministic)
}
func (m *ClrGC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClrGC.Merge(m, src)
}
func (m *ClrGC) XXX_Size() int {
	return xxx_messageInfo_ClrGC.Size(m)
}
func (m *ClrGC) XXX_DiscardUnknown() {
	xxx_messageInfo_ClrGC.DiscardUnknown(m)
}

var xxx_messageInfo_ClrGC proto.InternalMessageInfo

func (m *ClrGC) GetGen0CollectCount() int64 {
	if m != nil {
		return m.Gen0CollectCount
	}
	return 0
}

func (m *ClrGC) GetGen1CollectCount() int64 {
	if m != nil {
		return m.Gen1CollectCount
	}
	return 0
}

func (m *ClrGC) GetGen2CollectCount() int64 {
	if m != nil {
		return m.Gen2CollectCount
	}
	return 0
}

func (m *ClrGC) GetHeapMemory() int64 {
	if m != nil {
		return m.HeapMemory
	}
	return 0
}

type ClrThread struct {
	AvailableCompletionPortThreads int32    `protobuf:"varint,1,opt,name=AvailableCompletionPortThreads,proto3" json:"AvailableCompletionPortThreads,omitempty"`
	AvailableWorkerThreads         int32    `protobuf:"varint,2,opt,name=AvailableWorkerThreads,proto3" json:"AvailableWorkerThreads,omitempty"`
	MaxCompletionPortThreads       int32    `protobuf:"varint,3,opt,name=MaxCompletionPortThreads,proto3" json:"MaxCompletionPortThreads,omitempty"`
	MaxWorkerThreads               int32    `protobuf:"varint,4,opt,name=MaxWorkerThreads,proto3" json:"MaxWorkerThreads,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *ClrThread) Reset()         { *m = ClrThread{} }
func (m *ClrThread) String() string { return proto.CompactTextString(m) }
func (*ClrThread) ProtoMessage()    {}
func (*ClrThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{2}
}

func (m *ClrThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClrThread.Unmarshal(m, b)
}
func (m *ClrThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClrThread.Marshal(b, m, deterministic)
}
func (m *ClrThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClrThread.Merge(m, src)
}
func (m *ClrThread) XXX_Size() int {
	return xxx_messageInfo_ClrThread.Size(m)
}
func (m *ClrThread) XXX_DiscardUnknown() {
	xxx_messageInfo_ClrThread.DiscardUnknown(m)
}

var xxx_messageInfo_ClrThread proto.InternalMessageInfo

func (m *ClrThread) GetAvailableCompletionPortThreads() int32 {
	if m != nil {
		return m.AvailableCompletionPortThreads
	}
	return 0
}

func (m *ClrThread) GetAvailableWorkerThreads() int32 {
	if m != nil {
		return m.AvailableWorkerThreads
	}
	return 0
}

func (m *ClrThread) GetMaxCompletionPortThreads() int32 {
	if m != nil {
		return m.MaxCompletionPortThreads
	}
	return 0
}

func (m *ClrThread) GetMaxWorkerThreads() int32 {
	if m != nil {
		return m.MaxWorkerThreads
	}
	return 0
}

func init() {
	proto.RegisterType((*CLRMetric)(nil), "CLRMetric")
	proto.RegisterType((*ClrGC)(nil), "ClrGC")
	proto.RegisterType((*ClrThread)(nil), "ClrThread")
}

func init() {
	proto.RegisterFile("common/CLR.proto", fileDescriptor_a10d56830892247a)
}

var fileDescriptor_a10d56830892247a = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xb1, 0x9d, 0x06, 0xfa, 0x76, 0x29, 0x1a, 0x04, 0xd3, 0x43, 0x29, 0x39, 0x95, 0xc1,
	0xe4, 0x36, 0x1b, 0x3b, 0xec, 0xb6, 0x69, 0xac, 0x3b, 0x24, 0x23, 0x68, 0x1b, 0x85, 0xdd, 0x14,
	0xf5, 0xa1, 0x18, 0xcb, 0x7a, 0x46, 0x96, 0xd7, 0xe6, 0x13, 0xec, 0xbb, 0x6c, 0x1f, 0x6f, 0x5f,
	0x60, 0x58, 0xf1, 0xc2, 0x42, 0x1b, 0x7a, 0xb2, 0xf9, 0xfd, 0x7f, 0x7f, 0x3d, 0xd9, 0x3c, 0x38,
	0xd1, 0x54, 0xd7, 0xe4, 0x0a, 0x31, 0x97, 0xbc, 0xf1, 0x14, 0xe8, 0xf4, 0xf9, 0x40, 0xb6, 0x8f,
	0x2d, 0x9c, 0xb6, 0x70, 0x2c, 0xe6, 0x72, 0x81, 0xc1, 0x97, 0x9a, 0x31, 0x18, 0x85, 0xb2, 0xc6,
	0x3c, 0x39, 0x4f, 0x2e, 0x32, 0x19, 0xdf, 0xd9, 0x04, 0x32, 0xdd, 0x74, 0x79, 0x7a, 0x9e, 0x5c,
	0x3c, 0x9b, 0x8d, 0xb8, 0x58, 0x7e, 0x93, 0x3d, 0x60, 0x13, 0x48, 0x8d, 0xce, 0xb3, 0x88, 0xc7,
	0x5c, 0x58, 0x7f, 0x2d, 0x64, 0x6a, 0x34, 0x9b, 0xc2, 0x38, 0xac, 0x3d, 0xaa, 0xdb, 0x7c, 0x14,
	0x33, 0xe8, 0xb3, 0xaf, 0x91, 0xc8, 0x21, 0x99, 0xfe, 0x4e, 0xe0, 0x28, 0x36, 0xd8, 0x0b, 0x38,
	0xb9, 0x46, 0x77, 0x29, 0xc8, 0x5a, 0xd4, 0x41, 0x50, 0xe7, 0xc2, 0x30, 0xfd, 0x01, 0x1f, 0xdc,
	0xab, 0x3d, 0x37, 0xdd, 0xb9, 0x57, 0x8f, 0xb8, 0xb3, 0x3d, 0x37, 0xdb, 0xb9, 0x7b, 0x9c, 0x9d,
	0x01, 0x7c, 0x42, 0xd5, 0x2c, 0xb0, 0x26, 0xbf, 0x89, 0xb7, 0xce, 0xe4, 0x7f, 0x64, 0xfa, 0x27,
	0x81, 0xe3, 0xdd, 0x37, 0xb0, 0x8f, 0x70, 0xf6, 0xee, 0x87, 0x2a, 0xad, 0x5a, 0x59, 0x14, 0x54,
	0x37, 0x16, 0x43, 0x49, 0x6e, 0x49, 0x3e, 0x6c, 0x85, 0x36, 0xde, 0xff, 0x48, 0x3e, 0x61, 0xb1,
	0x37, 0x30, 0xd9, 0x19, 0x37, 0xe4, 0x2b, 0xf4, 0xff, 0xfa, 0x69, 0xec, 0x1f, 0x48, 0xd9, 0x5b,
	0xc8, 0x17, 0xea, 0xfe, 0xf1, 0xc9, 0x59, 0x6c, 0x1e, 0xcc, 0xfb, 0xbf, 0xb2, 0x50, 0xf7, 0xfb,
	0xd3, 0x46, 0xb1, 0xf3, 0x80, 0xbf, 0xff, 0x99, 0xc0, 0x25, 0x79, 0xc3, 0x55, 0xa3, 0xf4, 0x1a,
	0x79, 0x5b, 0x6d, 0xee, 0x94, 0xad, 0x4a, 0xd7, 0x93, 0x9a, 0x3b, 0x0c, 0x77, 0xe4, 0x2b, 0x6e,
	0x95, 0x33, 0x9d, 0x32, 0xc8, 0x95, 0x41, 0x17, 0x96, 0xc9, 0xf7, 0xd7, 0xa6, 0x0c, 0xeb, 0x6e,
	0xc5, 0x35, 0xd5, 0xc5, 0x07, 0x45, 0xc2, 0x52, 0x77, 0xfb, 0x72, 0xae, 0x56, 0x6d, 0x61, 0x68,
	0xd6, 0x56, 0x9b, 0xc2, 0x63, 0x43, 0x3e, 0xa0, 0x2f, 0x8c, 0x6f, 0xf4, 0xb0, 0x8b, 0xbf, 0xd2,
	0xd3, 0x2f, 0xd5, 0xe6, 0x66, 0x38, 0xff, 0xf3, 0xf6, 0xec, 0x65, 0xbf, 0x9f, 0x9a, 0xec, 0x6a,
	0x1c, 0x37, 0xf5, 0xd5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xb0, 0x71, 0x61, 0xd2, 0x02,
	0x00, 0x00,
}
