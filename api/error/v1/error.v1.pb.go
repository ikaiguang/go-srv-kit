// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/error/v1/error.v1.proto

package errorv1

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
	// UNKNOWN 常规
	ERROR_UNKNOWN               ERROR = 0
	ERROR_REQUEST_FAILED        ERROR = 1
	ERROR_RECORD_NOT_FOUND      ERROR = 2
	ERROR_RECORD_ALREADY_EXISTS ERROR = 3
	// CONTINUE Continue
	ERROR_CONTINUE            ERROR = 100
	ERROR_SWITCHING_PROTOCOLS ERROR = 101
	ERROR_PROCESSING          ERROR = 102
	ERROR_EARLY_HINTS         ERROR = 103
	// OK OK
	ERROR_OK                     ERROR = 200
	ERROR_CREATED                ERROR = 201
	ERROR_ACCEPTED               ERROR = 202
	ERROR_NON_AUTHORITATIVE_INFO ERROR = 203
	ERROR_NO_CONTENT             ERROR = 204
	ERROR_RESET_CONTENT          ERROR = 205
	ERROR_PARTIAL_CONTENT        ERROR = 206
	ERROR_MULTI_STATUS           ERROR = 207
	ERROR_ALREADY_REPORTED       ERROR = 208
	ERROR_I_M_USED               ERROR = 226
	// MULTIPLE_CHOICES MultipleChoices
	ERROR_MULTIPLE_CHOICES   ERROR = 300
	ERROR_MOVED_PERMANENTLY  ERROR = 301
	ERROR_FOUND              ERROR = 302
	ERROR_SEE_OTHER          ERROR = 303
	ERROR_NOT_MODIFIED       ERROR = 304
	ERROR_USE_PROXY          ERROR = 305
	ERROR_EMPTY306           ERROR = 306
	ERROR_TEMPORARY_REDIRECT ERROR = 307
	ERROR_PERMANENT_REDIRECT ERROR = 308
	// BAD_REQUEST Bad Request
	ERROR_BAD_REQUEST                     ERROR = 400
	ERROR_UNAUTHORIZED                    ERROR = 401
	ERROR_PAYMENT_REQUIRED                ERROR = 402
	ERROR_FORBIDDEN                       ERROR = 403
	ERROR_NOT_FOUND                       ERROR = 404
	ERROR_METHOD_NOT_ALLOWED              ERROR = 405
	ERROR_NOT_ACCEPTABLE                  ERROR = 406
	ERROR_PROXY_AUTH_REQUIRED             ERROR = 407
	ERROR_REQUEST_TIMEOUT                 ERROR = 408
	ERROR_CONFLICT                        ERROR = 409
	ERROR_GONE                            ERROR = 410
	ERROR_LENGTH_REQUIRED                 ERROR = 411
	ERROR_PRECONDITION_FAILED             ERROR = 412
	ERROR_REQUEST_ENTITY_TOO_LARGE        ERROR = 413
	ERROR_REQUEST_URI_TOO_LONG            ERROR = 414
	ERROR_UNSUPPORTED_MEDIA_TYPE          ERROR = 415
	ERROR_REQUESTED_RANGE_NOT_SATISFIABLE ERROR = 416
	ERROR_EXPECTATION_FAILED              ERROR = 417
	ERROR_TEAPOT                          ERROR = 418
	ERROR_MISDIRECTED_REQUEST             ERROR = 421
	ERROR_UNPROCESSABLE_ENTITY            ERROR = 422
	ERROR_LOCKED                          ERROR = 423
	ERROR_FAILED_DEPENDENCY               ERROR = 424
	ERROR_TOO_EARLY                       ERROR = 425
	ERROR_UPGRADE_REQUIRED                ERROR = 426
	ERROR_PRECONDITION_REQUIRED           ERROR = 428
	ERROR_TOO_MANY_REQUESTS               ERROR = 429
	ERROR_REQUEST_HEADER_FIELDS_TOO_LARGE ERROR = 431
	ERROR_UNAVAILABLE_FOR_LEGAL_REASONS   ERROR = 451
	// INTERNAL_SERVER Internal Server Error
	ERROR_INTERNAL_SERVER                 ERROR = 500
	ERROR_NOT_IMPLEMENTED                 ERROR = 501
	ERROR_BAD_GATEWAY                     ERROR = 502
	ERROR_SERVICE_UNAVAILABLE             ERROR = 503
	ERROR_GATEWAY_TIMEOUT                 ERROR = 504
	ERROR_HTTP_VERSION_NOT_SUPPORTED      ERROR = 505
	ERROR_VARIANT_ALSO_NEGOTIATES         ERROR = 506
	ERROR_INSUFFICIENT_STORAGE            ERROR = 507
	ERROR_LOOP_DETECTED                   ERROR = 508
	ERROR_NOT_EXTENDED                    ERROR = 510
	ERROR_NETWORK_AUTHENTICATION_REQUIRED ERROR = 511
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0:   "UNKNOWN",
		1:   "REQUEST_FAILED",
		2:   "RECORD_NOT_FOUND",
		3:   "RECORD_ALREADY_EXISTS",
		100: "CONTINUE",
		101: "SWITCHING_PROTOCOLS",
		102: "PROCESSING",
		103: "EARLY_HINTS",
		200: "OK",
		201: "CREATED",
		202: "ACCEPTED",
		203: "NON_AUTHORITATIVE_INFO",
		204: "NO_CONTENT",
		205: "RESET_CONTENT",
		206: "PARTIAL_CONTENT",
		207: "MULTI_STATUS",
		208: "ALREADY_REPORTED",
		226: "I_M_USED",
		300: "MULTIPLE_CHOICES",
		301: "MOVED_PERMANENTLY",
		302: "FOUND",
		303: "SEE_OTHER",
		304: "NOT_MODIFIED",
		305: "USE_PROXY",
		306: "EMPTY306",
		307: "TEMPORARY_REDIRECT",
		308: "PERMANENT_REDIRECT",
		400: "BAD_REQUEST",
		401: "UNAUTHORIZED",
		402: "PAYMENT_REQUIRED",
		403: "FORBIDDEN",
		404: "NOT_FOUND",
		405: "METHOD_NOT_ALLOWED",
		406: "NOT_ACCEPTABLE",
		407: "PROXY_AUTH_REQUIRED",
		408: "REQUEST_TIMEOUT",
		409: "CONFLICT",
		410: "GONE",
		411: "LENGTH_REQUIRED",
		412: "PRECONDITION_FAILED",
		413: "REQUEST_ENTITY_TOO_LARGE",
		414: "REQUEST_URI_TOO_LONG",
		415: "UNSUPPORTED_MEDIA_TYPE",
		416: "REQUESTED_RANGE_NOT_SATISFIABLE",
		417: "EXPECTATION_FAILED",
		418: "TEAPOT",
		421: "MISDIRECTED_REQUEST",
		422: "UNPROCESSABLE_ENTITY",
		423: "LOCKED",
		424: "FAILED_DEPENDENCY",
		425: "TOO_EARLY",
		426: "UPGRADE_REQUIRED",
		428: "PRECONDITION_REQUIRED",
		429: "TOO_MANY_REQUESTS",
		431: "REQUEST_HEADER_FIELDS_TOO_LARGE",
		451: "UNAVAILABLE_FOR_LEGAL_REASONS",
		500: "INTERNAL_SERVER",
		501: "NOT_IMPLEMENTED",
		502: "BAD_GATEWAY",
		503: "SERVICE_UNAVAILABLE",
		504: "GATEWAY_TIMEOUT",
		505: "HTTP_VERSION_NOT_SUPPORTED",
		506: "VARIANT_ALSO_NEGOTIATES",
		507: "INSUFFICIENT_STORAGE",
		508: "LOOP_DETECTED",
		510: "NOT_EXTENDED",
		511: "NETWORK_AUTHENTICATION_REQUIRED",
	}
	ERROR_value = map[string]int32{
		"UNKNOWN":                         0,
		"REQUEST_FAILED":                  1,
		"RECORD_NOT_FOUND":                2,
		"RECORD_ALREADY_EXISTS":           3,
		"CONTINUE":                        100,
		"SWITCHING_PROTOCOLS":             101,
		"PROCESSING":                      102,
		"EARLY_HINTS":                     103,
		"OK":                              200,
		"CREATED":                         201,
		"ACCEPTED":                        202,
		"NON_AUTHORITATIVE_INFO":          203,
		"NO_CONTENT":                      204,
		"RESET_CONTENT":                   205,
		"PARTIAL_CONTENT":                 206,
		"MULTI_STATUS":                    207,
		"ALREADY_REPORTED":                208,
		"I_M_USED":                        226,
		"MULTIPLE_CHOICES":                300,
		"MOVED_PERMANENTLY":               301,
		"FOUND":                           302,
		"SEE_OTHER":                       303,
		"NOT_MODIFIED":                    304,
		"USE_PROXY":                       305,
		"EMPTY306":                        306,
		"TEMPORARY_REDIRECT":              307,
		"PERMANENT_REDIRECT":              308,
		"BAD_REQUEST":                     400,
		"UNAUTHORIZED":                    401,
		"PAYMENT_REQUIRED":                402,
		"FORBIDDEN":                       403,
		"NOT_FOUND":                       404,
		"METHOD_NOT_ALLOWED":              405,
		"NOT_ACCEPTABLE":                  406,
		"PROXY_AUTH_REQUIRED":             407,
		"REQUEST_TIMEOUT":                 408,
		"CONFLICT":                        409,
		"GONE":                            410,
		"LENGTH_REQUIRED":                 411,
		"PRECONDITION_FAILED":             412,
		"REQUEST_ENTITY_TOO_LARGE":        413,
		"REQUEST_URI_TOO_LONG":            414,
		"UNSUPPORTED_MEDIA_TYPE":          415,
		"REQUESTED_RANGE_NOT_SATISFIABLE": 416,
		"EXPECTATION_FAILED":              417,
		"TEAPOT":                          418,
		"MISDIRECTED_REQUEST":             421,
		"UNPROCESSABLE_ENTITY":            422,
		"LOCKED":                          423,
		"FAILED_DEPENDENCY":               424,
		"TOO_EARLY":                       425,
		"UPGRADE_REQUIRED":                426,
		"PRECONDITION_REQUIRED":           428,
		"TOO_MANY_REQUESTS":               429,
		"REQUEST_HEADER_FIELDS_TOO_LARGE": 431,
		"UNAVAILABLE_FOR_LEGAL_REASONS":   451,
		"INTERNAL_SERVER":                 500,
		"NOT_IMPLEMENTED":                 501,
		"BAD_GATEWAY":                     502,
		"SERVICE_UNAVAILABLE":             503,
		"GATEWAY_TIMEOUT":                 504,
		"HTTP_VERSION_NOT_SUPPORTED":      505,
		"VARIANT_ALSO_NEGOTIATES":         506,
		"INSUFFICIENT_STORAGE":            507,
		"LOOP_DETECTED":                   508,
		"NOT_EXTENDED":                    510,
		"NETWORK_AUTHENTICATION_REQUIRED": 511,
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
	return file_api_error_v1_error_v1_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_api_error_v1_error_v1_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_api_error_v1_error_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_error_v1_error_v1_proto protoreflect.FileDescriptor

var file_api_error_v1_error_v1_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6b,
	0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd6, 0x0e, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x12, 0x11, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x18, 0x0a, 0x0e, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xa8, 0x03,
	0x12, 0x1a, 0x0a, 0x10, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1f, 0x0a, 0x15,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x12, 0x11, 0x0a,
	0x08, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x10, 0x64, 0x1a, 0x03, 0xa8, 0x45, 0x64,
	0x12, 0x1c, 0x0a, 0x13, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x53, 0x10, 0x65, 0x1a, 0x03, 0xa8, 0x45, 0x65, 0x12, 0x13,
	0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x66, 0x1a, 0x03,
	0xa8, 0x45, 0x66, 0x12, 0x14, 0x0a, 0x0b, 0x45, 0x41, 0x52, 0x4c, 0x59, 0x5f, 0x48, 0x49, 0x4e,
	0x54, 0x53, 0x10, 0x67, 0x1a, 0x03, 0xa8, 0x45, 0x67, 0x12, 0x0d, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0xc8, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x12, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0xc9, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xc9, 0x01, 0x12, 0x13, 0x0a, 0x08,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0xca, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xca,
	0x01, 0x12, 0x21, 0x0a, 0x16, 0x4e, 0x4f, 0x4e, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49,
	0x54, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0xcb, 0x01, 0x1a, 0x04,
	0xa8, 0x45, 0xcb, 0x01, 0x12, 0x15, 0x0a, 0x0a, 0x4e, 0x4f, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45,
	0x4e, 0x54, 0x10, 0xcc, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xcc, 0x01, 0x12, 0x18, 0x0a, 0x0d, 0x52,
	0x45, 0x53, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0xcd, 0x01, 0x1a,
	0x04, 0xa8, 0x45, 0xcd, 0x01, 0x12, 0x1a, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0xce, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xce,
	0x01, 0x12, 0x17, 0x0a, 0x0c, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x10, 0xcf, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xcf, 0x01, 0x12, 0x1b, 0x0a, 0x10, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0xd0,
	0x01, 0x1a, 0x04, 0xa8, 0x45, 0xd0, 0x01, 0x12, 0x13, 0x0a, 0x08, 0x49, 0x5f, 0x4d, 0x5f, 0x55,
	0x53, 0x45, 0x44, 0x10, 0xe2, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xe2, 0x01, 0x12, 0x1b, 0x0a, 0x10,
	0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x53,
	0x10, 0xac, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xac, 0x02, 0x12, 0x1c, 0x0a, 0x11, 0x4d, 0x4f, 0x56,
	0x45, 0x44, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x45, 0x4e, 0x54, 0x4c, 0x59, 0x10, 0xad,
	0x02, 0x1a, 0x04, 0xa8, 0x45, 0xad, 0x02, 0x12, 0x10, 0x0a, 0x05, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0xae, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xae, 0x02, 0x12, 0x14, 0x0a, 0x09, 0x53, 0x45, 0x45,
	0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0xaf, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xaf, 0x02, 0x12,
	0x17, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0xb0, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xb0, 0x02, 0x12, 0x14, 0x0a, 0x09, 0x55, 0x53, 0x45, 0x5f,
	0x50, 0x52, 0x4f, 0x58, 0x59, 0x10, 0xb1, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xb1, 0x02, 0x12, 0x13,
	0x0a, 0x08, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x33, 0x30, 0x36, 0x10, 0xb2, 0x02, 0x1a, 0x04, 0xa8,
	0x45, 0xb2, 0x02, 0x12, 0x1d, 0x0a, 0x12, 0x54, 0x45, 0x4d, 0x50, 0x4f, 0x52, 0x41, 0x52, 0x59,
	0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0xb3, 0x02, 0x1a, 0x04, 0xa8, 0x45,
	0xb3, 0x02, 0x12, 0x1d, 0x0a, 0x12, 0x50, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x45, 0x4e, 0x54, 0x5f,
	0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0xb4, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xb4,
	0x02, 0x12, 0x16, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x90, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x17, 0x0a, 0x0c, 0x55, 0x4e, 0x41,
	0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x91, 0x03, 0x1a, 0x04, 0xa8, 0x45,
	0x91, 0x03, 0x12, 0x1b, 0x0a, 0x10, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x92, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x92, 0x03, 0x12,
	0x14, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x93, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x14, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x94, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1d, 0x0a, 0x12, 0x4d,
	0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45,
	0x44, 0x10, 0x95, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x95, 0x03, 0x12, 0x19, 0x0a, 0x0e, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x96, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x96, 0x03, 0x12, 0x1e, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x41,
	0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x97, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x97, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x98, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x98,
	0x03, 0x12, 0x13, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x10, 0x99, 0x03,
	0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x12, 0x0f, 0x0a, 0x04, 0x47, 0x4f, 0x4e, 0x45, 0x10, 0x9a,
	0x03, 0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x4c, 0x45, 0x4e, 0x47, 0x54,
	0x48, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x9b, 0x03, 0x1a, 0x04, 0xa8,
	0x45, 0x9b, 0x03, 0x12, 0x1e, 0x0a, 0x13, 0x50, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x9c, 0x03, 0x1a, 0x04, 0xa8,
	0x45, 0x9c, 0x03, 0x12, 0x23, 0x0a, 0x18, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10,
	0x9d, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x9d, 0x03, 0x12, 0x1f, 0x0a, 0x14, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x55, 0x52, 0x49, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47,
	0x10, 0x9e, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x9e, 0x03, 0x12, 0x21, 0x0a, 0x16, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x9f, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x9f, 0x03, 0x12, 0x2a, 0x0a, 0x1f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x41, 0x54, 0x49, 0x53, 0x46, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0xa0, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xa0, 0x03, 0x12, 0x1d, 0x0a, 0x12, 0x45, 0x58, 0x50, 0x45,
	0x43, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa1,
	0x03, 0x1a, 0x04, 0xa8, 0x45, 0xa1, 0x03, 0x12, 0x11, 0x0a, 0x06, 0x54, 0x45, 0x41, 0x50, 0x4f,
	0x54, 0x10, 0xa2, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xa2, 0x03, 0x12, 0x1e, 0x0a, 0x13, 0x4d, 0x49,
	0x53, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0xa5, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xa5, 0x03, 0x12, 0x1f, 0x0a, 0x14, 0x55, 0x4e,
	0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x10, 0xa6, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xa6, 0x03, 0x12, 0x11, 0x0a, 0x06, 0x4c,
	0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0xa7, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xa7, 0x03, 0x12, 0x1c,
	0x0a, 0x11, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x50, 0x45, 0x4e, 0x44, 0x45,
	0x4e, 0x43, 0x59, 0x10, 0xa8, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xa8, 0x03, 0x12, 0x14, 0x0a, 0x09,
	0x54, 0x4f, 0x4f, 0x5f, 0x45, 0x41, 0x52, 0x4c, 0x59, 0x10, 0xa9, 0x03, 0x1a, 0x04, 0xa8, 0x45,
	0xa9, 0x03, 0x12, 0x1b, 0x0a, 0x10, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0xaa, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xaa, 0x03, 0x12,
	0x20, 0x0a, 0x15, 0x50, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0xac, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xac,
	0x03, 0x12, 0x1c, 0x0a, 0x11, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x53, 0x10, 0xad, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xad, 0x03, 0x12,
	0x2a, 0x0a, 0x1f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45,
	0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52,
	0x47, 0x45, 0x10, 0xaf, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xaf, 0x03, 0x12, 0x28, 0x0a, 0x1d, 0x55,
	0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4c,
	0x45, 0x47, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x53, 0x10, 0xc3, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0xc3, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0xf4, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4,
	0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x45, 0x44, 0x10, 0xf5, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12, 0x16, 0x0a,
	0x0b, 0x42, 0x41, 0x44, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x10, 0xf6, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0xf6, 0x03, 0x12, 0x1e, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0xf7, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0xf7, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xf8, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf8,
	0x03, 0x12, 0x25, 0x0a, 0x1a, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10,
	0xf9, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf9, 0x03, 0x12, 0x22, 0x0a, 0x17, 0x56, 0x41, 0x52, 0x49,
	0x41, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x53, 0x4f, 0x5f, 0x4e, 0x45, 0x47, 0x4f, 0x54, 0x49, 0x41,
	0x54, 0x45, 0x53, 0x10, 0xfa, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xfa, 0x03, 0x12, 0x1f, 0x0a, 0x14,
	0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x4f,
	0x52, 0x41, 0x47, 0x45, 0x10, 0xfb, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xfb, 0x03, 0x12, 0x18, 0x0a,
	0x0d, 0x4c, 0x4f, 0x4f, 0x50, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0xfc,
	0x03, 0x1a, 0x04, 0xa8, 0x45, 0xfc, 0x03, 0x12, 0x17, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f, 0x45,
	0x58, 0x54, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0xfe, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xfe, 0x03,
	0x12, 0x2a, 0x0a, 0x1f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x41, 0x55, 0x54, 0x48,
	0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49,
	0x52, 0x45, 0x44, 0x10, 0xff, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xff, 0x03, 0x1a, 0x04, 0xa0, 0x45,
	0xf4, 0x03, 0x42, 0x65, 0x0a, 0x15, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x42, 0x12, 0x4b, 0x69, 0x74,
	0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x31, 0x50,
	0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b,
	0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b,
	0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_error_v1_error_v1_proto_rawDescOnce sync.Once
	file_api_error_v1_error_v1_proto_rawDescData = file_api_error_v1_error_v1_proto_rawDesc
)

func file_api_error_v1_error_v1_proto_rawDescGZIP() []byte {
	file_api_error_v1_error_v1_proto_rawDescOnce.Do(func() {
		file_api_error_v1_error_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_error_v1_error_v1_proto_rawDescData)
	})
	return file_api_error_v1_error_v1_proto_rawDescData
}

var file_api_error_v1_error_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_error_v1_error_v1_proto_goTypes = []interface{}{
	(ERROR)(0), // 0: kit.api.error.errorv1.ERROR
}
var file_api_error_v1_error_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_error_v1_error_v1_proto_init() }
func file_api_error_v1_error_v1_proto_init() {
	if File_api_error_v1_error_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_error_v1_error_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_error_v1_error_v1_proto_goTypes,
		DependencyIndexes: file_api_error_v1_error_v1_proto_depIdxs,
		EnumInfos:         file_api_error_v1_error_v1_proto_enumTypes,
	}.Build()
	File_api_error_v1_error_v1_proto = out.File
	file_api_error_v1_error_v1_proto_rawDesc = nil
	file_api_error_v1_error_v1_proto_goTypes = nil
	file_api_error_v1_error_v1_proto_depIdxs = nil
}