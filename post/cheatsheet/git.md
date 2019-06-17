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
