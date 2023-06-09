package curlpkg

import (
	"crypto/tls"
	"io"
	"io/ioutil"
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
	// http
	httpReq, err = http.NewRequest(http.MethodPost, httpURL, body)
	if err != nil {
		return
	}
	return httpReq, err
}

// NewGetRequest Get请求
func NewGetRequest(httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	// http
	httpReq, err = http.NewRequest(http.MethodGet, httpURL, body)
	if err != nil {
		return
	}
	return httpReq, err
}

// NewRequest .
func NewRequest(httpMethod, httpURL string, body io.Reader) (httpReq *http.Request, err error) {
	// http
	httpReq, err = http.NewRequest(httpMethod, httpURL, body)
	if err != nil {
		return httpReq, err
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
	bodyBytes, err = ioutil.ReadAll(httpResp.Body)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return httpCode, bodyBytes, err
	}
	return httpCode, bodyBytes, err
}
