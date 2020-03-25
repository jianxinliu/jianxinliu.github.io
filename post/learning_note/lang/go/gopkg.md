# Go pkg

## builtin

默认导入的包，包含原生类型、数据结构、内建函数。

- `func close(c chan<- Type)` 用于关闭 channel， 该 channel 必须为双向的或只发送的。 应当只由发送者执行，而不应由接收者执行 

- `func delete(m map[Type]Type1, key Type)` 用于在 map 中删除实例。

- `func len(v Type) int`用于返回 v 的长度。取决于不同的类型：

    - 数组：v 中元素的数量。
    - 数组指针：*v 中元素的数量（即使 v 为 nil）。
    - 切片或映射：v 中元素的数量；若 v 为 nil，len(v) 即为零。
    - 字符串：v 中字节的数量。
    - 信道：信道缓存中队列（未读取）元素的数量；若 v 为 nil，len(v) 即为零。

- `func cap(v Type) int` 返回 v 的容量。取决于不同的类型：

    - 数组：v 中元素的数量（与 len(v) 相同）。
    - 数组指针：*v 中元素的数量（与 len(v) 相同）。
    - slice：在重新切片时，切片能够达到的最大长度；若 v 为 nil，len(v) 即为零。
    - channel：按照元素的单元，相应信道缓存的容量；若 v 为 nil，len(v) 即为零。

- `func new(Type) *Type` 用于各种类型的内存分配。

- `func make(Type, size IntegerType) Type` 用于内建类型（map、slice 和channel）的内存分配。

- `func copy(dst, src []Type) int` 用于复制 slice

- `func append(slice []Type, elems ...Type) []Type` 用于追加 slice。

- panic 和 recover 用于异常处理机制

- print 和 println 是底层打印函数，可以在不引入 fmt 包的情况下使用， is useful for bootstrapping and debugging 。但并不保证会一直存在于 go 语言中。

- complex、real 和 imag 全部用于处理复数

## bufio

 bufio 包实现了带缓存的I/O操作。 它封装了一个 io.Reader 或者 io.Writer 对象，另外创建了一个对象 （Reader 或者 Writer），这个对象也实现了一个接口，并提供缓冲和文档读写的帮助。 

### Scanner

#### 用法

```go
// Scanner 最简单的用法。
// Scanner 对象可以从 Reader 中得到，如：scanner := bufio.NewScanner(strings.NewReader(“”))
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    fmt.Println(scanner.Text()) // 返回最后一次 Scan 得到的 token 字符串
}
// After scanner.Scan returns false, the Err method will return any error that
// occurred during scanning, except that if it was io.EOF, Err will return nil.
if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
}
```

#### 分割函数

