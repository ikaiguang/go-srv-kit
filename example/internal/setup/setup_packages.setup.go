package setup

import (
	"fmt"
	"io"
	stdlog "log"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	writerutil "github.com/ikaiguang/go-srv-kit/kit/writer"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	loghelper "github.com/ikaiguang/go-srv-kit/log/helper"
	mysqlutil "github.com/ikaiguang/go-srv-kit/mysql"
	redisutil "github.com/ikaiguang/go-srv-kit/redis"
)

// up 启动手柄
type up struct {
	Config

	// loggerFileWriterMutex 日志文件写手柄
	loggerFileWriterMutex sync.Once
	loggerFileWriter      io.Writer

	// debugHelperCloseFnSlice debug工具
	debugHelperCloseFnSlice []func() error

	// loggerMutex 日志
	loggerMutex                  sync.Once
	logger                       log.Logger
	loggerCloseFnSlice           []func() error
	loggerHelperMutex            sync.Once
	loggerHelper                 log.Logger
	loggerHelperCloseFnSlice     []func() error
	loggerMiddlewareMutex        sync.Once
	loggerMiddleware             log.Logger
	loggerMiddlewareCloseFnSlice []func() error

	// mysqlGormMutex mysql gorm
	mysqlGormMutex sync.Once
	mysqlGormDB    *gorm.DB

	// redisClientMutex redis客户端
	redisClientMutex sync.Once
	redisClient      *redis.Client
}

// NewUpPackages .
func NewUpPackages(conf Config) Packages {
	return newUpHandler(conf)
}

// newUpHandler .
func newUpHandler(conf Config) *up {
	return &up{
		Config: conf,
	}
}

