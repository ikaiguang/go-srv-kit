package testping

import (
	"testing"

	"github.com/go-resty/resty/v2"
	testdata "github.com/ikaiguang/go-srv-kit/example/integration-test/testdata"
	curlutil "github.com/ikaiguang/go-srv-kit/kit/curl"
	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./example/integration-test/ping -test.run=TestHTTP_Ping_Hello
func TestHTTP_Ping_Hello(t *testing.T) {
	urlPath := "/api/v1/ping/hello"
	fullURL := testdata.GenURL(urlPath)

	restyClient := resty.New()
	if testdata.EnableTrace() {
		restyClient = restyClient.SetDebug(true).EnableTrace()
	}

	// 请求
	restyResponse, err := restyClient.R().Get(fullURL)
	require.Nil(t, err)
	defer func() {
		if restyResponse.RawResponse != nil {
			_ = restyResponse.RawResponse.Body.Close()
		}
	}()

	if ok := curlutil.IsSuccessCode(restyResponse.StatusCode()); !ok {
		t.Logf("http response code : %d\n", restyResponse.StatusCode())
		t.Logf("http response body content : %v\n", string(restyResponse.Body()))
		t.FailNow()
	}

	t.Logf("http response body lenght : %d\n", len(restyResponse.Body()))
	if testdata.EnablePrintResult() {
		t.Logf("http response body content : %v\n", string(restyResponse.Body()))
	}
}
