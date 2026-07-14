package writerpkg

import (
	"errors"
	"io"
	"math"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// 轮转日志参数
const (
	DefaultRotationTime            = time.Hour * 24      // 1天
	DefaultRotationSize            = 100 << 20           // 100M
	DefaultRotationStorageAge      = time.Hour * 24 * 30 // 30天
	DefaultRotationCounter         = 10086               // 10086个
	_defaultRotationFilenameSuffix = ".log"              // 文件名后缀
	bytesPerMegabyte               = 1 << 20
)

// ConfigRotate 轮转输出
type ConfigRotate struct {
	// 存储位置
	// Dir 文件夹
	Dir string
	// Filename 文件名(默认：${filename}_app.%Y%m%d.log)
	Filename string

	// RotateTime 兼容旧配置；lumberjack 不支持纯时间轮转
	// 如果设置 RotateTime 且未设置 RotateSize，将返回错误
	RotateTime time.Duration
	// RotateSize 轮询规则：按文件大小RotateSize(默认：52428800 # 100<<20 = 100M)
	// 轮询规则：默认为：RotateSize
	RotateSize int64

	// StorageAge 存储规则：n久(默认：30天)
	// 存储规则：默认为：StorageAge
	StorageAge time.Duration
	// StorageCounter 存储规则：n个(默认：10086个)
	// 存储规则：默认为：StorageAge
	StorageCounter uint
	// Compress 是否压缩归档日志文件
	Compress bool
}

// NewRotateFile 轮转输出
func NewRotateFile(cfg *ConfigRotate, configOpts ...Option) (writer io.Writer, err error) {
	if cfg == nil {
		return nil, errors.New("rotate file config is nil")
	}

	configOpt := &options{
		filenameSuffix: _defaultRotationFilenameSuffix,
	}
	for i := range configOpts {
		configOpts[i](configOpt)
	}

	if cfg.RotateTime > 0 && cfg.RotateSize <= 0 {
		return nil, errors.New("rotate time is not supported by lumberjack; set rotate size instead")
	}

	rotateSize := cfg.RotateSize
	if rotateSize <= 0 {
		rotateSize = DefaultRotationSize
	}

	storageAge := cfg.StorageAge
	if cfg.StorageCounter <= 0 && storageAge <= 0 {
		storageAge = DefaultRotationStorageAge
	}

	return &lumberjack.Logger{
		Filename:   filepath.Join(cfg.Dir, cfg.Filename+configOpt.filenameSuffix),
		MaxSize:    bytesToMegabytes(rotateSize),
		MaxAge:     durationToDays(storageAge),
		MaxBackups: int(cfg.StorageCounter),
		LocalTime:  true,
		Compress:   cfg.Compress,
	}, nil
}

func bytesToMegabytes(size int64) int {
	if size <= 0 {
		return 0
	}
	return int(math.Ceil(float64(size) / bytesPerMegabyte))
}

func durationToDays(duration time.Duration) int {
	if duration <= 0 {
		return 0
	}
	return int(math.Ceil(duration.Hours() / 24))
}
