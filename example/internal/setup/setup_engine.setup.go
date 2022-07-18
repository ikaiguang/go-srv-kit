package setup

import (
	"io"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// engines 引擎模块
type engines struct {
	Config

	// loggerPrefixFieldMutex 日志前缀
	loggerPrefixFieldMutex sync.Once
	loggerPrefixField      *LoggerPrefixField

	// loggerFileWriterMutex 日志文件写手柄
	loggerFileWriterMutex sync.Once
	loggerFileWriter      io.Writer

	// debugHelperCloseFnSlice debug工具
	debugHelperCloseFnSlice []io.Closer

	// loggerMutex 日志
	loggerMutex                  sync.Once
	logger                       log.Logger
	loggerCloseFnSlice           []io.Closer
	loggerHelperMutex            sync.Once
	loggerHelper                 log.Logger
	loggerHelperCloseFnSlice     []io.Closer
	loggerMiddlewareMutex        sync.Once
	loggerMiddleware             log.Logger
	loggerMiddlewareCloseFnSlice []io.Closer

	// mysqlGormMutex mysql gorm
	mysqlGormMutex sync.Once
	mysqlGormDB    *gorm.DB

	// postgresGormMutex mysql gorm
	postgresGormMutex sync.Once
	postgresGormDB    *gorm.DB

	// redisClientMutex redis客户端
	redisClientMutex sync.Once
	redisClient      *redis.Client
}

// NewEngine ...
func NewEngine(conf Config) *engines {
	return &engines{
		Config: conf,
	}
}
