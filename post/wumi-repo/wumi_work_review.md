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

Wireguard VPN 工具

socat https://www.hi-linux.com/posts/61543.html   用于在中转机上连接 AB 两个网段，将中转机上某个端口的流量转发到 B 网段的某台机器的某个端口，这样 A 网段的机器就可以通过访问中转机上的端口访问到 B 网段的机器端口。



## 🟩项目列表

-   [ ] phone-business： 云手机业务
-   [ ] turbolive-server： 直播加速底层 HTTP 服务
-   [ ] resource：资源层 api & rpc 服务
-   [x] usercenter： 用户中心 RPC & HTTP 服务
-   [ ] billing-rpc： 账单 RPC 服务
-   [ ] pathlive-rpc：直播快相关 RPC 服务
-   [ ] console-resource: 控制台前端
-   [ ] operation-api: 运营平台 HTTP 服务
-   [ ] operation-console: 运营平台前端
-   [ ] order: 订单服务
-   [ ] phone-bill-api: 手机订单 HTTP 服务
-   [ ] luci-app-turbolive: 软路由 lua 脚本
-   [ ] pathlive-router-tools： 软路由工具包
-   [ ] pathlive-luci: 软路由系统定制页面，及 pathlive 线路集成
-   [ ] phone-backend-api: 云手机后端服务


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
     1.   bgp_ip(border gateway protocol): 下车点的外网 IP， 分配出来的 EIP，公网 IP，最终应用运行的那个 IP（如 TK 能看到登录的 IP）
     2.   vpc_ip: 下车点的UCloud 网络的内网 IP，更快速，不算下车点的带宽




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



#### 创建连接（使用线路）

>   连接过程：
>
>   从国内的上车点（源 IP）到转发机（一般是香港的机器），再到下车点（国外的机器）
>
>   其中，上车点的端口和转发机上的端口一一对应。转发机到下车点的端口临时分配



`turbolive-server/CreateUOLConnection`

1.   从请求参数中获取 line Token, 并从 pathlive.line_tokens 表获取该 token 的记录
2.   根据 token 中记录的 lineId 到 pathlive.lines 表中查询 line 信息
3.   创建链接前，如果该线路之前存在连接（连接未关闭），则先关闭连接（`pathlive.connections.active = false`）
4.   尝试选择一个可用的中转机以及未被使用的端口 `forward_server.OccupyOnePort`
     1.   创建 WireGuard（一种 VPN） 配置。设置目标 DNS (DestDNS)(ini 文件配置好的)
     2.   循环 10 次尝试获取中转机、连接
          1.   **寻找中转机**：从 `pathlive.forward_servers` 表中选取可用的中转机。找不到则表示线路全忙，则中止寻找
               1.   起始区域和指定的起始区域一致的
               2.   当前中转机上的连接不超过 200 （run-mode.max_load 设定）的
          2.   **寻找可用端口**：从找到的转发机中找一个空闲的端口 `FindForwardServerPort`
               1.   获取配置 `run-mode.fs-start-port` 表示**起始端口**
               2.   找到该转发集群（forward_server）上最后一次使用的连接信息（ `pathlive.connections` 表），找到最后一个端口
               3.   找到该转发机上所有可用的连接，如果没有在用连接，则**起始端口**就是可用的，返回即可
               4.   按顺序，在 maxLoad 个端口中找一个没有在用的端口（在用的端口在查出来的所有 connecting 中有记录）
          3.    **wireguard 配置**：（转发机的上车点 IP, 找到的可用端口，作为 wireguard 的入口设置）
          4.   **生成 connection 对象**：并将该记录插入 `pathlive.connections` 表中（尝试插入 20 次，成功则停止）。connection 表会记录连接使用的**上车点信息，转发机信息，wireguard 配置**
          5.   检测刚创建的连接是否可用。（**此时转发机及其端口还没正式使用，只是在 DB 中记录将被使用**）
               1.   查询该转发机该端口上的启用连接是否只有一个（刚创建的那一个），若有多个，则说明该端口改转发机被占用，进入下一次尝试
               2.   查询该线路上的可用连接是否只有一个（刚创建的那个），若有多个，则说明该线路有多个连接在使用，则进入下一次尝试（**一个线路只支持一个连接的时候，需要检测线路的占用情况。但是当前采用的是后一个连接会挤掉前一个的策略，所以此处检测线路占用情况的功能可以去除**）
