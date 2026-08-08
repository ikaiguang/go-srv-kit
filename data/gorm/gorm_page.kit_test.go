package gormpkg

import (
	"strings"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
)

type queryTestModel struct {
	ID        uint64
	IsDeleted bool
}

func TestQueryHelpersGenerateSafeSQL(t *testing.T) {
	db, err := gorm.Open(tests.DummyDialector{}, &gorm.Config{DryRun: true})
	if err != nil {
		t.Fatalf("gorm.Open() error = %v", err)
	}

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx = QueryUndeletedData(tx.Model(&queryTestModel{}))
		tx = AssembleWheres(tx, []*Where{
			nil,
			{Field: "id", Operator: "= ?; DROP TABLE users", Placeholder: "? OR 1=1", Value: 7},
		})
		tx = AssembleOrders(tx, []*Order{nil, NewOrder("id", "asc; DROP TABLE users")})
		return tx.Find(&[]queryTestModel{})
	})

	if strings.Contains(strings.ToUpper(sql), "DROP TABLE") || strings.Contains(sql, "OR 1=1") {
		t.Fatalf("query contains unsafe SQL: %s", sql)
	}
	if !strings.Contains(sql, "is_deleted = 0") {
		t.Fatalf("query does not contain undeleted predicate: %s", sql)
	}
	if !strings.Contains(sql, "ORDER BY id desc") {
		t.Fatalf("query does not sanitize order direction: %s", sql)
	}
}

func TestPaginatorNilOption(t *testing.T) {
	if got := Paginator(nil, nil); got != nil {
		t.Fatalf("Paginator(nil, nil) = %v, want nil", got)
	}
}
