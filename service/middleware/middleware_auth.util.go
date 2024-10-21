package middlewareutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	contextpkg "github.com/ikaiguang/go-srv-kit/kratos/context"
)

// TransportServiceKind 通行类型
type TransportServiceKind string

const (
	TransportServiceKindALL           = "ALL"
	TransportServiceKindHTTP          = "HTTP"
	TransportServiceKindGRPC          = "GRPC"
	TransportServiceKindMethodGet     = "GET"
	TransportServiceKindMethodHead    = "HEAD"
	TransportServiceKindMethodPost    = "POST"
	TransportServiceKindMethodPut     = "PUT"
	TransportServiceKindMethodPatch   = "PATCH"
	TransportServiceKindMethodDelete  = "DELETE"
	TransportServiceKindMethodConnect = "CONNECT"
	TransportServiceKindMethodOptions = "OPTIONS"
	TransportServiceKindMethodTrace   = "TRACE"
)

func MergeWhitelist(whitelist ...map[string]TransportServiceKind) map[string]TransportServiceKind {
	list := make(map[string]TransportServiceKind)
	for _, item := range whitelist {
		for k, v := range item {
			list[k] = v
		}
	}
	return list
}

func (s TransportServiceKind) MatchServiceKind(ctx context.Context) bool {
	switch s {
	default:
		return true
	case TransportServiceKindALL:
		return true
	case TransportServiceKindHTTP:
		tr, ok := transport.FromServerContext(ctx)
		return ok && tr.Kind() == transport.KindHTTP
	case TransportServiceKindGRPC:
		tr, ok := transport.FromServerContext(ctx)
		return ok && tr.Kind() == transport.KindGRPC
	}
	//return false
}

// NewWhiteListMatcher 路由白名单
func NewWhiteListMatcher(whiteList map[string]TransportServiceKind) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		// operation
		if tsk, ok := whiteList[operation]; ok && tsk.MatchServiceKind(ctx) {
			return false
		}

		// http path
		if tr, ok := contextpkg.MatchHTTPServerContext(ctx); ok {
			if sk, ok := whiteList[tr.Request().URL.Path]; ok {
				if sk == TransportServiceKindALL || sk == "" || string(sk) == tr.Request().Method {
					return false
				}
			}
		}

		return true
	}
}

// NewAuthMiddleware 验证中间
func NewAuthMiddleware(authTokenRepo authpkg.AuthRepo, whiteList map[string]TransportServiceKind) (m middleware.Middleware, err error) {
	m = selector.Server(
		authpkg.Server(
			authTokenRepo.JWTSigningKeyFunc,
			authpkg.WithSigningMethod(authTokenRepo.JWTSigningMethod()),
			authpkg.WithClaims(authTokenRepo.JWTSigningClaims),
			authpkg.WithAccessTokenValidator(authTokenRepo.VerifyAccessToken),
		),
	).
		Match(NewWhiteListMatcher(whiteList)).
		Build()

	return m, err
}
