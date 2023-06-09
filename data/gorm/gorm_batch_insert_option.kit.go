package gormpkg

// BatchInsertOption 批量插入选项
type BatchInsertOption func(*batchInsertOptions)

// batchInsertOptions 批量插入选项
type batchInsertOptions struct {
	// isInsertIgnore 是否忽略重复插入
	// MySQL : INSERT IGNORE INTO ...
	// Postgres : ON CONFLICT ...
	isInsertIgnore bool

	// withConflictAction 是否执行冲突解决
	// MySQL : INSERT INTO ... VALUES (...) AS alias ON DUPLICATE KEY UPDATE a = alias.a
	// Postgres : ON CONFLICT(id) DO UPDATE SET column_2= CONCAT(test_table.column_2, excluded.column_2);
	withConflictAction bool
	// onConflictValueAlias 值的别名
	// MySQL : INSERT INTO ... VALUES (...) AS alias ON DUPLICATE KEY UPDATE a = alias.a
	// Postgres : 默认：excluded
	onConflictValueAlias string
	// onConflictTarget 条件
	// MySQL : ON DUPLICATE KEY
	// Postgres : ON CONFLICT (targetColumn, ...)
	onConflictTarget string
	// onConflictAction 执行冲突解决方案
	// MySQL : DO NOTHING
	onConflictAction string
	// onConflictPrepareData 冲突解决方案数据；DO UPDATE SET column_2 = ?
	onConflictPrepareData []interface{}
}

// BatchInsertConflictActionReq 批量插入冲突解决请求
type BatchInsertConflictActionReq struct {
	// OnConflictValueAlias 值的别名
	// MySQL : INSERT INTO ... VALUES (...) AS alias ON DUPLICATE KEY UPDATE a = alias.a
	// Postgres : 默认：excluded
	OnConflictValueAlias string
	// OnConflictTarget 条件
	// MySQL : ON DUPLICATE KEY
	// Postgres : ON CONFLICT (targetColumn, ...)
	OnConflictTarget string
	// OnConflictAction 执行冲突解决方案
	// MySQL : DO NOTHING
	OnConflictAction string
	// OnConflictPrepareData 冲突解决方案数据；DO UPDATE SET column_2 = ?
	OnConflictPrepareData []interface{}
}

const (
	DefaultBatchInsertConflictAlias = "excluded"
)

var (
	DefaultBatchInsertConflictActionForMySQL = BatchInsertConflictActionReq{
		OnConflictValueAlias:  "AS " + DefaultBatchInsertConflictAlias,
		OnConflictTarget:      "ON DUPLICATE KEY",
		OnConflictAction:      "UPDATE", // UPDATE column_a= excluded.column_a
		OnConflictPrepareData: nil,
	}
	DefaultBatchInsertConflictActionPostgres = BatchInsertConflictActionReq{
		OnConflictValueAlias:  "",
		OnConflictTarget:      "ON CONFLICT",   //  ON CONFLICT(id)
		OnConflictAction:      "DO UPDATE SET", // DO UPDATE SET column_2 = CONCAT(test_table.column_2, excluded.column_2)
		OnConflictPrepareData: nil,
	}
)

// WithBatchInsertIgnore 忽略重复插入
// INSERT IGNORE INTO ...
func WithBatchInsertIgnore() BatchInsertOption {
	return func(options *batchInsertOptions) {
		options.isInsertIgnore = true
	}
}

// WithBatchInsertConflictAction 执行冲突解决
// MySQL : INSERT INTO ... VALUES (...) AS alias ON DUPLICATE KEY UPDATE a = alias.a
// Postgres : ON CONFLICT(id) DO UPDATE SET column_2= CONCAT(test_table.column_2, excluded.column_2);
func WithBatchInsertConflictAction(req *BatchInsertConflictActionReq) BatchInsertOption {
	return func(options *batchInsertOptions) {
		options.withConflictAction = true
		options.onConflictValueAlias = req.OnConflictValueAlias
		options.onConflictTarget = req.OnConflictTarget
		options.onConflictAction = req.OnConflictAction
		options.onConflictPrepareData = req.OnConflictPrepareData
	}
}
