# BIO 

传统的IO模式，阻塞式IO

参考：[SnailClimb-JavaGuide-Java IO & NIO](https://github.com/Snailclimb/JavaGuide/blob/master/Java%E7%9B%B8%E5%85%B3/Java%20IO%E4%B8%8ENIO.md)

**按操作方式分类结构图：**

[Java IO，硬骨头也能变软](https://mp.weixin.qq.com/s?__biz=MzU4NDQ4MzU5OA==&mid=2247483981&idx=1&sn=6e5c682d76972c8d2cf271a85dcf09e2&chksm=fd98542ccaefdd3a70428e9549bc33e8165836855edaa748928d16c1ebde9648579d3acaac10#rd)

![按操作方式分类结构图](..\assets\001.jpg)

**按操作对象分类结构图**

![按操作对象分类结构图](..\assets\002.jpg)

## IO 流的分类

- 按流的方向分：输入流、输出流
- 按流的操作单元分：字节流、字符流
- 按流的角色分：节点流、处理流

## IO的基类

- InputStream\Reader:所有输入流的基类，前者是字节流，后者是字符流
- OutputStream\Writer:所有输出流的基类，前者是字节流，后者是字符流

