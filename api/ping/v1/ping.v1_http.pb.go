// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package pingv1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvPingPing = "/kit.api.ping.pingv1.SrvPing/Ping"

type SrvPingHTTPServer interface {
	Ping(context.Context, *PingReq) (*PingResp, error)
}

func RegisterSrvPingHTTPServer(s *http.Server, srv SrvPingHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/ping/{message}", _SrvPing_Ping0_HTTP_Handler(srv))
}

func _SrvPing_Ping0_HTTP_Handler(srv SrvPingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PingReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvPingPing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*PingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PingResp)
		return ctx.Result(200, reply)
	}
}

type SrvPingHTTPClient interface {
	Ping(ctx context.Context, req *PingReq, opts ...http.CallOption) (rsp *PingResp, err error)
}

type SrvPingHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvPingHTTPClient(client *http.Client) SrvPingHTTPClient {
	return &SrvPingHTTPClientImpl{client}
}

func (c *SrvPingHTTPClientImpl) Ping(ctx context.Context, in *PingReq, opts ...http.CallOption) (*PingResp, error) {
	var out PingResp
	pattern := "/api/v1/ping/{message}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvPingPing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
