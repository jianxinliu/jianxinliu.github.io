# ORA-30043: Invalid value ‘…’ specified for parameter ‘Undo_Management’

## 描述

在进行“将回退段的管理方式改为自动管理”的实验时，原本的SQL语句是

```SQL
alter system set undo_management=auto scope=spfile;
```

但是由于对段的管理方式不是很清楚，错写成了

```SQL
alter system set undo_management=local scope=spfile;
```
奇怪的是当时并为报任何异常，显示系统已修改，但当我第二次启动ORACLE并键入`startup`时,报如下异常

```SQL
ORA-30043: Invalid value 'LOCAL' specified for parameter 'Undo_Management'
```

## 解决

在[ORA.codes](http://ora.codes/ora-30043/)里查到如下信息：

#### ORA-30043: Invalid value ‘…’ specified for parameter ‘Undo_Management’

**Message:**
ORA-30043: Invalid value ‘*string*‘ specified for parameter ‘Undo_Management’

**Cause:**
the specified undo management mode is invalid

**Action:**
Correct the parameter value in the initialization file and retry the operation.

于是

```SQL
---若没有此步骤，直接打开的参数文件可能不是最新的，里面没有 *.undo_management 参数
create pfile from spfile
```

打开初始化文件，将里面的`*.undo_management='LOCAL'`修改为`*.undo_management='AUTO'`，再执行

```SQL
---应用更新后的参数文件启动数据库
create spfile from pfile
```

数据库就可以正常启动了。