package loggerutil

import (
	"context"
	"os"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"github.com/ikaiguang/go-srv-kit/kit/header"
	ippkg "github.com/ikaiguang/go-srv-kit/kit/ip"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	"go.opentelemetry.io/otel/trace"
)

type ServiceInfo struct {
	Project  string `json:"project"`
	Name     string `json:"name"`
	Env      string `json:"env"`
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

func (s *ServiceInfo) Kvs() []any {
	return []any{
		"app", s.Name,
		"ip", s.IP,
	}
}

func NewServiceInfo(appConfig *configpb.App) *ServiceInfo {
	res := &ServiceInfo{
		Project:  appConfig.GetProjectName(),
		Name:     appConfig.GetServerName(),
		Env:      appConfig.GetServerEnv(),
		Version:  appConfig.GetServerVersion(),
		Hostname: "",
		IP:       ippkg.LocalIP(),
	}
	res.Hostname, _ = os.Hostname()
	return res
}

type TracerInfo struct {
	Tracer log.Valuer
}

func (s *TracerInfo) Kvs() []any {
	return []any{
		"tracer", s.Tracer,
	}
}

func NewTracerInfo() *TracerInfo {
	res := &TracerInfo{
		Tracer: withTracerInfo(),
	}
	return res
}

func withTracerInfo() log.Valuer {
	return func(ctx context.Context) any {
		var (
			res = `{"trace_id":"`
		)

		// trace
		span := trace.SpanContextFromContext(ctx)
		tid := ""
		if span.HasTraceID() {
			tid = span.TraceID().String()
		} else if tr, ok := transport.FromServerContext(ctx); ok {
			tid = tr.RequestHeader().Get(headerpkg.RequestID)
		}
		res += tid + `"`

		// user
		if claims, ok := authpkg.GetAuthClaimsFromContext(ctx); ok && claims.Payload != nil {
			if claims.Payload.UserID > 0 {
				res += `,"user_id":"` + strconv.FormatUint(claims.Payload.UserID, 10) + `"`
			} else if claims.Payload.UserUuid != "" {
				res += `,"user_uuid":"` + claims.Payload.UserUuid + `"`
			}
		}
		return res + `}`
	}
}

type TracerInfoKvs struct {
	TraceId log.Valuer
	UserId  log.Valuer
}

func (s *TracerInfoKvs) Kvs() []any {
	return []any{
		"trace_id", s.TraceId,
		"x_uid", s.UserId,
	}
}

func NewTracerInfoKvs() *TracerInfoKvs {
	res := &TracerInfoKvs{
		TraceId: withTraceId(),
		UserId:  withUserId(),
	}
	return res
}

func withTraceId() log.Valuer {
	return func(ctx context.Context) any {
		// trace
		span := trace.SpanContextFromContext(ctx)
		tid := ""
		if span.HasTraceID() {
			return span.TraceID().String()
		} else if tr, ok := transport.FromServerContext(ctx); ok {
			tid = tr.RequestHeader().Get(headerpkg.RequestID)
		}
		return tid
	}
}

func withUserId() log.Valuer {
	return func(ctx context.Context) any {
		uid := ""
		if claims, ok := authpkg.GetAuthClaimsFromContext(ctx); ok && claims.Payload != nil {
			if claims.Payload.UserID > 0 {
				uid = strconv.FormatUint(claims.Payload.UserID, 10)
			} else if claims.Payload.UserUuid != "" {
				uid = claims.Payload.UserUuid
			}
		}
		return uid
	}
}
