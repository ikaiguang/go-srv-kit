// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package exampleservicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/ikaiguang/go-srv-kit/api/example/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvExampleExample = "/kit.api.exampleservicev1.SrvExample/Example"

type SrvExampleHTTPServer interface {
	Example(context.Context, *resources.ExampleReq) (*resources.ExampleResp, error)
}

func RegisterSrvExampleHTTPServer(s *http.Server, srv SrvExampleHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/example/example-testing", _SrvExample_Example0_HTTP_Handler(srv))
}

func _SrvExample_Example0_HTTP_Handler(srv SrvExampleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.ExampleReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvExampleExample)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Example(ctx, req.(*resources.ExampleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.ExampleResp)
		return ctx.Result(200, reply)
	}
}

type SrvExampleHTTPClient interface {
	Example(ctx context.Context, req *resources.ExampleReq, opts ...http.CallOption) (rsp *resources.ExampleResp, err error)
}

type SrvExampleHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvExampleHTTPClient(client *http.Client) SrvExampleHTTPClient {
	return &SrvExampleHTTPClientImpl{client}
}

func (c *SrvExampleHTTPClientImpl) Example(ctx context.Context, in *resources.ExampleReq, opts ...http.CallOption) (*resources.ExampleResp, error) {
	var out resources.ExampleResp
	pattern := "/api/v1/example/example-testing"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvExampleExample))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}