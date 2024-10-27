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

#### 线路流量是怎么走的

在连接创建好之后，有几个重要的点

1.   客户端 wireguard 配置

```txt
[Interface]
Address = 10.13.13.2   (标识当前短点 VPN IP)
PrivateKey = cMBcS9bDJ1u4aLp1VCf+we6t4fXfi5s7v93poGxv/XU=
DNS = 8.8.8.8
MTU = 1392

[Peer]
PublicKey = YnaDr0SXsp3OKD2i0XP3z7dUUHVGDlAs7+xQ3ZOHu2E=
Endpoint = <上车点 IP>:30258   (上车点 IP)
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 30
```

 此处配置，将用户客户端和上车点联通。（上车点作为 wireguard 的中继服务器）

2.   上车点预先运行好 socat ，将流量转发到中转机

使用脚本预先生成，具体 systemd 配置，在上车点机器根目录下 a.sh 文件内，配置大致如下：

```txt
[Unit]
Description=Socat Wireguard 30258

[Service]
Type=simple
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=socat-wireguard-30258

ExecStart=/usr/bin/socat -T 600 UDP4-LISTEN:30258,reuseaddr,fork UDP4:<中转机 IP>:30258
Restart=always

[Install]
WantedBy=multi-user.target
```

3.   中转机上的 socat 再将流量转发到下车点（**创建连接时部署**）



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
