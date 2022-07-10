# Vue

不支持 IE8 及其以下版本

## 引用的形式

1. `<script>` 便签引入,vue 会被注册成一个全局变量，UMD版本
2. npm



## 模板语法

### 文本

```vue
<span>{{message}}</span>
```

绑定的  `message` 属性发生变化，插值处的内容就会更新。通过使用 `v-once` 指令也能执行一次性的插值，后续的值更新不会影响到页面的展示。

```vue
<span v-once>{{message}}</span>
```

### html

```vue
<span v-html="rawhtml"></span>
```

`v-html` 指令可以将变量当做 HTML 渲染，而不是渲染成文本，但是会忽略解析属性中的数据绑定。

### 特性

当需要绑定是属性是 HTML 的特性时，不能使用 Mastache 语法，例如

```vue
<!-- 以下绑定是不起作用的 -->
<div id={{id}}>dynamicId</div>

<!-- 正确的做法是 -->
<div v-bind:id="dynamicId"></div>
```

### js 表达式

```vue
{{number + 1}}
{{ok ? 'y' : 'n'}}
{{message.split('').reverse().join('')}}
<div v-bind:id="'list-'+id"></div>
```

注意只能是表达式，不能是语句。不应该在模板表达式中视图访问用户定义的全局变量。

### 指令

指令是带有 `v-` 开头的特殊属性，指令的作用是当表达式的值改变时，将其产生的连带影响，响应式的作用于 DOM。

#### 参数

有一些指令可以接收参数，在指令名称之后用冒号表示，例如：`v-bind:href`，`v-on:click`，`v-bind:id`。

#### 动态参数

从 2.6.6 开始，可以使用方括号将表达式括起来作为指令的参数，如：

```vue
<a v-bind:[attrname]='url'>...</a>
```

动态参数预期会求出一个字符串，异常情况下值为 null，这个特殊的 null 值可以被用于显性的移除绑定。任何其他非字符串类型的值都会触发一个警告。

```vue
<!-- 移除绑定 -->
<a v-bind:null='df'>...</a>
<!-- 之后变为 -->
<a>...</a>
```

#### 修饰符

修饰符是以半角句号`.` 指明的特殊后缀，用于之处一个指令应当以特殊方式绑定，例如：`.prevent` 修饰符告诉 `v-on` 指令对于触发的时间调用 `event.preventDefault()`。

#### 缩写

`v-bind` 和 `v-on` 这两个指令可以缩写，

```vue
<a v-bind:href='url'>...</a>
<a :href='url'>...</a>

<a v-on:click='handle'>...</a>
<a @click='handle'>...</a>
```

## 计算属性和侦听器

### 计算属性

不用该在模板中进行大量的计算逻辑，应该使用计算属性。

```vue
var vm = new Vue({
	el:'#app',
    data:{
		message:'Hello'
    },
    computed:{
        reversedMessage:function(){// 尽量不要使用箭头函数，因为没哟 this ，引用不到数据
			return this.message.split('').reverse().join('')
        }
    }
})

<p>
    origin message:{{message}}
</p>
<p>
    computed message {{reversedMessage}}<!-- computed 属性将函数当做属性引用 -->
</p>
```

计算属性的函数被用作属性的 `getter` 函数，每当需要获取计算属性时，都会通过制定的函数尽心获取，若计算属性依赖于其他属性，当其他属性改变时，计算属性也会跟着变。

#### 计算属性缓存 VS 方法

上面的计算属性同样可以通过声明方法来达到目的：

```js
var vm = new Vue({
	el:'#app',
    data:{
		message:'Hello'
    },
    methods:{
        reversedMessage:function(){// 尽量不要使用箭头函数，因为没哟 this ，引用不到数据
			return this.message.split('').reverse().join('')
        }
    }
})
```

虽然达到的效果是一样的，然而，不同是是**计算属性是基于他们的响应式依赖进行缓存的**。只在相关响应式依赖发生改变时他们才会重新求值。这就意味着，只要被依赖的属性没有发生改变（没有依赖，也不会再次求值），多次访问 `reversedMessage` 计算属性会立即返回之前的计算结果的缓存，而不会再次执行函数。

相比之下，每当触发重新渲染时，调用方法将总会再次触发函数执行。

#### 计算属性 VS 侦听属性

