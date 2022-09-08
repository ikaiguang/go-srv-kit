package dbv1_0_0_example

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./example/cmd/migration/v1.0.0/admin -test.run=TestMigrator_CreateTableExample
func TestMigrator_CreateTableExample(t *testing.T) {
	err := upHandler.CreateTableExample()
	require.Nil(t, err)
}
