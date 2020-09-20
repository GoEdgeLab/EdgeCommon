// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: model_http_web.proto

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

type HTTPWeb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn                   bool   `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Root                   string `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	Charset                string `protobuf:"bytes,4,opt,name=charset,proto3" json:"charset,omitempty"`
	RequestHeaderPolicyId  int64  `protobuf:"varint,5,opt,name=requestHeaderPolicyId,proto3" json:"requestHeaderPolicyId,omitempty"`
	ResponseHeaderPolicyId int64  `protobuf:"varint,6,opt,name=responseHeaderPolicyId,proto3" json:"responseHeaderPolicyId,omitempty"`
}

func (x *HTTPWeb) Reset() {
	*x = HTTPWeb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_http_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPWeb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPWeb) ProtoMessage() {}

func (x *HTTPWeb) ProtoReflect() protoreflect.Message {
	mi := &file_model_http_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPWeb.ProtoReflect.Descriptor instead.
func (*HTTPWeb) Descriptor() ([]byte, []int) {
	return file_model_http_web_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPWeb) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HTTPWeb) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *HTTPWeb) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *HTTPWeb) GetCharset() string {
	if x != nil {
		return x.Charset
	}
	return ""
}

func (x *HTTPWeb) GetRequestHeaderPolicyId() int64 {
	if x != nil {
		return x.RequestHeaderPolicyId
	}
	return 0
}

func (x *HTTPWeb) GetResponseHeaderPolicyId() int64 {
	if x != nil {
		return x.ResponseHeaderPolicyId
	}
	return 0
}

var File_model_http_web_proto protoreflect.FileDescriptor

var file_model_http_web_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x77, 0x65, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xc9, 0x01, 0x0a, 0x07, 0x48,
	0x54, 0x54, 0x50, 0x57, 0x65, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x36,
	0x0a, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_http_web_proto_rawDescOnce sync.Once
	file_model_http_web_proto_rawDescData = file_model_http_web_proto_rawDesc
)

func file_model_http_web_proto_rawDescGZIP() []byte {
	file_model_http_web_proto_rawDescOnce.Do(func() {
		file_model_http_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_http_web_proto_rawDescData)
	})
	return file_model_http_web_proto_rawDescData
}

var file_model_http_web_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_http_web_proto_goTypes = []interface{}{
	(*HTTPWeb)(nil), // 0: pb.HTTPWeb
}
var file_model_http_web_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_http_web_proto_init() }
func file_model_http_web_proto_init() {
	if File_model_http_web_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_http_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPWeb); i {
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
			RawDescriptor: file_model_http_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_http_web_proto_goTypes,
		DependencyIndexes: file_model_http_web_proto_depIdxs,
		MessageInfos:      file_model_http_web_proto_msgTypes,
	}.Build()
	File_model_http_web_proto = out.File
	file_model_http_web_proto_rawDesc = nil
	file_model_http_web_proto_goTypes = nil
	file_model_http_web_proto_depIdxs = nil
}
