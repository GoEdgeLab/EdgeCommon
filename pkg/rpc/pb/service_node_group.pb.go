// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_node_group.proto

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

// 创建分组
type CreateNodeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64  `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNodeGroupRequest) Reset() {
	*x = CreateNodeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeGroupRequest) ProtoMessage() {}

func (x *CreateNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNodeGroupRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

func (x *CreateNodeGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateNodeGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *CreateNodeGroupResponse) Reset() {
	*x = CreateNodeGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeGroupResponse) ProtoMessage() {}

func (x *CreateNodeGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateNodeGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNodeGroupResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// 修改分组
type UpdateNodeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64  `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateNodeGroupRequest) Reset() {
	*x = UpdateNodeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeGroupRequest) ProtoMessage() {}

func (x *UpdateNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNodeGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UpdateNodeGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 删除分组
type DeleteNodeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *DeleteNodeGroupRequest) Reset() {
	*x = DeleteNodeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeGroupRequest) ProtoMessage() {}

func (x *DeleteNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNodeGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// 查询所有分组
type FindAllEnabledNodeGroupsWithNodeClusterIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64 `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
}

func (x *FindAllEnabledNodeGroupsWithNodeClusterIdRequest) Reset() {
	*x = FindAllEnabledNodeGroupsWithNodeClusterIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledNodeGroupsWithNodeClusterIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledNodeGroupsWithNodeClusterIdRequest) ProtoMessage() {}

func (x *FindAllEnabledNodeGroupsWithNodeClusterIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledNodeGroupsWithNodeClusterIdRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledNodeGroupsWithNodeClusterIdRequest) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllEnabledNodeGroupsWithNodeClusterIdRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

type FindAllEnabledNodeGroupsWithNodeClusterIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*NodeGroup `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *FindAllEnabledNodeGroupsWithNodeClusterIdResponse) Reset() {
	*x = FindAllEnabledNodeGroupsWithNodeClusterIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledNodeGroupsWithNodeClusterIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse) ProtoMessage() {}

func (x *FindAllEnabledNodeGroupsWithNodeClusterIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledNodeGroupsWithNodeClusterIdResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllEnabledNodeGroupsWithNodeClusterIdResponse) GetGroups() []*NodeGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

// 修改分组排序
type UpdateNodeGroupOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupIds []int64 `protobuf:"varint,1,rep,packed,name=groupIds,proto3" json:"groupIds,omitempty"`
}

func (x *UpdateNodeGroupOrdersRequest) Reset() {
	*x = UpdateNodeGroupOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeGroupOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeGroupOrdersRequest) ProtoMessage() {}

func (x *UpdateNodeGroupOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeGroupOrdersRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeGroupOrdersRequest) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNodeGroupOrdersRequest) GetGroupIds() []int64 {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

// 查找单个分组信息
type FindEnabledNodeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *FindEnabledNodeGroupRequest) Reset() {
	*x = FindEnabledNodeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledNodeGroupRequest) ProtoMessage() {}

func (x *FindEnabledNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledNodeGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type FindEnabledNodeGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *NodeGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *FindEnabledNodeGroupResponse) Reset() {
	*x = FindEnabledNodeGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledNodeGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledNodeGroupResponse) ProtoMessage() {}

func (x *FindEnabledNodeGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledNodeGroupResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledNodeGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_node_group_proto_rawDescGZIP(), []int{8}
}

func (x *FindEnabledNodeGroupResponse) GetGroup() *NodeGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

var File_service_node_group_proto protoreflect.FileDescriptor

var file_service_node_group_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x46, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x58, 0x0a,
	0x30, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x31, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x22, 0x3a, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x22,
	0x37, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x32, 0x9d, 0x04,
	0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a,
	0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x98, 0x01, 0x0a,
	0x29, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x15, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x59, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_node_group_proto_rawDescOnce sync.Once
	file_service_node_group_proto_rawDescData = file_service_node_group_proto_rawDesc
)

