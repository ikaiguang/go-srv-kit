package gormpkg

import (
	"context"
	"strings"

	"gorm.io/gorm"
)

const (
	DefaultOrderColumn = "id"   // default order column
	DefaultOrderAsc    = "asc"  // order direction : asc
	DefaultOrderDesc   = "desc" // order direction : desc
)

// ParseOrderDirection 排序方向
func ParseOrderDirection(orderDirection string) string {
	if orderDirection = strings.ToLower(orderDirection); orderDirection == DefaultOrderAsc {
		return DefaultOrderAsc
	}
	return DefaultOrderDesc
}

// Order 排序(例子：order by id desc)
type Order struct {
	// Field 排序的字段(例子：id)
	Field string
	// Order 排序的方向(例子：desc)
	Order string
}

// NewOrder order
func NewOrder(field, orderDirection string) *Order {
	return &Order{
		Field: field,
		Order: orderDirection,
	}
}

// AssembleOrders 组装排序
func AssembleOrders(db *gorm.DB, orders []*Order) *gorm.DB {
	if len(orders) == 0 {
		return db
	}

	for i := range orders {
		column := orders[i].Field
		if !IsValidField(column) {
			//column = DefaultOrderColumn
			column = "bad_order_from_invalid_column"
			if db.Logger != nil {
				db.Logger.Error(context.Background(), "invalid column(", orders[i].Field, ")")
			}
		}
		db = db.Order(column + " " + ParseOrderDirection(orders[i].Order))
	}
	return db
}

// UnsafeAssembleOrders 不安全的组装排序
func UnsafeAssembleOrders(db *gorm.DB, orders []*Order) *gorm.DB {
	if len(orders) == 0 {
		return db
	}

	for i := range orders {
		db = db.Order(orders[i].Field + " " + orders[i].Order)
	}
	return db
}
