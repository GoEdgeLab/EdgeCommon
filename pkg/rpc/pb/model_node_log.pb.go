// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: model_node_log.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type NodeLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role        string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Tag         string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Level       string `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	NodeId      int64  `protobuf:"varint,5,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	CreatedAt   int64  `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *NodeLog) Reset() {
	*x = NodeLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_node_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeLog) ProtoMessage() {}

func (x *NodeLog) ProtoReflect() protoreflect.Message {
	mi := &file_model_node_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeLog.ProtoReflect.Descriptor instead.
func (*NodeLog) Descriptor() ([]byte, []int) {
	return file_model_node_log_proto_rawDescGZIP(), []int{0}
}

func (x *NodeLog) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *NodeLog) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *NodeLog) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NodeLog) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *NodeLog) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *NodeLog) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_model_node_log_proto protoreflect.FileDescriptor

var file_model_node_log_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x9d, 0x01, 0x0a, 0x07, 0x4e,
	0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_node_log_proto_rawDescOnce sync.Once
	file_model_node_log_proto_rawDescData = file_model_node_log_proto_rawDesc
)

func file_model_node_log_proto_rawDescGZIP() []byte {
	file_model_node_log_proto_rawDescOnce.Do(func() {
		file_model_node_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_node_log_proto_rawDescData)
	})
	return file_model_node_log_proto_rawDescData
}

var file_model_node_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_node_log_proto_goTypes = []interface{}{
	(*NodeLog)(nil), // 0: pb.NodeLog
}
var file_model_node_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_node_log_proto_init() }
func file_model_node_log_proto_init() {
	if File_model_node_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_node_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeLog); i {
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
			RawDescriptor: file_model_node_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_node_log_proto_goTypes,
		DependencyIndexes: file_model_node_log_proto_depIdxs,
		MessageInfos:      file_model_node_log_proto_msgTypes,
	}.Build()
	File_model_node_log_proto = out.File
	file_model_node_log_proto_rawDesc = nil
	file_model_node_log_proto_goTypes = nil
	file_model_node_log_proto_depIdxs = nil
}
