[TOC]

# 服务网关

服务网关是微服务中不可或缺的部分，在它统一向外部系统提供RESTful 接口的过程中，还具有服务路由、负载均衡、和权限控制等功能。**相当于微服务系统和外部系统之间的一个代理**，为微服务架构提供前门保护，也将权限控制等重要的非业务逻辑迁移到服务网关层面，使得服务集群主体能够具备高可复用性和高可测试性。

Spring Cloud Netflix 中的 Zuul就是这样的一个服务网关。

### Netflix使用Zuul进行以下操作：

- 认证（权限控制）
- 洞察（insight。监控）
- 压力测试
- 金丝雀测试
- 动态路由
- 服务迁移（微服务和外部系统轻耦合，便于微服务迁移）
- 负载脱落
- 安全
- 静态响应处理
- 主动/主动流量管理

Zuul的规则引擎基本上允许任何JVM语言编写的规则和过滤器，内置Java和Groovy。

### 服务网关是什么

>  服务网关=路由器+过滤器

路由器：转发一切外部请求到后端的微服务

过滤器：在服务网关中可以完成很多横切功能，如权限校验、限流、监控等

## Zuul 过滤器

Zuul 大部分功能是通过过滤器来实现的，其中内置了四个标准过滤器类型，这些过滤器类型对应了典型的请求生命周期。

- **PRE**：这种过滤器在请求被路由之前调用。我们可利用这种过滤器实现身份验证、在集群中选择请求的微服务、记录调试信息等。
- **ROUTING**：这种过滤器将请求路由到微服务。这种过滤器用于构建发送给微服务的请求，并使用Apache HttpClient或Netfilx Ribbon请求微服务。
- **POST**：这种过滤器在路由到微服务以后执行。这种过滤器可用来为响应添加标准的HTTP Header、收集统计信息和指标、将响应从微服务发送给客户端等。
- **ERROR**：在其他阶段发生错误时执行该过滤器。

![Zuul](http://www.ymq.io/images/2017/SpringCloud/zuulFilter/11.png)

允许我们创建自定义的过滤器类型。例如,可以定制一种STATIC类型的过滤器，直接在Zuul中生成响应，而不将请求转发到后端的微服务。  

### 自定义过滤器

通过继承 `ZuulFilter` 类，并覆写其中的一些方法来实现自己的过滤逻辑。

```java
import javax.servlet.http.HttpServletRequest;
import com.netflix.zuul.ZuulFilter;
import com.netflix.zuul.context.RequestContext;
import com.netflix.zuul.exception.ZuulException;

/**
 * ZuulFilter 是Zuul中的核心组件，通过继承该抽象类，覆写几个关键方法以达到自定义调度请求的作用
 * 该过滤器的作用是：如果请求参数中没有 pwd 参数，则不对其进行路由，返回 400
 * 
 * 测试方法：直接访问 Zuul 的地址，这里是：http://localhost:8043/gethello?pwd=123
 * 1. http://localhost:8043/gethello,url 中没有 pwd 参数，不进行路由，得不到结果
 * 2. http://localhost:8043/gethello?pwd=123，URL中有pwd 参数，但是 /gethello 接口并不需要 pwd 参数，仍然没有结果，但是服务网关确实起作用了
 * 3. http://localhost:8043/，没有pwd，但是依然能返回主页，但是 console 显示此请求的返回码是 400，因为 / 并没有进行路由
 * 4. http://localhost:8043/?pwd=123，有 pwd ，正常返回。
 * @author ljx
 * @Date Jan 14, 2019 10:47:43 AM
 *
 */
public class JxFilter extends ZuulFilter{

	/**
	 * 判断是否该执行过滤,原则是：
	 * 	上一个过滤器过滤结果为成功才执行过滤，否则不过滤，直接跳过以下过滤，返回上一个过滤器不通过的过滤结果
	 * 也可以自定义是否过滤的逻辑。
	 */
	@Override
	public boolean shouldFilter() {
//		RequestContext preFilter = RequestContext.getCurrentContext();
//		return (boolean)preFilter.get("isSuccess");
		return true;
	}

	/**
	 * 此过滤器真正执行的方法
	 */
	@Override
	public Object run() throws ZuulException {
		RequestContext ctx = RequestContext.getCurrentContext();
		HttpServletRequest req = ctx.getRequest();
		if(req.getParameter("pwd") != null) {
			System.out.println("pwd:"+req.getParameter("pwd"));
			ctx.setSendZuulResponse(true);
			ctx.setResponseStatusCode(200);
			ctx.set("isSuccess",true);
			return null;
		}else {
			System.out.println("pwdss:"+req.getParameter("pwd"));
			ctx.setSendZuulResponse(false); //不对其进行路由
			ctx.setResponseStatusCode(400);
			ctx.set("isSuccess",false);
			return null;
		}
	}

	/**
	 * 定义过滤类型，比如 return "post"，则表示是 POST 过滤
	 */
	@Override
	public String filterType() {
		return "post";//routing,error,post
	}

	/**
	 * 过滤的优先级，数字越大，优先级越低
	 */
	@Override
	public int filterOrder() {
		return 0;
	}
}
```

实现自己的过滤器之后，需要在应用中开启过滤，在 Application中加入 `@Bean`注解，表示开启。

```java
@EnableZuulProxy
@SpringBootApplication
public class ZuulGatewayApplication {

	public static void main(String[] args) {
		SpringApplication.run(ZuulGatewayApplication.class, args);
	}

	/**
	 * 开启自定义的过滤器
	 * @return
	 */
	@Bean
	public JxFilter jxFilter() {
		return new JxFilter();
	}
}
```



## 集成到既有系统

集成成本非常低，只需要将原先直接访问微服务的方式改为通过网关访问即可，Zuul 会通过Eureka 注册中心去发现可用的服务，并以 LB 的方式去访问。