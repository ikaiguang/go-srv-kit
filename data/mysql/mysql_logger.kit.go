package mysqlutil

import (
	"gorm.io/gorm/logger"
)

// NewLogger 数据库日志
func NewLogger(conf *logger.Config, writers ...logger.Writer) logger.Interface {
	if len(writers) == 0 {
		return logger.New(NewStdWriter(), *conf)
	}
	w := &multiWriter{
		writers: writers,
	}
	return logger.New(w, *conf)
}
