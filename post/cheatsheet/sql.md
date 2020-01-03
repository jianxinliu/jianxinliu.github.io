# SQL

## DDL(Data Definition Language)
`create`,`drop`,`alter`.


|操作对象|创建|删除|修改|
|:---:|:---:|:---:|:---:|
|模式|create schema|drop schema|-|
|表|create table|drop table|alter table|
|视图|create view|drop view|-|
|索引|create index|drop index|alter index|


### CREATE

```sql
-- 模式定义
create schema <schemaName> authorizzation <userName>
-- 表定义
create table <tableName> (<columnName> <dataType> [cloumn_level_constraint] [,<columnName> <dataType>] [,<table_level_constraint>]);
-- 索引定义
create [unqiue] [cluster] index <indexName> on <tableName>(<columnName> [,<rank>][,<columnName> [,<rank>]]...);
```

### DROP

```sql
-- 删除模式
drop schema <schemaName> <casade|restrict>;
-- 删除表
drop table <tableName> [restrict|casade];
-- 删除索引
drop index <indexName>
```

### ALTER

```sql
-- add cloumn
alter table <tableName> add <columnName> <dataType>;
-- drop column
alter table <tableName> drop <columnName> <dataType>;
-- modify column
alter table modify column <columnName> <dataType>;
-- 修改索引
alter index <oldIndexName> rename to <newIndexName>;
```

DROP，
TRUNCATE，
COMMENT，
RENAME

## DML(Data Manipulation Language)

SELECT，
INSERT，
UPDATE，
DELETE，
MERGE，
CALL，
EXPLAIN， PLAN，
LOCK TABLE

outer join,inner join,join

### 聚合函数
 
min,max,count,avg,sum

## DCL(Data Control Language)

GRANT 授权，
REVOKE 取消授权

## TCL(Transaction Control Language)

SAVEPOINT 设置保存点，
ROLLBACK  回滚，
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

