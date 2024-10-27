
# linux 命令

### envsubst

env-substitute: 将文件中的变量使用环境变量进行替换

```txt
Hello user $USER in $DESKTOP_SESSION. It's time to say $HELLO!
```

> export HELLO="good morning"
> envsubst < welcome.txt
> Hello user joe in Lubuntu. It's time to say good morning!



### 🔲expect

https://www.cnblogs.com/saneri/p/10819348.html

https://linux.die.net/man/1/expect

expect常用命令总结:

>   spawn               交互程序开始后面跟命令或者指定程序
>   expect              获取匹配信息匹配成功则执行expect后面的程序动作
>   send exp_send       用于发送指定的字符串信息
>   exp_continue        在expect中多次匹配就需要用到
>   send_user           用来打印输出 相当于shell中的echo
>   exit                退出expect脚本
>   eof                 expect执行结束 退出
>   set                 定义变量
>   puts                输出变量
>   set timeout         设置超时时间

使用示例：

```shell
/usr/bin/expect << EOF
spawn ssh root@${router}
expect "*password:" {send "${pass}\r"}
expect "*#" {send "./${bin} xxx\r"}
expect "*#" {send "rm -f /root/router.init\r"}
expect "*#" {send "reboot\r"}
expect eof
EOF
```

实验：

1.   创建需要交互的程序 que.sh

```sh
#!/bin/bash
 
echo "Enter your name"
 
read $REPLY
 
echo "Enter your age"
 
read $REPLY
 
echo "Enter your salary"
 
read $REPLY
```

2.   使用 expect 进行交互

```sh
/usr/bin/expect <<  EOF
spawn ./que.sh
expect "Enter your name\r" {send "jianxin\r"}
expect "Enter your age\r" {send "14\r"}
expect "Enter your salary\r" {send "33333\r"}
EOF
```



### ping

ping 用于网络诊断，判断连通性

原理：一台设备给目标设备发送 ICMP 报文，等待其相应，并记录时间。所**耗费的时间喻示了路径长度**，**重复请求响应的一致性也表明了连接的质量**。ping 回答了两个问题： **是否有连接，连接的质量如何**。

常用选项：

```text
-c count 设置发送报文的数量，Unix 系统不指定则会一直发，windwows 默认发四次
-i wait 设置两次发送之间间隔的秒数。默认是 1s
-n 输出数字形式
```

### telnet

简化版的 ssh。也可用于连接远程机器。和 ping 相比，telnet 可以探测机器上的端口是否可用

### nmap

扫描机器上开放的端口机器被哪个程序占用

执行后可以发现。echo -> 7; ftp -> 21; ssh -> 22; telnet -> 23; smtp -> 25; http -> 80 ……

### traceroute

展现到某个网络节点需要经过的中间节点以及其连接情况

原理是通过给到达目标主机中间的所有节点发送 ICMP 请求，并计时。

命令格式： `traceroute [options] host/ip packetSize`

### Systemd

https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html

```shell
systemctl start/stop/restart/status <service>
```

配置文件位置：

1.   系统自启动时调用： /etc/systemd/system
2.   自定义：/usr/lib/systemd/system



### ifconfig

https://www.computerhope.com/unix/uifconfi.htm

ifconfig -> interface onfig 用来查看和操作网络配置的

不带参数，直接执行，可以查看本机所有网络接口的情况

```shell
# 查看某个网口的信息
ifconfig <interface name>
```

设置网络接口：

```shell
# 同时配置 ip 地址，子网掩码，广播地址
sudo ifconfig <interface name> <ip> netmask <mask> broadcast <broadcast>
```

**ifconfig 只能分配静态 IP， 动态 IP 需要使用 DHCP **



地址族说明：

-   inet: tcp/ip, 也叫 ipv4
-   inet6: ipv6

### sed

```shell
# 文本、文件替换

# 将 log.txt 文件中的 A 都替换成 B, 忽略大小写。并将结果写回 log.txt
# 不带 -i 则将替换结果输出
# -e 支持多个，标识替换的正则   -e '' -e ''
# -e 后面的字符串格式 's/pattern/replacement/flags'
sed -i -e 's/A/B/i' log.txt
```



### tc

使用 tc (traffic control) 限制主机带宽

tc man page https://man7.org/linux/man-pages/man8/tc.8.html

https://catbro666.github.io/posts/357ad3ec/



```shell
UPLOAD_SPEED=2
DOWNLOAD_SPEED=2

### add upload
ip link add dev ifb0 type ifb
ip link set ifb0 up
# redirect ingress to ifb0
tc qdisc add dev eth0 ingress handle ffff:
tc filter add dev eth0 parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0
# add qdisc
tc qdisc add dev ifb0 root handle 1:0 htb default 1
# add default class
tc class add dev ifb0 parent 1:0 classid 1:1 htb rate ${UPLOAD_SPEED}mbit ceil ${UPLOAD_SPEED}mbit
### add download
tc qdisc add dev eth0 root handle 1:0 htb default 1
tc class add dev eth0 parent 1:0 classid 1:1 htb rate ${DOWNLOAD_SPEED}mbit ceil ${DOWNLOAD_SPEED}mbit
```

更新带宽（以 ifb0 接口为例）：

```shell
# 先删除
tc class del dev ifb0 parent 1:0 classid 1:1 htb rate ${UPLOAD_SPEED}mbit ceil ${UPLOAD_SPEED}mbit
# 再增加
tc class add dev ifb0 parent 1:0 classid 1:1 htb rate ${UPLOAD_SPEED}mbit ceil ${UPLOAD_SPEED}mbit
```

### awk

https://www.bookstack.cn/read/bash-tutorial/docs-archives-commands-awk.md



### sort

`sort`命令将文本文件的所有行排序后输出。

```
$ sort file1.txt file2.txt file3.txt > final_sorted_list.txt
```

它的参数如下。

-   `-b` `--ignore-leading-blanks` 默认情况下，排序用的是每行的第一个字符。这个参数忽略每行开头的空格，从第一个非空白字符开始排序。
-   `-f` `--ignore-case` 让排序不区分大小写。
-   `-n` `--numeric-sort` 按照数值排序，而不是字符值，用于行首是数值的情况。
-   `-r` `--reverse` 按相反顺序排序。结果按照降序排列，而不是升序。
-   `-k` `--key=field1[,field2]` 指定按照每行的第几个字段（从1开始）排序，而不是按照行首字符排序。该属性可以多个连用，用于指定多重排序标准，还可以指定每个字段指定排序标准，这些值与全局属性一致，比如b（忽略开头的空格），n（数值排序），r（逆向排序）等等。
-   `-m` `--merge` 把每个参数看作是一个预先排好序的文件。把多个文件合并成一个排好序的文件，而没有执行额外的排序。
-   `-o` `--output=file` 把排好序的输出结果发送到文件，而不是标准输出。
-   `-t` `--field-separator=char` 定义字段分隔字符。默认情况下，字段由空格或制表符分隔。
-   `-u` 输出结果中删除重复行





### ssh

使用 pem 证书链接服务器。注意：本地 pem 证书的权限不能太大，否则会被拒绝连接

```
ssh -i /path/to/xxx.pem root@xx.xx.xx.xx
```

#### 隧道

将远端的端口通过隧道打通, 来达到连接本地端口可以访问远程端口上服务的能力. 如部署在集群内的服务,只能通过跳板机访问,则可以在本地和跳板机打一个隧道, 来方便本地访问. 

```shell
ssh -L 8091:127.0.0.1:8091 -L 9532:127.0.0.1:9532 -L 6560:127.0.0.1:6560 <目标机器,也可以是在 ssh config 中配置的名称>
```



#### SSH Config

https://linuxize.com/post/using-the-ssh-config-file/



具体如何配置可以参考 `man ssh_config`



配置大致格式：

```conf
Host hostname1
	SSH_OPTION value
	SSH_OPTION value
Host hostname2
	SSH_OPTION value
	SSH_OPTION value
Host * (匹配所有 Host)
	SSH_OPTION value
```

读取顺序是自上而下，一段段读取，先读取的 OPTION 优先级更高。



所以如果有相同的 Host 定义，先定义的生效。一般地，更精确的定义放在文件开头，更一般性的定义放在文件末尾。



### speedtest

```shell
curl -s https://raw.githubusercontent.com/sivel/speedtest-cli/master/speedtest.py | python -

wget --output-document=/dev/null http://speedtest.wdc01.softlayer.com/downloads/test500.zip
```



# Bash



## 行操作

https://www.bookstack.cn/read/bash-tutorial/docs-readline.md

相关快捷键

-   Ctrl + a: 光标到行首
-   Ctrl + e: 光标行尾 
-   Ctrl + w: 删除光标前面的单词
-   Ctrl  + u: 剪切光标到行首
-   Ctrl  + k: 剪切光标到行尾



# go-zero

https://go-zero.dev/cn/docs/goctl/goctl

### goctl

