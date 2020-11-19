使用十分简单，使用已有配置，在配置文件中指定日志的配置文件

配合 Lombok 的 @Slf4j 注解，直接使用 log 变量即可打 log

关键概念：Logger、Appenders and Layouts
此三个类用于控制不同类型不同级别的消息以什么格式输出到哪里。