5.   `DeployForwardServer` 对找到的转发机进行 ssh 配置下发部署（异步执行）
     1.   找到该连接对应的线路 & 转发机（relay,forward_server）
     2.   部署：
          1.   **在下车点部署 wireguard server**，用于 VPN（client 是 app 在直播时创建，连接这个 server ）
               1.   生成 wireGuard 白名单。根据线路的 serviceType 选择白名单（TK-1, TK-2）
                    1.   如果该账户有配置白名单覆盖的，则以这个配置的为准
                    2.   获取白名单详细列表 `GetWhiteListConf(ServiceType)`  （**在白名单中的地址是可以通过的，其他的都会被阻拦**）
                         1.   支持多个 serviceType, 用逗号分开
                         2.   到 `pathlive.white_lists` 表中查询指定 serviceType 的白名单，并拼成特定格式的字符

                    3.   将白名单生成写入 wireGuard 的白名单配置文件中的命令 `echo <whitelist> > /path/to/wireguard/whitelit.conf`

               2.   在下车点部署 wireGuard server, 尝试 10 次，任何一次部署成功就停止
                    1.   默认走香港代理，并且在多次尝试中，代理和直连的方式交替进行，可以在部署失败下次尝试时多一种选择
                    2.   ssh 连接到下车点的机器。（ip: line.bgpIp, port: 22）。
                         1.   如果使用代理。则使用内建的代理服务器进行连接（ini 中配置 xx-proxy）,通过代理连接下车点公网 IP（ip: line.bgpIp, port: 22）
                         2.   否则通过下车点的公网 IP 直连（ip: line.bgpIp, port: 22）

                    3.   运行命令，部署 wireGuard server (**本次连接会挤掉上一次的连接**)
                         1.   停止已有的 wireGuard server
                         2.   白名单配置写入配置文件
                         3.   重启 wireGuard server

          2.   **在中转机部署 socat,** 用于流量转发。尝试 10 次，任何一次部署成功就停止
               1.   选择中转机的目标 IP `GetDestRelayIp`， 最终中转到的 IP（下车点）
               2.   默认走代理， ssh 连接中转机
               3.   执行命令，部署 socat
                    1.   停止存在的 socat
                    2.   将中转机的目标 IP 写入 socat 的配置（最终流量转到的 IP） `GenerateRelaySystemd`
                    3.   重启 socat ，使配置生效

     3.   如果 wireguard server 和 socat 都部署好了。将此条连接的信息写入  `pathlive.connections` 表


#### 中断连接

`tubrolive-server/TerminateUOLConn`

将连接的 active 置为 false 即可。机器上安装的 wireguard & socat 下次使用的时候会覆盖





## operation-api

运营后端 api



### 线路抓包功能

最终实现：`pathlive-rpc/LineTcpDump` 

实现： 顺着中间节点一个个连过去

原因：当然可以直接通过公网连接下车点机器，执行 tcpdump 命令进行抓包，但是抓包的结果回传，如果直接在下车点上传，则会使用下车点机器的带宽，可能会影响用户的正常使用。所以抓包时，会顺着代理的一个个节点构建一条通道，最终抓包的结果顺着通道返回（通道一般会走 Ucloud 专线，也叫 UDPN,UCloud Dedicated Private Network）



#### 获取抓的包

`operation-api/tcpdump （get）` 获取生成的包。抓好的包会存放在 ofileos 上。`pathlive.tcpdump` 表存放文件名

## UserCenter

用户中心

用户类型(`usercenter/userinfologin.go`)：

