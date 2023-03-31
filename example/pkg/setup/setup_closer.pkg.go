package setuppkg

import (
	"fmt"
	stdlog "log"
	"strings"

	pkgerrors "github.com/pkg/errors"
)

// Close .
func (s *engines) Close() (err error) {
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
			stdlog.Printf("|*** 退出程序：发生Panic：\n%s\n", strings.Join(errInfos, "\n"))
		}
		stdlog.Printf("|*** 退出程序：发生Panic：%v\n", panicRecover)
	}()

	// 关闭配置
	stdlog.Println("|*** 退出程序：关闭：配置处理手柄")
	if err = s.Config.Close(); err != nil {
		errorPrefix := "Config.Close error : "
		errInfos = append(errInfos, errorPrefix+err.Error())
	}

	// 缓存
	if s.redisClient != nil {
		stdlog.Println("|*** 退出程序：关闭：Redis客户端")
		err := s.redisClient.Close()
		if err != nil {
			errorPrefix := "redisClient.Close error : "
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 数据库
	if s.mysqlGormDB != nil {
		stdlog.Println("|*** 退出程序：关闭：MySQL-GORM")
		errorPrefix := "mysqlGormDB.Close error : "
		connPool, err := s.mysqlGormDB.DB()
		if err != nil {
			errInfos = append(errInfos, errorPrefix+err.Error())
		} else if err = connPool.Close(); err != nil {
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}
	if s.postgresGormDB != nil {
		stdlog.Println("|*** 退出程序：关闭：Postgres-GORM")
		errorPrefix := "postgresGormDB.Close error : "
		connPool, err := s.postgresGormDB.DB()
		if err != nil {
			errInfos = append(errInfos, errorPrefix+err.Error())
		} else if err = connPool.Close(); err != nil {
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	//s.consulClient
	//s.jaegerTraceExporter

	// 雪花算法
	if s.snowflakeStopChannel != nil {
		stdlog.Println("|*** 退出程序：关闭：雪花算法")
		s.snowflakeStopChannel <- 1
		close(s.snowflakeStopChannel)
	}

	// debug
	if len(s.debugHelperCloseFnSlice) > 0 {
		stdlog.Println("|*** 退出程序：关闭：调试工具debugutil")
	}
	for i := range s.debugHelperCloseFnSlice {
		err = s.debugHelperCloseFnSlice[i].Close()
		if err != nil {
			errorPrefix := fmt.Sprintf("debugHelperCloseFnSlice[%d] error : ", i+1)
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 日志
	if len(s.loggerCloseFnSlice) > 0 {
		stdlog.Println("|*** 退出程序：关闭：日志输出实例")
	}
	for i := range s.loggerCloseFnSlice {
		err = s.loggerCloseFnSlice[i].Close()
		if err != nil {
			errorPrefix := fmt.Sprintf("loggerCloseFnSlice[%d] error : ", i+1)
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 日志工具
	if len(s.loggerHelperCloseFnSlice) > 0 {
		stdlog.Println("|*** 退出程序：关闭：日志输出工具")
	}
	for i := range s.loggerHelperCloseFnSlice {
		err := s.loggerHelperCloseFnSlice[i].Close()
		if err != nil {
			errorPrefix := fmt.Sprintf("loggerHelperCloseFnSlice[%d] error : ", i+1)
			errInfos = append(errInfos, errorPrefix+err.Error())
		}
	}

	// 中间件日志工具
	if len(s.loggerMiddlewareCloseFnSlice) > 0 {
		stdlog.Println("|*** 退出程序：关闭：中间件日志输出工具")
	}
	for i := range s.loggerMiddlewareCloseFnSlice {
		err = s.loggerMiddlewareCloseFnSlice[i].Close()
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
		stdlog.Println("|*** 退出程序：关闭：日志Writer")
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
