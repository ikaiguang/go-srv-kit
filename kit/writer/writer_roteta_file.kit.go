package writerpkg

import (
	"io"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// 轮转日志参数
const (
	DefaultRotationTime            = time.Hour * 24      // 1天
	DefaultRotationSize            = 50 << 20            // 50M
	DefaultRotationStorageAge      = time.Hour * 24 * 30 // 30天
	DefaultRotationCounter         = 10086               // 10086个
	_defaultRotationFilenameSuffix = "_%Y%m%d%H%M%S.log" // 文件名后缀
)

// ConfigRotate 轮转输出
type ConfigRotate struct {
	// 存储位置
	// Dir 文件夹
	Dir string
	// Filename 文件名(默认：${filename}_app.%Y%m%d.log)
	Filename string

	// RotateTime 轮询规则：n久(默认：86400s # 86400s = 1天)
	// 轮询规则：默认为：RotateTime
	RotateTime time.Duration
	// RotateSize 轮询规则：按文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	// 轮询规则：默认为：RotateTime
	RotateSize int64

	// StorageAge 存储规则：n久(默认：30天)
	// 存储规则：默认为：StorageAge
	StorageAge time.Duration
	// StorageCounter 存储规则：n个(默认：10086个)
	// 存储规则：默认为：StorageAge
	StorageCounter uint
}

// NewRotateFile 轮转输出
func NewRotateFile(cfg *ConfigRotate, configOpts ...Option) (writer io.Writer, err error) {
	var (
		configOpt = &options{
			filenameSuffix: _defaultRotationFilenameSuffix,
		}
		rotateOpts []rotatelogs.Option
	)
	for i := range configOpts {
		configOpts[i](configOpt)
	}

	// 轮询 时间 或 文件大小
	switch {
	case cfg.RotateTime > 0:
		rotateOpts = append(rotateOpts, rotatelogs.WithRotationTime(cfg.RotateTime))
	case cfg.RotateSize > 0:
		rotateOpts = append(rotateOpts, rotatelogs.WithRotationSize(cfg.RotateSize))
	default:
		rotateOpts = append(rotateOpts, rotatelogs.WithRotationTime(DefaultRotationTime))
	}

	// 存储 n个 或 n久
	switch {
	case cfg.StorageCounter > 0:
		rotateOpts = append(rotateOpts, rotatelogs.WithRotationCount(cfg.StorageCounter))
	case cfg.StorageAge > 0:
		rotateOpts = append(rotateOpts, rotatelogs.WithMaxAge(cfg.StorageAge))
	default:
		rotateOpts = append(rotateOpts, rotatelogs.WithMaxAge(DefaultRotationStorageAge))
	}

	// 写
	writer, err = rotatelogs.New(
		filepath.Join(cfg.Dir, cfg.Filename+configOpt.filenameSuffix),
		rotateOpts...,
	)
	if err != nil {
		return
	}
	return
}
