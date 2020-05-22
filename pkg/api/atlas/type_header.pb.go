// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_header.proto

package atlas

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Header describes header or the object
type Header struct {
	// Type + Name pair
	TypeName *TypeName `protobuf:"bytes,100,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// Types that are valid to be assigned to VersionOptional:
	//	*Header_Version
	VersionOptional isHeader_VersionOptional `protobuf_oneof:"version_optional"`
	// Types that are valid to be assigned to UuidOptional:
	//	*Header_Uuid
	UuidOptional isHeader_UuidOptional `protobuf_oneof:"uuid_optional"`
	// Types that are valid to be assigned to UuidReferenceOptional:
	//	*Header_UuidReference
	UuidReferenceOptional isHeader_UuidReferenceOptional `protobuf_oneof:"uuid_reference_optional"`
	// Types that are valid to be assigned to TimestampOptional:
	//	*Header_Ts
	TimestampOptional isHeader_TimestampOptional `protobuf_oneof:"timestamp_optional"`
	// Types that are valid to be assigned to DescriptionOptional:
	//	*Header_Description
	DescriptionOptional  isHeader_DescriptionOptional `protobuf_oneof:"description_optional"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f5b5f06ff97e4c, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetTypeName() *TypeName {
	if m != nil {
		return m.TypeName
	}
	return nil
}

type isHeader_VersionOptional interface {
	isHeader_VersionOptional()
}

type Header_Version struct {
	Version int32 `protobuf:"varint,200,opt,name=version,proto3,oneof"`
}

func (*Header_Version) isHeader_VersionOptional() {}

func (m *Header) GetVersionOptional() isHeader_VersionOptional {
	if m != nil {
		return m.VersionOptional
	}
	return nil
}

func (m *Header) GetVersion() int32 {
	if x, ok := m.GetVersionOptional().(*Header_Version); ok {
		return x.Version
	}
	return 0
}

type isHeader_UuidOptional interface {
	isHeader_UuidOptional()
}

type Header_Uuid struct {
	Uuid *UUID `protobuf:"bytes,300,opt,name=uuid,proto3,oneof"`
}

func (*Header_Uuid) isHeader_UuidOptional() {}

func (m *Header) GetUuidOptional() isHeader_UuidOptional {
	if m != nil {
		return m.UuidOptional
	}
	return nil
}

func (m *Header) GetUuid() *UUID {
	if x, ok := m.GetUuidOptional().(*Header_Uuid); ok {
		return x.Uuid
	}
	return nil
}

type isHeader_UuidReferenceOptional interface {
	isHeader_UuidReferenceOptional()
}

type Header_UuidReference struct {
	UuidReference *UUID `protobuf:"bytes,400,opt,name=uuid_reference,json=uuidReference,proto3,oneof"`
}

func (*Header_UuidReference) isHeader_UuidReferenceOptional() {}

func (m *Header) GetUuidReferenceOptional() isHeader_UuidReferenceOptional {
	if m != nil {
		return m.UuidReferenceOptional
	}
	return nil
}

func (m *Header) GetUuidReference() *UUID {
	if x, ok := m.GetUuidReferenceOptional().(*Header_UuidReference); ok {
		return x.UuidReference
	}
	return nil
}

type isHeader_TimestampOptional interface {
	isHeader_TimestampOptional()
}

type Header_Ts struct {
	Ts *timestamp.Timestamp `protobuf:"bytes,500,opt,name=ts,proto3,oneof"`
}

func (*Header_Ts) isHeader_TimestampOptional() {}

func (m *Header) GetTimestampOptional() isHeader_TimestampOptional {
	if m != nil {
		return m.TimestampOptional
	}
	return nil
}

func (m *Header) GetTs() *timestamp.Timestamp {
	if x, ok := m.GetTimestampOptional().(*Header_Ts); ok {
		return x.Ts
	}
	return nil
}

type isHeader_DescriptionOptional interface {
	isHeader_DescriptionOptional()
}

type Header_Description struct {
	Description string `protobuf:"bytes,600,opt,name=description,proto3,oneof"`
}

func (*Header_Description) isHeader_DescriptionOptional() {}

func (m *Header) GetDescriptionOptional() isHeader_DescriptionOptional {
	if m != nil {
		return m.DescriptionOptional
	}
	return nil
}

func (m *Header) GetDescription() string {
	if x, ok := m.GetDescriptionOptional().(*Header_Description); ok {
		return x.Description
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Header) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Header_Version)(nil),
		(*Header_Uuid)(nil),
		(*Header_UuidReference)(nil),
		(*Header_Ts)(nil),
		(*Header_Description)(nil),
	}
}