侦听属性：观察和响应 Vue 实例上的数据变动。

```js
var vm = new Vue({
	el:'#app',
    data:{
		firstName: 'Foo',
        lastName: 'Bar',
        fullName: 'Foo Bar'
    },
    watch:{
        firstName: function (val) {
          this.fullName = val + ' ' + this.lastName
        },
        lastName: function (val) {
          this.fullName = this.firstName + ' ' + val
        }
    },
    // 实际上此处的侦听属性被滥用，更好的应该是
    computed: {
        fullName: function () {
          return this.firstName + ' ' + this.lastName
        }
    }
})
```

#### 计算属性的 setter

计算属性默认只有 getter ,不过在需要的时候可以提供 setter

```js
computed:{
    fullName:{
        get:function(){
            return this.firstName + ' ' + this.lastName
        },
        set:function(newName){
            let names = newName.split(' ')
            this.firstName = names[0]
            this.lastName = names[1]
        }
    }
}
```

### 侦听器

虽然计算属性在大多数情况下更合适，但有时也需要一个自定义的侦听器。这就是为什么 Vue 通过 `watch` 选项提供了一个更通用的方法，来响应数据的变化。当需要在数据变化时执行异步或开销较大的操作时，这个方式是最有用的。

```html
<div id="watch-example">
  <p>
    Ask a yes/no question:
    <input v-model="question">
  </p>
  <p>{{ answer }}</p>
</div>

<!-- 因为 AJAX 库和通用工具的生态已经相当丰富，Vue 核心代码没有重复 -->
<!-- 提供这些功能以保持精简。这也可以让你自由选择自己更熟悉的工具。 -->
<script src="https://cdn.jsdelivr.net/npm/axios@0.12.0/dist/axios.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/lodash@4.13.1/lodash.min.js"></script>
<script>
var watchExampleVM = new Vue({
  el: '#watch-example',
  data: {
    question: '',
    answer: 'I cannot give you an answer until you ask a question!'
  },
  watch: {
    // 如果 `question` 发生改变，这个函数就会运行
    question: function (newQuestion, oldQuestion) {
      this.answer = 'Waiting for you to stop typing...'
      this.debouncedGetAnswer()
    }
  },
  created: function () {
    // `_.debounce` 是一个通过 Lodash 限制操作频率的函数。
    // 在这个例子中，我们希望限制访问 yesno.wtf/api 的频率
    // AJAX 请求直到用户输入完毕才会发出。想要了解更多关于
    // `_.debounce` 函数 (及其近亲 `_.throttle`) 的知识，
    // 请参考：https://lodash.com/docs#debounce
    this.debouncedGetAnswer = _.debounce(this.getAnswer, 500)
  },
  methods: {
    getAnswer: function () {
      if (this.question.indexOf('?') === -1) {
        this.answer = 'Questions usually contain a question mark. ;-)'
        return
      }
      this.answer = 'Thinking...'
      var vm = this
      axios.get('https://yesno.wtf/api')
        .then(function (response) {
          vm.answer = _.capitalize(response.data.answer)
        })
        .catch(function (error) {
          vm.answer = 'Error! Could not reach the API. ' + error
        })
    }
  }
})
</script>
```

`vm.$watch` API

## Class 与 Style 绑定



## 条件渲染

#### v-if,v-else,v-else-if



#### 使用 key 管理可复用的元素

Vue 会复用已有的元素，而不是从头开始渲染，但当需要重复渲染时，可以给元素添加 key 属性以将组件唯一化，强制渲染



#### v-show

根据条件切换 元素的 display 属性



## 列表渲染





## 事件处理

`v-on:event`

方法

```js
methods:{
    say:function(message='hello'){
        alert(message)
    }
}
```

绑定：

```vue
<button @:click="say">say hello</button>
```

调用：

```vue
<button @:click="say('hahah')">say hahah</button>
```

特殊变量 `$event`，可以访问原始 DOM 事件

### 事件修饰符

- `.stop`
- `.prevent`
- `.capture`
- `.self`
- `.once`
- `.passive`

### 按键修饰符

```vue
<input v-on:keyup.enter="submit" />
<!-- 只有在按 Enter 时才调用 submit  -->
```

## 表单输入绑定



## 组件基础

