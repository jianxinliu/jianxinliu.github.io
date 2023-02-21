# JVM 相关经验

## 参数释义与示例

JVM 参数模式

>   -XX:[+/-]<option\>=<value(string/number)\>

1.   `+` 表示开启该 option。如：`-XX:+UseG1GC` 表示使用 G1 垃圾收集器
2.   `-` 表示关闭该 option。
3.   value string。如： `-XX:HeapDumpPath="/path/to/file"`
4.   value number, 可带单位。如：`-XX:MaxPermSize=64m`

参考：https://blog.csdn.net/Dax1n/article/details/77163540

JVM 空间大小配置参数

|         参数         | 含义                | 示例值 |
| :------------------: | ------------------- | :----: |
|         -Xms         | 初始堆大小          | 1536M  |
|         -Xmx         | 最大堆空间          | 1536M  |
|         -Xmn         | 年轻代大小          |  512M  |
|         -Xss         | 每个线程的堆栈大小  |  256k  |
|  -XX:SurvivorRatio   | Survivor 区空间比例 |   6    |
|  -XX:MetaspaceSize   | 元空间大小          |  256M  |
| -XX:MaxMetaspaceSize | 最大元空间大小      |  256M  |

GC 参数

|       参数       | 含义 | 示例值 |
| :--------------: | ---- | :----: |
|   -XX:+UseG1GC   |      |        |
| -XX:+UseParNewGC |      |        |
|                  |      |        |
|                  |      |        |

GC log 配置

|                 参数                 | 含义             | 示例 |
| :----------------------------------: | ---------------- | :--: |
|      -Xloggc:<gc log file path>      |                  |      |
|         -XX:+PrintGCDetails          | 打印 GC 过程细节 |      |
|        -XX:+PrintGCDateStamps        |                  |      |
|        -XX:+PrintGCTimeStamps        | 打印 GC 停顿耗时 |      |
|           -XX:PrintGCCause           |                  |      |
|      -XX:+UseGCLogFileRotation       |                  |      |
| -XX:NumberOfGCLogFiles=\<file number\> |                  |      |
|    -XX:GCLogFileSize=\<file size\>M    |                  |      |

Dump 配置

|                   参数                   | 含义                       | 示例 |
| :--------------------------------------: | -------------------------- | :--: |
|     -XX:+HeapDumpOnOutOfMemoryError      | OOM 的时候输出当前堆栈信息 |      |
| -XX:HeapDumpPath=\<hprof file path\>.hprof | dump 堆栈的文件路径        |      |



## 调优经验




## 问题排查经验

排查 JVM 相关问题需要用到的命令

```sh
# 查看当前机器上所有的 Java 进程
jps [-l]

# 查看 Java 进程的 jvm 占用情况
# jstat 命令参考  https://docs.oracle.com/javase/7/docs/technotes/tools/share/jstat.html
jstat -gcutil <vmid> <print time interval in millsecond>
```



## 问题排查工具

官方提供的工具：https://docs.oracle.com/javase/8/docs/technotes/tools/

标准 jdk 工具：

-   基础工具：jar, java, javac, javadoc, javah, javap, jdb, jdeps
-   排查、监控、管理工具：jcmd, jconsole, jmc, jvisualvm

实验性 jdk 工具（未来版本可能会去除）

-   监控工具： jps, jstat, jstatd
-   排查工具：jinfo, jhat, jmap, jsadebugd, jstack
-   脚本工具：jrunscript

### jps

官方文档：https://docs.oracle.com/javase/7/docs/technotes/tools/share/jps.html

一般格式：

`jps [option]`

|   选项   | 含义                                                         |
| :------: | ------------------------------------------------------------ |
|    -q    | 控制不输出类名, jar 文件名以及传给 main 的参数，只输出 jvm id |
|    -m    | 输出传给 main 的参数                                         |
|    -l    | 输出类的全名，或者 jar 文件的路径名                          |
|    -v    | 输出传给 jvm 的参数                                          |
|    -V    | 输出通过 flags file 方式传给 jvm 的参数（-XX:Flags=<filename>） |
| -Joption | 通过该选项，传递运行参数给 jps 调用的 Java 程序              |

