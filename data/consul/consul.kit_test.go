package consulpkg

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestNewClientRejectsNilConfig(t *testing.T) {
	client, err := NewClient(nil)
	if err == nil {
		t.Fatal("NewClient(nil) error = nil, want error")
	}
	if client != nil {
		t.Fatalf("NewClient(nil) client = %v, want nil", client)
	}
}

func TestBuildConfigAppliesWriter(t *testing.T) {
	t.Setenv(api.HTTPSSLEnvName, "not-a-bool")
	var output bytes.Buffer

	_, err := buildConfig(&Config{}, WithWriter(&output))
	if err != nil {
		t.Fatalf("buildConfig() error = %v", err)
	}
	if !strings.Contains(output.String(), api.HTTPSSLEnvName) {
		t.Fatalf("logger output = %q, want %s warning", output.String(), api.HTTPSSLEnvName)
	}
}

func TestNewClientUsesReadOnlyProbe(t *testing.T) {
	var method string
	config := api.DefaultConfig()
	config.Address = "consul.test"
	config.HttpClient = &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		method = r.Method
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("")),
			Request:    r,
		}, nil
	})}
	client, err := api.NewClient(config)
	if err != nil {
		t.Fatalf("api.NewClient() error = %v", err)
	}
	if err = probeClient(client); err != nil {
		t.Fatalf("probeClient() error = %v", err)
	}
	if method != http.MethodGet {
		t.Fatalf("probe method = %q, want GET", method)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}
