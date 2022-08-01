package contextutil

import (
	"context"
	"net"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	iputil "github.com/ikaiguang/go-srv-kit/kit/ip"
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
func ClientIPFromHTTP(ctx context.Context, r *http.Request) (clientIP string) {
	// Check if we're running on a trusted platform, continue running backwards if error
	if defaultTrustedPlatform != "" {
		// Developers can define their own header of Trusted Platform or use predefined constants
		if clientIP = r.Header.Get(defaultTrustedPlatform); clientIP != "" {
			return clientIP
		}
	}

	ips := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
	for i := len(ips) - 1; i >= 0; i-- {
		if clientIP = strings.TrimSpace(ips[i]); iputil.IsValidIP(clientIP) {
			return clientIP
		}
	}

	ips = strings.Split(r.Header.Get("X-Real-Ip"), ",")
	for i := len(ips) - 1; i >= 0; i-- {
		if clientIP = strings.TrimSpace(ips[i]); iputil.IsValidIP(clientIP) {
			return clientIP
		}
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return clientIP
	}
	if clientIP = strings.TrimSpace(ip); iputil.IsValidIP(clientIP) {
		return clientIP
	}
	return clientIP
}

// ClientIPFromGRPC ...
func ClientIPFromGRPC(ctx context.Context) (clientIP string) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// Check if we're running on a trusted platform, continue running backwards if error
		if defaultTrustedPlatform != "" {
			// Developers can define their own header of Trusted Platform or use predefined constants
			if addrSlice := md.Get(strings.ToLower(defaultTrustedPlatform)); len(addrSlice) > 0 {
				clientIP = addrSlice[0]
				return clientIP
			}
		}

		ips := md.Get("x-forwarded-for")
		for i := len(ips) - 1; i >= 0; i-- {
			if clientIP = strings.TrimSpace(ips[i]); iputil.IsValidIP(clientIP) {
				return clientIP
			}
		}

		ips = md.Get("x-real-ip")
		for i := len(ips) - 1; i >= 0; i-- {
			if clientIP = strings.TrimSpace(ips[i]); iputil.IsValidIP(clientIP) {
				return clientIP
			}
		}
	}

	if pr, ok := peer.FromContext(ctx); ok {
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

// SetClientIpToContext put client ip into context
func SetClientIpToContext(ctx context.Context, clientIp string) context.Context {
	return context.WithValue(ctx, clientIPKey{}, clientIp)
}

// GetClientIpFromContext extract client ip from context
func GetClientIpFromContext(ctx context.Context) (clientIp string, ok bool) {
	clientIp, ok = ctx.Value(clientIPKey{}).(string)
	return
}
