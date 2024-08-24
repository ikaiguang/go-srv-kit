package gormpkg

import (
	stderrors "errors"
	"time"

	"gorm.io/gorm"
)

const (
	FieldID          = "id"
	FieldCreatedTime = "created_time"
	FieldUpdatedTime = "updated_time"
	FieldDeletedTime = "deleted_time"
	FieldIsDeleted   = "is_deleted"
)

var (
	_ = gorm.Model{}
)

type Model struct {
	Id          uint64    `gorm:"column:id;type:uint;autoIncrement;comment:ID" json:"id"`
	CreatedTime time.Time `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:time;not null;comment:更新时间" json:"updated_time"`
	IsDeleted   bool      `gorm:"column:is_deleted;type:uint;default:0;comment:是否已删除" json:"is_deleted"`
	DeletedTime time.Time `gorm:"column:deleted_time;type:time;comment:删除时间" json:"deleted_time"`
}

type ModelForMysql struct {
	Id          uint64    `gorm:"column:id;type:uint;autoIncrement;default:current_timestamp();comment:ID" json:"id"`
	CreatedTime time.Time `gorm:"column:created_time;type:time;not null;default:current_timestamp();comment:创建时间" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:time;not null;default:current_timestamp();comment:更新时间" json:"updated_time"`
	IsDeleted   bool      `gorm:"column:is_deleted;type:uint;default:0;comment:是否已删除" json:"is_deleted"`
	DeletedTime time.Time `gorm:"column:deleted_time;type:time;comment:删除时间" json:"deleted_time"`
}

type ModelForPostgres struct {
	Id          uint64    `gorm:"column:id;type:uint;autoIncrement;default:current_timestamp;comment:ID" json:"id"`
	CreatedTime time.Time `gorm:"column:created_time;type:time;not null;default:current_timestamp;comment:创建时间" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:time;not null;default:current_timestamp;comment:更新时间" json:"updated_time"`
	IsDeleted   bool      `gorm:"column:is_deleted;type:uint;default:0;comment:是否已删除" json:"is_deleted"`
	DeletedTime time.Time `gorm:"column:deleted_time;type:time;comment:删除时间" json:"deleted_time"`
}

// IsErrRecordNotFound ...
func IsErrRecordNotFound(err error) bool {
	return stderrors.Is(err, gorm.ErrRecordNotFound)
}

// QueryUndeletedData 未删除的数据
func QueryUndeletedData(dbConn *gorm.DB) *gorm.DB {
	return dbConn.Where(FieldIsDeleted, 0)
}

// QueryDeletedData 删除的数据
func QueryDeletedData(dbConn *gorm.DB) *gorm.DB {
	return dbConn.Where(FieldIsDeleted, 1)
}

func SetUpdateTime(updates map[string]interface{}) {
	updates[FieldUpdatedTime] = time.Now()
}

func SetCreateTime(updates map[string]interface{}) {
	updates[FieldCreatedTime] = time.Now()
}

func SoftDelete(dbConn *gorm.DB) *gorm.DB {
	return dbConn.UpdateColumns(map[string]interface{}{
		FieldIsDeleted:   1,
		FieldDeletedTime: time.Now(),
	})
}

func Deleted(dbConn *gorm.DB, value interface{}, conditions ...interface{}) *gorm.DB {
	return dbConn.Delete(value, conditions)
}
