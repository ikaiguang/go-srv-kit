package jaegerutil

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	stdhttp "net/http"
	"testing"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

// go test -v ./data/jaeger/ -count=1 -test.run=TestNewJaegerExporter_Xxx
// jaeger.kit_test.go:51: ==> httpResp StatusCode : 400
// jaeger.kit_test.go:55: ==> httpResp Body : Cannot parse content type: mime: no media type
func TestNewJaegerExporter_Xxx(t *testing.T) {
	var tests = []struct {
		name string
		conf *confv1.Base_JaegerTracer
	}{
		{
			name: "#WithHttpBasicAuth:NO",
			conf: &confv1.Base_JaegerTracer{
				Endpoint: "http://127.0.0.1:14268/api/traces",
			},
		},
		{
			name: "#WithHttpBasicAuth:YES",
			conf: &confv1.Base_JaegerTracer{
				Endpoint:          "http://127.0.0.1:14268/api/traces",
				WithHttpBasicAuth: true,
				Username:          "ikaiguang",
				Password:          "123456",
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			httpClient := stdhttp.Client{}
			httpURL := v.conf.Endpoint
			httpRequest, err := stdhttp.NewRequest(stdhttp.MethodPost, httpURL, nil)
			require.Nil(t, err)
			if v.conf.WithHttpBasicAuth {
				httpRequest.SetBasicAuth(v.conf.Username, v.conf.Password)
			}
			httpResp, err := httpClient.Do(httpRequest)
			if err != nil {
				t.Error("==> httpClient.Do(httpRequest) error :", err.Error())
			}
			require.Nil(t, err)
			t.Log("==> httpResp StatusCode :", httpResp.StatusCode)
			defer func() { _ = httpResp.Body.Close() }()
			httpBodyBytes, err := ioutil.ReadAll(httpResp.Body)
			require.Nil(t, err)
			t.Log("==> httpResp Body :", string(httpBodyBytes))

		})
	}
}
