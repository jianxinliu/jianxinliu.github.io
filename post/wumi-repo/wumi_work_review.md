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

asynQ https://pkg.go.dev/github.com/hibiken/asynq#section-readme


## phone-business(比较通用的构建流程)

### build

使用 docker 进行构建

1. 读取 /deploy/templates 下的配置文件。分为 preview & production
2. 将配置文件中的配置信息设置为环境变量
3. 使用 `envsubst` 命令将配置文件中的参数使用环境变量替换到 `config.tmpl` 中（`config.tmpl` 是配置的模板，只需要往里面填充对应的参数即可）
4. 将替换结果写入到 `phone-api.yaml` 下。（最终生效的配置文件是这个）

#### make 本地构建

使用 make 进行不同阶段的构建，有点类似前端项目中 package.json 中的 scripts 声明，写好对应的命令后直接点击执行即可。

项目中的 Makefile 支持的操作：

1.   goct 代码生成， model, api, rpc 等代码生成。有改动后，直接运行相关的命令即可生成代码
2.   项目构建
3.   docker builder & push

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



### Xphone 库表结构

1. Model 表： 支持的类型（配置表）：
     1. 金牌线路（golden）,金牌高清线路（goldenPlus）,银牌线路（Silvery），珀金线路（Platinum）
     2. UPhone Mini, Phone X, Phone Live
2. Product 表：产品表。记录所有的产品类型。属于配置表。通用版（universal）, 矩阵版（matrix），无网版（phone-noip）……



### pathlive 库表结构

1.   lines: 直播线路表，记录线路的起始、终止区域（上车点: 从国内节点上车，走专有通道到下车点、下车点：最终请求发出地是该 ip），ucloud 提供的 ip (**bgp_ip**, **vpc_ip**)



### 从创建一条线路开始



#### 创建云手机



1.   前端发送请求 `/order/create`, 接口实现在 `phone-bill-api` 项目 `ordercreatelogic.go`
     1.   查询当前用户的信息：`userCenterRpc/UserInfo`
     2.   子账号不支持创建订单
     3.   账号未实名制不支持创建订单
     4.   如果是矩阵版云手机，则限制 3 台起售



#### 创建线路

1.   **订单创建**：前端发送请求 `/order/create_path`, 接口实现在 `phone-bill-api` 项目 `createpathorderlogic.go`
     1.   查询当前用户的信息：`userCenterRpc/UserInfo`
     2.   子账号不支持创建订单
     3.   账号未实名制不支持创建订单
     4.   如果购买类型是“到月末”或者是“到下月末”，判断是否有资格
          1.   根据该用户是否在白名单中
          2.   如果是购买到月末，则根据当前距离月末是否大于 7 天
     5.   根据订单的产品 id 获取产品信息（调用 resourceRpc.GetProductInfo，实际上是去 product 表根据 product_id 查询记录）
     6.   根据产品 id 和区域 id 获取价格ID 。（调用 resourceRpc.GetResourcePrice，实际上是去 resource_price 表查询）（price_id 举例: //resource/products/path-golden/regions/uae-dubai）
     7.   根据 priceId 获取实际的价格（调用 billingRpc.GetPurchasePriceV2）(计价规则单独列举)
     8.   根据订单的 regionId 获取 region 信息（调用 resourceRpc.GetIpRegionInfo，实际上是去 ip_region 表根据 regionId 查询一条记录）
     9.   获取订单指定组的组信息。（调用 userCenterRpc.GroupInfo, 实际上是去 group 表查询）。如果未指定，则默认采用“默认分组”
     10.   order 表中插入一条记录
     11.   将当前订单信息通过异步队列发送（实际上使用的是 Redis）
           1.   事件名：asynq:order:timeout。任务 payload 携带信息： `<Order type>|<resource type path>|<order id>`
           2.   事件 Handler： `order.asnycqServer.HandleOrderPayTimeoutTask`
           3.   实际上是创建了一个超时未支付取消订单的任务（状态设置为 timeout）(1 分钟超时)
