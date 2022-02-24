package curlutil

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	pkgerrors "github.com/pkg/errors"
)

const (
	// header
	_contentType = "Content-Type"
	_userAgent   = "User-Agent"

	// timeout
	DefaultTimeout = time.Minute

	// content type
	ContentTypeJSON     = "application/json"
	ContentTypeJSONUtf8 = "application/json; charset=utf-8"
	ContentTypePB       = "application/x-protobuf"
)

// NewPostRequest Post请求
func NewPostRequest(httpURL string, body io.Reader) (req *http.Request, err error) {
	// http
	httpReq, err := http.NewRequest(http.MethodPost, httpURL, body)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return
	}
	return httpReq, err
}

// NewGetRequest Get请求
func NewGetRequest(httpURL string, body io.Reader) (req *http.Request, err error) {
	// http
	httpReq, err := http.NewRequest(http.MethodGet, httpURL, body)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return
	}
	return httpReq, err
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}
	httpClient.Timeout = options.timeout

	return httpClient
}

// Do 请求
func Do(httpReq *http.Request, opts ...Option) (httpCode int, bodyBytes []byte, err error) {
	httpClient := NewHTTPClient(opts...)
	defer httpClient.CloseIdleConnections()

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return httpCode, bodyBytes, err
	}
	defer func() { _ = httpResp.Body.Close() }()

	return response(httpResp)
}

// Default http.DefaultClient
func Default(httpReq *http.Request) (httpCode int, bodyBytes []byte, err error) {
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		err = pkgerrors.WithStack(err)
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
		err = pkgerrors.WithStack(err)
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
	bodyBytes, err = ioutil.ReadAll(httpResp.Body)
	if err != nil {
		if err == io.EOF {
			err = nil
		} else {
			err = pkgerrors.WithStack(err)
		}
		return httpCode, bodyBytes, err
	}
	return httpCode, bodyBytes, err
}
