// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: api/auth/v1/auth.v1.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status 枚举值
type StatusEnum_Status int32

const (
	// UNSPECIFIED 未指定
	StatusEnum_UNSPECIFIED StatusEnum_Status = 0
	// INITIAL 初始状态
	StatusEnum_INITIAL StatusEnum_Status = 1
	// ENABLE 启用
	StatusEnum_ENABLE StatusEnum_Status = 2
	// DISABLE 禁用
	StatusEnum_DISABLE StatusEnum_Status = 3
	// WHITELIST 白名单
	StatusEnum_WHITELIST StatusEnum_Status = 4
	// BLACKLIST 黑名单
	StatusEnum_BLACKLIST StatusEnum_Status = 5
	// DELETED 已删除
	StatusEnum_DELETED StatusEnum_Status = 6
)

// Enum value maps for StatusEnum_Status.
var (
	StatusEnum_Status_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "INITIAL",
		2: "ENABLE",
		3: "DISABLE",
		4: "WHITELIST",
		5: "BLACKLIST",
		6: "DELETED",
	}
	StatusEnum_Status_value = map[string]int32{
		"UNSPECIFIED": 0,
		"INITIAL":     1,
		"ENABLE":      2,
		"DISABLE":     3,
		"WHITELIST":   4,
		"BLACKLIST":   5,
		"DELETED":     6,
	}
)

func (x StatusEnum_Status) Enum() *StatusEnum_Status {
	p := new(StatusEnum_Status)
	*p = x
	return p
}

func (x StatusEnum_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusEnum_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_auth_v1_auth_v1_proto_enumTypes[0].Descriptor()
}

func (StatusEnum_Status) Type() protoreflect.EnumType {
	return &file_api_auth_v1_auth_v1_proto_enumTypes[0]
}

