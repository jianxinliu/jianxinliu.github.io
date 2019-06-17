# SQL

## DDL(Data Definition Language)

CREATE
### ALTER
```sql
-- add cloumn
alter table table_name add column_name data_type;

-- drop column
alter table table_name drop column column_name;

-- modify column
alter table modify column column_name datatype
```
DROP
TRUNCATE
COMMENT
RENAME

## DML(Data Manipulation Language)

SELECT
INSERT
UPDATE
DELETE
MERGE
CALL
EXPLAIN PLAN
LOCK TABLE

## DCL(Data Control Language)

GRANT 授权
REVOKE 取消授权

## TCL(Transaction Control Language)

SAVEPOINT 设置保存点
ROLLBACK  回滚
SET TRANSACTION

# MySQL Commands

```shell
mysql -uuname -ppwd

show databases;

use database_name;

show tables;

# 查看创建表的语句
show create table table_name;
```

