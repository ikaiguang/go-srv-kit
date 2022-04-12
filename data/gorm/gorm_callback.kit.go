package gormutil

import (
	"gorm.io/gorm"
)

// RemoveCallback 移除回调
// 参考文档: https://gorm.io/docs/write_plugins.html#Callbacks
// 参考func: callbacks.RegisterDefaultCallbacks
func RemoveCallback(db *gorm.DB) error {
	//return db.Callback().Create().Remove("gorm:create")
	return nil
}