func (x StatusEnum_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusEnum_Status.Descriptor instead.
func (StatusEnum_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{2, 0}
}

type PlatformEnum_Platform int32

const (
	// UNKNOWN 未知
	PlatformEnum_UNKNOWN PlatformEnum_Platform = 0
	// COMPUTER 电脑端
	PlatformEnum_COMPUTER PlatformEnum_Platform = 1
	// MOBILE 移动端
	PlatformEnum_MOBILE PlatformEnum_Platform = 2
	// ANDROID 安卓系统
	PlatformEnum_ANDROID PlatformEnum_Platform = 100
	// IOS 苹果系统
	PlatformEnum_IOS PlatformEnum_Platform = 101
	// WEB 网页
	PlatformEnum_WEB PlatformEnum_Platform = 200
	// H5 网页
	PlatformEnum_H5 PlatformEnum_Platform = 201
	// APPLET 小程序
	PlatformEnum_APPLET PlatformEnum_Platform = 203
	// APP 应用
	PlatformEnum_APP PlatformEnum_Platform = 204
	// ANDROID_APP 安卓应用
	PlatformEnum_ANDROID_APP PlatformEnum_Platform = 205
	// IOS_APP 苹果应用
	PlatformEnum_IOS_APP PlatformEnum_Platform = 206
	// IPAD_APP 平板应用
	PlatformEnum_IPAD_APP PlatformEnum_Platform = 207
)

// Enum value maps for PlatformEnum_Platform.
var (
	PlatformEnum_Platform_name = map[int32]string{
		0:   "UNKNOWN",
		1:   "COMPUTER",
		2:   "MOBILE",
		100: "ANDROID",
		101: "IOS",
		200: "WEB",
		201: "H5",
		203: "APPLET",
		204: "APP",
		205: "ANDROID_APP",
		206: "IOS_APP",
		207: "IPAD_APP",
	}
	PlatformEnum_Platform_value = map[string]int32{
		"UNKNOWN":     0,
		"COMPUTER":    1,
		"MOBILE":      2,
		"ANDROID":     100,
		"IOS":         101,
		"WEB":         200,
		"H5":          201,
		"APPLET":      203,
		"APP":         204,
		"ANDROID_APP": 205,
		"IOS_APP":     206,
		"IPAD_APP":    207,
	}
)

func (x PlatformEnum_Platform) Enum() *PlatformEnum_Platform {
	p := new(PlatformEnum_Platform)
	*p = x
	return p
}

func (x PlatformEnum_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatformEnum_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_api_auth_v1_auth_v1_proto_enumTypes[1].Descriptor()
}

func (PlatformEnum_Platform) Type() protoreflect.EnumType {
	return &file_api_auth_v1_auth_v1_proto_enumTypes[1]
}

func (x PlatformEnum_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatformEnum_Platform.Descriptor instead.
func (PlatformEnum_Platform) EnumDescriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{3, 0}
}

// LimitType 登录限制类型
type LimitTypeEnum_LimitType int32

const (
	// UNLIMITED 无限制
	LimitTypeEnum_UNLIMITED LimitTypeEnum_LimitType = 0
	// ONLY_ONE 同一账户仅允许登录一次
	// 方案1：验证码...可强制登录
	// 方案2：强制登录，然后提示"强制下线/在其他终端登录"
	LimitTypeEnum_ONLY_ONE LimitTypeEnum_LimitType = 1
	// PLATFORM_ONE 同一账户每个平台都可登录一次
	// 方案1：验证码...可强制登录
	// 方案2：强制登录，然后提示"强制下线/在其他终端登录"
	LimitTypeEnum_PLATFORM_ONE LimitTypeEnum_LimitType = 2
)

// Enum value maps for LimitTypeEnum_LimitType.
var (
	LimitTypeEnum_LimitType_name = map[int32]string{
		0: "UNLIMITED",
		1: "ONLY_ONE",
		2: "PLATFORM_ONE",
	}
	LimitTypeEnum_LimitType_value = map[string]int32{
		"UNLIMITED":    0,
		"ONLY_ONE":     1,
		"PLATFORM_ONE": 2,
	}
)

func (x LimitTypeEnum_LimitType) Enum() *LimitTypeEnum_LimitType {
	p := new(LimitTypeEnum_LimitType)
	*p = x
	return p
}

func (x LimitTypeEnum_LimitType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LimitTypeEnum_LimitType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_auth_v1_auth_v1_proto_enumTypes[2].Descriptor()
}

func (LimitTypeEnum_LimitType) Type() protoreflect.EnumType {
	return &file_api_auth_v1_auth_v1_proto_enumTypes[2]
}

func (x LimitTypeEnum_LimitType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LimitTypeEnum_LimitType.Descriptor instead.
func (LimitTypeEnum_LimitType) EnumDescriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{4, 0}
}

// TokenType 令牌类型
type TokenTypeEnum_TokenType int32

const (
	// DEFAULT 默认
	TokenTypeEnum_DEFAULT TokenTypeEnum_TokenType = 0
	// SERVICE 服务直连
	TokenTypeEnum_SERVICE TokenTypeEnum_TokenType = 1
	// ADMIN 管理后台
	TokenTypeEnum_ADMIN TokenTypeEnum_TokenType = 2
	// API 通用api
	TokenTypeEnum_API TokenTypeEnum_TokenType = 3
	// WEB 通用web
	TokenTypeEnum_WEB TokenTypeEnum_TokenType = 4
	// APP 应用
	TokenTypeEnum_APP TokenTypeEnum_TokenType = 5
	// H5 h5应用
	TokenTypeEnum_H5 TokenTypeEnum_TokenType = 6
	// MANAGER 管理员
	TokenTypeEnum_MANAGER TokenTypeEnum_TokenType = 7
)

// Enum value maps for TokenTypeEnum_TokenType.
var (
	TokenTypeEnum_TokenType_name = map[int32]string{
		0: "DEFAULT",
		1: "SERVICE",
		2: "ADMIN",
		3: "API",
		4: "WEB",
		5: "APP",
		6: "H5",
		7: "MANAGER",
	}
	TokenTypeEnum_TokenType_value = map[string]int32{
		"DEFAULT": 0,
		"SERVICE": 1,
		"ADMIN":   2,
		"API":     3,
		"WEB":     4,
		"APP":     5,
		"H5":      6,
		"MANAGER": 7,
	}
)

func (x TokenTypeEnum_TokenType) Enum() *TokenTypeEnum_TokenType {
	p := new(TokenTypeEnum_TokenType)
	*p = x
	return p
}

func (x TokenTypeEnum_TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenTypeEnum_TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_auth_v1_auth_v1_proto_enumTypes[3].Descriptor()
}

func (TokenTypeEnum_TokenType) Type() protoreflect.EnumType {
	return &file_api_auth_v1_auth_v1_proto_enumTypes[3]
}

func (x TokenTypeEnum_TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenTypeEnum_TokenType.Descriptor instead.
func (TokenTypeEnum_TokenType) EnumDescriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{5, 0}
}

// Payload 授权信息
type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id 唯一id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// uid 唯一id
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// tt : token type 令牌类型
	Tt TokenTypeEnum_TokenType `protobuf:"varint,3,opt,name=tt,proto3,enum=kit.api.authv1.TokenTypeEnum_TokenType" json:"tt,omitempty"`
	// lp : login platform 平台信息
	Lp PlatformEnum_Platform `protobuf:"varint,4,opt,name=lp,proto3,enum=kit.api.authv1.PlatformEnum_Platform" json:"lp,omitempty"`
	// lt : limit type 登录限制类型
	Lt LimitTypeEnum_LimitType `protobuf:"varint,5,opt,name=lt,proto3,enum=kit.api.authv1.LimitTypeEnum_LimitType" json:"lt,omitempty"`
	// st : signing time 授权时间
	// 用途1：可用于判断： 强制登录，登出其他账户
	St *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=st,proto3" json:"st,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payload) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Payload) GetTt() TokenTypeEnum_TokenType {
	if x != nil {
		return x.Tt
	}
	return TokenTypeEnum_DEFAULT
}

