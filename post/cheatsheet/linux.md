# tar

```shell
# 解压 tar.gz 文件
tar -zxvf source_file target_file

# 压缩并打包文件为 tar.gz
tar -zcvf source_file

# 不解压，查看文件
tar -tf source_file
```

# Vim

```shell
# quit or without saving
:q[!]

# quit with saving changes
:wq 
:x

# search
/pattern sflgkfg 

# undo
u

# 删除光标处的一个字符，可以在 x 前加数量表示删除多个
x

# 删除光标后的一个单词
dw

# 复制，粘贴
y(yank)
p

# undo 
u
```

# Bash Veriables

```shell
# 打印环境变量
env

# 打印某个环境变量，如 $PATH
echo $NAME

# 暂时设置某个环境变量,如：export PATH=$PATH:/usr/foo/bar
export NAME=value
```

# Fiel Permissions
```shell
# 修改文件的权限
chmod mode file

# 修改整个文件夹下文件的权限
chmod -R mode folder

# mode 说明
# 4 == read(r)
# 2 == write(w)
# 1 == execute(x)
```
# Process Management
```shell
# process snapshot
ps

# real time process
top

# kill process with pid
kill -9 pid

# kill process with name
pkill name

# kill all process starts with name
killall name
```

# IO Redirection
```shell
# 将命令执行结果输出到文件（改变输出流到文件）
cmd > file

# 执行命令不输出（将输出丢弃）
cmd > /dev/null

# 追加输出到文件
cmd >> file

# 重定向错误流
cmd 2> file

# 重定向标准输出流到标准错误流
cmd 1>&2

# 重定向标准错误流到标准输出流
cmd 2>&1

# 重定向所有流到文件

cmd &> file

```