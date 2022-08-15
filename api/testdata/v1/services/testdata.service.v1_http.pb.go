// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package testdataservicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/ikaiguang/go-srv-kit/api/testdata/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvTestdataDelete = "/kit.api.testdataservicev1.SrvTestdata/Delete"
const OperationSrvTestdataGet = "/kit.api.testdataservicev1.SrvTestdata/Get"
const OperationSrvTestdataPatch = "/kit.api.testdataservicev1.SrvTestdata/Patch"
const OperationSrvTestdataPost = "/kit.api.testdataservicev1.SrvTestdata/Post"
const OperationSrvTestdataPut = "/kit.api.testdataservicev1.SrvTestdata/Put"
const OperationSrvTestdataWebsocket = "/kit.api.testdataservicev1.SrvTestdata/Websocket"

type SrvTestdataHTTPServer interface {
	Delete(context.Context, *resources.TestReq) (*resources.TestResp, error)
	Get(context.Context, *resources.TestReq) (*resources.TestResp, error)
	Patch(context.Context, *resources.TestReq) (*resources.TestResp, error)
	Post(context.Context, *resources.TestReq) (*resources.TestResp, error)
	Put(context.Context, *resources.TestReq) (*resources.TestResp, error)
	Websocket(context.Context, *resources.TestReq) (*resources.TestResp, error)
}

func RegisterSrvTestdataHTTPServer(s *http.Server, srv SrvTestdataHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/testdata/websocket", _SrvTestdata_Websocket0_HTTP_Handler(srv))
	r.GET("/api/v1/testdata/get", _SrvTestdata_Get0_HTTP_Handler(srv))
	r.PUT("/api/v1/testdata/put", _SrvTestdata_Put0_HTTP_Handler(srv))
	r.POST("/api/v1/testdata/post", _SrvTestdata_Post0_HTTP_Handler(srv))
	r.DELETE("/api/v1/testdata/post", _SrvTestdata_Delete0_HTTP_Handler(srv))
	r.PATCH("/api/v1/testdata/post", _SrvTestdata_Patch0_HTTP_Handler(srv))
}

func _SrvTestdata_Websocket0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.TestReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestdataWebsocket)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Websocket(ctx, req.(*resources.TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Get0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.TestReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestdataGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*resources.TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Put0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.TestReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestdataPut)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Put(ctx, req.(*resources.TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Post0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.TestReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestdataPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Post(ctx, req.(*resources.TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Delete0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.TestReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestdataDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*resources.TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Patch0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.TestReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvTestdataPatch)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Patch(ctx, req.(*resources.TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.TestResp)
		return ctx.Result(200, reply)
	}
}

type SrvTestdataHTTPClient interface {
	Delete(ctx context.Context, req *resources.TestReq, opts ...http.CallOption) (rsp *resources.TestResp, err error)
	Get(ctx context.Context, req *resources.TestReq, opts ...http.CallOption) (rsp *resources.TestResp, err error)
	Patch(ctx context.Context, req *resources.TestReq, opts ...http.CallOption) (rsp *resources.TestResp, err error)
	Post(ctx context.Context, req *resources.TestReq, opts ...http.CallOption) (rsp *resources.TestResp, err error)
	Put(ctx context.Context, req *resources.TestReq, opts ...http.CallOption) (rsp *resources.TestResp, err error)
	Websocket(ctx context.Context, req *resources.TestReq, opts ...http.CallOption) (rsp *resources.TestResp, err error)
}

type SrvTestdataHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvTestdataHTTPClient(client *http.Client) SrvTestdataHTTPClient {
	return &SrvTestdataHTTPClientImpl{client}
}

func (c *SrvTestdataHTTPClientImpl) Delete(ctx context.Context, in *resources.TestReq, opts ...http.CallOption) (*resources.TestResp, error) {
	var out resources.TestResp
	pattern := "/api/v1/testdata/post"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvTestdataDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Get(ctx context.Context, in *resources.TestReq, opts ...http.CallOption) (*resources.TestResp, error) {
	var out resources.TestResp
	pattern := "/api/v1/testdata/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvTestdataGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Patch(ctx context.Context, in *resources.TestReq, opts ...http.CallOption) (*resources.TestResp, error) {
	var out resources.TestResp
	pattern := "/api/v1/testdata/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvTestdataPatch))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Post(ctx context.Context, in *resources.TestReq, opts ...http.CallOption) (*resources.TestResp, error) {
	var out resources.TestResp
	pattern := "/api/v1/testdata/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvTestdataPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Put(ctx context.Context, in *resources.TestReq, opts ...http.CallOption) (*resources.TestResp, error) {
	var out resources.TestResp
	pattern := "/api/v1/testdata/put"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvTestdataPut))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Websocket(ctx context.Context, in *resources.TestReq, opts ...http.CallOption) (*resources.TestResp, error) {
	var out resources.TestResp
	pattern := "/api/v1/testdata/websocket"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvTestdataWebsocket))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}