func (x *Payload) GetLp() PlatformEnum_Platform {
	if x != nil {
		return x.Lp
	}
	return PlatformEnum_UNKNOWN
}

func (x *Payload) GetLt() LimitTypeEnum_LimitType {
	if x != nil {
		return x.Lt
	}
	return LimitTypeEnum_UNLIMITED
}

func (x *Payload) GetSt() *timestamppb.Timestamp {
	if x != nil {
		return x.St
	}
	return nil
}

// Auth 验证数据、缓存数据、...
type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data 存储 用户数据、其他数据
	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// payload 授权信息
	Payload *Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// secret 密码
	Secret string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	// status 状态
	Status StatusEnum_Status `protobuf:"varint,4,opt,name=status,proto3,enum=kit.api.authv1.StatusEnum_Status" json:"status,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{1}
}

func (x *Auth) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Auth) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Auth) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Auth) GetStatus() StatusEnum_Status {
	if x != nil {
		return x.Status
	}
	return StatusEnum_UNSPECIFIED
}

// StatusEnum 用户状态
type StatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusEnum) Reset() {
	*x = StatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusEnum) ProtoMessage() {}

func (x *StatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusEnum.ProtoReflect.Descriptor instead.
func (*StatusEnum) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{2}
}

// Platform 平台标识
type PlatformEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlatformEnum) Reset() {
	*x = PlatformEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformEnum) ProtoMessage() {}

func (x *PlatformEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformEnum.ProtoReflect.Descriptor instead.
func (*PlatformEnum) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{3}
}

// LimitTypeEnum 登录限制类型
type LimitTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LimitTypeEnum) Reset() {
	*x = LimitTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitTypeEnum) ProtoMessage() {}

func (x *LimitTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitTypeEnum.ProtoReflect.Descriptor instead.
func (*LimitTypeEnum) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{4}
}

// TokenTypeEnum 令牌类型
type TokenTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TokenTypeEnum) Reset() {
	*x = TokenTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenTypeEnum) ProtoMessage() {}

func (x *TokenTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenTypeEnum.ProtoReflect.Descriptor instead.
func (*TokenTypeEnum) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_v1_proto_rawDescGZIP(), []int{5}
}

var File_api_auth_v1_auth_v1_proto protoreflect.FileDescriptor

var file_api_auth_v1_auth_v1_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x69, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x02, 0x74, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x74, 0x74, 0x12, 0x35,
	0x0a, 0x02, 0x6c, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6b, 0x69, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x02, 0x6c, 0x70, 0x12, 0x37, 0x0a, 0x02, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x6c, 0x74, 0x12, 0x2a,
	0x0a, 0x02, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x73, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x78, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0x6a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x48, 0x49, 0x54, 0x45, 0x4c, 0x49, 0x53, 0x54, 0x10,
	0x04, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x05,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x06, 0x22, 0xb1, 0x01,
	0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xa0,
	0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50,
	0x55, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x64, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x65, 0x12, 0x08, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10,
	0xc8, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x48, 0x35, 0x10, 0xc9, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x41,
	0x50, 0x50, 0x4c, 0x45, 0x54, 0x10, 0xcb, 0x01, 0x12, 0x08, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x10,
	0xcc, 0x01, 0x12, 0x10, 0x0a, 0x0b, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50,
	0x50, 0x10, 0xcd, 0x01, 0x12, 0x0c, 0x0a, 0x07, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x10,
	0xce, 0x01, 0x12, 0x0d, 0x0a, 0x08, 0x49, 0x50, 0x41, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x10, 0xcf,
	0x01, 0x22, 0x4b, 0x0a, 0x0d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0x3a, 0x0a, 0x09, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x22, 0x71,
	0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22,
	0x60, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45,
	0x42, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02,
	0x48, 0x35, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10,
	0x07, 0x42, 0x54, 0x0a, 0x0e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x76, 0x31, 0x42, 0x0c, 0x4b, 0x69, 0x74, 0x41, 0x70, 0x69, 0x41, 0x75, 0x74, 0x68, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76,
	0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_auth_v1_auth_v1_proto_rawDescOnce sync.Once
	file_api_auth_v1_auth_v1_proto_rawDescData = file_api_auth_v1_auth_v1_proto_rawDesc
)

func file_api_auth_v1_auth_v1_proto_rawDescGZIP() []byte {
	file_api_auth_v1_auth_v1_proto_rawDescOnce.Do(func() {
		file_api_auth_v1_auth_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_v1_auth_v1_proto_rawDescData)
	})
	return file_api_auth_v1_auth_v1_proto_rawDescData
}

var file_api_auth_v1_auth_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_auth_v1_auth_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_auth_v1_auth_v1_proto_goTypes = []interface{}{
	(StatusEnum_Status)(0),        // 0: kit.api.authv1.StatusEnum.Status
	(PlatformEnum_Platform)(0),    // 1: kit.api.authv1.PlatformEnum.Platform
	(LimitTypeEnum_LimitType)(0),  // 2: kit.api.authv1.LimitTypeEnum.LimitType
	(TokenTypeEnum_TokenType)(0),  // 3: kit.api.authv1.TokenTypeEnum.TokenType
	(*Payload)(nil),               // 4: kit.api.authv1.Payload
	(*Auth)(nil),                  // 5: kit.api.authv1.Auth
	(*StatusEnum)(nil),            // 6: kit.api.authv1.StatusEnum
	(*PlatformEnum)(nil),          // 7: kit.api.authv1.PlatformEnum
	(*LimitTypeEnum)(nil),         // 8: kit.api.authv1.LimitTypeEnum
	(*TokenTypeEnum)(nil),         // 9: kit.api.authv1.TokenTypeEnum
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 11: google.protobuf.Any
}
var file_api_auth_v1_auth_v1_proto_depIdxs = []int32{
	3,  // 0: kit.api.authv1.Payload.tt:type_name -> kit.api.authv1.TokenTypeEnum.TokenType
	1,  // 1: kit.api.authv1.Payload.lp:type_name -> kit.api.authv1.PlatformEnum.Platform
	2,  // 2: kit.api.authv1.Payload.lt:type_name -> kit.api.authv1.LimitTypeEnum.LimitType
	10, // 3: kit.api.authv1.Payload.st:type_name -> google.protobuf.Timestamp
	11, // 4: kit.api.authv1.Auth.data:type_name -> google.protobuf.Any
	4,  // 5: kit.api.authv1.Auth.payload:type_name -> kit.api.authv1.Payload
	0,  // 6: kit.api.authv1.Auth.status:type_name -> kit.api.authv1.StatusEnum.Status
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_auth_v1_auth_v1_proto_init() }
func file_api_auth_v1_auth_v1_proto_init() {
	if File_api_auth_v1_auth_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_auth_v1_auth_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_api_auth_v1_auth_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_api_auth_v1_auth_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusEnum); i {
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
		file_api_auth_v1_auth_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformEnum); i {
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
		file_api_auth_v1_auth_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitTypeEnum); i {
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
		file_api_auth_v1_auth_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenTypeEnum); i {
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
			RawDescriptor: file_api_auth_v1_auth_v1_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_auth_v1_auth_v1_proto_goTypes,
		DependencyIndexes: file_api_auth_v1_auth_v1_proto_depIdxs,
		EnumInfos:         file_api_auth_v1_auth_v1_proto_enumTypes,
		MessageInfos:      file_api_auth_v1_auth_v1_proto_msgTypes,
	}.Build()
	File_api_auth_v1_auth_v1_proto = out.File
	file_api_auth_v1_auth_v1_proto_rawDesc = nil
	file_api_auth_v1_auth_v1_proto_goTypes = nil
	file_api_auth_v1_auth_v1_proto_depIdxs = nil
}
