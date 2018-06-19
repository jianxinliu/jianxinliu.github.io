# Linux 命令行大全

**说明：{此中内容即命令的用法举例}<表示常用参数>**
## linux 知识：

文件权限：文件所属者、文件所属用户组、其他人的权限
用户权限表示：

- r :read：打开文件或者使用ls 查看目录下的内容
- w :write: 修改文件内容，在目录中创建或删除文件
- x :excute: 将文件作为程序运行或使用cd 进入目录

文件权限表示：-rw-r--r--
解读：- rw- r-- r--，分为四组，第一组是类型，-表示文件、d表示目录、l表示链接，l后的属性始终是rwxrwxrwx，是伪属性，真正的属性由链接指向的文件确定

链接：硬链接，软链接（符号链接），链接类似windows的快捷方式

后面三组分别对应文件的三个权限，以及其对应的用户权限。
例如：-rw-r--r--表示：这是一个文件，文件所属者拥有读写权限、文件所属用户组拥有读权限、其他人拥有读权限	

## 简单命令：

linux 命令格式：
​		command -options arguments
获取帮助：

- man command
- help functionName
  - info bash
- command --help

file:查看文件类型，使用空格分隔多个文件
{file test.txt 

test.txt: ASCII text}

**pwd**:present working directory

**cd**  :change directory /表示root用户主目录、~表示当前用户主目录、-表示上一次操作的目录
{cd /}{cd ~}{cd -}{cd path}

**ls**  :list {ls a directory}

- <-R > recursive,递归地查看指定文件及其子文件的内容
- <-X > 根据后缀，按字典序排序结果
- <-F > 在文件后加符号以区别可执行文件和目录
- <-RXF > 连接多个选项
- <-1 >one,show in single line
- <-l >long

**less**: 分屏查看信息(可查看一切文件) 方向键翻页，/ to search something    q to quit

'>': 重定向符号，如在命令后接重定向，则表示将命令的执行结果输出到文件，> 后接文件名

'>>'：从文件尾部追加内容，单个定向符会覆盖文件内容（如果文件内有内容的话）
{date > date.txt}将当前日期输出到文件 date.txt中
{> /path/newfile}创建新文件
{> /path/oldfile}清空此文件的内容

date:显示日期   cal：显示当月的日历，高亮当天

**mkdir**:make directory <-v>:verbose 显示详细过程
<-p>:同时创建子目录

rm :remove 删除文件

- <-v>:verbose
- <-i>:interactive
- <-r>:recursive
- <-f>:force
- <-rf>:combine two options,完全删除，难以恢复，可以删除非空目录

rmdir:删除空目录

mv :move ,默认会覆盖同名文件
<-r>移动子文件     -v  -i  都适用
{mv test.txt text}修改文件名

ln:line 创建链接，<-s>:symbolic创建软链接
{ln file lineName}

type:{type cd}
which:{which ls}
man:manual{man ls},程序的手册文档分成多个部分

1. 用户命令
2. 内核系统调用的程序接口
3. c库函数程序接口
4. 特殊文件，如设备节点和驱动程序
5. 文件格式
6. 游戏和娱乐，例如屏幕保护程序
7. 其他杂项
8. 系统管理命令

   {man 5 passwd}				
   whatis:{whatis ls}

## 重定向：I/O重定向，上文有提到

- 2> ：将标准错误流重定向。2是shell在内部用文件描述符的索引，同时，0表示标准输入流、1表示标准输出流
- &>:将标准输出和标准错误流同时重定向
- |:管道，通过管道符，将前一个命令的标准输出链接到后一个命令的标准输入{ls -l /usr/bin | less}

## 扩展：路径名扩展

echo d*,列出所有以d开头的文件或目录
花括号扩展
应用：mkdir {2009..2010}-0{1..9} {2009..2010}-{10..12} 按月创建以2009-01到2010-12为名的文件夹

## 权限：

- chmod:change mode,修改权限（文件模式），只有文件所有者或者root用户才能执行此命令
- chown:change owner 修改所有者
- chgro:change group 修改文件所属用户组

#### chmod使用方法：八进制法和符号法，

符号法：{chmod u+x test.txt}给test.txt 的所有者加上对此文件的执行权限

g:group,o:others,a:all,u:user
+:加权限，-:减权限，=:除了此权限，若有其他权限，则去掉
指定多种权限时，用逗号分隔

#### chown使用方法：

- chown [owner][:group] file ...
- {chown bob:users file}将file文件的所有者改为用户bob，并将其所属用户组改为users
- {chown :users}将file文件的所属用户组改为users,所有者不变
- {chown bob}将file文件的所有者改为用户bob，所属用户组不变

#### umask:设置默认权限

- su:以其他用户和组ID的身份来运行shell
- su [-[l]][user] ,没有任何参数，默认启动root用户
- sudo:以另一个用户的身份执行命令

