// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_http_access_log.proto

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

// 创建访问日志
type CreateHTTPAccessLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessLogs []*HTTPAccessLog `protobuf:"bytes,1,rep,name=accessLogs,proto3" json:"accessLogs,omitempty"`
}

func (x *CreateHTTPAccessLogsRequest) Reset() {
	*x = CreateHTTPAccessLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_access_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPAccessLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPAccessLogsRequest) ProtoMessage() {}

func (x *CreateHTTPAccessLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_access_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPAccessLogsRequest.ProtoReflect.Descriptor instead.
func (*CreateHTTPAccessLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_http_access_log_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHTTPAccessLogsRequest) GetAccessLogs() []*HTTPAccessLog {
	if x != nil {
		return x.AccessLogs
	}
	return nil
}

type CreateHTTPAccessLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateHTTPAccessLogsResponse) Reset() {
	*x = CreateHTTPAccessLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_access_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPAccessLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPAccessLogsResponse) ProtoMessage() {}

func (x *CreateHTTPAccessLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_access_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPAccessLogsResponse.ProtoReflect.Descriptor instead.
func (*CreateHTTPAccessLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_http_access_log_proto_rawDescGZIP(), []int{1}
}

// 列出往前的单页访问日志
type ListHTTPAccessLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId           string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`                      // 上一页请求ID，可选
	ServerId            int64  `protobuf:"varint,2,opt,name=serverId,proto3" json:"serverId,omitempty"`                       // 服务ID
	Size                int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`                               // 单页条数
	Day                 string `protobuf:"bytes,4,opt,name=day,proto3" json:"day,omitempty"`                                  // 日期，格式YYYYMMDD
	Reverse             bool   `protobuf:"varint,5,opt,name=reverse,proto3" json:"reverse,omitempty"`                         // 是否反向查找，可选
	HasError            bool   `protobuf:"varint,6,opt,name=hasError,proto3" json:"hasError,omitempty"`                       // 是否有错误，可选
	FirewallPolicyId    int64  `protobuf:"varint,7,opt,name=firewallPolicyId,proto3" json:"firewallPolicyId,omitempty"`       // WAF策略ID，可选
	FirewallRuleGroupId int64  `protobuf:"varint,8,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"` // WAF分组ID，可选
	FirewallRuleSetId   int64  `protobuf:"varint,9,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`     // WAF规则集ID，可选
}

func (x *ListHTTPAccessLogsRequest) Reset() {
	*x = ListHTTPAccessLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_access_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHTTPAccessLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHTTPAccessLogsRequest) ProtoMessage() {}

