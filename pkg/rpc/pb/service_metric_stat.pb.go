// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_metric_stat.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 上传统计数据
type UploadMetricStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricStats []*MetricStat `protobuf:"bytes,1,rep,name=metricStats,proto3" json:"metricStats,omitempty"`
}

func (x *UploadMetricStatsRequest) Reset() {
	*x = UploadMetricStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_metric_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadMetricStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMetricStatsRequest) ProtoMessage() {}

func (x *UploadMetricStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_metric_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMetricStatsRequest.ProtoReflect.Descriptor instead.
func (*UploadMetricStatsRequest) Descriptor() ([]byte, []int) {
	return file_service_metric_stat_proto_rawDescGZIP(), []int{0}
}

func (x *UploadMetricStatsRequest) GetMetricStats() []*MetricStat {
	if x != nil {
		return x.MetricStats
	}
	return nil
}

// 计算指标数据数量
type CountMetricStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricItemId int64 `protobuf:"varint,1,opt,name=metricItemId,proto3" json:"metricItemId,omitempty"`
}

func (x *CountMetricStatsRequest) Reset() {
	*x = CountMetricStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_metric_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountMetricStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountMetricStatsRequest) ProtoMessage() {}

func (x *CountMetricStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_metric_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountMetricStatsRequest.ProtoReflect.Descriptor instead.
func (*CountMetricStatsRequest) Descriptor() ([]byte, []int) {
	return file_service_metric_stat_proto_rawDescGZIP(), []int{1}
}

func (x *CountMetricStatsRequest) GetMetricItemId() int64 {
	if x != nil {
		return x.MetricItemId
	}
	return 0
}

// 读取单页指标数据
type ListMetricStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricItemId int64 `protobuf:"varint,1,opt,name=metricItemId,proto3" json:"metricItemId,omitempty"`
	Offset       int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size         int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListMetricStatsRequest) Reset() {
	*x = ListMetricStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_metric_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMetricStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMetricStatsRequest) ProtoMessage() {}

func (x *ListMetricStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_metric_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMetricStatsRequest.ProtoReflect.Descriptor instead.
func (*ListMetricStatsRequest) Descriptor() ([]byte, []int) {
	return file_service_metric_stat_proto_rawDescGZIP(), []int{2}
}

func (x *ListMetricStatsRequest) GetMetricItemId() int64 {
	if x != nil {
		return x.MetricItemId
	}
	return 0
}

func (x *ListMetricStatsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListMetricStatsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListMetricStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricStats []*MetricStat `protobuf:"bytes,1,rep,name=metricStats,proto3" json:"metricStats,omitempty"`
}

func (x *ListMetricStatsResponse) Reset() {
	*x = ListMetricStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_metric_stat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMetricStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMetricStatsResponse) ProtoMessage() {}

func (x *ListMetricStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_metric_stat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMetricStatsResponse.ProtoReflect.Descriptor instead.
func (*ListMetricStatsResponse) Descriptor() ([]byte, []int) {
	return file_service_metric_stat_proto_rawDescGZIP(), []int{3}
}

func (x *ListMetricStatsResponse) GetMetricStats() []*MetricStat {
	if x != nil {
		return x.MetricStats
	}
	return nil
}

var File_service_metric_stat_proto protoreflect.FileDescriptor

var file_service_metric_stat_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x18, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x22, 0x4b, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x32, 0xe9,
	0x01, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50,
	0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x0f, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_metric_stat_proto_rawDescOnce sync.Once
	file_service_metric_stat_proto_rawDescData = file_service_metric_stat_proto_rawDesc
)

func file_service_metric_stat_proto_rawDescGZIP() []byte {
	file_service_metric_stat_proto_rawDescOnce.Do(func() {
		file_service_metric_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_metric_stat_proto_rawDescData)
	})
	return file_service_metric_stat_proto_rawDescData
}

