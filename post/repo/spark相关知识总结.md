## 常用参数说明

|             参数             |                             说明                             | 默认值 |
| :--------------------------: | :----------------------------------------------------------: | :----: |
|       spark.cores.max        |               整个集群可以使用的最大 CPU 核数                |        |
|     spark.executor.cores     | 如果配置了，则物理机上有多少可用 CPU（N）,则执行器个数为 N/cores。同事，该参数可以理解成任务执行时线程池的大小 |        |
|    spark.executor.memory     | 执行器节点能够使用的最大内存，执行器下的每个任务线程，瓜分这部分内存 |   1g   |
|    spark.memory.fraction     | 执行器内，用于执行计算和存储的内存划分比例。用于指定执行内存的比例 |  0.6   |
| spark.memory.storageFraction |                 执行器内存里用于存储的百分比                 |  0.5   |
|                              |                                                              |        |
|                              |                                                              |        |
|                              |                                                              |        |
|                              |                                                              |        |



### 资源计算

|                    |                           含义说明                           |                           计算公式                           |                       备注                        |
| :----------------: | :----------------------------------------------------------: | :----------------------------------------------------------: | :-----------------------------------------------: |
|         M          |              执行器可用内存（execution memory）              | $executor.memory * memory.fraction * ( 1 - storageFraction)$ |                                                   |
|         N          |                  executor.cores(线程池大小)                  |                   $spark.executors.cores$                    |                                                   |
|         D          |                       数据量（如 20G）                       |                                                              |                                                   |
|         P          |                    并行度（task 的数量）                     |                                                              |                                                   |
|      资源匹配      | $D / P$ 是需求。$(\frac{1}{2}\frac{M}{N}, \frac{M}{N})$ 是系统能提供的内存资源。二者在同一个数量级，就基本能欧满足资源需求，以及较好的利用率 |       $D/P <==> (\frac{1}{2}\frac{M}{N}, \frac{M}{N})$       | 经验谈：$\frac{M}{N}$ 在 200Mb 比较适合大部分场景 |
| 每个任务的资源需求 |               相当于每个任务需要处理多大的数据               |                            $D/P$                             |                                                   |
|   系统分配的资源   |           在配置下，能够分给每个任务的内存资源范围           |           $(\frac{1}{2}\frac{M}{N}, \frac{M}{N})$            |                                                   |
|                    |                                                              |                                                              |                                                   |
|                    |                                                              |                                                              |                                                   |
|                    |                                                              |                                                              |                                                   |



## EDA spark 可能的优化点

1. 硬件相关参数调整：调整 driver 的内存，执行器的内存和 cores, 寻找合适的并行度
2. shuffle 相关参数调整： spark.sql.autoBroadcastJoinThreshold 决定何时方式 shuffle join 而改用 broad cast join, 当前是 10M, 可调节至 2G。scala 编程中采用广播变量
3. job 等待超时阈值调低（有些任务一直在 waitting）
4. 更改 spark 的默认序列化方式为 Kryo 
5. 适当降低 Executor 内存 RDD 存储占用的比率（默认 60%， shuffle 20%， APP 可用内存 20%）(目前应用并没有用到该功能，可以适当减少，给程序运行和计算让出空间)  
	spark.memory.fraction=0.8(0.6)
	spark.memory.storageFraction=0.2(0.5) 
6. spark.sql.codegen -> true, 编译 SQL 语句为二进制，在大型，重复使用 SQL 的情况下性能优异
7. spark.cleaner.periodicGC.interval 默认 30min, 可以调整为更小，更频繁的触发内存清理工作
8. spark.cleaner.referenceTracking.blocking 默认阻塞，可以关闭阻塞，开启并行清理

## 相关资料收集

https://blog.csdn.net/u010002184/article/details/111737069    spark 3.* 内存配置说明  统一内存分配 vs 静态内存分配

https://www.cnblogs.com/zhouyc/p/13562858.html

https://bbs.huaweicloud.com/blogs/184507

https://www.programmersought.com/article/64963553068/

