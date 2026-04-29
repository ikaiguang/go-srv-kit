package mongopkg

// Level 日志级别，与 github.com/go-kratos/kratos/v2/log.Level 值兼容。
// 两者底层类型均为 int8，枚举值一一对应，可通过类型转换互转。
type Level int8

const (
	LevelDebug Level = iota - 1
	LevelInfo
	LevelWarn
	LevelError
)

// Logger 日志接口，方法签名与 github.com/go-kratos/kratos/v2/log.Logger 对应。
//
// 注意：由于 Go 接口要求参数类型完全匹配，kratos 的 log.Logger（参数为 log.Level）
// 不能直接赋值给本接口（参数为本包的 Level）。请使用 LogAdapter 进行转换：
//
//	import kratoslog "github.com/go-kratos/kratos/v2/log"
//	var adapter mongopkg.Logger = mongopkg.LogAdapter(func(level mongopkg.Level, keyvals ...any) error {
//	    return kratosLogger.Log(kratoslog.Level(level), keyvals...)
//	})
type Logger interface {
	Log(level Level, keyvals ...any) error
}

// LogAdapter 日志函数适配器，方便将外部 Logger 转换为本包的 Logger 接口。
type LogAdapter func(level Level, keyvals ...any) error

// Log 实现 Logger 接口。
func (f LogAdapter) Log(level Level, keyvals ...any) error {
	return f(level, keyvals...)
}
