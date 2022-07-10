# Go Web 先修知识

## GOPATH

[GOROOT & GOPATH](https://my.oschina.net/achun/blog/134002)
go中是没有项目这个概念的，只有包。
用于存放使用 `go get` 下载的 Go 依赖，这个环境变量支持多个目录，使用与系统相关的分隔符分割。此时，`go get` 的内容默认放在第一个目录。

以上 `GOPATH` 目录约定有三个子目录：

1. src.存放源码(.go,.c,.h,.s)
2. pkg.存放编译后生成的文件（.a） (go build)
3. bin.存放编译后生成的可执行文件 (go install)

新建代码包一般都是在 `src` 下新建文件夹，文件夹的名称一般是代码包的名称，允许多级目录。如 `beego` 的包名是 `github.com/astaxie`。`src`即是包管理的根目录。`import github.com/astaxie` 会先到 GOPATH 下查找依赖，若没有，则去 GOROOT 查找。

go 命令

### go get

`go get` 命令获取远程包。支持多个开源社区，如 github，googlecode,bitbucket……不同的平台需要不同的源码控制工具。github 采用 git。本质上可以理解为通过 git `clone` 代码到本地 `src` 目录，然后执行 `go install` 到 `bin` 目录。

### go build , go install & go run

go build 用于测试编译，编译过程中会发生包的链接，即编译与之相关联的包。命令后可接文件名，用于编译指定文件。不接文件，则遍历当前路径下所有文件。更多：`go help build` & `go help install`。
go install 和 go build 类似，只是前者或将生成的结果文件移到 pkg 或 bin 中。

go run 用于编译并运行 go 程序。

### go clean

移除当前源码包中编译生成的文件。一般需要提交代码时使用 go clean 清理和源码管理无关的编译生成文件。

### go fmt

go 风格格式化文件。`gofmt -w src` 格式化整个项目（或某个文件），并将格式化结果写入文件（`-w`）。

### go test

自动读取源码目录下名为 `*_test.go` 的文件，生成并运行测试用的可执行文件。`go help testflag`

## go module

 https://zhuanlan.zhihu.com/p/60703832 

最新的包管理工具。v1.13 默认开启。 可以通过 `go env` 命令返回值的 `GOMOD` 字段是否为空来判断是否已经开启了 gomod，如果没有开启，可以通过设置环境变量 `export GO111MODULE=on` 开启。 

### 开始使用

初始化，生成 go .mod 文件，类似 npm 的 package.json 文件

```sh
// 项目根目录
go mod init
// 生成 go.mod 文件（也可用于从老项目中迁移到 go module），内容如下：
// module jianxin/gomod (当前模块名称)
//
// go 1.13 （版本）
```

使用 go module 之后，在运行 `go run\build\test`等命令时，若缺少依赖，则会自动下载。之后会生成 go.sum 文件，作用类似 npm 的 package-lock.json 文件。

 gomod 不会在 `$GOPATH/src` 目录下保存下载包的源码，而是把源码和链接库保存在 `$GOPATH/pkg/mod` 目录下。 

### 包管理命令

#### 安装依赖

依然使用 go get ，但是可以指定版本，如：`go get foo@v1.2.2`，指定 git 的分支，tag ，hash ，如：`go get foo@master or foo@tag or foo@e3702bed2`。 gomod 除了遵循语义化版本原则外，还遵循最小版本选择原则。

 如果下载所有依赖可以使用 `go mod download` 命令。 

#### 升降级依赖

```sh
// 查看所有已升级依赖版本
go list -u -m all
// 升级次级或补丁版本号
go get -u <pkg>
// 仅升级补丁版本号
go get -u=patch <pkg>
// 升降级版本号，可以使用比较运算符控制
go get foo@'<v1.6.2'
```

#### 移除依赖

```sh
// 移出所有代码中不需要的包
go mod tidy
// 修改 go.mod 配置文件的内容
go mod edit --droprequire=golang.org/x/crypto
```

#### 查看 & 格式化依赖

```sh
// 直接查看 go.mod 或者
go list -m [-json] all

// 格式化 go.mod 文件
go mod edit -fmt
```

## go 基础

### go 关键字

go 一共有 25 个关键字。列举特有的关键字：

```go
func select defer go map chan fallthrough
```

### go 程序

go 程序是通过 package 来组织的。一个可独立运行的 go 程序，一定包含一个 main 包，main 包中一定包含一个入口函数 main，main 函数无参数无返回。即 `main.main()` 函数是每一个可独立运行程序的入口点。

注意点：main 包并不一定是指一个名为 main 的目录，而是在 程序的开头声明 `package main` 即可。

### go 语法

参考示例程序。syntaxExamples

# go Web

go 访问数据库

### go database/sql 接口

go 定义了`database/sql`接口而并未实现。

#### sql.Register

用于注册数据库驱动。在 init 函数中调用驱动，进行注册。

```go
var d = Driver{proto:"tcp",raddr:"127.0.0.1:3306"}
func init(){
    Register("SET NAMES utf8")
    sql.Register("mymysql",&d)
}
```

#### driver.Driver

是一个数据库驱动的接口，其中有一个方法：`Open(name string) (Conn,error)`。返回的 Conn 只能用于一次 goroutine 的操作。

#### driver.Conn

是一个数据库连接的接口，这个 Conn 同样只能用在一个 goroutine 中。接口中定义了一系列方法：

1. `Prepare(query string) (Stmt,error)`。返回与当前连接相关的执行 SQL 语句的状态，使用 statement 进行查询、删除等操作。
2. `Close() error`。关闭连接，释放资源。内部有连接池的实现。
3. `Begin() (Tx,error)`。开启一个事务。

#### driver.Stmt

是一种准备好的状态，和 Conn 相关联，也只能应用在一个 goroutine 中。一系列方法：

1. `Close() error`.关闭当前的连接状态。
2. `NumInput() int`.返回当前预留参数的个数。
3. `Exec(args []Value) (Result, error)`.执行 `Prepare` 准备好的 SQL，传入参数执行 update\insert 等操作会返回 Result 数据。
4. `Query(args []Value) (Rows,error)`.执行 `Prepare` 准备好的 SQL，传入参数执行 select 操作返回 Rows 结果集。

#### driver.Tx

事务处理，包含提交和回滚操作。

1. `Commit() error`.
2. `Rollback() error`.

#### driver.Execer

这是一个 Conn 可选择实现的接口，如果没有定义，调用 DB.Exec 时，会首先调用 Prepare 返回的 Stmt，然后执行 Stmt 的 Exec。有方法 `Exec(query string,args []Value) (Result,error)`。

#### driver.Result

执行 update、insert 时返回的结果，有方法：

1. `LastInsertId() (int64,error)`。执行插入时得到的自增 ID
2. `RowsAffected() (int64,error)`。返回 query 操作影响的条数。

#### driver.Rows

查询时返回的结果集

1. `Columns() []string`。返回所查询的字段
2. `Close() error`。关闭 Rows 迭代器
3. `Next(dest []Value) error`。返回下一条数据，数据值赋值给 dest。

#### driver.Value

是空接口，可容纳任何数据，要么是 nil ，要么是以下任意一种：`int64,float64,bool,[]byte,string,time.Time`。

### go 使用 mysql

建议驱动 https://github.com/Go-SQL-Driver/MySQL；

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main(){
    // 
    db,err := sql.Open("mysql","uname:pwd@/dbname")

    // insert & update & delete
    stmt,err:= db.Prepare("insert into user(name) value(?)")
    stmt.Exec("jack")
    // stme.LastInsertId(),stmt.RowsAffected()

    // query
    rows,err := db.Query(<query sql>)
    for rows.Next() {
        var name string
        rows.Scan(&name)
    }
    db.Close()
}
```

### go 使用 redis

驱动 https://github.com/hoisie/redis.go

```go
import "github.com/astaxie/goredis"

func main(){
    var client goredis.Client

    // string
    client.Set("a",[]byte("hello"))
    val,_ := client.Get("a")
    client.Del("a")

    // list
    vals := []string{"a","b","c"}
    for _,v := range vals {
        client.Rpush("1",[]byte(v))
    }
    dbvals,_ := client.Lrange("1",0,4)
    for i , v := range dbvals {
        //
    }
}
```

## Session & Cookie

```go
http.SetCookie(w ResponseWriter, cookie *Cookie)

r.Cookie(cname) // or r.Cookies()
```

go 未实现 Session，需手动写。

## 文本处理类

### json

```go
import "encoding.json"

// 只解析导出字段
type Student struct {
    Name string `json:"name"`
    Age string `json:age`
}

type Students struct {
    Stus []Student `json:stus`
}

func main()  {
    var stus Students
    str := `{"stus":[
        {"name":"jianxin","age":"14"},
        {"name":"jianxinliu","age":"24"}
        ]}`
    json.Unmarshal([]byte(str),&stus)
    fmt.Println(stus)

    // stus.Stus = append(stus.Stus, Student{Name:"jack",Age:35})

    b ,err := json.Marshal(stus)
    fmt.Println(string(b))
}
```

### 文件操作

#### 目录操作

1. `os.Mkdir(name string,prem FileMode) error`。创建单一目录，并设置权限，形式为r=4,w=2,x=1
2. `os.MkdirAll(path string,perm FileMode) error`。按路径创建目录，可嵌套
3. `os.Remove(name string) error`。删除目录，目录必须为空
4. `os.RemoveAll(path string) error`。删除多级子目录

#### 文件操作

1. `os.Create(name string) (file *File,err Error)`。创建默认权限为 0666 的文件。
2. `os.NewFile(fd uintptr,name string) *File`。根据文件描述符创建文件。
3. `os.Open(name string) (file *File,err Error)`。以只读的方式打开文件，内部调用 OpenFile。
4. `os.OpenFile(name string,flat int,perm uint32) (file *File,err Error)`。支持选择打开文件的方式，用 flag 设置。
5. `func (file *File) Write(b []byte) (n int,err Error)`。写入 byte 信息到文件。
6. `func (file *File) WriteAt(b []byte,off int64) (n int,err Error)`。在指定位置开始写入 byte 类型的信息。
7. `func (file *File) WriteString(s string) (ret int,err Error)`
8. `func (file *File) Read(b []byte) (n int,err Error)`。读数据到 b 。
9. `func (file *File) ReadAt(b []byte,off int64) (n int,err Error)`。从 off 开始读取数据到 b 。
10. `Remove(name string) Error`。和删除文件夹是同一个函数。

### 字符串处理

来自 strings 包。

1. `Contains(s,substr string) bool`。Contains("abc","c") true
2. `Join(a []string,sep string) string`。Join(["a","b"],",") "a,b"
3. `Index(s,sep string) int`。like indexOf
4. `Repeat(s string,count int) string`。Repeat("a",3) "aaa"
5. `Replace(s,old,new string,n int) string`。Replace("aaa","a","b",2) "bba"。n 表示替换的次数，< 0 则全部替换
6. `Split(s,sep string) []string`。
7. `Trim(s string,cutset string) string`。去除指定的字符串
8. `Fields(s string) []string`。去除空格符，并按空格进行 split

来自 strconv 包

1. Append* 系列将其他类型转换wi字符串再添加到现有字符串。AppendInt,AppendBool,AppendQuote……
2. Format* 系列将其他类型转换为字符串。FormatInt,FormatUint,Itoa……
3. Parse* 系列将字符串转换为其他类型。ParseInt,ParseBool,ParseFloat……

## Web 服务

### Socket

### WebSocket

### REST

### RPC

## 安全与加密

## 错误处理，调试 & 测试

### 错误处理

Error 对象
用法：errors.New() 将字符串转换为符合error接口的对象。

```go
func Sqrt(f float64) (float64,error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // ...
}
```

实现 Error 方法即可自定义错误类型。

### GDB 调试

### 测试

测试用例统一在一个文件中，该文件的要求：

1. 文件名以 `_test.go` 结尾
2. import "testing"
3. 测试用例函数以 `Test` 开头，且后续命名注意区分，Testin 不符合，TestIn 才符合
4. 测试用例按函数定义顺序执行
5. 测试函数 `TestXXX(t *testing.T)` 是正确的函数声明
6. 函数中通过 `t.Error,Errorf,FailNow,Fatal，FatalIf` 说明测试不通过,调用 `Log`记录测试信息。

在文件夹下执行 `go test [-v]` 执行测试用例

### 压力测试

压力测试函数的格式为 `func BenchmarkXXX(b *testing.B)`.

1. `go test -test.bench="test_name_regex"` 执行压力测试
2. 压力测试用例中，循环体内使用 `testing.B.N` 作为循环次数
3. 文件名也是以 `_test.go` 结尾。

## 部署与维护

### 日志

[seelog](https://github.com/cihub/seelog)

# Go 并发编程

# 参考

《Go Web编程》谢孟军
《Go 并发编程实战》郝林