func file_service_node_group_proto_rawDescGZIP() []byte {
	file_service_node_group_proto_rawDescOnce.Do(func() {
		file_service_node_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_node_group_proto_rawDescData)
	})
	return file_service_node_group_proto_rawDescData
}

var file_service_node_group_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_node_group_proto_goTypes = []interface{}{
	(*CreateNodeGroupRequest)(nil),                            // 0: pb.CreateNodeGroupRequest
	(*CreateNodeGroupResponse)(nil),                           // 1: pb.CreateNodeGroupResponse
	(*UpdateNodeGroupRequest)(nil),                            // 2: pb.UpdateNodeGroupRequest
	(*DeleteNodeGroupRequest)(nil),                            // 3: pb.DeleteNodeGroupRequest
	(*FindAllEnabledNodeGroupsWithNodeClusterIdRequest)(nil),  // 4: pb.FindAllEnabledNodeGroupsWithNodeClusterIdRequest
	(*FindAllEnabledNodeGroupsWithNodeClusterIdResponse)(nil), // 5: pb.FindAllEnabledNodeGroupsWithNodeClusterIdResponse
	(*UpdateNodeGroupOrdersRequest)(nil),                      // 6: pb.UpdateNodeGroupOrdersRequest
	(*FindEnabledNodeGroupRequest)(nil),                       // 7: pb.FindEnabledNodeGroupRequest
	(*FindEnabledNodeGroupResponse)(nil),                      // 8: pb.FindEnabledNodeGroupResponse
	(*NodeGroup)(nil),                                         // 9: pb.NodeGroup
	(*RPCSuccess)(nil),                                        // 10: pb.RPCSuccess
}
var file_service_node_group_proto_depIdxs = []int32{
	9,  // 0: pb.FindAllEnabledNodeGroupsWithNodeClusterIdResponse.groups:type_name -> pb.NodeGroup
	9,  // 1: pb.FindEnabledNodeGroupResponse.group:type_name -> pb.NodeGroup
	0,  // 2: pb.NodeGroupService.createNodeGroup:input_type -> pb.CreateNodeGroupRequest
	2,  // 3: pb.NodeGroupService.updateNodeGroup:input_type -> pb.UpdateNodeGroupRequest
	3,  // 4: pb.NodeGroupService.deleteNodeGroup:input_type -> pb.DeleteNodeGroupRequest
	4,  // 5: pb.NodeGroupService.findAllEnabledNodeGroupsWithNodeClusterId:input_type -> pb.FindAllEnabledNodeGroupsWithNodeClusterIdRequest
	6,  // 6: pb.NodeGroupService.updateNodeGroupOrders:input_type -> pb.UpdateNodeGroupOrdersRequest
	7,  // 7: pb.NodeGroupService.findEnabledNodeGroup:input_type -> pb.FindEnabledNodeGroupRequest
	1,  // 8: pb.NodeGroupService.createNodeGroup:output_type -> pb.CreateNodeGroupResponse
	10, // 9: pb.NodeGroupService.updateNodeGroup:output_type -> pb.RPCSuccess
	10, // 10: pb.NodeGroupService.deleteNodeGroup:output_type -> pb.RPCSuccess
	5,  // 11: pb.NodeGroupService.findAllEnabledNodeGroupsWithNodeClusterId:output_type -> pb.FindAllEnabledNodeGroupsWithNodeClusterIdResponse
	10, // 12: pb.NodeGroupService.updateNodeGroupOrders:output_type -> pb.RPCSuccess
	8,  // 13: pb.NodeGroupService.findEnabledNodeGroup:output_type -> pb.FindEnabledNodeGroupResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_node_group_proto_init() }
func file_service_node_group_proto_init() {
	if File_service_node_group_proto != nil {
		return
	}
	file_models_model_node_group_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_node_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeGroupRequest); i {
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
		file_service_node_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeGroupResponse); i {
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
		file_service_node_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeGroupRequest); i {
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
		file_service_node_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeGroupRequest); i {
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
		file_service_node_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledNodeGroupsWithNodeClusterIdRequest); i {
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
		file_service_node_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledNodeGroupsWithNodeClusterIdResponse); i {
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
		file_service_node_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeGroupOrdersRequest); i {
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
		file_service_node_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledNodeGroupRequest); i {
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
		file_service_node_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledNodeGroupResponse); i {
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
			RawDescriptor: file_service_node_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_node_group_proto_goTypes,
		DependencyIndexes: file_service_node_group_proto_depIdxs,
		MessageInfos:      file_service_node_group_proto_msgTypes,
	}.Build()
	File_service_node_group_proto = out.File
	file_service_node_group_proto_rawDesc = nil
	file_service_node_group_proto_goTypes = nil
	file_service_node_group_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeGroupServiceClient is the client API for NodeGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeGroupServiceClient interface {
	// 创建分组
	CreateNodeGroup(ctx context.Context, in *CreateNodeGroupRequest, opts ...grpc.CallOption) (*CreateNodeGroupResponse, error)
	// 修改分组
	UpdateNodeGroup(ctx context.Context, in *UpdateNodeGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除分组
	DeleteNodeGroup(ctx context.Context, in *DeleteNodeGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询所有分组
	FindAllEnabledNodeGroupsWithNodeClusterId(ctx context.Context, in *FindAllEnabledNodeGroupsWithNodeClusterIdRequest, opts ...grpc.CallOption) (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse, error)
	// 修改分组排序
	UpdateNodeGroupOrders(ctx context.Context, in *UpdateNodeGroupOrdersRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个分组信息
	FindEnabledNodeGroup(ctx context.Context, in *FindEnabledNodeGroupRequest, opts ...grpc.CallOption) (*FindEnabledNodeGroupResponse, error)
}

type nodeGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeGroupServiceClient(cc grpc.ClientConnInterface) NodeGroupServiceClient {
	return &nodeGroupServiceClient{cc}
}

func (c *nodeGroupServiceClient) CreateNodeGroup(ctx context.Context, in *CreateNodeGroupRequest, opts ...grpc.CallOption) (*CreateNodeGroupResponse, error) {
	out := new(CreateNodeGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGroupService/createNodeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) UpdateNodeGroup(ctx context.Context, in *UpdateNodeGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeGroupService/updateNodeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) DeleteNodeGroup(ctx context.Context, in *DeleteNodeGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeGroupService/deleteNodeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) FindAllEnabledNodeGroupsWithNodeClusterId(ctx context.Context, in *FindAllEnabledNodeGroupsWithNodeClusterIdRequest, opts ...grpc.CallOption) (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse, error) {
	out := new(FindAllEnabledNodeGroupsWithNodeClusterIdResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGroupService/findAllEnabledNodeGroupsWithNodeClusterId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) UpdateNodeGroupOrders(ctx context.Context, in *UpdateNodeGroupOrdersRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeGroupService/updateNodeGroupOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) FindEnabledNodeGroup(ctx context.Context, in *FindEnabledNodeGroupRequest, opts ...grpc.CallOption) (*FindEnabledNodeGroupResponse, error) {
	out := new(FindEnabledNodeGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGroupService/findEnabledNodeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeGroupServiceServer is the server API for NodeGroupService service.
type NodeGroupServiceServer interface {
	// 创建分组
	CreateNodeGroup(context.Context, *CreateNodeGroupRequest) (*CreateNodeGroupResponse, error)
	// 修改分组
	UpdateNodeGroup(context.Context, *UpdateNodeGroupRequest) (*RPCSuccess, error)
	// 删除分组
	DeleteNodeGroup(context.Context, *DeleteNodeGroupRequest) (*RPCSuccess, error)
	// 查询所有分组
	FindAllEnabledNodeGroupsWithNodeClusterId(context.Context, *FindAllEnabledNodeGroupsWithNodeClusterIdRequest) (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse, error)
	// 修改分组排序
	UpdateNodeGroupOrders(context.Context, *UpdateNodeGroupOrdersRequest) (*RPCSuccess, error)
	// 查找单个分组信息
	FindEnabledNodeGroup(context.Context, *FindEnabledNodeGroupRequest) (*FindEnabledNodeGroupResponse, error)
}

// UnimplementedNodeGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeGroupServiceServer struct {
}

func (*UnimplementedNodeGroupServiceServer) CreateNodeGroup(context.Context, *CreateNodeGroupRequest) (*CreateNodeGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeGroup not implemented")
}
func (*UnimplementedNodeGroupServiceServer) UpdateNodeGroup(context.Context, *UpdateNodeGroupRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeGroup not implemented")
}
func (*UnimplementedNodeGroupServiceServer) DeleteNodeGroup(context.Context, *DeleteNodeGroupRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeGroup not implemented")
}
func (*UnimplementedNodeGroupServiceServer) FindAllEnabledNodeGroupsWithNodeClusterId(context.Context, *FindAllEnabledNodeGroupsWithNodeClusterIdRequest) (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledNodeGroupsWithNodeClusterId not implemented")
}
func (*UnimplementedNodeGroupServiceServer) UpdateNodeGroupOrders(context.Context, *UpdateNodeGroupOrdersRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeGroupOrders not implemented")
}
func (*UnimplementedNodeGroupServiceServer) FindEnabledNodeGroup(context.Context, *FindEnabledNodeGroupRequest) (*FindEnabledNodeGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledNodeGroup not implemented")
}

func RegisterNodeGroupServiceServer(s *grpc.Server, srv NodeGroupServiceServer) {
	s.RegisterService(&_NodeGroupService_serviceDesc, srv)
}

func _NodeGroupService_CreateNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).CreateNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGroupService/CreateNodeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).CreateNodeGroup(ctx, req.(*CreateNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_UpdateNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).UpdateNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGroupService/UpdateNodeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).UpdateNodeGroup(ctx, req.(*UpdateNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_DeleteNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).DeleteNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGroupService/DeleteNodeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).DeleteNodeGroup(ctx, req.(*DeleteNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_FindAllEnabledNodeGroupsWithNodeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledNodeGroupsWithNodeClusterIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).FindAllEnabledNodeGroupsWithNodeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGroupService/FindAllEnabledNodeGroupsWithNodeClusterId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).FindAllEnabledNodeGroupsWithNodeClusterId(ctx, req.(*FindAllEnabledNodeGroupsWithNodeClusterIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_UpdateNodeGroupOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeGroupOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).UpdateNodeGroupOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGroupService/UpdateNodeGroupOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).UpdateNodeGroupOrders(ctx, req.(*UpdateNodeGroupOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_FindEnabledNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).FindEnabledNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGroupService/FindEnabledNodeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).FindEnabledNodeGroup(ctx, req.(*FindEnabledNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeGroupService",
	HandlerType: (*NodeGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNodeGroup",
			Handler:    _NodeGroupService_CreateNodeGroup_Handler,
		},
		{
			MethodName: "updateNodeGroup",
			Handler:    _NodeGroupService_UpdateNodeGroup_Handler,
		},
		{
			MethodName: "deleteNodeGroup",
			Handler:    _NodeGroupService_DeleteNodeGroup_Handler,
		},
		{
			MethodName: "findAllEnabledNodeGroupsWithNodeClusterId",
			Handler:    _NodeGroupService_FindAllEnabledNodeGroupsWithNodeClusterId_Handler,
		},
		{
			MethodName: "updateNodeGroupOrders",
			Handler:    _NodeGroupService_UpdateNodeGroupOrders_Handler,
		},
		{
			MethodName: "findEnabledNodeGroup",
			Handler:    _NodeGroupService_FindEnabledNodeGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_group.proto",
}
