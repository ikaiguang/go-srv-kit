package curlpkg

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultTimeout timeout
	DefaultTimeout = time.Minute

	// ContentType header
	ContentType         = "Content-Type"
	ContentTypeJSON     = "application/json"
	ContentTypeJSONUtf8 = "application/json; charset=utf-8"
	ContentTypePB       = "application/x-protobuf"

	// UserAgent Accept header
	UserAgent = "User-Agent"
)

// NewPostRequest Post请求
func NewPostRequest(httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewPostRequestContext(context.Background(), httpURL, body)
}

// NewGetRequest Get请求
func NewGetRequest(httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewGetRequestContext(context.Background(), httpURL, body)
}

// NewPutRequest Put请求
func NewPutRequest(httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewPutRequestContext(context.Background(), httpURL, body)
}

// NewPatchRequest Patch请求
func NewPatchRequest(httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewPatchRequestContext(context.Background(), httpURL, body)
}

// NewDeleteRequest Delete请求
func NewDeleteRequest(httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewDeleteRequestContext(context.Background(), httpURL, body)
}

// NewRequest .
func NewRequest(httpMethod, httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewRequestContext(context.Background(), httpMethod, httpURL, body)
}

// NewPostRequestContext Post请求（支持 Context）
func NewPostRequestContext(ctx context.Context, httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewRequestContext(ctx, http.MethodPost, httpURL, body)
}

// NewGetRequestContext Get请求（支持 Context）
func NewGetRequestContext(ctx context.Context, httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewRequestContext(ctx, http.MethodGet, httpURL, body)
}

// NewPutRequestContext Put请求（支持 Context）
func NewPutRequestContext(ctx context.Context, httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewRequestContext(ctx, http.MethodPut, httpURL, body)
}

// NewPatchRequestContext Patch请求（支持 Context）
func NewPatchRequestContext(ctx context.Context, httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewRequestContext(ctx, http.MethodPatch, httpURL, body)
}

// NewDeleteRequestContext Delete请求（支持 Context）
func NewDeleteRequestContext(ctx context.Context, httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	return NewRequestContext(ctx, http.MethodDelete, httpURL, body)
}

// NewRequestContext 创建 HTTP 请求（支持 Context）
func NewRequestContext(ctx context.Context, httpMethod, httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return http.NewRequestWithContext(ctx, httpMethod, httpURL, body)
}

// NewHTTPClient http客户端
func NewHTTPClient(opts ...Option) *http.Client {
	// 可选项
	options := options{
		timeout: DefaultTimeout,
	}
	for _, o := range opts {
		o(&options)
	}

	// client
	httpClient := &http.Client{Timeout: options.timeout}

	// 仅在显式启用时跳过 TLS 证书验证
	if options.insecureSkipVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient.Transport = tr
	}

	return httpClient
}

// Do 请求
func Do(httpReq *http.Request, opts ...Option) (httpCode int, bodyBytes []byte, err error) {
	httpClient := NewHTTPClient(opts...)
	defer httpClient.CloseIdleConnections()

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return httpCode, bodyBytes, err
	}
	defer func() { _ = httpResp.Body.Close() }()

	return response(httpResp)
}

// Default http.DefaultClient
func Default(httpReq *http.Request) (httpCode int, bodyBytes []byte, err error) {
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return httpCode, bodyBytes, err
	}
	defer func() { _ = httpResp.Body.Close() }()

	return response(httpResp)
}

// DoWithClient 请求一次后关闭连接
func DoWithClient(httpClient *http.Client, httpReq *http.Request) (httpCode int, bodyBytes []byte, err error) {
	// 哪里开启，哪里关闭
	//defer httpClient.CloseIdleConnections()

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return httpCode, bodyBytes, err
	}
	defer func() { _ = httpResp.Body.Close() }()

	return response(httpResp)
}

// response
func response(httpResp *http.Response) (httpCode int, bodyBytes []byte, err error) {
	//defer func() { _ = httpResp.Body.Close() }()

	// resp
	httpCode = httpResp.StatusCode
	bodyBytes, err = io.ReadAll(httpResp.Body)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return httpCode, bodyBytes, err
	}
	return httpCode, bodyBytes, err
}