```js
Vue.component('cmptName',{
    data:function(){// data 必须是一个函数，这样Vue 才能拥有每个组件数据的副本
        return {
            count:0
        }
    },
    template:`<button v-on:click="count++">You clicked me {{ count }} times.</button>`
})
```

组件是可复用的Vue实例，在模板中应用`<cmptName>` 。因为组件是可复用的 Vue 实例，所以组件与 `new Vuew` 接受相同是选项，例外的是 `el` 这样的根实例特有的选项。

### 组件的复用

组件的 `data` 选项必须是一个函数，这样每个实例都可以维护一份属于该组件的数据（被返回对象的独立拷贝）

```js
data:function(){
    return {
        count:0
    }
}
// 也可以这样写
data(){
    return {
        count:0;
    }
}
```

### 组件的组织

自定义的组件需要注册到 Vue 实例中，有两种注册类型：**全局注册** 和 **局部注册**。

```js
// 全局注册，注册之后可以用在任何新创建的 Vue 根实例中，必须发生在根实例被创建之前
Vue.component('my-component-name',{
    // option
})
// 第一个参数就是组件的名字
// 如果直接在 DOM 中使用一个组件，推荐遵循 W3C 的规范，字母全小写且必须包含一个连字符

// 局部注册。通过 js 对象定义组件
new Vue({
    el:'#app',
    components:{
        'component-a':ComponentA
    }
})

// 基础组件的自动化全局注册
// Globally register all base components for convenience, because they
// will be used very frequently. Components are registered using the
// PascalCased version of their file name.

import Vue from 'vue'
import upperFirst from 'lodash/upperFirst'
import camelCase from 'lodash/camelCase'

// https://webpack.js.org/guides/dependency-management/#require-context
const requireComponent = require.context(
  // Look for files in the current directory
  '.',
  // Do not look in subdirectories
  false,
  // Only include "_base-" prefixed .vue files
  /_base-[\w-]+\.vue$/
)

// For each matching file name...
requireComponent.keys().forEach((fileName) => {
  // Get the component config
  const componentConfig = requireComponent(fileName)
  // Get the PascalCase version of the component name
  const componentName = upperFirst(
    camelCase(
      fileName
        // Remove the "./_" from the beginning
        .replace(/^\.\/_/, '')
        // Remove the file extension from the end
        .replace(/\.\w+$/, '')
    )
  )
  // Globally register the component
  Vue.component(componentName, componentConfig.default || componentConfig)
})
```

### 通过 prop 向子组件传递数据

当一个值传递给一个 prop 特性时，它就变成那个组件实例的一个属性。

```js
Vue.component('blog-post',{
    props:['title','author'],
    template:'<h3>{{title}}-{{author}}</h3>'
})
```

可以使用 v-bind 来动态传递 prop

### 单个根元素

每个组件必须只能有一个根元素

可以使用 `<template>` 或者 `<div>`，包裹组件

### 监听子组件事件

Vue 实例提供了一个自定义事件的系统来解决这个问题。父组件可以像处理原生 DOM 事件一样通过 `v-on` 监听子组件实例的事件：

```vue
<blog-post v-on:my-event="evtHandler"></blog-post>
```

同时子组件可以通过内建的 `$emit(evtName)` 方法来触发一个事件。

```vue
<button @click="$emit('my-event')">
    ...
</button>
```

这样一来就不存在监听子组件事件的问题了，完全就是平常的监听事件，只不过这个事件来自其子组件。

#### 使用事件触发来向父组件传值

子组件抛出值

```vue
<button @click="$emit('my-event',0.1)">
    ...
</button>
```

父组件接受值

```vue
<blog-post v-on:my-event=evtHandle($event)></blog-post>
```

不支持越级传递，否则接收不到事件。

### 通过插槽分发内容

`<slot>` html 的占位符，可以为 innerHTML 占位，如：

```html
<ErrorBox>
	Something Wrong Happend!
</ErrorBox>
```

```js
Vue.component('ErrorBox',{
    template:`
		<div class="demo-alert-box">
          <strong>Error!</strong>
          <slot></slot>
        </div>
	`
})

// 效果
Error!Something Wrong Happend!
```

### 动态组件



# Vuex

是一个为整个应用提供共享数据存储的库。Vuex 确保了数据的变换都是通过可控的接口进行的。