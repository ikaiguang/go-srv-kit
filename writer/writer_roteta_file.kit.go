package writerutil

import (
	"io"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
)

// 轮转日志参数
const (
	DefaultRotationStorageAge      = time.Hour * 24 * 30     // 30天
	DefaultRotationSize            = 50 << 20                // 50M
	_defaultRotationFilenameSuffix = "_app.%Y%m%d%H%M%S.log" // 文件名后缀
)

// ConfigRotate 轮转输出
type ConfigRotate struct {
	// 存储位置
	// Dir 文件夹
	Dir string
	// Filename 文件名(默认：${filename}_app.%Y%m%d.log)
	Filename string

	// RotateTime 轮询规则：n久 或 文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	RotateTime time.Duration
	// RotateSize 轮询规则：按文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	RotateSize int64

	// StorageCounter 存储：n个 或 有效期StorageAge(默认：30天)
	StorageCounter uint
	// StorageAge 存储n久(默认：30天)
	StorageAge time.Duration
}

// NewRotateFile 轮转输出
func NewRotateFile(cfg *ConfigRotate) (writer io.Writer, err error) {
	var opts []rotatelogs.Option

	// 轮询 时间 或 文件大小
	switch {
	case cfg.RotateTime > 0:
		opts = append(opts, rotatelogs.WithRotationTime(cfg.RotateTime))
	case cfg.RotateSize > 0:
		opts = append(opts, rotatelogs.WithRotationSize(cfg.RotateSize))
	default:
		opts = append(opts, rotatelogs.WithRotationSize(DefaultRotationSize))
	}

	// 存储 n个 或 n久
	switch {
	case cfg.StorageCounter > 0:
		opts = append(opts, rotatelogs.WithRotationCount(cfg.StorageCounter))
	case cfg.StorageAge > 0:
		opts = append(opts, rotatelogs.WithMaxAge(cfg.StorageAge))
	default:
		opts = append(opts, rotatelogs.WithMaxAge(DefaultRotationStorageAge))
	}

	// 写
	writer, err = rotatelogs.New(
		filepath.Join(cfg.Dir, cfg.Filename+_defaultRotationFilenameSuffix),
		opts...,
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
