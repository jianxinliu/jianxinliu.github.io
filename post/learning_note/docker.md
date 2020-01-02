# Docker

概念：docker是一个给开发者、系统管理员提供开发、发布、运行应用的平台。在Linux中在容器中部署应用叫做容器化。容器并不是一个新兴事物，但是在容器中可以如此简单的部署应用是。



容器化技术越来越流行是因为：

- 灵活：就算最复杂的应用也可以被容器化
- 轻量级：多个容器分享和共用同一个宿主内核
- 可替换：支持热发布、更新和升级
- 可移植：本地构建之后，直接移动到云端或其他地方
- 可扩展
- 可堆叠



## 镜像和容器（Images and Containers）

**镜像**：是一个包含一切运行应用时所需要的资源的可运行的包，包括：代码、运行时环境、库、环境变量和配置文件。

**容器**：容器是镜像的运行时实例。容器通过运行镜像来启动。当镜像运行在内存里的时候，可以使用命令`docker ps `来查看当前运行了那些容器。

**镜像和容器的关系：**未运行时叫镜像，运行在内存中并由docker管理的叫容器，是程序运行的容器。



## 容器和虚拟机

容器在Linux上以原生的方式运行，和其他容器共享宿主机的内核。容器运行在一个独立的进程中。

虚拟机则完全是一个寄生操作系统，通过虚拟层和宿主OS交互。虚拟机也需要更多的资源。





## 命令



`docker search` 搜索可用镜像

`docker pull` 下载容器镜像。镜像都是按照`用户名/镜像名`的方式存储，所以需要写完整。

`docker run`命令由有两个参数，一个是镜像名，一个是要在镜像中运行的命令。如：

```sh
docker run learn/tutorial echo "hello world"
```

```sh
//在tutorial容器里安装ping程序,-y参数阻止 apt-get 进入交互模式，在docker环境下是无法响应这种交互的。
docker run learn/tutorial apt-get install -y ping
```

当对容器做了修改之后，可以将对容器的修改保存下来，下次直接以最新的状态来运行docker。在docker中保存状态的过程称之为committing,它保存新旧状态之间的区别，从而产生一个新版本。

#### 目标：

首先使用**docker ps -l**命令获得安装完ping命令之后容器的id。然后把这个镜像保存为learn/ping。

#### 提示：

1. 运行docker commit，可以查看该命令的参数列表。
2. 你需要指定要提交保存容器的ID。(译者按：通过docker ps -l 命令获得)
3. 无需拷贝完整的id，通常来讲最开始的三至四个字母即可区分。（译者按：非常类似git里面的版本号)

#### 正确的命令：

**docker commit 698 learn/ping**

执行完docker commit命令之后，会返回新版本镜像的id号。 

运行新命令

```sh
docker run learn/ping ping www.google.com
```

`docker ps`可以查看当前运行的容器

`docker inspect`可以更详细的查看某容器的信息

例如：

```sh
$ docker ps 
ID               IMAGE                    
efedcsf34sd      learn/ping:lastest
$ docker inspect efe
```

