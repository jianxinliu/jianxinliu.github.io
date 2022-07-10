# Jquery

[JQuery参考手册](http://www.w3school.com.cn/jquery/jquery_reference.asp)

# jquery特性

- HTML元素选取
- HTML 元素操作
- CSS 操作
- HTML 事件函数
- JavaScript 特效和动画
- HTML DOM 遍历和修改
- AJAX
- Utilities

## 向页面添加 jquery（CDN）方式

**提示**：使用谷歌或微软的 jQuery，有一个很大的优势：

许多用户在访问其他站点时，已经从谷歌或微软加载过 jQuery。所有结果是，当他们访问您的站点时，会从缓存中加载 jQuery，这样可以减少加载时间。同时，大多数 CDN 都可以确保当用户向其请求文件时，会从离用户最近的服务器上返回响应，这样也可以提高加载速度。

**Google 的 CDN**

```HTML
<head>
<script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs
/jquery/1.4.0/jquery.min.js"></script>
</head>
```

**提示**：通过 Google CDN 来获得最新可用的版本：
如果您观察上什么的 Google URL - 在 URL 中规定了 jQuery 版本 (1.8.0)。如果您希望使用最新版本的 jQuery，也可以从版本字符串的末尾（比如本例 1.8）删除一个数字，谷歌会返回 1.8 系列中最新的可用版本（1.8.0、1.8.1 等等），或者也可以只剩第一个数字，那么谷歌会返回 1 系列中最新的可用版本（从 1.1.0 到 1.9.9）。
**Microsoft  的 CDN**

```html
<head>
<script type="text/javascript" src="http://ajax.microsoft.com/ajax/jquery
/jquery-1.4.min.js"></script>
</head>
```

# Jquery 基本语法

基础语法是：$(selector).action()

- 美元符号定义 Jquery
- 选择符（selector）“查询”和“查找” HTML 元素
- jQuery 的 action() 执行对元素的操作

## 文档就绪函数

```javascript
$(document).ready(function(){
  ---- jQuery functions go here ----
})
```

这是为了防止文档在完全加载（就绪）之前运行 jQuery 代码。

如果在文档没有完全加载之前就运行函数，操作可能失败。下面是两个具体的例子：

- 试图隐藏一个不存在的元素
- 获得未完全加载的图像的大小

# jQuery 元素选择器

jQuery 使用 CSS 选择器来选取 HTML 元素。

$("p") 选取 <p> 元素。

$("p.intro") 选取所有 class="intro" 的 <p> 元素。

$("p#demo") 选取所有 id="demo" 的 <p> 元素。

## jQuery 属性选择器

jQuery 使用 XPath 表达式来选择带有给定属性的元素。

$("[href]") 选取所有带有 href 属性的元素。

$("[href='#']") 选取所有带有 href 值等于 "#" 的元素。

$("[href!='#']") 选取所有带有 href 值不等于 "#" 的元素。

$("[href\$='.jpg']") 选取所有 href 值以 ".jpg" 结尾的元素。

## jQuery CSS 选择器

jQuery CSS 选择器可用于改变 HTML 元素的 CSS 属性。

下面的例子把所有 p 元素的背景颜色更改为红色：

### 实例

```javascript
$("p").css("background-color","red");
```

## 更多的选择器实例

| 语法                   | 描述                                         |
| -------------------- | ------------------------------------------ |
| $(this)              | 当前 HTML 元素                                 |
| $("p")               | 所有 <p> 元素                                  |
| $("p.intro")         | 所有 class="intro" 的 <p> 元素                  |
| $(".intro")          | 所有 class="intro" 的元素                       |
| $("#intro")          | id="intro" 的元素                             |
| $("ul li:first")     | 每个 <ul> 的第一个 <li> 元素                       |
| $("[href\$='.jpg']") | 所有带有以 ".jpg" 结尾的属性值的 href 属性               |
| $("div#intro .head") | id="intro" 的 <div> 元素中的所有 class="head" 的元素 |

# Jquery 事件

## jQuery 名称冲突

jQuery 使用 $ 符号作为 jQuery 的简介方式。

某些其他 JavaScript 库中的函数（比如 Prototype）同样使用 $ 符号。

jQuery 使用名为 noConflict() 的方法来解决该问题。

*var jq=jQuery.noConflict()*，帮助您使用自己的名称（比如 jq）来代替 $ 符号。

## 书写规则

由于 jQuery 是为处理 HTML 事件而特别设计的，那么当您遵循以下原则时，您的代码会更恰当且更易维护：

- 把所有 jQuery 代码置于事件处理函数中
- 把所有事件处理函数置于文档就绪事件处理器中
- 把 jQuery 代码置于单独的 .js 文件中
- 如果存在名称冲突，则重命名 jQuery 库

## jQuery 事件

下面是 jQuery 中事件方法的一些例子：

| Event 函数                        | 绑定函数至                   |
| ------------------------------- | ----------------------- |
| $(document).ready(function)     | 将函数绑定到文档的就绪事件（当文档完成加载时） |
| $(selector).click(function)     | 触发或将函数绑定到被选元素的点击事件      |
| $(selector).dblclick(function)  | 触发或将函数绑定到被选元素的双击事件      |
| $(selector).focus(function)     | 触发或将函数绑定到被选元素的获得焦点事件    |
| $(selector).mouseover(function) | 触发或将函数绑定到被选元素的鼠标悬停事件    |

# jQuery 效果 - 隐藏和显示

## jQuery hide() 和 show()

通过 jQuery，您可以使用 hide() 和 show() 方法来隐藏和显示 HTML 元素：

```javascript
$("#hide").click(function(){
  $("p").hide();
});

$("#show").click(function(){
  $("p").show();
});
```

### 语法：

```
$(selector).hide(speed,callback);

$(selector).show(speed,callback);
```

可选的 speed 参数规定隐藏/显示的速度，可以取以下值："slow"、"fast" 或毫秒。

可选的 callback 参数是隐藏或显示完成后所执行的函数名称。

下面的例子演示了带有 speed 参数的 hide() 方法：

### 实例

```
$("button").click(function(){
  $("p").hide(1000);
});
```

## jQuery toggle()

通过 jQuery，您可以使用 toggle() 方法来切换 hide() 和 show() 方法。

显示被隐藏的元素，并隐藏已显示的元素：

### 实例

```
$("button").click(function(){
  $("p").toggle();
});
```

### 语法：

```
$(selector).toggle(speed,callback);
```

可选的 speed 参数规定隐藏/显示的速度，可以取以下值："slow"、"fast" 或毫秒。

可选的 callback 参数是 toggle() 方法完成后所执行的函数名称。

# jQuery 效果 - 淡入淡出

## jQuery Fading 方法

通过 jQuery，您可以实现元素的淡入淡出效果。

jQuery 拥有下面四种 fade 方法：

可选的 speed 参数规定效果的时长。它可以取以下值："slow"、"fast" 或毫秒。

fadeTo() 方法中必需的 opacity 参数将淡入淡出效果设置为给定的不透明度（值介于 0 与 1 之间）。

可选的 callback 参数是 fading 完成后所执行的函数名称。

- fadeIn() 淡入

```javascript
$(selector).fadeIn(speed,callback);
```

- fadeOut() 淡出

```
$(selector).fadeOut(speed,callback);
```

- fadeToggle() 淡入和淡出切换

```
$(selector).fadeToggle(speed,callback);
```

- fadeTo() 渐变为给定的不透明度（值介于 0 与 1 之间)

