[TOC]

# Feign 

Feign 是一个声明式的伪Http客户端，默认集成了 Ribbon，并和 Eureka 结合，默认实现了负载均衡的效果。

支持 Hystrix机器fallback,支持Ribbon，从而不需要显式的调用这二者。

**Feign 具有如下特性：**

- 可插拔的注解支持，包括`Feign`注解和`JAX-RS`注解
- 支持可插拔的`HTTP`编码器和解码器
- 支持`Hystrix`和它的`Fallback`
- 支持`Ribbon`的负载均衡
- 支持`HTTP`请求和响应的压缩`Feign`是一个声明式的`Web Service`客户端，它的目的就是让`Web Service`调用更加简单。**它整合了`Ribbon`和`Hystrix`，从而不再需要显式地使用这两个组件**。`Feign`还提供了`HTTP`请求的模板，通过编写简单的接口和注解，就可以定义好`HTTP`请求的参数、格式、地址等信息。接下来，`Feign`会完全代理`HTTP`的请求，我们只需要像调用方法一样调用它就可以完成服务请求。





## 集成到原有系统

