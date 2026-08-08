package registrypkg

import "testing"

func TestNewConsulRegistryRejectsNilClient(t *testing.T) {
	registry, err := NewConsulRegistry(nil)
	if err == nil {
		t.Fatal("NewConsulRegistry(nil) error = nil, want error")
	}
	if registry != nil {
		t.Fatalf("NewConsulRegistry(nil) registry = %v, want nil", registry)
	}
}