```
$(selector).fadeTo(speed,opacity,callback);
```

# jQuery 效果 - 滑动

## jQuery 滑动方法

通过 jQuery，您可以在元素上创建滑动效果。

jQuery 拥有以下滑动方法：

- slideDown()

```
$(selector).slideDown(speed,callback);
```

- slideUp()

```
$(selector).slideUp(speed,callback);
```

- slideToggle()

```
$(selector).slideToggle(speed,callback);
```

# jQuery 效果 - 动画

## jQuery 动画 - animate() 方法

jQuery animate() 方法用于创建自定义动画。

### 语法：

```
$(selector).animate({params},speed,callback);
```

必需的 params 参数定义形成动画的 CSS 属性。

可选的 speed 参数规定效果的时长。它可以取以下值："slow"、"fast" 或毫秒。

可选的 callback 参数是动画完成后所执行的函数名称。

下面的例子演示 animate() 方法的简单应用；它把 <div> 元素移动到左边，直到 left 属性等于 250 像素为止：

### 实例

```
$("button").click(function(){
  $("div").animate({left:'250px'});
}); 
```

**提示**：默认地，所有 HTML 元素都有一个静态位置，且无法移动。

如需对位置进行操作，要记得首先把元素的 CSS position 属性设置为 relative、fixed 或 absolute！

## jQuery animate() - 操作多个属性

请注意，生成动画的过程中可同时使用多个属性：

### 实例

```
$("button").click(function(){
  $("div").animate({
    left:'250px',
    opacity:'0.5',
    height:'150px',
    width:'150px'
  });
}); 
```

