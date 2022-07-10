参考书[《Redis设计与实现》黄健宏]( https://book.douban.com/subject/25900156/ )、[Redis in Action]( https://book.douban.com/subject/10597898/ )

# 一、数据结构与对象

## SDS

- Redis 只会使用 C 字符串作为字面量，在大多数情况下，Redis 使用 SDS（Simple Dynamic String，简单动态字符串）作为字符串表示。
- 比起 C 字符串，SDS 具有以下优点：
  1. **常数复杂度获取字符串长度**。
  2. 杜绝缓冲区溢出。
  3. 减少修改字符串长度时所需的内存重分配次数。
  4. **二进制安全**。
  5. 兼容部分 C 字符串函数。

## 链表

- 链表被广泛用于实现 Redis 的各种功能，比如列表键、发布与订阅、慢查询、监视器等。
- 每个链表节点由一个 listNode 结构来表示，每个节点都有一个**指向前置节点和后置节点的指针**，所以 Redis 的链表实现是**双端链表**。
- 每个链表使用一个 list 结构来表示，这个结构带有**表头节点指针、表尾节点指针，以及链表长度等信息**。
- 因为链表表头节点的前置节点和表尾节点的后置节点都指向 NULL，所以 Redis 的链表实现是**无环链表**。
- 通过为链表设置不同的类型特定函数，Redis 的链表可以用于**保存各种不同类型的值**。

## 字典

- 字典被广泛用于实现 Redis 的各种功能，其中包括数据库和哈希键。
- Redis 中的字典使用**哈希表作为底层实现**，**每个字典带有两个哈希表，一个平时使用，另一个仅在进行 rehash 时使用**。
- 当字典被用作数据库的底层实现，或者哈希键的底层实现时，Redis 使用 MurmurHash2 算法来计算键的哈希值。
- 哈希表使用**链地址法来解决键冲突**，被分配（**头插**）到同一个索引上的多个键值对会连接成一个单向链表。
- 在对哈希表进行扩展或者收缩操作时，程序需要将现有哈希表包含的所有键值对 rehash 到新哈希表里面，并且这个 rehash 过程并不是一次性地完成的，而是**渐进式地完成的**。

重点：

1. 渐进式 rehash。
   1. 指的是 rehash 状态会持续一段时间，并且是非连续的。也就是说 rehash 的过程会**为了防止因 rehash 导致服务器停止服务而主动中断**，**分多次、渐进式的完成 rehash 的过程**。rehash 的状态由 `rehashidx` 字段指明，为 -1 时表示不是 rehash 状态，其余数字则指示当前 rehash 在 `ht[0]` 中进行到的位置。在渐进式 rehash 期间，每一次对字典的 CRUD 操作都会触发 rehash 操作。
   2. **rehash 操作的实质**。字典内部维护两个 哈希表 `ht[0]` & `ht[1]`，其中，后者用于 rehash，前者用于平常存储数据。rehash 时，分多次的将 `ht[0]` 中的数据 rehash 到 `ht[1]` 中，全部迁移完成之后，释放`ht[0]`， 将 `ht[1]` 设置为 `ht[0]`，再新创建一个哈希表挂在 `ht[1]`上以备下次 rehash。
2. 链地址解决冲突。使用头插法加入冲突的新数据。为什么？因为因冲突产生的链表没有尾指针，若实行尾插，则时间复杂度为 O(n)，而查找时又由于程序的时间局部性原理，刚插入的数据很快就会被再次访问，所以会有更大的概率去访问表尾，时间复杂度也是 O(n)；而**使用头插法，插入和查询时间复杂度都是常数级的**。

## 跳表

- 跳跃表是有序集合的底层实现之一。
- Redis 的跳跃表实现由 zskiplist 和 zskiplistNode 两个结构组成，其中 zskiplist 用于保存跳跃表信息（比如表头节点、表尾节点、长度），而 zskiplistNode 则用于表示跳跃表节点。
- 每个跳跃表节点的层高都是 1 至 32 之间的随机数。（由抛硬币算法决定）
- 在同一个跳跃表中，多个节点可以包含相同的分值，但每个节点的成员对象必须是唯一的。
- 跳跃表中的节点按照分值大小进行排序，当分值相同时，节点按照成员对象的大小进行排序。

[跳表参考]( https://mp.weixin.qq.com/s?src=11&timestamp=1583242386&ver=2194&signature=zHw6vFj4Y0shs0BmOMFYpnptLix4*d04inIFwbKI5nUEDz9YtEj0TwH1PmCZLSnrI5djjcUfY9Gd4KJX*nT0hik8mH0nsg-QYb6p6JpBntzTCxDPF1MIh3XUutLDjbpc&new=1 )

在大部分情况下，跳跃表的效率可以和平衡树相媲美，并且因为跳跃表的实现比平衡树要来得更为简单，所以有不少程序都使用跳跃表来代替平衡树。

## 整数集合

- 整数集合是集合键的底层实现之一。
- 整数集合的底层实现为数组，这个数组以有序、无重复的方式保存集合元素，在有需要时，程序会根据新添加元素的类型，改变这个数组的类型。
- 升级操作为整数集合带来了操作上的灵活性，并且尽可能地节约了内存。
- 整数集合只支持升级操作，不支持降级操作。

## 压缩列表

- 压缩列表是一种为节约内存而开发的顺序型数据结构。
- 压缩列表被用作列表键和哈希键的底层实现之一。
- 压缩列表可以包含多个节点，每个节点可以保存一个字节数组或者整数值。
- 添加新节点到压缩列表，或者从压缩列表中删除节点，可能会引发连锁更新操作，但这种操作出现的几率并不高

## 对象

Redis 并没有直接使用以上的数据结构实现数据库，而是在此基础上构建了一个对象系统，包含：**字符串对象、列表对象、哈希对象、集合对象和有序集合对象**。

- Redis数据库中的每个键值对的键和值都是一个对象。
- Redis共有字符串、列表、哈希、集合、有序集合五种类型的对象，每种类型的对象至少都有两种或以上的编码方式，不同的编码可以在不同的使用场景上优化对象的使用效率。
- 服务器在执行某些命令之前，会先检查给定键的类型能否执行指定的命令，而检查一个键的类型就是检查键的值对象的类型。
- Redis的对象系统带有引用计数实现的内存回收机制，当一个对象不再被使用时，该对象所占用的内存就会被自动释放。
- Redis会共享值为0到9999的字符串对象。
- 对象会记录自己的最后一次被访问的时间，这个时间可以用于计算对象的空转时间。

# 二、单机数据库的实现

## 数据库

默认情况下，Redis 服务器会创建 16 个数据库（配置文件中的 `databases 16`指定），**每个客户端都有自己的目标数据库**，所以 Redis 命令是有操作对象的。默认情况下，客户端的目标数据库为 0 号，但可以通过 `SELECT` 命令切换数据库，不同数据库间数据不共享。如：

```sh
127.0.0.1:6379> set name jx # 默认 0 号数据库
OK
127.0.0.1:6379> get name
"jx"
127.0.0.1:6379> select 1 # 切换至 1 号数据库
OK
127.0.0.1:6379[1]> get name
(nil) # 获取不到值
127.0.0.1:6379[1]> set name kl
OK
127.0.0.1:6379[1]> get name
"kl"
127.0.0.1:6379[1]> select 0 # 切换回 0 号
OK
127.0.0.1:6379> get name
"jx" # 0 号的值
127.0.0.1:6379>
```

目标数据库是由客户端`redisDb *db`指针维护的，记录自己连接的是哪个数据库。所以 `SELECT` 命令的本质就是修改这个指针，使其指向不同的数据库。

注意：

- 多个客户端连接，都是默认连接 0 号数据库的，而不会自动递增的选择其目标数据库。故**使用编程语言客户端时，要明确连接的是哪个数据库。**
- redis 数据库更像是一种**命名空间**，不支持单独加密码，不支持单独命名，只以索引的方式区分不同数据库。
- redis 不同数据库不适宜存储不同应用的数据，但可以存储不同环境的数据。如：0 号存储生产数据，1 号存储测试数据。但不适宜 0 号存储 A 应用数据，1 号存储 B 应用数据。因为 **redis 并没有在不同的数据库上进行权限控制，而是使用同一个权限**。不同应用的数据应该使用不同的 redis 实例进行存储。
- 单体情况下才支持数据库切换，集群模式下，只有一个数据库 db0 ，故不支持切换

### 数据库键空间

redis 数据库中使用 dict 保存所有的键值对，故称这个 dictionary 为键空间。对数据库的 CRUD 和其余一些针对 redis 本身的操作实际上都是对这个字典的操作。

**键空间的键就是数据库的键**，每个键都是一个字符串对象；**键空间的值就是数据库的值**，可以是字符串对象、列表对象、哈希表对象、集合对象和有序集合对象中的任意一种 Redis 对象。

### 键过期及删除策略

可以对键空间的键设置生存时间，以秒或者毫秒计，键过期之后，服务器自动删除生存时间为 0 的键。过期时间存储在过期字典中。主要有以下**三种过期键删除策略**：

1. **定时删除**。<u>在设置键过期时间的同时设置定时器，定时器在键过期时间来临时执行键删除。</u>虽然定时删除会及时的删除过期键，对内存友好，但**当内存不紧张而 CPU 时间紧张时，定时删除无疑是给 CPU “添乱”**，特别是在过期键较多的情况下。并且创建定时器需要用到 redis 服务器中的时间事件，其实现方式是无序列表，查找一个事件的时间是 O(n)，因此并不能高效的处理大量时间事件。
2. **惰性删除**。<u>放任键过期不管，但当每次从键空间中获取键时，都检查键是否过期，若过期则删除。</u>对 CPU 友好而对内存不友好。可以保证只有在过期键非删除不可的情况下才删除，且只涉及当前键，故对 CPU 友好。但缺点是过期键依然会占用内存，特别是存在大量过期键且一段时间内没有被访问到的情况。因为无用的垃圾占用了大量内存，甚至可以看做是一种内存泄漏。
3. **定期删除**。<u>每隔一段时间，对数据库进行检查，删除过期键。</u>是上两种策略的整合和折中。每隔一段时间进行过期键删除操作，并限制操作的时长和频率来减少对 CPU 时间的影响。**难点在于确定操作的时长和频率**，这是一个需要权衡的事情。

实际上， **Redis 服务器使用惰性删除和定期删除两种策略**。

### RDB、AOF和复制功能对过期键的处理

#### RDB

写入：当执行 `SAVE` 或者 `BGSAVE` 创建新的 RDB 文件时，已过期的键不会被保存进新的 RDB 文件中。

载入：当 redis 服务器载入 RDB 文件时：

- 如果服务器以主模式运行，则会对键是否过期进行检查，过期键不会被载入。
- 如果服务器以从模式运行，则不会对键是否过期进行检查，即所有键都会被载入数据库。但当主从同步时，从数据库的数据会被清空，故从服务器不会受 RDB 文件中过期键的影响。

#### AOF

写入：只要键还在键空间中，不论是否过期，都会被 AOF 文件记录。当过期键被惰性删除或定时删除时， AOF 会被追加一条 `DEL `命令，来显示记录键已被删除。所以过期键就算被保存至 AOF 文件，也会被记录删除。

重写：同主模式下的 RDB 写入策略，过期键不会被重写到新的 AOF 文件。

**AOF 重写：一种不停机更新 AOF 文件的机制，重写后的 AOF 文件包含重建当前数据集需要的最少命令。重写之后的精简 AOF 文件就不会无限增长了。**实际上是**根据当前数据库状态来重新写一个 AOF 文件并重命名**，而不管之前 AOF 文件的状态，也不需要对其进行读写。重写时，直接读取键空间的键和值，用一条命令记录键值对。

#### 复制

当服务器运行在复制模式下，过期键的删除动作由主服务器控制。主服务器在删除时，会向所有从服务器发送 `DEL` ，从服务器才删除过期键，否则不删除，这可以**保证主从一致**。**从服务器在遇到对过期键的请求时，不会删除，而是会无视过期与否，依旧返回值**。这是复制模式设计存在的缺陷，需要使用者衡量选择使用哪种模式。

## RDB 持久化

redis 快照持久化方式，可手动触发快照，也可设置定期保存快照。手动触发快照的命令是 `SAVE` & `BGSAVE`。这二者的区别是前者会**阻塞 redis 服务器进程，客户端发送的所有命令都会被拒绝**，直到 rdb 文件创建完毕，而后者会创建子进程在后台执行，不阻塞服务器。在 `BGSAVE`执行期间，若客户端再次发送 `SAVE` 则会被拒绝，是为了防止父进程和子进程同时执行 `rdbSave`而产生竞争条件，这是不允许的。同样的，再次发送`BGSAVE` 也会被拒绝，因为也会产生竞争条件。除此之外，`BGSAVE` 和 `BGREWRITEAOF` 也是不能同时执行的。**因为二者都是由子进程执行，都会产生大量的磁盘写入操作，于性能不友好**。

RDB 文件的读入是在 redis 服务器启动时，检测到 RDB 文件的存在，自动完成的。

相较于 AOF 文件的高优先级，只有在未开启 AOF 模式下，服务器会使用 RDB 文件恢复数据库。

另外可以通过设置时间间隔自动触发快照生成。如设置 `save 900 1` 则表示在 900 秒内至少有一次数据库修改，则执行 `BGSAVE` 。

## AOF 持久化

Append Only File 方式，通过保存 redis 服务器执行的命令来记录数据库的状态。AOF 持久化的效率和安全性解释：

```
# The fsync() call tells the Operating System to actually write data on disk
# instead of waiting for more data in the output buffer. Some OS will really flush
# data on disk, some other OS will just try to do it ASAP.

# Redis supports three different modes:

# no: don't fsync, just let the OS flush the data when it wants. Faster.
# always: fsync after every write to the append only log. Slow, Safest.
# everysec: fsync only one time every second. Compromise.

# The default is "everysec", as that's usually the right compromise between
# speed and data safety. It's up to you to understand if you can relax this to
# "no" that will let the operating system flush the output buffer when
# it wants, for better performances (but if you can live with the idea of
# some data loss consider the default persistence mode that's snapshotting),
# or on the contrary, use "always" that's very slow but a bit safer than
# everysec.

appendfsync everysec

——来自 redis.conf
```

### RDB 和 AOF 效率与安全性辩证

RDB 的快照方式可以将数据恢复到某一时间点，而在快照保存时间间隔内的数据则不可恢复，这就有很大的安全性问题（if a crush were to happen before a snapshot,you'd lose any data changed since the last snapshot）。而 AOF 的方式虽然几乎保证最少的数据丢失，但将日志的写磁盘的速度缓慢，且 AOF 文件容易体积膨胀，浪费存储空间且启动时加载时间长。**故常常将二者结合起来，先使用 RDB 文件将数据库恢复到某一时间点，再使用 AOF 恢复部分快照时间间隔内的数据。**

## 事件

Redis 服务器是一个事件驱动程序，服务器需要处理以下两类事件：

- 文件事件（file event）：Redis 服务器通过**套接字**与客户端（或者其他 Redis 服务器）进行连接，而**文件事件就是服务器对套接字操作的抽象**。服务器与客户端（或者其他服务器）的通信会产生相应的文件事件，而服务器则通过监听并处理这些事件来完成一系列网络通信操作。
- 时间事件（time event）：Redis 服务器中的一些操作（比如 serverCron 函数）需要在给定的时间点执行，而时间事件就是服务器对这类定时操作的抽象。

重点：

- Redis 服务器是一个事件驱动程序，服务器处理的事件分为**时间事件**和**文件事件**两类。
- 文件事件处理器是基于 **Reactor** 模式实现的网络通信程序。
- 文件事件是对套接字操作的抽象：每次套接字变为可应答（acceptable）、可写（writable）或者可读（readable）时，相应的文件事件就会产生。
- 文件事件分为 AE_READABLE 事件（读事件）和 AE_WRITABLE 事件（写事件）两类。
- 时间事件分为定时事件和周期性事件：定时事件只在指定的时间到达一次，而周期性事件则每隔一段时间到达一次。
- 服务器在一般情况下只执行 serverCron 函数一个时间事件，并且这个事件是周期性事件。
- 文件事件和时间事件之间是合作关系，服务器会轮流处理这两种事件，并且处理事件的过程中也不会进行抢占。
- 时间事件的实际处理时间通常会比设定的到达时间晚一些。

待深入。

## 客户端

……

## 服务器

[Redis 单线程]( https://zhuanlan.zhihu.com/p/34438275 )

Redis 为什么这么快：

1. 完全基于内存。
2. 采用单线程，避免了不必要的上下文切换，也不用考虑锁的问题，及其加锁放锁死锁产生的问题。
3. 多路 I/O  复用模型，非阻塞 IO。

**多路 I/O 复用模型**：多路指的是多个网络连接，复用指的是复用同一个线程。

**单线程**：指的是**在处理网络请求时只有一个线程**。

……

### 服务器初始化

一个Redis服务器从启动到能够接受客户端的命令请求，需要经过一系列的初始化和设置过程，比如初始化服务器状态，接受用户指定的服务器配置，创建相应的数据结构和网络连接等等。启动步骤：

1. 初始化服务器状态结构。对服务器进行一些默认配置，如端口号、运行架构、持久化方式、创建命令表……此处的配置可被下一步覆盖。
2. 载入配置选项。应用启动服务器时指定的参数或者 `redis.conf` 文件的配置。
3. 初始化服务器数据结构。除了第一步创建的命令表，此处还会创建:
   1. `server.clients`。记录连上的客户端。
   2. `server.db`。记录服务器包含的数据库。
   3. ……
   4. 执行完毕后，打印 redis logo
4. 还原数据库状态。载入 RDB 或 AOF 文件，回复数据库。若开启 AOF 则优先使用 AOF 文件还原数据库，否则使用 RDB 文件还原。
5. 执行事件循环，等待接收客户端连接。

# 三、多机数据库的实现

## 复制

可以通过执行 `SLAVEOF` 命令或者设置 `slaveof` 选项，让一个服务器去复制（replicate）另一个服务器，我们称被复制的服务器为主服务器（master），而对主服务器进行复制的服务器则被称为从服务器（slave）。

实验：

```sh
redis-server
# 默认启动 6379

---------------------
redis-server --port 6380
# 启动第二个实例

---------------------
redis-cli -h 127.0.0.1 -p 6379
slaveof 127.0.0.1 6380
# 6379 为从，6380 为主

---------------------
redis-cli -p 6380
set name jianxin

---------------------
# 切回 6379 的客户端，展示主从复制
keys *
> name
```

### 旧版复制功能的实现

复制功能分为**同步**和**命令传播**两个操作：

- 同步操作用于将从服务器的数据库状态更新至主服务器当前所处的数据库状态。
- 命令传播操作则用于在主服务器的数据库状态被修改，导致主从服务器的数据库状态出现不一致时，让主从服务器的数据库重新回到一致状态。

#### **同步过程：**

1. slave 发送 `SYNC` 给 mater
2. master 接收到命令，执行 `BGSAVE`，生成一个 RDB 文件，并将执行期间接收到的写命令保存至**缓冲区**。
3. master 持久化完毕后，将生成的 RDB 文件发送给 slave，slave 通过 RDB 文件重建数据库至 master 执行 `BGSAVE` 时的状态。
4. master 再将缓冲区内的写命令发送给 slave ，slave 执行这些写命令，将自身数据库状态同步至 master 当前的状态。

#### **命令传播**

即是 master 执行了导致数据库状态发生变化的命令，导致主从不一致，此时 **master 需要将自身执行的<u>命令传播</u>给 slave**，slave 同样执行了相同的命令之后，再次回到主从一致的状态。

#### **旧版复制功能的缺陷**

当处于命令传播阶段（slave 正常的接收 master 的命令并执行），slave 断线导致主从不一致，当 slave 再次连接上 master 时，执行 `SYNC` 命令实现主从一致。此时问题出现，假设在断线前，双方保持主从一致的状态很久，生成的 RDB 文件巨大，而断线到重新连上， master 只执行了三个命令，也就是说 slave 只落后了三个命令，要实现主从一致， slave 只需要再执行这三个命令即可，但是 `SYNC` 命令会让 master 产生一个完整的 RDB 文件，里面包含了巨量却对 slave 同步来说不必要的键值对，slave 和 master 都需要做很多无用功。 

所以 **`SYNC` 命令是一个非常耗费资源的操作。**体现如下：

1. master 需要生成巨大的 RDB 文件，会**耗费 CPU、内存和磁盘 I/O资源**。
2. master 需要将生成的 RDB 发送给 slave，会**耗费大量网络资源（带宽和流量）**。
3. slave 接收到巨大的 RDB，**载入数据库需要很长的时间**，且处于**阻塞不能接受命令**的状态。

### 新版复制功能的实现

为了解决以上问题，只需要增加一个**部分重同步**的功能用于处理断线后重复制即可。部分重同步：master 将 slave 断开期间执行的写命令发送给 slave 即可。**Redis 2.8 版本**开始，使用 `PSYNC` 代替 `SYNV`，`PSYNC` 有两种模式：**完整重同步**和**部分重同步**。完整重同步和 `SYNC` 基本一致。 

**部分重同步的实现**：

1. 我的设想：master 监测 slave 的状态，一旦哪个断线，就记录“slave x 在  a 号命令出断线”，当 slave x 向 master 发送`PSYNC` 时，mater 识别是哪个 slave，并发送器对应缺失数据。这样好麻烦啊，master 这么累，还是 master 吗？
2. 实际上：master 和 slave 共同维护自己所发送/接收的数据偏移量，如 master 的 `offset=10086`，则各个 slave 的 `offset =10086`，此时 master 向各个 slave 发送长度为 33 字节的数据，则各自的 `offset = 10119`。通过对比偏移量就很容易知道是否处于主从同步状态。并且 master 需要再向 slave 发多少数据也很容易计算出来。

实际上，master 为了发送 slave 缺失的数据，还维护了一个**复制积压缓冲区**，大小默认为 1 M，是一个固定长度的字节队列。master 向 slave 传播的命令都会在入这队列（因为定长，最先入队的数据会因队长不够而出队），队列也维护了偏移量， slave 请求 `PSYNC` 时，发送自己的 offset，master 进行比较，若缺失的数据还在缓冲区内（缺失 1 M 以内的数据），则 master 回复 `+COUNTINE` ，表示数据同步将以部分重同步的方式进行，master 从缓冲区内取出缺失的数据发回；若确实的数据不完全在缓冲区内（缺失 1 M 以上的数据），则 master 执行完整重同步。

即为关键的是复制积压缓冲区的大小设置，太小则不能发挥出部分重同步的作用。配置说明(redis.conf)：

```
# Set the replication backlog size. The backlog is a buffer that accumulates
# slave data when slaves are disconnected for some time, so that when a slave
# wants to reconnect again, often a full resync is not needed, but a partial
# resync is enough, just passing the portion of data the slave missed while
# disconnected.

# The bigger the replication backlog, the longer the time the slave can be
# disconnected and later be able to perform a partial resynchronization.

# The backlog is only allocated once there is at least a slave connected.

# repl-backlog-size 1mb
```

除了 offset 和 repl-backlog 之外，实现部分重同步还需要**服务器 ID（run ID）**。是用来在同步时判别 slave 是第一次连上这个 master ，还是断线重连，**如果是第一次连，那直接执行完整重同步，若是断线重连，则可以尝试部分重同步**。**实现原理**是：master 会在 slave 第一次复制时发送自己的 run ID 给 slave 保存，当 slave 断线重连某一个 master 时会带上 run ID，master 进行检查，若和自己相同，则说明该 slave 是断线重连，若不相同，则说明该 slave 是第一次连。

### 复制的实现步骤

1. 设置主服务器的地址和端口
2. 建立套接字连接
3. 发送PING命令确认连接顺畅
4. 身份验证
5. 发送端口信息，用于 master 打印而已
6. 同步（复制工作）
7. 命令传播（复制完成，交流常态化）

### 心跳检测

在命令传播阶段，slave 每秒都会向 master 发送`replconf ack <slave_replcation_offset>`。有三个作用：

#### 1. 检测主从之间的网络连接。

master 会记录 slave 上次发送 `replconf ack` 到当前的时间，用 `lag=0` 表示 salve 刚发送过，其他正数 x 则表示 x 秒之前发送过。若长时间未发送，则说明主从之间的网络连接出现故障。

#### 2. 辅助实现 min-slaves 选项。

```
# It is possible for a master to stop accepting writes if there are less than
# N slaves connected, having a lag less or equal than M seconds.

# The N slaves need to be in "online" state.

# The lag in seconds, that must be <= the specified value, is calculated from
# the last ping received from the slave, that is usually sent every second.

# This option does not GUARANTEE that N replicas will accept the write, but
# will limit the window of exposure for lost writes in case not enough slaves
# are available, to the specified number of seconds.

# For example to require at least 3 slaves with a lag <= 10 seconds use:

min-slaves-to-write 3
min-slaves-max-lag 10
```

#### 3. 检测命令丢失。

因为 `replconf ack` 命令会携带 slave 的 offset，若出现 master 发送给 slave 的命令丢失，则会造成 slave 的 offset 与 master 的不一致。在 slave 向 master 发送 `replconf ack` 命令时，master 会检测到命令丢失，则会将 salve 缺失的数据重新发送。这个补发的过程和部分重复制机制实现原理非常相似，主要差异在于此时的补发数据是在 slave 没有断线的情况下发生的。

**补发机制**是在 Redis 2.8 之后才有的，故为了保证复制时主从一致，需要确保使用 2.8 及以上的版本。

## Sentinel

哨兵，是 Redis 的高可用解决方案。由一个或多个 sentinel 实例组成 sentinel 系统，用于监测任意多个 master 及其下属 slave，主要作用是**在 master 下线时，自动将其下属的某个 slave 升级为 master 继续服务**，同时设置被选中的 slave 为其他 salve 的 master，并继续监测下线的旧 master，若重新上线，则将其作为新 master 的 slave。

```sh
# 启动 sentinel
redis-sentiel /path/to/sentinel.conf
# or
redis-server /path/to/sentinel.conf --sentinel
```

当一个Sentinel启动时，它需要执行以下步骤：

1. 初始化服务器。sentinel 本质上是一个运行在特殊模式下的 redis 服务器，但是 sentinel 不载入数据库文件。
2. 将普通 Redis 服务器使用的代码替换成 Sentinel 专用代码。sentinel 特有的数据结构和命令。
3. 初始化 Sentinel 状态。
4. 根据给定的配置文件，初始化 Sentinel 的监视主服务器列表。
5. 创建连向主服务器的网络连接。

### 获取主从服务器信息

sentinel 通过向主从服务器**每十秒一次的频率**发送 `INFO` 命令获取主从服务器的信息。对于 master ，`INFO` 返回的信息足够 sentinel 知晓 runID,role,slaves 等信息；对于 salve，`INFO` 返回的信息足够 sentinel 知晓 runID ,role ,master_host&master_port ,salve_repl_offset 等信息。

### 向主从服务器发送信息

默认情况下，sentinel 以**两秒一次的频率**通过命令连接向所哟被监视的主从服务器发送该格式的命令：

```shell
PUBLISH __sentinel__:hello "<s_ip>,<s_port>,<s_runid>,<s_epoch>,<m_name>,<m_ip>,<m_port>,<m_eroch>"
```

该命令向服务器的`__sentinel__:hello` 频道发送一条消息。其中 s 开头的是 sentinel 自身的信息，m 开头的是 master 的信息。若 sentinel 正在监视的是 slave ，则 m 开头的信息是指 slave 正在复制的 master 的信息。

### 接收来自主从服务器的频道信息

当 Sentinal 与主从服务器建立订阅连接后，Sentinel 会通过订阅连接，向服务器发送

```shell
SUBSCRIBE __sentinel__:hello
```

即订阅 `__sentinal__:hello` 这个频道，订阅关系会一直持续到与服务器断开连接为止。

#### 多个 Sentinel

若多个 sentinel 监视同一个服务器时，其中任何一个 sentinel 向服务器的 `__sential:hello__` 频道发送一条信息，所有订阅了该频道的 sentinel 都会收到这条信息（包括发送信息的 sentinel 自己，但自己会通过对比 sentinel 运行 ID 来对比是否是自己，若是则丢弃）。

**监视同一个服务器的多个 Sentinel 之间会通过发送和分析频道信息来自动感知其他 sentinel 的存在**，所以可以创建命令连接，组成 sentinel 网络。

### 检测主观下线状态

检测对象包括 master、slave和其他 sentinel。主观下线是指被检测对象的主动下线。标志是被检测对象在 `down-after-milliseconds` 时间内，持续向 sentinel 返回无效回复，则 sentinel 认为该实例进入主观下线状态。用户设置的 `down-after-milliseconds` 的值，不仅会被 sentinel 用来判断 master 的主观下线状态，而且也会用来判断 master 下属的 slave 和监视 master 的所有 sentinel 的主观下线状态。

默认情况下，sentinel 会以每秒一次的频率向所有与它创建了命令连接的实例（即被观测对象）发送 PING 命令，根据实例返回的 回复判断实例是否在线。

1. 有效回复：`+PONG`,`-LOADING`,`-MASTERDOWN`
2. 无效回复：除**有效回复之外的回复**或在**指定时间内没有回复**。

### 检测客观下线状态

若主服务器客观下线，则 sentinel 可以将其进行故障移除，但是过程是民主的。

因为不同的 sentinel 可能载入不同的配置（如：`down-after-milliseconds` 投票下线 master 的 sentinel 数量……），所以决定下线 master 的操作需要根据多个 sentinel 的意见来决定。

当一个 sentinel 检测到 master 主观下线之后，会判断其是否真的下线了，通过向同样监视该 master 的其他 sentinel 发送：

```shell
SENTINEL is-master-down-by-addr <ip> <port> <current_epoch> <runid>
```

进行询问，当从其他 sentinel **接收到足量的已下线判断**后，才会对 master 进行故障转移操作。

### 选举领头 sentinel

当一个 master 被判断为客观下线之后，监视该 master 的各个 sentinel 会进行商议，选举出一个领头 sentinel ，对下线 master 进行故障转移操作。（Sentinel 系统选举领头 Sentinel 的方法是对 Raft 算法的领头选举方法的实现。）

选举领头 sentinel 的规则和方法：

1. 所有的 sentinel 都有被选为领头的资格。
2. 每次选举之后，不论是否成功，所有 sentinel 的配置纪元（configuration epoch）都会自增一次。
3. 在一个配置纪元内，所有 sentinel 都有一次将某个 sentinel设置为局部领头的机会，并且局部领头一旦设置，在该纪元内则不能再更改。
4. 每个发现 master 进入客观下线的 sentinel 都会要求其他 sentinel 将自己设置为局部领头。
5. 当一个 sentinelA 向其他 sentinel 发送 `sentinel is-master-down-by-addr` ，其中的 runid 参数是自己的 runid 时，则表示 sentinelA 要求其他 sentinel 将自己设置为局部领头。
6. 先到先得原则。最先向选民 sentinel 发送设置要求（拉票）的 sentinel 将成为该选民 sentinel 的**意向局部领头**（投票），后来的其他 sentinel 的设置要求都会被拒绝。
7. 局部领头当选的要求是有半数以上的选票。
8. 每个配置纪元只会产生一个局部领头。
9. 若在给定的时限内没有一个 sentinel 当选，则会在一定时间内再次进行选举，直到选出为止。

### 故障转移

该操作包含三个动作：

1. 在已下线  master 的所有 salve 中选择一个作为新的 master
2. 让其他 salve 改为复制新的 master 
3. 将已下线的 master 作为新的 master 的 salve，当旧的 master再次上线时，它将作为新 master 的 salve。

#### 选新的 master

选新的 master 的标准是该 salve **状态良好、数据完整**。然后向其发送 `salveof no one` ，将其转换为 master。

**挑选规则：**

1. 去除已经下线或断线的 slave，保证剩余 salve 都是正常在线的
2. 去除最近五秒都没有回复领头 sentinel 的 `INFO` 命令的 salve，保证剩余 salve 都是最近成功通信过的
3. 去除所有与已下线 master 断开连接超过一定时间（`down-after-milliseconds * 10`）的 salve，保证剩下的 salve 保存的数据都是比较新的。

**选取方法**：之后，领头 sentinel 根据 slave 的优先级进行排序，选取优先级最高者。若有多个优先级同的 salve，按照其复制偏移量进行排序，选取复制偏移量最大的 salve ，以保证其数据是最新的。若复制偏移量仍然相同，则根据其 ID 进行排序，选取 ID 最小的 salve（**此时选谁都一样了**）。

#### 修改 salve 的复制目标

新的 master 选举出来之后，就需要让其余 salve 都改为复制新的 master，这一步通过向其发送 `SLAVEOF` 命令完成。

#### 将旧 master 变为新 master 的 salve

此时的 `SLAVEOF` 命令会等旧 master 上线时再发送。

## 集群

Redis 集群是 Redis 提供的分布式数据库方案，**集群通过分片（sharding）来进行数据共享**，并提供复制和故障转移功能。

### 节点

刚开始时，每个节点相互独立，都属于只包含自己的集群中，要组建可用的集群，需要将各节点连接起来。

```shell
# 连接节点
cluster meet <ip> <port>
# 列出集群中的节点。最开始时，只会列出自己
cluster nodes
```

启用集群，需要开启节点的集群模式。在配置文件中设置 `cluster-enabled yes` 的节点才能作为集群节点。在本机启用集群时，要注意多个节点的端口等差异化的数据、文件需要不同。

### 槽指派

集群通过分片方式保存数据库，将数据库分为 16384 (2^14) 个槽（slot），每个键都属于这其中的一个，被集群中的多个节点分别处理。一个节点可处理 [0,16384] 个槽，**若所有槽都有节点在处理，则集群处于上线状态；相反，若有任何一个槽没有得到处理，则集群处于下线状态。**所以，就算使用 `cluster meet` 构建好集群，集群依然处于下线状态。需要进行**槽指派**，使所有槽都有节点处理。通过 `cluster addsolts <slot> [...slots]` 可以为节点指派一个或多个槽。

[为什么是 16384 个槽? Redis 作者的回答]( https://github.com/antirez/redis/issues/2576 )

### 在集群中执行命令

因为多个节点分管不同的槽，当命令落到其他节点的槽上时，当前节点不会执行命令，而是会通过返回 `MOVED` 命令重定向到指定节点进行执行。重定向的过程在集群模式下会自动进行。另外，通过 `cluster keyslot "<key>"` 可以查看键所属的槽。

### 重新分片

重新分片可以**将已经指派给某节点的任意数量的槽改派给其他节点**，并且槽所属的键值对也会被迁移到新节点。

重新分片过程，集群不需要下线。

重新分片操作使用 redis-trib 完成。

### ASK 错误

这个错误是指，**在重新分片的过程中，属于一个槽的一部分键值对在旧节点，另一部分在新节点**，当客户端请求执行与键有关的命令时，若对应键存在当前节点，则直接执行；若**不存在，则返回 ASK 错误**，指引客户端到正在导入槽的目标节点。

### 复制与故障转移

集群的故障转移和 sentinel 机制类似。区别在于，sentinel 是内部选举，选民和候选都是 sentinel，选出新的 sentinel 执行故障转移操作；cluster 是外部选举，候选是下线节点 master 的 slaves ，而选民则是cluster 中的其他节点，选出新的 master 继续服务。

```shell
# 向一个节点发送命令,可以让接收命令的节点成为 node_id 所指定节点的 slave，并开始对 master 进行复制
cluster replicate <node_id>
```

#### 故障检测

集群中的节点会定期的向集群中的其他节点发送消息，以确认对方是否在线，若在规定时间内对方没有回复，则主观的认为对方下线，也称为疑似下线（probable fail）。同样的，若超过半数以上的主节点认为某个主节点疑似下线，则该主节点被标记为下线，将该主节点标记为下线的节点会发送广播，收到广播的节点会立刻将该节点也标记为下线。

#### 故障转移

上文提及过，集群的故障转移和 sentinel 机制类似，也是主从结构中的 master 下线，则选出一个 slave 充当新的 master 继续服务。大致流程为：

1. 复制下线主节点的所有从节点里面，会有一个从节点被选中。
2. 被选中的从节点会执行`SLAVEOF no one`命令，成为新的主节点。
3. 新的主节点会撤销所有对已下线主节点的槽指派，并将这些槽全部指派给自己。
4. 新的主节点向集群广播一条 PONG 消息，这条 PONG 消息可以让集群中的其他节点立即知道这个节点已经由从节点变成了主节点，并且这个主节点已经接管了原本由已下线节点负责处理的槽。
5. 新的主节点开始接收和自己负责处理的槽有关的命令请求，故障转移完成。

# 四、独立功能的实现

## 发布与订阅

发布订阅模式的实现参看 [pm-mq]( https://github.com/jianxinliu/pm/tree/master/src/main/java/com/minister/pm/mq)

## 事务

Redis 通过``MULTI`、`EXEC`、`WATCH`等命令来实现事务（transaction）功能。事务提供了一种将**多个命令请求打包**（命令入队），然后**一次性、按顺序地执行多个命令**的机制，并且在事务执行期间，服务器不会中断事务而改去执行其他客户端的命令请求，它会将事务中的所有命令都执行完毕，然后才去处理其他客户端的命令请求（**因为 redis 是单线程的，以此来保证事务的隔离性**）。

同时，通过 `watch` 命令实现乐观锁，监听键的修改，若键被修改，则拒绝执行事务，以此保证原子性。

redis 的事务同样具有 ACID 特性。但 redis 事务不支持回滚。

> - Redis commands can fail only if called with a wrong syntax (and the problem is not detectable during the command queueing), or against keys holding the wrong data type: this means that in practical terms a failing command is the result of a programming errors, and a kind of error that is very likely to be detected during development, and not in production.
> 
> - Redis is internally simplified and faster because it does not need the ability to roll back.
>   
>     ​                                                                                                                                                                             https://redis.io/topics/transactions 

## Lua脚本

## 慢查询日志

执行时间超过 `slowlog-log-slower-than` 设定的值的命令会被记录到慢查询日志中，慢查询日志需要是一个可指定长度（`slowlog-max-len`）的队列，若队满，最先入队的日志会被新入队的日志挤出。

`slowlog get` 获取日志，`slowlog len` 查询日志数量，`slowlog reset` 清楚日志。

## 监视器

客户端执行 `monitor` 命令可将自身变为监视器，实时接收并打印出服务器当前处理的命令请求的相关信息。