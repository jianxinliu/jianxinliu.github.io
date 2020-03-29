# .gitignore 语法

- 空行或是以`#`开头的行即注释行将被忽略。

- 可以在前面添加正斜杠`/`来避免递归,下面的例子中可以很明白的看出来与下一条的区别。

- 可以在后面添加正斜杠`/`来忽略文件夹，例如`build/`即忽略build文件夹。

- 可以使用`!`来否定忽略，即比如在前面用了`*.apk`，然后使用`!a.apk`，则这个a.apk不会被忽略。（声明例外）

- `*`用来匹配零个或多个字符，如`*.[oa]`忽略所有以".o"或".a"结尾，`*~`忽略所有以`~`结尾的文件（这种文件通常被许多编辑器标记为临时文件）；`[]`用来匹配括号内的任一字符，如`[abc]`，也可以在括号内加连接符，如`[0-9]`匹配0至9的数；`?`用来匹配单个字符。

Spring start.spring.io 给出的模板

```.gitignore
HELP.md
target/ # 忽略整个文件夹
!.mvn/wrapper/maven-wrapper.jar
!**/src/main/** # 该文件夹下不被忽略
!**/src/test/**

### STS ###
.apt_generated
.classpath
.factorypath
.project
.settings
.springBeans
.sts4-cache

### IntelliJ IDEA ###
.idea
*.iws # 任何以 .iws 结尾的文件
*.iml
*.ipr

### NetBeans ###
/nbproject/private/ # 非递归的忽略 nbproject/private/ 文件夹
/nbbuild/
/dist/
/nbdist/
/.nb-gradle/
build/

### VS Code ###
.vscode/
```

[ A collection of useful .gitignore templates ]( https://github.com/github/gitignore )

# yaml 语法

```yaml
# 注释

# 对象
animal:pets # or
animal:{name:pets}

# 数组
animal:
    - cat
    - dog
    - fish
# or 
animal:[cat,dog,fish]

# 纯量
data: 123.3
data2: hello
happy: ~ # ~ 表示 null
live: true
# # 时间日期，转换成 js 之后，会进行转换 new Date(value)
iso8601: 2001-12-14t21:59:43.10-05:00 
date: 1976-07-31
# # 数据类型强转
e: !!str 123  
f: !!str true  # 转成 js 为：{e:'123',f:'true'}

# 引用,锚点`&` 和 别名 `*`
defaults: &defaults # 给defaults 起别名，便于引用
  adapter:  postgres
  host:     localhost

development:
  database: myapp_development
  <<: *defaults  # << 表示合并到当前，* 引用 &defaults
```

# npm 操作命令

```shell
# 初始化项目，创建 package.json
npm init [--yes] # --yes 表示全部使用默认配置

# 全局安装依赖
npm install -g

# 安装依赖到项目,并修改 package.json 文件（添加依赖到 dependencies 属性）
npm install <module> --save

# 安装开发依赖，并修改 package.json 文件（添加依赖到 devDependencies 属性）
npm install <module> --save-dev

# 安装备选依赖（fallback）（添加依赖到 optionalDencies 属性）
npm instal <module> --save-optional

# 查看安装目录，本地 & 全局
npm root [-g]

# 卸载模块，本地 & 全局
npm uninstall [-g]

# 更新模块，本地 & 全局
npm update [-g]

# 查看模块信息
npm info <module>
```

