package contextpkg

import (
	"context"
	"net"
	"strings"

	"github.com/go-kratos/kratos/v3/transport/http"
	ippkg "github.com/ikaiguang/go-srv-kit/kit/v3/ip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// ClientIP 获取客户端IP
func ClientIP(ctx context.Context) string {
	// http
	tr, ok := MatchHTTPServerContext(ctx)
	if ok {
		return ClientIPFromHTTP(ctx, tr.Request())
	}
	// grpc
	return ClientIPFromGRPC(ctx)
}

// ClientIPFromHTTP ...
func ClientIPFromHTTP(_ context.Context, r *http.Request) (clientIP string) {
	if r == nil {
		return ""
	}
	// Check if we're running on a trusted platform, continue running backwards if error
	if platformHeader := trustedPlatform(); platformHeader != "" {
		// Developers can define their own header of Trusted Platform or use predefined constants
		clientIP = strings.TrimSpace(r.Header.Get(platformHeader))
		if ippkg.IsValidIP(clientIP) {
			return clientIP
		}
	}

	ips := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
	for i := len(ips) - 1; i >= 0; i-- {
		if clientIP = strings.TrimSpace(ips[i]); ippkg.IsValidIP(clientIP) {
			return clientIP
		}
	}

	ips = strings.Split(r.Header.Get("X-Real-Ip"), ",")
	for i := len(ips) - 1; i >= 0; i-- {
		if clientIP = strings.TrimSpace(ips[i]); ippkg.IsValidIP(clientIP) {
			return clientIP
		}
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return clientIP
	}
	if clientIP = strings.TrimSpace(ip); ippkg.IsValidIP(clientIP) {
		return clientIP
	}
	return clientIP
}

// ClientIPFromGRPC ...
func ClientIPFromGRPC(ctx context.Context) (clientIP string) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// Check if we're running on a trusted platform, continue running backwards if error
		if platformHeader := trustedPlatform(); platformHeader != "" {
			// Developers can define their own header of Trusted Platform or use predefined constants
			if addrSlice := md.Get(strings.ToLower(platformHeader)); len(addrSlice) > 0 {
				clientIP = strings.TrimSpace(addrSlice[0])
				if ippkg.IsValidIP(clientIP) {
					return clientIP
				}
			}
		}

		ips := md.Get("x-forwarded-for")
		for i := len(ips) - 1; i >= 0; i-- {
			if clientIP = strings.TrimSpace(ips[i]); ippkg.IsValidIP(clientIP) {
				return clientIP
			}
		}

		ips = md.Get("x-real-ip")
		for i := len(ips) - 1; i >= 0; i-- {
			if clientIP = strings.TrimSpace(ips[i]); ippkg.IsValidIP(clientIP) {
				return clientIP
			}
		}
	}

	if pr, ok := peer.FromContext(ctx); ok && pr.Addr != nil {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			clientIP = tcpAddr.IP.String()
			return clientIP
		}
		clientIP = pr.Addr.String()
		return clientIP
	}
	return clientIP
}

// clientIPKey ...
type clientIPKey struct{}

// SetClientIPToContext puts a client IP into context.
func SetClientIPToContext(ctx context.Context, clientIP string) context.Context {
	return context.WithValue(ctx, clientIPKey{}, clientIP)
}

// GetClientIPFromContext extracts a client IP from context.
func GetClientIPFromContext(ctx context.Context) (clientIP string, ok bool) {
	clientIP, ok = ctx.Value(clientIPKey{}).(string)
	return
}

// Deprecated: use SetClientIPToContext instead.
//
//nolint:revive // compatibility alias
func SetClientIpToContext(ctx context.Context, clientIP string) context.Context {
	return SetClientIPToContext(ctx, clientIP)
}

// Deprecated: use GetClientIPFromContext instead.
//
//nolint:revive // compatibility alias
func GetClientIpFromContext(ctx context.Context) (clientIP string, ok bool) {
	return GetClientIPFromContext(ctx)
}