1. [`func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)`](http://docscn.studygolang.com/pkg/bufio/#ScanBytes)
2. [`func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)`](http://docscn.studygolang.com/pkg/bufio/#ScanLines)
3. [`func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)`](http://docscn.studygolang.com/pkg/bufio/#ScanRunes)
4. [`func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)`](http://docscn.studygolang.com/pkg/bufio/#ScanWords)

用于 `Scanner` 的分割函数，按不同单位对 `data` 进行分割，都符合 `SplitFunc`。

`type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)`。 

可自定义分割函数，只要符合 `SplitFunc` 定义即可，调用 `scanner.Split(splitFunc)` 替换默认分割器为自定义分割器，也可以调用该函数显式指定分割器:`scanner.Split(bufio.ScanWords)`。

分割器是 `Scanner` 内部调用的函数：`advance, token, err := s.split(s.buf[s.start:s.end], s.err != nil)`。

### Reader

 Reader 实现了对一个 io.Reader 对象的缓冲读。 

1. [`func NewReader(rd io.Reader) *Reader`](http://docscn.studygolang.com/pkg/bufio/#NewReader)

2. [`func NewReaderSize(rd io.Reader, size int) *Reader`](http://docscn.studygolang.com/pkg/bufio/#NewReaderSize)

3. [`func (b *Reader) Buffered() int`](http://docscn.studygolang.com/pkg/bufio/#Reader.Buffered) 当前 buf 缓存的字节数 `return b.w - b.r`。

4. [`func (b *Reader) Discard(n int) (discarded int, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.Discard) 读取时跳过接下来的 n 个字节。

5. [`func (b *Reader) Peek(n int) ([]byte, error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.Peek) 取 n 个字节但不移动指针。

6. [`func (b *Reader) Read(p []byte) (n int, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.Read) 将数据读入 p ，若 `len(p) >= len(b.buf)` 则直接读入 p (``b.rd.Read(p)`) ，避免复制；否则读入 buf (`b.rd.Read(b.buf)`)。其中 `r,w` 指针分别表示当前**读取位**和**写入位**，故求 buf 大小的函数 `Buffered() int {return b.w - b.r}`。

7. [`func (b *Reader) ReadByte() (c byte, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.ReadByte)

8. [`func (b *Reader) ReadBytes(delim byte) (line []byte, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.ReadBytes)

9. [`func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.ReadLine)

10. [`func (b *Reader) ReadRune() (r rune, size int, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.ReadRune)

11. [`func (b *Reader) ReadSlice(delim byte) (line []byte, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.ReadSlice)

12. [`func (b *Reader) ReadString(delim byte) (line string, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.ReadString)

13. [`func (b *Reader) Reset(r io.Reader)`](http://docscn.studygolang.com/pkg/bufio/#Reader.Reset)

14. [`func (b *Reader) UnreadByte() error`](http://docscn.studygolang.com/pkg/bufio/#Reader.UnreadByte)

15. [`func (b *Reader) UnreadRune() error`](http://docscn.studygolang.com/pkg/bufio/#Reader.UnreadRune)

16. [`func (b *Reader) WriteTo(w io.Writer) (n int64, err error)`](http://docscn.studygolang.com/pkg/bufio/#Reader.WriteTo)

### Writer

 Writer实现了io.Writer对象的缓存。

#### 用法：

```go
w := bufio.NewWriter(os.Stdout)
fmt.Fprint(w, "Hello, ")
fmt.Fprint(w, "world!")
w.Flush() // Don't forget to flush!
```

1. [`func NewWriter(w io.Writer) *Writer`](http://docscn.studygolang.com/pkg/bufio/#NewWriter)
2. [`func NewWriterSize(w io.Writer, size int) *Writer`](http://docscn.studygolang.com/pkg/bufio/#NewWriterSize)
3. [`func (b *Writer) Available() int`](http://docscn.studygolang.com/pkg/bufio/#Writer.Available) 返回 buf 中未使用的字节数，`return len(b.buf) - b.n`。
4. [`func (b *Writer) Buffered() int`](http://docscn.studygolang.com/pkg/bufio/#Writer.Buffered) 返回 buf 中数据的个数, `return b.n`。
5. [`func (b *Writer) Flush() error`](http://docscn.studygolang.com/pkg/bufio/#Writer.Flush) 将 buf 中的数据写入到底层的 io.Writer。
6. [`func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)`](http://docscn.studygolang.com/pkg/bufio/#Writer.ReadFrom)
7. [`func (b *Writer) Reset(w io.Writer)`](http://docscn.studygolang.com/pkg/bufio/#Writer.Reset) 清除当前 buf 中的数据，存在的错误，重定向输出到 w。
8. [`func (b *Writer) Write(p []byte) (nn int, err error)`](http://docscn.studygolang.com/pkg/bufio/#Writer.Write)
9. [`func (b *Writer) WriteByte(c byte) error`](http://docscn.studygolang.com/pkg/bufio/#Writer.WriteByte)
10. [`func (b *Writer) WriteRune(r rune) (size int, err error)`](http://docscn.studygolang.com/pkg/bufio/#Writer.WriteRune)
11. [`func (b *Writer) WriteString(s string) (int, error)`](http://docscn.studygolang.com/pkg/bufio/#Writer.WriteString)

### ReaderWriter

实现了 io.ReaderWriter ，维护了指向 Reader 和 Writer 的指针

```go
type ReadWriter struct {
	*Reader
	*Writer
}
```

## bytes

和 `strings` 包类似，用于协助处理 byte 切片。

 http://docscn.studygolang.com/pkg/bytes/ 

## container

提供三种容器的实现：堆，双向链表和循环链表。

### heap

为实现了 heap 接口的类型提供操作，默认是一个小顶堆。堆可以用于实现优先级队列（不过要逆转排序）。

### list

实现了一个双向链表。

### ring



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