package gormutil

import (
	"context"

	"gorm.io/gorm"
)

const (
	DefaultPlaceholder = "?" // param placeholder
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
	Value interface{}
}

// NewWhere where
func NewWhere(field, operator string, value interface{}) *Where {
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
		if !IsValidField(column) {
			//column = DefaultOrderColumn
			column = "bad_where_from_invalid_column"
			if db.Logger != nil {
				db.Logger.Error(context.Background(), "invalid column(", wheres[i].Field, ")")
			}
		}
		db = db.Where(column+" "+wheres[i].Operator+" "+wheres[i].Placeholder, wheres[i].Value)
	}
	return db
}

// UnsafeAssembleWheres 不安全的组装条件
func UnsafeAssembleWheres(db *gorm.DB, wheres []*Where) *gorm.DB {
	if len(wheres) == 0 {
		return db
	}
	for i := range wheres {
		db = db.Where(wheres[i].Field+" "+wheres[i].Operator+" "+wheres[i].Placeholder, wheres[i].Value)
	}
	return db
}
