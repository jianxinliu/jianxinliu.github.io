[TOC]

[参考](https://github.com/Snailclimb/JavaGuide/blob/master/Java%E7%9B%B8%E5%85%B3/Multithread/AQS.md)

# AQS

AQS （`AbstractQueuedSynchronizer`抽象队列型同步器）,是一个用来构建锁和同步器的框架，使用AQS能简单且高效地构造出应用广泛的大量的同步器，比如ReentrantLock，Semaphore，其他的诸如ReentrantReadWriteLock，SynchronousQueue，FutureTask等等皆是基于AQS的。当然，我们自己也能利用AQS非常轻松容易地构造出符合我们自己需求的同步器。



AQS核心思想是，如果被请求的共享资源**空闲**，则将当前请求资源的线程设置为有效的工作线程，并且将共享资源设置为锁定状态。如果被请求的共享资源被占用，那么就需要一套线程阻塞等待以及被唤醒时锁分配的机制，这个机制AQS是用CLH队列锁实现的，即将暂时获取不到锁的线程加入到队列中。

**AQS定义两种资源共享方式** 

- **Exclusive**（独占）：只有一个线程能执行，如ReentrantLock。又可分为公平锁和非公平锁： 
  - 公平锁：按照线程在队列中的排队顺序，先到者先拿到锁
  - 非公平锁：当线程要获取锁时，无视队列顺序直接去抢锁，谁抢到就是谁的
- **Share**（共享）：多个线程可同时执行，如Semaphore/CountDownLatch。Semaphore、CountDownLatch、 CyclicBarrier、ReadWriteLock 

## Semaphore 信号量

和 `synchronized`、`ReentrantLock` 不同的是，Semaphore 允许多个线程访问同一个资源,常常用作限制对一个资源的访问线程数。

```java
import java.util.concurrent.Semaphore;
/**
 * 模拟公交车(火车)
 * 信号量： 信号量维护了一些许可证，每一个 acquire() 方法在获取到许可证之前，如果当前没有可用的许可  证，就会被阻塞，然后才能拿到许可证。 每个release() 方法添加一个许可证，潜在得释放一个被阻塞的获取者。但是，并没有实际的许可证，Semaphore 只是维护了可用许可证的数量。
 * 信号量常常用来限制访问某些资源（物理的或逻辑的）的线程数量
 * @author ljx
 * @Date Jan 16, 2019 10:34:51 AM
 */
public class Bus {

	private final static int SEAT_MAX = 11;

	private Semaphore tickets = new Semaphore(SEAT_MAX);
	private boolean[] used = new boolean[SEAT_MAX];

    // 上车买票，然后得到一个座位
	public int getSeat() throws InterruptedException {
		tickets.acquire();
		return getAviliableSeat();
	}
	// 下车
	public void debus(int i) {
		markAsUnused(i);
		tickets.release();
	}

	private synchronized int getAviliableSeat() {
		for (int i = 0; i < SEAT_MAX; i++) {
			if (!this.used[i]) {
				used[i] = true;
				return i + 1;
			}
		}
		return -1;
	}

	private synchronized void markAsUnused(int i) {
		this.used[i - 1] = false;
	}
}
```

[参考、更多](https://blog.csdn.net/qq_19431333/article/details/70212663)

Semaphore是信号量，用于管理一组资源。其内部是基于AQS的共享模式，AQS的状态表示许可证的数量，在许可证数量不够时，线程将会被挂起；而一旦有一个线程释放一个资源，那么就有可能重新唤醒等待队列中的线程继续执行。

## CountDownLatch 倒计时器

CountDownLatch是一个同步工具类，它允许一个或多个线程一直等待，直到其他线程的操作执行完后再执行。

### CountDownLatch 的三种典型用法

> ① 某一线程在开始运行前等待n个线程执行完毕。将 CountDownLatch 的计数器初始化为n ：new CountDownLatch(n) ，每当一个任务线程执行完毕，就将计数器减1 countdownlatch.countDown()，当计数器的值变为0时，在CountDownLatch上 await() 的线程就会被唤醒。一个典型应用场景就是启动一个服务时，主线程需要等待多个组件加载完毕，之后再继续执行。
>
> ② 实现多个线程开始执行任务的最大并行性。注意是并行性，不是并发，强调的是多个线程在某一时刻同时开始执行。类似于赛跑，将多个线程放到起点，等待发令枪响，然后同时开跑。做法是初始化一个共享的 CountDownLatch 对象，将其计数器初始化为 1 ：new CountDownLatch(1) ，多个线程在开始执行任务前首先 coundownlatch.await()，当主线程调用 countDown() 时，计数器变为0，多个线程同时被唤醒。
>
> ③ 死锁检测：一个非常方便的使用场景是，你可以使用n个线程访问共享资源，在每次测试阶段的线程数目是不同的，并尝试产生死锁。

#### 应用1

```java
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * CountDownLatch 应用 1：喝茶线程需要等到卖茶叶和烧开水线程完成之后才能执行
 * @author ljx
 * @Date Jan 16, 2019 11:31:13 AM
 */
public class CountDownLatchTest {

	static class DrinkTea implements Runnable{		
		private CountDownLatch cdl;
		public DrinkTea(CountDownLatch cdl) {
			this.cdl = cdl;
		}
		@Override
		public void run() {
			try {
				cdl.await();
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
			System.out.println("drink tea");
		}
		
	}
	
	static class BuyTea implements Runnable{
		private CountDownLatch cdl;
		public BuyTea(CountDownLatch cdl) {
			this.cdl = cdl;
		}
		@Override
		public void run() {
			try {
				Thread.sleep(2000);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
			cdl.countDown();
			System.out.println("go outside buy tea");
		}
	}
	
	static class BoilWater implements Runnable{
		private CountDownLatch cdl;
		public BoilWater(CountDownLatch cdl) {
			this.cdl = cdl;
		}
		@Override
		public void run() {
			try {
				Thread.sleep(3000);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
			cdl.countDown();
			System.out.println("gu lu gu lu...");
		}
	}
	
	public static void main(String[] args) {
		ExecutorService exe = Executors.newFixedThreadPool(3);
		CountDownLatch cdl = new CountDownLatch(2);
		
		exe.execute(new DrinkTea(cdl));
		exe.execute(new BuyTea(cdl));
		exe.execute(new BoilWater(cdl));
		
		exe.shutdown();
	}
}
```

#### 应用2

```java
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * CountDownLatch 应用 2 ：多任务并行，多个线程在同一时刻执行 类似赛跑，发令枪响，所有运动员开始跑
 * 
 * @author ljx
 * @Date Jan 16, 2019 2:07:45 PM
 *
 */
public class CountDownLatchTest2 {

	static class Runner implements Runnable {
		private CountDownLatch cdl;
		private String name;// 跑者姓名
		private int during;// 跑完所花的时间

		public Runner(CountDownLatch cdl, String name, int during) {
			this.cdl = cdl;
			this.name = name;
			this.during = during;
		}

		@Override
		public void run() {
			try {
				// I am Ready !!!
				cdl.await();
				System.out.println(name + " run...");
				Thread.sleep(during); // running hard
				System.out.println(name + " run end....");
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}

	public static void main(String[] args) {
		ExecutorService exe = Executors.newFixedThreadPool(5);
		CountDownLatch cdl = new CountDownLatch(1);
		// Ready !!!
		for (int i = 0; i < 5; i++) {
			exe.execute(new Runner(cdl, "r" + (i + 1), (i + 1) * 1000));
		}
		System.out.println("...... Pang .....");
		// Go !!!
		cdl.countDown();
		exe.shutdown();
	}
}
```

### `CountDownLatch` 的不足

`CountDownLatch`是一次性的，计数器的值只能在构造方法中初始化一次，之后没有任何机制再次对其设置值，当`CountDownLatch`使用完毕后，它不能再次被使用。

## CyclicBarrier 循环栅栏

之所以叫循环，是因为 CyclicBarrier 可以通过 reset 方法进行重用。

 ```java
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * 循环栅栏（屏障）。和 CountDownLatch 非常类似，可以实现线程间的等待，但功能更复杂和强大
 * 作用是，让一组线程达到同一个屏障（同步点）时被阻塞，直到最后一个线程到达同步点，才会打开屏障。
 * 
 * 实现游戏中十个玩家的加载，必须每个玩家都加载完成才会开始游戏。
 * @author ljx
 * @Date Jan 16, 2019 2:41:45 PM
 */
public class CyclicBarrierTest {

	static class Player implements Runnable {

		private CyclicBarrier cb;
		private String name;
		private int loadLatey;

		public Player(CyclicBarrier cb, String name, int ll) {
			this.cb = cb;
			this.name = name;
			this.loadLatey = ll;
		}

		@Override
		public void run() {
			try {
				Thread.sleep(loadLatey);
				System.out.println(name + " Ready....");
				cb.await();
			} catch (InterruptedException e) {
				e.printStackTrace();
			} catch (BrokenBarrierException e) {
				e.printStackTrace();
			}
		}

		public static void main(String[] args) {
			final int TEN = 10;
			ExecutorService exe = Executors.newFixedThreadPool(TEN);
			// CyclicBarrier(int parties, Runnable barrierAction)，指定线程到达屏障时的操作，优先执行barrierAction（parties 个线程）
			CyclicBarrier cb = new CyclicBarrier(TEN,() -> {
				System.out.println("欢迎来到 王者荣耀");
			});
			for (int i = 1; i <= TEN; i++) {
				exe.execute(new Player(cb, "P" + i, i * 1000));
			}			
			exe.shutdown();
		}
	}
}
 ```

## CyclicBarrier 和 CountDownLatch 比较

![CyclicBarrier 和 CountDownLatch 比较](https://camo.githubusercontent.com/5c19d9e66ffaf3d7193b01948279db9b9b3b98d3/68747470733a2f2f6d792d626c6f672d746f2d7573652e6f73732d636e2d6265696a696e672e616c6979756e63732e636f6d2f4a6176612532302545372541382538422545352542412538462545352539312539382545352542462538352545352541342538372545462542432539412545352542392542362545352538462539312545372539462541352545382541462538362545372542332542422545372542422539462545362538302542422545372542422539332f4151533333332e706e67) 

CountDownLatch基于AQS；CyclicBarrier基于锁和Condition。本质上都是依赖于volatile和CAS实现的。 

**参考：**

- <https://blog.csdn.net/u010185262/article/details/54692886>
- <https://blog.csdn.net/tolcf/article/details/50925145?utm_source=blogxgwz0>