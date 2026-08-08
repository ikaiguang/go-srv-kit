package gormpkg

import (
	"context"
	"strings"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
)

type batchInsertTestRepo struct {
	table   string
	columns []string
	rows    int
}

func (s *batchInsertTestRepo) TableName() string { return s.table }
func (s *batchInsertTestRepo) Len() int          { return s.rows }
func (s *batchInsertTestRepo) InsertColumns() ([]string, string) {
	return s.columns, strings.TrimSuffix(strings.Repeat("?, ", len(s.columns)), ", ")
}
func (s *batchInsertTestRepo) InsertValues(args *BatchInsertValueArgs) ([]any, []string) {
	values := make([]any, 0, (args.StepEnd-args.StepStart)*len(s.columns))
	placeholders := make([]string, 0, args.StepEnd-args.StepStart)
	for i := args.StepStart; i < args.StepEnd; i++ {
		placeholders = append(placeholders, "("+args.InsertPlaceholder+")")
		for range s.columns {
			values = append(values, i)
		}
	}
	return values, placeholders
}

func TestBatchInsertValidation(t *testing.T) {
	db, err := gorm.Open(tests.DummyDialector{}, &gorm.Config{DryRun: true})
	if err != nil {
		t.Fatalf("gorm.Open() error = %v", err)
	}

	tests := []struct {
		name    string
		db      *gorm.DB
		repo    BatchInsertRepo
		wantErr string
	}{
		{name: "nil db", repo: &batchInsertTestRepo{}, wantErr: "db is nil"},
		{name: "nil repo", db: db, wantErr: "repo is nil"},
		{name: "invalid table", db: db, repo: &batchInsertTestRepo{table: "users; DROP TABLE users", columns: []string{"id"}, rows: 1}, wantErr: "table name is invalid"},
		{name: "empty columns", db: db, repo: &batchInsertTestRepo{table: "users", rows: 1}, wantErr: "columns are empty"},
		{name: "invalid column", db: db, repo: &batchInsertTestRepo{table: "users", columns: []string{"id) VALUES (1)"}, rows: 1}, wantErr: "column name is invalid"},
		{name: "valid", db: db, repo: &batchInsertTestRepo{table: "app.users", columns: []string{"id", "name"}, rows: 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := BatchInsertWithContext(context.Background(), tt.db, tt.repo, nil)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("BatchInsertWithContext() error = %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("BatchInsertWithContext() error = %v, want containing %q", err, tt.wantErr)
			}
		})
	}
}

func TestWithBatchInsertConflictActionNil(t *testing.T) {
	options := &batchInsertOptions{}
	WithBatchInsertConflictAction(nil)(options)
	if options.withConflictAction {
		t.Fatal("WithBatchInsertConflictAction(nil) enabled conflict action")
	}
}