[spark 官方优化思路](https://spark.apache.org/docs/latest/tuning.html)

[美团 spark 基本优化思路](https://tech.meituan.com/2016/04/29/spark-tuning-basic.html)

[long running spark application debugging](https://www.channable.com/tech/debugging-a-long-running-apache-spark-application-a-war-story)

## spark 笔记

逻辑分层
application
	jobs
		stages
			tasks
物理分层
driver
	workers
		executors
		
spark-shell attach to master : spark-shell --master spark://<host>:7077		
		

## 调优经验

> In the end, our cluster performance problems were solved by two simple configuration changes. 
> 
> First, we had to make sure that the garbage collector in the Spark driver program was triggered more often. 
> 
> his would make sure that the queue with cleanup tasks would get filled as soon as a cluster-wide resource like e.g. an RDD  was not needed any more and could be expediently removed from all workers. 
> 
> Second, we had to make sure that the cleanup thread was  running in non-blocking mode. Otherwise, it could simply not keep up with the number of cleanup tasks that we were generating.	
> **spark.cleaner.periodicGC.interval**
> 
> **spark.cleaner.referenceTracking.blocking=false**


## spark join strategies: 	

[spark  join 策略](https://towardsdatascience.com/strategies-of-spark-join-c0e7b4572bcf)

### Broadcasted hash join。

什么是 hash join。使用参与 join 的两表中较小的一个表，根据 join_key 构造 hash table， 然后循环另一个大表，去一一匹配。（这是适用于 = 的 join_key）
Broadcasted hash join, 将其中一个较小的 表（rdd） 复制一份，发送到各个 worker 节点，避免了 shuffle， worker 节点上的 task 共享该小表。 spark 只有在一方是小表的情况下才会选择该策略。

### shuffle hash join。

什么是 shuffle。由于 spark 应用是集群模式，会存在多个工作节点，数据都保存在不同的工作节点上，进行 join 时， 可能节点 A 需要存储在节点 B 上的某个数据
若进行频繁交换，则效率低下。此时可进行 hash shuffle, 将整体数据按照一定的 hash_key 进行分组，相同 hash_key 值的数据会被 shuffle 到同一个节点上，这样后续的操作就可以避免频繁的数据交换，从而提升效率。
但是当数据集巨大时，shuffle 产生的节点间数据传输量也是巨大的，也会成为性能瓶颈。 同时巨大的数据也会导致需要维护一个巨大的 hash table， 同样也是高内存消耗的。

### shuffle sort-merge join  https://www.hadoopinrealworld.com/how-does-shuffle-sort-merge-join-work-in-spark/

自 spark 2.3 后，是默认的 join 策略，使用与 join 双方都是大表的情况。该策略分三个阶段：
1. shuffle 根据 join-key 将两边的数据 shuffle 到各个节点（join-key 相同的值到同一个节点） 
2. sort  join 两边的数据集按照 join-key 进行排序
3. merge 因为经过 sort 节点，数据集是按 join-key 有序的, 所以在 merge 节点，join 操作就不需要遍历整个数据集去寻找符合的值。

### Cartesian Join

使用笛卡尔积计算两个关系的连接，也叫 Broadcast nested loop join，是一个嵌套循环结构，效率较差，是 spark 的 fallback 策略。

spark 如何选择 join 策略：
> If it is an ‘=’ join:
> Look at the join hints, in the following order:
>
> 	1. Broadcast Hint: Pick broadcast hash join if the join type is supported.
>	
> 	2. Sort merge hint: Pick sort-merge join if join keys are sortable.
>	
> 	3. shuffle hash hint: Pick shuffle hash join if the join type is supported.
>	
> 	4. shuffle replicate NL hint: pick cartesian product if join type is inner like.
>   If there is no hint or the hints are not applicable
>
> 	5. Pick broadcast hash join if one side is small enough to broadcast, and the join type is supported.
>	
> 	6. Pick shuffle hash join if one side is small enough to build the local hash map, and is much smaller than the other side, and spark.sql.join.preferSortMergeJoin is false.
>	
> 	7. Pick sort-merge join if join keys are sortable.
>	
> 	8. Pick cartesian product if join type is inner .
>	
> 	9. Pick broadcast nested loop join as the final solution. It may OOM but there is no other choice.
>


>   If it’s not ‘=’ join:
>   Look at the join hints, in the following order:
>
> 	10. broadcast hint: pick broadcast nested loop join.
>	
> 	11. shuffle replicate NL hint: pick cartesian product if join type is inner like.
>    If there is no hint or the hints are not applicable
>
> 	12. Pick broadcast nested loop join if one side is small enough to broadcast.
>	
> 	13. Pick cartesian product if join type is inner like.
>	
> 	14. Pick broadcast nested loop join as the final solution. It may OOM but we don’t have any other choice.

### Alluxio 设置文件 TTL
```java
fileSystem.createDirectory(
                    getAlluxioUrl(fileExportInfo.getFilePath()),
                    CreateDirectoryPOptions.newBuilder()
                    .setCommonOptions(FileSystemMasterCommonPOptions.newBuilder().setTtl(11L).build())
                    .setRecursive(true).build());
```