## 进程：

**ps** :查看进程信息,只输出和当前终端会话相关的进程信息，在ps命令执行瞬间及其状态的快照x，显示所有进程信息，不管由哪个终端控制
结果解读：TTY列：进程的控制终端

**top**:动态查看进程信息（默认3秒刷新）

**STAT**:进程状态。S:sleeping/stoped,R:run/ready,D:waiting for I/O

**T**:terminated暂停，Z:zombie无效的，<:高优先级进程，N:nice低优先级（友好进程）进程aux（不带前置连字符）,显示更多信息
结果解读：%CPU:cpu使用占比，%MEM:内存使用占比，VSZ:虚拟好用内存大小，RSS:实际使用内存大小（KB）,START:进程开启的时间

**kill**:中断进程，发送中断信号到进程，进程自己关闭
kill [-signal] PID ...   man kill

## 环境：
shell在环境中存储了两种变量，分别是环境变量和shell变量，但这在bash中基本没区别。
printenv [var name] 输出环境变量[值]
echo $USER 显示环境变量的值

## VI：命令模式下使用
- :q：退出

- :q!：强制退出

- i:进入插入模式（编辑模式）

- Esc:从插入模式返回命令模式，按两次返回最初状态

- :：保存文本

- :w：写入磁盘

- vim file.txt ：创建新文件

- o：在当前行的上方插入一行

- O：在当前行的下方插入一行

- u：撤销操作

- 删除：d实际上是剪切文本

  - x：删除光标处的字符
  - 3x：删除光标处的字符和后两个字符
  - dd：删除一行
  - 5dd：当前行和之后4行
  - dW：当前字符到下一单词的起始，当光标在单词的第一个字母下，可删除这个单词
  - d$：当前字符到行尾
  - d0：当前字符到行首

- 剪切、复制、粘贴

- y进行复制

  - yy:复制当前行
  - 。。。使用方法类似d命令
  - p进行粘贴
  - 合并行：J

- 行内搜索：fa：在当前行搜索a，并将光标移动到搜索到的下一指定字符，;重复上一次搜索

- 文件内搜索：  /something  ,n 重复此搜索

- 查找与替换：:%s/old/new/gc

  - :表示启动一条ex命令
  - %表示却低估操作范围为全文件，也可以用1,5表示搜索范围是第一行到第五行，若不指定，则在单飞倩行生效s指定具体操作，替换操作（搜索和替换）

  - /old/new表示搜索和替换的文本

  - g表示global，对搜索到的每一个进行替换
  - c表示进行询问确定(询问的次数和查找到的结果数相同)，此时显示结果
  - y表示yes
  - n表示no
  - a表示all,执行所有替换
  - q表示quit，停止替换
  - l表示last，此次是最后一次替换，执行后退出


编辑多个文件：

- vi file1 file2 ...   ：打开多个文件

- :n 切换到下一个文件（ex命令）

- :N 切换到上一个文件

- :buffers 查看当前正在编辑的文件列表
  - :e fileName 载入文件

- :r ，read a file 将文件读入当前光标处
## 软件包管理

在库里搜索包：
Debain:
```shell
apt-get update
apt-cache search search_string
```
Red Hat(CentOS,Fedora):	
```shell
yum search packageName
```
安装库内的包：
Debain:
```shell
[apt-get update]apt-get install packageName
```
Red Hat(CentOS,Fedora):	
```shell
yum install packageName
```
删除软件包：
Debain:
```shell
apt-get remove packageName
```
Red Hat(CentOS,Fedora):	
```shell
yum erase packageName
```
更新库中的软件包：
Debain:
```shell
[apt-get update]apt-get upgrade
```
Red Hat(CentOS,Fedora):	
```shell
yum update
```
列出已安装的软件包列表：
Debain:
```shell
dkcp --list
```
Red Hat(CentOS,Fedora):	
```shell
rpm -qa
```
判断软件包是否安装：
Debain:
```shell
dpck --status packageName
```
Red Hat(CentOS,Fedora):	
```shell
rpm -q packageName
```
显示已安装软件包的相关信息：
Debain:
```shell
apt-cache show packageName
```
Red Hat(CentOS,Fedora):	
```shell
yum info packageName
```
查看某个文件具体由哪个软件包安装得到：
Debain:
```shell
dpck --search fileName
```
Red Hat(CentOS,Fedora):	
```shell
rpm -qf fileName
```

## 存储介质：

		挂载、卸载存储设备
## 网络：

		ping :向网络主机发送特殊数据包以确定网络链接是否正常
		wget:非交互式下载器{wget www.baidu.com}  man wget

## 文件搜索：

		find      man find
## 归档和备份：

		gzip:压缩文件，默认压缩文件取代原来的文件 man gzip 
		gunzip:解压文件
		tar:归档文件

## shell 脚本编写

​	
​	