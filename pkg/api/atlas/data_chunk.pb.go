// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data_chunk.proto

//*
// DataChunk represents one chunk (block,single piece) of data send used by DataChunks() function in Control Plane

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

// DataChunk is a chunk of data transferred as a single piece.
// Can be part of bigger data, transferred by smaller chunks.
type DataChunk struct {
	// Header describes this chunk
	Header *Metadata `protobuf:"bytes,100,opt,name=header,proto3" json:"header,omitempty"`
	// Data is the purpose of the whole data chunk type.
	// May contain any arbitrary sequence of bytes no longer than 2^32.
	Data []byte `protobuf:"bytes,200,opt,name=data,proto3" json:"data,omitempty"`
	// Types that are valid to be assigned to TransportMetadataOptional:
	//	*DataChunk_TransportMetadata
	TransportMetadataOptional isDataChunk_TransportMetadataOptional `protobuf_oneof:"transport_metadata_optional"`
	// Types that are valid to be assigned to PayloadMetadataOptional:
	//	*DataChunk_PayloadMetadata
	PayloadMetadataOptional isDataChunk_PayloadMetadataOptional `protobuf_oneof:"payload_metadata_optional"`
	XXX_NoUnkeyedLiteral    struct{}                            `json:"-"`
	XXX_unrecognized        []byte                              `json:"-"`
	XXX_sizecache           int32                               `json:"-"`
}

func (m *DataChunk) Reset()      { *m = DataChunk{} }
func (*DataChunk) ProtoMessage() {}
func (*DataChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39e17118bd74b9a, []int{0}
}

func (m *DataChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataChunk.Unmarshal(m, b)
}
func (m *DataChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataChunk.Marshal(b, m, deterministic)
}
func (m *DataChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataChunk.Merge(m, src)
}
func (m *DataChunk) XXX_Size() int {
	return xxx_messageInfo_DataChunk.Size(m)
}
func (m *DataChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_DataChunk.DiscardUnknown(m)
}

var xxx_messageInfo_DataChunk proto.InternalMessageInfo

func (m *DataChunk) GetHeader() *Metadata {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DataChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type isDataChunk_TransportMetadataOptional interface {
	isDataChunk_TransportMetadataOptional()
}

type DataChunk_TransportMetadata struct {
	TransportMetadata *Metadata `protobuf:"bytes,300,opt,name=transport_metadata,json=transportMetadata,proto3,oneof"`
}

func (*DataChunk_TransportMetadata) isDataChunk_TransportMetadataOptional() {}

func (m *DataChunk) GetTransportMetadataOptional() isDataChunk_TransportMetadataOptional {
	if m != nil {
		return m.TransportMetadataOptional
	}
	return nil
}

func (m *DataChunk) GetTransportMetadata() *Metadata {
	if x, ok := m.GetTransportMetadataOptional().(*DataChunk_TransportMetadata); ok {
		return x.TransportMetadata
	}
	return nil
}

type isDataChunk_PayloadMetadataOptional interface {
	isDataChunk_PayloadMetadataOptional()
}

type DataChunk_PayloadMetadata struct {
	PayloadMetadata *Metadata `protobuf:"bytes,400,opt,name=payload_metadata,json=payloadMetadata,proto3,oneof"`
}

func (*DataChunk_PayloadMetadata) isDataChunk_PayloadMetadataOptional() {}

func (m *DataChunk) GetPayloadMetadataOptional() isDataChunk_PayloadMetadataOptional {
	if m != nil {
		return m.PayloadMetadataOptional
	}
	return nil
}

func (m *DataChunk) GetPayloadMetadata() *Metadata {
	if x, ok := m.GetPayloadMetadataOptional().(*DataChunk_PayloadMetadata); ok {
		return x.PayloadMetadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DataChunk) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DataChunk_TransportMetadata)(nil),
		(*DataChunk_PayloadMetadata)(nil),
	}
}

func init() {
	proto.RegisterType((*DataChunk)(nil), "atlas.DataChunk")
}

func init() { proto.RegisterFile("data_chunk.proto", fileDescriptor_d39e17118bd74b9a) }

var fileDescriptor_d39e17118bd74b9a = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0x2c, 0x49,
	0x8c, 0x4f, 0xce, 0x28, 0xcd, 0xcb, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c,
	0xc9, 0x49, 0x2c, 0x96, 0xe2, 0xcb, 0x4d, 0x2d, 0x49, 0x04, 0x49, 0x42, 0x84, 0x95, 0xda, 0x99,
	0xb8, 0x38, 0x5d, 0x12, 0x4b, 0x12, 0x9d, 0x41, 0x4a, 0x85, 0xd4, 0xb9, 0xd8, 0x32, 0x52, 0x13,
	0x53, 0x52, 0x8b, 0x24, 0x52, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xf8, 0xf5, 0xc0, 0xba, 0xf4, 0x7c,
	0xa1, 0x9a, 0x82, 0xa0, 0xd2, 0x42, 0xc2, 0x5c, 0x2c, 0x20, 0xbe, 0xc4, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x9e, 0x20, 0x30, 0x47, 0xc8, 0x91, 0x4b, 0xa8, 0xa4, 0x28, 0x31, 0xaf, 0xb8, 0x20, 0xbf,
	0xa8, 0x24, 0x1e, 0x66, 0x8f, 0xc4, 0x1a, 0x26, 0xac, 0x46, 0x79, 0x30, 0x04, 0x09, 0xc2, 0x55,
	0xc3, 0x04, 0x85, 0x6c, 0xb9, 0x04, 0x0a, 0x12, 0x2b, 0x73, 0xf2, 0x13, 0x53, 0x10, 0x06, 0x4c,
	0x60, 0xc6, 0x6e, 0x00, 0x63, 0x10, 0x3f, 0x54, 0x2d, 0x4c, 0xc8, 0x49, 0x96, 0x4b, 0x1a, 0xd3,
	0x05, 0xf1, 0xf9, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x89, 0x39, 0x4e, 0xd2, 0x5c, 0x92, 0xe8, 0xa6,
	0xc3, 0x25, 0x93, 0xd8, 0xc0, 0x01, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x8a, 0x75,
	0xe4, 0x3b, 0x01, 0x00, 0x00,
}