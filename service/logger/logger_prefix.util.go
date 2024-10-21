package loggerutil

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	ippkg "github.com/ikaiguang/go-srv-kit/kit/ip"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
	"go.opentelemetry.io/otel/trace"
	"os"
	"strconv"
)

type ServiceInfo struct {
	Project  string `json:"project"`
	Name     string `json:"name"`
	Env      string `json:"env"`
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

func (s *ServiceInfo) Kvs() []interface{} {
	buf, _ := json.Marshal(s)
	return []interface{}{
		"app", string(buf),
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

func (s *TracerInfo) Kvs() []interface{} {
	return []interface{}{
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
	return func(ctx context.Context) interface{} {
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

		// span
		res += `,"span_id":"`
		spanId := ""
		if span.HasSpanID() {
			spanId = span.SpanID().String()
		}
		res += spanId + `"`

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