var file_service_metric_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_metric_stat_proto_goTypes = []interface{}{
	(*UploadMetricStatsRequest)(nil), // 0: pb.UploadMetricStatsRequest
	(*CountMetricStatsRequest)(nil),  // 1: pb.CountMetricStatsRequest
	(*ListMetricStatsRequest)(nil),   // 2: pb.ListMetricStatsRequest
	(*ListMetricStatsResponse)(nil),  // 3: pb.ListMetricStatsResponse
	(*MetricStat)(nil),               // 4: pb.MetricStat
	(*RPCSuccess)(nil),               // 5: pb.RPCSuccess
	(*RPCCountResponse)(nil),         // 6: pb.RPCCountResponse
}
var file_service_metric_stat_proto_depIdxs = []int32{
	4, // 0: pb.UploadMetricStatsRequest.metricStats:type_name -> pb.MetricStat
	4, // 1: pb.ListMetricStatsResponse.metricStats:type_name -> pb.MetricStat
	0, // 2: pb.MetricStatService.uploadMetricStats:input_type -> pb.UploadMetricStatsRequest
	1, // 3: pb.MetricStatService.countMetricStats:input_type -> pb.CountMetricStatsRequest
	2, // 4: pb.MetricStatService.listMetricStats:input_type -> pb.ListMetricStatsRequest
	5, // 5: pb.MetricStatService.uploadMetricStats:output_type -> pb.RPCSuccess
	6, // 6: pb.MetricStatService.countMetricStats:output_type -> pb.RPCCountResponse
	3, // 7: pb.MetricStatService.listMetricStats:output_type -> pb.ListMetricStatsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_metric_stat_proto_init() }
func file_service_metric_stat_proto_init() {
	if File_service_metric_stat_proto != nil {
		return
	}
	file_models_model_metric_stat_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_metric_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadMetricStatsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_metric_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountMetricStatsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_metric_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMetricStatsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_metric_stat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMetricStatsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_metric_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_metric_stat_proto_goTypes,
		DependencyIndexes: file_service_metric_stat_proto_depIdxs,
		MessageInfos:      file_service_metric_stat_proto_msgTypes,
	}.Build()
	File_service_metric_stat_proto = out.File
	file_service_metric_stat_proto_rawDesc = nil
	file_service_metric_stat_proto_goTypes = nil
	file_service_metric_stat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetricStatServiceClient is the client API for MetricStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricStatServiceClient interface {
	// 上传统计数据
	UploadMetricStats(ctx context.Context, in *UploadMetricStatsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算指标数据数量
	CountMetricStats(ctx context.Context, in *CountMetricStatsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 读取单页指标数据
	ListMetricStats(ctx context.Context, in *ListMetricStatsRequest, opts ...grpc.CallOption) (*ListMetricStatsResponse, error)
}

type metricStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricStatServiceClient(cc grpc.ClientConnInterface) MetricStatServiceClient {
	return &metricStatServiceClient{cc}
}

func (c *metricStatServiceClient) UploadMetricStats(ctx context.Context, in *UploadMetricStatsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MetricStatService/uploadMetricStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricStatServiceClient) CountMetricStats(ctx context.Context, in *CountMetricStatsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.MetricStatService/countMetricStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricStatServiceClient) ListMetricStats(ctx context.Context, in *ListMetricStatsRequest, opts ...grpc.CallOption) (*ListMetricStatsResponse, error) {
	out := new(ListMetricStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.MetricStatService/listMetricStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricStatServiceServer is the server API for MetricStatService service.
type MetricStatServiceServer interface {
	// 上传统计数据
	UploadMetricStats(context.Context, *UploadMetricStatsRequest) (*RPCSuccess, error)
	// 计算指标数据数量
	CountMetricStats(context.Context, *CountMetricStatsRequest) (*RPCCountResponse, error)
	// 读取单页指标数据
	ListMetricStats(context.Context, *ListMetricStatsRequest) (*ListMetricStatsResponse, error)
}

// UnimplementedMetricStatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricStatServiceServer struct {
}

func (*UnimplementedMetricStatServiceServer) UploadMetricStats(context.Context, *UploadMetricStatsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadMetricStats not implemented")
}
func (*UnimplementedMetricStatServiceServer) CountMetricStats(context.Context, *CountMetricStatsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountMetricStats not implemented")
}
func (*UnimplementedMetricStatServiceServer) ListMetricStats(context.Context, *ListMetricStatsRequest) (*ListMetricStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetricStats not implemented")
}

func RegisterMetricStatServiceServer(s *grpc.Server, srv MetricStatServiceServer) {
	s.RegisterService(&_MetricStatService_serviceDesc, srv)
}

func _MetricStatService_UploadMetricStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadMetricStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricStatServiceServer).UploadMetricStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricStatService/UploadMetricStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricStatServiceServer).UploadMetricStats(ctx, req.(*UploadMetricStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricStatService_CountMetricStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountMetricStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricStatServiceServer).CountMetricStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricStatService/CountMetricStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricStatServiceServer).CountMetricStats(ctx, req.(*CountMetricStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricStatService_ListMetricStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricStatServiceServer).ListMetricStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricStatService/ListMetricStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricStatServiceServer).ListMetricStats(ctx, req.(*ListMetricStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricStatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MetricStatService",
	HandlerType: (*MetricStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "uploadMetricStats",
			Handler:    _MetricStatService_UploadMetricStats_Handler,
		},
		{
			MethodName: "countMetricStats",
			Handler:    _MetricStatService_CountMetricStats_Handler,
		},
		{
			MethodName: "listMetricStats",
			Handler:    _MetricStatService_ListMetricStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_metric_stat.proto",
}
