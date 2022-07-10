[TOC]

# RPC

https://www.zhihu.com/question/25536695

https://www.zhihu.com/question/25536695/answer/417707733

RPC(Reomte Procedure Call，远程过程调用)——**分布式系统的通信技术**。即在 A 服务器上调用 B 服务器上的应用提供的函数（方法），应为不在同一个内存空间，所以不能直接调用，只能通过网络来表达调用的语义和传递调用的数据。

理解：其实和 HTTP 的方式一样，只是 HTTP 使用 HTTP 协议，在内部系统间调用时，依旧需要忍受 HTTP 协议的信息冗余。而 RPC 的方式使用自定义的传输协议。

> 使用RPC要解决的问题：
>
> - 通讯问题。通讯协议，连接
> - 寻址问题。找到目标服务器
> - 序列化和反序列化问题。多次序列化和反序列化操作。
>
> RPC协议：CORBA，Java RMI，Web Service的RPC风格，Hessian ，Thrift，甚至 REST API。
>
> 当两个物理分离的系统之间需要建立逻辑上的关联时，RPC 是常见的技术手段之一。除RPC外，常见的**多系统数据交互方案**还有分布式消息队列、HTTP请求调用、数据库和分布式缓存等。
>
> 其中 HTTP 和 RPC 调用是没有经过中间件的，是端到端的直接数据交互。HTTP 调用也可以看成是一种特殊的 RPC 调用，只不过传统意义上的 RPC 是长连接数据交互，而 HTTP 是短连接。
>
> RPC 在我们熟知的各种中间件中都有它的身影。 Nginx/Redis/MySQL/Dubbo/Hadoop/Spark/Tensorflow等都是在 RPC 的基础上构建的，是广义的 RPC ，指分布式系统的通信技术。

分布式意味着隔离，隔离意味着通信，通信意味着 RPC 的存在。RPC 涉及两个协议：对象序列化协议和调用控制协议。

## Dubbo

doc http://dubbo.apache.org/zh-cn/

[分布式 | Dubbo 架构设计详解](https://mp.weixin.qq.com/s/q8S3Ihas0KXVMfbdNjau0w)

[Dubbo入门---搭建一个最简单的Demo框架](https://blog.csdn.net/noaman_wgs/article/details/70214612)

dubbo 是一个 RPC 框架，提供了三大核心能力：面向接口的远程方法调用，智能容错和负载均衡，以及服务自动注册和发现。

![img](http://dubbo.apache.org/img/architecture.png)

| 节点        | 角色说明                               |
| ----------- | -------------------------------------- |
| `Provider`  | 暴露服务的服务提供方                   |
| `Consumer`  | 调用远程服务的服务消费方               |
| `Registry`  | 服务注册与发现的注册中心               |
| `Monitor`   | 统计服务的调用次数和调用时间的监控中心 |
| `Container` | 服务运行容器                           |

示例提供者安装：`git clone https://github.com/apache/incubator-dubbo.git`

安装：

```sh
git clone https://github.com/apache/incubator-dubbo.git
cd incubator-dubbo
运行 dubbo-demo-provider中的org.apache.dubbo.demo.provider.Provider
如果使用Intellij Idea 请加上-Djava.net.preferIPv4Stack=true
```

配置：

```sh
resource/META-INFO.spring/dubbo-demo-provider.xml
修改其中的dubbo:registry，替换成真实的注册中心地址，推荐使用zookeeper
```

start.dubbo.io

http://jm.taobao.org/2018/06/13/%E5%BA%94%E7%94%A8/