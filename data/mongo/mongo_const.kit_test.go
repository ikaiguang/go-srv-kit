package mongopkg

import "testing"

func TestOperationConstants(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{name: "query", got: OperationElemMatch, want: "$elemMatch"},
		{name: "update", got: OperationSetOnInsert, want: "$setOnInsert"},
		{name: "pipeline", got: OperationSetWindowFields, want: "$setWindowFields"},
		{name: "expression", got: OperationToString, want: "$toString"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("operation = %q, want %q", tt.got, tt.want)
			}
		})
	}

	if OperationNotIn != OperationNin {
		t.Fatalf("OperationNotIn = %q, want alias of OperationNin %q", OperationNotIn, OperationNin)
	}
}
