**Logback**

上手使用十分简单，使用已有配置，在配置文件中指定日志的配置文件



## 关键概念：Logger、Appenders and Layouts

此三个类用于控制不同类型不同级别的消息以什么格式输出到哪里。

### Logger

配合 Lombok 的 `@Slf4j` 注解，直接使用 log 变量即可打 log。该方法是 `Logger logger = LoggerFactory.getLogger(xxx.xxx.Class.class);` 的简化形式，后者可以获得和当前类全限定名同名的 `Logger`，这样做是为了获得唯一命名的 Logger（因为`getLogger` 方法传入相同的名称只会得到相同的实例，也可以方便地获知日志信息来源）。

### Appender

在 Logback 的语境里，日志的输出地叫做 Appender。Logback  支持将日志输出到很多不同的地方，如 Console、文件、远程服务器、各种数据库、消息队列等，并且支持设置多个输出地。

### Layout

Layout 用于设置输出的日志格式，需要关联到 Appender 上。

### Configuration

logback 支持使用 java 代码配置和 基于 xml 或 Groovy 的文件配置。

Logback 配置文件查找顺序(在 classpath 里找)，一个未找到则继续往下找：

1. `logback-test.xml`
2. `logback.groovy`
3. `logback.xml`
4. 使用 Logback 自己的配置，查找 `META-INF\services\ch.qos.logback.classic.spi.Configurator` 的实现类
5. 使用默认的 `BasicConfigurator` ，将日志输出到控制台

所以不配置 logback 也是可以直接上手使用的。