**提示**：可以用 animate() 方法来操作所有 CSS 属性吗？

是的，几乎可以！不过，需要记住一件重要的事情：**当使用 animate() 时，必须使用 Camel 标记法书写所有的属性名**，比如，必须使用 paddingLeft 而不是 padding-left，使用 marginRight 而不是 margin-right，等等。

同时，色彩动画并不包含在核心 jQuery 库中。

如果需要生成颜色动画，您需要从 jQuery.com 下载 Color Animations 插件。

## jQuery animate() - 使用相对值

也可以定义相对值（该值相对于元素的当前值）。需要在值的前面加上 += 或 -=：

### 实例

```
$("button").click(function(){
  $("div").animate({
    left:'250px',
    height:'+=150px',
    width:'+=150px'
  });
});
```

## jQuery animate() - 使用预定义的值

您甚至可以把属性的动画值设置为 "show"、"hide" 或 "toggle"：

### 实例

```
$("button").click(function(){
  $("div").animate({
    height:'toggle'
  });
});
```

## jQuery animate() - 使用队列功能

**默认地，jQuery 提供针对动画的队列功能**。

这意味着如果您在彼此之后编写多个 animate() 调用，jQuery 会创建包含这些方法调用的“内部”队列。然后逐一运行这些 animate 调用。

### 实例 1

隐藏，如果您希望在彼此之后执行不同的动画，那么我们要利用队列功能：

```
$("button").click(function(){
  var div=$("div");
  div.animate({height:'300px',opacity:'0.4'},"slow");
  div.animate({width:'300px',opacity:'0.8'},"slow");
  div.animate({height:'100px',opacity:'0.4'},"slow");
  div.animate({width:'100px',opacity:'0.8'},"slow");
});
```

### 实例 2

下面的例子把 <div> 元素移动到右边，然后增加文本的字号：

```
$("button").click(function(){
  var div=$("div");
  div.animate({left:'100px'},"slow");
  div.animate({fontSize:'3em'},"slow");
});
```

## jQuery Callback 函数

当动画 100% 完成后，即调用 Callback 函数。

### 典型的语法：

```
$(selector).hide(speed,callback)
```

*callback* 参数是一个在 hide 操作完成后被执行的函数。

```
$("p").hide(1000,function(){
alert("The paragraph is now hidden");
});
```

# jQuery - Chaining(链式语法)

例子：下面的例子把 css(), slideUp(), and slideDown() 链接在一起。"p1" 元素首先会变为红色，然后向上滑动，然后向下滑动

```
$("#p1").css("color","red")
  .slideUp(2000)
  .slideDown(2000);
```

# jQuery - 获得内容和属性

## jQuery DOM 操作

## 获得内容 - text()、html() 以及 val()

三个简单实用的用于 DOM 操作的 jQuery 方法：

- text() - <u>设置或返回</u>所选元素的**文本内容**
- html() - <u>设置或返回</u>所选元素的**内容（包括 HTML 标记）**
- val() - <u>设置或返回</u>**表单字段的值**

下面的例子演示如何通过 jQuery text() 和 html() 方法来获得内容：

### 实例

```
$("#btn1").click(function(){
  alert("Text: " + $("#test").text());
});
$("#btn2").click(function(){
  alert("HTML: " + $("#test").html());
});
```

下面的例子演示如何通过 jQuery val() 方法获得输入字段的值：

### 实例

```
$("#btn1").click(function(){
  alert("Value: " + $("#test").val());
});
```

## 获取属性 - attr()

jQuery attr() 方法用于获取属性值。

下面的例子演示如何获得链接中 href 属性的值：

### 实例

```
$("button").click(function(){
  alert($("#w3s").attr("href"));
});
```

# jQuery - 设置内容和属性

text("")

html("")

val("")

## text()、html() 以及 val() 的回调函数

上面的三个 jQuery 方法：text()、html() 以及 val()，同样拥有回调函数。回调函数由两个参数：被选元素列表中当前元素的下标，以及原始（旧的）值。然后以函数新值返回您希望使用的字符串。

下面的例子演示带有回调函数的 text() 和 html()：

### 实例

```
$("#btn1").click(function(){
  $("#test1").text(function(i,origText){
    return "Old text: " + origText + " New text: Hello world!
    (index: " + i + ")";
  });
});

$("#btn2").click(function(){
  $("#test2").html(function(i,origText){
    return "Old html: " + origText + " New html: Hello <b>world!</b>
    (index: " + i + ")";
  });
});
```

## 设置属性 - attr()

jQuery attr() 方法也用于设置/改变属性值。

当需要设置多个属性时，可以将多个属性以数组形式传入

