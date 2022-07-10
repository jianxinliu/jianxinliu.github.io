# Eureka

https://cloud.spring.io/spring-cloud-netflix/multi/multi__service_discovery_eureka_clients.html

# Spring Boot



# Spring Cloud

 

## Spring Cloud Config



## Spring Cloud Netflix



# 断路器 Hystrix 客户端

see ./Hystrix.md

# 客户端负载平衡器 Ribbon



# 声明式REST客户端 Feign	

使用Feign可以很方便的在已有的REST服务上添加自己的代码，使用FeignClient访问已有服务或者远端服务器，自己的代码只需要和FeignClient交互即可。FeignClient配置有Hystrix 熔断器，当远端返回失败时调用熔断器返回结果，从而保护远端服务。

Feign 在前端发来的请求和后端API之间加入一层间接引用，以便实现分离，分离之后便可以对请求和响应做很作事情：负载均衡，熔断，压缩，日志。

Feign使用Jersey和CXF等工具为REST或SOAP服务编写java客户机。此外，Feign允许您在Apache HC等http库之上编写自己的代码。Feign通过可定制的解码器和错误处理将您的代码与http API连接起来，而这些解码器和错误处理可以被编写到任何基于文本的http API上。 

使用 Feign 时，Spring Cloud 集成了	Ribbon 和 Eureka来提供负载均衡的HTTP客户端。

```java
@FeignClent("stores")
public interface StoreClient{
    @RequestMapping("/stores")
    List<Store> getStores();
}
```

`@FeignClient`注解的参数是一个任意定的名字（上面的是 stores），用于创建 Ribbon 负载均衡器。Ribbon 客户端将会去寻找一个叫 stores 的物理服务，如果应用了 Eureka 客户端，那 Ribbon 客户端将会在 Eureka 注册中心去解析（寻找）指定服务。



Spring Cloud 的 Feign 支持的核心概念就是命名的客户端。

# 路由和过滤器 Zuul



# Jenkins





# Docker

see ./docker.md

# Kubernetes



# angularjs

see ./angularjs/angularjsLearning.md



----

----

----

----



# velocity

​	模板引擎，like jsp,freemaker

# quartz

​	作业调度框架

# cxf

​	Web services

# jibx

​	绑定 XML 数据到 Java 对象的框架

# ibatis

​	半自动化的ORM

# appfuse

​	入门级 J2EE框架

# axis2

​	SOAP引擎，提供创建服务器端、客户端和网关SOAP操作的基本框架。

# WSDL

​	Web Services Description Language

​	

# svn:

|svn|is like|git|note|
|:---|:---:|:---|:---:|
|svn checkout|   is like  | git clone|check out the repository to local|
|svn commit|     is like |  git commit|option `-m` is also available to add comments|
|svn diff |            is like|   git diff|show the modified lines|
|svn update|      is like |  git pull & git merge|default to update to the lastest version|
|svn status   |     is like |  git status||
|svn add   |         is like  |git add||
|svn revert|is like|git reset HEAD|undo to the preview status,use the option `-R`to recursive revert the folder|
|svn cat,svn list,svn log||||
|svn copy|is like|git branch|copy a new branch,`cd` to the working directory|
|svn merge|is like|git merge|a hiddern parameter is pwd|





# Angular
