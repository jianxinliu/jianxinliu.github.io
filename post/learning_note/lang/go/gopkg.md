# Go pkg

## default

默认导入的包，包含原生类型、数据结构、内建函数。

- close 用于channel 通讯。使用它来关闭 channel
- delete 用于在 map 中删除实例。
- len 和 cap 可用于不同的类型， len 用于返回字符串、slice 和数组的长度.
- new 用于各种类型的内存分配。
- make 用于内建类型（map、slice 和channel）的内存分配。
- copy 用于复制 slice
- append 用于追加 slice。
- panic 和recover 用于异常处理机制
- print 和 println 是底层打印函数，可以在不引入 fmt 包的情况下使用。
- complex、real 和 imag 全部用于处理复数

## bufio

 bufio 包实现了带缓存的I/O操作。 它封装了一个 io.Reader 或者 io.Writer 对象，另外创建了一个对象 （Reader 或者 Writer），这个对象也实现了一个接口，并提供缓冲和文档读写的帮助。 

## bytes



## container



## crypto



## database



## encoding



## errors



## fmt



## io



## math



## net



## os



## reflect



## sort、strings 、strconv、text



## sync



## time