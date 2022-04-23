package migrationuitl

import (
	"fmt"
	"sync"
)

var (
	migrations = &migrationRegister{}
)

// migrationRegister 注册迁移
type migrationRegister struct {
	sync.Mutex
	m   map[string]MigrationRepo
	ids []string
}

// MustRegisterMigrate 注册迁移
func MustRegisterMigrate(repo MigrationRepo) {
	migrations.Lock()
	defer migrations.Unlock()

	// 已注册
	if _, ok := migrations.m[repo.Identifier()]; ok {
		panic("migration already registered: " + repo.Identifier())
	}
	migrations.m[repo.Identifier()] = repo
	migrations.ids = append(migrations.ids, repo.Identifier())
}

// migrateUpError 迁移失败错误
func migrateUpError(identifier string, err error) error {
	return fmt.Errorf("migration up failed: identifier=%s \n\t error: %s", identifier, err)
}

// migrateDownError 迁移失败错误
func migrateDownError(identifier string, err error) error {
	return fmt.Errorf("migration down failed: identifier=%s \n\t error: %s", identifier, err)
}

// migrateNotFoundError 迁移失败错误
func migrateNotFoundError(identifier string) error {
	return fmt.Errorf("migration not found: identifier=%s", identifier)
}

// MigrateALL 执行迁移
func MigrateALL() error {
	migrations.Lock()
	defer migrations.Unlock()

	for i := range migrations.ids {
		repo := migrations.m[migrations.ids[i]]
		if err := repo.Up(); err != nil {
			return migrateUpError(repo.Identifier(), err)
		}
	}
	return nil
}

// Migrate 执行迁移
func Migrate(ids ...string) error {
	migrations.Lock()
	defer migrations.Unlock()

	for i := range ids {
		_, ok := migrations.m[ids[i]]
		if !ok {
			return migrateNotFoundError(ids[i])
		}
	}
	for i := range ids {
		repo := migrations.m[ids[i]]
		if err := repo.Up(); err != nil {
			return migrateDownError(repo.Identifier(), err)
		}
	}
	return nil
}

// RollbackALL 回滚迁移
func RollbackALL() error {
	migrations.Lock()
	defer migrations.Unlock()

	for i := range migrations.ids {
		repo := migrations.m[migrations.ids[i]]
		if err := repo.Down(); err != nil {
			return migrateDownError(repo.Identifier(), err)
		}
	}
	return nil
}

// Rollback 回滚迁移
func Rollback(ids ...string) error {
	migrations.Lock()
	defer migrations.Unlock()

	for i := range ids {
		_, ok := migrations.m[ids[i]]
		if !ok {
			return migrateNotFoundError(ids[i])
		}
	}
	for i := range ids {
		repo := migrations.m[ids[i]]
		if err := repo.Down(); err != nil {
			return migrateUpError(repo.Identifier(), err)
		}
	}
	return nil
}

// MigrateRepos 执行迁移
func MigrateRepos(repos ...MigrationRepo) error {
	for i := range repos {
		if err := repos[i].Up(); err != nil {
			return migrateUpError(repos[i].Identifier(), err)
		}
	}
	return nil
}

// RollbackRepos 执行回滚
func RollbackRepos(repos ...MigrationRepo) error {
	for i := range repos {
		if err := repos[i].Down(); err != nil {
			return migrateDownError(repos[i].Identifier(), err)
		}
	}
	return nil
}