func (x *ListHTTPAccessLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_access_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHTTPAccessLogsRequest.ProtoReflect.Descriptor instead.
func (*ListHTTPAccessLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_http_access_log_proto_rawDescGZIP(), []int{2}
}

func (x *ListHTTPAccessLogsRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ListHTTPAccessLogsRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *ListHTTPAccessLogsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListHTTPAccessLogsRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ListHTTPAccessLogsRequest) GetReverse() bool {
	if x != nil {
		return x.Reverse
	}
	return false
}

func (x *ListHTTPAccessLogsRequest) GetHasError() bool {
	if x != nil {
		return x.HasError
	}
	return false
}

func (x *ListHTTPAccessLogsRequest) GetFirewallPolicyId() int64 {
	if x != nil {
		return x.FirewallPolicyId
	}
	return 0
}

func (x *ListHTTPAccessLogsRequest) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

func (x *ListHTTPAccessLogsRequest) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

type ListHTTPAccessLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessLogs []*HTTPAccessLog `protobuf:"bytes,1,rep,name=accessLogs,proto3" json:"accessLogs,omitempty"`
	RequestId  string           `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
	HasMore    bool             `protobuf:"varint,3,opt,name=hasMore,proto3" json:"hasMore,omitempty"`
}

func (x *ListHTTPAccessLogsResponse) Reset() {
	*x = ListHTTPAccessLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_access_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHTTPAccessLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHTTPAccessLogsResponse) ProtoMessage() {}

func (x *ListHTTPAccessLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_access_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHTTPAccessLogsResponse.ProtoReflect.Descriptor instead.
func (*ListHTTPAccessLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_http_access_log_proto_rawDescGZIP(), []int{3}
}

func (x *ListHTTPAccessLogsResponse) GetAccessLogs() []*HTTPAccessLog {
	if x != nil {
		return x.AccessLogs
	}
	return nil
}

func (x *ListHTTPAccessLogsResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ListHTTPAccessLogsResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

// 查找单个日志
type FindHTTPAccessLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *FindHTTPAccessLogRequest) Reset() {
	*x = FindHTTPAccessLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_access_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindHTTPAccessLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindHTTPAccessLogRequest) ProtoMessage() {}

func (x *FindHTTPAccessLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_access_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindHTTPAccessLogRequest.ProtoReflect.Descriptor instead.
func (*FindHTTPAccessLogRequest) Descriptor() ([]byte, []int) {
	return file_service_http_access_log_proto_rawDescGZIP(), []int{4}
}

func (x *FindHTTPAccessLogRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type FindHTTPAccessLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessLog *HTTPAccessLog `protobuf:"bytes,1,opt,name=accessLog,proto3" json:"accessLog,omitempty"`
}

func (x *FindHTTPAccessLogResponse) Reset() {
	*x = FindHTTPAccessLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_access_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindHTTPAccessLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindHTTPAccessLogResponse) ProtoMessage() {}

func (x *FindHTTPAccessLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_access_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindHTTPAccessLogResponse.ProtoReflect.Descriptor instead.
func (*FindHTTPAccessLogResponse) Descriptor() ([]byte, []int) {
	return file_service_http_access_log_proto_rawDescGZIP(), []int{5}
}

func (x *FindHTTPAccessLogResponse) GetAccessLog() *HTTPAccessLog {
	if x != nil {
		return x.AccessLog
	}
	return nil
}

var File_service_http_access_log_proto protoreflect.FileDescriptor

var file_service_http_access_log_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x50, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xbd, 0x02, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61,
	0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61,
	0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74,
	0x49, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22, 0x38, 0x0a, 0x18,
	0x46, 0x69, 0x6e, 0x64, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x54,
	0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x32, 0x98, 0x02, 0x0a, 0x14, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a,
	0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x6c, 0x69, 0x73, 0x74,
	0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a,
	0x11, 0x66, 0x69, 0x6e, 0x64, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x6f, 0x67, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x54, 0x54, 0x50,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_access_log_proto_rawDescOnce sync.Once
	file_service_http_access_log_proto_rawDescData = file_service_http_access_log_proto_rawDesc
)

func file_service_http_access_log_proto_rawDescGZIP() []byte {
	file_service_http_access_log_proto_rawDescOnce.Do(func() {
		file_service_http_access_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_access_log_proto_rawDescData)
	})
	return file_service_http_access_log_proto_rawDescData
}

var file_service_http_access_log_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_http_access_log_proto_goTypes = []interface{}{
	(*CreateHTTPAccessLogsRequest)(nil),  // 0: pb.CreateHTTPAccessLogsRequest
	(*CreateHTTPAccessLogsResponse)(nil), // 1: pb.CreateHTTPAccessLogsResponse
	(*ListHTTPAccessLogsRequest)(nil),    // 2: pb.ListHTTPAccessLogsRequest
	(*ListHTTPAccessLogsResponse)(nil),   // 3: pb.ListHTTPAccessLogsResponse
	(*FindHTTPAccessLogRequest)(nil),     // 4: pb.FindHTTPAccessLogRequest
	(*FindHTTPAccessLogResponse)(nil),    // 5: pb.FindHTTPAccessLogResponse
	(*HTTPAccessLog)(nil),                // 6: pb.HTTPAccessLog
}
var file_service_http_access_log_proto_depIdxs = []int32{
	6, // 0: pb.CreateHTTPAccessLogsRequest.accessLogs:type_name -> pb.HTTPAccessLog
	6, // 1: pb.ListHTTPAccessLogsResponse.accessLogs:type_name -> pb.HTTPAccessLog
	6, // 2: pb.FindHTTPAccessLogResponse.accessLog:type_name -> pb.HTTPAccessLog
	0, // 3: pb.HTTPAccessLogService.createHTTPAccessLogs:input_type -> pb.CreateHTTPAccessLogsRequest
	2, // 4: pb.HTTPAccessLogService.listHTTPAccessLogs:input_type -> pb.ListHTTPAccessLogsRequest
	4, // 5: pb.HTTPAccessLogService.findHTTPAccessLog:input_type -> pb.FindHTTPAccessLogRequest
	1, // 6: pb.HTTPAccessLogService.createHTTPAccessLogs:output_type -> pb.CreateHTTPAccessLogsResponse
	3, // 7: pb.HTTPAccessLogService.listHTTPAccessLogs:output_type -> pb.ListHTTPAccessLogsResponse
	5, // 8: pb.HTTPAccessLogService.findHTTPAccessLog:output_type -> pb.FindHTTPAccessLogResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_http_access_log_proto_init() }
func file_service_http_access_log_proto_init() {
	if File_service_http_access_log_proto != nil {
		return
	}
	file_model_http_access_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_access_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPAccessLogsRequest); i {
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
		file_service_http_access_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPAccessLogsResponse); i {
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
		file_service_http_access_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHTTPAccessLogsRequest); i {
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
		file_service_http_access_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHTTPAccessLogsResponse); i {
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
		file_service_http_access_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindHTTPAccessLogRequest); i {
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
		file_service_http_access_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindHTTPAccessLogResponse); i {
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
			RawDescriptor: file_service_http_access_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_access_log_proto_goTypes,
		DependencyIndexes: file_service_http_access_log_proto_depIdxs,
		MessageInfos:      file_service_http_access_log_proto_msgTypes,
	}.Build()
	File_service_http_access_log_proto = out.File
	file_service_http_access_log_proto_rawDesc = nil
	file_service_http_access_log_proto_goTypes = nil
	file_service_http_access_log_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPAccessLogServiceClient is the client API for HTTPAccessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPAccessLogServiceClient interface {
	// 创建访问日志
	CreateHTTPAccessLogs(ctx context.Context, in *CreateHTTPAccessLogsRequest, opts ...grpc.CallOption) (*CreateHTTPAccessLogsResponse, error)
	// 列出单页访问日志
	ListHTTPAccessLogs(ctx context.Context, in *ListHTTPAccessLogsRequest, opts ...grpc.CallOption) (*ListHTTPAccessLogsResponse, error)
	// 查找单个日志
	FindHTTPAccessLog(ctx context.Context, in *FindHTTPAccessLogRequest, opts ...grpc.CallOption) (*FindHTTPAccessLogResponse, error)
}

type hTTPAccessLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPAccessLogServiceClient(cc grpc.ClientConnInterface) HTTPAccessLogServiceClient {
	return &hTTPAccessLogServiceClient{cc}
}

func (c *hTTPAccessLogServiceClient) CreateHTTPAccessLogs(ctx context.Context, in *CreateHTTPAccessLogsRequest, opts ...grpc.CallOption) (*CreateHTTPAccessLogsResponse, error) {
	out := new(CreateHTTPAccessLogsResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPAccessLogService/createHTTPAccessLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPAccessLogServiceClient) ListHTTPAccessLogs(ctx context.Context, in *ListHTTPAccessLogsRequest, opts ...grpc.CallOption) (*ListHTTPAccessLogsResponse, error) {
	out := new(ListHTTPAccessLogsResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPAccessLogService/listHTTPAccessLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPAccessLogServiceClient) FindHTTPAccessLog(ctx context.Context, in *FindHTTPAccessLogRequest, opts ...grpc.CallOption) (*FindHTTPAccessLogResponse, error) {
	out := new(FindHTTPAccessLogResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPAccessLogService/findHTTPAccessLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPAccessLogServiceServer is the server API for HTTPAccessLogService service.
type HTTPAccessLogServiceServer interface {
	// 创建访问日志
	CreateHTTPAccessLogs(context.Context, *CreateHTTPAccessLogsRequest) (*CreateHTTPAccessLogsResponse, error)
	// 列出单页访问日志
	ListHTTPAccessLogs(context.Context, *ListHTTPAccessLogsRequest) (*ListHTTPAccessLogsResponse, error)
	// 查找单个日志
	FindHTTPAccessLog(context.Context, *FindHTTPAccessLogRequest) (*FindHTTPAccessLogResponse, error)
}

// UnimplementedHTTPAccessLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPAccessLogServiceServer struct {
}

func (*UnimplementedHTTPAccessLogServiceServer) CreateHTTPAccessLogs(context.Context, *CreateHTTPAccessLogsRequest) (*CreateHTTPAccessLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPAccessLogs not implemented")
}
func (*UnimplementedHTTPAccessLogServiceServer) ListHTTPAccessLogs(context.Context, *ListHTTPAccessLogsRequest) (*ListHTTPAccessLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHTTPAccessLogs not implemented")
}
func (*UnimplementedHTTPAccessLogServiceServer) FindHTTPAccessLog(context.Context, *FindHTTPAccessLogRequest) (*FindHTTPAccessLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindHTTPAccessLog not implemented")
}

func RegisterHTTPAccessLogServiceServer(s *grpc.Server, srv HTTPAccessLogServiceServer) {
	s.RegisterService(&_HTTPAccessLogService_serviceDesc, srv)
}

func _HTTPAccessLogService_CreateHTTPAccessLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPAccessLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAccessLogServiceServer).CreateHTTPAccessLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPAccessLogService/CreateHTTPAccessLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAccessLogServiceServer).CreateHTTPAccessLogs(ctx, req.(*CreateHTTPAccessLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPAccessLogService_ListHTTPAccessLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHTTPAccessLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAccessLogServiceServer).ListHTTPAccessLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPAccessLogService/ListHTTPAccessLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAccessLogServiceServer).ListHTTPAccessLogs(ctx, req.(*ListHTTPAccessLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPAccessLogService_FindHTTPAccessLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindHTTPAccessLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAccessLogServiceServer).FindHTTPAccessLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPAccessLogService/FindHTTPAccessLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAccessLogServiceServer).FindHTTPAccessLog(ctx, req.(*FindHTTPAccessLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTPAccessLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPAccessLogService",
	HandlerType: (*HTTPAccessLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPAccessLogs",
			Handler:    _HTTPAccessLogService_CreateHTTPAccessLogs_Handler,
		},
		{
			MethodName: "listHTTPAccessLogs",
			Handler:    _HTTPAccessLogService_ListHTTPAccessLogs_Handler,
		},
		{
			MethodName: "findHTTPAccessLog",
			Handler:    _HTTPAccessLogService_FindHTTPAccessLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_access_log.proto",
}