```text
// 账号类型分这些
// - 直销客户
// - 渠道1：UCloud SML、UCLOUD EIU、UCLOUD XX、王磊、以及离职销售组成的兼职小团队
//                  这类客户会给我们带来终端客户，也会带来代理商。客户和收入都走我们这里
//                  主账号和子账号分别用于看整合和名下销售的客户情况，通过邀请码关联客户
//                  我们一般给这类渠道固定折扣，折扣之上属于渠道的利润。暂定长期返毛利
// - 渠道2：UCLOUD离职销售或朋友
//                  和渠道1很类似，带来的客户有可能是终端客户，也有可能是代理商。客户和收入都走我们这里
//                  区别在于返佣方式不一样，渠道1侧重于长期合作，暂定长期返；渠道2则是短期行为，比如返三个月内每个月月销30%
// - 代理：在我们这里体现为单个客户，他的客户我们看不见
//                  直播快有一些代理，比如正心，我们看不到正心的终端客户情况
//                  代理自行维护客户关系，并解决客户售后问题
```

### RPC

包含功能： 

1.   用户
2.   群组：group 表
3.   鉴权：使用 Redis 做 token 过期控制（一个月过期， 2592000 秒）
4.   实名认证： user 表 & certification 表
5.   资源过期通知：资源快到期或者已经到期，通过短信通知用户及时充值。此处不管是否自动续费的逻辑，如果设置的自动续费，但是余额不足，也需要进行通知。构造短信模板
     1.   到期前 3 天进行通知，到期后 4 天进行资源删除
6.   渠道用户管理（代理商）（user 表 channel 相关字段）
7.   系统通知：企业微信群机器人通知

### API

-   [ ] ⚠️需要了解渠道 & 代理相关的业务逻辑

包含功能：

1.   微信、QQ第三方登录，获取 UnionId 存到 user 表里
2.   发票服务：发票信息管理 & 卡票可开管理，发送通知，提醒相关人员进行开票操作
     1.   invoice 表：发票及金额表
     2.   invoice_info 表： 发票本身的信息，一般是已经开具的发票记录
3.   🟪渠道（代理商）渠道 A: Ucloud 销售，直接返现金；渠道 B: 
     1.   渠道 A：
          1.   获取渠道信息`channel.GetChannelDataByUserId`（指定用户的这个渠道）：所有下级用户，折扣，返利，所有可用资源（线路 & 云手机）
          2.   获取渠道消费信息`channel.GetChannelConsumptionByInviter`: **指定时间段内**的所有下级用户的线路 & 手机资源及其时间段内的消费信息。消费信息包括：
               1.   订单金额，续费金额……计算方式：`GetConsumptionAndCalcTime`
          3.   获取渠道用户返利 `GetChannelBonusByInviter`: 指定时间段内的所有下级用户的返利总和。计算方式：使用用户购买 & 充值的总金额进行计算
          4.   获取渠道用户的支付信息 `GetChannelPaymentByInviter`
     2.   渠道 B: 分为直销和代理。逻辑类似，只是返利不一样
4.   子账号：
     1.   创建`createsublogic.go`：账号名：用户名 + @主账号
     2.   删除 `deletesublogic.go`：同时回收其名下资源
     3.   登录`subloginlogic.go`: 密码做特殊处理，记录登录
5.   老推新活动：用户表里，channel_inviter 为空，并且有具体的邀请人 inviter 为邀请人的 userId
6.   用户充值：
7.   登录注册`registerlogic.go`： 登录和注册活动都需要记录到表里。支持邀请码，渠道分销及其渠道定级（channel_invite_level）
8.   验证码：验证码存在 Redis 中，以手机号为 key,  3 分钟过期
9.   用户余额





## Order & phone-bill-api

用于处理各种订单 & 账单的 api



### Order

**现在该项目中的接口基本没怎么使用，主要用来提供订单相关的实体**

#### 退款

refund 表

重要字段说明：

| 字段名 |    含义    |                备注                 |
| :----: | :--------: | :---------------------------------: |
| total  | 总退款金额 | 单位：分；提供给其他服务时叫 refund |
| count  |  退款数量  |                                     |
|  all   | 订单总金额 | 单位：分；提供给其他服务时叫 total  |

退款时，计算好应退金额，在 refund 表中记录一笔，发送异步通知：

