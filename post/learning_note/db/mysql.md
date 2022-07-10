# Mysql 基本篇

## 数据库操作

```sql
-- 显示所有数据库
SHOW databases;
-- 创建数据库
CREATE DATABASE name;
-- 改变当前数据库(选择数据库)
USE database_name;
-- 删除数据库
DROP DATABASE name;
```

### 存储引擎

存储引擎就是指表的类型。为了提高MySQL DBMS的使用效率和灵活性，可以根据实际需要选择存储引擎。存储引擎决定了表的类型，即如何存储和索引数据、是否支持事务等，同时，存储引擎也决定了表的物理存储方式。

```sql
-- MySQL 支持的引擎
SHOW ENGINES;
```

可以看到，MySQL支持 9 种引擎，其中 InnoDB 是默认的。

InnoDB 支持**事务**，**回滚**，**崩溃修复能力**，**多版本并发控制的事务安全**，**行级锁**，**AUTO_INCREMENT**和**外键**。是最常用的引擎。其缺点是读写效率稍差，占据的数据空间稍大。MySQL 的其他存储引擎均不支持事务。

可以在 my.ini 或 my.cnf 中修改默认引擎，需要停止正在运行的 MySQL 服务。或者使用语句

```sql
SET DEFAULT_STORAGE_ENGINE=MyISAM;
SHOW VARIABLE LIKE '%storage_engine%';
```

一般有需要事务支持都选择 InnoDB 引擎，但 MyISAM 和 MEMORY 引擎也有应用场景，如下：

1. MyISAM 引擎出入数据快，空间和内存使用较低，如果表**主要是用于插入和读取记录**，那选择 MyISAM 能有更高的效率。如果**应用的完整性、并发性要求不高**，也可以选择 MyISAM。
2. MEMORY 引擎所有数据都在内存中，基于 HASH，处理速度快，但安全性不高，如果需要**很快的读写速度**，可以使用此引擎，但是对表大小有限制，官方建议是用于**临时表的存储**。

需要注意的是，同一数据库中不同的表可使用不同的引擎，故可针对表的特性来选择存储引擎。

## 表操作

数据表的设计，有一些基本的原则和理念：

1. 标准化和规范化。主要是三种范式，这在数据库原理中有说明。
2. 数据驱动。采用数据驱动而非硬编码的方式，可以增强系统的灵活性和扩展性。如：角色权限管理的信息可以存放在数据库中，这样就可以做到很细致的权限控制。事实上，如果过程是数据驱动的，就相当于把相当大的责任交给用户，由用户来维护自己的工作流程。
3. 考虑各种变化。考虑表中哪些字段未来可能发生变化。
4. 表和表的关系。一对一，一对多，多对多。

表操作主要是表的 CRUD ，在 cheatsheet 中的 SQL 语句中有叙述。

## 索引

索引是一种特殊的数据库结构，可以用来提高查询速度。所以表中的一列或多列的值进行排序的一种结构。

索引类型：顺序文件上的索引、B+树索引、散列索引和位图索引。

- 顺序文件上的索引。是针对按指定属性值升序或降序存储的关系，在该属性上建立顺序索引文件，该文件由属性值和相应的元组指针组成
- B+树索引。**将索引属性组织成 B+ 树的形式，叶结点为属性值和相应元组指针。**
- 散列（hash）索引。将索引值按散列函数映射到相应的桶中。
- 位图索引。用位向量记录索引属性中可能出现的值，每一个位向量对应一个可能值。

### 索引的分类

