// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: kratos/app/app_response.kit.proto

package apppkg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RuntimeEnvEnum_RuntimeEnv int32

const (
	// UNKNOWN 未知
	RuntimeEnvEnum_UNKNOWN RuntimeEnvEnum_RuntimeEnv = 0
	// LOCAL 本地开发
	RuntimeEnvEnum_LOCAL RuntimeEnvEnum_RuntimeEnv = 1
	// DEVELOP 开发环境
	RuntimeEnvEnum_DEVELOP RuntimeEnvEnum_RuntimeEnv = 2
	// TESTING 测试环境
	RuntimeEnvEnum_TESTING RuntimeEnvEnum_RuntimeEnv = 3
	// PREVIEW 预发布 环境
	RuntimeEnvEnum_PREVIEW RuntimeEnvEnum_RuntimeEnv = 4
	// PRODUCTION 生产环境
	RuntimeEnvEnum_PRODUCTION RuntimeEnvEnum_RuntimeEnv = 5
)

// Enum value maps for RuntimeEnvEnum_RuntimeEnv.
var (
	RuntimeEnvEnum_RuntimeEnv_name = map[int32]string{
		0: "UNKNOWN",
		1: "LOCAL",
		2: "DEVELOP",
		3: "TESTING",
		4: "PREVIEW",
		5: "PRODUCTION",
	}
	RuntimeEnvEnum_RuntimeEnv_value = map[string]int32{
		"UNKNOWN":    0,
		"LOCAL":      1,
		"DEVELOP":    2,
		"TESTING":    3,
		"PREVIEW":    4,
		"PRODUCTION": 5,
	}
)

func (x RuntimeEnvEnum_RuntimeEnv) Enum() *RuntimeEnvEnum_RuntimeEnv {
	p := new(RuntimeEnvEnum_RuntimeEnv)
	*p = x
	return p
}

func (x RuntimeEnvEnum_RuntimeEnv) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RuntimeEnvEnum_RuntimeEnv) Descriptor() protoreflect.EnumDescriptor {
	return file_kratos_app_app_response_kit_proto_enumTypes[0].Descriptor()
}

func (RuntimeEnvEnum_RuntimeEnv) Type() protoreflect.EnumType {
	return &file_kratos_app_app_response_kit_proto_enumTypes[0]
}

func (x RuntimeEnvEnum_RuntimeEnv) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RuntimeEnvEnum_RuntimeEnv.Descriptor instead.
func (RuntimeEnvEnum_RuntimeEnv) EnumDescriptor() ([]byte, []int) {
	return file_kratos_app_app_response_kit_proto_rawDescGZIP(), []int{0, 0}
}

// RuntimeEnvEnum app运行环境
type RuntimeEnvEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RuntimeEnvEnum) Reset() {
	*x = RuntimeEnvEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kratos_app_app_response_kit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeEnvEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeEnvEnum) ProtoMessage() {}

func (x *RuntimeEnvEnum) ProtoReflect() protoreflect.Message {
	mi := &file_kratos_app_app_response_kit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeEnvEnum.ProtoReflect.Descriptor instead.
func (*RuntimeEnvEnum) Descriptor() ([]byte, []int) {
	return file_kratos_app_app_response_kit_proto_rawDescGZIP(), []int{0}
}

// Response 响应
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason    string            `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message   string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	RequestId string            `protobuf:"bytes,4,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Data      *anypb.Any        `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kratos_app_app_response_kit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_kratos_app_app_response_kit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_kratos_app_app_response_kit_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Response) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ResponseData data
type ResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseData) Reset() {
	*x = ResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kratos_app_app_response_kit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseData) ProtoMessage() {}

func (x *ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_kratos_app_app_response_kit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseData.ProtoReflect.Descriptor instead.
func (*ResponseData) Descriptor() ([]byte, []int) {
	return file_kratos_app_app_response_kit_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_kratos_app_app_response_kit_proto protoreflect.FileDescriptor

var file_kratos_app_app_response_kit_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x61, 0x70, 0x70,
	0x70, 0x6b, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d,
	0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0x5b, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x76, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c,
	0x4f, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f,
	0x50, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x04, 0x12, 0x0e, 0x0a,
	0x0a, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x22, 0x99, 0x02,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x69, 0x74,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x70, 0x6b, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x53, 0x0a,
	0x0e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x70, 0x6b, 0x67, 0x42,
	0x0c, 0x4b, 0x69, 0x74, 0x41, 0x70, 0x70, 0x41, 0x70, 0x70, 0x50, 0x6b, 0x67, 0x50, 0x01, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69,
	0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74,
	0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x3b, 0x61, 0x70, 0x70, 0x70,
	0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kratos_app_app_response_kit_proto_rawDescOnce sync.Once
	file_kratos_app_app_response_kit_proto_rawDescData = file_kratos_app_app_response_kit_proto_rawDesc
)

func file_kratos_app_app_response_kit_proto_rawDescGZIP() []byte {
	file_kratos_app_app_response_kit_proto_rawDescOnce.Do(func() {
		file_kratos_app_app_response_kit_proto_rawDescData = protoimpl.X.CompressGZIP(file_kratos_app_app_response_kit_proto_rawDescData)
	})
	return file_kratos_app_app_response_kit_proto_rawDescData
}

var file_kratos_app_app_response_kit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kratos_app_app_response_kit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kratos_app_app_response_kit_proto_goTypes = []interface{}{
	(RuntimeEnvEnum_RuntimeEnv)(0), // 0: kit.app.apppkg.RuntimeEnvEnum.RuntimeEnv
	(*RuntimeEnvEnum)(nil),         // 1: kit.app.apppkg.RuntimeEnvEnum
	(*Response)(nil),               // 2: kit.app.apppkg.Response
	(*ResponseData)(nil),           // 3: kit.app.apppkg.ResponseData
	nil,                            // 4: kit.app.apppkg.Response.MetadataEntry
	(*anypb.Any)(nil),              // 5: google.protobuf.Any
}
var file_kratos_app_app_response_kit_proto_depIdxs = []int32{
	5, // 0: kit.app.apppkg.Response.data:type_name -> google.protobuf.Any
	4, // 1: kit.app.apppkg.Response.metadata:type_name -> kit.app.apppkg.Response.MetadataEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kratos_app_app_response_kit_proto_init() }
func file_kratos_app_app_response_kit_proto_init() {
	if File_kratos_app_app_response_kit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kratos_app_app_response_kit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeEnvEnum); i {
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
		file_kratos_app_app_response_kit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_kratos_app_app_response_kit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseData); i {
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
			RawDescriptor: file_kratos_app_app_response_kit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kratos_app_app_response_kit_proto_goTypes,
		DependencyIndexes: file_kratos_app_app_response_kit_proto_depIdxs,
		EnumInfos:         file_kratos_app_app_response_kit_proto_enumTypes,
		MessageInfos:      file_kratos_app_app_response_kit_proto_msgTypes,
	}.Build()
	File_kratos_app_app_response_kit_proto = out.File
	file_kratos_app_app_response_kit_proto_rawDesc = nil
	file_kratos_app_app_response_kit_proto_goTypes = nil
	file_kratos_app_app_response_kit_proto_depIdxs = nil
}
