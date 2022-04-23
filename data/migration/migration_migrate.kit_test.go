package migrationuitl

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// go test -v -count=1 ./data/migration -test.run=TestMigrateALL
func TestMigrateALL(t *testing.T) {
	var (
		testMg   = NewCreateTable(dbConn.Migrator(), &TestMigration{})
		normalMg = NewCreateTable(dbConn.Migrator(), &Migration{})
	)
	var args = []struct {
		name string
		repo MigrationRepo
	}{
		{
			name: "#创建表：" + testMg.Identifier(),
			repo: testMg,
		},
		{
			name: "#创建表：" + normalMg.Identifier(),
			repo: normalMg,
		},
	}

	for i := range args {
		MustRegisterMigrate(args[i].repo)
	}

	// 迁移
	err := MigrateALL()
	require.Nil(t, err)
	err = Migrate(testMg.Identifier(), normalMg.Identifier())
	require.Nil(t, err)
	err = MigrateRepos(testMg, normalMg)
	require.Nil(t, err)
	// 回滚
	err = RollbackALL()
	require.Nil(t, err)
	err = Rollback(testMg.Identifier(), normalMg.Identifier())
	require.Nil(t, err)
	err = RollbackRepos(testMg, normalMg)
	require.Nil(t, err)
}
