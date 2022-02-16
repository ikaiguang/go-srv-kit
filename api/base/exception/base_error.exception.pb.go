// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/base/exception/base_error.exception.proto

package exception

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

// ERROR .
type ERROR int32

const (
	// UNKNOWN 未知
	ERROR_UNKNOWN ERROR = 0
	// CONTINUE Continue
	ERROR_CONTINUE   ERROR = 100
	ERROR_PROCESSING ERROR = 102
	// OK OK
	ERROR_OK      ERROR = 200
	ERROR_CREATED ERROR = 201
	// MULTIPLE_CHOICES MultipleChoices
	ERROR_MULTIPLE_CHOICES ERROR = 300
	// BAD_REQUEST Bad Request
	ERROR_BAD_REQUEST ERROR = 400
	// UNAUTHORIZED Unauthorized
	ERROR_UNAUTHORIZED ERROR = 401
	// FORBIDDEN Forbidden
	ERROR_FORBIDDEN ERROR = 403
	// NOT_FOUND Not Found
	ERROR_NOT_FOUND ERROR = 404
	// METHOD_NOT_ALLOWED Method Not Allowed
	ERROR_METHOD_NOT_ALLOWED ERROR = 405
	// REQUEST_TIMEOUT Request Timeout
	ERROR_REQUEST_TIMEOUT ERROR = 408
	// TOO_MANY_REQUESTS Too Many Requests
	ERROR_TOO_MANY_REQUESTS ERROR = 429
	// INTERNAL_SERVER Internal Server Error
	ERROR_INTERNAL_SERVER ERROR = 500
	// NOT_IMPLEMENTED Not Implemented
	ERROR_NOT_IMPLEMENTED ERROR = 501
	// BAD_GATEWAY Bad Gateway
	ERROR_BAD_GATEWAY ERROR = 502
	// GATEWAY_TIMEOUT Gateway Timeout
	ERROR_GATEWAY_TIMEOUT ERROR = 504
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0:   "UNKNOWN",
		100: "CONTINUE",
		102: "PROCESSING",
		200: "OK",
		201: "CREATED",
		300: "MULTIPLE_CHOICES",
		400: "BAD_REQUEST",
		401: "UNAUTHORIZED",
		403: "FORBIDDEN",
		404: "NOT_FOUND",
		405: "METHOD_NOT_ALLOWED",
		408: "REQUEST_TIMEOUT",
		429: "TOO_MANY_REQUESTS",
		500: "INTERNAL_SERVER",
		501: "NOT_IMPLEMENTED",
		502: "BAD_GATEWAY",
		504: "GATEWAY_TIMEOUT",
	}
	ERROR_value = map[string]int32{
		"UNKNOWN":            0,
		"CONTINUE":           100,
		"PROCESSING":         102,
		"OK":                 200,
		"CREATED":            201,
		"MULTIPLE_CHOICES":   300,
		"BAD_REQUEST":        400,
		"UNAUTHORIZED":       401,
		"FORBIDDEN":          403,
		"NOT_FOUND":          404,
		"METHOD_NOT_ALLOWED": 405,
		"REQUEST_TIMEOUT":    408,
		"TOO_MANY_REQUESTS":  429,
		"INTERNAL_SERVER":    500,
		"NOT_IMPLEMENTED":    501,
		"BAD_GATEWAY":        502,
		"GATEWAY_TIMEOUT":    504,
	}
)

func (x ERROR) Enum() *ERROR {
	p := new(ERROR)
	*p = x
	return p
}

func (x ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_api_base_exception_base_error_exception_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_api_base_exception_base_error_exception_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_api_base_exception_base_error_exception_proto_rawDescGZIP(), []int{0}
}

var File_api_base_exception_base_error_exception_proto protoreflect.FileDescriptor

var file_api_base_exception_base_error_exception_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x65, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xaa, 0x03, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x12, 0x11, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a,
	0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x11, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55,
	0x45, 0x10, 0x64, 0x1a, 0x03, 0xa8, 0x45, 0x64, 0x12, 0x13, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x66, 0x1a, 0x03, 0xa8, 0x45, 0x66, 0x12, 0x0d, 0x0a,
	0x02, 0x4f, 0x4b, 0x10, 0xc8, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x12, 0x0a, 0x07,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0xc9, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xc9, 0x01,
	0x12, 0x1b, 0x0a, 0x10, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x43, 0x48, 0x4f,
	0x49, 0x43, 0x45, 0x53, 0x10, 0xac, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xac, 0x02, 0x12, 0x16, 0x0a,
	0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x90, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x17, 0x0a, 0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f,
	0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x91, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x14,
	0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x93, 0x03, 0x1a, 0x04,
	0xa8, 0x45, 0x93, 0x03, 0x12, 0x14, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x94, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1d, 0x0a, 0x12, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0x95, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x95, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x98, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x98, 0x03, 0x12, 0x1c, 0x0a, 0x11, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e,
	0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53, 0x10, 0xad, 0x03, 0x1a, 0x04, 0xa8,
	0x45, 0xad, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0xf4, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12,
	0x1a, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x45, 0x44, 0x10, 0xf5, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12, 0x16, 0x0a, 0x0b, 0x42,
	0x41, 0x44, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x10, 0xf6, 0x03, 0x1a, 0x04, 0xa8,
	0x45, 0xf6, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xf8, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf8, 0x03, 0x1a,
	0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x54, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75,
	0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3b, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_base_exception_base_error_exception_proto_rawDescOnce sync.Once
	file_api_base_exception_base_error_exception_proto_rawDescData = file_api_base_exception_base_error_exception_proto_rawDesc
)

func file_api_base_exception_base_error_exception_proto_rawDescGZIP() []byte {
	file_api_base_exception_base_error_exception_proto_rawDescOnce.Do(func() {
		file_api_base_exception_base_error_exception_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_base_exception_base_error_exception_proto_rawDescData)
	})
	return file_api_base_exception_base_error_exception_proto_rawDescData
}

var file_api_base_exception_base_error_exception_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_base_exception_base_error_exception_proto_goTypes = []interface{}{
	(ERROR)(0), // 0: api.base.exception.ERROR
}
var file_api_base_exception_base_error_exception_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_base_exception_base_error_exception_proto_init() }
func file_api_base_exception_base_error_exception_proto_init() {
	if File_api_base_exception_base_error_exception_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_base_exception_base_error_exception_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_base_exception_base_error_exception_proto_goTypes,
		DependencyIndexes: file_api_base_exception_base_error_exception_proto_depIdxs,
		EnumInfos:         file_api_base_exception_base_error_exception_proto_enumTypes,
	}.Build()
	File_api_base_exception_base_error_exception_proto = out.File
	file_api_base_exception_base_error_exception_proto_rawDesc = nil
	file_api_base_exception_base_error_exception_proto_goTypes = nil
	file_api_base_exception_base_error_exception_proto_depIdxs = nil
}
