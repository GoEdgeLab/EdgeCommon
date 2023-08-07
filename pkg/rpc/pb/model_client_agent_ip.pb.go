// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_client_agent_ip.proto

package pb

import (
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

type ClientAgentIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip          string       `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Ptr         string       `protobuf:"bytes,3,opt,name=ptr,proto3" json:"ptr,omitempty"`
	ClientAgent *ClientAgent `protobuf:"bytes,30,opt,name=clientAgent,proto3" json:"clientAgent,omitempty"`
}

func (x *ClientAgentIP) Reset() {
	*x = ClientAgentIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_client_agent_ip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAgentIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAgentIP) ProtoMessage() {}

func (x *ClientAgentIP) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_client_agent_ip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAgentIP.ProtoReflect.Descriptor instead.
func (*ClientAgentIP) Descriptor() ([]byte, []int) {
	return file_models_model_client_agent_ip_proto_rawDescGZIP(), []int{0}
}

func (x *ClientAgentIP) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClientAgentIP) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ClientAgentIP) GetPtr() string {
	if x != nil {
		return x.Ptr
	}
	return ""
}

func (x *ClientAgentIP) GetClientAgent() *ClientAgent {
	if x != nil {
		return x.ClientAgent
	}
	return nil
}

var File_models_model_client_agent_ip_proto protoreflect.FileDescriptor

var file_models_model_client_agent_ip_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x74,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x74, 0x72, 0x12, 0x31, 0x0a, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_client_agent_ip_proto_rawDescOnce sync.Once
	file_models_model_client_agent_ip_proto_rawDescData = file_models_model_client_agent_ip_proto_rawDesc
)

func file_models_model_client_agent_ip_proto_rawDescGZIP() []byte {
	file_models_model_client_agent_ip_proto_rawDescOnce.Do(func() {
		file_models_model_client_agent_ip_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_client_agent_ip_proto_rawDescData)
	})
	return file_models_model_client_agent_ip_proto_rawDescData
}

var file_models_model_client_agent_ip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_client_agent_ip_proto_goTypes = []interface{}{
	(*ClientAgentIP)(nil), // 0: pb.ClientAgentIP
	(*ClientAgent)(nil),   // 1: pb.ClientAgent
}
var file_models_model_client_agent_ip_proto_depIdxs = []int32{
	1, // 0: pb.ClientAgentIP.clientAgent:type_name -> pb.ClientAgent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_client_agent_ip_proto_init() }
func file_models_model_client_agent_ip_proto_init() {
	if File_models_model_client_agent_ip_proto != nil {
		return
	}
	file_models_model_client_agent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_client_agent_ip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAgentIP); i {
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
			RawDescriptor: file_models_model_client_agent_ip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_client_agent_ip_proto_goTypes,
		DependencyIndexes: file_models_model_client_agent_ip_proto_depIdxs,
		MessageInfos:      file_models_model_client_agent_ip_proto_msgTypes,
	}.Build()
	File_models_model_client_agent_ip_proto = out.File
	file_models_model_client_agent_ip_proto_rawDesc = nil
	file_models_model_client_agent_ip_proto_goTypes = nil
	file_models_model_client_agent_ip_proto_depIdxs = nil
}
