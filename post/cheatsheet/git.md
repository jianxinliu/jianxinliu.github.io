# Git Commands

## Base and setting

```shell
# initial a folder as git repository
git init

# config
git config [--global] user.name[email] 'xxx'

```

## 基本操作

```shell

# add changes to staged
git add [file]

# commit changes to current branch of local repository
git commit -m 'commit message'

# show status of working tree
git status

# 查看提交记录
git log [--graph]
```

## 远程

```shell
# 本地仓库关联远程仓库
git remote add origin repo_url

# 更新本地仓库到远程
git push origin branch_name

# clone
git clone repo_url

# 抓取远程到本地
git pull
```

## 分支

```shell
# 创建并切换分支，以下命令相当于 git branch name && git checkout name
git checkout -b branch_name

# 查看当前分支
git branch

# 切换分支
git checkout branch_name

#  合并某个分支到当前分支
git merge branch_name

# 删除分支，合并之后就可删除
git branch -d branch_name
```

## 版本控制最佳实践

> - **鼓励频繁地提交** 。SVN的初学者可能会有一种想法，他们保留代码一直在本地修改，直到代码确定没有问题了才提交。但最佳实践是频繁地提交，而不要等到代码没问题了再一次性提交。对于可能损坏主干原则的代码，不要直接提交到主干，而是创建一个分支，在分支中频繁提交。
> - **确定分支流程**。基本上所有的特性和较大的bug修复都应该使用分支来修改。
> - **定义主干原则，并且坚守它**。我们团队的主干原则是“主干对应的代码必须是可以发布并且不会产生bug的”，如果不能保证新增的或者修改的代码符合这一原则，就在分支提交代码。任何人破坏这一原则引起bug，就请大家吃饭。
> - **不要把逻辑的修改和代码格式化操作混在一起**。如果您做了一些代码格式化的操作，就单独提交这次修改。比如您去掉了代码中所有的空行，那就单独提交一个commit，然后再做一些逻辑的修改，再提交。这样可以避免“天哪，所有的东西都不一样了”，出现问题之后更容易追溯。
> - **不相干的代码分开提交**。也就是说不要在一次提交里修复两个bug。
> - **保持工作代码库的“干净”**。如果您有文件不想也不需要提交，就加入到忽略列表（ignorelist）。不需要提交的文件包括编译后文件、配置文件和第三方依赖等。这样的好处是，您每次打开SVN提交界面，如果没有修改过任何代码，就会看见一个空的列表，如果修改过代码，就显示修改过的代码。这能提醒您不要漏掉任何需要提交的文件
>
> 《Web 全栈工程师的自我修养》——版本控制