// Close .
func (s *up) Close() (err error) {
	// 退出程序
	stdlog.Println("|==================== 退出程序 开始 ====================|")
	defer stdlog.Println("|==================== 退出程序 结束 ====================|")

	var errInfos []string
	defer func() {
		if len(errInfos) > 0 {
			err = pkgerrors.New(strings.Join(errInfos, "；\n"))
		}
	}()

	// 发生Panic
	defer func() {
		panicRecover := recover()
		if panicRecover == nil {
			return
		}

		// Panic
		if len(errInfos) > 0 {
			stdlog.Printf("|*** 退出程序 发生Panic：\n%s\n", strings.Join(errInfos, "\n"))
		}
		stdlog.Printf("|*** 退出程序 发生Panic：%v\n", panicRecover)
	}()

	// 缓存
	if s.redisClient != nil {
		stdlog.Println("|*** 退出程序：关闭Redis客户端")
		err := s.redisClient.Close()
		if err != nil {
			errorPrefix := "redisClient.Close error : "
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 数据库
	if s.mysqlGormDB != nil {
		stdlog.Println("|*** 退出程序：关闭MySQL-GORM")
		errorPrefix := "mysqlGormDB.Close error : "
		connPool, err := s.mysqlGormDB.DB()
		if err != nil {
			errInfos = append(errInfos, errorPrefix+err.Error())
		} else if err = connPool.Close(); err != nil {
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// debug
	if len(s.debugHelperCloseFnSlice) > 0 {
		stdlog.Println("|*** 退出程序：关闭调试工具debugutil")
	}
	for i := range s.debugHelperCloseFnSlice {
		err := s.debugHelperCloseFnSlice[i]()
		if err != nil {
			errorPrefix := fmt.Sprintf("debugHelperCloseFnSlice[%d] error : ", i+1)
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 日志
	if len(s.loggerCloseFnSlice) > 0 {
		stdlog.Println("|*** 退出程序：关闭日志输出实例")
	}
	for i := range s.loggerCloseFnSlice {
		err := s.loggerCloseFnSlice[i]()
		if err != nil {
			errorPrefix := fmt.Sprintf("loggerCloseFnSlice[%d] error : ", i+1)
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 日志工具
	if len(s.loggerHelperCloseFnSlice) > 0 {
		stdlog.Println("|*** 退出程序：关闭日志输出工具")
	}
	for i := range s.loggerHelperCloseFnSlice {
		err := s.loggerHelperCloseFnSlice[i]()
		if err != nil {
			errorPrefix := fmt.Sprintf("loggerHelperCloseFnSlice[%d] error : ", i+1)
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 中间件日志工具
	if len(s.loggerMiddlewareCloseFnSlice) > 0 {
		stdlog.Println("|*** 退出程序：关闭中间件日志输出工具")
	}
	for i := range s.loggerMiddlewareCloseFnSlice {
		err := s.loggerMiddlewareCloseFnSlice[i]()
		if err != nil {
			errorPrefix := fmt.Sprintf("loggerMiddlewareCloseFnSlice[%d] error : ", i+1)
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// Writer
	type closer interface {
		Close() error
	}
	if writerCloser, ok := s.loggerFileWriter.(closer); ok {
		stdlog.Println("|*** 退出程序：关闭Writer")
		if err := writerCloser.Close(); err != nil {
			errorPrefix := "loggerFileWriter.Close error : "
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 有错误
	if len(errInfos) > 0 {
		err = pkgerrors.New(strings.Join(errInfos, "；\n"))
		return err
	}
	return err
}

// LoggerFileWriter 文件日志写手柄
func (s *up) LoggerFileWriter() (io.Writer, error) {
	var err error
	s.loggerFileWriterMutex.Do(func() {
		s.loggerFileWriter, err = s.setupLoggerFileWriter()
	})
	if err != nil {
		return nil, err
	}
	if s.loggerFileWriter != nil {
		return s.loggerFileWriter, err
	}

	s.loggerFileWriter, err = s.setupLoggerFileWriter()
	if err != nil {
		return nil, err
	}
	return s.loggerFileWriter, err
}

// Logger 日志处理示例
func (s *up) Logger() (log.Logger, []func() error, error) {
	var (
		err error
	)
	s.loggerMutex.Do(func() {
		s.logger, s.loggerCloseFnSlice, err = s.setupLogger()
	})
	if err != nil {
		return nil, nil, err
	}
	if s.logger != nil {
		return s.logger, s.loggerCloseFnSlice, err
	}
	s.logger, s.loggerCloseFnSlice, err = s.setupLogger()
	if err != nil {
		return nil, nil, err
	}
	return s.logger, s.loggerCloseFnSlice, err
}

// LoggerHelper 日志处理示例
func (s *up) LoggerHelper() (log.Logger, []func() error, error) {
	var err error
	s.loggerHelperMutex.Do(func() {
		s.loggerHelper, s.loggerHelperCloseFnSlice, err = s.setupLoggerHelper()
	})
	if err != nil {
		return nil, nil, err
	}
	if s.loggerHelper != nil {
		return s.loggerHelper, s.loggerHelperCloseFnSlice, err
	}
	s.loggerHelper, s.loggerHelperCloseFnSlice, err = s.setupLoggerHelper()
	if err != nil {
		return nil, nil, err
	}
	return s.loggerHelper, s.loggerHelperCloseFnSlice, err
}

// LoggerMiddleware 中间件的日志处理示例
func (s *up) LoggerMiddleware() (log.Logger, []func() error, error) {
	var err error
	s.loggerMiddlewareMutex.Do(func() {
		s.loggerMiddleware, s.loggerMiddlewareCloseFnSlice, err = s.setupLoggerMiddleware()
	})
	if err != nil {
		return nil, nil, err
	}
	if s.loggerMiddleware != nil {
		return s.loggerMiddleware, s.loggerMiddlewareCloseFnSlice, err
	}
	s.loggerMiddleware, s.loggerMiddlewareCloseFnSlice, err = s.setupLoggerMiddleware()
	if err != nil {
		return nil, nil, err
	}
	return s.loggerMiddleware, s.loggerMiddlewareCloseFnSlice, err
}

// LoggerFileWriter 文件日志写手柄
func (s *up) MysqlGormDB() (*gorm.DB, error) {
	var err error
	s.mysqlGormMutex.Do(func() {
		s.mysqlGormDB, err = s.setupMysqlGormDB()
	})
	if err != nil {
		return nil, err
	}
	if s.mysqlGormDB != nil {
		return s.mysqlGormDB, err
	}

	s.mysqlGormDB, err = s.setupMysqlGormDB()
	if err != nil {
		return nil, err
	}
	return s.mysqlGormDB, err
}

// RedisClient redis 客户端
func (s *up) RedisClient() (*redis.Client, error) {
	var err error
	s.redisClientMutex.Do(func() {
		s.redisClient, err = s.setupRedisClient()
	})
	if err != nil {
		return nil, err
	}
	if s.redisClient != nil {
		return s.redisClient, err
	}
	s.redisClient, err = s.setupRedisClient()
	if err != nil {
		return nil, err
	}
	return s.redisClient, err
}

// setupDebugUtil 设置调试工具
func (s *up) setupDebugUtil() error {
	if !s.Config.IsDebugMode() {
		return nil
	}
	stdlog.Printf("|*** 加载调试工具debugutil")
	syncFn, err := debugutil.Setup()
	if err != nil {
		return err
	}
	s.debugHelperCloseFnSlice = append(s.debugHelperCloseFnSlice, syncFn)
	return err
}

// setupLogHelper 设置日志工具
func (s *up) setupLogHelper() (closeFnSlice []func() error, err error) {
	loggerInstance, closeFnSlice, err := s.LoggerHelper()
	if err != nil {
		return closeFnSlice, err
	}
	if loggerInstance == nil {
		stdlog.Println("|*** 未加载日志工具")
		return closeFnSlice, err
	}

	// 日志
	loggerConfig := s.Config.LoggerConfig()
	if s.Config.EnableLoggingConsole() && loggerConfig.Console != nil {
		stdlog.Println("|*** 加载日志工具：日志输出到控制台")
	}
	if s.Config.EnableLoggingFile() && loggerConfig.File != nil {
		stdlog.Println("|*** 加载日志工具：日志输出到文件")
	}

	loghelper.Setup(loggerInstance)
	return closeFnSlice, err
}

// setupLoggerFileWriter 启动日志文件写手柄
func (s *up) setupLoggerFileWriter() (io.Writer, error) {
	loggerConfig := s.Config.LoggerConfig()
	if !s.Config.EnableLoggingFile() || loggerConfig.File == nil {
		stdlog.Println("|*** 加载日志工具：虚拟的文件写手柄")
		return writerutil.NewDummyWriter()
	}
	rotateConfig := &writerutil.ConfigRotate{
		Dir:            loggerConfig.File.Dir,
		Filename:       loggerConfig.File.Filename,
		RotateTime:     loggerConfig.File.RotateTime.AsDuration(),
		RotateSize:     loggerConfig.File.RotateSize,
		StorageCounter: uint(loggerConfig.File.StorageCounter),
		StorageAge:     loggerConfig.File.StorageAge.AsDuration(),
	}
	return writerutil.NewRotateFile(rotateConfig)
}

// setupLogger 初始化日志输出实例
func (s *up) setupLogger() (logger log.Logger, closeFnSlice []func() error, err error) {
	skip := logutil.DefaultCallerSkip + 1
	return s.setupLoggerWithCallerSkip(skip)
}

// setupLoggerHelper 初始化日志工具输出实例
func (s *up) setupLoggerHelper() (logger log.Logger, closeFnSlice []func() error, err error) {
	skip := logutil.DefaultCallerSkip + 2
	return s.setupLoggerWithCallerSkip(skip)
}

// setupLoggerMiddleware 初始化中间价的日志输出实例
func (s *up) setupLoggerMiddleware() (logger log.Logger, closeFnSlice []func() error, err error) {
	skip := logutil.DefaultCallerSkip - 1
	return s.setupLoggerWithCallerSkip(skip)
}

// setupLoggerWithCallerSkip 初始化日志输出实例
func (s *up) setupLoggerWithCallerSkip(skip int) (logger log.Logger, closeFnSlice []func() error, err error) {
	// loggers
	var loggers []log.Logger

	// 配置
	loggerConfig := s.Config.LoggerConfig()
	if loggerConfig == nil {
		return logger, closeFnSlice, err
	}

	// 日志 输出到控制台
	stdLogger, err := logutil.NewDummyLogger()
	if err != nil {
		return logger, closeFnSlice, err
	}
	if s.Config.EnableLoggingConsole() && loggerConfig.Console != nil {
		stdLoggerConfig := &logutil.ConfigStd{
			Level:      logutil.ParseLevel(loggerConfig.Console.Level),
			CallerSkip: skip,
		}
		stdLoggerImpl, err := logutil.NewStdLogger(stdLoggerConfig)
		if err != nil {
			return logger, closeFnSlice, err
		}
		closeFnSlice = append(closeFnSlice, stdLoggerImpl.Sync)
		stdLogger = stdLoggerImpl
	}
	loggers = append(loggers, stdLogger)

	// 日志 输出到文件
	if s.Config.EnableLoggingFile() && loggerConfig.File != nil {
		// file logger
		fileLoggerConfig := &logutil.ConfigFile{
			Level:      logutil.ParseLevel(loggerConfig.File.Level),
			CallerSkip: skip,

			Dir:      loggerConfig.File.Dir,
			Filename: loggerConfig.File.Filename,

			RotateTime: loggerConfig.File.RotateTime.AsDuration(),
			RotateSize: loggerConfig.File.RotateSize,

			StorageCounter: uint(loggerConfig.File.StorageCounter),
			StorageAge:     loggerConfig.File.StorageAge.AsDuration(),
		}
		writer, err := s.LoggerFileWriter()
		if err != nil {
			return logger, closeFnSlice, err
		}
		fileLogger, err := logutil.NewFileLogger(
			fileLoggerConfig,
			logutil.WithWriter(writer),
		)
		closeFnSlice = append(closeFnSlice, fileLogger.Sync)
		if err != nil {
			return logger, closeFnSlice, err
		}
		loggers = append(loggers, fileLogger)
	}

	// 日志工具
	if len(loggers) == 0 {
		return logger, closeFnSlice, err
	}
	return log.MultiLogger(loggers...), closeFnSlice, err
}

// setupMysqlGormDB mysql gorm 数据库
func (s *up) setupMysqlGormDB() (*gorm.DB, error) {
	if s.Config.MySQLConfig() == nil {
		stdlog.Println("|*** 加载MySQL-GORM：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载MySQL-GORM：...")

	// logger writer
	var (
		writers []logger.Writer
		opts    []mysqlutil.Option
	)
	if s.Config.EnableLoggingConsole() {
		writers = append(writers, mysqlutil.NewStdWriter())
	}
	if s.Config.EnableLoggingFile() {
		writer, err := s.LoggerFileWriter()
		if err != nil {
			return nil, err
		}
		writers = append(writers, mysqlutil.NewJSONWriter(writer))
	}
	if len(writers) > 0 {
		opts = append(opts, mysqlutil.WithWriters(writers...))
	}
	return mysqlutil.NewMysqlDB(s.Config.MySQLConfig(), opts...)
}

// setupRedisClient redis 客户端
func (s *up) setupRedisClient() (*redis.Client, error) {
	if s.Config.RedisConfig() == nil {
		stdlog.Println("|*** 加载Redis客户端：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载Redis客户端：...")

	return redisutil.NewDB(s.Config.RedisConfig())
}
