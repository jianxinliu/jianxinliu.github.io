**Logback**

上手使用十分简单，使用已有配置，在配置文件中指定日志的配置文件



## 关键概念：Logger、Appenders and Layouts

此三个类用于控制不同类型不同级别的消息以什么格式输出到哪里。

### Logger

配合 Lombok 的 `@Slf4j` 注解，直接使用 log 变量即可打 log。该方法是 `Logger logger = LoggerFactory.getLogger(xxx.xxx.Class.class);` 的简化形式，后者可以获得和当前类全限定名同名的 `Logger`，这样做是为了获得唯一命名的 Logger（因为`getLogger` 方法传入相同的名称只会得到相同的实例，也可以方便地获知日志信息来源）。

### Appender

在 Logback 的语境里，日志的输出地叫做 Appender。Logback  支持将日志输出到很多不同的地方，如 Console、文件、远程服务器、各种数据库、消息队列等，并且支持设置多个输出地。

### Layout

Layout 用于设置输出的日志格式，需要关联到 Appender 上。默认是 `%m%n` 表示输出程序打的日志并且换行。

layout pattern 参考([log4j v1.2 PatternLayout doc](https://logging.apache.org/log4j/1.2/apidocs/org/apache/log4j/PatternLayout.html))

[另一个带示例的参考](https://blog.csdn.net/guoquanyou/article/details/5689652)

| Conversion Character | Effect                                                       |
| -------------------- | ------------------------------------------------------------ |
| **c**                | Used to output the category of the logging event. The category conversion specifier can be optionally followed by *precision specifier*, that is a decimal constant in brackets.If a precision specifier is given, then only the corresponding number of right most components of the category name will be printed. By default the category name is printed in full.For example, for the category name "a.b.c" the pattern **%c{2}** will output "b.c". |
| **C**                | Used to output the fully qualified class name of the caller issuing the logging request. This conversion specifier can be optionally followed by *precision specifier*, that is a decimal constant in brackets.If a precision specifier is given, then only the corresponding number of right most components of the class name will be printed. By default the class name is output in fully qualified form.For example, for the class name "org.apache.xyz.SomeClass", the pattern **%C{1}** will output "SomeClass".**WARNING** Generating the caller class information is slow. Thus, use should be avoided unless execution speed is not an issue. |
| **d**                | Used to output the date of the logging event. The date conversion specifier may be followed by a *date format specifier* enclosed between braces. For example, **%d{HH:mm:ss,SSS}** or **%d{dd MMM yyyy HH:mm:ss,SSS}**. If no date format specifier is given then ISO8601 format is assumed.The date format specifier admits the same syntax as the time pattern string of the [`SimpleDateFormat`](http://java.sun.com/j2se/1.4.2/docs/api/java/text/SimpleDateFormat.html?is-external=true). Although part of the standard JDK, the performance of `SimpleDateFormat` is quite poor.For better results it is recommended to use the log4j date formatters. These can be specified using one of the strings "ABSOLUTE", "DATE" and "ISO8601" for specifying [`AbsoluteTimeDateFormat`](https://logging.apache.org/log4j/1.2/apidocs/org/apache/log4j/helpers/AbsoluteTimeDateFormat.html), [`DateTimeDateFormat`](https://logging.apache.org/log4j/1.2/apidocs/org/apache/log4j/helpers/DateTimeDateFormat.html) and respectively [`ISO8601DateFormat`](https://logging.apache.org/log4j/1.2/apidocs/org/apache/log4j/helpers/ISO8601DateFormat.html). For example, **%d{ISO8601}** or **%d{ABSOLUTE}**.These dedicated date formatters perform significantly better than [`SimpleDateFormat`](http://java.sun.com/j2se/1.4.2/docs/api/java/text/SimpleDateFormat.html?is-external=true). |
| **F**                | Used to output the file name where the logging request was issued.**WARNING** Generating caller location information is extremely slow and should be avoided unless execution speed is not an issue. |
| **l**                | Used to output location information of the caller which generated the logging event.The location information depends on the JVM implementation but usually consists of the fully qualified name of the calling method followed by the callers source the file name and line number between parentheses.The location information can be very useful. However, its generation is *extremely* slow and should be avoided unless execution speed is not an issue. |
| **L**                | Used to output the line number from where the logging request was issued.**WARNING** Generating caller location information is extremely slow and should be avoided unless execution speed is not an issue. |
| **m**                | Used to output the application supplied message associated with the logging event. |
| **M**                | Used to output the method name where the logging request was issued.**WARNING** Generating caller location information is extremely slow and should be avoided unless execution speed is not an issue. |
| **n**                | Outputs the platform dependent line separator character or characters.This conversion character offers practically the same performance as using non-portable line separator strings such as "\n", or "\r\n". Thus, it is the preferred way of specifying a line separator. |
| **p**                | Used to output the priority of the logging event.            |
| **r**                | Used to output the number of milliseconds elapsed from the construction of the layout until the creation of the logging event. |
| **t**                | Used to output the name of the thread that generated the logging event. |
| **x**                | Used to output the NDC (nested diagnostic context) associated with the thread that generated the logging event. |
| **X**                | Used to output the MDC (mapped diagnostic context) associated with the thread that generated the logging event. The **X** conversion character *must* be followed by the key for the map placed between braces, as in **%X{clientNumber}** where `clientNumber` is the key. The value in the MDC corresponding to the key will be output.See [`MDC`](https://logging.apache.org/log4j/1.2/apidocs/org/apache/log4j/MDC.html) class for more details. |
| **%**                | The sequence %% outputs a single percent sign.               |

### Configuration

logback 支持使用 java 代码配置和 基于 xml 或 Groovy 的文件配置。

Logback 配置文件查找顺序(在 classpath 里找)，一个未找到则继续往下找：

1. `logback-test.xml`
2. `logback.groovy`
3. `logback.xml`
4. 使用 Logback 自己的配置，查找 `META-INF\services\ch.qos.logback.classic.spi.Configurator` 的实现类
5. 使用默认的 `BasicConfigurator` ，将日志输出到控制台

所以不配置 logback 也是可以直接上手使用的。

### 示例

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Configuration status="debug" name="eda" dest="/aplog/log/eda/log4j2.log">
    <Properties>
        <Property name="applicationName">${spring:spring.application.name}</Property>
        <Property name="LOG_HOME">${spring:matrix.log.dir}/${applicationName}
        </Property>
        <Property name="defaultPatternLayout" >
            %d{yyyy-MM-dd HH:mm:ss.SSS} %-5level [%t] [%X{evtUsr}][%X{traceId}] %logger{36} - %msg%n : %m%n${LOG_EXCEPTION_CONVERSION_WORD:-%wEx}
        </Property>

        <!-- for monitor -->
        <Property name="MONITOR_LOG_FILE_PATH">${spring:matrix.monitor.file.path}</Property>
        <Property name="MONITOR_LOG_HOME">${MONITOR_LOG_FILE_PATH}/${applicationName}</Property>
        <Property name="monitorPatternLayout">
            %m%n
        </Property>
    </Properties>
    <Appenders>
        <Console name="Console" target="SYSTEM_OUT">
            <PatternLayout pattern="%d{yyyy-MM-dd HH:mm:ss.SSS} %-5level [%t] [%X{evtUsr}][%X{traceId}] %logger{36} - %msg%n"/>
        </Console>
        <!-- <JsonLayout compact="true" locationInfo="true" complete="false"
        eventEol="true"></JsonLayout>-->
        <RollingFile name="infoLog" fileName="${LOG_HOME}/info.log"
                     filePattern="${LOG_HOME}/info.log.%d{yyyy-MM-dd}">
            <ThresholdFilter level="INFO" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${defaultPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
            <DefaultRolloverStrategy max="30"/>
        </RollingFile>
        <RollingFile name="errorLog" fileName="${LOG_HOME}/error.log"
                     filePattern="${LOG_HOME}/error.log.%d{yyyy-MM-dd}">
            <ThresholdFilter level="ERROR" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${defaultPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
            <DefaultRolloverStrategy max="30"/>
        </RollingFile>
        <RollingFile name="warnLog" fileName="${LOG_HOME}/warn.log"
                     filePattern="${LOG_HOME}/warn.log.%d{yyyy-MM-dd}">
            <ThresholdFilter level="WARN" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${defaultPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
            <DefaultRolloverStrategy max="30"/>
        </RollingFile>
        <RollingFile name="sqlLog" fileName="${LOG_HOME}/sql.log"
                     filePattern="${LOG_HOME}/sql.log.%d{yyyy-MM-dd}">
            <ThresholdFilter level="INFO" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${defaultPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
            <DefaultRolloverStrategy max="30"/>
        </RollingFile>
        <RollingFile name="traceIdInfo" fileName="${LOG_HOME}/traceIdInfo.log"
                     filePattern="${LOG_HOME}/traceIdInfo.log.%d{yyyy-MM-dd}">
            <ThresholdFilter level="INFO" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${defaultPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
        </RollingFile>

        <!-- for monitor -->
        <RollingFile name="DEBUG_LOG_MONITOR" fileName="${MONITOR_LOG_HOME}/debug.log"
                     filePattern="${MONITOR_LOG_HOME}/debug.log.%d{yyyy-MM-dd}.%i.log">
            <ThresholdFilter level="DEBUG" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${monitorPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
            <DefaultRolloverStrategy max="30"/>
        </RollingFile>

        <RollingFile name="INFO_LOG_MONITOR" fileName="${MONITOR_LOG_HOME}/info.log"
                     filePattern="${MONITOR_LOG_HOME}/info.log.%d{yyyy-MM-dd}.%i.log">
            <ThresholdFilter level="INFO" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${monitorPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
            <DefaultRolloverStrategy max="30"/>
        </RollingFile>

        <RollingFile name="WARN_LOG_MONITOR" fileName="${MONITOR_LOG_HOME}/warn.log"
                     filePattern="${MONITOR_LOG_HOME}/warn.log.%d{yyyy-MM-dd}.%i.log">
            <ThresholdFilter level="WARN" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${monitorPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
            <DefaultRolloverStrategy max="30"/>
        </RollingFile>

        <RollingFile name="ERROR_LOG_MONITOR" fileName="${MONITOR_LOG_HOME}/error.log"
                     filePattern="${MONITOR_LOG_HOME}/error.log.%d{yyyy-MM-dd}.%i.log">
            <ThresholdFilter level="ERROR" onMatch="ACCEPT" onMismatch="DENY"/>
            <PatternLayout pattern="${monitorPatternLayout}"/>
            <Policies>
                <TimeBasedTriggeringPolicy interval="1" />
                <SizeBasedTriggeringPolicy size="200MB"/>
            </Policies>
            <DefaultRolloverStrategy max="30"/>
        </RollingFile>

        <Async name="ASYNC_MONITOR" includeLocation="true">
            <AppenderRef ref="DEBUG_LOG_MONITOR"/>
            <AppenderRef ref="INFO_LOG_MONITOR"/>
            <AppenderRef ref="WARN_LOG_MONITOR"/>
            <AppenderRef ref="ERROR_LOG_MONITOR"/>
        </Async>

        <Async name="Async">
            <AppenderRef ref="infoLog"/>
            <AppenderRef ref="errorLog"/>
            <AppenderRef ref="warnLog"/>
            <AppenderRef ref="Console"/>
        </Async>
        <Async name="AsyncSql">
            <AppenderRef ref="sqlLog"/>
            <AppenderRef ref="Console"/>
        </Async>
        <Async name="AsyncTrace">
            <AppenderRef ref="traceIdInfo"/>
            <AppenderRef ref="infoLog"/>
            <AppenderRef ref="Console"/>
        </Async>
    </Appenders>
    <Loggers>
        <Logger name="jdbc.sqltiming" level="INFO" additivity="false">
            <AppenderRef ref="AsyncSql"/>
        </Logger>
        <Logger name="systemTraceInfo" level="INFO" additivity="false">
            <AppenderRef ref="AsyncTrace"/>
        </Logger>
        <Logger name="com.navi.eda.edaservice.metrics" level="INFO" additivity="false">
            <AppenderRef ref="ASYNC_MONITOR"/>
        </Logger>
        <Root level="INFO" >
            <AppenderRef ref="Async"/>
        </Root>
    </Loggers>
</Configuration>
```