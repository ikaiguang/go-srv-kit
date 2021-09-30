// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: conf.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Bootstrap 配置引导
type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// app application
	App *App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	// log 日志
	Log *Log `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	// data 数据
	Data *Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Bootstrap) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// App application
type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// env app 环境
	Env string `protobuf:"bytes,1,opt,name=env,proto3" json:"env,omitempty"`
	// name app名字
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// version app版本
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// endpoints app站点
	Endpoints []string `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// metadata 元数据
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1}
}

func (x *App) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *App) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *App) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Log 日志
type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// console 输出到控制台
	Console *Log_Console `protobuf:"bytes,1,opt,name=console,proto3" json:"console,omitempty"`
	// file 输出到文件
	File *Log_File `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Log) GetConsole() *Log_Console {
	if x != nil {
		return x.Console
	}
	return nil
}

func (x *Log) GetFile() *Log_File {
	if x != nil {
		return x.File
	}
	return nil
}

// Data 数据
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mysql *Data_MySQL `protobuf:"bytes,1,opt,name=mysql,proto3" json:"mysql,omitempty"`
	Redis *Data_Redis `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetMysql() *Data_MySQL {
	if x != nil {
		return x.Mysql
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

// Console 输出到控制台
type Log_Console struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enable 是否启用
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// level 日志级别；DEBUG、INFO、WARN、ERROR、FATAL
	Level string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Log_Console) Reset() {
	*x = Log_Console{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log_Console) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log_Console) ProtoMessage() {}

func (x *Log_Console) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log_Console.ProtoReflect.Descriptor instead.
func (*Log_Console) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Log_Console) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Log_Console) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

// File 输出到文件
type Log_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enable 是否启用
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// level 日志级别；DEBUG、INFO、WARN、ERROR、FATAL
	Level string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	// dir 存储目录
	Dir string `protobuf:"bytes,3,opt,name=dir,proto3" json:"dir,omitempty"`
	// filename 文件名(默认：${filename}_app.%Y%m%d%H%M%S.log)
	Filename string `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	// rotate_time 轮询规则：n久 或 文件大小rotate_size(默认：52428800 # 50<<20 = 50M)
	RotateTime *durationpb.Duration `protobuf:"bytes,5,opt,name=rotate_time,json=rotateTime,proto3" json:"rotate_time,omitempty"`
	// rotate_size 轮询规则：按文件大小rotate_size(默认：52428800 # 50<<20 = 50M)
	RotateSize int64 `protobuf:"varint,6,opt,name=rotate_size,json=rotateSize,proto3" json:"rotate_size,omitempty"`
	// storage_counter 存储：n个 或 有效期storage_age(默认：2592000s = 30天)
	StorageCounter uint32 `protobuf:"varint,7,opt,name=storage_counter,json=storageCounter,proto3" json:"storage_counter,omitempty"`
	// storage_age 存储n久(默认：2592000s = 30天)
	StorageAge *durationpb.Duration `protobuf:"bytes,8,opt,name=storage_age,json=storageAge,proto3" json:"storage_age,omitempty"`
}

func (x *Log_File) Reset() {
	*x = Log_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log_File) ProtoMessage() {}

func (x *Log_File) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log_File.ProtoReflect.Descriptor instead.
func (*Log_File) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Log_File) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Log_File) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Log_File) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *Log_File) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Log_File) GetRotateTime() *durationpb.Duration {
	if x != nil {
		return x.RotateTime
	}
	return nil
}

func (x *Log_File) GetRotateSize() int64 {
	if x != nil {
		return x.RotateSize
	}
	return 0
}

func (x *Log_File) GetStorageCounter() uint32 {
	if x != nil {
		return x.StorageCounter
	}
	return 0
}

func (x *Log_File) GetStorageAge() *durationpb.Duration {
	if x != nil {
		return x.StorageAge
	}
	return nil
}

// MySQL MySQL
type Data_MySQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dsn string `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty"`
	// conn_max_idle_time 设置连接空闲的最长时间
	SlowThreshold *durationpb.Duration `protobuf:"bytes,2,opt,name=slow_threshold,json=slowThreshold,proto3" json:"slow_threshold,omitempty"`
	LoggerEnable  bool                 `protobuf:"varint,3,opt,name=logger_enable,json=loggerEnable,proto3" json:"logger_enable,omitempty"`
	// logger_level 日志级别；值：DEBUG、INFO、WARN、ERROR、FATAL
	LoggerLevel string `protobuf:"bytes,4,opt,name=logger_level,json=loggerLevel,proto3" json:"logger_level,omitempty"`
	// conn_max_active 连接可复用的最大时间
	ConnMaxActive uint32 `protobuf:"varint,5,opt,name=conn_max_active,json=connMaxActive,proto3" json:"conn_max_active,omitempty"`
	// conn_max_lifetime 可复用的最大时间
	ConnMaxLifetime *durationpb.Duration `protobuf:"bytes,6,opt,name=conn_max_lifetime,json=connMaxLifetime,proto3" json:"conn_max_lifetime,omitempty"`
	// conn_max_idle 连接池中空闲连接的最大数量
	ConnMaxIdle uint32 `protobuf:"varint,7,opt,name=conn_max_idle,json=connMaxIdle,proto3" json:"conn_max_idle,omitempty"`
	// conn_max_idle_time 设置连接空闲的最长时间
	ConnMaxIdleTime *durationpb.Duration `protobuf:"bytes,8,opt,name=conn_max_idle_time,json=connMaxIdleTime,proto3" json:"conn_max_idle_time,omitempty"`
}