2.   **支付**：代码路径（`phone-bill-api/orderpathlogic.go`）
     1.   获取订单信息
     2.   根据支付类型选择不同的处理方式
          1.   支付宝（app/web）
          2.   微信（app/web）
          3.   余额
               1.   `billingRpc/UserWalletDeduct` 更新用户钱包余额
               2.   更新支付信息(payment 表)
3.   **订单处理**：代码路径（`phone-bill-api/common.go -> OrderPaid`）
     1.   区分是手机订单还是线路订单，续费订单还是充值订单
     2.   线路订单：将任务 `asynq:order:process:v2` 加入 `resource` 队列，payload 携带 orderId。resource 队列 handler: `resource/asyncqserver.go.HandleOrderProcessTask`
          1.   根据订单号查找订单。
          2.   根据付费类型，计算此订单的有效期（开始、结束）。（`billingRpc/CalcExpireDate`）
          3.   创建 resource。`createresourcelogic.CreateResource`
               1.    **创建线路**：`create.go.CreateResource` 。调用 `turbolive-server/CreateUOLLine`，得到资源 id
                    1.   请求参数：ChargeType, Quantity, SourceRegion, DestRegion, LineId, LineType
                    2.   校验 DestRegion 是否是当前所支持的目的区域（ini 配置文件中的 dest-region section 配置）
                    3.   根据目的区域决定上车点（sourceRegion） `SourceRegionDecision` 默认从广州上车
                    4.   在 pathlive.lines 表创建 line 记录（循环创建 3 次，成功则停止，**为啥？**）
                    5.   如果创建成功，将 line 的相关信息更新到该条记录上。（循环更新 3 次，成功则停止，**为啥？**）
                         1.   如果更新成功，从 ini 配置文件中 `run-mode` section 读取配置是否需要创建 UHost
                              1.   创建 UHost & EIP，创建成功则保存 UHost 相关信息到 line 记录上 （循环创建 10 次，成功则停止，**为啥？**）。并且根据 uhostId 异步获取 IP
                                   1.   异步创建 UHost，尝试 3 次
                                   2.   异步创建 EIP
                                        1.   如果是铂金线路，尝试 3 次。使用 EIP 管理策略选出一个 EIP， 逻辑参考：**EIP 管理** 章节
                                        2.   其他线路，则分配 EIP。重试 5 次进行分配 IP `allocateEip`。（原因是：<u>因指定了多个可申请的IP段，以防止某个段没IP了，或者UCloud后台把某个段锁定等情况，导致 allocateEip 失败，多重试几次</u>）
                                        3.   判断 UHostId & EIPId 都创建成功。有则将 EIP 绑定到 UHost 上。
                                             1.   如果是铂金路线，需要将 EIP 重置为使用中的状态（`eip.UsedTime = now`），并且将默认带宽增加到 10M
                                        4.   否则，将创建好的 UHost 或者 EIP 对应删除掉（非铂金线路的 EIP 需要释放）
                              2.   否则需要回滚
                         2.   否则需要回滚
                    6.   创建不成功需要**回滚** `rollback`。删除 line 记录。**计费、资源等回滚操作由网关层公共服务统一实现**
               2.   所有资源 id 写入 resouce 表
               3.   每个 id 都作为一个任务，写入 resouce 队列，事件名：`asynq:resource:create_check`。事件handler: `resouce/asynqserver.HandleCreateCheckTask`
                    1.   根据 resourceId 找到 resource 表中的 resource 信息
                    2.   根据 resourceId 请求 `turbolive-server/DescribeUOLLine` 找到 resource 对应的线路信息（ip、区域）
                    3.   如果该线路可用（line.Enable）。在 resouce 表中标记该资源的状态为运行中，并设置该资源的过期检查任务
                         1.   `utils.SetResourceExpire`。异步进行。向 resource 队列注册一个资源过期的事件 `asynq:resource:expired`， handler 为 `resource/asynqserver.HandleExpiredTask`，执行时间为过期时间，即在过期时间到的时候执行资源清理工作
                              1.   再次检查过期时间是否在当前时间之前，否则再次设置过期检查事件 `utils.SetResourceExpire`（**过期时间可能被续费等操作更新**）
                              2.   检查当前资源是否有设置自动续费，有则自动续费 `resource/asynqserver.autoRenew`
                              3.   将该资源做关机操作 `PowerOffResource`，实际执行的是 `turbolive-server/SetUOLLineStatus`
                              4.   更新 resource 表，标识该资源已到期
                              5.   注册 7 天后删除该资源的事件 `utils.SetResourceExpireDelete`（如果 7 天内重新续费，则该事件被删除）
                                   1.   事件： `asynq:resource:expired_delete` handler：`resource/asynqserver.HandleExpiredDeleteTask` ，七天后执行
                                   2.   再次检查过期时间 + 7 天是否在当前时间之前，否则再次设置过期删除事件 `utils.SetResourceExpireDelete` （**理由同上**）
                                   3.   此时不管设定的过期时间是否真的过期，过期则设置为"**已删除**"状态，否则设置为"**提前删除**"状态。
                                        1.   调用 `turbolive-server/DeleteUOLLine` 执行实际删除操作
                                             1.   从 pathlive.lines 表中获取线路信息
                                             2.   更新其状态为不可用，vpcIp 为"删除中"
                                             3.   异步执行删除。
                                                  1.   如果是铂金线路，则不删除 EIP, 只是将线路的带宽变为 1M(**连接 UCloud 操作**)， 并标记 lastReleaseTime 为当前时间
                                                  2.   否则删除云主机（**UHost**）
                                                  3.   逻辑删除 lines 表记录（删除十次，一次成功则停止，**为啥？**）
                                        2.   resource 表做逻辑删除（`resouce.Deleted = 1`）
                    4.   如果不可用。查看该任务的重试次数，如果超过 90 次还未成功，则标记当前资源创建状态为失败，更新到 resource 表中

