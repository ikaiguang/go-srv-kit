package logutil

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// logger config
const (
	_defaultCallerSkipStd  = 3
	_defaultCallerSkipFile = 3
)

// Config 配置¬
type Config struct {
	Std  ConfigStd
	File ConfigFile
}

// ConfigStd 标准输出
type ConfigStd struct {
	Enable     bool      // 是否启用
	Level      log.Level // 日志级别
	CallerSkip int       // 日志 runtime caller skips(默认：_defaultCallerSkipStd)
}

// ConfigFile 输出到文件
type ConfigFile struct {
	Enable     bool      // 是否启用
	Level      log.Level // 日志级别
	CallerSkip int       // 日志 runtime caller skips(默认：_defaultCallerSkipFile)

	// 存储位置
	// Dir 文件夹
	Dir string
	// Filename 文件名(默认：${filename}_app.%Y%m%d.log)
	Filename string

	// 轮询：n久 或 文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	// RotateTime n久
	RotateTime time.Duration
	// RotateSize 文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	RotateSize int64

	// 存储：n个 或 有效期StorageAge(默认：30天)
	// StorageSize n个
	StorageSize uint
	// StorageAge 存储n久(默认：30天)
	StorageAge time.Duration
}
