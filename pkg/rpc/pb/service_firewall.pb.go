// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_firewall.proto

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

// 组合看板数据
type ComposeFirewallGlobalBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComposeFirewallGlobalBoardRequest) Reset() {
	*x = ComposeFirewallGlobalBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_firewall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeFirewallGlobalBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeFirewallGlobalBoardRequest) ProtoMessage() {}

func (x *ComposeFirewallGlobalBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_firewall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeFirewallGlobalBoardRequest.ProtoReflect.Descriptor instead.
func (*ComposeFirewallGlobalBoardRequest) Descriptor() ([]byte, []int) {
	return file_service_firewall_proto_rawDescGZIP(), []int{0}
}

type ComposeFirewallGlobalBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountDailyLogs         int64                                                           `protobuf:"varint,1,opt,name=countDailyLogs,proto3" json:"countDailyLogs,omitempty"`
	CountDailyBlocks       int64                                                           `protobuf:"varint,2,opt,name=countDailyBlocks,proto3" json:"countDailyBlocks,omitempty"`
	CountDailyCaptcha      int64                                                           `protobuf:"varint,3,opt,name=countDailyCaptcha,proto3" json:"countDailyCaptcha,omitempty"`
	CountWeeklyBlocks      int64                                                           `protobuf:"varint,4,opt,name=countWeeklyBlocks,proto3" json:"countWeeklyBlocks,omitempty"`
	HttpFirewallRuleGroups []*ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat `protobuf:"bytes,30,rep,name=httpFirewallRuleGroups,proto3" json:"httpFirewallRuleGroups,omitempty"`
	DailyStats             []*ComposeFirewallGlobalBoardResponse_DailyStat                 `protobuf:"bytes,31,rep,name=dailyStats,proto3" json:"dailyStats,omitempty"`
	HourlyStats            []*ComposeFirewallGlobalBoardResponse_HourlyStat                `protobuf:"bytes,32,rep,name=hourlyStats,proto3" json:"hourlyStats,omitempty"`
}

func (x *ComposeFirewallGlobalBoardResponse) Reset() {
	*x = ComposeFirewallGlobalBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_firewall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeFirewallGlobalBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeFirewallGlobalBoardResponse) ProtoMessage() {}

func (x *ComposeFirewallGlobalBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_firewall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeFirewallGlobalBoardResponse.ProtoReflect.Descriptor instead.
func (*ComposeFirewallGlobalBoardResponse) Descriptor() ([]byte, []int) {
	return file_service_firewall_proto_rawDescGZIP(), []int{1}
}

func (x *ComposeFirewallGlobalBoardResponse) GetCountDailyLogs() int64 {
	if x != nil {
		return x.CountDailyLogs
	}
	return 0
}

func (x *ComposeFirewallGlobalBoardResponse) GetCountDailyBlocks() int64 {
	if x != nil {
		return x.CountDailyBlocks
	}
	return 0
}

func (x *ComposeFirewallGlobalBoardResponse) GetCountDailyCaptcha() int64 {
	if x != nil {
		return x.CountDailyCaptcha
	}
	return 0
}

func (x *ComposeFirewallGlobalBoardResponse) GetCountWeeklyBlocks() int64 {
	if x != nil {
		return x.CountWeeklyBlocks
	}
	return 0
}

func (x *ComposeFirewallGlobalBoardResponse) GetHttpFirewallRuleGroups() []*ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat {
	if x != nil {
		return x.HttpFirewallRuleGroups
	}
	return nil
}

func (x *ComposeFirewallGlobalBoardResponse) GetDailyStats() []*ComposeFirewallGlobalBoardResponse_DailyStat {
	if x != nil {
		return x.DailyStats
	}
	return nil
}

func (x *ComposeFirewallGlobalBoardResponse) GetHourlyStats() []*ComposeFirewallGlobalBoardResponse_HourlyStat {
	if x != nil {
		return x.HourlyStats
	}
	return nil
}

// 发送告警(notify)消息
type NotifyHTTPFirewallEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId                int64 `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	HttpFirewallPolicyId    int64 `protobuf:"varint,2,opt,name=httpFirewallPolicyId,proto3" json:"httpFirewallPolicyId,omitempty"`
	HttpFirewallRuleGroupId int64 `protobuf:"varint,3,opt,name=httpFirewallRuleGroupId,proto3" json:"httpFirewallRuleGroupId,omitempty"`
	HttpFirewallRuleSetId   int64 `protobuf:"varint,4,opt,name=httpFirewallRuleSetId,proto3" json:"httpFirewallRuleSetId,omitempty"`
	CreatedAt               int64 `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *NotifyHTTPFirewallEventRequest) Reset() {
	*x = NotifyHTTPFirewallEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_firewall_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyHTTPFirewallEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyHTTPFirewallEventRequest) ProtoMessage() {}

func (x *NotifyHTTPFirewallEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_firewall_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyHTTPFirewallEventRequest.ProtoReflect.Descriptor instead.
func (*NotifyHTTPFirewallEventRequest) Descriptor() ([]byte, []int) {
	return file_service_firewall_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyHTTPFirewallEventRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *NotifyHTTPFirewallEventRequest) GetHttpFirewallPolicyId() int64 {
	if x != nil {
		return x.HttpFirewallPolicyId
	}
	return 0
}

func (x *NotifyHTTPFirewallEventRequest) GetHttpFirewallRuleGroupId() int64 {
	if x != nil {
		return x.HttpFirewallRuleGroupId
	}
	return 0
}

func (x *NotifyHTTPFirewallEventRequest) GetHttpFirewallRuleSetId() int64 {
	if x != nil {
		return x.HttpFirewallRuleSetId
	}
	return 0
}

func (x *NotifyHTTPFirewallEventRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpFirewallRuleGroup *HTTPFirewallRuleGroup `protobuf:"bytes,1,opt,name=httpFirewallRuleGroup,proto3" json:"httpFirewallRuleGroup,omitempty"`
	Count                 int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat) Reset() {
	*x = ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_firewall_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat) ProtoMessage() {}

func (x *ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_firewall_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat.ProtoReflect.Descriptor instead.
func (*ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat) Descriptor() ([]byte, []int) {
	return file_service_firewall_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat) GetHttpFirewallRuleGroup() *HTTPFirewallRuleGroup {
	if x != nil {
		return x.HttpFirewallRuleGroup
	}
	return nil
}

func (x *ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ComposeFirewallGlobalBoardResponse_HourlyStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hour         string `protobuf:"bytes,1,opt,name=hour,proto3" json:"hour,omitempty"`
	CountLogs    int64  `protobuf:"varint,2,opt,name=countLogs,proto3" json:"countLogs,omitempty"`
	CountCaptcha int64  `protobuf:"varint,3,opt,name=countCaptcha,proto3" json:"countCaptcha,omitempty"`
	CountBlocks  int64  `protobuf:"varint,4,opt,name=countBlocks,proto3" json:"countBlocks,omitempty"`
}

func (x *ComposeFirewallGlobalBoardResponse_HourlyStat) Reset() {
	*x = ComposeFirewallGlobalBoardResponse_HourlyStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_firewall_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeFirewallGlobalBoardResponse_HourlyStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeFirewallGlobalBoardResponse_HourlyStat) ProtoMessage() {}

func (x *ComposeFirewallGlobalBoardResponse_HourlyStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_firewall_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeFirewallGlobalBoardResponse_HourlyStat.ProtoReflect.Descriptor instead.
func (*ComposeFirewallGlobalBoardResponse_HourlyStat) Descriptor() ([]byte, []int) {
	return file_service_firewall_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ComposeFirewallGlobalBoardResponse_HourlyStat) GetHour() string {
	if x != nil {
		return x.Hour
	}
	return ""
}

func (x *ComposeFirewallGlobalBoardResponse_HourlyStat) GetCountLogs() int64 {
	if x != nil {
		return x.CountLogs
	}
	return 0
}

func (x *ComposeFirewallGlobalBoardResponse_HourlyStat) GetCountCaptcha() int64 {
	if x != nil {
		return x.CountCaptcha
	}
	return 0
}

func (x *ComposeFirewallGlobalBoardResponse_HourlyStat) GetCountBlocks() int64 {
	if x != nil {
		return x.CountBlocks
	}
	return 0
}

type ComposeFirewallGlobalBoardResponse_DailyStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day          string `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
	CountLogs    int64  `protobuf:"varint,2,opt,name=countLogs,proto3" json:"countLogs,omitempty"`
	CountCaptcha int64  `protobuf:"varint,3,opt,name=countCaptcha,proto3" json:"countCaptcha,omitempty"`
	CountBlocks  int64  `protobuf:"varint,4,opt,name=countBlocks,proto3" json:"countBlocks,omitempty"`
}

func (x *ComposeFirewallGlobalBoardResponse_DailyStat) Reset() {
	*x = ComposeFirewallGlobalBoardResponse_DailyStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_firewall_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeFirewallGlobalBoardResponse_DailyStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeFirewallGlobalBoardResponse_DailyStat) ProtoMessage() {}

func (x *ComposeFirewallGlobalBoardResponse_DailyStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_firewall_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeFirewallGlobalBoardResponse_DailyStat.ProtoReflect.Descriptor instead.
func (*ComposeFirewallGlobalBoardResponse_DailyStat) Descriptor() ([]byte, []int) {
	return file_service_firewall_proto_rawDescGZIP(), []int{1, 2}
}

func (x *ComposeFirewallGlobalBoardResponse_DailyStat) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ComposeFirewallGlobalBoardResponse_DailyStat) GetCountLogs() int64 {
	if x != nil {
		return x.CountLogs
	}
	return 0
}

func (x *ComposeFirewallGlobalBoardResponse_DailyStat) GetCountCaptcha() int64 {
	if x != nil {
		return x.CountCaptcha
	}
	return 0
}

func (x *ComposeFirewallGlobalBoardResponse_DailyStat) GetCountBlocks() int64 {
	if x != nil {
		return x.CountBlocks
	}
	return 0
}

var File_service_firewall_proto protoreflect.FileDescriptor

var file_service_firewall_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x21, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x85, 0x07, 0x0a, 0x22, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4c, 0x6f,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c,
	0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x12, 0x78, 0x0a, 0x16, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x52, 0x16, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x50, 0x0a, 0x0a, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x0a, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x53, 0x0a, 0x0b,
	0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x0b, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x1a, 0x82, 0x01, 0x0a, 0x19, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x4f, 0x0a, 0x15, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x15, 0x68, 0x74, 0x74, 0x70, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x84, 0x01, 0x0a, 0x0a, 0x48, 0x6f, 0x75, 0x72, 0x6c,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x1a, 0x81, 0x01,
	0x0a, 0x09, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x22, 0xfe, 0x01, 0x0a, 0x1e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x14, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14,
	0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x34,
	0x0a, 0x15, 0x68, 0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x68,
	0x74, 0x74, 0x70, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x32, 0xcd, 0x01, 0x0a, 0x0f, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x17, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x54, 0x54,
	0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x22,
	0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_firewall_proto_rawDescOnce sync.Once
	file_service_firewall_proto_rawDescData = file_service_firewall_proto_rawDesc
)

func file_service_firewall_proto_rawDescGZIP() []byte {
	file_service_firewall_proto_rawDescOnce.Do(func() {
		file_service_firewall_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_firewall_proto_rawDescData)
	})
	return file_service_firewall_proto_rawDescData
}

var file_service_firewall_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_firewall_proto_goTypes = []interface{}{
	(*ComposeFirewallGlobalBoardRequest)(nil),                            // 0: pb.ComposeFirewallGlobalBoardRequest
	(*ComposeFirewallGlobalBoardResponse)(nil),                           // 1: pb.ComposeFirewallGlobalBoardResponse
	(*NotifyHTTPFirewallEventRequest)(nil),                               // 2: pb.NotifyHTTPFirewallEventRequest
	(*ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat)(nil), // 3: pb.ComposeFirewallGlobalBoardResponse.HTTPFirewallRuleGroupStat
	(*ComposeFirewallGlobalBoardResponse_HourlyStat)(nil),                // 4: pb.ComposeFirewallGlobalBoardResponse.HourlyStat
	(*ComposeFirewallGlobalBoardResponse_DailyStat)(nil),                 // 5: pb.ComposeFirewallGlobalBoardResponse.DailyStat
	(*HTTPFirewallRuleGroup)(nil),                                        // 6: pb.HTTPFirewallRuleGroup
	(*RPCSuccess)(nil),                                                   // 7: pb.RPCSuccess
}
var file_service_firewall_proto_depIdxs = []int32{
	3, // 0: pb.ComposeFirewallGlobalBoardResponse.httpFirewallRuleGroups:type_name -> pb.ComposeFirewallGlobalBoardResponse.HTTPFirewallRuleGroupStat
	5, // 1: pb.ComposeFirewallGlobalBoardResponse.dailyStats:type_name -> pb.ComposeFirewallGlobalBoardResponse.DailyStat
	4, // 2: pb.ComposeFirewallGlobalBoardResponse.hourlyStats:type_name -> pb.ComposeFirewallGlobalBoardResponse.HourlyStat
	6, // 3: pb.ComposeFirewallGlobalBoardResponse.HTTPFirewallRuleGroupStat.httpFirewallRuleGroup:type_name -> pb.HTTPFirewallRuleGroup
	0, // 4: pb.FirewallService.composeFirewallGlobalBoard:input_type -> pb.ComposeFirewallGlobalBoardRequest
	2, // 5: pb.FirewallService.notifyHTTPFirewallEvent:input_type -> pb.NotifyHTTPFirewallEventRequest
	1, // 6: pb.FirewallService.composeFirewallGlobalBoard:output_type -> pb.ComposeFirewallGlobalBoardResponse
	7, // 7: pb.FirewallService.notifyHTTPFirewallEvent:output_type -> pb.RPCSuccess
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_firewall_proto_init() }
func file_service_firewall_proto_init() {
	if File_service_firewall_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_http_firewall_rule_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_firewall_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeFirewallGlobalBoardRequest); i {
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
		file_service_firewall_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeFirewallGlobalBoardResponse); i {
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
		file_service_firewall_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyHTTPFirewallEventRequest); i {
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
		file_service_firewall_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeFirewallGlobalBoardResponse_HTTPFirewallRuleGroupStat); i {
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
		file_service_firewall_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeFirewallGlobalBoardResponse_HourlyStat); i {
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
		file_service_firewall_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeFirewallGlobalBoardResponse_DailyStat); i {
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
			RawDescriptor: file_service_firewall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_firewall_proto_goTypes,
		DependencyIndexes: file_service_firewall_proto_depIdxs,
		MessageInfos:      file_service_firewall_proto_msgTypes,
	}.Build()
	File_service_firewall_proto = out.File
	file_service_firewall_proto_rawDesc = nil
	file_service_firewall_proto_goTypes = nil
	file_service_firewall_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FirewallServiceClient is the client API for FirewallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FirewallServiceClient interface {
	// 组合看板数据
	ComposeFirewallGlobalBoard(ctx context.Context, in *ComposeFirewallGlobalBoardRequest, opts ...grpc.CallOption) (*ComposeFirewallGlobalBoardResponse, error)
	// 发送告警(notify)消息
	NotifyHTTPFirewallEvent(ctx context.Context, in *NotifyHTTPFirewallEventRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type firewallServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFirewallServiceClient(cc grpc.ClientConnInterface) FirewallServiceClient {
	return &firewallServiceClient{cc}
}

func (c *firewallServiceClient) ComposeFirewallGlobalBoard(ctx context.Context, in *ComposeFirewallGlobalBoardRequest, opts ...grpc.CallOption) (*ComposeFirewallGlobalBoardResponse, error) {
	out := new(ComposeFirewallGlobalBoardResponse)
	err := c.cc.Invoke(ctx, "/pb.FirewallService/composeFirewallGlobalBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firewallServiceClient) NotifyHTTPFirewallEvent(ctx context.Context, in *NotifyHTTPFirewallEventRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.FirewallService/notifyHTTPFirewallEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirewallServiceServer is the server API for FirewallService service.
type FirewallServiceServer interface {
	// 组合看板数据
	ComposeFirewallGlobalBoard(context.Context, *ComposeFirewallGlobalBoardRequest) (*ComposeFirewallGlobalBoardResponse, error)
	// 发送告警(notify)消息
	NotifyHTTPFirewallEvent(context.Context, *NotifyHTTPFirewallEventRequest) (*RPCSuccess, error)
}

// UnimplementedFirewallServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFirewallServiceServer struct {
}

func (*UnimplementedFirewallServiceServer) ComposeFirewallGlobalBoard(context.Context, *ComposeFirewallGlobalBoardRequest) (*ComposeFirewallGlobalBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeFirewallGlobalBoard not implemented")
}
func (*UnimplementedFirewallServiceServer) NotifyHTTPFirewallEvent(context.Context, *NotifyHTTPFirewallEventRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyHTTPFirewallEvent not implemented")
}

func RegisterFirewallServiceServer(s *grpc.Server, srv FirewallServiceServer) {
	s.RegisterService(&_FirewallService_serviceDesc, srv)
}

func _FirewallService_ComposeFirewallGlobalBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeFirewallGlobalBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirewallServiceServer).ComposeFirewallGlobalBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FirewallService/ComposeFirewallGlobalBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirewallServiceServer).ComposeFirewallGlobalBoard(ctx, req.(*ComposeFirewallGlobalBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirewallService_NotifyHTTPFirewallEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyHTTPFirewallEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirewallServiceServer).NotifyHTTPFirewallEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FirewallService/NotifyHTTPFirewallEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirewallServiceServer).NotifyHTTPFirewallEvent(ctx, req.(*NotifyHTTPFirewallEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FirewallService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FirewallService",
	HandlerType: (*FirewallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "composeFirewallGlobalBoard",
			Handler:    _FirewallService_ComposeFirewallGlobalBoard_Handler,
		},
		{
			MethodName: "notifyHTTPFirewallEvent",
			Handler:    _FirewallService_NotifyHTTPFirewallEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_firewall.proto",
}
