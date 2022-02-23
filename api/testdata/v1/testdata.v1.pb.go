// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/testdata/v1/testdata.v1.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// TestReq 请求
type TestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 请求消息
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestReq) Reset() {
	*x = TestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_testdata_v1_testdata_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestReq) ProtoMessage() {}

func (x *TestReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_testdata_v1_testdata_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestReq.ProtoReflect.Descriptor instead.
func (*TestReq) Descriptor() ([]byte, []int) {
	return file_api_testdata_v1_testdata_v1_proto_rawDescGZIP(), []int{0}
}

func (x *TestReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// TestResp 响应
type TestResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应消息
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestResp) Reset() {
	*x = TestResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_testdata_v1_testdata_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResp) ProtoMessage() {}

func (x *TestResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_testdata_v1_testdata_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResp.ProtoReflect.Descriptor instead.
func (*TestResp) Descriptor() ([]byte, []int) {
	return file_api_testdata_v1_testdata_v1_proto_rawDescGZIP(), []int{1}
}

func (x *TestResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_testdata_v1_testdata_v1_proto protoreflect.FileDescriptor

var file_api_testdata_v1_testdata_v1_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xae, 0x07,
	0x0a, 0x0b, 0x53, 0x72, 0x76, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x12, 0x64, 0x0a,
	0x09, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x58, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x5b, 0x0a,
	0x03, 0x50, 0x75, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x1a, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x70, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5d, 0x0a, 0x04, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22,
	0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x32, 0x15, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x29,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d,
	0x74, 0x6f, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x72, 0x0a, 0x0e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2d, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x28, 0x01, 0x12,
	0x7d, 0x0a, 0x13, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x28, 0x01, 0x30, 0x01, 0x42, 0x47,
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76,
	0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_testdata_v1_testdata_v1_proto_rawDescOnce sync.Once
	file_api_testdata_v1_testdata_v1_proto_rawDescData = file_api_testdata_v1_testdata_v1_proto_rawDesc
)

func file_api_testdata_v1_testdata_v1_proto_rawDescGZIP() []byte {
	file_api_testdata_v1_testdata_v1_proto_rawDescOnce.Do(func() {
		file_api_testdata_v1_testdata_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_testdata_v1_testdata_v1_proto_rawDescData)
	})
	return file_api_testdata_v1_testdata_v1_proto_rawDescData
}

var file_api_testdata_v1_testdata_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_testdata_v1_testdata_v1_proto_goTypes = []interface{}{
	(*TestReq)(nil),  // 0: api.testdata.v1.TestReq
	(*TestResp)(nil), // 1: api.testdata.v1.TestResp
}
var file_api_testdata_v1_testdata_v1_proto_depIdxs = []int32{
	0, // 0: api.testdata.v1.SrvTestdata.Websocket:input_type -> api.testdata.v1.TestReq
	0, // 1: api.testdata.v1.SrvTestdata.Get:input_type -> api.testdata.v1.TestReq
	0, // 2: api.testdata.v1.SrvTestdata.Put:input_type -> api.testdata.v1.TestReq
	0, // 3: api.testdata.v1.SrvTestdata.Post:input_type -> api.testdata.v1.TestReq
	0, // 4: api.testdata.v1.SrvTestdata.Delete:input_type -> api.testdata.v1.TestReq
	0, // 5: api.testdata.v1.SrvTestdata.Patch:input_type -> api.testdata.v1.TestReq
	0, // 6: api.testdata.v1.SrvTestdata.ServerToClient:input_type -> api.testdata.v1.TestReq
	0, // 7: api.testdata.v1.SrvTestdata.ClientToServer:input_type -> api.testdata.v1.TestReq
	0, // 8: api.testdata.v1.SrvTestdata.BidirectionalStream:input_type -> api.testdata.v1.TestReq
	1, // 9: api.testdata.v1.SrvTestdata.Websocket:output_type -> api.testdata.v1.TestResp
	1, // 10: api.testdata.v1.SrvTestdata.Get:output_type -> api.testdata.v1.TestResp
	1, // 11: api.testdata.v1.SrvTestdata.Put:output_type -> api.testdata.v1.TestResp
	1, // 12: api.testdata.v1.SrvTestdata.Post:output_type -> api.testdata.v1.TestResp
	1, // 13: api.testdata.v1.SrvTestdata.Delete:output_type -> api.testdata.v1.TestResp
	1, // 14: api.testdata.v1.SrvTestdata.Patch:output_type -> api.testdata.v1.TestResp
	1, // 15: api.testdata.v1.SrvTestdata.ServerToClient:output_type -> api.testdata.v1.TestResp
	1, // 16: api.testdata.v1.SrvTestdata.ClientToServer:output_type -> api.testdata.v1.TestResp
	1, // 17: api.testdata.v1.SrvTestdata.BidirectionalStream:output_type -> api.testdata.v1.TestResp
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_testdata_v1_testdata_v1_proto_init() }
func file_api_testdata_v1_testdata_v1_proto_init() {
	if File_api_testdata_v1_testdata_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_testdata_v1_testdata_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestReq); i {
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
		file_api_testdata_v1_testdata_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResp); i {
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
			RawDescriptor: file_api_testdata_v1_testdata_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_testdata_v1_testdata_v1_proto_goTypes,
		DependencyIndexes: file_api_testdata_v1_testdata_v1_proto_depIdxs,
		MessageInfos:      file_api_testdata_v1_testdata_v1_proto_msgTypes,
	}.Build()
	File_api_testdata_v1_testdata_v1_proto = out.File
	file_api_testdata_v1_testdata_v1_proto_rawDesc = nil
	file_api_testdata_v1_testdata_v1_proto_goTypes = nil
	file_api_testdata_v1_testdata_v1_proto_depIdxs = nil
}