1. 普通索引。不附加任何条件，可以在任何数据类型上创建，其值的约束取决于字段本身的约束。
2. 唯一性索引。使用 `UNIQUE` 参数设置唯一索引，限制了字段的值是唯一的，主键是特殊的唯一性索引。
3. 全文索引。使用 `FULLTEXT` 参数设置全文索引。只能创建在 `CHAR`,`VARCHAR`,`TEXT`类型的字段上，**可以提高大量字符串类型数据查询的速度**。
4. 单列索引。只要保证对应一个字段即可，以上三种索引都可称为单列索引。
5. 多列索引。一个索引对应多个字段，但是**只有查询条件中使用了这些字段中的第一个字段时才会应用该索引**。
6. 空间索引。使用 `SPATIAL` 参数设置空间索引。只能建立在空间数据类型上，包括`GEOMETRY`,`POINT`,`LINESTRING`,`POLYGON`等，目前只有 MyISAM 引擎支持空间类型。

### 索引的设计原则

要考虑的是在哪些字段上建什么类型的索引。

1. 选择唯一性索引。唯一性索引的值是唯一的，可以快速确定一条记录。
2. 为经常需要排序、分组、联合操作的字段建索引。建索引之后可以有效的避免排序操作。
3. 为常作为查询条件的字段建索引。
4. 限制索引的数目。索引也需要占用磁盘空间，而且修改表时，对索引的更新和重建也是麻烦事。
5. 尽量使用数据量少的索引。如`CHAR(100)`类型的字段索引速度慢于`CHAR(10)`的字段。
6. 尽量使用值的前缀来索引。
7. 及时删除不再使用或很少使用的索引。

### 操作索引

```sql
-- 创建索引
CREATE INDEX inex_name ON table_name(prop_name[(length)][ASC|DESC]);
-- 建表时建索引
CREATE TABLE tb_name(
    name char(10),
    INDEX idx_name(name)
);
-- 通过修改表定义创建
ALTER TABLE tb_name ADD INDEX|KEY idx_name(prop_name[(length)][ASC|DESC]);
-- 多列索引
CREATE INDEX inex_name ON table_name(prop_name1[(length)][ASC|DESC],prop_name2[(length)][ASC|DESC],……);
-- 使用索引与否对比,如：
EXPLAIN SELECT * FROM student WHERE name = 'jack';
-- 删除索引
DROP INDEX idx_name ON tb_name;
```

```
未使用索引
mysql> explain select * from student where name='jack';
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+ 
| id | select_type | table   | type | possible_keys | key  | key_len | ref  | rows | Extra       |
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+ 
|  1 | SIMPLE      | student | ALL  | NULL          | NULL | NULL    | NULL | 3276 | Using where |
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+
使用索引
mysql> create index idx_name on student(name);
mysql> explain select * from student where name='jack';
+----+-------------+---------+------+---------------+----------+---------+-------+------+-------------+
| id | select_type | table   | type | possible_keys | key      | key_len | ref   | rows | Extra       |
+----+-------------+---------+------+---------------+----------+---------+-------+------+-------------+
|  1 | SIMPLE      | student | ref  | idx_name      | idx_name | 25      | const | 1023 | Using where |
+----+-------------+---------+------+---------------+----------+---------+-------+------+-------------+
```

## 视图

视图是从一个表或多个表中导出来的**虚拟表**，就像一个窗口，只可以看到特定的数据，可以保护数据安全性。视图的数据不是一直存在的，而是**根据查询时动态生成**的。视图的**创建删除不影响基本表**，但**更新会影响基本表**。**当视图来自多个基本表时，不允许添加和删除数据**。

```sql
-- 创建视图(简化版)
CREATE VIEW view_name as SELECT statement
-- 创建(修改)视图（完整版）
CREATE[OR REPLACE][ALGORITHM=[UNDEFINED|MERGE|TEMPLATE]]
VIEW view_name(column_list)
AS SELECT  statement
[WITH [CASCADED|LOCAL]CHECK OPTION]
-- ALGORITHM=[UNDEFINED|MERGE|TEMPLATE]，选择创建视图的算法，UNDEFINED自动选择，MERGE将使用视图语句与定义视图语句合并，TEMPLATE将视图结果存如临时表，用临时表来执行语句
-- [CASCADED|LOCAL]，CASCADED为默认值，表示更新视图时要满足相关视图和表的条件；LOCAL表示更新时满足视图本身定义的条件即可。

-- 视图可以和表一样被查看
SHOW TABLES;
DESC view_name;
```

