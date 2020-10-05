// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_diff_task.proto

package atlas

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

type DiffTask struct {
	A                    *S3Address `protobuf:"bytes,100,opt,name=a,proto3" json:"a,omitempty"`
	B                    *S3Address `protobuf:"bytes,200,opt,name=b,proto3" json:"b,omitempty"`
	Meta                 *S3Address `protobuf:"bytes,300,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DiffTask) Reset()         { *m = DiffTask{} }
func (m *DiffTask) String() string { return proto.CompactTextString(m) }
func (*DiffTask) ProtoMessage()    {}
func (*DiffTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0411a244925965, []int{0}
}

func (m *DiffTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiffTask.Unmarshal(m, b)
}
func (m *DiffTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiffTask.Marshal(b, m, deterministic)
}
func (m *DiffTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiffTask.Merge(m, src)
}
func (m *DiffTask) XXX_Size() int {
	return xxx_messageInfo_DiffTask.Size(m)
}
func (m *DiffTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DiffTask.DiscardUnknown(m)
}

var xxx_messageInfo_DiffTask proto.InternalMessageInfo

func (m *DiffTask) GetA() *S3Address {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *DiffTask) GetB() *S3Address {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *DiffTask) GetMeta() *S3Address {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*DiffTask)(nil), "atlas.DiffTask")
}

func init() { proto.RegisterFile("type_diff_task.proto", fileDescriptor_1e0411a244925965) }

var fileDescriptor_1e0411a244925965 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xa9, 0x2c, 0x48,
	0x8d, 0x4f, 0xc9, 0x4c, 0x4b, 0x8b, 0x2f, 0x49, 0x2c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0x12, 0x05, 0x4b, 0x16, 0x1b, 0xc7, 0x27, 0xa6,
	0xa4, 0x14, 0xa5, 0x16, 0x17, 0x43, 0x64, 0x95, 0x8a, 0xb8, 0x38, 0x5c, 0x32, 0xd3, 0xd2, 0x42,
	0x12, 0x8b, 0xb3, 0x85, 0xe4, 0xb8, 0x18, 0x13, 0x25, 0x52, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x04,
	0xf4, 0xc0, 0xba, 0xf4, 0x82, 0x8d, 0x1d, 0x21, 0xca, 0x83, 0x18, 0x13, 0x85, 0xe4, 0xb9, 0x18,
	0x93, 0x24, 0x4e, 0x30, 0xe2, 0x52, 0x90, 0x24, 0xa4, 0xca, 0xc5, 0x92, 0x9b, 0x5a, 0x92, 0x28,
	0xb1, 0x86, 0x09, 0x87, 0x1a, 0xb0, 0x74, 0x12, 0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x29, 0x15, 0x7e, 0x93, 0xb0, 0x00, 0x00, 0x00,
}
