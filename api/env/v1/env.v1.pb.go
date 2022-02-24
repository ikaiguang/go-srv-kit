// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/env/v1/env.v1.proto

package envv1

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

// Env app 环境
type Env int32

const (
	// UNKNOWN 未知
	Env_UNKNOWN Env = 0
	// DEVELOP 开发环境
	Env_DEVELOP Env = 1
	// TESTING 测试环境
	Env_TESTING Env = 2
	// PREVIEW 预发布 环境
	Env_PREVIEW Env = 3
	// PRODUCTION 生产环境
	Env_PRODUCTION Env = 4
)

// Enum value maps for Env.
var (
	Env_name = map[int32]string{
		0: "UNKNOWN",
		1: "DEVELOP",
		2: "TESTING",
		3: "PREVIEW",
		4: "PRODUCTION",
	}
	Env_value = map[string]int32{
		"UNKNOWN":    0,
		"DEVELOP":    1,
		"TESTING":    2,
		"PREVIEW":    3,
		"PRODUCTION": 4,
	}
)

func (x Env) Enum() *Env {
	p := new(Env)
	*p = x
	return p
}

func (x Env) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Env) Descriptor() protoreflect.EnumDescriptor {
	return file_api_env_v1_env_v1_proto_enumTypes[0].Descriptor()
}

func (Env) Type() protoreflect.EnumType {
	return &file_api_env_v1_env_v1_proto_enumTypes[0]
}

func (x Env) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Env.Descriptor instead.
func (Env) EnumDescriptor() ([]byte, []int) {
	return file_api_env_v1_env_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_env_v1_env_v1_proto protoreflect.FileDescriptor

var file_api_env_v1_env_v1_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x76, 0x2e, 0x65, 0x6e, 0x76, 0x76, 0x31, 0x2a, 0x49, 0x0a, 0x03, 0x45, 0x6e, 0x76, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45,
	0x57, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x04, 0x42, 0x43, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x2e, 0x65,
	0x6e, 0x76, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x2f,
	0x76, 0x31, 0x3b, 0x65, 0x6e, 0x76, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_env_v1_env_v1_proto_rawDescOnce sync.Once
	file_api_env_v1_env_v1_proto_rawDescData = file_api_env_v1_env_v1_proto_rawDesc
)

func file_api_env_v1_env_v1_proto_rawDescGZIP() []byte {
	file_api_env_v1_env_v1_proto_rawDescOnce.Do(func() {
		file_api_env_v1_env_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_env_v1_env_v1_proto_rawDescData)
	})
	return file_api_env_v1_env_v1_proto_rawDescData
}

var file_api_env_v1_env_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_env_v1_env_v1_proto_goTypes = []interface{}{
	(Env)(0), // 0: api.env.envv1.Env
}
var file_api_env_v1_env_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_env_v1_env_v1_proto_init() }
func file_api_env_v1_env_v1_proto_init() {
	if File_api_env_v1_env_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_env_v1_env_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_env_v1_env_v1_proto_goTypes,
		DependencyIndexes: file_api_env_v1_env_v1_proto_depIdxs,
		EnumInfos:         file_api_env_v1_env_v1_proto_enumTypes,
	}.Build()
	File_api_env_v1_env_v1_proto = out.File
	file_api_env_v1_env_v1_proto_rawDesc = nil
	file_api_env_v1_env_v1_proto_goTypes = nil
	file_api_env_v1_env_v1_proto_depIdxs = nil
}