## 触发器

有事件触发预定义操作，事件包括`INSERT`、`UPDATE`、`DELETE`语句。应用场景如：

1. 新员工入职后离职后，员工总数都应该相应的加减
2. 学生毕业后删除学生数据，则其借书数据也应该同时被删除。

```sql
-- 创建触发器
CREATE TRIGGER trigger_name BEFORE|AFTER trigger_event ON tb_name FOR EACH ROW trigger_statement;
-- trigger_event的值为INSERT、UPDATE、DELETE，FOR EACH ROW表示任何一条记录上的操作满足触发条件都会触发，trigger_statement表示触发之后执行的语句

-- 一个触发器的例子
create trigger tg_log_remove after delete on student for each row insert into logger(del_date) values(now());

-- 删除触发器
DROP TRIGGER tg_name;
SHOW TRIGGERS;
```

触发器注意点：

1. 同一个表相同的事件只能创建一个触发器。此处 `before insert`和 `after insert` 是不同类型的触发器。
2. 不需要的触发器要及时删除，否则会影响数据。
3. 触发器有不可控的因素，故应更多的使用编程手段进行操作。

## 事务和锁

InnoDB 使用 REDO 和 UNDO 日志支持事务。

### REDO 日志

事务执行时需要将执行的事务日志写入 REDO 日志，当每条更新 SQL 执行时，首先将 REDO 日志写入缓冲区，当执行 `COMMIT` 时，日志缓冲区的内容将被刷新到磁盘。日志缓冲区的刷新方式和间隔时间可以通过 my.ini 中参数 `innodb_flush_log_at_trx_commit`控制。在 MySQL 崩溃恢复时会重新执行 REDO 日志中的记录。

### UNDO 日志

UNDO 日志主要用于事务异常时的数据回滚，具体就是复制事务前的数据库内容到 UNDO 缓冲区，然后在合适的时间将内容刷新到磁盘。

### 事务隔离级别

SQL 标准制定了四种隔离级别，用于指定事务中的哪些数据改变其他事务可见，哪些不可见。低级别的隔离可以支持更高的并发处理吗，占用的资源也更少。 InnoDB 系统级事务隔离级别可以使用以下语句设置：

```sql
-- 未提交读（read-uncommitted）
SET GLOBAL TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
-- 提交读（read-committed）
SET GLOBAL TRANSACTION ISOLATION LEVEL READ COMMITTED;
-- 可重复读（repeatable-read）
SET GLOBAL TRANSACTION ISOLATION LEVEL REPEATABLE READ;
-- 可串行化（serializable）
SET GLOBAL TRANSACTION ISOLATION LEVEL SERIALIZABLE;
```

#### read-uncommitted 读取未提交的内容

此级别中，所有事务都能看到其他事务未提交的数据。一般很少用，读取未提交的数据称为脏读。

### read-committed 读取提交内容

是大多数系统默认的级别，但不是 MySQL 的默认级别。这中隔离级别可能会导致不可重复读的情况。假设如下情况，事务 A 和事务 B同时开启，此时，事务 B 修改数据**并提交**，则事务 A 前后会读取到不同的数据。

### repeatable-read 可重复读

是 MySQL 默认的隔离级别，可保证在并发环境下，多个事务可读取到同样的数据。理论上可能导致幻读，即数据行的增减。如一个事务 A 对所有数据行做了修改，此时事务 B 增加了一行数据，则事务 A 会发现有一行数据没有被修改到。InnoDB 通过多版本比并发控制（Multi_Version_Concurrency Control,MVCC）解决该问题。

