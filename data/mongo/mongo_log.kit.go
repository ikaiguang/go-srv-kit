package mongopkg

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/event"
)

// NewMonitor ...
func NewMonitor(logger log.Logger) *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			_ = log.WithContext(ctx, logger).Log(log.LevelDebug,
				"request_id", evt.RequestID,
				"database", evt.DatabaseName,
				"command_name", evt.CommandName,
				"command_script", evt.Command.String(),
			)
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			_ = log.WithContext(ctx, logger).Log(log.LevelDebug,
				"request_id", evt.RequestID,
				"command_name", evt.CommandName,
				"exec_duration", evt.Duration.String(),
				// "result", evt.Reply.String(),
			)
		},
	}
}
