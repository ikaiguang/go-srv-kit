package mongopkg

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/event"
	"time"
)

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
func NewMonitor(logger log.Logger, opts ...MonitorOption) *event.CommandMonitor {
	options := &monitorOption{}
	for i := range opts {
		opts[i](options)
	}
	loggerHandler := log.NewHelper(log.With(logger, "module", "mongo-driver-monitor"))
	return &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			loggerHandler.WithContext(ctx).Debugw(
				"request_id", evt.RequestID,
				"database", evt.DatabaseName,
				"command_name", evt.CommandName,
				"command", evt.Command.String(),
			)
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			kvs := []interface{}{
				"request_id", evt.RequestID,
				"database", evt.DatabaseName,
				"command_name", evt.CommandName,
				"duration", evt.Duration.String(),
				//"result", evt.Reply.String(),
			}

			if options.slowThreshold > 0 && evt.Duration >= options.slowThreshold {
				loggerHandler.WithContext(ctx).Warnw(kvs...)
			} else {
				loggerHandler.WithContext(ctx).Debugw(kvs...)
			}
		},
		Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
			loggerHandler.WithContext(ctx).Warnw(
				"request_id", evt.RequestID,
				"database", evt.DatabaseName,
				"command_name", evt.CommandName,
				"duration", evt.Duration.String(),
				"error", evt.Failure,
			)
		},
	}
}