下面的例子演示如何改变（设置）链接中 href 属性的值：

### 实例

```
$("button").click(function(){
  $("#w3s").attr("href","http://www.w3school.com.cn/jquery");
});
```

## attr() 的回调函数

jQuery 方法 attr()，也提供回调函数。回调函数由两个参数：被选元素列表中当前元素的下标，以及原始（旧的）值。然后以函数新值返回您希望使用的字符串。**同其他三个方法**

下面的例子演示带有回调函数的 attr() 方法：

### 实例

```
$("button").click(function(){
  $("#w3s").attr("href", function(i,origValue){
    return origValue + "/jquery";
  });
});
```

# jQuery - 添加元素

## 添加新的 HTML 内容

我们将学习用于添加新内容的四个 jQuery 方法：元素可以以参数的形式，追加无数个$("p").append(txt1,txt2,txt3); 

- append() - 在被选元素的结尾插入内容
- prepend() - 在被选元素的开头插入内容
- after() - 在被选元素之后插入内容
- before() - 在被选元素之前插入内容

# jQuery - 删除元素

## 删除元素/内容

如需删除元素和内容，一般可使用以下两个 jQuery 方法：

- remove() - 删除被选元素（及其子元素），可接受参数对元素进行过滤
- empty() - 从被选元素中删除子元素

# jQuery - 获取并设置 CSS 类

## jQuery 操作 CSS

jQuery 拥有若干进行 CSS 操作的方法。我们将学习下面这些：

- addClass() - 向被选元素添加一个或多个类
- removeClass() - 从被选元素删除一个或多个类
- toggleClass() - 对被选元素进行添加/删除类的切换操作
- css() - 设置或返回样式属性

# jQuery - css() 方法

## jQuery css() 方法

css() 方法设置或返回被选元素的一个或多个样式属性。

## 返回 CSS 属性

如需返回指定的 CSS 属性的值，请使用如下语法：

```
css("propertyname");
```

## 设置 CSS 属性

如需设置指定的 CSS 属性，请使用如下语法：将数组作为参数传入可以设置多个属性

```
css("propertyname","value");
```

# jQuery - 尺寸

## jQuery 尺寸 方法

jQuery 提供多个处理尺寸的重要方法：

- width() 设置或返回元素的宽度（不包括内边距、边框或外边距）
- height() 设置或返回元素的高度（不包括内边距、边框或外边距）
- innerWidth() 返回元素的宽度（包括内边距）。
- innerHeight() 返回元素的高度（包括内边距）。
- outerWidth() 返回元素的宽度（包括内边距和边框）。
- outerHeight() 返回元素的高度（包括内边距和边框）
- outerWidth(true) 方法返回元素的宽度（包括内边距、边框和外边距）
- outerHeight(true) 方法返回元素的高度（包括内边距、边框和外边距）。

# jQuery遍历

## 祖先（用于向上遍历DOM树）

- parent() 返回节点的直接父节点
- parents() 返回节点的所有祖先节点，一直向上到文档根节点（<html>）可选用参数来过滤
- parentUntil(node) 返回介于两者之间的祖先节点

## 后代 （用于向下遍历DOM树）

- children() 返回节点的所有直接子元素，可选用参数来过滤
- find() 返回节点的所有后代元素，直到最后一个后代，可选用参数来过滤，find("*")返回所有后代

## 同胞 （用于水平遍历DOM树）

- siblings() 返回所有同胞元素，可选用参数来过滤
- next() 以当前节点为标准，返回下一个同胞元素
- nextAll()  以当前节点为标准，返回所有跟随的同胞元素
- nextUntil() 返回介于两个给定参数之间的所有跟随的同胞元素
- prev()  prev(), prevAll() 以及 prevUntil() 方法的工作方式与上面的方法类似，只不过方向相反而已：它们返回的是前面的同胞元素（在 DOM 树中沿着同胞元素向后遍历，而不是向前）。
- prevAll()
- prevUntil()

# jQuery 遍历 - 过滤

## 缩写搜索元素的范围

三个最基本的过滤方法是：**first()**, **last()** 和**eq()**，它们允许您基于其在一组元素中的位置来选择一个特定的元素。

其他过滤方法，比如**filter()**和 **not()** 允许您选取匹配或不匹配某项指定标准的元素。缩写搜索元素的范围

三个最基本的过滤方法是：first(), last() 和 eq()，它们允许您基于其在一组元素中的位置来选择一个特定的元素。

其他过滤方法，比如 filter() 和 not() 允许您选取匹配或不匹配某项指定标准的元素。
