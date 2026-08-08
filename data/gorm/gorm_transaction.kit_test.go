package gormpkg

import (
	"context"
	"errors"
	"testing"

	"gorm.io/gorm"
)

func TestCommitAndErrRollbackPreservesResultError(t *testing.T) {
	resultErr := errors.New("business operation failed")
	rollbackErr := errors.New("rollback failed")
	tx := &transaction{err: rollbackErr}

	err := tx.CommitAndErrRollback(context.Background(), resultErr)
	if !errors.Is(err, resultErr) {
		t.Fatalf("CommitAndErrRollback() error = %v, want original result error", err)
	}
	if !errors.Is(err, rollbackErr) {
		t.Fatalf("CommitAndErrRollback() error = %v, want rollback error", err)
	}
}

func TestNewTransactionRejectsNilArguments(t *testing.T) {
	if err := NewTransaction(context.Background(), nil).Commit(context.Background()); err == nil {
		t.Fatal("NewTransaction() nil DB error = nil")
	}
	if err := NewTransaction(nil, &gorm.DB{}).Rollback(context.Background()); err == nil {
		t.Fatal("NewTransaction() nil context error = nil")
	}
	if err := (&transaction{tx: &gorm.DB{}}).Commit(nil); err == nil {
		t.Fatal("Commit(nil) error = nil")
	}
}
