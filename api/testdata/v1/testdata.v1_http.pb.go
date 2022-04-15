// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package testdatav1

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

type SrvTestdataHTTPServer interface {
	Delete(context.Context, *TestReq) (*TestResp, error)
	Get(context.Context, *TestReq) (*TestResp, error)
	Patch(context.Context, *TestReq) (*TestResp, error)
	Post(context.Context, *TestReq) (*TestResp, error)
	Put(context.Context, *TestReq) (*TestResp, error)
	Websocket(context.Context, *TestReq) (*TestResp, error)
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
		var in TestReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/kit.api.testdata.testdatav1.SrvTestdata/Websocket")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Websocket(ctx, req.(*TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Get0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/kit.api.testdata.testdatav1.SrvTestdata/Get")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Put0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/kit.api.testdata.testdatav1.SrvTestdata/Put")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Put(ctx, req.(*TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Post0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/kit.api.testdata.testdatav1.SrvTestdata/Post")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Post(ctx, req.(*TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Delete0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/kit.api.testdata.testdatav1.SrvTestdata/Delete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestResp)
		return ctx.Result(200, reply)
	}
}

func _SrvTestdata_Patch0_HTTP_Handler(srv SrvTestdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/kit.api.testdata.testdatav1.SrvTestdata/Patch")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Patch(ctx, req.(*TestReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestResp)
		return ctx.Result(200, reply)
	}
}

type SrvTestdataHTTPClient interface {
	Delete(ctx context.Context, req *TestReq, opts ...http.CallOption) (rsp *TestResp, err error)
	Get(ctx context.Context, req *TestReq, opts ...http.CallOption) (rsp *TestResp, err error)
	Patch(ctx context.Context, req *TestReq, opts ...http.CallOption) (rsp *TestResp, err error)
	Post(ctx context.Context, req *TestReq, opts ...http.CallOption) (rsp *TestResp, err error)
	Put(ctx context.Context, req *TestReq, opts ...http.CallOption) (rsp *TestResp, err error)
	Websocket(ctx context.Context, req *TestReq, opts ...http.CallOption) (rsp *TestResp, err error)
}

type SrvTestdataHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvTestdataHTTPClient(client *http.Client) SrvTestdataHTTPClient {
	return &SrvTestdataHTTPClientImpl{client}
}

func (c *SrvTestdataHTTPClientImpl) Delete(ctx context.Context, in *TestReq, opts ...http.CallOption) (*TestResp, error) {
	var out TestResp
	pattern := "/api/v1/testdata/post"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/kit.api.testdata.testdatav1.SrvTestdata/Delete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Get(ctx context.Context, in *TestReq, opts ...http.CallOption) (*TestResp, error) {
	var out TestResp
	pattern := "/api/v1/testdata/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/kit.api.testdata.testdatav1.SrvTestdata/Get"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Patch(ctx context.Context, in *TestReq, opts ...http.CallOption) (*TestResp, error) {
	var out TestResp
	pattern := "/api/v1/testdata/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/kit.api.testdata.testdatav1.SrvTestdata/Patch"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Post(ctx context.Context, in *TestReq, opts ...http.CallOption) (*TestResp, error) {
	var out TestResp
	pattern := "/api/v1/testdata/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/kit.api.testdata.testdatav1.SrvTestdata/Post"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Put(ctx context.Context, in *TestReq, opts ...http.CallOption) (*TestResp, error) {
	var out TestResp
	pattern := "/api/v1/testdata/put"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/kit.api.testdata.testdatav1.SrvTestdata/Put"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvTestdataHTTPClientImpl) Websocket(ctx context.Context, in *TestReq, opts ...http.CallOption) (*TestResp, error) {
	var out TestResp
	pattern := "/api/v1/testdata/websocket"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/kit.api.testdata.testdatav1.SrvTestdata/Websocket"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
