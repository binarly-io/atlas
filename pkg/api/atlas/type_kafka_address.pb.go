// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_kafka_address.proto

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

type KafkaAddress struct {
	Topic                string   `protobuf:"bytes,100,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition            int32    `protobuf:"varint,200,opt,name=partition,proto3" json:"partition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaAddress) Reset()         { *m = KafkaAddress{} }
func (m *KafkaAddress) String() string { return proto.CompactTextString(m) }
func (*KafkaAddress) ProtoMessage()    {}
func (*KafkaAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_171407632322b2b7, []int{0}
}

func (m *KafkaAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaAddress.Unmarshal(m, b)
}
func (m *KafkaAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaAddress.Marshal(b, m, deterministic)
}
func (m *KafkaAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaAddress.Merge(m, src)
}
func (m *KafkaAddress) XXX_Size() int {
	return xxx_messageInfo_KafkaAddress.Size(m)
}
func (m *KafkaAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaAddress.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaAddress proto.InternalMessageInfo

func (m *KafkaAddress) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *KafkaAddress) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func init() {
	proto.RegisterType((*KafkaAddress)(nil), "atlas.KafkaAddress")
}

func init() { proto.RegisterFile("type_kafka_address.proto", fileDescriptor_171407632322b2b7) }

var fileDescriptor_171407632322b2b7 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xa9, 0x2c, 0x48,
	0x8d, 0xcf, 0x4e, 0x4c, 0xcb, 0x4e, 0x8c, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x56, 0x72, 0xe6, 0xe2, 0xf1,
	0x06, 0xc9, 0x3a, 0x42, 0x24, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0xf2, 0x0b, 0x32, 0x93, 0x25, 0x52,
	0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x59, 0x2e, 0xce, 0x82, 0xc4, 0xa2, 0x92, 0xcc,
	0x92, 0xcc, 0xfc, 0x3c, 0x89, 0x13, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x08, 0x91, 0x24, 0x36,
	0xb0, 0x91, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xa5, 0x6f, 0x41, 0x6e, 0x00, 0x00,
	0x00,
}
