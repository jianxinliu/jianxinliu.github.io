# tar

```shell
# 解压 tar.gz 文件
tar -zxvf target_file.tar.gz

# 压缩并打包文件为 tar.gz
tar -zcvf target_file.tar.gz  source_file

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
dw， dd 删除当前行； {x,y}d 删除 x~y 行

# 复制，粘贴
y(yank) yy 复制整行； nyy 复制从当前行开始往下的 n 行
p（可以先 o 另起一行）

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

# File Permissions

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

ps -ef  (-e select all process.Identical to -A)(-f do full format listing, -H show process hierarchy)

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



# file

```shell
find -name filename // 按名称查找文件

whereis file/command

grep pattern          params: [-i]:忽略大小写; [-v] 反选 ;[-a] 处理二进制文件; [-R] 递归方式


# remote to local
scp uname@remote_host:<path_to_file> <local_path>

# local to remote
scp local_file uname@remote_host:<remote_path>
```





# network

```shell
// 查看端口使用情况（any one）
sudo lsof -i -P -n | grep LISTEN
sudo netstat -tulpn | grep LISTEN
sudo lsof -i:22 ## see a specific port such as 22 ##
sudo nmap -sTU -O IP-address-Here
```





# sed、awk

sed for stream editor



# Disk

```sh
# disk useage in human style
du -sh *
```



