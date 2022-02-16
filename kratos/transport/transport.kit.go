package transportutil

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"
)

// TransportKind 服务类型
func TransportKind(ctx context.Context) (kind transport.Kind, ok bool) {
	trInfo, ok := transport.FromServerContext(ctx)
	if !ok {
		return kind, ok
	}
	kind = trInfo.Kind()

	return kind, ok
}
