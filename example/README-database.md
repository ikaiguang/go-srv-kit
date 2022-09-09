# 数据库

```shell

# mysql
CREATE DATABASE srv_example DEFAULT CHARSET utf8mb4;
CREATE USER 'srv_user'@'%' IDENTIFIED BY 'Mysql.123456';
GRANT ALL PRIVILEGES ON srv_example.* TO "srv_user"@"%";
FLUSH PRIVILEGES;

# psql
CREATE DATABASE ligai_crm;
CREATE USER srv_user WITH PASSWORD 'Postgres.123456';
GRANT ALL PRIVILEGES ON DATABASE srv_example TO srv_user;

```

## 数据库迁移

有一下两种方式

1. 运行单独的程序，生成数据表
2. 在主程序配置中启用自动迁移；配置位置：`server.setting.enable_migrate_db`

```shell

# 例子
# go run ./cmd/migration/... -conf=./configs
# 运行 admin
# make migrate app=admin
go run ./app/admin/service/cmd/migration/... -conf=./app/admin/service/configs
# 运行 user
# make migrate app=user
go run ./app/user/service/cmd/migration/... -conf=./app/user/service/configs

```
