# I/O

传统的 IO 也叫 BIO，是阻塞式的 IO，读取线程需要等待 IO 完成才返回。还有两种分别是 AIO（异步 IO ） 和 NIO（非阻塞 IO）

## NIO

参考：

- [Java NIO浅析-美团技术团队](https://tech.meituan.com/2016/11/04/nio.html)
- [IBM Developer-Java NIO](https://www.ibm.com/developerworks/cn/education/java/j-nio/j-nio.html)
- [SnailClimb-JavaGuide-Java I/O&NI/O](https://github.com/Snailclimb/JavaGuide/blob/master/Java%E7%9B%B8%E5%85%B3/Java%20IO%E4%B8%8ENIO.md)
- [并发编程网-Java NIO 系列教程](http://ifeve.com/java-nio-all/)
- [云栖社区-理解Java NIO](https://yq.aliyun.com/articles/2371)

NIO（Non-blocking I/O,在 java 领域，也叫 New I/O）,是一种同步非阻塞的 I/O 模型，也是 I/O 多路复用的基础。

**BIO是面向流的，NIO是面向缓冲区的。**

**使用多线程的本质：**

1. 利用多核
2. 当 I/O 阻塞系统，但 CPU 空闲时，可以利用多线程使用 CPU 资源。

NIO核心组件：

1. Channels
2. Buffers
3. Selectors

### Buffer

例子1：

```java
ByteBuffer buffer = ByteBuffer.allocate(1024);
buffer.put(12);
buffer.put(122);
buffer.flip();// 必须有，如果不调用，则读不出数据，因为get、put操作都是将内部 position前移，所以需要 flip 将position重置
buffer.get();// 12
buffer.get();//122
```

例子2：

```java 
FileInputStream fin = new FileInputStream("someFiel.txt");
FileChannel chnl = fin.getChannel();
ByteBuffer buf = ByteBuffer.allocate(1024);
chnl.read(buf);// 给 buffer 赋值
buf.flip();
while(buf.hasRemaining())
	System.out.print((char)buf.get());

// 如果是很大的数据量，需要这样让一个 Buffer 循环的读，只需要在一次读完之后清空 buffer 就可以了
while(chnl.read(buf) != -1){
    buf.flip();
    while(buf.hasRemaining()){
        System.out.print((char)buf.get());
    }
    buf.clear();
}
```

#### Buffer 内部状态控制

Buffer 实际上就是一块内存，这块内存被 NIO Buffer 管理，并提供一系列方法用于更简单的操作这块内存，其底层数据结构是数组。初始状态，position = 0，limit = capacity = 数组容量，且三个指针具有如下关系：$position \le  limit \le capacity$

**position:**

- 当往缓冲区写数据时，position 记录写了多少数据，准确的说应该是 position 指向下一个写入的数据在数组中应该存放的位置

- 在从缓冲区中读取数据时，position 指向下一个读取的数据是来自数组中的那个位置,也就是说 position 能够记录从已经从缓冲区中获取了多少数据。

    总的来说：**position 的作用就是指向下一个被操作的数据的位置**。

**limit:**

- 在写入数据时，limit 记录还有多少空间可供写入（默认limit = capcity）
- 在读取数据时，limit 记录有多少数据需要取出。

在读取之前需要调用 `flip()` 方法转换读写模式，源码如下：

```java
 public final Buffer flip() {
     limit = position;// 将 limit 设置为 position 的位置
     position = 0;// position 设置为 0（开始读取时，第一个读取的数据位于 0 这个位置）
     mark = -1;
     return this;
 }
```

**capacity:**底层数组的大小

#### Buffer 读写

**Reads the byte at this buffer's current position, and then increments the position.**

```java
/**
     * Relative <i>get</i> method.  Reads the byte at this buffer's
     * current position, and then increments the position.
     *
     * @return  The byte at the buffer's current position
     *
     * @throws  BufferUnderflowException
     *          If the buffer's current position is not smaller than its limit
     */
public abstract byte get();
```

**Writes the given byte into this buffer at the current position, and then increments the position.** 

```java
  /**
     * Relative <i>put</i> method&nbsp;&nbsp;<i>(optional operation)</i>.
     *
     * <p> Writes the given byte into this buffer at the current
     * position, and then increments the position. </p>
     *
     * @param  b
     *         The byte to be written
     *
     * @return  This buffer
     *
     * @throws  BufferOverflowException
     *          If this buffer's current position is not smaller than its limit
     *
     * @throws  ReadOnlyBufferException
     *          If this buffer is read-only
     */
    public abstract ByteBuffer put(byte b);
```

Buffer 的读写操作都会操作 position 指针。

### Channel

NIO中所有的IO都是从 Channel 开始的

#### Channel的种类

- FileChannel：用于文件的读写
- SocketChannel & ServerSocketChannel：前者用于 TCP 数据的读写，一般是客户端实现；后者可以监听 TCP 请求，一般是服务器实现
- DatagramChannel：用于 UDP 的数据读写
- Scatter & Gather
    - Scatter：将从一个 Channel 读取的信息分散到多个 Buffer 中去，Buffer 数组，如：`ScatteringByteChannel`
    - Gather：将多个 Buffer 的内容按照顺序发送到一个 Channel,如：`GatheringByteChannel`

Channel 是双向的，可读也可写，可以异步读写，一般都是基于Buffer进行读写

### Selector

选择器（多路复用器）。用于检查一个或多个 NIO Channel（通道）的状态是否处于可读、可写。

使用 Selector 的好处在于： 使用更少的线程来就可以来处理通道了， 相比使用多个线程，避免了线程上下文切换带来的开销。

#### Selector的使用方法

- Selector 的创建

```
Selector selector = Selector.open();
```

- 注册 Channel 到 Selector(Channel必须是非阻塞的)

```
channel.configureBlocking(false);
SelectionKey key = channel.register(selector, Selectionkey.OP_READ);
```

- SelectionKey 介绍

    一个 SelectionKey 键表示了一个特定的通道对象和一个特定的选择器对象之间的注册关系。

- 从 Selector 中选择 channel(Selecting Channels via a Selector)

    选择器维护注册过的通道的集合，并且这种注册关系都被封装在 SelectionKey 当中.

- 停止选择的方法

    wakeup() 方法和 close() 方法。





### 内存映射

这个功能主要是为了提高大文件的读写速度而设计的。内存映射文件(memory-mappedfile)能让你创建和修改那些大到无法读入内存的文件。有了内存映射文件，你就可以认为文件已经全部读进了内存，然后把它**当成一个非常大的数组来访问**了。将文件的一段区域映射到内存中，比传统的文件处理速度要快很多。内存映射文件它虽然最终也是要从磁盘读取数据，但是它并不需要将数据读取到OS内核缓冲区，而是直接将进程的用户私有地址空间中的一部分区域与文件对象建立起映射关系，就好像直接从内存中读、写文件一样，速度当然快了。