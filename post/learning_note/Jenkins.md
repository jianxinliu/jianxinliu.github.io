# Jenkins

https://www.w3cschool.cn/jenkins/

## Jenkins Starter

[廖雪峰博客](https://www.liaoxuefeng.com/article/001463233913442cdb2d1bd1b1b42e3b0b29eb1ba736c5e000)

此demo在 Ubuntu 里搭建。

直接下载 Jenkins 的 war 包运行。运行  war 包之后，一切都在浏览器里操作。

## What is Jenkins

是一个持续集成（Continuous Integration，CI），持续交付（Continuous Delivery,CD）服务器。

## What is Pipeline in Jenkins

Jenkins 的 pipeline 是一系列插件的集合。

```groovy
Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any ---------- 1

    stages {
        stage('Build') {---------------2
            steps { -----------------3
                sh 'make' -----------------4
            }
        }
        stage('Test'){
            steps {
                sh 'make check'
                junit 'reports/**/*.xml' -----------------5
            }
        }
        stage('Deploy') {
            steps {
                sh 'make publish'
            }
        }
    }
}
```

1. `agent` 表示Jenkins 应该为Pipeline的这一部分分配一个执行者和工作区。
2. `stage`描述了这条Pipeline的一个阶段
3. `steps` 描述了要在其中运行的步骤
4. `sh` 执行给定的shell命令
5. `junit` 是由 JUnit 插件提供了用于聚合测试报告的Pipeline步骤。

Jenkins 从根本上讲是一种支持多种自动化模式的自动化引擎。Pipeline 在 Jenkins上添加了一套强大的自动化工具，**支持从简单的连续集成到全面的连续输送Pipeline的用例**。通过**建模一系列相关任务**，用户可以利用Pipeline的许多功能。

- 代码：Pipeline以代码的形式实现，通常被检入源代码控制，使团队能够编辑，审查和迭代其传送流程。
- 耐用：Pipeline可以在计划和计划外重新启动Jenkins管理时同时存在。
- Pausable：Pipeline可以选择停止并等待人工输入或批准，然后再继续Pipeline运行。
- 多功能：Pipeline支持复杂的现实世界连续交付要求，包括并行分叉/连接，循环和执行工作的能力。
- 可扩展：Pipeline插件支持其DSL的自定义扩展 以及与其他插件集成的多个选项。

[Pipeline](https://www.w3cschool.cn/jenkins/jenkins-epas28oi.html)由多个步骤组成，允许您构建，测试和部署应用程序。Jenkins Pipeline允许您**以简单的方式组合多个步骤**，可以帮助您**建模任何类型的自动化过程**。

想像一个“一步一步”，就像执行单一动作的单一命令一样。当一个步骤成功时，它移动到下一步。当步骤无法正确执行时，Pipeline将失败。

当Pipeline中的所有步骤成功完成后，Pipeline被认为已成功执行。

## Blue Ocean the Jenkins GUI
