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