// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_user_bill.proto

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

type UserBill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User        *User   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Type        string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	TypeName    string  `protobuf:"bytes,4,opt,name=typeName,proto3" json:"typeName,omitempty"`
	Description string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Amount      float32 `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Month       string  `protobuf:"bytes,7,opt,name=month,proto3" json:"month,omitempty"`
	IsPaid      bool    `protobuf:"varint,8,opt,name=isPaid,proto3" json:"isPaid,omitempty"`
	PaidAt      int64   `protobuf:"varint,9,opt,name=paidAt,proto3" json:"paidAt,omitempty"`
	Code        string  `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	CanPay      bool    `protobuf:"varint,11,opt,name=canPay,proto3" json:"canPay,omitempty"`
	DayFrom     string  `protobuf:"bytes,12,opt,name=dayFrom,proto3" json:"dayFrom,omitempty"`
	DayTo       string  `protobuf:"bytes,13,opt,name=dayTo,proto3" json:"dayTo,omitempty"`
	PricePeriod string  `protobuf:"bytes,14,opt,name=pricePeriod,proto3" json:"pricePeriod,omitempty"`
	IsOverdue   bool    `protobuf:"varint,15,opt,name=isOverdue,proto3" json:"isOverdue,omitempty"` // 是否已逾期
}

func (x *UserBill) Reset() {
	*x = UserBill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_bill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBill) ProtoMessage() {}

func (x *UserBill) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_bill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBill.ProtoReflect.Descriptor instead.
func (*UserBill) Descriptor() ([]byte, []int) {
	return file_models_model_user_bill_proto_rawDescGZIP(), []int{0}
}

func (x *UserBill) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserBill) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserBill) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserBill) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *UserBill) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UserBill) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UserBill) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *UserBill) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *UserBill) GetPaidAt() int64 {
	if x != nil {
		return x.PaidAt
	}
	return 0
}

func (x *UserBill) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserBill) GetCanPay() bool {
	if x != nil {
		return x.CanPay
	}
	return false
}

func (x *UserBill) GetDayFrom() string {
	if x != nil {
		return x.DayFrom
	}
	return ""
}

func (x *UserBill) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *UserBill) GetPricePeriod() string {
	if x != nil {
		return x.PricePeriod
	}
	return ""
}

func (x *UserBill) GetIsOverdue() bool {
	if x != nil {
		return x.IsOverdue
	}
	return false
}

var File_models_model_user_bill_proto protoreflect.FileDescriptor

var file_models_model_user_bill_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x69, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x70, 0x61, 0x69, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61,
	0x6e, 0x50, 0x61, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x50,
	0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x61, 0x79, 0x54, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x79,
	0x54, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x64, 0x75,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x64,
	0x75, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_models_model_user_bill_proto_rawDescOnce sync.Once
	file_models_model_user_bill_proto_rawDescData = file_models_model_user_bill_proto_rawDesc
)

func file_models_model_user_bill_proto_rawDescGZIP() []byte {
	file_models_model_user_bill_proto_rawDescOnce.Do(func() {
		file_models_model_user_bill_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_bill_proto_rawDescData)
	})
	return file_models_model_user_bill_proto_rawDescData
}

var file_models_model_user_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_bill_proto_goTypes = []interface{}{
	(*UserBill)(nil), // 0: pb.UserBill
	(*User)(nil),     // 1: pb.User
}
var file_models_model_user_bill_proto_depIdxs = []int32{
	1, // 0: pb.UserBill.user:type_name -> pb.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_user_bill_proto_init() }
func file_models_model_user_bill_proto_init() {
	if File_models_model_user_bill_proto != nil {
		return
	}
	file_models_model_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_bill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBill); i {
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
			RawDescriptor: file_models_model_user_bill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_bill_proto_goTypes,
		DependencyIndexes: file_models_model_user_bill_proto_depIdxs,
		MessageInfos:      file_models_model_user_bill_proto_msgTypes,
	}.Build()
	File_models_model_user_bill_proto = out.File
	file_models_model_user_bill_proto_rawDesc = nil
	file_models_model_user_bill_proto_goTypes = nil
	file_models_model_user_bill_proto_depIdxs = nil
}
