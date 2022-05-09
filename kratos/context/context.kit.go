package contextutil

import (
	"context"
	"net"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
	iputil "github.com/ikaiguang/go-srv-kit/kit/ip"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// TrustedPlatform 信任的平台
var (
	defaultTrustedPlatform = headerutil.RemoteAddr
)

// SetTrustedPlatform 设置信任的平台
func SetTrustedPlatform(platformHeader string) {
	defaultTrustedPlatform = platformHeader
}

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
	// Check if we're running on a trusted platform, continue running backwards if error
	if defaultTrustedPlatform != "" {
		// Developers can define their own header of Trusted Platform or use predefined constants
		if addr := ctx.Header().Get(defaultTrustedPlatform); addr != "" {
			return addr
		}
	}

	ips := strings.Split(ctx.Header().Get("X-Forwarded-For"), ",")
	for i := len(ips) - 1; i >= 0; i-- {
		if clientIP := strings.TrimSpace(ips[i]); iputil.IsValidIP(clientIP) {
			return clientIP
		}
	}

	ips = strings.Split(ctx.Header().Get("X-Real-Ip"), ",")
	for i := len(ips) - 1; i >= 0; i-- {
		if clientIP := strings.TrimSpace(ips[i]); iputil.IsValidIP(clientIP) {
			return clientIP
		}
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(ctx.Request().RemoteAddr))
	if err != nil {
		return ""
	}
	if clientIP := strings.TrimSpace(ip); iputil.IsValidIP(clientIP) {
		return clientIP
	}
	return ""
}

// ClientIPFromGRPC ...
func ClientIPFromGRPC(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// Check if we're running on a trusted platform, continue running backwards if error
		if defaultTrustedPlatform != "" {
			// Developers can define their own header of Trusted Platform or use predefined constants
			if addrSlice := md.Get(strings.ToLower(defaultTrustedPlatform)); len(addrSlice) > 0 {
				return addrSlice[0]
			}
		}

		ips := md.Get("x-forwarded-for")
		for i := len(ips) - 1; i >= 0; i-- {
			if clientIP := strings.TrimSpace(ips[i]); iputil.IsValidIP(clientIP) {
				return clientIP
			}
		}

		ips = md.Get("x-real-ip")
		for i := len(ips) - 1; i >= 0; i-- {
			if clientIP := strings.TrimSpace(ips[i]); iputil.IsValidIP(clientIP) {
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