func (x *Data_MySQL) Reset() {
	*x = Data_MySQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_MySQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_MySQL) ProtoMessage() {}

func (x *Data_MySQL) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_MySQL.ProtoReflect.Descriptor instead.
func (*Data_MySQL) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Data_MySQL) GetDsn() string {
	if x != nil {
		return x.Dsn
	}
	return ""
}

func (x *Data_MySQL) GetSlowThreshold() *durationpb.Duration {
	if x != nil {
		return x.SlowThreshold
	}
	return nil
}

func (x *Data_MySQL) GetLoggerEnable() bool {
	if x != nil {
		return x.LoggerEnable
	}
	return false
}

func (x *Data_MySQL) GetLoggerLevel() string {
	if x != nil {
		return x.LoggerLevel
	}
	return ""
}

func (x *Data_MySQL) GetConnMaxActive() uint32 {
	if x != nil {
		return x.ConnMaxActive
	}
	return 0
}

func (x *Data_MySQL) GetConnMaxLifetime() *durationpb.Duration {
	if x != nil {
		return x.ConnMaxLifetime
	}
	return nil
}

func (x *Data_MySQL) GetConnMaxIdle() uint32 {
	if x != nil {
		return x.ConnMaxIdle
	}
	return 0
}

func (x *Data_MySQL) GetConnMaxIdleTime() *durationpb.Duration {
	if x != nil {
		return x.ConnMaxIdleTime
	}
	return nil
}

// Redis redis
type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Username     string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password     string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Db           uint32               `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`
	DialTimeout  *durationpb.Duration `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,6,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	// conn_max_active 连接的最大数量
	ConnMaxActive uint32 `protobuf:"varint,8,opt,name=conn_max_active,json=connMaxActive,proto3" json:"conn_max_active,omitempty"`
	// conn_max_lifetime 连接可复用的最大时间
	ConnMaxLifetime *durationpb.Duration `protobuf:"bytes,9,opt,name=conn_max_lifetime,json=connMaxLifetime,proto3" json:"conn_max_lifetime,omitempty"`
	// conn_max_idle 连接池中空闲连接的最大数量
	ConnMaxIdle uint32 `protobuf:"varint,10,opt,name=conn_max_idle,json=connMaxIdle,proto3" json:"conn_max_idle,omitempty"`
	// conn_max_idle_time 设置连接空闲的最长时间
	ConnMaxIdleTime *durationpb.Duration `protobuf:"bytes,11,opt,name=conn_max_idle_time,json=connMaxIdleTime,proto3" json:"conn_max_idle_time,omitempty"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() uint32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Data_Redis) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Data_Redis) GetConnMaxActive() uint32 {
	if x != nil {
		return x.ConnMaxActive
	}
	return 0
}

func (x *Data_Redis) GetConnMaxLifetime() *durationpb.Duration {
	if x != nil {
		return x.ConnMaxLifetime
	}
	return nil
}

func (x *Data_Redis) GetConnMaxIdle() uint32 {
	if x != nil {
		return x.ConnMaxIdle
	}
	return 0
}

func (x *Data_Redis) GetConnMaxIdleTime() *durationpb.Duration {
	if x != nil {
		return x.ConnMaxIdleTime
	}
	return nil
}

var File_conf_proto protoreflect.FileDescriptor

