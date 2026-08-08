package registrypkg

import "testing"

func TestNewEtcdRegistryRejectsNilClient(t *testing.T) {
	registry, err := NewEtcdRegistry(nil)
	if err == nil {
		t.Fatal("NewEtcdRegistry(nil) error = nil, want error")
	}
	if registry != nil {
		t.Fatalf("NewEtcdRegistry(nil) registry = %v, want nil", registry)
	}
}
