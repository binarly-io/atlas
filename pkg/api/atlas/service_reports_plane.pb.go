// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_reports_plane.proto

package atlas

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("service_reports_plane.proto", fileDescriptor_e5e5f15723b7379f)
}

var fileDescriptor_e5e5f15723b7379f = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x29, 0x8e, 0x2f, 0xc8, 0x49, 0xcc,
	0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0x92,
	0x2c, 0xa9, 0x2c, 0x80, 0x29, 0x88, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x81, 0xa8, 0x90,
	0x12, 0x47, 0x95, 0x2a, 0xc8, 0xa9, 0x84, 0x48, 0x18, 0xb9, 0x72, 0xf1, 0x04, 0x41, 0x4c, 0x0c,
	0x00, 0x19, 0x28, 0x64, 0xca, 0xc5, 0x0e, 0xe5, 0x0b, 0x89, 0xe8, 0x81, 0x8d, 0xd5, 0x83, 0xf0,
	0x83, 0x20, 0xe6, 0x49, 0x09, 0xa1, 0x89, 0x16, 0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x4d,
	0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xf6, 0x7a, 0xdf, 0xa7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportsPlaneClient is the client API for ReportsPlane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportsPlaneClient interface {
	Reports(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error)
}

type reportsPlaneClient struct {
	cc grpc.ClientConnInterface
}

func NewReportsPlaneClient(cc grpc.ClientConnInterface) ReportsPlaneClient {
	return &reportsPlaneClient{cc}
}

func (c *reportsPlaneClient) Reports(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error) {
	out := new(ReportReply)
	err := c.cc.Invoke(ctx, "/atlas.ReportsPlane/Reports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportsPlaneServer is the server API for ReportsPlane service.
type ReportsPlaneServer interface {
	Reports(context.Context, *ReportRequest) (*ReportReply, error)
}

// UnimplementedReportsPlaneServer can be embedded to have forward compatible implementations.
type UnimplementedReportsPlaneServer struct {
}

func (*UnimplementedReportsPlaneServer) Reports(ctx context.Context, req *ReportRequest) (*ReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reports not implemented")
}

func RegisterReportsPlaneServer(s *grpc.Server, srv ReportsPlaneServer) {
	s.RegisterService(&_ReportsPlane_serviceDesc, srv)
}

func _ReportsPlane_Reports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsPlaneServer).Reports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.ReportsPlane/Reports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsPlaneServer).Reports(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportsPlane_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atlas.ReportsPlane",
	HandlerType: (*ReportsPlaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reports",
			Handler:    _ReportsPlane_Reports_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_reports_plane.proto",
}
