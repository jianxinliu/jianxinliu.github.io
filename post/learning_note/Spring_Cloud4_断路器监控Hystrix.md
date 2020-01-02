[TOC]

# Hystrix 断路器

[Hystrix wiki](https://github.com/Netflix/Hystrix/wiki)

断路器就是熔断机制，和电路中的熔断一样。电路中使用熔点较低的材料作为电路中的一部分，当电路发生异常，导致导体温度过高，则熔点低的材料便会被熔断，从而触发断路，保护整体电路及其相关电器。

分布式系统中的断路器也一样，当某处服务发生故障或延迟过久，断路器能马上失败，给请求返回失败，防止因部分失败而影响整个系统。

## 1.Overview

典型的分布式系统都是由多个服务共同协作。服务之间的调用而出现的失败和延迟响应是常有的事，，如果一个服务失败了，可能会影响到其他服务的性能，甚至造成其他服务不可用，更严重的可能会造成整个应用的瘫痪。Hystrix 使得应用变得更有弹性、容错能力更强。

Hystrix 控制服务间通信，提供容错和延迟容错，通过隔离故障服务和阻止故障的级联效应来提高系统的整体弹性。

Hystrix 通过隔离和封装对服务的调用来提供错误和延迟的容错。

```java
// 接口调用失败了，调用实现类 HystrixClientFallback 里的方法
@FeignClient(name = "service-account", fallback = HystrixClientFallback.class )
public interface AccountFeignClient {
     // 通过用户商户编号查询用户账户信息
	@RequestMapping(value = "/account/getById", method = RequestMethod.POST)
	BaseResMessage<AccountDto> getByAccountId(@RequestParam("accountId") Long id);

	// 内部类-断路器
	@Component
	static class HystrixClientFallback implements AccountFeignClient{
		@Override
		public BaseResMessage<AccountDto> getByAccountId(Long id) {
             // 返回预定义信息
			return new BaseResMessage<AccountDto>(0,"失败");
		}
	}
}
```

## 2. Hystrix 应用场景

如果所有服务都正常，那么请求流程是这样的：

![img](https://github.com/Netflix/Hystrix/wiki/images/soa-1-640.png) 

但当其中一个服务产生延迟，就会阻塞所有涉及请求这个服务的请求。应用程序中每一个可能导致网络请求的点都可能导致潜在的失败。比失败更糟糕的是，这些应用程序还可能导致服务之间的延迟增加，从而备份队列、线程和其他系统资源，从而在整个系统中导致更多的级联失败。 

![img](https://github.com/Netflix/Hystrix/wiki/images/soa-2-640.png) 

![img](https://github.com/Netflix/Hystrix/wiki/images/soa-3-640.png) 

## 3.Hystrix 的做法

- 把所有对外部系统或依赖的调用封装在 `HystrixCommand ` 或 `HystrixObservablecCommand` 对象里，这些对象通常都是运行在单独的线程里。
- 限制调用的时长不能超过预定义的阈值。这些阈值可以针对每个依赖进行设置。
- 为每个依赖维护一个小型的线程池（or semaphore），当线程池满了，再发往这个依赖的请求就会被马上拒绝，而不是加入等待队列。
- 度量（measuring）成功、失败（客户端抛出的异常），超时，线程拒绝。
- 如果错误率超过阈值，手动或自动的触发断路器，在一段时间内停止对特定服务的所有请求
- 在请求失败、请求被拒绝、超时、短路时执行回退逻辑。
- 几乎实时的监控指标（metrics）和配置的修改

![img](https://github.com/Netflix/Hystrix/wiki/images/soa-4-isolation-640.png) 

图中的Dependency I 因为某些原因失败，变得不可用，所有对 I 的调用都会失败。当对 I 的调用失败达到一个特定的阀值(5秒之内发生20次失败是Hystrix定义的缺省值), **链路就会被处于open状态**， **之后所有所有对服务 I 的调用都不会被执行， 取而代之的是由断路器提供的一个表示链路open的Fallback消息**. Hystrix提供了相应机制，可以让开发者定义这个Fallbak消息.

**open的链路阻断了瀑布式错误**， 可以让被淹没或者错误的服务有时间进行修复。这个fallback可以是另外一个Hystrix保护的调用, 静态数据，或者合法的空值. Fallbacks可以组成链式结构，所以，最底层调用其它业务服务的第一个Fallback返回静态数据.

## 4. How it works

https://github.com/Netflix/Hystrix/wiki/How-it-Works

用户的请求不再直接访问服务，而是通过线程池中的空闲线程来访问服务。

###4.1 执行流程

下面的执行流程图展示了一个请求发往一个服务的过程。

![img](https://github.com/Netflix/Hystrix/wiki/images/hystrix-command-flow-chart-640.png) 

上图的详细解释：

1. [Construct a `HystrixCommand` or `HystrixObservableCommand` Object](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow1)
2. [Execute the Command](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow2)
3. [Is the Response Cached?](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow3)
4. [Is the Circuit Open?](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow4)
5. [Is the Thread Pool/Queue/Semaphore Full?](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow5)
6. [`HystrixObservableCommand.construct()` or `HystrixCommand.run()`](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow6)
7. [Calculate Circuit Health](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow7)
8. [Get the Fallback](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow8)
9. [Return the Successful Response](https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow9)

### 4.2 断路器



### 4.3 隔离





## 5. How To Use





## 记一个 Feign 中 Hystrix 的问题

写demo测试Feign，目标是触发熔断，但怎么都触发不了，网上说的触发方式是让服务提供者下线（也就是停服务），我试过，一停就报异常，Ribbon 的 LB 说找不到服务提供者（异常信息后面再补）。



说个前提：此时我的classpath 里依然有 Ribbon的依赖，按理说Feign是集成了Ribbon的，不知道这会不会有影响，当然最后成功的时候是移除了 Ribbon的依赖的。

排错之路漫漫啊，一开始怀疑 Eureka 的问题，怀疑 Eureka 迅速的将失效的服务(停止服务) `deregister`，所以 LB 直接报错找不到服务，导致根本没有走到断路，然后就到处找怎样才能让Eureka晚点把不可用的服务清理出注册列表呢。然后就发现了Eureka的自我保护模式。自我保护模式具体是怎样的详见具体笔记。开启自我保护模式之后，依然没有走到断路器，然后就在怀疑是不是保护的时间太短了，哪里参数配置不对。

由来一次重新编译服务消费者时，突然报

```
No fallback instance of type class * found for feign client *
```

就是说feign client 没有 fallback 的实例。我就纳闷了，明明实现了啊。就像这样：

```java
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

import com.jx.service.controller.ICallerClient.HystrixClientFallBack;

@FeignClient(name = "eureka-provider", fallback = HystrixClientFallBack.class)
public interface ICallerClient {
	/**
	 * 测试 Hystrix ，对应的服务提供者会故意延迟返回
	 */
	@RequestMapping(value = "/hystrix", method = RequestMethod.GET)
	String forHystrix();

	/**
	 * inner class . 断路器
	 */
	static class HystrixClientFallBack implements ICallerClient {
		@Override
		public String forHystrix() {
			return "Some thing is broken, it is return from Hystrix!";
		}
	}
}
```

明明 `implements` 了client ，怎么就没有实现呢，我发现`@FeignClient`的来源和我参考的例子不一样，我的是`org.springframework.cloud.openfeign.FeignClient`，而参考的例子是 `org.springframework.cloud.netflix.feign.FeignClient`，我还去找这两者有什么差异。

后来突然灵机一现，拿了这句去检索：

```
No fallback instance of type class * found for feign client *
```

然后就发现了碰到同样问题的人，人家的解决方案是这样的：

1. 开启hystrix(**Feign Client 默认是不会启用 Hystrix 的，需要配置**)

```
feign.hystrix.enabled=true
```

2. Fallback类需要注解@Component

我发现我一个都没做，第一个配置之前没见过没有加还尚可原谅，**第二个就不可饶恕了，完全是粗心**。



想错了，使用断路的是服务消费者，和服务提供者在不在线没关系。实际上，就算服务提供者从来都没有在Eureka中注册过，访问服务消费者时，若发现访问不到服务提供者，也会走断路器。