### jstat

官方文档： https://docs.oracle.com/javase/7/docs/technotes/tools/share/jstat.html

一般格式：

`jstat [一般选项 | 输出选项 vmid [输出时间间隔[s|ms] [输出条数]]]`

参数表：

|   参数   | 说明                                                        |                            可选值                            |
| :------: | ----------------------------------------------------------- | :----------------------------------------------------------: |
| 一般选项 | 如果指定了一般选项，则再不能指定其他输出选项和参数          |       -help: 帮助信息<br>-options: 展示支持的输出选项        |
| 输出选项 | 决定了 jstat 命令输出内容的格式。格式为表格                 | -t：在表格第一列展示一个时间列，表示该 jvm 启动到现在经过的秒数<br>-h: 设置输出多少行内容就输出一行表头(`-h 3` 表示每输出 3 行就追加一行表头)<br>-J: 给 Java application launcher 设置启动参数 |
|   vmid   | virtual machine id。指定本机 jvm 进程号。也可以是远程机上的 |                                                              |
| 时间间隔 | 间隔多久采集一次。默认单位是毫秒，也可以通过 s 设置为秒     |                                                              |
| 输出条数 | 总共采集多少次，默认是无限次                                |                                                              |
| 统计选项 | 包含在输出选项中，决定命令输出的内容                        | 可通过 `jstat -options` 展示所有支持的选项。具体内容以及表格字段解释参考：[官方文档对应章节](https://docs.oracle.com/javase/7/docs/technotes/tools/share/jstat.html#statoption) |
|          |                                                             |                                                              |

统计选项：

1.   -class. 展示类加载的相关信息
2.   -compiler. 展示 JIT 编译器相关信息
3.   -gc. 展示当前堆状况
4.   -gccapacity. 展示各个年代区及其空间信息
5.   -gccause. 展示 gc 概况及其上次或本次 gc 的原因
6.   -gcnew. 展示新生代的情况
7.   -gcnewcapacity. 展示新生代的大小和空间
8.   -gcold. 展示老年代情况
9.   -gcoldcapacity. 展示老年代大小
10.   -gcpermcapacity. 展示永久区情况
11.   -gcutil. 展示 gc 概况
12.   -printcompilation. Hotspot 编译的方法统计

### jinfo

查看 jvm 配置信息

一般格式： `jinfo <pid>`

### jmap

展示 jvm 堆内存信息

一般格式：`jmap [options] pid`

Options:

|                选项                 | 说明                                                         |
| :---------------------------------: | ------------------------------------------------------------ |
|               无选项                | 直接在控制台打印堆内存内对象信息                             |
| -dump:[live,]format=b,file=filename | 导出堆内存信息到二进制文件 filename.hprof。live 如果指定了，则再 dump 之前会进行一次 gc, 确保 dump 出来的只有还在存活的对象 |
|                -heap                | 打印 GC 所使用的堆内存信息。                                 |
|            -histo[:live]            | 打印堆内存对象容量和个数的直方图（histogram）,live 选项同上  |
|              -clstats               | 打印类加载器相关的统计信息                                   |

### jhat

分析JVM 堆内存。分析 hprof 文件，并启用一个 web 服务器来展示。可以在浏览器上查看堆内存信息。内置 OQL(Object query language) 来动态查询内存信息。

一般格式：`jhat [options] heap-dump-file`

options

| 选项  | 含义                      |
| :---: | ------------------------- |
| -port | 指定 web 服务器监听的端口 |
|       |                           |
|       |                           |
|       |                           |
|       |                           |

### jstack

打印指定线程的堆栈信息。

一般格式：`jstack [options] pid`

Options:

| 选项 | 含义                                          |
| :--: | --------------------------------------------- |
|  -F  | 当此命令无响应时，强制 dump 堆栈信息          |
|  -l  | long listing。多打印一些锁的相关信息          |
|  -m  | 混合模式，一般适用于 Java 和 C/C++ 混合的情况 |























