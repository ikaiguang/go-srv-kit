package dbv1_0_0

import (
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	dbv1_0_0_example "github.com/ikaiguang/go-srv-kit/example/cmd/migration/v1.0.0/example"
)

// Upgrade update
func Upgrade(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) (err error) {
	//var (
	//	serverVersion     = "v1.0.0"
	//	migrateIdentifier = serverVersion + ":migrate"
	//)
	//// 已进行数据库迁移
	//dataModel, _, err := migrateRepo.QueryOneByIdentifier(migrateIdentifier)
	//if err != nil {
	//	return
	//}
	//if dataModel.Id > 0 {
	//	return
	//}
	//// 记录数据库迁移
	//defer func() {
	//	if err == nil {
	//		err = migrateRepo.CreateDefaultRecord(serverVersion, migrateIdentifier)
	//	}
	//}()

	// admin
	err = dbv1_0_0_example.Upgrade(dbConn, migrateRepo)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	return err
}
