# [廖雪峰GIT教程](https://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000)

## 设置Git

**创建提交的身份:**`git config --global user.name 'xxx'`
​			    `git config --global user.email 'xxx'`

## 创建版本库

**初始化仓库：**`git init`，将当前目录变成git可以管理的仓库，此时，此目录下多出一个`.git`目录

**添加文件：**`git add <hello.txt>`  # 把所有要提交的文件修改放到缓存区,此`hello.txt`文件需要放在当前仓库中,也就是`git init`命令执行的目录。可以多次`add`文件，然后一起`commit`

**提交文件：**`git commit -m 'add a file'` # 把暂存区的所有内容（一次可以提交多个文件）提交到当前分支,引号内为此次提交的说明，**修改文件后需要再次提交**

## 时光机穿梭

**查看仓库状态：**`git status` #掌握仓库当前状态,即当前`working tree`的状态，有没有文件没有提交、被修改……

**查看文件的修改内容：**`git diff ***.txt` #查看文件修改内容,查看某文件的修改内容

### 版本回退

**查看提交历史：**`git log`，显示从最近到最远的提交日志。 包括版本号、提交者、提交日期和提交时添加的说明。一个简洁版本`git log --pretty=oneline`

**回退：**`git reset`。`HEAD`表示当前版本，`HEAD^`表示上一个版本，`HEAD^^`表示上上个版本，`HEAD~100`表示前100个版本