1.   通知 Key: `asynq:order:refund`
2.   队列： `payment` 
3.   payload: refundId
4.   处理逻辑: `phone-bill-api/asyncserver.HandleOrderRefundTask`  (**该退款逻辑已启用。用于原先创建资源失败时自动退款，现在没有自动退款，只有运营平台可以操作退款**)
     1.   根据 refundId 查询 refund 详情（创建退款单时写入的记录）
     2.   根据订单信息找到这个订单的支付信息
          1.   确认是否已经支付，以及支付金额
          2.   确认支付方式，原路退回



### phone-bill-api

主要有订单处理、支付处理、续费、充值、计价等功能



#### 续费——线路

`phone-bill-api/orderrenewlogic.OrderRenew`  支持批量续费、支持手机 & 线路续费。此处只关注线路续费逻辑。列举需要关注的几个特殊点。

1.   创建 **parentId**。因为支持批量续费，但是每个资源又都属于一个 order，故使用一个 parent order 将本次续费的所有资源关联起来，parent 也存储续费总额。
2.   req.CreateOrder。 前端控制本次是否真正创建 order 记录。若指定不创建，一般是用于计算本次续费的总价。
3.   查询线路价格时，需要 productId & ipRegion 两个数据，这两个值在 line 中本身是有记录的。但是不直接使用，而是再查询数据库的原因是，产品和地区都可能下架，如果续费时，产品或区域已经下架，则不能续费。
4.   和新建订单一样。计算好价格，订单表新增记录后，等待用户支付，设置一个超时取消订单的任务



## pathlive-ops-api

### 换 IP 

代码： `pathlive-ops-api/change-eip.go`

线路换 ip 。大致流程为：

1.   解绑原先的 EIP 和 uhost 的绑定
2.   申请新 EIP
     1.   查询当前可用的 IP 段（eip_segments 表，查询同区域，同线路类型的可用段，默认与原 ip 不在同一个网段，随机选择一个）
     2.   调用 ucloud api 申请 EIP
3.   绑定新 EIP 和原先的 uhost



### 客户端实时 ping 结果

代码：`pathlive-ops-api/ping-client.go`

原理：连接到下车点，执行 ping 命令，得到返回结果

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

常用参数：

1.   

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



## 计算机网络



### 子网掩码

参考： https://www.bookstack.cn/read/network-basic/7.md

子网掩码的表示： 154.71.150.42/22 表示 154.71.150.42 这个 IP 的子网掩码是 255.255.252.0。计算方法如下：

22 表示 32 位的子网掩码中，前 22 位都是 1, 后 10 位是 0， 即 `11111111 11111111 11111100 00000000`，这样转成十进制的就是 `255.255.252.0`

前 22 位都是 1 ，表示这些位被**掩盖**了，不能用于表示该子网下的主机，即剩下能表示主机的位只剩 10 位，就是说，这个子网的这个 ID 下能表示的主机数是 $2^{10} = 1024$ , 是这个子网段下主机数最多的子网。

可以看出，这是一个 B 类网络，前 16 位表示网络号，22 - 16 = 6 位表示子网，也就是这个 B 类网络下，能有的子网段数量是 $2^6 = 64$ 个，总共能容纳的主机数： $\Sigma^{i}_{1 \le i \le 6}{2^i * 2^{16-i}}$

`154.71.150.42` 这个 IP 对应的二进制表示为： `10011010 1000111 100101/10 101010`， 因为其子网掩码为 22, 可以看出这是一个 B 类网络，则其前 16 （$\lfloor22 / 8\rfloor * 8 = 16$）位是不动的，并且其子网段总共有 6 ($22 \% 8 = 6$) 位



## Kubernetes



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



## Docker cheatsheet

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
```

#### container

```sh
docker container -h

docker container attach <id>

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



https://yeasy.gitbook.io/docker_practice/

## SSH Config

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



## 缓存 & Db

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



## 🔲Frp

https://sspai.com/post/52523



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

## Golang 异步确认 & 超时控制

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



## TikTok 简易爬虫实现

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

