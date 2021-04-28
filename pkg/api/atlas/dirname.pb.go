// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dirname.proto

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

// Dirname represents abstract directory name.
type Dirname struct {
	// Dirname
	Dirname              string   `protobuf:"bytes,100,opt,name=dirname,proto3" json:"dirname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dirname) Reset()      { *m = Dirname{} }
func (*Dirname) ProtoMessage() {}
func (*Dirname) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e715b44ad67bc1b, []int{0}
}

func (m *Dirname) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dirname.Unmarshal(m, b)
}
func (m *Dirname) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dirname.Marshal(b, m, deterministic)
}
func (m *Dirname) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dirname.Merge(m, src)
}
func (m *Dirname) XXX_Size() int {
	return xxx_messageInfo_Dirname.Size(m)
}
func (m *Dirname) XXX_DiscardUnknown() {
	xxx_messageInfo_Dirname.DiscardUnknown(m)
}

var xxx_messageInfo_Dirname proto.InternalMessageInfo

func (m *Dirname) GetDirname() string {
	if m != nil {
		return m.Dirname
	}
	return ""
}

func init() {
	proto.RegisterType((*Dirname)(nil), "atlas.Dirname")
}

func init() { proto.RegisterFile("dirname.proto", fileDescriptor_3e715b44ad67bc1b) }

var fileDescriptor_3e715b44ad67bc1b = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xc9, 0x2c, 0xca,
	0x4b, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c,
	0x56, 0x52, 0xe6, 0x62, 0x77, 0x81, 0x88, 0x0b, 0x49, 0x70, 0xb1, 0x43, 0x95, 0x48, 0xa4, 0x28,
	0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x49, 0x6c, 0x60, 0x2d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0xdb, 0x9b, 0x50, 0x43, 0x00, 0x00, 0x00,
}
