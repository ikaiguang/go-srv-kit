package gormpkg

import (
	"context"

	"gorm.io/gorm"
)

const (
	DefaultPlaceholder     = "?" // param placeholder
	invalidWhereColumnName = "bad_where_from_invalid_column"
)

// Where 条件；例：where id = ?(where id = 1)
type Where struct {
	// Field 字段
	Field string
	// Operator 运算符
	Operator string
	// Placeholder 占位符
	Placeholder string
	// Value 数据
	Value any
}

// NewWhere where
func NewWhere(field, operator string, value any) *Where {
	return &Where{
		Field:       field,
		Operator:    operator,
		Placeholder: DefaultPlaceholder,
		Value:       value,
	}
}

// AssembleWheres 组装条件
func AssembleWheres(db *gorm.DB, wheres []*Where) *gorm.DB {
	if len(wheres) == 0 {
		return db
	}
	for i := range wheres {
		column := wheres[i].Field
		if !IsValidColumnName(column) {
			column = invalidWhereColumnName
			if db.Logger != nil {
				db.Logger.Error(context.Background(), "invalid column(", wheres[i].Field, ")")
			}
		}
		db = db.Where(column+" "+wheres[i].Operator+" "+wheres[i].Placeholder, wheres[i].Value)
	}
	return db
}

// UnsafeAssembleWheres 不安全的组装条件
// WARNING: 此函数不验证列名，可能导致 SQL 注入。仅在确认输入安全时使用。
func UnsafeAssembleWheres(db *gorm.DB, wheres []*Where) *gorm.DB {
	if len(wheres) == 0 {
		return db
	}
	for i := range wheres {
		db = db.Where(wheres[i].Field+" "+wheres[i].Operator+" "+wheres[i].Placeholder, wheres[i].Value)
	}
	return db
}
