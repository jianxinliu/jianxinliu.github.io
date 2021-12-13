# Regular Expression in JavaScript



toc

1. 什么是正则
2. 基本概念
3. 常用方法
4. 然后



# 什么是正则

Regular Expression(规则 表达式)，即[描述某种规则的表达式](https://zh.wikipedia.org/wiki/%E6%AD%A3%E5%88%99%E8%A1%A8%E8%BE%BE%E5%BC%8F#:~:text=%E7%9A%84%E6%84%8F%E6%80%9D%EF%BC%8CRegular-,Expression,-%E5%8D%B3%E2%80%9C%E6%8F%8F%E8%BF%B0%E6%9F%90%E7%A7%8D)。这里的规则也可以被看做是“模式”。



# 基本概念



# 常用方法

## 代码替换

JavaScript 中访问对象属性有两种方法：

1. object.property
2. object['property']

二者的效果是一样的。如果需要把代码中的 2 都替换成 1 的写法，手动替换是一项繁重的工作。使用正则就会简单的多。

首先识别模式：['<property>']，property 是字母或数字，是一个单词，而包裹的是引号和方括号。则可以写出表达式：

`/\[['"](\w*)['"]\]/`。其中

1.  `['"]` 可以匹配单引号或双引号。
2. `(\w*)` 表示匹配任意个字符，并且分为一组，方便替换时引用。
3. `\[`  和 `\]` 表示转义。

替换的表达式为 `.$1` 。表示将匹配到的字符用“ `.` + **匹配到的第一组的内容**”替换，其中 `$` 表示对正则表达式内分组的引用。在代码里可以使用 `string.replace(reg or string, replacement)`，在编辑器里，则可以直接替换。





Refs:

1. https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Guide/Regular_Expressions
2. https://zh.javascript.info/regexp-introduction