var file_conf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x12, 0x1e, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61,
	0x70, 0x70, 0x12, 0x1e, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c,
	0x6f, 0x67, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd8, 0x01, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xbc, 0x03, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x1a,
	0x37, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0xa4, 0x02, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a,
	0x0b, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x61,
	0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x67, 0x65, 0x22,
	0xda, 0x07, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x05, 0x6d, 0x79, 0x73, 0x71,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x79, 0x53, 0x51, 0x4c, 0x52, 0x05, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x12, 0x29, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0xfe,
	0x02, 0x0a, 0x05, 0x4d, 0x79, 0x53, 0x51, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x73, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x73, 0x6c,
	0x6f, 0x77, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73,
	0x6c, 0x6f, 0x77, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x45, 0x0a, 0x11,
	0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f,
	0x69, 0x64, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e,
	0x4d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x1a,
	0xfa, 0x03, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x64, 0x62, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6e,
	0x4d, 0x61, 0x78, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x63, 0x6f, 0x6e,
	0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78,
	0x49, 0x64, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x6e, 0x4d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x57, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6b,
	0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61,
	0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_proto_rawDescOnce sync.Once
	file_conf_proto_rawDescData = file_conf_proto_rawDesc
)

func file_conf_proto_rawDescGZIP() []byte {
	file_conf_proto_rawDescOnce.Do(func() {
		file_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_proto_rawDescData)
	})
	return file_conf_proto_rawDescData
}

var file_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: conf.v1.Bootstrap
	(*App)(nil),                 // 1: conf.v1.App
	(*Log)(nil),                 // 2: conf.v1.Log
	(*Data)(nil),                // 3: conf.v1.Data
	nil,                         // 4: conf.v1.App.MetadataEntry
	(*Log_Console)(nil),         // 5: conf.v1.Log.Console
	(*Log_File)(nil),            // 6: conf.v1.Log.File
	(*Data_MySQL)(nil),          // 7: conf.v1.Data.MySQL
	(*Data_Redis)(nil),          // 8: conf.v1.Data.Redis
	(*durationpb.Duration)(nil), // 9: google.protobuf.Duration
}
var file_conf_proto_depIdxs = []int32{
	1,  // 0: conf.v1.Bootstrap.app:type_name -> conf.v1.App
	2,  // 1: conf.v1.Bootstrap.log:type_name -> conf.v1.Log
	3,  // 2: conf.v1.Bootstrap.data:type_name -> conf.v1.Data
	4,  // 3: conf.v1.App.metadata:type_name -> conf.v1.App.MetadataEntry
	5,  // 4: conf.v1.Log.console:type_name -> conf.v1.Log.Console
	6,  // 5: conf.v1.Log.file:type_name -> conf.v1.Log.File
	7,  // 6: conf.v1.Data.mysql:type_name -> conf.v1.Data.MySQL
	8,  // 7: conf.v1.Data.redis:type_name -> conf.v1.Data.Redis
	9,  // 8: conf.v1.Log.File.rotate_time:type_name -> google.protobuf.Duration
	9,  // 9: conf.v1.Log.File.storage_age:type_name -> google.protobuf.Duration
	9,  // 10: conf.v1.Data.MySQL.slow_threshold:type_name -> google.protobuf.Duration
	9,  // 11: conf.v1.Data.MySQL.conn_max_lifetime:type_name -> google.protobuf.Duration
	9,  // 12: conf.v1.Data.MySQL.conn_max_idle_time:type_name -> google.protobuf.Duration
	9,  // 13: conf.v1.Data.Redis.dial_timeout:type_name -> google.protobuf.Duration
	9,  // 14: conf.v1.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	9,  // 15: conf.v1.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	9,  // 16: conf.v1.Data.Redis.conn_max_lifetime:type_name -> google.protobuf.Duration
	9,  // 17: conf.v1.Data.Redis.conn_max_idle_time:type_name -> google.protobuf.Duration
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_conf_proto_init() }
func file_conf_proto_init() {
	if File_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
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
		file_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log_Console); i {
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
		file_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log_File); i {
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
		file_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_MySQL); i {
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
		file_conf_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
			RawDescriptor: file_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_proto_goTypes,
		DependencyIndexes: file_conf_proto_depIdxs,
		MessageInfos:      file_conf_proto_msgTypes,
	}.Build()
	File_conf_proto = out.File
	file_conf_proto_rawDesc = nil
	file_conf_proto_goTypes = nil
	file_conf_proto_depIdxs = nil
}
