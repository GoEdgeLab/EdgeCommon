// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username               string         `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Fullname               string         `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Mobile                 string         `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Tel                    string         `protobuf:"bytes,5,opt,name=tel,proto3" json:"tel,omitempty"`
	Email                  string         `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Remark                 string         `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	IsOn                   bool           `protobuf:"varint,8,opt,name=isOn,proto3" json:"isOn,omitempty"`
	CreatedAt              int64          `protobuf:"varint,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	RegisteredIP           string         `protobuf:"bytes,12,opt,name=registeredIP,proto3" json:"registeredIP,omitempty"`
	IsVerified             bool           `protobuf:"varint,13,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	IsRejected             bool           `protobuf:"varint,14,opt,name=isRejected,proto3" json:"isRejected,omitempty"`
	RejectReason           string         `protobuf:"bytes,15,opt,name=rejectReason,proto3" json:"rejectReason,omitempty"`
	IsDeleted              bool           `protobuf:"varint,16,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	IsIndividualIdentified bool           `protobuf:"varint,17,opt,name=isIndividualIdentified,proto3" json:"isIndividualIdentified,omitempty"`
	IsEnterpriseIdentified bool           `protobuf:"varint,18,opt,name=isEnterpriseIdentified,proto3" json:"isEnterpriseIdentified,omitempty"`
	OtpLogin               *Login         `protobuf:"bytes,19,opt,name=otpLogin,proto3" json:"otpLogin,omitempty"` // OTP认证
	NodeCluster            *NodeCluster   `protobuf:"bytes,10,opt,name=nodeCluster,proto3" json:"nodeCluster,omitempty"`
	Features               []*UserFeature `protobuf:"bytes,11,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_models_model_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *User) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *User) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *User) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *User) GetRegisteredIP() string {
	if x != nil {
		return x.RegisteredIP
	}
	return ""
}

func (x *User) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *User) GetIsRejected() bool {
	if x != nil {
		return x.IsRejected
	}
	return false
}

func (x *User) GetRejectReason() string {
	if x != nil {
		return x.RejectReason
	}
	return ""
}

func (x *User) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *User) GetIsIndividualIdentified() bool {
	if x != nil {
		return x.IsIndividualIdentified
	}
	return false
}

func (x *User) GetIsEnterpriseIdentified() bool {
	if x != nil {
		return x.IsEnterpriseIdentified
	}
	return false
}

func (x *User) GetOtpLogin() *Login {
	if x != nil {
		return x.OtpLogin
	}
	return nil
}

func (x *User) GetNodeCluster() *NodeCluster {
	if x != nil {
		return x.NodeCluster
	}
	return nil
}

func (x *User) GetFeatures() []*UserFeature {
	if x != nil {
		return x.Features
	}
	return nil
}

var File_models_model_user_proto protoreflect.FileDescriptor

var file_models_model_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x04, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x49, 0x50, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x49, 0x50, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x69, 0x73,
	0x49, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x49, 0x6e,
	0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x69, 0x73, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x6f, 0x74,
	0x70, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x08, 0x6f, 0x74, 0x70, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_models_model_user_proto_rawDescOnce sync.Once
	file_models_model_user_proto_rawDescData = file_models_model_user_proto_rawDesc
)

func file_models_model_user_proto_rawDescGZIP() []byte {
	file_models_model_user_proto_rawDescOnce.Do(func() {
		file_models_model_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_proto_rawDescData)
	})
	return file_models_model_user_proto_rawDescData
}

var file_models_model_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_proto_goTypes = []interface{}{
	(*User)(nil),        // 0: pb.User
	(*Login)(nil),       // 1: pb.Login
	(*NodeCluster)(nil), // 2: pb.NodeCluster
	(*UserFeature)(nil), // 3: pb.UserFeature
}
var file_models_model_user_proto_depIdxs = []int32{
	1, // 0: pb.User.otpLogin:type_name -> pb.Login
	2, // 1: pb.User.nodeCluster:type_name -> pb.NodeCluster
	3, // 2: pb.User.features:type_name -> pb.UserFeature
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_user_proto_init() }
func file_models_model_user_proto_init() {
	if File_models_model_user_proto != nil {
		return
	}
	file_models_model_node_cluster_proto_init()
	file_models_model_user_feature_proto_init()
	file_models_model_login_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_models_model_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_proto_goTypes,
		DependencyIndexes: file_models_model_user_proto_depIdxs,
		MessageInfos:      file_models_model_user_proto_msgTypes,
	}.Build()
	File_models_model_user_proto = out.File
	file_models_model_user_proto_rawDesc = nil
	file_models_model_user_proto_goTypes = nil
	file_models_model_user_proto_depIdxs = nil
}
