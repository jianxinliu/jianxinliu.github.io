[TOC]

# 单例

在整个运行环境中，单例类只允许存在一个对象，实现思路是将对象实例化的权利掌握在单例类自己的手中。

## 立即实例化

```java
public class Singleton{
    private static Singleton instance = new Singleton();
    private Singleton(){}
    public Singleton getInstance(){
		return instance;
    }
}
```

## 懒实例化

```java
public class Singleton{
    private static Singleton instance;
    private Singleton(){}
    public Singleton getInstance(){
        if(instance == null){// 懒初始化
            instance = new Singleton();
        }
        return instance;
    }
}
```

上面的代码不是线程安全的。假设实例化是一个耗时的操作（构造器中 sleep），而且在线程A正在执行  `instance = new Singleton()`时，线程 B 调用 `getInstance`，此时线程 B访问到的 instance 变量是未被初始化的，则线程 B 也会执行一次 初始化操作，这样就有了两个实例。

### DCL（double check locking）

双重锁校验。

```java
public class Singleton{
    private static Singleton instance;
    private Singleton(){}
    public Singleton getInstance(){
        if(instance == null){// 懒初始化
            synchronized(Singleton.class){
                if(instance == null){
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
```

上述代码仍然有存在问题的可能性。因为 `instance = new Singleton()`不是原子性的，这一行代码的执行过程如下：

1. 为 instance 分配内存
2. 初始化 instance（new Singleton()）
3. 将 instance 指向分配好的内存

因为 jvm 有指令重排的特性，所以上面的执行顺序并不能得到保证，若线程 A 的执行顺序为 132，当执行到 13 时，线程 B 进入，**此时线程 B 读取到的 instance  实际上还没被初始化，但已经指向了一块内存，所以线程 B 返回了一个残缺的 instance 对象**，没有被初始化为真正的 Singleton 类。

使用 volatile 关键字可以保证线程间的可见性，保证线程在进行读操作时世界去主存读取而不是从缓存中读取。

# synchronized 关键字底层原理