[--hard参数](http://blog.csdn.net/carolzhang8406/article/details/49761927)

`git reset --hard HEAD^` #回退到上一个版本，此时执行`git log`命令，显示版本变少了一个，只要记得回退之前的版本号（**使用`git reflog`显示每一次操作的版本号**），还可以用`git reset --hard 版本号 `回退到之前的版本，版本好不需要写全，git会自动查找

**git内部维护一个链表，HEAD表示当前版本，当进行版本切换时，只是链表头指针进行移动，并更新工作区文件**

**查看历史命令：**`git reflog` 

**小结：**

- `HEAD`指向的版本就是当前版本，因此，Git允许我们在版本的历史之间穿梭，使用命令`git reset --hard commit_id`。
- 穿梭前，用`git log`可以查看提交历史，以便确定要回退到哪个版本。
- 要重返未来，用`git reflog`查看命令历史，以便确定要回到未来的哪个版本。


### 工作区和暂存区


**工作区（Working Directory）：**就是`git init`执行时的目录
**版本库（Repository）** 工作区的隐藏目录`.git`：

	stage(index) 暂存区,git add 命令就是将提交的修改放到stage，commit 命令一次性将stage中的内容提交到分支
	master Git自动创建的分支
	HEAD 指针，指向master
![git 原理图](DontMove.jpg)

### 管理修改

**Git跟踪管理的是修改（包括新增文件），而不是文件。每次修改，如果不`add`到缓存区，就不会加入到`commit`中**

### 撤销修改

**查看工作区和版本库里最新版本的区别：**`git diff HEAD -- readme.md`
**撤销在工作区的修改，提交之后则不可撤销，只能版本回退（前提是还未将本地版本库推送到远程）：**`git checkout -- <file>` 用版本库的版本替换工作区的版本，无论是工作区的修改还是删除，都可以'一键还原'，丢弃工作区的修改（让文件回到最近一次的git commit或git add时的状态）

**撤销暂存区的修改，重新放回工作区。**`git reset HEAD <file>`

### 删除文件

**工作区中删除文件：**`rm file`。若删除的文件已经提交到版本库，删除文件后，可从版本库恢复.`git status`会显示哪些文件被删，此时可选择：

- **从版本库中删除文件：**`git rm file`，并且`git commit`

-  **从版本库恢复文件，但是会丢失最近一次提交后修改的内容：**`git checkout -- file`


## 远程仓库

本地Git仓库通过ssh加密和github仓库之间传输，设置如下：

1. 创建SSH Key :`ssh-keygen -t rsa -C 'user@example.com' `

   你需要把邮件地址换成你自己的邮件地址，然后一路回车，使用默认值即可，由于这个Key也不是用于军事目的，所以也无需设置密码。

   如果一切顺利的话，可以在用户主目录里找到`.ssh`目录，里面有`id_rsa`和`id_rsa.pub`两个文件，这两个就是SSH Key的秘钥对，`id_rsa`是私钥，不能泄露出去，`id_rsa.pub`是公钥，可以放心地告诉任何人。

2. 登录github添加公钥

### 添加远程库

**本地仓库关联远程仓库：**`git remote add origin git@github.com:username/repostery.git` 远程库的名字为origin

第一次使用Git的clone或者push命令连接GitHub时需确认

**把本地仓库（实际是将master分支）推送到远程：**`git push -u origin master` #第一次把当前分支master推送到远程，-u参数不但推送，而且将本地的分支和远程的分支关联起来,在以后的推送或者拉取时就可以简化命令。

**把当前分支master推送到远程:**`git push origin master` 只要本地作了修改，就可以推送到远程

### 克隆远程库

**从远程库克隆到本地库:**`git clone git@github.com:username/repostery.git` 

	git支持多种协议，包括https，但通过试试支持原生git协议速度最快
## 分支管理

分支在实际中有什么用呢？假设你准备开发一个新功能，但是需要两周才能完成，第一周你写了50%的代码，如果立刻提交，由于代码还没写完，不完整的代码库会导致别人不能干活了。如果等代码全部写完再一次提交，又存在丢失每天进度的巨大风险。

你创建了一个属于你自己的分支，别人看不到，还继续在原来的分支上正常工作，而你在自己的分支上干活，想提交就提交，直到开发完毕后，再一次性合并到原来的分支上，这样，既安全，又不影响别人工作。

### 创建与合并分支

Git鼓励大量使用分支完成任务

**创建并切换分支:**`git checkout -b dev` 相当于`git branch dev` 和`git checkout dev `一起执行

**查看当前分支，当前分支前有个*号:**`git branch` 

**创建分支:**`git branch <name>`

**切换分支:**`git checkout <name>` 

**合并某个分支到当前分支:**`git merge <name>` 

**删除分支:**`git branch -d <name>` 合并之后就可删除

### 解决冲突

当需要合并的分支都有了各自新的提交，Git无法快速合并，此时会产生冲突，需要先手动解决冲突后再提交

`git status`查看冲突的文件。`cat file`查看文件内容

**查看分支合并图:**`git log --graph`

### 分支管理策略

使用Fast forword模式，删除分支后，会丢掉分支信息

**禁用Fast forward(快进模式：直接将master指针指向合并的分支)合并dev分支:**`git merge --no-ff -m 'message' dev` 

	本次合并要创建新的commit，所以要加上-m参数，把commit描述写进去
	Fast forward合并不可查看合并记录
**分支管理策略：**

1.  master分支应该是非常稳定的，仅用来发布新版本，平时不能在master分支上工作
2.  在dev分支上工作，每个人再在dev分支上开分支

团队合作的分支情况：

![团队合作的分支情况](git_branch.png)

### Bug分支

**隐藏当前工作现场，等恢复后继续工作**`git stash` 
**查看stash记录**`git stash list` 
**仅恢复现场，不删除stash内容**`git stash apply` 你可以多次stash，恢复的时候，先用`git stash list`查看，然后恢复指定的stash，用命令：`git stash apply stash@{0}`

**删除stash内容**`git stash drop` 
**恢复现场的同时删除stash内容**`git stash pop` 

### Feature分支

每添加一个新功能，最好新建一个feature分支，在上面开发，完成后，合并，最后，删除该feature分支。

**强行删除某个未合并的分支**`git branch -D <name>` 

### 多人合作

**查看远程仓库**`git remote`
**查看远程库详细信息**`git remote -v` 

#### 推送分支

将该分支上所有本地提交推送到远程 `git push origin master/dev`

#### 抓取分支

**抓取远程提交**`git pull` ,需要先建立本地分支和远程分支的关联

**在本地创建和远程分支对应的分支**`git checkout -b branch-name origin/branch-name` 
**建立本地分支和远程分支的关联**`git branch --set-upstream branch-name origin/branch-name` 

多人合作的工作模式通常是这样的：

1. 首先，尝试用`git push origin branch-name`推送自己的修改
2. 如果推送失败，则因为远程分支比本地更新，需要先用`git pull`试图合并
3. 如果合并有冲突，则手动解决冲突，并在本地提交
4. 没有冲突或者解决冲突后就可以用`git push origin branch-name`推送成功。

如果`git pull`提示“no tracking information”，则说明本地分支和远程分支的链接关系没有创建，用命令`git branch --set-upstream branch-name origin/branch-name`。

## 标签管理

发布一个版本时，我们通常先在版本库中打一个标签（tag就是一个有意义的名字，跟某个commit绑定在一起），这样，就唯一确定了打标签时刻的版本。将来无论什么时候，取某个标签的版本，就是把那个打标签的时刻的历史版本取出来。所以，**标签也是版本库的一个快照**。但其实它就是指向某个commit的指针。

### 创建标签

切换到需要打标签的分支上。**默认标签是打在最新提交的commit上的**

`git tag v1.0` **给当前分支最新的commit打上"v1.0"的标签**
`git tag v0.9 36df530` **给历史提交的commit打标签**
`git tag -a v0.1 -m 'version 0.1 released' 3628164` **-a指定标签名，-m指定说明文字**，创建带有说明的标签
`git tag -s <tagname> -m 'blabla'` **可以用PGP签名标签**
`git tag` **查看所有标签**,按字母排序
`git show v1.0` **查看标签信息**

### 操作标签

`git tag -d v0.1` **删除标签**,创建的标签村在本地
`git push origin <tagname>` **推送某个标签到远程**
`git push origin --tags` **推送所有尚未推送的本地标签**
**删除远程标签**

1.  `git tag -d v0.2` #先删除本地标签
    2. `git push origin :refs/tags/v0.2` #删除远程标签

## 使用GitHub

参与开源项目bootstrap：

1. 在bootstrap项目中点Fork,在自己的账号克隆了一个bootstrap仓库
2. 在自己的账号下clone到本地

一定要从自己的账号下clone仓库，这样你才能推送修改。如果从bootstrap的作者的仓库地址`git@github.com:twbs/bootstrap.git`克隆，因为没有权限，你将不能推送修改。

如果你想修复bootstrap的一个bug，或者新增一个功能，立刻就可以开始干活，干完后，往自己的仓库推送。

如果你希望bootstrap的官方库能接受你的修改，你就可以在GitHub上发起一个pull request。当然，对方是否接受你的pull request就不一定了。

## 自定义git

**让Git显示颜色：**`git config --global color.ui true`

### 忽略特殊文件

编写.gitignore文件（将要忽略的文件名放入此文件中）来忽略某些文件，此文件本身要放到版本库内，并可对其做版本管理

忽略文件的原则是：

1. 忽略操作系统自动生成的文件，比如缩略图等；
2. 忽略编译生成的中间文件、可执行文件等，也就是如果一个文件是通过另一个文件自动生成的，那自动生成的文件就没必要放进版本库，比如Java编译产生的`.class`文件；
3. 忽略你自己的带有敏感信息的配置文件，比如存放口令的配置文件。



`git add -f hello.pyc` **-f参数强制添加到Git**
`git check-ignore -v hello.pyc`　**检查.gitignore文件的规则**

### 配置别名

`git config --global alias.st status`,将`git status`简写成`git st`

简写命令
`git config --global alias.co checkout` #简写checkout命令
`git config --global alias.st status`
`git config --global alias.ci commit`
`git config --global alias.br branch`
`git config --global alias.unstage 'reset HEAD'` #撤销暂存区的修改
`git config --global alias.last 'log -1'` #查看最近一次的提交
`git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"`



配置文件
**--global参数时针对当前用户起作用，如果不加，仅针对当前仓库起作用**
每个仓库的Git配置文件在 .git/config 文件中.要删除别名，打开文件将alias后的对应部分删除
当前用户的Git配置文件在用户主目录下的 .gitconfig 文件中

### 搭建Git服务器

1、安装`git sudo apt install git`
2、创建git用户 `sudo adduser git`
3、创建证书登录，将所有需要登录的用户的公钥导入到/home/git/.ssh/authorized_keys文件，每行一个
4、初始化Git仓库

	在仓库目录下输入命令 sudo git init --bare sample.git 创建裸仓库（没有工作区）
	把owner改为git sudo chown -R git:git sample.git
5、禁用shell登录，修改/etc/passwd文件
	git:x:1001:1001:,,,:/home/git:/bin/bash
	改为：
	git:x:1001:1001:,,,:/home/git:/usr/bin/git-shell

6、克隆远程仓库
git clone git@server:/srv/sample.git
