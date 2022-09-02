package jaegerutil

import (
	"fmt"
	"go.opentelemetry.io/otel/exporters/jaeger"
	stdhttp "net/http"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

// NewJaegerExporter ...
func NewJaegerExporter(conf *confv1.Data_JaegerTracer, opts ...Option) (*jaeger.Exporter, error) {
	return NewExporter(conf, opts...)
}

// NewExporter jaeger.Exporter
func NewExporter(conf *confv1.Data_JaegerTracer, opts ...Option) (*jaeger.Exporter, error) {
	var jaegerOptions []jaeger.CollectorEndpointOption
	if conf.Endpoint != "" {
		jaegerOptions = append(jaegerOptions, jaeger.WithEndpoint(conf.Endpoint))
	}
	if conf.WithHttpBasicAuth {
		jaegerOptions = append(jaegerOptions, jaeger.WithUsername(conf.Username))
		jaegerOptions = append(jaegerOptions, jaeger.WithPassword(conf.Password))
	}

	isValidConnection, err := checkConnection(conf)
	if err != nil {
		err = fmt.Errorf("check connection error : %w", err)
		return nil, err
	}
	if !isValidConnection {
		err = fmt.Errorf("invalid jaeger endpoint")
		return nil, err
	}

	return jaeger.New(jaeger.WithCollectorEndpoint(jaegerOptions...))
}

// checkConnection 检查链接可用性
func checkConnection(conf *confv1.Data_JaegerTracer) (isValid bool, err error) {
	httpClient := stdhttp.Client{}
	defer httpClient.CloseIdleConnections()
	httpRequest, err := stdhttp.NewRequest(stdhttp.MethodPost, conf.Endpoint, nil)
	if err != nil {
		return isValid, err
	}
	if conf.WithHttpBasicAuth {
		httpRequest.SetBasicAuth(conf.Username, conf.Password)
	}
	httpResp, err := httpClient.Do(httpRequest)
	if err != nil {
		return isValid, err
	}
	defer func() { _ = httpResp.Body.Close() }()

	// 有效的链接
	isValid = true
	//httpBodyBytes, err := ioutil.ReadAll(httpResp.Body)
	//if err != nil {
	//	return isValid, err
	//}
	//_ = httpResp.StatusCode
	//_ = httpBodyBytes
	return isValid, err
}
