# 在 Linux 手动安装 python

参考： https://www.cnblogs.com/lsdb/p/11935843.html

## 下载

下载 python 压缩包

## 安装

```shell
# 以python 3.8，安装到/opt/python38Env为例
# 和常规的编译安装过程一样

# 解压
tar -zxf Python-3.8.0.tgz

# 进入目录
cd Python-3.8.0

# 预编译
./configure --prefix=/opt/python38Env

# 编译
make

# 将编译结果复制到预编译时指定的路径
make install

```

## python import 从哪儿导入

python 程序中 import 包时从 `sys.path` 指向的那些目录下 import。

默认情况下 `sys.path[0]` 是当前被运行python文件所处的目录（当没有被执行文件时为空），其他是当前所使用 python 命令的 `../lib/` 预设置的文件(夹)。

对于用户而言有两个办法修改 `sys.path`，一种方法是python会将环境变量 PYTHONPATH（冒号分隔）解析加到 `sys.path`，所以要加入的目录直接在 `~/.bashrc` 等文件中赋给 PYTHONPATH 即可。另外一种方法是 `sys.path` 本质就是一个列表，所以可以直接在python代码中使用 `sys.path.insert()`、`sys.path.append()` 进行添加。

### 复制依赖库

复制项目需要的依赖库到 `<python_path>/lib/python3.8/site-packages/`

如果是虚拟环境，则复制到虚拟环境里

## 创建虚拟环境

```shell
# 创建虚拟环境
# 含义：python3 调用 venv 模块，创建一个名叫 test_env 的虚拟环境
# 本质上是把 Python-3.8.0 文件夹复制一份到当前目录下，并重命名为 pyvenv
# 创建的虚拟环境默认是在当前执行创建命令的目录下
python -m venv pyvenv

# 使用新建的虚拟环境
source pyvenv/bin/activate

# 退出上边激活的虚拟环境
# 本质是 pyvenv/bin/activate 中的 deactivate 方法
deactivate

# 删除虚拟环境
# 毕竟只是创建了个文件夹，所以要删创建的虚拟环境，直接把整个文件夹删除即可
rm -rf pyvenv
```

## 依赖及版本

```shell
pip install numpy==1.20.1 pandas==1.2.4 matplotlib==3.3.4 scipy==1.6.2 -i https://pypi.tuna.tsinghua.edu.cn/simple
```

## 可能的问题

### ModuleNotFoundError: No module named '_ctypes'

参考： https://www.pythonpool.com/modulenotfounderror-no-module-named-_ctypes-solved/

```shell
yum install libffi-devel -y
make install

# 安装完 libffi-devel 库之后，需要重新编译 python
```