[指令大全](https://go-zero.dev/cn/docs/goctl/commands)

代码生成工具

- api 服务生成
- rpc 服务生成
- model 代码生成
- 模板生成

#### 安装方式

使用此命令将 goctl 安装到 `$GOPATH/bin` 下，手动将此路径加入环境变量
```shell
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest 
```



#### api 文件编写

[api 文件语法](https://github.com/zeromicro/zero-doc/blob/main/go-zero.dev/cn/api-grammar.md)

##### 1. type 声明请求与响应结构体

```api
type StuUpdateReq {
  // json 的参数以请求体的方式传入
	Name string `json:"Name"`
	// form 的参数，post 时是一个 form, get 时是 url 参数的形式
	Age int     `form:"Age"`
	// 该参数以 path 的方式传入
	Id int      `path:"Id"`
	// header 的参数以请求头的方式传入
	UserId string  `header:"UserId"`
}

type StuUpdateResp {
  BaseResponse
  Ret bool `json:"Ret"`
}
```

##### 2. server 声明基础信息

```api
@server (
	// 路由分组
	perfix: app/v1/stu
	// 加载中间件, 对应的中间件结构体名称为 CustomJwtMiddleware，此处可以省略 Middleware 后缀
	middleware: CustomJwt
)
```

##### 3. service 声明具体的路由定义

```api
service stu-api {
	// 指定 handler, 对应 handler 函数的名称为 stuUpdateHandler（同样是省略 Handler 后缀）
	@handler stuUpdate
	// 定义此路由的相关信息：请求方法，请求 path, 入参与返回值
	post /update (StuUpdateReq) returns(StuUpdateResp)
}
```



>   新增时，修改 api 文件后，再次运行 goctl 生成命令即可增量式生成新增的代码

#### api 服务生成

大致步骤：


1. 编写 xx.api 文件
2. 运行 `goctl api go -api xx.api -dir . -style gozero` 自动生成项目目录。具体命令使用方式参考[官方文档](https://go-zero.dev/cn/docs/goctl/api)
3. 运行 `go mod tidy` 自动寻找依赖，写到 go.mod 中，并进行依赖下载
4. 编写具体的逻辑

或者直接使用直接生成项目结构和基本代码：

```shell
goctl api new <service name>
```



#### rpc 服务生成



#### model 生成

https://go-zero.dev/cn/docs/goctl/model

支持通过 MySQL ddl 生成代码：

```shell
goctl model mysql ddl -src="<path to ddl.sql>" -dir="<path for model code>" -c
```

通过 datasource 生成：

```shell
goctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/database" -table="*" -dir="./model"
```

自动生成基本 CRUD 代码结构。



#### 其他服务生成

https://go-zero.dev/cn/docs/goctl/other

-   docker file `goctl docker`
-   Kubenetes 部署文件 `goctl kube`



### config

包含以下配置：

1.   `rest.RestConf`: 主要包含 restful api 相关的配置，具体内容可参考该结构体
2.   `zrpc.RpcConf`: 主要包含 rpc 相关的配置，具体内容可参考该结构体



### jwt

go-zero 的 jwt 中间件，在登录后可以根据请求头中带的 token, 解析出用户名，并将用户名放在每个请求的 context 中

大致原理是：

​	1.	**拦截请求**：中间件拦截每一个传入的 HTTP 请求，检查是否包含有效的 JWT 令牌。

​	2.	**解析和验证 JWT**：使用配置中的密钥解析 JWT 令牌，并验证其有效性（如签名、过期时间等）。

​	3.	**提取用户信息**：从 JWT 的声明（claims）中提取 userId 等信息。

​	4.	**注入上下文**：将提取的 userId 添加到请求的上下文中，以便后续处理程序可以访问。

```go
func getJwtToken(secretKey, userId string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
  // 声明携带的字段
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
```

路由定义中，需要声明 Jwt 中间件

```go
server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/account/add",
				Handler: ops.AddAccountHandler(serverCtx),
			},
		},
  	// 启用 jwt
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/ops"),
	)
```



### go-zero 接口参数加解密

```go
package utils

import (
	"encoding/base64"

	"github.com/zeromicro/go-zero/core/codec"
)

// 加密 message。 key 必须是 base64 格式。返回的密文是 base64 格式的
func EncryptBase64(key string, message []byte) (string, error) {
	messageBase64 := base64.StdEncoding.EncodeToString(message)
	return codec.EcbEncryptBase64(key, messageBase64)
}

// 解密 cipher。 key 必须是 base64 格式
func DecryptBase64(key string, cipher string) ([]byte, error) {
	message, err := codec.EcbDecryptBase64(key, cipher)
	if err != nil {
		return []byte{}, err
	}
	bys, err := base64.StdEncoding.DecodeString(message)
	return bys, err
}
```

```go
package middleware

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"utils"
)

type CryptoMiddleware struct {
	KeyBase64 string
}

func NewCryptoMiddleware(keyBase64 string) *CryptoMiddleware {
	return &CryptoMiddleware{
		KeyBase64: keyBase64,
	}
}

func (m *CryptoMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 加密返回值
		cw := newCryptionResponseWriter(w)
		defer cw.flush([]byte(m.KeyBase64))

		// 解密请求体
		if err := decryptionRequest(m.KeyBase64, r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// 响应体重写
		next.ServeHTTP(cw, r)
	}
}

// https://github.com/zeromicro/go-zero/blob/master/rest/handler/cryptionhandler.go
func decryptionRequest(key string, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	plantText, err := utils.DecryptBase64(key, string(body))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	buf.Write(plantText)
	r.Body = io.NopCloser(&buf)
	return nil
}

type cryptionResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func newCryptionResponseWriter(w http.ResponseWriter) *cryptionResponseWriter {
	return &cryptionResponseWriter{
		ResponseWriter: w,
		buf:            new(bytes.Buffer),
	}
}

func (w *cryptionResponseWriter) Flush() {
	if flusher, ok := w.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

func (w *cryptionResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Hijack implements the http.Hijacker interface.
// This expands the Response to fulfill http.Hijacker if the underlying http.ResponseWriter supports it.
func (w *cryptionResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacked, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hijacked.Hijack()
	}

	return nil, nil, errors.New("server doesn't support hijacking")
}

func (w *cryptionResponseWriter) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}

func (w *cryptionResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *cryptionResponseWriter) flush(key []byte) {
	if w.buf.Len() == 0 {
		return
	}

	content, err := utils.EncryptBase64(string(key), w.buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if n, err := io.WriteString(w.ResponseWriter, content); err != nil {
		logx.Errorf("write response failed, error: %s", err)
	} else if n < len(content) {
		logx.Errorf("actual bytes: %d, written bytes: %d", len(content), n)
	}
}
```





## go-zero 相关代码学习



### syncx 包

代码位置：go-zero/core/syncx/singleflight.go





# gRPC



### Protobuf 文件格式

https://protobuf.dev/programming-guides/proto3/

#### message

定义一个消息传递的实体

```protobuf
message BaseResp {
  string requestId = 1;
  int32 code = 2;
  string message = 3;
}

message StuUpdateReq {
  string name = 1;
  int32 age = 2;
  int32 id = 3;
  string userId = 4;
}

message StuUpdateResp {
  BaseResp baseResp = 1;
  bool ret = 2;
}
```

Message 中包含的重点：

1.   类型：可以是简单的类型，也可以是复杂类型
2.   字段唯一码（Field Numbers）：**一旦这个 message 被使用，唯一码就不应该再更改。1-15 使用一个字节存储，包含常用的字段，15 开外的使用两个字节存储。一般需要在 15 以内保留一些空间用于增加。最小值是 1。**序列号定义可以采用 `10, 20, 30` 这样带间隔的，方便后续增加。
3.   字段修饰符：
     1.   `singular`: 默认修饰符，表示每个字段只能存在 0 个或 1 个。
     2.   `optional`: 指定一个字段是否是必须的。未设置值时，不会被序列化；设置了值时，可以被序列化和反序列化
     3.   `repeated`: 表示这个字段可以出现一次或者多次，会保留其原本的顺序。一般用来定义数组。`repeated string ids = 1;` 表示 ids 是一个字符串数组。
     4.   `map`: 表示一个 map 结构 

#### service

格式：

```protobuf
service Rpc {
  rpc Ping(Request) returns(Response);
  rpc UpdateStu(StuUpdateReq) returns(StuUpdateResp);
}
```



### grpcurl

使用 curl 的方式访问 grpc 服务接口 https://github.com/fullstorydev/grpcurl

基本用法: (localhost:8888 存在 grpc 服务 api.Service )

`grpcurl [flags] [address] [list|describe] [symbol]`

-   列举服务器支持的所有 grpc 服务

    `grpcurl localhost:8888 list`

-   列举服务支持的方法: 

    `grpcurl --plaintext localhost:8888 list api.Service`

-   调用指定方法(并添加 json 格式的请求体):

    `grpcurl --plaintext -d '{"user": "aa", "pass": "11"}' localhost:8888 api.Service/AddUser`

-   增加请求头

    `-H header:value -H header1:value1 ...`

-   描述方法,及其参数情况

    `grpc localhost:8888 describe api.Service/AddUser`





# 计算机网络



### 子网掩码

参考： https://www.bookstack.cn/read/network-basic/7.md

子网掩码的表示： 154.71.150.42/22 表示 154.71.150.42 这个 IP 的子网掩码是 255.255.252.0。计算方法如下：

22 表示 32 位的子网掩码中，前 22 位都是 1, 后 10 位是 0， 即 `11111111 11111111 11111100 00000000`，这样转成十进制的就是 `255.255.252.0`

前 22 位都是 1 ，表示这些位被**掩盖**了，不能用于表示该子网下的主机，即剩下能表示主机的位只剩 10 位，就是说，这个子网的这个 ID 下能表示的主机数是 $2^{10} = 1024$ , 是这个子网段下主机数最多的子网。

可以看出，这是一个 B 类网络，前 16 位表示网络号，22 - 16 = 6 位表示子网，也就是这个 B 类网络下，能有的子网段数量是 $2^6 = 64$ 个，总共能容纳的主机数： $\Sigma^{i}_{1 \le i \le 6}{2^i * 2^{16-i}}$

`154.71.150.42` 这个 IP 对应的二进制表示为： `10011010 1000111 100101/10 101010`， 因为其子网掩码为 22, 可以看出这是一个 B 类网络，则其前 16 （$\lfloor22 / 8\rfloor * 8 = 16$）位是不动的，并且其子网段总共有 6 ($22 \% 8 = 6$) 位



### NAT

Network Address Translation



SNAT Source Network Address Translation https://www.juniper.net/documentation/en_US/contrail20/topics/task/configuration/snat-vnc.html



### DNS

-   查询域名的 DNS 解析结果： `nslookup <domain>`
-   使用指定 DNS 服务器解析域名：`dig @<dns server> domain`。 可用来验证配置的 DNS 解析是否正常工作





# Kubernetes

## Kubernetes 概念

[Kubernetes 实践手册](https://www.bookstack.cn/read/kubernetes-handbook/guide-kubectl-cheatsheet.md)

[Kubernetes 云原生架构应用手册](https://www.bookstack.cn/read/kubernetes-handbook-201910/README.md)



目标：管理跨多个主机的容器，提供基本的部署、维护以及应用的伸缩。

基本概念：https://yeasy.gitbook.io/docker_practice/kubernetes/concepts

-   节点 Node：是运行 kubernetes 的主机
    -   可以是物理主机，也可以是虚拟机，每个节点都需要运行一些必要的服务以运行容器，如 docker, kubelet, 代理服务……
    -   容器状态用来描述节点当前的状态。主要有：Running, Pending, 
-   容器组 Pod: 一个 Pod 是由若干个容器组成的容器组，同个组内的容器共享相同的存储卷
-   容器组生命周期 Pod-states: 是容器所有状态的集合。包括：pod 类型，pod 生命周期，事件，重启策略，replication controllers
-   副本控制器 Replication controllers: 负责指定数量的 pod 在同一时间一起运行
-   服务 services: 是 pod 的高级抽象，同时提供外部访问 pod 的策略
-   卷 volumes: 就是一个目录
-   标签 labels: 用来连接一组对象，比如 pod。标签可以用来组织和选择子对象
-   接口权限: 端口、IP和代理的防火墙规则
-   web 界面： 可以通过 ui 操控 kubernetes
-   cli 命令： Kubectl



## 实践



### 使用 cloud-native-sandbox 在本地运行

安装 https://github.com/rootsongjc/cloud-native-sandbox.git 

下载仓库代码，参考以下代码进行操作： https://github.com/rootsongjc/cloud-native-sandbox


### 使用 minikube 在本地运行

使用 Minikube 在本地运行 kubernetes 单机版。[安装方式](https://minikube.sigs.k8s.io/docs/start/) 

通过 minikube 启动 k8s

`minikube start` 会自动安装 k8s，并且可以使用 kubectl 进行控制

`kubectl cluster-info` 命令可以看到当前通过 minikube 启动的 k8s 集群

`kubectl get po -A` 展示当前集群上的 pod

`minikube dashboard` 启动 k8s 控制面 UI



### kubectl 常用命令

https://cloud.tencent.com/developer/article/1638810

kubectl 常用选项

>   -   `kubectl options` 展示所有选项
>   -   此处的选项，可以传给任意子命令

-   -n, --namespace=''。设置本次 cli 命令请求的 namespace

#### get 获取资源信息

-   -o, --output 指定输出格式。json,yaml, wide……

-   -l: selector label selector, =, ==, !=

-   --sort-by='': 按指定字段排序，字段可以通过指定输出为 json 来看。格式为： `--sort-by='{.status.podIP}'` 按 pod 的 IP 排序

-   -A, --all-namespace=false. 如果指定，表示列出所有 namespace 下的资源，不指定，则只列出当前 namespace 下的

```she
kubectl get pods -o wide --sort-by='{.status.podIP}'
```

**常用资源类型列表**

`kubectl api-resources` 可列出所有资源类型

-   namespace, ns
-   nodes, no
-   presistenctVolumes, pv
-   pods, po
-   replicationControllers, rc
-   Services, svc
-   daemonSets, ds
-   replicaSets, rs
-   statefulSets, sts
-   Cronjobs, cj
-   Events, ev

##### pod

-n: namespace 指定 namespace



### 实验 k8s app 版本回滚

1.   使用 go-zero 时间一个简单的 echo api server， 监听 8888 端口
2.   将程序打包成镜像，并上传 docker hub（k8s 镜像需要从某个 registry 中拉取，可以在本地起，也可以直接 push 到 docker.io）
     1.   `docker buildx build -f ./app/Dockerfile -t xiawan12/docker-starter:002 . --push ` 推到 docker.io xiawan12 这个账号下
3.   通过配置文件，启动 k8s。`kubectl create namespace local-test`,  `kubectl apply -f test-docker.yml`
4.   `minikube node list` 查看当前集群的 ip <hostIP> (`kubectl -n local-test describe pods` Node)
5.   浏览器访问： `http://<hostIp>:30004/from/me` (根据具体 api server 接口调整)
6.   调整版本： 
     1.   edit deployment: `kubectl -n local-test edit deployment -f test-docker.yml`。修改 image 到对应的版本。保存后自动生效
     2.   set image: `kubectl set image deployment/test-docker test-docker=xiawan12/docker-starter:001 -n local-test ` 执行后自动生效
     3.   Rollout:
          1.   查看所有版本：`kubectl -n local-test rollout history deployment test-docker`
          2.   回滚至上一版本：`kubectl -n local-test rollout undo deployment test-docker ` **优点：可在很紧急，并且明确回滚到上一个版本就可以解决问题的情况下，立即恢复服务，而不用管具体哪个版本**
          3.   回滚至指定版本： `kubectl -n local-test rollout undo deployment test-docker --to-revision=<>` **缺点：需要明确知道每个版本号代表的功能**
7.   验证版本功能
8.   停止 `kubectl -n local-test delete deployment/test-docker` 
9.   **重启** `kubectl -n local-test rollout restart deployment/test-docker`

test-docker.yml 内容

```xml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-docker
  namespace: local-test
  labels:
    app: test-docker
spec:
  replicas: 2
  selector:
    matchLabels:
      app: test-docker
  template:
    metadata:
      labels:
        app: test-docker
    spec:
      containers:
      - name: test-docker
        image: xiawan12/docker-starter:001
        ports:
        - containerPort: 8888

        resources:
          limits:
            cpu: 2
            memory: 1Gi
          requests:
            cpu: 200m
            memory: 256Mi
        
---
kind: Service
apiVersion: v1
metadata:
  labels:
    app: test-docker
  name: test-docker
  namespace: local-test
spec:
  type: NodePort
  ports:
  - name: http
    port: 8888
    protocol: TCP
    targetPort: 8888
    nodePort: 30004
  selector:
    app: test-docker
```



### 标签与选择器

#### 标签

https://kubernetes.io/zh-cn/docs/concepts/overview/working-with-objects/labels/

k8s 标签系统，可以用来给各种资源增加标签，资源之间连接时，也通过标签来选择匹配。是一种**松耦合的组织系统资源的方式**。

标签可以用来标记资源，对资源分组，配合选择器，就可以让资源之间精准协作。一般以**键值对**的形式存在，如 app=demo。

如，在 demo app 下部署两个应用，一个是 api 服务，需要对外开放，另一个是 cronjob 不需要对外开放，则配置可以这样设置

```yaml
# service
Service
	selector:
		app: demo
		role: api
	
-----------------
# 因为 role: api 的匹配，service 的流量只会导到这个 pod
apiPod
	image：api:v1
	label:
		app: demo  # 起分组功能
		role: api  # 起标记功能
		
---------

cronPod
	image: api:v1
	label:
		app: demo
		role: cronjob
```

lable 进行匹配时，多个 lable 之间是 **与** 的关系

**也可以对 node 打标签，可以做到在指定的 node 部署 pod**



标签键命名规则：

1.   由前缀+名称组成，用 `/` 分割，如： `kubenetes.io/`。
2.   名称是必须的,len < 64，前缀省略表示该标签对用户私有
3.   名称组成：字母或数字开头结尾，`-._` 是允许的符号
4.   前缀若指定，必须是由 `.` 分割的一系列标签，后跟 `/` 以示结束

标签操作：

更多示例参考`kubectl lable --help`

```shell
# 给 pod test 加 prod=true 的标签
kubectl lable pods test prod=true
# 给 pod test 更新 prod 标签为 false（overwrite 不存在的 label 会报错）
kubectl lable pods test prod=true --overwrite
# 给 pod test 删除 prod 标签
kubectl label pods test prod-
```



#### 选择器

用来匹配资源，符合选择器规则的资源才会被调用或者使用。

```shell
# 查询资源的 label
kubectl get service --show-labels

# 使用选择器(-l 后跟选择器)
kubectl get svc -l 'app in (api, cronjob)' --show-labels
# 多维选择器(选择版本不是 1， 且 app = api 的 service)
kubectl get svc -l version!=1,'app in (api)' --show-labels
```

选择器运算符：

1.   分两类：**基于等值**，**基于集合**。多组选择器用`,` 分割，多组之间是 `&&` 的关系。需要注意的是：<u>否定选择会匹配键名的补集</u>
2.   **等值类**支持的运算符： `=`, `==`, `!=`。如：`tier != frontend` 会匹配**所有键名等于 `tire` 并且值不等于 `frontend`** 加上**所有键名不是 `tire`** 的资源
3.   集合类支持的运算符：`in`, `notin`, `exists`，并且可以只用在键名上，。如：
     1.   `env in (dev, prod)` 匹配所有键等于 `env` 并且值是 `dev` 或者 `prod` 的资源
     2.   `tire notin (frontend, backend)` 匹配**所有键等于 `tire` 并且值不等 `frontend`, `backend`** 加上**所有没有 `tire` 键**的资源。
     3.   `partition` 匹配所有键是 `partition` 的资源，而不管其值是什么
     4.   `!partition` 匹配所有键**不是** `partition` 的资源，而不管其值是什么



### 节点驱逐 pod

```shell
# 驱逐节点上的所有 pod, daemonset 类型的除外。并给这个节点打上污点，不会再被调度到
kubectl drain --ignore-daemonsets <节点名称>
```

可能需要一点时间，会等 pod 做完收尾工作才算结束



如果有被驱逐，但是没被删掉的 pod, 需要手动删除

`kubectl get pods -n <> --field-selector=status.phase=Failed | grep Evicted | awk '{print $1}' | xargs kubectl delete pod -n <>` **找出 Evicted 的节点并删除**

### 服务回滚、重启、扩缩容

-   回滚：
    -   回滚至上一版本：`kubectl -n local-test rollout undo deployment test-docker ` **优点：可在很紧急，并且明确回滚到上一个版本就可以解决问题的情况下，立即恢复服务，而不用管具体哪个版本**
    -   回滚至指定版本： `kubectl -n local-test rollout undo deployment test-docker --to-revision=<>` **缺点：需要明确知道每个版本号代表的功能**
-   重启：
    -   重启指定 pod。如果是以 deploy 的方式起的，可以直接 delete 这个 pod, deploy 会自动起一个新的 pod。delete deploy 的话，再重启就需要有配置文件了，相当于 stop & start, 而不是 restart。
    -   重启整个资源。 `kubectl rollout restart RESOURCE`。会自动 scale down 0 & scale up 到指定副本数
-   扩缩容：
    -   `kubectl scale --replicas=n RESOURCE`。适用于 deployment, replica set, replication controller, or stateful set 这些资源。



### 服务亲和性，反亲和性配置

https://kubernetes.io/zh-cn/docs/concepts/scheduling-eviction/assign-pod-node/

此配置可以让副本尽量分布在不同的 node 上，应对压力时，可以让集群内的每台服务器都最大化利用起来

```yaml
spec:
  affinity:
          podAntiAffinity:
            preferredDuringSchedulingIgnoredDuringExecution:
              - weight: 50
                podAffinityTerm:
                  labelSelector:
                    matchExpressions:
                      - key: app
                        operator: In
                        values:
                          - ${APP}
                  topologyKey: "kubernetes.io/hostname"
```





### 集群问题排查

-   `kubectl top node`， `kubectl -n <> pod [-A]` 查看 node 或者 pod 的 CPU， 内存占用情况。可以通过看服务对资源的要求，来设置 request
    -   `kubectl top node --sort-by [cpu|memory] [--sum]`



### 从运行中的资源生成配置文件

一般可以用来备份资源

```shell
kubectl get deployment my-deployment -o yaml > deployment.yaml
```

这样生成的配置文件， 可能包含一些不需要的信息，比如 status、metadata 中的 creationTimestamp、resourceVersion 等字段。如果你希望创建一个新的资源配置文件，可以手动移除这些字段。



### 拉取私有仓库镜像

```shell
kubectl -n <ns> create secret generic <secret-name> \
    --from-file=.dockerconfigjson=<path/to/.docker/config.json> \
    --type=kubernetes.io/dockerconfigjson
```

https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/



# 网站增加 HTTPS 支持

**最好是直接到 cloudflare 上购买域名。DNS， SSL 证书，一站式搞定，还免费**

域名购买：[godaddy.com](godaddy.com), 支持设置 DNS 转发。不过尽量使用二级域名指向服务器 IP，一级域名会被默认设置一些 DNS。



域名买好，并配好 DNS 后，还需要 SSL 证书，有付费的，也有免费的。这里以免费的 [let's Encrypto](https://letsencrypt.org/zh-cn/getting-started/) 为例

使用更简单的脚本： [acme.sh](https://github.com/acmesh-official/acme.sh) 按照说明，一步步执行就好

如果方便增加 DNS记录，则推荐使用 DNS 方式获取证书。(如果 DNS 服务商有API，则推荐使用 API，如 goDaddy 可以配置  `--dns dns_gd`)



```shell
acme.sh --issue -d <domain> --dns dns_gd
```



证书获取到之后，如果要集成到 k8s Ingress 上，则需要：

```shell
# 新增 secrets
kubectl create secret tls <tls-name> -n <namespce> \
--cert=<abslute path to crt file> \
--key=<abslute path to key file>

# 这里使用的 cert 文件最好是生成的 fullchain.cer ，否则一些应用会报链式验证失败
```

Ingress 配置这个 ssl

https://kubernetes.io/docs/concepts/services-networking/ingress/#tls

```yaml
Kind: Ingress
....
spec:
  ingressClassName: nginx
  tls:
  - hosts:
  		- "ssl host" (aa.bb.com) 需要和下面 rules 配的域名完全一致
    secretName: secret name
  rules:
  - host: aa.bb.com
```

### 验证

https://www.digicert.com/help/ 可以验证网站的证书是否正常

也可以用命令行：

```shell
openssl s_client -debug -connect <hostname>:443
```

使用非 fullchain 的证书，验证时会输出类似 "unable to verify the first certificate" 的异常（如果是 https 接口的话，可以用 postman 之类的工具发起请求，也能测出类似结果）

使用 fullchain 证书，输出的结果：

```text
Certificate chain
 0 s:CN=*.aa.com
   i:C=AT, O=ZeroSSL, CN=ZeroSSL ECC Domain Secure Site CA
   a:PKEY: id-ecPublicKey, 256 (bit); sigalg: ecdsa-with-SHA384
   v:NotBefore: Apr  3 00:00:00 2024 GMT; NotAfter: Jul  2 23:59:59 2024 GMT
 1 s:C=AT, O=ZeroSSL, CN=ZeroSSL ECC Domain Secure Site CA
   i:C=US, ST=New Jersey, L=Jersey City, O=The USERTRUST Network, CN=USERTrust ECC Certification Authority
   a:PKEY: id-ecPublicKey, 384 (bit); sigalg: ecdsa-with-SHA384
   v:NotBefore: Jan 30 00:00:00 2020 GMT; NotAfter: Jan 29 23:59:59 2030 GMT
 2 s:C=US, ST=New Jersey, L=Jersey City, O=The USERTRUST Network, CN=USERTrust ECC Certification Authority
   i:C=GB, ST=Greater Manchester, L=Salford, O=Comodo CA Limited, CN=AAA Certificate Services
   a:PKEY: id-ecPublicKey, 384 (bit); sigalg: RSA-SHA384
   v:NotBefore: Mar 12 00:00:00 2019 GMT; NotAfter: Dec 31 23:59:59 2028 GMT
```



自动更新 k8s 证书

```shell
#!/bin/bash

# 查看 secret 使用证书的 serial, 可以与更新后的证书文件进行对比
# openssl x509 -noout -serial -in <(kubectl -n ${namespace} get secret/${secret_name} -o jsonpath='{.data.tls\.crt}' | base64 -d)

domain=zrqsmcx.top
namespace=sdk-h5
cert_dir=/home/ubuntu/ingress/ssl
cert_file=${cert_dir}/fullchain.cer
key_file=${cert_dir}/${domain}.key
secret_name=zrqsmcx.top1

if [ "$(openssl x509 -noout -serial -in ${cert_file})" != "$(openssl x509 -noout -serial -in <(kubectl -n ${namespace} get secret/${secret_name} -o jsonpath='{.data.tls\.crt}' | base64 -d))" ]; then
    kubectl create secret tls ${secret_name} -n ${namespace} --cert=${cert_file} --key=${key_file} --dry-run=client -o yaml | kubectl apply -f -
    echo 'secret renew'
else
    echo 'no need renew secret'
fi
```





# Docker

[Docker 从入门到实践](https://yeasy.gitbook.io/docker_practice/repository/registry_auth)



## cheatsheet

https://www.runoob.com/docker/docker-command-manual.html

```sh
# list images
docker images

# run an image
docker run ...
# 以运行 mysql 为例
docker run -itd --name mysql-local -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
# i interactive
# t tty
# d detach run in background
# --privileged 给容器内的程序提升执行权限
# --rm 容器停止时删除
# --restart=always 容器启动失败时自动重启，也可以设置重启次数


# redis
docker run -itd --name redis-local -p 6379:6379 redis

# with env 
docker run -e k=v <name>

# 查看当前运行容器的状态
docker stats
```

#### image

```sh
# 查找镜像
docker search <image name>

# 拉取镜像
docker pull <image name>:<version|lastest> (不带版本，默认拉取最新的)

# 列出安装的所有镜像
docker images

# 删除镜像
docker rmi <image name>

docker tag source_image[:tag] target_image[:tag]

docker push <hub server>
```

#### container

```sh
docker container -h

docker container attach <id>

# 在容器内部执行
docker exec -it <id> <cmd>

# 查看容器信息
docker container inspect <id>

# 查看所有容器
docker ps -a

# 启动一个停止的容器 cid: 容器 id
docker start <cid>

# 查看容器的日志  docker logs --help
docker logs -f <cid>

# 查看容器的端口映射
docker port <cid>

# 查看容器中运行的进程情况
docker top <cid>
```

#### 🔲Dockerfile

https://www.runoob.com/docker/docker-dockerfile.html



https://yeasy.gitbook.io/docker_practice/  以下内容主要参考该文档

#### 管理数据

docker 中的数据管理主要有两种方式：

1.   数据卷（volumes）
2.   挂载主机目录（bind mounts）

##### 数据卷

数据卷是和容器分开的，可以独立于容器的生命周期，也可以挂载到多个容器上。

```shell
# 创建数据卷
docker volume create my-volumn
# 查看所有数据卷
docker volume ls
# 查看指定数据卷信息
docker volume inspect <volume_name>
# 给容器指定数据卷。使用 --mount 来将数据卷挂载到容器里(可一次挂载多个)
docker run -d --name web --mount source=my-volume,target=/usr/share/nginx/html nginx:alpine

# 删除（volume 独立于容器生命周期，容器删除不会删除 volume（除非 docker rm -v）, 未被引用的 volume 也不会主动删除）
docker volume rm my-volume
# 清理无用的 volume 来精简空间
docker volume prune
```

##### 挂载主机目录

就是将主机的目录或者文件挂载到容器里，使用方式：`--mount type=bind,source=<host absolute path>,target=<container path>[,readonly]`

一般用来测试，可以通过操作本地文件来达到操作容器中文件的目的。测试比较方便。一个场景是，如果容器内的 app 可以实时读取配置文件的内容的变更，则可以把主机上这个文件挂载到容器里，通过在主机上调整配置文件

#### 使用网络

容器使用网络主要是两种方式：

1.   外部访问容器
2.   容器互联

##### 外部访问容器

该种方式主要是通过设定**端口映射**来实现容器内部的网络应用访问网络。可以通过 1. `-P` 随机分配端口映射 2. `-p <host port>:<container port>` 来指定容器和主机的端口映射。

如：`docker run -d --name nginx -p 80:80 nginx:alpine` 可以将 NGINX 运行在容器内，并通过主机的 80 端口访问。

`-p` 的一般格式：

1.   `ip:hostPort:containerPort`。映射指定地址的指定端口到容器的端口。如 `127.0.0.1:80:80` 映射 127.0.0.1 的 80 端口到容器的 80 端口
2.   `ip::containerPort`。映射指定地址的所有端口到容器的端口
3.   `hostPort:containerPort`。映射本地所有接口的端口到容器的端口（**常用**）

`-p` 可以使用多次来映射多个端口

##### 容器互联

要实现容器之间网络互联，一般会在 run 的时候通过 `--link` 来链接容器。

```shell
docker run -itd --name cc --link c1
```

但更好的做法是将网络独立出来。使用 `docker network create` 创建一个网络，在将需要互联的容器都关联到同一个网络即可

```shell
# create network
docker network create ap_net

# 启动两个容器
docker run -itd --name ap0 alpine ash
docker run -itd --name ap1 alpine ash

# 在容器里 ping 另一个容器，会发现网络不通
ping ap1

# 将 ap0 ap1 加入网络 ap_net
docker network connect ap_net ap0
docker network connect ap_net ap1

# 再次 ping 发现可以 ping 通
```

```shell
docker network --help
Usage:  docker network COMMAND

Manage networks

Commands:
  connect     Connect a container to a network
  create      Create a network
  disconnect  Disconnect a container from a network
  inspect     Display detailed information on one or more networks
  ls          List networks
  prune       Remove all unused networks
  rm          Remove one or more networks
```

使用 `docker network inspect <network>` 可以查看这个网络的详情，包括有哪些容器在使用这个网络。

(实际上，最终还是通过宿主机的 **iptables** 来控制容器间的网络)



### docker hub

Local registry: `docker run -d -p 5000:5000 --restart=always --name local_registry registry:latest`

1.   Login: `docker login -u <username> -p <password>`
2.   docker.io
     1.   `docker tag local_image:tag_name username/repository_name:tag_name`
     2.   `docker push username/repository_name:tag_name`
3.   local registry
     1.   `docker tag local_image:tag_name registry_address/repository_name:tag_name`
     2.   `docker push registry_address/repository_name:tag_name`



## Docker Compose

`docker compose [command]`

支持同时运行多个容器。使用 `docker-compose.yaml` 文件定义项目以及服务。

概念：

-   服务：实际运行的容器
-   项目：多个服务组成。在 docker-compose.yaml 文件中声明



### 实验

1.   创建一个可以持续运行一段时间的脚本
2.   使用 Dockerfile 将这个脚本做成一个镜像
3.   使用 Docker Componse 将这个镜像运行成多个服务，并自动扩缩容



#### 持续运行的脚本

WORKDIR: ./docker

```shell
#!/bin/ash

echo 'hhhhh'

# 暂停，便于后面的操作
sleep 600

echo 'container done'
```

#### 将脚本做成镜像

WORKDIR: ./docker

```dockerfile
FROM alpine:3.16

ENV NAME=hello AGE=11

WORKDIR /app

COPY /app.sh .

RUN chmod +x app.sh

ENTRYPOINT [ "/app/app.sh" ]
```

```txt
├── docker
│   ├── Dockerfile
│   └── app.sh
├── docker-compose.yaml
```



#### Compose It

```yaml
version: "3"

services:
  ap1:
    build: ./docker
    networks: 
        - ap_net

  ap2:
    build: ./docker
    networks: 
        - ap_net
    depends_on:
    		- ap1

networks:
  ap_net:
```

#### Run

```shell
# start all service
$ docker compose up -d
[+] Running 3/3
 ✔ Network docker-starter_ap_net   Created
 ✔ Container docker-starter-ap1-1  Started
 ✔ Container docker-starter-ap2-1  Started
 
$ docker container ls -a
CONTAINER ID   IMAGE              COMMAND       CREATED          STATUS       PORTS     NAMES
f8e29fb8becd   docker-starter-ap2 "/app/app.sh" 6 seconds ago    Up 5 seconds           docker-starter-ap2-1
2ff6891de6ac   docker-starter-ap1 "/app/app.sh" 6 seconds ago    Up 5 seconds           docker-starter-ap1-1

# scala
$ docker compose up -d --scale ap1=2
[+] Running 3/3
 ✔ Container docker-starter-ap1-1  Running
 ✔ Container docker-starter-ap2-1  Running
 ✔ Container docker-starter-ap1-2  Started
 
 $ docker container ls -a
 CONTAINER ID   IMAGE              COMMAND          CREATED        STATUS       PORTS NAMES
cfe654dcdabc   docker-starter-ap1  "/app/app.sh"    1 minute ago   Up 2 minutes       docker-starter-ap1-2
f8e29fb8becd   docker-starter-ap2  "/app/app.sh"    2 minutes ago  Up 2 minutes       docker-starter-ap2-1
2ff6891de6ac   docker-starter-ap1  "/app/app.sh"    2 minutes ago  Up 2 minutes       docker-starter-ap1-1

# scala
$ docker compose up -d --scale ap1=1
[+] Running 2/2
 ✔ Container docker-starter-ap1-1  Running
 ✔ Container docker-starter-ap2-1  Running
 
 $ docker container ls -a
CONTAINER ID   IMAGE              COMMAND       CREATED          STATUS       PORTS     NAMES
f8e29fb8becd   docker-starter-ap2 "/app/app.sh" 6 seconds ago    Up 5 seconds           docker-starter-ap2-1
2ff6891de6ac   docker-starter-ap1 "/app/app.sh" 6 seconds ago    Up 5 seconds           docker-starter-ap1-1

# test container link
$ docker exec -it docker-starter-ap1-1 sh
/app # ping docker-starter-ap2-1
PING docker-starter-ap2-1 (192.168.107.3): 56 data bytes
64 bytes from 192.168.107.3: seq=0 ttl=64 time=0.394 ms
64 bytes from 192.168.107.3: seq=1 ttl=64 time=0.179 ms
^C
--- docker-starter-ap2-1 ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max = 0.179/0.286/0.394 ms
/app # 

# stop
$ docker compose down
[+] Running 3/3
 ✔ Container docker-starter-ap2-1  Removed
 ✔ Container docker-starter-ap1-1  Removed
 ✔ Network docker-starter_ap_net   Removed
```



## Docker Swarm

类似 k8s 的集群管理与编排工具。

**节点**：运行 docker 的**宿主机**被看作是 docker swarm 的一个节点。节点分为**管理节点**（manager）和**工作节点**（worker）。docker swarm 命令基本只能在管理节点执行。可以有多个管理节点，但只会有一个 leader, 通过 raft 协议选举。

**服务和任务**： 任务是 swarm 中最小的调度单位，也就是 docker 的容器。服务是一系列任务的集合。

建议直接使用 k8s。



# 缓存 & Db

带缓存的数据库操作。

1.   在写入时，先写入数据库，再写入缓存
2.   在删除时，先删除数据库，再删除缓存
3.   查询，则是先查询缓存，不中再查询数据库

以 go-zero/cachedsql.go 中的方法为例：

```go
// ExecCtx runs given exec on given keys, and returns execution result.
func (cc CachedConn) ExecCtx(ctx context.Context, exec ExecCtxFn, keys ...string) (
	sql.Result, error) {
  // 先执行数据库操作
	res, err := exec(ctx, cc.db)
	if err != nil {
		return nil, err
	}

  // 再执行缓存操作。
	if err := cc.DelCacheCtx(ctx, keys...); err != nil {
		return nil, err
	}

	return res, nil
}
```

如果顺序反了：

1.   写入的情况下。如果写入数据库失败，还要再回滚缓存，并且，如果有线程读到了缓存，就相当于读取到了一笔不存在的记录
2.   删除情况下。如果缓存先被清除，但还没来得及写数据库，此时有线程读取，肯定是读取到数据库中的旧记录，但是这个数据是即将被删除的，所以也发生脏读



### MySQL



常用命令：

```sql
# 删除 unique key
alter table <table> drop index <index_name>
# 修改 unique key
# 先删除，再添加
alter table <table> add unique key <name>(fields...)
```





#### 联合索引

形式： `index <索引名> on <表名> (col_1, col_2, col_3，……)`

索引名，需要见名知意，如这个索引是给某个功能加的，就可以直接用功能描述来命名

多列组成联合索引，一般用于比较固定的查询，如 ETL ，报表等 SQL 固定，执行频繁的场景。

因为联合索引生效的前提是，`(col_1, col_2, col_3， ……)` 需要**从左到右**依次命中才能使用完整的索引，中间任何一个未命中都会停止走索引匹配。同时，遇到范围查询（`>, <, between, like`）也会停止匹配。

索引除了对 `where`  子句里的过滤条件生效，也会对分组条件生效。



#### 日志清理

```shell
# 清理当前所有的 binlog。也可以指定时间
PURGE BINARY LOGS BEFORE now();
```



#### mysql 问题排查

##### 查看当前在执行的事务

`SELECT trx_mysql_thread_id, trx_state, trx_query, trx_requested_lock_id, trx_tables_locked, trx_rows_locked, trx_isolation_level, trx_started FROM INFORMATION_SCHEMA.INNODB_TRX;` 

-   可以知道事务状态 trx_state
-   开启时间 trx_started
-   执行的 SQL trx_query
-   锁 ID trx_requested_lock_id
-   锁表情况 trx_tables_locked
-   行锁情况 trx_rows_locked
-   事务隔离级别 trx_isolation_level
-   mysql 线程 ID trx_mysql_thread_id

##### 查看当前打开的表

`SHOW OPEN TABLES` 

如果 In_use =1 表示此表当前有锁



#### 数据处理

-   插入更新

`insert ignore into <table> (<columns>) values ... on duplicate key update field1=values(fields), fields2=4 `  将指定记录插入表中, 如果主键不存在, 则插入, 否则, 按照 update 后的定义, 更新指定字段. 字段值可以使用 `values(xxx)` 引用更新的值. 

`ignore` 用来忽略错误, 如忽略主键重复错误. 

-   rowAffected

 https://dev.mysql.com/doc/c-api/8.4/en/mysql-affected-rows.html 这个官方文档, 说明了一条 `insert`, `update`, 或是  `delete` 语句实际影响的行数, 以及在不同情况下, "影响" 的定义. 

-   -   update: 返回的是实际发生改变的行
    -   `replace` : 如果一行记录替换了一条旧的,那返回的值是 2. 替换的逻辑是删除 + 插入,所以是 2 行
    -   `insert ... on duplicate key update`:  插入一条会记作 1, 更新逻辑同 `replace` . 所以想知道这条语句实际影响的行数,应该分开 `insert  ingore into` + `insert ignore into ... on duplicate key update` , 这样, 更新的实际行数才不会混入到插入的行数中. 





# 🔲Frp

https://sspai.com/post/52523





# Golang 异步确认 & 超时控制

```go
package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RouterRecordReq struct {
	RouterId    string
	LineId      string
	CheckUnbind bool
}

var checkInterval = 10 * time.Second
var timeOut = 18 * checkInterval

func (l *RouterAsyncLogic) RecordRouterAsync(in *RouterRecordReq) {
	go func() {
		timeOutChan := time.After(timeOut)
		doRecord := false
		for {
			time.Sleep(checkInterval)
			// 超时，或者 checkRouter 返回 true, 就停止循环
			done := l.keepCheckUntil(timeOutChan, func() bool {
				checked, err := l.checkRouter(in.RouterId, in.LineId, in.CheckUnbind)
				// 执行出错或者返回 true，都不再循环
				if err != nil || checked {
					if checked {
						// 返回 true , 则记录历史
						doRecord = true
					}
					return true
				}
				return false
			})
			if done {
				break
			}
		}
		op := utils.Ternary(in.CheckUnbind, "解绑", "绑定")
		// 记录路由器操作历史之前，先确定"绑定/解绑"操作成功再记录历史
		if doRecord {
			operationType := utils.Ternary(in.CheckUnbind,
				resourceserver.RouterOperation_UNBIND,
				resourceserver.RouterOperation_BIND)
			_, _ = NewRouterHistoryLogic(l.ctx, l.svcCtx).RouterHistory(&resourceserver.RouterHistoryRequest{
				OperationType: operationType,
				RouterId:      in.RouterId,
			})
			l.Logger.Infof("%s 路由器操作历史, routerId: %s", op, in.RouterId)
		} else {
			l.Logger.Errorf("%s 路由器操作历史失败, routerId: %s", op, in.RouterId)
		}
	}()
}

func (l *RouterAsyncLogic) keepCheckUntil(timeOutChan <-chan time.Time, predictor func() bool) bool {
	select {
	case <-timeOutChan:
		return true
	case <-time.After(500 * time.Millisecond):
		return predictor()
	}
}

// checkRouter 检查 router 的状态。checkUnbind 表示是否检查 router 的未绑定状态；否则就检查绑定状态
// 返回 bool, 表示是否满足指定的状态
// 在绑定的过程中，可能会再次被绑定或者解绑。
// 暂时在页面控制二者必须顺序发生。单个重复发生时，并不影响记录操作历史的准确性
func (l *RouterAsyncLogic) checkRouter(routerId, lineId string, checkUnbind bool) (bool, error) {
	router, err := l.svcCtx.PathliveRpc.GetRouter(l.ctx, &pathlive.GetRouterRequest{
		RouterId: routerId,
	})
	if err != nil {
		return false, err
	}
	r := router.Router
	if checkUnbind {
		// 当前是否是解绑状态
		return r.LineInfo == nil || r.LineInfo.LineId == "", nil
	} else {
		return r.LineInfo != nil && r.LineInfo.LineId == lineId, nil
	}
}
```



# TikTok 简易爬虫实现

tiktok web 页面，为各种爬虫准备了一份数据，就是其页面源码中，一个 id 为 `SIGI_STATE` 的 script 里的 json 数据。实际上，tiktok web 页面使用 sigi 框架，并且配合 SSR 将 sigi 应用的 state 保存在了 dom 里，相当于 vue 的 data。这个 state 里包含了用户的相关信息，用户发布的视频等等信息。

所以需要做的就是拉取 web 页面，解析出这个 json, 并且获取感兴趣的字段。

第一步，访问 tk 页面。tk 是限制了访问区域的，比如国内以及想干的大部分 ip 都不能够访问。所以第一步就是需要有一台能够访问 tk 的机器。

第二步，在这台机器上使用 curl 访问 tk 主页

第三步，从 html 页面中解析出 json

第四步，从 JSON 中提取感兴趣的字段

### 代码实现

一：跳板机。因为 tk 对访问的区域敏感，所以准备了多个区域的多台机器备用。查询时，可以选择发起访问的区域

```go
// region: ip
var TkDestIpMap = map[string][string]{}
```

使用 ssh 工具，连接到指定区域，并执行命令。这里写一个简易的 ssh 工具

```go
// 现在太菜，后面补
```

连接上之后，就可以执行命令了。但是因为命令比较多，而且使用 shell 编写也比较麻烦，所以使用 golang 编写，再打包成可执行命令，然后只需要触发一下就可以了。



二：获取 HTML 并提取感兴趣的字段

编写一个 golang 命令行工具

```go
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var tkId = flag.String("t", "", "tk user id")

// 使用一下命令，将此 go 程序编译成可执行程序（这里编译后的可执行程序名为 fetch。 使用方式为 ./fetch -t <tk user id>）
// build: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./fetch ./parseTkState.go
func main() {
	flag.Parse()

	sw := bytes.Buffer{}
	command := exec.Command("curl", "-s", "https://www.tiktok.com/@"+*tkId)
	command.Stdout = &sw
	err := command.Run()
	if err != nil {
		_ = fmt.Errorf("curl tk failed %v", err)
		return
	}

  // 使用 curl 获取到的 tk 主页 html
	html := sw.String()
	script, err := extraJsonInScript(html)
	if err != nil {
		_ = fmt.Errorf("extra json failed %v", err)
		return
	}
  // 将结果输出到 stdout，便于调用者获取
	fmt.Print(script)
}

// 从 html 中解析出含有用户信息的 json
var scriptsReg = regexp.MustCompile(`<script\s+.*?>(.*?)</script>`)
func extraJsonInScript(html string) (string, error) {
	ret := ""
	rets := scriptsReg.FindAllStringSubmatch(html, -1)
	for _, v := range rets {
		isStateScript := strings.Contains(v[0], `id="SIGI_STATE"`)
		if len(v) > 1 && isStateScript && json.Valid([]byte(v[1])) {
			state := TKState{}
			jsonStr := strings.Trim(v[1], " ")
			err := json.Unmarshal([]byte(jsonStr), &state)
			if err != nil {
				continue
			}
			userMap := state.UserModule.Users
			if userMap == nil || len(userMap) == 0 {
				return "", errors.New("用户不存在或账号已注销")
			}
			// 返回什么内容，由 TKState 决定
			stateStr, err := json.Marshal(state)
			if err != nil {
				continue
			}
			ret = string(stateStr)
			break
		}
	}
	return ret, nil
}

// 省略这个结构体的内容。具体内容可以手动把 tk 主页的 json 拉出来看，并且使用工具转换成结构体即可
type TKState struct {
}
```

有了 fetch 这个可执行程序，调用方就很简单了。

```go
sh := NewSSHHelper(sshConf)
cmd := fmt.Sprintf("./fetch -t %s", url.PathEscape(tiktokId))
Logger.Infof("show tk user cmd: %s", cmd)
sshRet, err := sh.RunCMD(cmd)
if err != nil {
  l.Logger.Errorf("curl failed, %v", err)
}
return sshRet
```

但是因为访问 tk 主页是个网络请求行为，所以不得不考虑超时问题。以下是处理超时的逻辑：

```go
sh := NewSSHHelper(sshConf)
// 接收 fetch 命令结构的 channel
retChan := make(chan string, 1)
// 异步执行，让 main 进入 select 流程 
go func() {
  cmd := fmt.Sprintf("./fetch -t %s", url.PathEscape(tiktokId))
  Logger.Infof("show tk user cmd: %s", cmd)
  sshRet, err := sh.RunCMD(cmd)
  if err != nil {
    l.Logger.Errorf("curl failed, %v", err)
    // 出错了写入空值，后面会判断
    retChan <- ""
  }
  retChan <- sshRet
}()

// 经典的 golang 超时控制结构
select {
  case <-time.After(45 * time.Second):
  	return "", status.New(codes.DeadlineExceeded, "超时").Err()
  case sshRet := <-retChan:
    if sshRet == "" {
      return "", status.New(codes.Unknown, "解析失败").Err()
    } else {
      return sshRet, nil
    }
}
```

至此，一个简易的 tk 爬虫便能够跑起来了。

但是需要注意的是，fetch 程序是运行在能够访问 tk 的机器上的，而 fetch 程序的调用者，需要通过 ssh 连接到这台机器上去触发。并且，因为有很多区域，每个区域都有一个主机，所以 fetch 程序的部署也是一个繁琐的事情。

一开始写了一个 shell 脚本，循环所有的机器列表，一个个通过 scp 把编译好的 fetch 程序部署上去。虽然也能用，但是由于机器数量巨大，机器分布在全球，访问时间长短不一，脚本又不能并行，所以就执行得很慢。

后经同事指点，了解了 ansible 这个工具。那是真好用。

所以现在的部署脚本就是：

```shell
# build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./fetch ./parseTkState.go
echo 'build done'
sleep 2

# deploy
ansible tk -i ./ansible.ini -m copy -a "src=./fetch dest=~/"
echo 'copy done'

ansible tk -i ./ansible.ini -m file -a "path=/root/fetch mode=0755"
echo 'deploy done'
```

ansible.ini 就是配置机器列表，大致长这样：

```ini
[tk]
hostname1 ansible_password=yyy
hostname2
...
[tk:vars]
ansible_connection=ssh
ansible_user=root
ansible_password=xxx
```



# 前端 & js

## 通过 CSS 给 HTML 元素加水印

HTML 结构：

```html
<div class="container">
    <div :class="'ad' + (i || '') + ' ad'" v-for="(ad, i) in 3">
        {{watermark}}
    </div>
    ...
</div>
```

CSS:

```scss
.container {
  // 重要
  position: relative;

  --marker-right: 40%;
  --marker-top: 50%;

  .ad1 {
    --marker-right: 60%;
    --marker-top: 30%;
  }

  .ad2 {
    --marker-right: 20%;
    --marker-top: 70%;
  }

  .ad {
    font-weight: bold;
    text-align: center;
    width: 300px;
    // 重要
    position: absolute;
    right: var(--marker-right);
    top: var(--marker-top);
    opacity: 0.15;
    rotate: -36deg;
    user-select: none;
    overflow: hidden;
    pointer-events: none;
  }
}
```



## js 实现限流

```ts
async function ratelimiter<R>(fns: Array<() => Promise<R>>, cap: number = 10) {
    let cut = []
    let lastIndex = 0
    while ((cut = fns.slice(lastIndex, lastIndex + cap)).length) {
        lastIndex += cut.length
      	// 等这一批请求全部执行完再执行下一批，也可以增加间隔停留
        await Promise.all(cut.map(c => c()))
        console.log("done batch... ", cut.length);
    }
}

// 使用
await ratelimiter(tableList.map(row => () => doHardWork(row)), 10)

async function doHardWork(row) {
  // 耗时的请求
}
```

## JS/TS Map 对象序列化与反序列化

JS  中的 Map 对象，直接使用 `JSON.stringify` 序列化时，不能正确按预想中的变成 `{}` 对象的形式，所以需要特殊处理。

序列化函数

```ts
// 将对象中 Map 对象序列化成带标记的对象
function mapStringifyReplacer(key: any, value: any) {
    if (value instanceof Map) {
        return {
            dataType: "Map",
            value: Array.from(value.entries())
        }
    } else {
        return value
    }
}
JSON.stringify(obj, mapStringifyReplacer)
// {a: [1,2]} => {a: {dataType: 'Map', value: ['a', [1,2]]}}
```

反序列化：

```ts
function mapParseReceiver(key: any, value: any) {
    if (typeof value === 'object' && value !== null) {
        if (value.dataType === 'Map') {
            return new Map(value.value);
        }
    }
    return value;
}
JSON.parse(str, mapParseReceiver)
// {a: {dataType: 'Map', value: ['a', [1,2]]}} => map 对象
```

若不需要反序列化，反而是要序列化成对象的格式，方便和外部系统协作，则序列化函数可以这样写：

```ts
function mapStringifyToObjReplacer(key: any, value: any) {
    if (value instanceof Map) {
        let obj = {}
        for (const iterator of value.entries()) {
            // @ts-ignore
            obj[iterator[0]] = iterator[1]
        }
        return obj
    } else {
        return value
    }
}
```

## JS 中将生成器生成的 Promise 同步化

参考： [**Javascript 中通过 yield 和 promise 使异步变同步**](https://blog.51cto.com/u_15283585/2957703)

生成器教程： https://zh.javascript.info/generators

核心原理： 利用递归对生成器进行迭代，每次**遇见 Promise 则在 then 中再次递归**，则可以保证下次执行一定是在 promise resolve 之后的。

```ts
function awaitable(gen: Generator) {
    const item = gen.next()
    const { value, done } = item
    if (value instanceof Promise) {
        value.then(() => awaitable(gen))
    } else {
      	// 其他类型的值，正常迭代
        awaitable(gen)
    }
    if (done) {
        return item.value
    }
}
```

可以再根据需求调整此函数





# Vim

https://coolshell.cn/articles/5426.html

实验的时候最好是找一些代码来操作

### 命令模式相关命令：

#### 编辑相关：

-   x -> 删除当前光标所在的一个字符
-   dd -> 删除当前行，并存入到剪切板。就相当于剪切功能 （dd 中间可加数字，表示要剪切的行数）
-   p -> 粘贴剪切板的内容到当前行
-   y -> 复制当前行
-   gu -> 变小写，gU 变大写。开启大小写转换选择，后面跟光标移动操作来确定选择哪些字符进行大小写转换。如：gUe -> 将当前光标到单词结尾的字符变大写
-   Ctrl V -> 块选择模式。
    -   Ctrl V -> 进入块选择模式
    -   jhkl, ^$，上下左右 等等方式进行选择
    -   I -> 插入
    -   ESC 退出模式，并将输入的内容应用到选择的行

##### 配合光标移动

-   ye -> 从当前光标复制到当前单词结尾。同样的，w, b,W, B 都可以结合来复制单词中的一部分

##### 重复

-   . -> 重复上次的命令
-   N<command> -> 重复某个命令 N 次
-   100iabcd [ESC] -> 插入 `abcd ` 100 次

#### 插入模式：

-   a -> 在光标处插入（在光标处切换到插入模式）
-   A -> 在行尾插入
-   o -> 在当前行后插入
-   O -> 在当前行前插入
-   cw -> 替换从光标到单词结尾的所有字符（删除当前位置到单词结尾的所有字符，并切换到插入模式）

#### 光标移动：

-   0 -> 到行首（不论开头是不是 blank）
-   $ -> 到行尾（不论开头是不是 blank）
-   ^ -> 到行首（到第一个不是 blank 的字符）
-   g_ -> 到行尾（到第最后一个不是 blank 的字符）
-   G -> 到文件尾， gg -> 到文件开头
-   <N>G -> 直接跳到第N行开头，也可以直接 `:N` 。gg 到第一行，G 到最后一行
-   w -> 移动到下一个单词的开头
-   e -> 移动到下一个单词的结尾
-   b -> 移动到上一个单词的开头
    -   注：`w/e/b` 对应的 `W/E/B` 功能相似，只是大写的会认为单词是用 blank 分割的，小写的会认为单词是由字母数字下划线组成的（其他符号则认为是单词的分割符）
-   f<char> -> 在本行内，移动到下一个 char 的位置 (find)
-   t<char> -> 在本行内，移动到下一个 char 之前的位置（till）
-   % -> 在括号的开闭符号间移动，支持 ( [ {    (需要先把光标移动到其中一个括号上)
-   `*/#` -> 匹配光标所在的单词，并移动到上/下一个匹配的单词，`*` 是下一个，`#`是上一个。（实际上是用搜索实现的，只不过会自动匹配查找关键词，查找一次后，用 n/N 都可以继续查找）
-   Ctrl O -> 回到光标的上一个位置，可以无视文件、tab 页、窗口
-   Ctrl I -> 回到光标的下一个位置，可以无视文件、tab 页、窗口

#### Undo/redo：

-   u -> undo
-   Ctrl + r -> redo

#### 文件操作：

-   :e <path/to/file> -> 打开新的文件
-   :bn / :bp -> 上下切换打开的多个文件
-   :w -> 文件编辑后存盘， 若后带文件路径，则会保存到指定文件名
-   :saveas <path/to/file> -> 另存为
-   :x / :q -> 退出

#### 宏录制

-   qa -> 开启录制宏 a (可以是其他名字)， q -> 结束录制
-   @a -> 回放宏 a
-   @@ -> 回放最近创建的一个宏，前面可叠加数量来重复操作

例子：

在一个只有一行 1 的文本中进行如下操作：

1.   qa 开启录制
2.   Yp 复制粘贴一行
3.   Ctrl a 将当前行数 + 1
4.   q 结束录制
5.   100@a 将创建 102 行按顺序排好的数字

#### 可视化选择

开启：

-   v -> 可视化选择
-   V -> 可视化行选择
-   Ctrl v -> 可视化块选择

可视化选择的操作：

-   J -> 把所选的行 join 起来
-   <, > -> 进行增减缩进
-   = -> 自动缩进（目前还不明确缩进的规则是啥，缩进出来不好看）

#### 执行 shell 命令

-   :r!<command> -> 将 command 在 shell 中的执行结果读取并写入到当前位置
-   :pwd -> 展示当前工作目录

#### 分屏

-   :split -> 开启横向分屏
-   :vsplit -> 开启纵向分屏
-   :sp filename -> 上下分割，并打开一个新文件
-   :vsp filename -> 左右分割，并打开一个新文件
-   :q -> 关闭当前分屏
-   Ctrl w -> 窗口操作
    -   hjkl -> 上下左右选择窗口(方向键也可) (对应大写的功能是移动窗口，此时不能用方向键了)
    -   _ -> 横向最大化窗口
    -   | -> 纵向最大化窗口
    -   = -> 所有窗口尺寸一样
    -   `-` -> 横向减小尺寸
    -   `+` -> 横向增加尺寸
    -   c -> 关闭当前窗口, 如果只剩最后一个窗口，则关闭失败
    -   q -> 如果只剩最后一个分屏，则关闭 VIM
    -   s -> 上下分割当前文件
    -   v -> 左右分割当前文件

用 vim 打开多个文件并分屏展示：

-   vim -on file1 file2 -> 水平分屏
-   vim -On file1 file2 -> 垂直分屏

n 是数量，可以指定，也可以写 n 来更加文件数量自动判断

如：vim -On file1 -> 打开一个文件，不分屏

### 插入模式相关命令：

-   Ctrl n/ Ctrl p -> 自动提示。输入单词开头，会出现候选





# Loki on Grafana 记录

参考 grafana 文档进行配置，grafana 配置面板，数据源选择 loki ， 发现需要填一个 url，即 loki 的服务地址。先运行 loki 服务。

https://medium.com/@amolbansal1234/how-to-install-loki-and-grafana-in-kubernetes-cluster-through-helm-chart-dae514d7f1c

https://ezeugwagerrard.com/blog/Deploy-A-Scalable-Loki-Instance-To-Kubernetes-Via-Helm

###  安装 helm （如果没有的话）



### 通过 Helm 安装 loki chart 

>    loki-stack 虽然已经不再维护了，但还是最简单的使用方式

1.   添加源 `helm repo add grafana https://grafana.github.io/helm-charts`
2.   `helm repo update`
3.   在 k8s 集群上部署 `helm upgrade --install loki --namespace=loki-stack grafana/loki-stack`

部署完成后，`kubectl -n loki-stack get all` 可以看到 helm 自动运行了一下服务：

```shell
$ kubectl -n loki-stack get all
NAME                      READY   STATUS    RESTARTS   AGE
pod/loki-0                1/1     Running   0          26m
pod/loki-promtail-24v89   1/1     Running   0          26m
pod/loki-promtail-4lgmz   1/1     Running   0          26m
pod/loki-promtail-g7xql   1/1     Running   0          26m
pod/loki-promtail-j4crf   1/1     Running   0          26m
pod/loki-promtail-mtxx2   1/1     Running   0          26m
pod/loki-promtail-pbb89   1/1     Running   0          26m
pod/loki-promtail-qxmvh   1/1     Running   0          26m
pod/loki-promtail-trdzk   1/1     Running   0          26m

NAME                      TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/loki              ClusterIP   172.17.220.240   <none>        3100/TCP   26m
service/loki-headless     ClusterIP   None             <none>        3100/TCP   26m
service/loki-memberlist   ClusterIP   None             <none>        7946/TCP   26m

NAME                           DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/loki-promtail   8         8         8       8            8           <none>          26m

NAME                    READY   AGE
statefulset.apps/loki   1/1     26m
```

有收集 pod 日志并发往 loki 的 promtail，对对外提供接口的 service。

### grafana 配置 loki 作为数据源

loki 的 url 可以通过 `kubectl -n loki-stack get svc/loki` 查看，或者上面 `get all` 时也输出了，是 `http://loki.loki-stack:3100`，配置到 grafana 上，就可以进行查询了。

### 自定义 helm chart 里 loki 的配置

Values.yaml

```yaml
loki:
  enabled: true
  # 开启本地持久化，会自动创建 pvc
  persistence:
    enabled: true
    size: 20Gi
    storageClassName: csi-udisk-rssd # pvc class name
  isDefault: true
  url: http://{{(include "loki.serviceName" .)}}:{{ .Values.loki.service.port }}
  readinessProbe:
    httpGet:
      path: /ready
      port: http-metrics
    initialDelaySeconds: 45
  livenessProbe:
    httpGet:
      path: /ready
      port: http-metrics
    initialDelaySeconds: 45
  datasource:
    jsonData: "{}"
    uid: ""
  # 直接在这里写 loki 的各种配置
  auth_enabled: false
  chunk_store_config:
    max_look_back_period: 0s
  compactor:
    shared_store: filesystem
    working_directory: /data/loki/boltdb-shipper-compactor
    compaction_interval: 30m
    retention_enabled: true
    retention_delete_delay: 12h
    retention_delete_worker_count: 50
    delete_request_store: filesystem
  ingester:
    chunk_block_size: 262144
    chunk_idle_period: 15m
    chunk_retain_period: 1m
    lifecycler:
      ring:
        replication_factor: 1
    max_transfer_retries: 0
    wal:
      dir: /data/loki/wal
  limits_config:
    retention_period: 72h
    enforce_metric_name: false
    max_entries_limit_per_query: 5000
    reject_old_samples: true
    reject_old_samples_max_age: 168h
  memberlist:
    join_members:
    - 'loki-memberlist'
  schema_config:
    configs:
    - from: "2020-10-24"
      index:
        period: 24h
        prefix: index_
      object_store: filesystem
      schema: v11
      store: boltdb-shipper
  server:
    grpc_listen_port: 9095
    http_listen_port: 3100
  storage_config:
    boltdb_shipper:
      active_index_directory: /data/loki/boltdb-shipper-active
      cache_location: /data/loki/boltdb-shipper-cache
      cache_ttl: 24h
      shared_store: filesystem
    filesystem:
      directory: /data/loki/chunks
  table_manager:
    retention_deletes_enabled: true
    retention_period: 336h

promtail:
  enabled: true
  config:
    logLevel: info
    serverPort: 3101
    clients:
      - url: http://{{ .Release.Name }}:3100/loki/api/v1/push

grafana:
  enabled: false
  sidecar:
    datasources:
      label: ""
      labelValue: ""
      enabled: true
      maxLines: 1000
  image:
    tag: 10.3.3

prometheus:
  enabled: false
  isDefault: false
  url: http://{{ include "prometheus.fullname" .}}:{{ .Values.prometheus.server.service.servicePort }}{{ .Values.prometheus.server.prefixURL }}
  datasource:
    jsonData: "{}"
```

使用配置部署 loki :

```shell
helm upgrade --install loki --namespace loki-stack --values loki-values.yaml grafana/loki-stack
```

验证 pvc 是否起作用：

```shell
# loki 默认会花在 emptyDir 作为存储，所以重启后，之前的日志会丢失。如果 pvc 生效的话，重启后日志还在
helm uninstall loki --namespace loki-stack
```



# 限流控制

需求： API 接口对发来的请求做限制。根据请求带的唯一性 ID 作为区分，假设每个 ID 每 24H 内最多只允许通过三次。

初始想法：ID 第一次请求来的时候，ID 作为 key，允许次数作为 value 存到 redis，并设置过期时间为 24H，每发来一个请求，如果通过，则 value 加一。value 大于等于 3 时拒绝请求。直到这个 ID key 过期。

这种方法存在的问题是，在极端情况下，短时间内会出现最大 5 次的请求被通过。这显然是不符合预期的。

>   第一个 24H |O-------------------O-O|
>
>   第二个 24H |O-O-O------------------|



知道一种由 TCP 滑动窗口演变而来，比较适合这种场景的方法。基于 redis zset 实现，大致步骤如下：

使用带有效期的 zset 存储允许通过的次数，score 是请求发生的时间戳

1.   `zremrangebyscore` 清除有效期外的值。`zremrangebyscore key -inf now-<limit seconds>`。这一步可以保证时间窗口移动之后，有效期外的次数不会限制新时间窗口下的次数
2.   `zcard key` 判断记录的次数是否达到限制。这里获取到的次数，是在有效期内通过的请求次数。
3.   如果请求被通过。则使用 `zadd key now nowString` 记录，并更新 key 的过期时间为新的 24H `expire key <limit seconds>`。更新 key 的过期时间为新的 24H 表示时间窗口的滑动

大致示意图

>   第一次请求 |O----------------------|
>
>   第二次请求       |O----------------------|
>
>   第三次请求               |O----------------------|
>
>   第四次请求: 因为次数达到 3 ，被拦截
>
>   第五次请求    |-----------------------||O----------------------|     请求发生在第一次请求的超时时间之外，但由于每次第一步先移除有效期外的次数，此时第一次请求记录的数被清除，在当前时间段内，已通过的请求只有两次，所以第五次请求也被允许通过





# Android

[安卓基础概念及相关源码](https://github.com/jeanboydev/Android-ReadTheFuckingSourceCode/blob/master/article/android/basic/01_activity.md)

[ADB 相关命令](https://adbshell.com/commands/adb-shell-pm-list-packages)



问题排查：
```shell
# 查看 APP 的内存占用情况
dumpsys meminfo <package name>

# 查看系统中的 activity
dumpsys activity activities

# 查看系统日志
logcat
```



# Git

ubuntu 上更新 Git 版本

```shell
sudo add-apt-repository -y ppa:git-core/ppa
sudo apt-get update
sudo apt-get install git -y
```



# Nginx

相关命令：

```shell
# 启动
systemctl start nginx
# 重载配置
sytemctl reload nginx
# 停止
systemctl stop nginx
# 验证配置文件
nginx -t -c <file>
```

### 配置相关

[nginx 在线配置](https://nginxconfig.org/)

一般 nginx 会使用默认的 `/etc/nginx/nginx.conf` 作为配置文件，这个文件里定义了 `http` 等顶级模块，并且会引入 `/etc/nginx/conf.d/*.conf` 所以可以把业务相关的 `server` 配置放到这个目录下，以 `.conf` 结尾。



另外需要注意的是， `/etc/nginx/nginx.conf`  使用的默认用户是 `www-data` 可能对某些文件，或者反向代理时的配置读取不到。可以视情况而定换成 root。



#### 反向代理配置

```conf
server {
	listen 80;
	listen 443 ssl; # HTTPS 支持

	# SSL 证书
	ssl_certificate      /root/deploys/apps/nginx/ssl/id.pem;
  ssl_certificate_key  /root/deploys/apps/nginx/ssl/id.key;

	# 监听这个域名的请求
	server_name admin.test.com;
	access_log /var/log/nginx/access.ops.log;
	error_log /var/log/nginx/err.ops.log;		

	# 路径配置
	location / {
		proxy_redirect off;
		
		# 反向代理
		proxy_pass http://localhost:3002;
	}
	
	location /api/ {
		proxy_redirect off;
		proxy_pass http://localhost:8888;
		proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
	}
}

server {
	listen 80;
	listen 443 ssl;

	ssl_certificate      /root/deploys/apps/nginx/ssl/id.pem;
	ssl_certificate_key  /root/deploys/apps/nginx/ssl/id.key;

	server_name web.test.com;
	access_log /var/log/nginx/access.ssp.log;
	error_log /var/log/nginx/err.ssp.log;

	# gzip
	gzip            on;
	gzip_vary       on;
	gzip_proxied    any;
	gzip_comp_level 6;
	gzip_types      text/plain text/css text/xml application/json application/javascript application/rss+xml application/atom+xml image/svg+xml;
	
	location / {
		proxy_pass http://localhost:3001;
	}
}
```