参考-> [Synchronized](https://github.com/Snailclimb/JavaGuide/blob/master/Java%E7%9B%B8%E5%85%B3/synchronized.md)

## synchronized 块

以如下程序为例

```java
public class SynchronizedDemo {
	public void method() {
		synchronized (this) {
			System.out.println("synchronized 代码块");
		}
	}
}
```

将以上程序编译后，再执行

```shell
$ javap -c -s -v -l SynchronizedDemo.class
# 或者
$ javap -c -s -v -l SynchronizedDemo.class > javap.txt
```

从中可以看到

```java
public void method();
    descriptor: ()V
    flags: ACC_PUBLIC
    Code:
      stack=2, locals=3, args_size=1
         0: aload_0
         1: dup
         2: astore_1
         3: monitorenter             // monitor enter,进入监控
         4: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
         7: ldc           #3                  // String code block
         9: invokevirtual #4                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
        12: aload_1
        13: monitorexit             // monitor exit ,退出监控
        14: goto          22
        17: astore_2
        18: aload_1
        19: monitorexit
        20: aload_2
        21: athrow
        22: return
```

从以上可以看出synchronized 关键字能起到同步作用的原因：

> synchronized 同步语句块的实现使用的是 monitorenter 和 monitorexit 指令，其中 monitorenter 指令指向同步代码块的开始位置，monitorexit 指令则指明同步代码块的结束位置。 当执行 monitorenter 指令时，线程试图获取锁也就是获取 monitor(**monitor对象存在于每个Java对象的对象头中，synchronized 锁便是通过这种方式获取锁的，也是为什么Java中任意对象可以作为锁的原因**) 的持有权.当计数器为0则可以成功获取，获取后将锁计数器设为1也就是加1。相应的在执行 monitorexit 指令后，将锁计数器设为0，表明锁被释放。如果获取对象锁失败，那当前线程就要阻塞等待，直到锁被另外一个线程释放为止。

## synchronized 方法

> synchronized 修饰的方法并没有 monitorenter 指令和 monitorexit 指令，取得代之的确实是 ACC_SYNCHRONIZED 标识，该标识指明了该方法是一个同步方法，**JVM 通过该 ACC_SYNCHRONIZED 访问标志来辨别一个方法是否声明为同步方法，从而执行相应的同步调用**。
>
> 在 Java 早期版本中，synchronized 属于重量级锁，效率低下，因为监视器锁（monitor）是依赖于底层的操作系统的 Mutex Lock 来实现的，Java 的线程是映射到操作系统的原生线程之上的。**如果要挂起或者唤醒一个线程，都需要操作系统帮忙完成，而操作系统实现线程之间的切换时需要从用户态转换到内核态**，这个状态之间的转换需要相对比较长的时间，时间成本相对较高，这也是为什么早期的 synchronized 效率低的原因。庆幸的是在 Java 6 之后 Java 官方对从 JVM 层面对synchronized 较大优化，所以现在的 synchronized 锁效率也优化得很不错了。JDK1.6对锁的实现引入了大量的优化，如自旋锁、适应性自旋锁、锁消除、锁粗化、偏向锁、轻量级锁等技术来减少锁操作的开销。
>



参考文章的后续部分都不错，就不照搬了，直接看原文比较好。

## 实现自己的同步锁

其实实现同步锁的原理非常简单，最简单的只需要一个变量就可以达到控制同步的目的。

```java
/**
 * 模拟锁的实现原理（synchronized）
 * @author ljx
 * @Date Jan 14, 2019 9:21:10 PM
 */
public class Lock {
	/**
	 * 存在于每个对象头中，所以Java中可以使用任何对象作为锁
	 * 原理：
	 * 	线程试图获取锁，也就是试图获取 minitor 的持有权,当计数器的值为 0 时，获取锁成功，
	 * 	获取锁后，将计数器的值加 1 。（在可重入锁中，可以对计数器再执行加的操作）。
	 *  释放锁：将计数器的值减 1 。（在可重如锁中，直到计数器的值为 0 才算释放成功）。
	 */
	static class Monitor{
		public int value = 0;
		public String owner;
		public Monitor() {}
		public Monitor(String o) {
			this.owner = o;
		}
	}
	
	public Monitor monitor = new Monitor();
	
	/**
	 * 获取锁，获取不到返回false
	 * @param o
	 * @return
	 */
	public boolean lock(String o) {
		if(monitor.value != 0) {
//			System.out.println("get lock fail!");
			return false;
		}else {
			monitor.owner = o;
			monitor.value++;
//			System.out.println("the lock is own to:"+o);
			return true;
		}
	}
	
	/**
	 * 获取锁，一定能获取，暂时获取不到就等待,类似自旋锁
	 * @param owner
	 */
	public void lock1(String owner) {
		if(monitor.value != 0) {
//			System.out.println(owner+"...");
			while(monitor.value != 0) {
				// 轮询，每 300 ms 查看一下状态
				try {
					Thread.sleep(300);
				} catch (InterruptedException e) {
					e.printStackTrace();
				}
				System.out.print(owner + ".  ");
			}
		}
		System.out.println(owner+" get...");
		monitor.owner = owner;
		monitor.value++;
	}
	
	public void unlock() {
		monitor.owner = "";
		monitor.value--;
	}
}
```



## 感悟

技术从原始粗放进化到精细化的管理和认知。



竞争锁失败后的处理：1.轮询，查看锁的状态。2,。阻塞挂起，等待通知。这二者都是耗时和耗资源的，实际上都是同步的操作。

Go 语言的哲学:

> 不要通过共享内存来通信，而应该通过通信来共享内存 





信号量。。。基本操作系统原理



# Volatile

**参考《java并发编程艺术——13章》**

是轻量级的 Synchronized ,**在多处理器开发中保证了共享变量的可见性**。可见性就是当一个线程修改一个共享变量时，另外的线程能读到这个被修改的值。

官方定义：

> java 语言允许线程访问共享变量，为了确保共享变量可以被准确一致的更新，线程应该确保通过排它锁单独的获取这个变量。java 语言提供了 volatile,在某些情况下，比锁更方便。如果一个变量被声明成 volatile, java 线程内存模型确保所有线程看到这个变量都是一致的。



## Why volatile

如果使用恰当，volatile 比 synchronized 	的使用和执行成本会更低，因为它不会引起线程的下文切换和调度。



## volatile 的实现原理

通过观察使用 volatile 修饰变量对应的汇编代码，可以发现，有 volatile 变量修饰的共享变量进行写操作时会多出一个 lock 前缀指令，lock 前缀指令在多核处理器下回引发两件事情：

1. 将当前缓存行的数据回写到系统内存
2. 这个写会内存的操作会引起在其他CPU里缓存了该地址的数据无效。

> 处理器为了提高效率，不是直接和内存通讯，而是将内存的数据拿到内部缓存（L1,L2或其他）后再进行操作，但操作完之后不知道何时回写回到内存，如果对声明了 volatile 变量进行写操作， JVM 就会想处理器发出一条 Lock 前缀的指令，将这个变量所在的缓存行的数据写回到内存。但是就算写回到内存，如果其他处理器缓存的值还是旧的，再执行计算操作时机会出问题。所以在多处理器下，为了保证缓存一致，都会实现缓存一致性协议，每个处理器通过嗅探在总线上传播的数据来判断自己缓存的值是否过期了，当处理器发现自己缓存行对应的地址被修改，就会将当前缓存行设置成无效状态，当处理器需要这个数据进行计算时，就会强制重新从内存将新的值加载到内部缓存。

总的来说，volatile  的原理就是：

**Lock 前缀指令会引起处理器缓存回写到内存**

**一个处理器的缓存回写到内存会导致其他处理器的缓存失效**

### 缓存行

> 除了我们在代码中使用的同步锁和jvm自己内置的同步锁外，还有一种隐藏的锁就是缓存行，它也被称为性能杀手。 
> 在多核cup的处理器中，每个cup都有自己独占的一级缓存、二级缓存，甚至还有一个共享的三级缓存，为了提高性能，cpu读写数据是以缓存行为最小单元读写的；32位的cpu缓存行为32字节，64位cup的缓存行为64字节，这就导致了一些问题。 
> 例如，多个不需要同步的变量因为存储在连续的32字节或64字节里面，当需要其中的一个变量时，就将它们作为一个缓存行一起加载到某个cup-1私有的缓存中（**虽然只需要一个变量，但是cpu读取会以缓存行为最小单位，将其相邻的变量一起读入**），被读入cpu缓存的变量相当于是对主内存变量的一个拷贝，也相当于变相的将在同一个缓存行中的几个变量加了一把锁，这个缓存行中任何一个变量发生了变化，当cup-2需要读取这个缓存行时，就需要先将cup-1中被改变了的整个缓存行更新回主存（即使其它变量没有更改），然后cup-2才能够读取，而cup-2可能需要更改这个缓存行的变量与cpu-1已经更改的缓存行中的变量是不一样的，所以**这相当于给几个毫不相关的变量加了一把同步锁**； 
> 为了防止伪共享，不同jdk版本实现方式是不一样的： 
>
> 1. 在jdk1.7之前会 将需要独占缓存行的变量前后添加一组long类型的变量，依靠这些无意义的数组的填充做到一个变量自己独占一个缓存行； 
> 2. 在jdk1.7因为jvm会将这些没有用到的变量优化掉，所以采用继承一个声明了好多long变量的类的方式来实现； 
> 3. 在jdk1.8中通过添加sun.misc.Contended注解来解决这个问题，若要使该注解有效必须在jvm中添加以下参数： 
>   -XX:-RestrictContended
>
> sun.misc.Contended注解会在变量前面添加128字节的padding将当前变量与其他变量进行隔离； 
> 关于什么是缓存行，jdk是如何避免缓存行的，网上有非常多的解释，在这里就不再深入讲解了；
> --------------------- 
> 作者：朱清震 
> 来源：CSDN 
> 原文：https://blog.csdn.net/zqz_zqz/article/details/70233767 
> 版权声明：本文为博主原创文章，转载请附上博文链接！





# 锁的升级

更多-> [java 中的锁 -- 偏向锁、轻量级锁、自旋锁、重量级锁](https://blog.csdn.net/zqz_zqz/article/details/70233767)

java se 1.6为了减少获得锁和释放锁带来的性能损耗，引入了偏向锁和轻量级锁，锁以在 java se 1.6 中，锁有四种状态：无锁状态，偏向锁状态，轻量级锁和重量级锁状态。它会随着竞争情况逐级提升，但不能降级。



## 偏向锁

**缘由：**大多数情况下，不仅不存在多线程竞争，而且总是由同一线程多次获得，为了让线程获得锁的代价更低而引入偏向锁。偏向锁会偏向于第一个访问锁的线程，如果在运行过程中，同步锁只有一个线程访问，不存在多线程争用的情况，则线程是不需要触发同步的，这种情况下，就会给线程加一个偏向锁。

**How to :**当一个线程访问同步块并获取锁时，会先在对象头和栈帧中的锁记录里存储锁**偏向的线程id**，以后该线程在进入和退出同步块时不需要花费CAS（Compare And Swap）操作来获取和释放锁,只需要测试一下对象头里是否存储着指向当前线程的偏向锁，如果测试成功(锁记录里有该线程的 id)，则表明线程已经获得了锁，如果测试失败，则需要再测试一下Mark Word 中的偏向锁的标识是否设置为 1 （表示偏向锁开启），如果没有设置，则使用CAS竞争锁，如果设置了则尝试使用CAS将对象头的偏向锁指向当前线程(锁记录增加一条当前线程 id)。

**注：**Mark Word 在对象头中，存储对象的 hashcode 或锁信息。



简单说，偏向锁在遇到新新线程来获取锁时，~~会先确认眼神，看是不是曾经见过的人，如果是，则不用走流程，直接哭~~。



## 轻量级锁

