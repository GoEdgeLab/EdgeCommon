// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_ns_domain.proto

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

// DNS域名
type NSDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                   string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsOn                   bool             `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`
	CreatedAt              int64            `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	IsDeleted              bool             `protobuf:"varint,5,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	Version                int64            `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	TsigJSON               []byte           `protobuf:"bytes,7,opt,name=tsigJSON,proto3" json:"tsigJSON,omitempty"`
	NsDomainGroupIds       []int64          `protobuf:"varint,8,rep,packed,name=nsDomainGroupIds,proto3" json:"nsDomainGroupIds,omitempty"`
	Status                 string           `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	UserId                 int64            `protobuf:"varint,10,opt,name=userId,proto3" json:"userId,omitempty"`                                // 用户ID
	RecordsHealthCheckJSON []byte           `protobuf:"bytes,11,opt,name=recordsHealthCheckJSON,proto3" json:"recordsHealthCheckJSON,omitempty"` // 健康检查设置
	NsCluster              *NSCluster       `protobuf:"bytes,30,opt,name=nsCluster,proto3" json:"nsCluster,omitempty"`
	User                   *User            `protobuf:"bytes,31,opt,name=user,proto3" json:"user,omitempty"`
	NsDomainGroups         []*NSDomainGroup `protobuf:"bytes,32,rep,name=nsDomainGroups,proto3" json:"nsDomainGroups,omitempty"`
}

func (x *NSDomain) Reset() {
	*x = NSDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ns_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSDomain) ProtoMessage() {}

func (x *NSDomain) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ns_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSDomain.ProtoReflect.Descriptor instead.
func (*NSDomain) Descriptor() ([]byte, []int) {
	return file_models_model_ns_domain_proto_rawDescGZIP(), []int{0}
}

func (x *NSDomain) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NSDomain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NSDomain) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NSDomain) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NSDomain) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *NSDomain) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *NSDomain) GetTsigJSON() []byte {
	if x != nil {
		return x.TsigJSON
	}
	return nil
}

func (x *NSDomain) GetNsDomainGroupIds() []int64 {
	if x != nil {
		return x.NsDomainGroupIds
	}
	return nil
}

func (x *NSDomain) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NSDomain) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *NSDomain) GetRecordsHealthCheckJSON() []byte {
	if x != nil {
		return x.RecordsHealthCheckJSON
	}
	return nil
}

func (x *NSDomain) GetNsCluster() *NSCluster {
	if x != nil {
		return x.NsCluster
	}
	return nil
}

func (x *NSDomain) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *NSDomain) GetNsDomainGroups() []*NSDomainGroup {
	if x != nil {
		return x.NsDomainGroups
	}
	return nil
}

var File_models_model_ns_domain_proto protoreflect.FileDescriptor

var file_models_model_ns_domain_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x6e, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6e, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce,
	0x03, 0x0a, 0x08, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x73, 0x69,
	0x67, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x73, 0x69,
	0x67, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x10, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x36, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x16, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x2b, 0x0a, 0x09, 0x6e, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x4e, 0x53, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x09, 0x6e, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0e, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x0e, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ns_domain_proto_rawDescOnce sync.Once
	file_models_model_ns_domain_proto_rawDescData = file_models_model_ns_domain_proto_rawDesc
)

func file_models_model_ns_domain_proto_rawDescGZIP() []byte {
	file_models_model_ns_domain_proto_rawDescOnce.Do(func() {
		file_models_model_ns_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ns_domain_proto_rawDescData)
	})
	return file_models_model_ns_domain_proto_rawDescData
}

var file_models_model_ns_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ns_domain_proto_goTypes = []interface{}{
	(*NSDomain)(nil),      // 0: pb.NSDomain
	(*NSCluster)(nil),     // 1: pb.NSCluster
	(*User)(nil),          // 2: pb.User
	(*NSDomainGroup)(nil), // 3: pb.NSDomainGroup
}
var file_models_model_ns_domain_proto_depIdxs = []int32{
	1, // 0: pb.NSDomain.nsCluster:type_name -> pb.NSCluster
	2, // 1: pb.NSDomain.user:type_name -> pb.User
	3, // 2: pb.NSDomain.nsDomainGroups:type_name -> pb.NSDomainGroup
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_ns_domain_proto_init() }
func file_models_model_ns_domain_proto_init() {
	if File_models_model_ns_domain_proto != nil {
		return
	}
	file_models_model_ns_cluster_proto_init()
	file_models_model_ns_domain_group_proto_init()
	file_models_model_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ns_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSDomain); i {
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
			RawDescriptor: file_models_model_ns_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ns_domain_proto_goTypes,
		DependencyIndexes: file_models_model_ns_domain_proto_depIdxs,
		MessageInfos:      file_models_model_ns_domain_proto_msgTypes,
	}.Build()
	File_models_model_ns_domain_proto = out.File
	file_models_model_ns_domain_proto_rawDesc = nil
	file_models_model_ns_domain_proto_goTypes = nil
	file_models_model_ns_domain_proto_depIdxs = nil
}
