package curlpkg

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

// RequestParam 请求参数
type RequestParam struct {
	// Timeout 请求超时；默认 DefaultTimeout = 1分钟
	Timeout time.Duration
}

// Do 执行请求
func Do(httpReq *http.Request, requestParam *RequestParam) (httpCode int, bodyBytes []byte, err error) {
	// client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	defer httpClient.CloseIdleConnections()

	// 请求
	if requestParam.Timeout < 1 {
		requestParam.Timeout = DefaultTimeout
	}
	httpClient.Timeout = requestParam.Timeout
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return
	}
	defer func() { _ = httpResp.Body.Close() }()

	// resp
	httpCode = httpResp.StatusCode
	bodyBytes, err = ioutil.ReadAll(httpResp.Body)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return
	}
	return
}
