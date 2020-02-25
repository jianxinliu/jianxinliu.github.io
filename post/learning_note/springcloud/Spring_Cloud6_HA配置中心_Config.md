[TOC]

# Spring Cloud Config 高可用的分布式配置中心

分布式配置中心的作用：方便管理众多服务的配置文件，及时更新。

和 Eureka 一样分为两个角色：

- Config Server
- Config Client

## 简介

> `SpringCloudConfig`就是我们通常意义上的配置中心，把应用原本放在本地文件的配置抽取出来放在中心服务器，从而能够提供更好的管理、发布能力。`SpringCloudConfig`分服务端和客户端，服务端负责将`git svn`中存储的配置文件发布成`REST`接口，客户端可以从服务端REST接口获取配置。但客户端并不能主动感知到配置的变化，从而主动去获取新的配置，这需要每个客户端通过`POST`方法触发各自的`/refresh`。
>
> `SpringCloudBus`通过一个轻量级消息代理连接分布式系统的节点。这可以用于广播状态更改（如配置更改）或其他管理指令。`SpringCloudBus`提供了通过`POST`方法访问的`endpoint/bus/refresh`，这个接口通常由`git`的钩子功能调用，用以通知各个`SpringCloudConfig`的客户端去服务端更新配置。
>
> 注意：这是工作的流程图，实际的部署中`SpringCloudBus`并不是一个独立存在的服务，这里单列出来是为了能清晰的显示出工作流程。
>
> 下图是`SpringCloudConfig`结合`SpringCloudBus`实现分布式配置的工作流

![SpringCloudConfig](http://www.ymq.io/images/2017/SpringCloud/config/1.png) 

工作流程：(git backend profile)

1. 配置变化 -> git Web Hook通知 Config Server
2. 通过 Spring Cloud Bus 给 Config Client 发消息
3. Config Client 主动发起请求，reload config
4. Config Server pull 变化的配置（按需更新配置？）

或者：

指定配置的位置，通过重新构建 config Server 工程去重新发现新配置。（File System backend profile）

## 资源发布

Spring Cloud Config 将配置资源发布成一个小型的 SpringBoot Application，将配置暴露成 RESTful API，这样便可以通过 URL 去访问所需要的配置。这个小型的应用提供如下形式的资源：

```
/{application}/{profile}[/{label}]
/{application}-{profile}.yml
/{label}/{application}-{profile}.yml
/{application}-{profile}.properties
/{label}/{application}-{profile}.properties
```

例如：有一个服务，application.name = config-server,其特定位置有`application-dev.yml`和 `application-test.yml`两个配置文件，则启动`config-server` 服务之后，可以通过 `/config-server/test`去访问`application-test.yml` 的配置，返回如下：

```json
{
    "name":"config-server",
    "profiles":[
        "test"
    ],
    "label":null,
    "version":null,
    "state":null,
    "propertySources":[
        {
            "name":"classpath:/config/application-test.yml",
            "source":{
                "server.port":8094,
                "eureka.client.serviceUrl.defaultZone":"http://localhost:8659/eureka/",
                "spring.application.name":"eureka-provider",
                "app.name":"shu ting a",
                "app.word":"Hello from shu ting ya",
                "jx.friend.name":"mark jackson",
                "jx.friend.age":123,
                "jx.friend.address":"china..."
            }
        }
    ]
}
```



## 资源定位策略

### git backend

这是默认策略。去Git仓库获取配置资源

### File System backend profile

但可以配置成到本地类路径或本地文件系统获取配置资源。需要添加的配置

```yaml
spring:
  profiles:
    active: native
  cloud:
    config:
      server:
        native:
          searchLocations: classpath:/config
```



## 集成既有系统

在服务中添加 `@EnableAutoConfiguration` 注解，将服务相关的配置移到 配置中心

#### 服务端配置

启动类上加：
```java
@EnableConfigServer   // -----------------------
@SpringBootApplication
public class ConfigApplication {
	public static void main(String[] args) {
		SpringApplication.run(ConfigApplication.class, args);
    }
}
```

#### 客户端配置：

启动类上加：

```java
@EnableAutoConfiguration  // --------------  
@SpringBootApplication
public class ZuulGatewayApplication {
	public static void main(String[] args) {
		SpringApplication.run(ZuulGatewayApplication.class, args);
	}
}
```

启动时是这样的：

>**Fetching config from server at : http://127.0.0.1:8082**
>Located environment: name=service-zuul, profiles=[dev], label=service-zuul, version=null, state=null
>Located property source: CompositePropertySource {name='configService', propertySources=[MapPropertySource {**name='classpath:/config/service-zuul-dev.yml'**}]}
>**The following profiles are active: dev**
>Endpoint ID 'service-registry' contains invalid characters, please migrate to a valid format.
>Endpoint ID 'hystrix.stream' contains invalid characters, please migrate to a valid format.
>BeanFactory id=eedf591e-076f-3fab-86d5-7b28b2aa4119
>Bean 'org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration' of type [org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration$$EnhancerBySpringCGLIB$$2fae14b3] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
>**Tomcat initialized with port(s): 8043 (http)**
>Starting service [Tomcat]
>Starting Servlet Engine: Apache Tomcat/9.0.13

服务里剩余的配置：

```yaml
spring: 
  application:
    name: service-zuul
  cloud: 
    config: 
      uri: http://127.0.0.1:8082
      profile: dev
      fail-fast: true
      label: service-zuul
#  profiles:     如果需要，可以通过此属性定义指向特定的文件
#  	active: dev
```

配置中心的 `bootstrap.yml`
```yaml
spring:
  application:
    name: config-server
  profiles:
    active: native
  cloud:
    config:
      server:
        native:
          searchLocations: classpath:/config    
server: 
  port: 8082
```

在配置中心找配置文件的过程：

- 通过`cloud.config.uri`连接配置中心
- 在配置中心的 `bootstrap.yml`中找到配置文件的位置是 `classpath:/config`
- 在 `classpath:/config` 下找 `/{application}/{profile}`,此URL对应 `service-zuul-dev.yml`