func init() {
	proto.RegisterType((*Header)(nil), "atlas.Header")
}

func init() {
	proto.RegisterFile("type_header.proto", fileDescriptor_05f5b5f06ff97e4c)
}

var fileDescriptor_05f5b5f06ff97e4c = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4e, 0xf3, 0x30,
	0x10, 0xc5, 0xbf, 0xa4, 0x7f, 0x3e, 0x3a, 0x15, 0x14, 0x4c, 0x05, 0xa1, 0x2c, 0xa8, 0xca, 0x26,
	0x0b, 0x70, 0x25, 0x10, 0x17, 0xa8, 0x58, 0x84, 0x0d, 0x0b, 0xab, 0x5d, 0x47, 0x6e, 0x33, 0x2d,
	0x91, 0x92, 0x38, 0xb2, 0x1d, 0xa4, 0xde, 0x82, 0x83, 0x70, 0x10, 0x8e, 0xc0, 0x31, 0x58, 0x70,
	0x00, 0x64, 0x27, 0x69, 0x8a, 0xd8, 0xd8, 0xd6, 0x6f, 0x9e, 0xe7, 0x3d, 0x3d, 0x38, 0xd1, 0xdb,
	0x1c, 0xc3, 0x17, 0xe4, 0x11, 0x4a, 0x9a, 0x4b, 0xa1, 0x05, 0xe9, 0x70, 0x9d, 0x70, 0x35, 0xba,
	0xda, 0x08, 0xb1, 0x49, 0x70, 0x6a, 0xe1, 0xb2, 0x58, 0x4f, 0x75, 0x9c, 0xa2, 0xd2, 0x3c, 0xcd,
	0x4b, 0xdd, 0xe8, 0xd4, 0x7e, 0x35, 0x47, 0xc6, 0x53, 0xac, 0xe0, 0xc0, 0xc2, 0xa2, 0x88, 0xa3,
	0x12, 0x4c, 0xbe, 0x5c, 0xe8, 0x06, 0x76, 0x3d, 0xb9, 0x81, 0x9e, 0x9d, 0x1a, 0xb9, 0x17, 0x8d,
	0x1d, 0xbf, 0x7f, 0x37, 0xa0, 0xd6, 0x8c, 0xce, 0xb7, 0x39, 0x3e, 0xf3, 0x14, 0xd9, 0x81, 0xae,
	0x5e, 0xe4, 0x12, 0xfe, 0xbf, 0xa2, 0x54, 0xb1, 0xc8, 0xbc, 0x0f, 0x67, 0xec, 0xf8, 0x9d, 0xe0,
	0x1f, 0xab, 0x09, 0x99, 0x40, 0xdb, 0x78, 0x78, 0xef, 0xae, 0x5d, 0xd3, 0xaf, 0xd6, 0x2c, 0x16,
	0x4f, 0x8f, 0x81, 0xc3, 0xec, 0x8c, 0x3c, 0xc0, 0x91, 0xb9, 0x43, 0x89, 0x6b, 0x94, 0x98, 0xad,
	0xd0, 0x7b, 0x6b, 0xfd, 0x55, 0xbb, 0xec, 0xd0, 0xa8, 0x58, 0x2d, 0x22, 0xb7, 0xe0, 0x6a, 0xe5,
	0x7d, 0x97, 0xd2, 0x11, 0x2d, 0x5b, 0xa0, 0x75, 0x0b, 0x74, 0x5e, 0xb7, 0x10, 0xb4, 0x98, 0xab,
	0x15, 0xb9, 0x86, 0x7e, 0x84, 0x6a, 0x25, 0xe3, 0x5c, 0x9b, 0xa8, 0x9f, 0xed, 0xb1, 0xe3, 0xf7,
	0x82, 0x36, 0xdb, 0xa7, 0x33, 0x02, 0xc7, 0x55, 0xf2, 0x50, 0x58, 0xc2, 0x93, 0xd9, 0x00, 0xac,
	0x71, 0x03, 0x2e, 0xe0, 0xfc, 0x77, 0xde, 0x66, 0x34, 0x04, 0xb2, 0x6b, 0xbf, 0xa1, 0x67, 0x30,
	0xdc, 0x33, 0xd9, 0xf1, 0x65, 0xd7, 0x86, 0xbd, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x70,
	0x83, 0x69, 0xdc, 0x01, 0x00, 0x00,
}