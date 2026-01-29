# 数据库迁移技能

## 描述
创建和管理数据库迁移脚本，支持 MySQL/PostgreSQL 数据库的版本控制和迁移。

## 参数
- `migration_name`: 迁移名称（如：add_user_table）
- `database_type`: 数据库类型（mysql/postgresql）
- `up_sql`: Up 迁移 SQL
- `down_sql`: Down 回滚 SQL

## 执行步骤
1. 在 `internal/data/schema/migrations/` 创建迁移文件
2. 命名格式：`{timestamp}_{migration_name}.up.sql` / `.down.sql`
3. 运行迁移工具
4. 验证迁移结果

## 迁移文件模板
```sql
-- {timestamp}_{migration_name}.up.sql
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(36) NOT NULL,
  `username` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uuid` (`uuid`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- {timestamp}_{migration_name}.down.sql
DROP TABLE IF EXISTS `users`;
```