##### UHost  & EIP 

云主机用来走流量，EIP 作为该流量的起始地址



##### EIP 管理

`eip_manage.FindOneUsableEipWithRegion`

从 `pathlive.eips` 表中查找一个指定区域的 EIP， 并且满足条件： `usable = true and last_release_time < now - eip_lock_days ` ，并从中随机选择一个

其中：usable 是非时间条件管控，可能是出于其他原因需要禁用该 EIP, last_release_time 是时间条件的管控 



##### 计价规则

`billingRpc.GetPurchasePriceV2` 

1.   参数：UserId, PricingId, ChargeType, Count(购买数量)



#### 创建链接（使用线路）

>   连接过程：
>
>   从国内的上车点（源 IP）到转发机（一般是香港的机器），再到下车点（国外的机器）
>
>   其中，上车点的端口和转发机上的端口一一对应。转发机到下车点的端口临时分配



`turbolive-server/CreateUOLConnection`

1.   从请求参数中获取 line Token, 并从 pathlive.line_tokens 表获取该 token 的记录
2.   根据 token 中记录的 lineId 到 pathlive.lines 表中查询 line 信息
3.   创建链接前，如果该线路之前存在连接（连接未关闭），则先关闭连接（`pathlive.connections.active = false`）
4.   尝试选择一个未被使用的端口 `forward_server.OccupyOnePort`
     1.   创建 WireGuard（一种 VPN） 配置
     2.   循环 10 次尝试获取
          1.   从 `pathlive.forward_servers` 表中选取可用的中转机。找不到则表示线路全忙
               1.   起始区域和指定的起始区域一致的
               2.   当前中转机上的连接不超过 200 （run-mode.max_load 设定）的
          2.   从找到的转发机中找一个空闲的端口 `FindForwardServerPort`
               1.   获取配置 `run-mode.fs-start-port` 表示起始端口
               2.   找到该转发集群（forward_server）上最后一次使用的连接信息 `pathlive.connections` 表
               3.    
          3.   `DeployForwardServer` 对找到的转发机进行 ssh 配置下发部署（异步执行）

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