InnoDB 的 MVCC 机制：为每个数据行增加两个隐含值，记录行的创建时间和过期时间。每一行存储时间发生时的系统版本号，每个查询根据事务的版本号来查询结果。这样查询指定版本号可解决不可重复读的问题，因为其他事务增加或减少的数据是处于更新的版本，而按照初始次查询的版本号再次查询则不会涉及更新的数据。

### serializable 可串行化

最高的隔离级别，通过强行的事务排序，解决幻读的问题，实现方法是在每个读的数据行上加共享锁。这个级别会导致大量的超时和锁竞争，一般不推荐使用。

## InnoDB 锁机制

关于锁的类型和锁的粒度在数据库原理中有叙述。

# MySQL 高级应用

## 用户安全管理

### 权限表

通过权限表来控制不同用户拥有的权限，控制用户对数据库的访问。权限表存在`mysql` 数据库中，权限表中最重要的就是`db`和`user`表，还有 `tables_priv`,`columns_priv`,`procs_priv`等表。

#### user 表

是最重要的一个权限表，有 45 个字段，可分为四类：用户列、权限列、安全列和资源控制列。

1. 用户列。主要是 `Host`,`User`,`Password`。
2. 权限列。决定了用户的权限，描述在全局范围内用户允许对数据和数据库的操作，这些字段基本是以 `_priv`结尾的。
3. 安全列。有 6 个字段，连个 ssl 相关，两个 x509相关，两个授权加密。
4. 资源控制列。限制用户使用的资源。
   1. max_question.每小时允许执行的查询次数
   2. max_update.每小时允许执行的更新次数
   3. max_connections.每小时允许执行的连接次数
   4. max_user_connections.用户同时允许建立连接的次数。

#### db 表

存储用户对某个数据的操作权限，决定用户能从哪个主机存取哪个数据库。字段大致可分为两类：用户类和权限列。

1. 用户列。`Host`,`Db`,`User`。
2. 权限列。这些字段基本是以 `_priv`结尾的。

#### 其他权限表

主要是：`tables_priv`,`columns_priv`,`procs_priv`。

用于对表、列、存储过程（函数）设置操作权限。

### 账户管理

#### 登入退出 MySQL 服务器

```shell
mysql [-h hostname|hostIP -P port] -u username -p password [db_name -e sql_statement]
```

#### 用户管理

```sql
-- 创建用户
CREATE USER uname[IDENTIFIED BY [PASSWORD]'pwd'];
-- uname构成：uanme@host
-- 若密码是字符串，可不用 PASSWORD 关键字

-- 插入mysql.user创建用户
INSERT INTO mysql.user(Host,User,authentication_string) VALUES(...);-- PASSWORD() 函数给密码加密

-- 删除用户
DROP USER uname
-- 也可操作 mysql.user 表来删除用户

-- root 用户修改自己的密码
mysqladmin -u uname -p password "new pwd";
-- 也可操作 mysql.user 表修改，注意新密码应该用 PASSWORD(new_pwd) 函数加密

-- root 用户修改普通用户密码
SET PASSEORF FOR 'uname'@'host'=PASSWORD(new_pwd);
-- 普通用户自己修改密码
SET PASSWORD=PASSWORD(new_pwd);
```

## 日志管理

MySQL 日志记录了数据库的日常操作的和错误信息，分为二进制文件、错误日志、通用查询日志和慢查询日志。

- 二进制文件：记录所有更改数据的语句，可用于数据复制。
- 错误日志：记录MySQL运行期间出现的问题。
- 查询日志：记录建立的客户端连接和执行的语句。
- 慢查询日志：记录所有执行时间超过`long_query_time`的查询，或未使用索引的查询。

默认只启动错误日志。启用日志会降低数据库的性能，MySQL会花很多时间记录日志，日志也会占用大量磁盘空间。

进阶参考：《MYSQL内核_InnoDB存储引擎》 姜承尧