# Zookeeper

是什么：用于分布式系统的分布式协调服务。分布式应用可以在其基础上构建同步、配置管理、分组（groups）、命名空间（naming） 等高级服务（It exposes a simple set of primitives that distributed applications can build upon to implement higher level services for synchronization, configuration maintenance, and groups and naming.）。

设计目标：

简单：zk 允许分布式进程通过像文件系统一样的**共享的层级式命名空间**（shared hierarchical namespace）进行协作。命名空间由叫 znode  的数据注册器组成。虽然很像用于存储的文件和文件系统，但有所不同的是， znode 的数据都在内存中。

可复制： zk 