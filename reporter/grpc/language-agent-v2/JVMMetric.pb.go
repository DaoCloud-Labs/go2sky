// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent-v2/JVMMetric.proto

package language_agent_v2

import (
	context "context"
	fmt "fmt"
	common "github.com/DaoCloud-Labs/go2sky/reporter/grpc/common"
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

type JVMMetricCollection struct {
	EnvCode              string              `protobuf:"bytes,1,opt,name=envCode,proto3" json:"envCode,omitempty"`
	Metrics              []*common.JVMMetric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	ServiceInstanceId    int32               `protobuf:"varint,3,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *JVMMetricCollection) Reset()         { *m = JVMMetricCollection{} }
func (m *JVMMetricCollection) String() string { return proto.CompactTextString(m) }
func (*JVMMetricCollection) ProtoMessage()    {}
func (*JVMMetricCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bd38fe036677f3, []int{0}
}

func (m *JVMMetricCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMMetricCollection.Unmarshal(m, b)
}
func (m *JVMMetricCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMMetricCollection.Marshal(b, m, deterministic)
}
func (m *JVMMetricCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMMetricCollection.Merge(m, src)
}
func (m *JVMMetricCollection) XXX_Size() int {
	return xxx_messageInfo_JVMMetricCollection.Size(m)
}
func (m *JVMMetricCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMMetricCollection.DiscardUnknown(m)
}

var xxx_messageInfo_JVMMetricCollection proto.InternalMessageInfo

func (m *JVMMetricCollection) GetEnvCode() string {
	if m != nil {
		return m.EnvCode
	}
	return ""
}

func (m *JVMMetricCollection) GetMetrics() []*common.JVMMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *JVMMetricCollection) GetServiceInstanceId() int32 {
	if m != nil {
		return m.ServiceInstanceId
	}
	return 0
}

func init() {
	proto.RegisterType((*JVMMetricCollection)(nil), "JVMMetricCollection")
}

func init() {
	proto.RegisterFile("language-agent-v2/JVMMetric.proto", fileDescriptor_a5bd38fe036677f3)
}

var fileDescriptor_a5bd38fe036677f3 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x6b, 0xc2, 0x30,
	0x18, 0x87, 0x57, 0x65, 0x13, 0xb3, 0xcb, 0x56, 0xc7, 0x28, 0x9e, 0x9c, 0xec, 0xe0, 0x41, 0x13,
	0xa8, 0x1f, 0x60, 0xb0, 0x8e, 0x81, 0x32, 0x87, 0x28, 0x38, 0xd8, 0x2d, 0xa6, 0x2f, 0xb1, 0xb4,
	0xc9, 0x5b, 0x92, 0x58, 0xf1, 0xbc, 0xcf, 0xb0, 0x2f, 0xb1, 0x4f, 0x39, 0x6c, 0xad, 0x3b, 0xb8,
	0x53, 0xc8, 0xf3, 0xfe, 0xfb, 0xf1, 0x90, 0x87, 0x8c, 0x6b, 0xb9, 0xe5, 0x12, 0x46, 0x5c, 0x82,
	0x76, 0xa3, 0x22, 0x64, 0xd3, 0xd5, 0x6c, 0x06, 0xce, 0x24, 0x82, 0xe6, 0x06, 0x1d, 0x76, 0x3b,
	0x02, 0x95, 0x42, 0xcd, 0xaa, 0xe7, 0x08, 0x6f, 0x8e, 0x70, 0xba, 0x9a, 0x55, 0xa4, 0xff, 0xe5,
	0x91, 0xce, 0x69, 0x34, 0xc2, 0x2c, 0x03, 0xe1, 0x12, 0xd4, 0x7e, 0x40, 0x5a, 0xa0, 0x8b, 0x08,
	0x63, 0x08, 0xbc, 0x9e, 0x37, 0x68, 0x2f, 0xea, 0xaf, 0xff, 0x48, 0x5a, 0xaa, 0xec, 0xb6, 0x41,
	0xa3, 0xd7, 0x1c, 0x5c, 0x87, 0x84, 0x9e, 0x16, 0x2c, 0xea, 0x92, 0x3f, 0x24, 0xb7, 0x16, 0x4c,
	0x91, 0x08, 0x98, 0x68, 0xeb, 0xb8, 0x16, 0x30, 0x89, 0x83, 0x66, 0xcf, 0x1b, 0x5c, 0x2e, 0xce,
	0x0b, 0xe1, 0x2b, 0xb9, 0xff, 0xdb, 0x01, 0x39, 0x1a, 0xb7, 0xac, 0x7a, 0xfc, 0x21, 0x69, 0x89,
	0x2a, 0x95, 0x7f, 0x47, 0xff, 0x09, 0xda, 0x6d, 0xd3, 0x08, 0x95, 0xe2, 0x3a, 0xb6, 0xfd, 0x8b,
	0xe7, 0x6f, 0x8f, 0x8c, 0xd1, 0x48, 0xca, 0x73, 0x2e, 0x36, 0x40, 0x6d, 0xba, 0xdf, 0xf1, 0x2c,
	0x4d, 0xf4, 0x81, 0x28, 0xaa, 0xc1, 0xed, 0xd0, 0xa4, 0xb4, 0x96, 0x47, 0x4b, 0x79, 0xb4, 0x08,
	0xe7, 0xde, 0xe7, 0x93, 0x4c, 0xdc, 0x66, 0xbb, 0xa6, 0x02, 0x15, 0x7b, 0xe1, 0x18, 0x65, 0xb8,
	0x8d, 0x47, 0x6f, 0x7c, 0x6d, 0x99, 0xc4, 0xd0, 0xa6, 0x7b, 0x66, 0xca, 0x58, 0x60, 0x98, 0x34,
	0xb9, 0x60, 0x67, 0xfe, 0x7f, 0x1a, 0xdd, 0x65, 0xba, 0xff, 0x38, 0x5e, 0x7b, 0xaf, 0x2e, 0xcd,
	0x0f, 0x8a, 0x05, 0x66, 0xeb, 0xab, 0x52, 0xf6, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0xec, 0x0a,
	0x7a, 0x2e, 0xb8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JVMMetricReportServiceClient is the client API for JVMMetricReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JVMMetricReportServiceClient interface {
	Collect(ctx context.Context, in *JVMMetricCollection, opts ...grpc.CallOption) (*common.Commands, error)
}

type jVMMetricReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJVMMetricReportServiceClient(cc grpc.ClientConnInterface) JVMMetricReportServiceClient {
	return &jVMMetricReportServiceClient{cc}
}

func (c *jVMMetricReportServiceClient) Collect(ctx context.Context, in *JVMMetricCollection, opts ...grpc.CallOption) (*common.Commands, error) {
	out := new(common.Commands)
	err := c.cc.Invoke(ctx, "/JVMMetricReportService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JVMMetricReportServiceServer is the server API for JVMMetricReportService service.
type JVMMetricReportServiceServer interface {
	Collect(context.Context, *JVMMetricCollection) (*common.Commands, error)
}

// UnimplementedJVMMetricReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJVMMetricReportServiceServer struct {
}

func (*UnimplementedJVMMetricReportServiceServer) Collect(ctx context.Context, req *JVMMetricCollection) (*common.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterJVMMetricReportServiceServer(s *grpc.Server, srv JVMMetricReportServiceServer) {
	s.RegisterService(&_JVMMetricReportService_serviceDesc, srv)
}

func _JVMMetricReportService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JVMMetricCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JVMMetricReportServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JVMMetricReportService/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JVMMetricReportServiceServer).Collect(ctx, req.(*JVMMetricCollection))
	}
	return interceptor(ctx, in, info, handler)
}

var _JVMMetricReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "JVMMetricReportService",
	HandlerType: (*JVMMetricReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _JVMMetricReportService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent-v2/JVMMetric.proto",
}
