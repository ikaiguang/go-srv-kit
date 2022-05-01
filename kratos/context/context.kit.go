package contextutil

import (
	"context"
	"net"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// MatchHTTPContext 匹配
func MatchHTTPContext(ctx context.Context) (http.Context, bool) {
	httpCtx, ok := ctx.(http.Context)
	return httpCtx, ok
}

// ClientIP 获取客户端IP
func ClientIP(ctx context.Context) string {
	httpCtx, ok := MatchHTTPContext(ctx)
	if ok {
		return ClientIPFromHTTP(httpCtx)
	}
	return ClientIPFromGRPC(ctx)
}

// ClientIPFromHTTP ...
func ClientIPFromHTTP(ctx http.Context) string {
	ips := strings.Split(ctx.Header().Get("X-Forwarded-For"), ",")
	for i := len(ips) - 1; i >= 0; i-- {
		if clientIP, isValid := isValidIP(ips[i]); isValid {
			return clientIP
		}
	}

	ips = strings.Split(ctx.Header().Get("X-Real-Ip"), ",")
	for i := len(ips) - 1; i >= 0; i-- {
		if clientIP, isValid := isValidIP(ips[i]); isValid {
			return clientIP
		}
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(ctx.Request().RemoteAddr))
	if err != nil {
		return ""
	}
	if remoteIP := net.ParseIP(ip); remoteIP == nil {
		return remoteIP.String()
	}
	return ""
}

// ClientIPFromGRPC ...
func ClientIPFromGRPC(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ips := md.Get("x-forwarded-for")
		for i := len(ips) - 1; i >= 0; i-- {
			if clientIP, isValid := isValidIP(ips[i]); isValid {
				return clientIP
			}
		}

		ips = md.Get("x-real-ip")
		for i := len(ips) - 1; i >= 0; i-- {
			if clientIP, isValid := isValidIP(ips[i]); isValid {
				return clientIP
			}
		}
	}

	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			return tcpAddr.IP.String()
		}
		return pr.Addr.String()
	}
	return ""
}

// isValidIP 有效的ip
func isValidIP(ip string) (clientIP string, isValid bool) {
	clientIP = strings.TrimSpace(ip)
	isValid = net.ParseIP(ip) != nil
	return clientIP, isValid
}
