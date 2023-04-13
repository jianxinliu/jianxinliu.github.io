# Biz


## 代码仓库结构

console-source  前端

phone-business 后端

resource 资源层（rpc 调用者）

operation-console 运营平台前端

turbolive-server 直播快服务端

pathlive-rpc 直播快服务端（新）


资源层调用底层 xxRpc 实现


## 使用到的库

- go-zero Go 微服务
- grpc


## misc

github cicd pipeline 

k8s 构建

docker dockerfile 

make makefile


## phone-business(比较通用的构建流程)

### build

使用 docker 进行构建

1. 读取 /deploy/templates 下的配置文件。分为 preview & production
2. 将配置文件中的配置信息设置为环境变量
3. 使用 `envsubst` 命令将配置文件中的参数使用环境变量替换到 `config.tmpl` 中（`config.tmpl` 是配置的模板，只需要往里面填充对应的参数即可）
4. 将替换结果写入到 `phone-api.yaml` 下。（最终生效的配置文件是这个）

### deploy

使用 github pipeline 进行部署

使用的配置文件： `.github/workflows`

1. checkout: 拉取最新代码
2. 生成 image tag(git 提交的相关信息 + 时间)
3. 构建 docker 镜像 & push
4. 将 `/deploy/templates/deploy.tmpl` 中的变量替换为环境变量中的值(主要是 pipeline 中 xxx-deploy.yaml 中定义的变量)，写入到 `deploy.yaml` 中，作为 k8s 的部署配置 
5. 部署到 k8s

#### api

提供的 http 接口，面向前端页面，实际的逻辑还是通过 rpc 调用 

#### rpc

phone-business 的 rpc 调用端

在进行一些业务逻辑的处理之后，调用 xxBackend rpc


## tubrolive-server

使用 gin 编写的 web 服务，后迁移到 pathlive-rpc ,使用 go-zero 

## resource 资源层

go-zero 开发模式



# Tech




## linux 命令

### envsubst

env-substitute: 将文件中的变量使用环境变量进行替换

```txt
Hello user $USER in $DESKTOP_SESSION. It's time to say $HELLO!
```

> export HELLO="good morning"
> envsubst < welcome.txt
> Hello user joe in Lubuntu. It's time to say good morning!



## go-zero

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





## gRPC

