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

### join

#### `inner join` VS `cross join`

`inner join` 效果等同于使用 `from`多表查询，等同于 `cross join`，等同于 `join` , 即：

```sql
--- 效果同，得两张表的笛卡尔积
select * from user inner join student;
select * from user,student;
select * from user cross join student;
select * from user join student;
```

#### `left(right) join` == `left(right) outer join`

效果相同。

#### `left join` VS `right join`

```sql
--- 以 T1 的行数为准，rows(T2) < rows(T1),则留空
select * from T1 left join T2;

--- 以 T2 的行数为准，rows（T1）< rows(T2),则留空
select * from T1 right join T2;
```

JOIN 按照功能可分为如下三类：

- `INNER JOIN`（内连接，或等值连接）：获取两个表中字段匹配关系的记录；
- `LEFT JOIN`（左连接）：获取左表中的所有记录，即使在右表没有对应匹配的记录；
- `RIGHT JOIN`（右连接）：与 LEFT JOIN 相反，用于获取右表中的所有记录，即使左表没有对应匹配的记录。

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

