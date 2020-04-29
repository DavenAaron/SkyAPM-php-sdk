// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/CLR.proto

package v1

import (
	fmt "fmt"
	v1 "github.com/SkyAPM/SkyAPM-php-sdk/reporter/network/common/v1"
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
	Cpu                  *v1.CPU    `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
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

func (m *CLRMetric) GetCpu() *v1.CPU {
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
	proto.RegisterType((*CLRMetric)(nil), "skywalking.network.protocol.common.v1.CLRMetric")
	proto.RegisterType((*ClrGC)(nil), "skywalking.network.protocol.common.v1.ClrGC")
	proto.RegisterType((*ClrThread)(nil), "skywalking.network.protocol.common.v1.ClrThread")
}

func init() { proto.RegisterFile("common/CLR.proto", fileDescriptor_a10d56830892247a) }

var fileDescriptor_a10d56830892247a = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0xea, 0xda, 0x30,
	0x14, 0xc7, 0x69, 0xeb, 0x4f, 0x58, 0x76, 0x23, 0x19, 0x8c, 0xe2, 0x85, 0x88, 0x6c, 0x20, 0x32,
	0x52, 0xeb, 0x60, 0x17, 0xc3, 0x9b, 0xad, 0x30, 0xbd, 0xb0, 0xa3, 0x64, 0x1b, 0xc2, 0xee, 0x62,
	0x16, 0x6a, 0x69, 0x9a, 0x94, 0x18, 0xff, 0xbd, 0xd2, 0xf6, 0x5c, 0x7b, 0x82, 0xbd, 0xc0, 0x30,
	0x8d, 0x9d, 0x45, 0x65, 0xde, 0xc9, 0xf7, 0x7c, 0x3e, 0xe7, 0x78, 0x7a, 0x02, 0x3a, 0x54, 0x16,
	0x85, 0x14, 0x41, 0xb4, 0xc0, 0xa8, 0x54, 0x52, 0x4b, 0xf8, 0x7a, 0x93, 0x1f, 0xf7, 0x84, 0xe7,
	0x99, 0x48, 0x91, 0x60, 0x7a, 0x2f, 0x55, 0x5e, 0x55, 0xa8, 0xe4, 0xa8, 0xa2, 0xd1, 0x2e, 0xec,
	0xbe, 0xb0, 0xa2, 0x4d, 0x0c, 0x31, 0xf8, 0xed, 0x80, 0x67, 0xd1, 0x02, 0xc7, 0x4c, 0xab, 0x8c,
	0x42, 0x08, 0x5a, 0x3a, 0x2b, 0x98, 0xef, 0xf4, 0x9d, 0xa1, 0x87, 0xcd, 0x6f, 0x38, 0x05, 0x1e,
	0x2d, 0xb7, 0xbe, 0xdb, 0x77, 0x86, 0xcf, 0x27, 0x23, 0xf4, 0xd0, 0x2c, 0x14, 0x25, 0xdf, 0xf0,
	0x49, 0x83, 0x53, 0xe0, 0xa6, 0xd4, 0xf7, 0x8c, 0xfc, 0xe6, 0x51, 0x99, 0xab, 0x59, 0x84, 0xdd,
	0x94, 0xc2, 0x39, 0x68, 0xeb, 0xb5, 0x62, 0xe4, 0x87, 0xdf, 0x32, 0x1d, 0xc6, 0x8f, 0x77, 0xf8,
	0x6a, 0x3c, 0x6c, 0xfd, 0xc1, 0x2f, 0x07, 0x3c, 0x99, 0xbe, 0x70, 0x04, 0x3a, 0x33, 0x26, 0xc6,
	0x91, 0xe4, 0x9c, 0x51, 0x1d, 0xc9, 0xad, 0xd0, 0x76, 0xdf, 0xab, 0xdc, 0xb2, 0x61, 0x83, 0x75,
	0x6b, 0x36, 0xbc, 0xc1, 0x4e, 0x1a, 0xac, 0x57, 0xb3, 0x8d, 0x1c, 0xf6, 0x00, 0x98, 0x33, 0x52,
	0xc6, 0xac, 0x90, 0xea, 0x68, 0x76, 0xf3, 0xf0, 0x45, 0x32, 0xf8, 0x73, 0xba, 0xca, 0x79, 0x07,
	0xf8, 0x09, 0xf4, 0x3e, 0xec, 0x48, 0xc6, 0xc9, 0x8a, 0xb3, 0x48, 0x16, 0x25, 0x67, 0x3a, 0x93,
	0x22, 0x91, 0x4a, 0x57, 0xc0, 0xc6, 0xfc, 0xff, 0x27, 0xfc, 0x1f, 0x0a, 0xbe, 0x03, 0x2f, 0x6b,
	0x62, 0x29, 0x55, 0xce, 0xd4, 0xd9, 0x77, 0x8d, 0x7f, 0xa7, 0x0a, 0xdf, 0x03, 0x3f, 0x26, 0x87,
	0xdb, 0x93, 0x3d, 0x63, 0xde, 0xad, 0x9f, 0xbe, 0x4a, 0x4c, 0x0e, 0xcd, 0x69, 0x2d, 0xe3, 0x5c,
	0xe5, 0x1f, 0xf7, 0x60, 0x2c, 0x55, 0x8a, 0x48, 0x49, 0xe8, 0x9a, 0x5d, 0x5e, 0x9a, 0x94, 0x45,
	0x7d, 0x6d, 0x4e, 0x44, 0xba, 0x25, 0x29, 0x43, 0x24, 0x65, 0x42, 0x27, 0xce, 0xf7, 0x57, 0xff,
	0xc0, 0xc0, 0x42, 0xc1, 0x19, 0x0a, 0x0c, 0x14, 0xec, 0xc2, 0x9f, 0x6e, 0xf7, 0x4b, 0x7e, 0x5c,
	0xda, 0x7e, 0x9f, 0x2b, 0x2c, 0xb1, 0x0f, 0x67, 0xd5, 0x36, 0x4f, 0xe8, 0xed, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x30, 0x45, 0x65, 0x5b, 0x03, 0x00, 0x00,
}