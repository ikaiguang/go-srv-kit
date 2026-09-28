package registrypkg

import "testing"

func TestRegistryConstructorsRejectNilClient(t *testing.T) {
	constructors := map[string]func() (bool, error){
		"NewRegistry": func() (bool, error) {
			registry, err := NewRegistry(nil)
			return registry == nil, err
		},
		"NewConsulRegistry": func() (bool, error) {
			registry, err := NewConsulRegistry(nil)
			return registry == nil, err
		},
	}
	for name, constructor := range constructors {
		t.Run(name, func(t *testing.T) {
			isNil, err := constructor()
			if err == nil {
				t.Fatalf("%s(nil) error = nil, want error", name)
			}
			if !isNil {
				t.Fatalf("%s(nil) registry is non-nil, want nil", name)
			}
		})
	}
}
