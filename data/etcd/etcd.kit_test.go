package etcdpkg

import "testing"

func TestBuildClientConfigRejectsNilConfig(t *testing.T) {
	config, err := buildClientConfig(nil)
	if err == nil {
		t.Fatal("buildClientConfig(nil) error = nil, want error")
	}
	if config != nil {
		t.Fatalf("buildClientConfig(nil) config = %v, want nil", config)
	}
}

func TestBuildClientConfigRejectsInvalidCAPEM(t *testing.T) {
	config, err := buildClientConfig(&Config{CaCert: []byte("not a PEM certificate")})
	if err == nil {
		t.Fatal("buildClientConfig(invalid CA) error = nil, want error")
	}
	if config != nil {
		t.Fatalf("buildClientConfig(invalid CA) config = %v, want nil", config)
	}
}

func TestBuildClientConfigAppliesInsecureSkipVerifyWithoutCustomCA(t *testing.T) {
	config, err := buildClientConfig(&Config{InsecureSkipVerify: true})
	if err != nil {
		t.Fatalf("buildClientConfig() error = %v", err)
	}
	if config.TLS == nil {
		t.Fatal("buildClientConfig() TLS = nil, want TLS config")
	}
	if !config.TLS.InsecureSkipVerify {
		t.Fatal("buildClientConfig() InsecureSkipVerify = false, want true")
	}
}

func TestNewClientRejectsNilConfig(t *testing.T) {
	client, err := NewClient(nil)
	if err == nil {
		t.Fatal("NewClient(nil) error = nil, want error")
	}
	if client != nil {
		t.Fatalf("NewClient(nil) client = %v, want nil", client)
	}
}
