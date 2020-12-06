[TOC]

# Spring Cloud (一) 服务的注册与发现（Eureka）

https://www.w3cschool.cn/spring_cloud/spring_cloud-2hgl2ixf.html

Spring Cloud是一个基于Spring Boot实现的云应用开发工具，它为**基于JVM的云应用开发中涉及的配置管理、服务发现、断路器、智能路由、微代理、控制总线、全局锁、决策竞选、分布式会话和集群状态管理等**操作提供了一种简单的开发方式 .

## 微服务

[微服务架构-博客园](http://www.cnblogs.com/imyalost/p/6792724.html)

## 服务治理

由于**Spring Cloud为服务治理做了一层抽象接口**，所以在Spring Cloud应用中可以支持多种不同的服务治理框架，比如：Netflix Eureka、Consul、Zookeeper。在Spring Cloud服务治理抽象层的作用下，我们可以无缝地切换服务治理实现，并且不影响任何其他的服务注册、服务发现、服务调用等逻辑。 

## Spring Cloud Eureka

实现服务治理

### Eureka Server

提供服务注册和发现。

分两个角色：Eureka Server 和 Eureka Client.

- Eureka Server 用来提供服务注册和发现
- Eureka Client 用来给一个服务提供把自己注册到 Eureka Server 的能力。

#### 添加依赖

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-eureka-server</artifactId>
</dependency>
```

#### 开启服务注册

通过 `@EnableEurekaServer`注解启动一个服务注册中心提供给其他应用进行对话。

```java
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.server.EnableEurekaServer;
@SpringBootApplication
@EnableEurekaServer
public class EurekaServerApplication {
    public static void main(String[] args) {
        SpringApplication.run(EurekaServerApplication.class, args);
    }
}
```

#### 添加配置

在默认配置中，服务注册中心也会将自己作为客户端来尝试注册自己，所以需要禁用注册中心的客户端注册行为，在 `application.yml` 中添加：

```yml
registerWithEureka: false
fetchRegistry: false
```

完整配置：

```yml
server:
  port: 8761
eureka:
  instance:
    hostname: localhost
  client:
    registerWithEureka: false
    fetchRegistry: false
    serviceUrl:
      defaultZone: http://${eureka.instance.hostname}:${server.port}/eureka/
```



## Service Provider

- 服务提供方
- 将自身服务注册到Eureka 注册中心，从而使服务消费者可以找到

#### 添加依赖

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-eureka-server</artifactId>
</dependency>
```

#### 开启服务注册

在应用主类上添加 `@EnableEurekaClient` 注解，这个注解只有Eureka 可用，也可以用 `@EnableDiscoveryClient`。

```java
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
@SpringBootApplication
@EnableEurekaClient
@RestController
public class EurekaProviderApplication {
    @RequestMapping("/")
    public String home() {
        return "Hello world";
    }
    public static void main(String[] args) {
        SpringApplication.run(EurekaProviderApplication.class, args);
    }
}
```

#### 添加配置

需要配置才能找到Eureka 服务器，完整配置如下：

```yml
eureka:
	clent:
		serviceUrl:
			defaultZone:http://localhost:9001/eureka/
			
spring:	
	application:
		name:eureka-provider
		
server:
	port:8089
```

其中`defaultZone`是一个魔术字符串后备值，为任何不表示首选项的客户端提供服务URL（即它是有用的默认值）。 通过`spring.application.name`属性，我们可以指定微服务的名称后续在调用的时候只需要使用该名称就可以进行服务的访问 

# Eureka 自我保护模式

自我保护模式是Eureka为了提高分区容错性而采用的一种手段，当网络情况不佳服务于注册中心间的心跳不佳，有些**注册中心会认为这些服务已经不可用了，便会直接将其剔除（deregistry）**。在有watch机制支持的注册中心里这样问题不大, 当网络恢复正常, watcher可以重新把服务注册回来. 但Eureka Server并不支持watch机制, 它会认为这是网络原因造成的服务暂时表现的不健康, 但是Consumer可能是可以正常消费的. 所以不会把这些服务剔除, 这就是自我保护模式。

进入保护模式会干什么：

1. Eureka Server不再从注册列表中移除因为长时间没收到心跳而应该过期的服务。
2. Eureka Server仍然能够接受新服务的注册和查询请求，但是不会被同步到其它节点上，保证当前节点依然可用。
3. 当网络稳定时，当前Eureka Server新的注册信息会被同步到其它节点中。

  保护模式会给我们带来一些干扰, 比如Windows机器上启动了一个服务, 注册到Eureka, 当停止的时候Windows是直接杀进程的, 所以不会发生Deregistry, 这种残留的服务就会导致Eureka进入自我保护模式, 导致服务一会通一会不通. 我们可以通过配置

```
eureka.server.enable-self-preservation = false
```

来禁用自我保护模式, 不过官方不建议这么做.

 

 ## `@SpringBootApplication` 与 `@SpringCloudApplication`

`@SpringCloudApplication` 注解在 `@SpringBootApplication`的基础上，增加了服务发现`@EnableDiscoveryClient`和熔断器
`@EnableCircuitBreaker`的支持。



`@SpringCloudApplication`注解的实现

```java
/**
 * @author Spencer Gibb
 */
@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Inherited
@SpringBootApplication
@EnableDiscoveryClient
@EnableCircuitBreaker
public @interface SpringCloudApplication {
}
// SpringCloudApplication 注解是一个空实现，只是增加了注解的支持，省去了写的麻烦
```

`@SpringBootApplication` 注解的实现

```java
@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Inherited
@SpringBootConfiguration
@EnableAutoConfiguration
@ComponentScan(excludeFilters = {
		@Filter(type = FilterType.CUSTOM, classes = TypeExcludeFilter.class),
		@Filter(type = FilterType.CUSTOM, classes = AutoConfigurationExcludeFilter.class) })
public @interface SpringBootApplication {
    ...
}
```

