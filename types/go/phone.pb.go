// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: phone.proto

package userpb

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

type Phone_PhoneType int32

const (
	Phone_UNKNOWN Phone_PhoneType = 0
	Phone_MOBILE  Phone_PhoneType = 1
	Phone_HOME    Phone_PhoneType = 2
	Phone_WORK    Phone_PhoneType = 3
)

// Enum value maps for Phone_PhoneType.
var (
	Phone_PhoneType_name = map[int32]string{
		0: "UNKNOWN",
		1: "MOBILE",
		2: "HOME",
		3: "WORK",
	}
	Phone_PhoneType_value = map[string]int32{
		"UNKNOWN": 0,
		"MOBILE":  1,
		"HOME":    2,
		"WORK":    3,
	}
)

func (x Phone_PhoneType) Enum() *Phone_PhoneType {
	p := new(Phone_PhoneType)
	*p = x
	return p
}

func (x Phone_PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Phone_PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_phone_proto_enumTypes[0].Descriptor()
}

func (Phone_PhoneType) Type() protoreflect.EnumType {
	return &file_phone_proto_enumTypes[0]
}

func (x Phone_PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Phone_PhoneType.Descriptor instead.
func (Phone_PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_phone_proto_rawDescGZIP(), []int{0, 0}
}

type Phone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string          `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type   Phone_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=user.Phone_PhoneType" json:"type,omitempty"`
}

func (x *Phone) Reset() {
	*x = Phone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Phone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phone) ProtoMessage() {}

func (x *Phone) ProtoReflect() protoreflect.Message {
	mi := &file_phone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phone.ProtoReflect.Descriptor instead.
func (*Phone) Descriptor() ([]byte, []int) {
	return file_phone_proto_rawDescGZIP(), []int{0}
}

func (x *Phone) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Phone) GetType() Phone_PhoneType {
	if x != nil {
		return x.Type
	}
	return Phone_UNKNOWN
}

var File_phone_proto protoreflect.FileDescriptor

var file_phone_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x38, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f,
	0x42, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x03, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x65, 0x6f, 0x61,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_phone_proto_rawDescOnce sync.Once
	file_phone_proto_rawDescData = file_phone_proto_rawDesc
)

func file_phone_proto_rawDescGZIP() []byte {
	file_phone_proto_rawDescOnce.Do(func() {
		file_phone_proto_rawDescData = protoimpl.X.CompressGZIP(file_phone_proto_rawDescData)
	})
	return file_phone_proto_rawDescData
}

var file_phone_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_phone_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_phone_proto_goTypes = []interface{}{
	(Phone_PhoneType)(0), // 0: user.Phone.PhoneType
	(*Phone)(nil),        // 1: user.Phone
}
var file_phone_proto_depIdxs = []int32{
	0, // 0: user.Phone.type:type_name -> user.Phone.PhoneType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_phone_proto_init() }
func file_phone_proto_init() {
	if File_phone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_phone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Phone); i {
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
			RawDescriptor: file_phone_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_phone_proto_goTypes,
		DependencyIndexes: file_phone_proto_depIdxs,
		EnumInfos:         file_phone_proto_enumTypes,
		MessageInfos:      file_phone_proto_msgTypes,
	}.Build()
	File_phone_proto = out.File
	file_phone_proto_rawDesc = nil
	file_phone_proto_goTypes = nil
	file_phone_proto_depIdxs = nil
}