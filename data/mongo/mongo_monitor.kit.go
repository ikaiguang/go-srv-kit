package mongopkg

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/event"
)

// moduleKey 模块标识常量
const moduleKey = "module"

// moduleMonitor 监控模块名称常量
const moduleMonitor = "mongo-driver-monitor"

type monitorOption struct {
	slowThreshold time.Duration
}

type MonitorOption func(*monitorOption)

// WithSlowThreshold 慢查询阈值
func WithSlowThreshold(slowThreshold time.Duration) MonitorOption {
	return func(o *monitorOption) {
		o.slowThreshold = slowThreshold
	}
}

// NewMonitor ...
func NewMonitor(logger Logger, opts ...MonitorOption) *event.CommandMonitor {
	options := &monitorOption{}
	for i := range opts {
		opts[i](options)
	}
	return &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			_ = logger.Log(LevelInfo,
				moduleKey, moduleMonitor,
				"request_id", evt.RequestID,
				"database", evt.DatabaseName,
				"command_name", evt.CommandName,
				"command", evt.Command.String(),
			)
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			kvs := []any{
				moduleKey, moduleMonitor,
				"request_id", evt.RequestID,
				"database", evt.DatabaseName,
				"command_name", evt.CommandName,
				"duration", evt.Duration.String(),
			}

			if options.slowThreshold > 0 && evt.Duration >= options.slowThreshold {
				_ = logger.Log(LevelWarn, kvs...)
			} else {
				_ = logger.Log(LevelDebug, kvs...)
			}
		},
		Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
			_ = logger.Log(LevelWarn,
				moduleKey, moduleMonitor,
				"request_id", evt.RequestID,
				"database", evt.DatabaseName,
				"command_name", evt.CommandName,
				"duration", evt.Duration.String(),
				"error", evt.Failure,
			)
		},
	}
}
