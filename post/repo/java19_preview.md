# Jdk 19 尝鲜

## 运行脚本

因当前好多特性都处于预览状态，所以运行时需要添加参数，使用脚本`run.sh`简化。

```sh
# usage: run.sh ClassName
rm -f *.class

if test $1; then
    javac --release 19 --enable-preview --add-modules jdk.incubator.concurrent $1.java
    java --enable-preview --add-modules jdk.incubator.concurrent $1
fi

sleep 1
rm -f *.class
```

## 基本特性

```java
import java.util.HashMap;

public class First {
	public static void main(String[] args) {
		System.out.println("hello java 19!");

		// 获取预备配好大小的 hashMap
		// 实现原理就是内部进行计算 size / 0.75 (去除 loadFactor 的影响，直接创建指定大小的 hashMap)
		HashMap<String, String> map = HashMap.newHashMap(10);
		System.out.println(map.size());

		// switch 中的模式匹配 (--enable-preview)
		compute(4);
		compute("Aadsf");
		compute("AA");
		compute(new Position(3, 6));
		compute(new Coordinate(5, 3));
		compute(new ServiceAImpl());

		// 使用 swtich 模式匹配实现的策略模式
		Chart chart = new LineChart();
		drawChartStrategy(chart);
    }

	// javac --release 19 --enable-preivew First.java
	// java --enable-preview First
	private static void compute(Object operator) {
		switch (operator) {
		    // 类型 + 条件匹配
			case String s when s.length() > 5 -> System.out.println(s.toUpperCase());
			case String s -> System.out.println(s.toLowerCase());
		    case Integer i -> System.out.println(i * i);
			case Position pos -> System.out.println(pos.x + "-" + pos.y);
			// 类型解构
			case Coordinate(int x,int y) -> System.out.println(x + "-" + y);
		    case ServiceA ser -> ser.say();
			case ServiceB ser -> ser.talk();
			default -> {}
		}
	}

    private static void drawChartStrategy(Chart chart) {
		switch (chart) {
			case LineChart c -> {
				c.type = "line";
				c.upType();
			}
		    case BarChart c -> {
				c.type = "bar";
				c.lowerType();
			}
			default -> {
				chart.type = "line";
			}
		}
		chart.draw();
	}

    public record Position(int x, int y) {}

	public record Coordinate(int x, int y) {}

}

interface ServiceA {
	void say();
}

interface ServiceB {
	void talk();
}

class ServiceAImpl implements ServiceA {
	@Override
	public void say() {
		System.out.println("service A say");
	}
}

class Chart {
    String type;

    void draw() {
		System.out.println("draw chart, type: " + type);
    }
}

class LineChart extends Chart{
    
	void upType() {
		super.type = super.type.toUpperCase();
	}
}

class BarChart extends Chart{
    void lowerType() {
		super.type = super.type.toLowerCase();
	}
}
```

## 结构化并发

```java
import java.time.Duration;
import java.util.Random;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Executor;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

import jdk.incubator.concurrent.StructuredTaskScope;

public class Concurrent {
    static Order loadOrder() throws Exception {
		System.out.println("load order  " + Thread.currentThread().getName());
		Thread.sleep(Duration.ofSeconds(new Random().nextInt(1, 5)));
		return new Order("aa");
    }
    static Customer loadCustomer() throws Exception {
		// throw new Exception(":sdf");
		return new Customer("nu");
    }
    static Inventory loadInventory() throws Exception {
		System.out.println("load inventory");
		Thread.sleep(Duration.ofSeconds(new Random().nextInt(1, 5)));
		return new Inventory(new Random().nextInt(1, 5));
    }

     public static void main(String[] args) throws ExecutionException, InterruptedException {

		// javac --enable-preview -source 19 --add-modules jdk.incubator.concurrent
		// Concurrent.java
		// java --enable-preview --add-modules jdk.incubator.concurrent Concurrent

		// ecah task executed in a virtual thread
		try (var scope = new StructuredTaskScope.ShutdownOnFailure()) {
		    Future<Order> ordeFuture = scope.fork(Concurrent::loadOrder);
		    Future<Customer> customerFuture = scope.fork(Concurrent::loadCustomer);
		    // task canceled
		    Future<Inventory> inventoryFuture = scope.fork(Concurrent::loadInventory);

		    scope.join();

		    // 如果未调用该方法，任务执行过程中有异常会报： Task completed with exception。同时任务也会全部终止执行
		    scope.throwIfFailed();

		    System.out.println(ordeFuture.resultNow());
		    System.out.println(customerFuture.resultNow());
		    System.out.println(inventoryFuture.resultNow());
		}
    }
}

record Order(String oid) {
}

record Customer(String name) {
}

record Inventory(Integer cnt) {
}

```

## 虚拟线程

```java
import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.ThreadLocalRandom;

public class VThread {
    public static void main(String[] args) throws Exception {
		// create virtual thread to run single task
		Thread.startVirtualThread(VThread::platformThread);
		Thread.ofVirtual().start(VThread::vThread);
		
		Thread.currentThread().join(Duration.ofSeconds(20));
    }

     // cost 10098ms
    private static void platformThread() {
		try {
		    threads(Executors.newFixedThreadPool(100), "platform");
		} catch (Exception e) {
		    e.printStackTrace();
		}
    }

    // cost 1423ms
    private static void vThread() {
		try {
		    // create one new virtual thread per task
		    threads(Executors.newVirtualThreadPerTaskExecutor(), "virtual");
		} catch (Exception e) {
		    e.printStackTrace();
		}
    }

    private static void threads(ExecutorService executor, String type) throws Exception {
		List<Task> tasks = new ArrayList<>();
		for (int i = 0; i < 1_000; i++) {
		    tasks.add(new Task(i));
		}

		long time = System.currentTimeMillis();

		List<Future<Integer>> futures = executor.invokeAll(tasks);

		long sum = 0;
		for (Future<Integer> future : futures) {
		    sum += future.get();
		}

		time = System.currentTimeMillis() - time;
		System.out.println(type + " done : sum = " + sum + "; time = " + time + " ms");

		executor.shutdown();
    }
}

class Task implements Callable<Integer> {

    private final int number;

    public Task(int number) {
		this.number = number;
    }

    @Override
    public Integer call() {
		// System.out.printf(
		// "Thread %s - Task %d waiting...%n", Thread.currentThread().getName(),
		// number);
		try {
		    Thread.sleep(1000);
		} catch (InterruptedException e) {
		    // System.out.printf(
		    // "Thread %s - Task %d canceled.%n", Thread.currentThread().getName(), number);
		    return -1;
		}
		// System.out.printf(
		// "Thread %s - Task %d finished.%n", Thread.currentThread().getName(), number);
		return ThreadLocalRandom.current().nextInt(100);
    }
}
```