# Vue

## lifecycle

![vue-lifecycle](./vue-lifecycle.png)

## Vuex

### Mutation VS Action

Mutation 用于向 Store 提交对状态的修改，而且建议通过 Mutation 而不是直接修改 state。

[Action ](https://vuex.vuejs.org/zh/guide/actions.html) 的作用类似 Mutation ，只是，Mutation 提交的是 commit ，而 Action 提交的是 Mutation，可执行异步，而 Mutation 必须同步执行。

```javascript
const store = new Vuex.Store({
  state: {
    count: 0
  },
  mutations: {
    increment (state) {
        // 提交状态修改
      state.count++
    }
  },
  actions: {
    increment (context) {
        // 提交 Mutation
      context.commit('increment')
    },
    incrementAsync({ commit },amount){
        setTimeout(() => {
            commit('increment',amount)
        },1000)
    }
  }
})

// 触发 Mutation
store.commit('increment')

// 触发 Action
store.dispatch('incrementAsync',{
    amount:10
})
```

### Module

支持将 store 分割成多个 module，每个 module 都拥有自己的 mutation、action、state、getter，或者嵌套子 module。

访问 module：**将 module 当做  `store.state` 中的一个属性来访问**。如:

```js
const moduleA = {
  namespaced:true, // 使用模块名作为 namespace 访问 state、action……
  state: { ... },
  mutations: { ... },
  actions: { 
    // state 是模块内部状态，rootState 才是 store 的 state
    // call : dispath('moduleA/incrementIfOddOnRootSum')
  	incrementIfOddOnRootSum ({ state, commit, rootState }) {
      if ((state.count + rootState.count) % 2 === 1) {
        commit('increment')
      }
    }
  },
  getters: { ... }
}

const moduleB = {
  state: { ... },
  mutations: { ... },
  actions: { ... }
}

const store = new Vuex.Store({
  modules: {
    a: moduleA,
    b: moduleB
  }
})

store.state.a // -> moduleA 的状态
store.state.b // -> moduleB 的状态
```

## Vue 自定义右键菜单

https://github.com/vmaimone/vue-context-menu

```html
<div @contextmenu.prevent="$refs.ctxMenu.open">
  ...
</div>

<context-menu id="context-menu" ref="ctxMenu" @ctx-open="onCtxOpen">
  <li @click="doSomething(...)">option 1</li>
  <li class="disabled">option 2</li>
  <li>option 3</li>
</context-menu>
```

```js
// 程序触发 & 携带数据
that.$refs.ctxMenu.open('', data);
```

## 组件通信

### 父子组件通信

父组件传值到子组件：`props`；子组件传值到父组件：子组件通过`$emit`触发事件，父组件监听得到数据

### 兄弟组件或非父子组件通信

推荐通过 Vuex 通信。普通传值很简单，基于 Vuex 也可以实现类似组件内部的响应式功能。接收数据的组件监听 store 中的特定数据变化即可：

```js
this.$store.watch(
	(state,getters) => state.xxxState,
	(newValue,oldValue) => {
		// 
	}
)
```

eventBus 通信。没有组件关系的限制，新创建一个 Vue 实例作为第三方，提供事件的触发与派发。

```js
let vue = new Vue()

// component A
vue.$emit('eventA',{payload})

// component B
vue.$on('eventA',payload => {})
```

选用什么通信方式取决于组件的关系和传输数据的大小。

父子组件通信自不必说，其他关系的组件通信需要考虑传输数据的大小，因为通信实质上是数据的复制传输，若数据量很大，复制是一件很重量级的工作，倒不如共享同一份数据（Vuex）来的轻巧。

## Vue 动态组件

Vue 通过给组件添加 `is` 属性来动态的渲染不同的组件，一般用法是：

```html
<keep-alive>
	<component :is='dynamicComponentName' ref='child'></component>
</keep-alive>
```

`<keep-alive>` 可以是组件在切换后还保留状态，而不会重新渲染。`is` 会让 `<component>` 元素随着 `dynamicComponentName` 的变化而动态渲染，该变量指明的是动态组件的 `name` 属性。`ref` 常常出现在动态组件上，一般用于父组件对不同子组件的统一操作；如不同子组件都需要处理同一份数据，但是处理方式不同，变可以使用该方法（if 判断不优雅）。

**is** 属性还可以用来避免 DOM 内模板的限制：如 

```html
<table>
  <blog-post-row></blog-post-row>
</table>
```

这个自定义组件 `<blog-post-row>` 会被作为无效的内容提升到外部，并导致最终渲染结果出错。幸好这个特殊的 `is` attribute 给了我们一个变通的办法：

```html
<table>
    <tr is="blog-post-row"></tr> 
</table>
```

需要注意的是**如果我们从以下来源使用模板的话，这条限制是\*不存在\*的**：

- 字符串 (例如：`template: '...'`)
- [单文件组件 (`.vue`)](https://cn.vuejs.org/v2/guide/single-file-components.html)
- [`<script type="text/x-template">`](https://cn.vuejs.org/v2/guide/components-edge-cases.html#X-Templates)

## 例子

以一个简单的日志记录为例，主界面有 tab 页，选择不同级别的日志进行记录，底部展示当前所有日志记录。

```text
<button @click='changeLog(logBtn.type)' v-for='logBtn in logList'>{{logBtn.text}}</button>
<keep-alive>
	<component :is='currentLogger' ref='log'></component>
</keep-alive>

logList:[{
	type: 0,
	text:'Error',
	componentName: 'ErrorLogger'
},……]

changeLog(type){
	this.currentLogger = this.logList[type].componentName
}

addLog(){
	this.logTable.push(this.$refs.log.add())
}
```

Error logger

```text
name: ErrorLogger

add(){
	return `<span style='color: red'>${logMsg}</span>`
}
```

Warning logger、Success logger……

# CSS

css 选择器上的`/deep/` 修饰：https://stackoverflow.com/questions/25609678/what-do-deep-and-shadow-mean-in-a-css-selector

HTML5 Web Components offer full encapsulation of CSS styles.

This means that:

- styles defined within a component cannot leak out and effect the rest of the page
- styles defined at the page level do not modify the component's own styles

 

- 修改第三方组件的样式时，应该注意影响范围，最好在第三方组件外围包裹元素，再使用 CSS 子元素选择器选择第三方组件。

## CSS 如何处理非法样式



跳过不处理。CSS 在快速发展，若不能兼容“错误”样式，而是报错，整个罢工不渲染，肯定是不长久的。对于不支持新特性的老式浏览器来说，这一点及其重要。这样，就有了一种处理浏览器兼容的方式——给同一个块写两个或更多的样式，依照浏览器支持程度顺序写，如：

```css
.block {
    width:300px,
    width:calc(90% - 30px) /*浏览器支持 calc 函数，则覆盖 width:300px,若不支持则位置原先设置*/
}
```



## [层叠、优先级与继承](https://developer.mozilla.org/zh-CN/docs/Learn/CSS/Building_blocks/Cascade_and_inheritance)

理解**层叠**（cascade）是理解 CSS 的关键。当同一块被设置的不同的样式，就发生了冲突，该使用哪个设置，就是层叠体现的地方——**后来者会覆盖前者，优先级高的会覆盖优先级低的**。

对于如何定义**优先级**，有一个总的原则——具体。如：`元素选择器 < 类选择器 < id 选择器 <  元素内 style`。随着优先级的提高，选择器所表示的范围越来越小，选择出的元素越来越具体，设置的对应规则自然应该被优先展示。使用 `!important` 可以提升优先级。

一些设置在父元素上的样式会被子元素**继承**，有些则不能。如给 `body` 元素设置字体颜色，则整个页面的字体颜色都会随之改变（除非设置更加具体的规则进行覆盖），这就是所有页面元素都继承了 `body` 这个父元素的样式。像宽度这种属性则不会默认继承父类，否则就很难达到想要的效果了，**哪些属性属于默认继承很大程度上是由常识决定的**。

以上这些概念一起控制 CSS 规则应该作用于哪些元素。

### 理解继承

CSS 提供了 `inherit`,`initial`,`unset` 属性值来控制继承，每个 CSS 属性都支持这些值。

- `inherit`:开启继承
- `initial`:设置属性值和浏览器默认样式相同。如果浏览器默认样式中未设置且该属性是自然继承的，那么会设置为 `inherit` 
- `unset`:将属性重置为自然值，也就是如果属性是自然继承那么就是 `inherit`，否则和 `initial`一样

## [盒模型](https://developer.mozilla.org/zh-CN/docs/Learn/CSS/Building_blocks/The_box_model)



# Echarts

## 一个性能问题

Vue + Echarts 可能的一个性能问题是将 ECharts 对象挂载到 Vue 的 data 中。ECharts 实例是一个巨大的对象，若挂载到 Vue 的 data 中，会严重占用内存，影响图表的渲染速度。

```js
// somewhere beyond vue object
var echartInstance = null
// inside vue object
echartInstance = echart.init(document.getElementById('#id'));
echartInstance.setOption(options);
this.$once('hook:beforeDestroy', function() {
    echart.dispose(echartInstance);
});
```

ref:https://github.com/apache/incubator-echarts/issues/7234



# 回调

当需要在回调函数中使用当前作用域中的变量时，会出现访问不到的情况，如：

```js
this.unwatch = this.$store.watch(
    (state, getters) => state.showXAxis,
    (newV, oldV) => {
        // echartInstance 访问不到
        echartInstance.setOption({
            xAxis: {
                show: newV
            }
        });
    }
);
```

更改：

```js
const echartInstanceCopy = echartInstance;
this.unwatch = this.$store.watch(
    (state, getters) => state.showXAxis,
    (newV, oldV) => {
        ....
```

同样的还有` this ` 的问题，不光箭头函数中访问不到正确的 `this` ，在回调函数中也不能正确的访问 `this` 指向的 vue 实例，也需要提前做局部缓存。（在基于 `webpack` 的应用中，在箭头函数中可以正确的访问到 `this`,那是因为 `webpack`在编译时自动将箭头函数作用域中的 `this` 替换为外围的 `this`）

```js
this.showChart(10);
const that = this;
echartInstance.on('click', function(e) {
    // 访问不到正确的 this 
    that.$emit('point-event', e);
});

// webpack 优化
echartInstance.on('click', function(e) {
    // 访问不到正确的 this 
    this.$emit('point-event', e);
});
// webpack 优化为
let _this = this
echartInstance.on('click', function(e) {
    _this.$emit('point-event', e);
});
```

箭头函数中的 `this` ：箭头函数没有自己的 `this`，但会引用父级的 `this`。所以在箭头函数中还是可以使用 `this` 的，只要是访问的调用者确是当前父级即可。

# java8 stream api

`groupingby` 可以对集合按指定的键进行分组，如集合：

```js
[
    {name:'jack',age:12,addr:'shanghai',country:'China'},
    {name:'rose',age:24,addr:'shanghai',country:'UK'},
    {name:'robin',age:12,addr:'beijing',country:'China'},
    {name:'pony',age:12,addr:'guangzhou',country:'Ca'}
] = User
```

```java
// 单条件分组
Map<String,List<User>> collect = list.stream().collect(Collectors.groupingby(User::getAddr()));

// 双条件分组
Map<Pair<String,String>,List<User>> collect = list.stream().collect(Collectors.groupingby(u -> Pair.of(u.getAddr(),u.getCountry()));
                                                                    
// 多条件分组
Map<String,List<User>> collect = list.stream().collect(Collectors.groupingby(u -> getGroupbyKey(u)));
             
// 任意键组合
String getGroupbyKey(User u){
    return String.format("%s_%s_%s",u.getAddr(),u.getCountry(),u.getName());
}
```

[Oracle Java Tutorial](https://docs.oracle.com/javase/tutorial/index.html)



# 文档生成工具

https://docsify.js.org/#/    https://github.com/docsifyjs/docsify

https://www.showdoc.cc/

https://github.com/phachon/mm-wiki



# Js Promise

参考之前的一篇[博文](https://jianxinliu.github.io/post/learning_note/f2e/ES6_Promise.html)。

```javascript
async function getAll(){
    let promises = listColumn.map(col => {
        return (col => {
            return new Promise((resolve, reject) => {
                api.query(`SELECT DISTINCT ${field} FROM ${tableName} WHERE ${field} IS NOT NULL`)
                    .then(res => {
                        if (res.data) {
                            resolve({[field]: res.data})
                        } else {
                            reject('error')
                        }
                    })
            })
        })(col)
    })
    let results = await Promise.all(promises)
    ......
}
```

## 微任务队列

```javascript
let promise = Promise.resolve()
promise.then(res => console.log('promise done!'))
consoloe.log('code finished!')
```

上面这段代码的输出是这样的：

```
code finished!
promise done!
```

为什么立即 resolve 的代码执行依然会排在直接输出语句之后呢？因为 `Promise` 中，不论是 `then`,`catch`或 `finally` 语句块中的内容，都不会立即执行，而是会加入微任务队列中，直到 js 引擎没有其它任务在运行时（**宏任务队列**），才会从队列中取出任务执行。在上面的例子中执行到第 2 句时， `console.log('promise done!')`被放入微任务队列中，js 引擎接着执行第 2 句，之后 js 引擎没有任务执行了，才从微任务队列中取出 `console.log('promise done!')` 执行。若要保证执行的顺序符合“直觉”，可将需要被顺序执行的代码依次使用 `then` 去调用，这样所有任务都会被加入队列，依次执行。

[ref-微任务队列](https://zh.javascript.info/microtask-queue)、[事件循环与宏任务队列](https://zh.javascript.info/event-loop)

## `async` & `await`

async/await 是以更舒适的方式使用 promise 的一种特殊语法，同时它也非常易于理解和使用。

async 用于修饰函数，保证该函数一定返回一个 `Promise`，即使最终函数没有返回 `Promise`，处理结果也会被封装进一个 resolved 的`Promise` 中返回。所以：

```js
async function foo(){
    return 1
}
async function bar(){
    return Promise().resolve(1)
}

foo().then(log) // 1
bar().then(log) // 1
```



# 函数错误处理

参考 go 语言函数的错误处理

```go
ret,error := someFn()
if(error){
    // 处理错误
} else {
    // 处理返回结果
}
```

```javascript
function foo(fields){
    if(fields.length < 1){
        return {
            ret:null,
            error:'fields empty!'
        }
    }
    // 正常处理流程
    
    // 最终返回
    let something = ''
    return {
        ret:something,
        error:null
    }
}

// 函数调用
let {ret,error} = foo(['a'])
if(error){
    alert(error)
} else {
    // handle result
}
// 或者对于输出返回值类型的函数，还可以更简便地处理
console.log(ret || error)
```

go 语言函数的返回值支持自定义变量名，即`let {ret,error} = foo(['a'])`中的 `ret` 可以自定义，即函数返回两个值，第一个是返回值，第二个是错误信息，但 js 不行，必须获取函数返回的那个变量，尤其是使用解构。

```js
// js 内置的错误对象
throw new Error('ErrMsg')

// 自定义错误对象
throw 'custom error'
```



# Web 端透视表生成的关键点

Excel 的透视表功能强大，支持多个维度对原始数据进行透视，从而获得更精确的信息。在 Web 实现该功能，看似及其困难，实际上关键技术还是常见的那些。有两种方案，一种是纯前端生成，一种是前端生成 SQL 后端得出透视表。纯前端生成的方案除了在数据量巨大时可能出现瓶颈外，其余都堪称完美。

## 透视表生成原理

对一个表格的数据生成透视表，实际上可以看成是 SQL 执行 `group by` 的操作，无论透视表是加行维度还是列维度，在 `group by` 中没有区别，都是增加 `group by` 的维度，不同的只是在展示上有区别。

## 纯前端

不使用 SQL，前端也可以进行 `group by`，知名 JavaScript 工具类库 `lodash`有针对数组的 `groupby` 函数，但该函数只支持一个维度的分组，不像 SQL 的 `group by` 后可接任意多字段。

```javascript
var _ = require('lodash')

Array.prototype.groupBy = function(fields) {
  return groupBy(this, fields, 0)
}

// 通过递归的方式实现多级 group by
function groupBy(dataSource, fields, i) {
  let ret = _.groupBy(dataSource, d => d[fields[i]])
  if (i < fields.length - 1) {
    i += 1
    for (let p in ret) {
      ret[p] = groupBy(ret[p], fields, i)
    }
  }
  return ret
}
```

如此 groupby 出来的数据，结构和透视表完全一样，之后其中的统计列需要再根据 groupby 出的组数据进行计算。这样的好处是生成透视表及其方便，还同时保留了原始数据，有利于后期对生成的统计数据快速钻取原始数据。

## 前端拼接 SQL

当数据量巨大，可以由前端根据操作生成 SQL ，再由后端执行 SQL 生成透视表的数据。实际上透视表也可以由 SQL 生成。如有学生数据，行维度选择国家和性别，列维度选择班级和城市，统计其 IQ 和 体重的均值，则生成其透视表数据的 SQL 是这样的。注：假设班级只有 `1` 和 `2`，城市只有 `beijing` 和 `shangha` 。

```sql
select country,
	sex,
	case when class = 1 and city = 'beijing' then avg(iq) end "1_beijing_IQ",
	case when class = 1 and city = 'shangha' then avg(iq) end "1_shangha_IQ",
	case when class = 2 and city = 'beijing' then avg(iq) end "2_beijing_IQ",
	case when class = 2 and city = 'shangha' then avg(iq) end "2_shangha_IQ",
	case when class = 2 and city = 'beijing' then avg(weight) end "2_beijing_weight",
	case when class = 1 and city = 'shangha' then avg(weight) end "1_shangha_weight",
	case when class = 1 and city = 'beijing' then avg(weight) end "1_beijing_weight",
	case when class = 2 and city = 'shangha' then avg(weight) end "2_shangha_weight"
from student
group by country, sex, class, city
```

从可以看到 `group by` 子句中可以看出，无论行维度还是列维度，都加入 group by 即可，只是展现不同而已。

由于这样的 SQL 生成的结果是会出现空值的，故再嵌套一层 `group by` 即可压缩结果集。

```sql
select country,
		sex,
		max("1_beijing_IQ") as "1_beijing_IQ" ,
		max("1_shangha_IQ") as "1_shangha_IQ" ,
		max("2_beijing_IQ") as "2_beijing_IQ" ,
		max("2_shangha_IQ") as "2_shangha_IQ" ,
		max("2_beijing_weight") as "2_beijing_weight",
		max("1_shangha_weight") as "1_shangha_weight",
		max("1_beijing_weight") as "1_beijing_weight",
		max("2_shangha_weight") as "2_shangha_weight"
from(
		select country,
				sex,
				case when class = 1 and city = 'beijing' then avg(iq) end "1_beijing_IQ",
				case when class = 1 and city = 'shangha' then avg(iq) end "1_shangha_IQ",
				case when class = 2 and city = 'beijing' then avg(iq) end "2_beijing_IQ",
				case when class = 2 and city = 'shangha' then avg(iq) end "2_shangha_IQ",
				case when class = 2 and city = 'beijing' then avg(weight) end "2_beijing_weight",
				case when class = 1 and city = 'shangha' then avg(weight) end "1_shangha_weight",
				case when class = 1 and city = 'beijing' then avg(weight) end "1_beijing_weight",
				case when class = 2 and city = 'shangha' then avg(weight) end "2_shangha_weight"
		from student
		group by country, sex, class, city
	) d
group by country,sex
order by country,sex
```

现在 SQL 是有了，但如何根据界面的操作来生成这个 SQL 呢？注意到内部 SQL 的 `case when` 的个数，是由列维度和统计维度的值的进行全排列的个数决定的。即班级有值：1,2；城市有值：beijing,shangha；统计 IQ  和 weight 。则全排列为

```
1 beijing iq
1 beijing weight
1 shangha iq
1 shangha weight
2 beijing iq
2 beijing weight
2 shangha iq
2 shangha weight
```

有了这个全排列的结果，则生成 SQL 就没什么问题了。全排列的函数参照[博客](https://blog.csdn.net/djcxym/article/details/79359057)，

```javascript
let indexes = []
function all_permutations(martix) {
  let res = []
  allPermutations(martix, res, 0)
  indexes = []
  return res
}
// 全排列
function allPermutations(martix, res, level) {
  let oneCase = []
  if (level < martix.length) {
    for (let i = 0; i < martix[level].length; i++) {
      indexes[level] = i
      allPermutations(martix, res, level + 1)
    }
  } else {
    for (let i = 0; i < martix.length; i++) {
      oneCase.push(martix[i][indexes[i]])
    }
    res.push(oneCase)
  }
}
```

全排列，也称笛卡尔积问题，简便方案：[两行代码](https://stackoverflow.com/questions/12303989/cartesian-product-of-multiple-arrays-in-javascript),[电商 SKU](https://juejin.im/post/5ee838cc6fb9a047ea45ef48)

```js
const f = (a, b) => [].concat(...a.map(d => b.map(e => [].concat(d, e)))); // 将问题简化为两个元素
const cartesian = (a, b, ...c) => (b ? cartesian(f(a, b), ...c) : a); // 递归处理复杂的情况
let output = cartesian([1,2],[10,20],[100,200,300]);
output:
[ [ 1, 10, 100 ],
  [ 1, 10, 200 ],
  [ 1, 10, 300 ],
  [ 1, 20, 100 ],
  [ 1, 20, 200 ],
  [ 1, 20, 300 ],
  [ 2, 10, 100 ],
  [ 2, 10, 200 ],
  [ 2, 10, 300 ],
  [ 2, 20, 100 ],
  [ 2, 20, 200 ],
  [ 2, 20, 300 ] ]
```



## 多级表头 HTML 的生成

基于 elementUI 的多级表头，需要嵌套 `<el-table-column>` 元素，这对于动态生成的多级表头极不友好。最开始想到的是使用 jsx 来生成嵌套的元素，但在 vue 里动态添加元素的打开方式似乎不对，故没有成功，后参看此[博客](https://blog.csdn.net/liub37/article/details/82906141)，才明白不用写代码也可以实现嵌套组件的生成。

该博客的思路是新建一个组件，由这个组件自己负责嵌套生成 `<el-table-column>` 元素。

MyColumn.vue 组件

```vue
<template>
  <el-table-column :prop="col.prop" :label="col.label" align="left">
    <template v-if="col.children">
      <my-column v-for="(item, index) in col.children" :key="index" :col="item"></my-column>
    </template>
  </el-table-column>
</template>

<script>
import Vue from 'vue'
import { TableColumn } from 'element-ui'
Vue.use(TableColumn)
export default {
  name: 'MyColumn',
  props: {
    col: {
      type: Object
    }
  }
}
</script>
<style scoped></style>
```

使用该组件需要把列信息的结构更改为：

```json
[
    {
        prop: 'date',
        label: '日期'
    },
    {
        label: '配送信息',
        children: [
            {
                prop: 'name',
                label: '姓名'
            },
            {
                label: '地址',
                children: [
                    {
                        prop: 'province',
                        label: '省份'
                    },
                    {
                        prop: 'city',
                        label: '市区'
                    },
                    {
                        prop: 'address',
                        label: '地址'
                    }
                ]
            }
        ]
    }
]
```

使用时只需要把更改结构后的数据循环后传给myColumn 组件即可

```html
<my-column v-for="(item, index) in tableHeadSource" :key="index" :col="item"></my-column>
```

#  JavaScript 扩展函数

利用对象的 `prototype` 即可为对象扩展函数，如给 Array 扩展一个获取数组最后元素的函数

```js
Array.prototype.last = function(){
    return this[this.length - 1]
}
```

- 不能使用箭头函数定义。因为箭头函数中没有 `this` ，而扩展函数还需要 `this` 来指向调用者
- 使用 `this` 指向调用者，如 `[1,2,3].last()`,函数中的 `this` 便可以指向 `[1,2,3]` 这个数据
- 这样扩展之后，相当于数组多了一个函数，全局可用。

**重要：** 不要随意污染全局空间，危害参见<a href='#prototype'>JS Clean Code</a>

# JavaScript 数据操作

函数式编程和命令式编程的不同。

[函数式编程简介](https://mp.weixin.qq.com/s?__biz=MjM5ODQ2MDIyMA==&mid=402307374&idx=1&sn=2ff35dc5bcadab0bbeae626f48f4e18e#rd) [抽象的能力](https://zhuanlan.zhihu.com/p/20617201) 

遍历，过滤，查找，分组，排序，映射，归约…… 

```js
let classNo = [1,2,3,4]
let students = [{name:'jack',age:12,addr:'beijing'},
              {name:'rose',age:12,addr:'beijing'},
              {name:'mary',age:25,addr:'sahnghai'},
              {name:'pony',age:24,addr:'sahnghai'},
              {name:'robin',age:24,addr:'gaungzhou'}]

classNo.forEach(v => console.log(v)) // => 1 2 3 4
classNo.map(v => console.log(v))     // => 1 2 3 4
students.map(v => console.log(v.name))  // => ['jack','rose','mary','pony','robin']

// 实际上 map 的主要作用不是遍历，而是遍历的过程中，对所遍历值的操作
let student_age = students.map(v => v.age) // => [12,12,25,24,24]
let student_age_formatted = students.map(v => v.age + '岁') => ['12岁','12岁','25岁','24岁','24岁']

let student_from_beijing = students.filter(v => v.addr === 'beijing') // => find jack & rose

let rose = students.find(v => v.name === 'rose') // find rose
let roseIndex = students.findIndex(v => v.name === 'rose') // 1

let range_by_age_asc = students.sort((a,b) => a.age - b.age)
let range_by_age_desc = students.sort((a,b) => b.age - a.age)

let avg_age = students.reduce((a,b) => a + b.age,0) / students.length // => 19.4
let min = (...args) => 
	Math.min(...args.map(v => Array.isArray(v) ? v.map(f => Number(f)) : Number(v)).flat(Infinity))
min([1,2,3],6,[4,67,78],9)

// array deep copy
let studentCopy1 = Object.assign([],student) // 若 copy 的对象的属性都是原生类型，则可深拷贝。若有引用，则 copy 的是引用。
let studentCopy2 = JSON.parse(JSON.stringify(student))
let studentCopy3 = student.map(stu => Object.assign({},stu))
let studentCopy4 = student.map(stu => ({...stu}))
```

# Java 8 Stream api 数据操作

```java
import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.Arrays;
import java.util.DoubleSummaryStatistics;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;
import java.util.stream.Stream;

/**
 * @author jianxinliu
 * @date 2020/07/11 11:43
 * @description
 */
public class StreamTest {
    static Student defaultStu = new Student("default", 0, "xxx", 100000, "");
    static Student[] students = new Student[]{
            new Student("jack", 13, "shanghai", 123, ""),
            new Student("rose", 24, "shanghai", 124, ""),
            new Student("pony", 34, "guangzhou", 135, ""),
            new Student("robin", 35, "beijing", 143, "")
    };

    public static void main(String[] args) {
        List<Student> stuList = Arrays.asList(students);
        // just for each
        stuList.forEach(System.out::println);
        // change list properties
        List<String> stus = stuList.stream().map(stu -> stu.note = "hello").collect(Collectors.toList());
        stus.forEach(System.out::println);

        // get cities
        stuList.stream().map(Student::getAddr).distinct().collect(Collectors.toList()).forEach(System.out::println);

        double totalAge = stuList.stream().map(Student::getAge).reduce((a, b) -> a + b).orElse(0).doubleValue();
        System.out.println(totalAge);

        double avgIQ = stuList.stream().filter(v -> v.getAge() > 13).mapToDouble(Student::getIq).average().orElse(0);
        System.out.println(avgIQ);

        boolean allAdult = stuList.stream().allMatch(v -> v.getAge() > 18);
        System.out.println(allAdult);

        DoubleSummaryStatistics iqStatistics = stuList.stream().filter(v -> v.getAge() > 18).mapToDouble(Student::getIq).summaryStatistics();
        System.out.println(iqStatistics.getAverage());

        System.out.println(getStudentByName("edsion").orElse(defaultStu));
        System.out.println(getStudentByName("edsion").orElseThrow(IllegalArgumentException::new));
    }

    static Optional<Student> getStudentByName(String name) {
        List<Student> stuList = Arrays.asList(students);
        Stream<Student> limit = stuList.stream().filter(v -> v.getName() == name).limit(1);
        if (limit.findAny().isPresent()) {
            return Optional.of(limit.findAny().get());
        } else {
            return Optional.empty();
        }
    }
}

@Data
@AllArgsConstructor
class Student {
    String name;
    Integer age;
    String addr;
    Integer iq;
    String note;
}
```

# js 文件读入写出

HTML 5 的 File API 和 Blob 对象给 Web 页面提供了读写本地文件的能力。

读取本地文件内容

```js
export function readFile(callback) {
  let inputEle = document.createElement('input')
  inputEle.setAttribute('id', 'tempInput')
  inputEle.setAttribute('type', 'file')
  inputEle.setAttribute('accept', 'text/txt')
  inputEle.setAttribute("style", "display: none")
  const handleValue = () => {
    let file = inputEle.files[0]
    let blob = new Blob([file], {type: "text/plain;charset=utf-8"})
    blob.text().then(text => {
	  // text 即文件内容的文本格式
      callback(text)
      document.body.removeChild(inputEle)
    })
  }
  inputEle.addEventListener('change', handleValue, true)
  document.body.appendChild(inputEle)
  inputEle.click()
}
```

写入文件使用 File-saver 依赖，更方便全面的操作，实际上也是使用 File API  和 Blob 对象。

```js
export function buildAndSave(fileName) {
  let content = {}
  let settingJson = JSON.stringify(content)
  const blob = new Blob([NOTE, settingJson], {type: "text/plain;charset=utf-8"})
  FileSaver.saveAs(blob, fileName + '.txt')
}
```

# 计算

## 程序计算

数据可以被计算，程序也可以被计算。

 使用 `eval` 计算出表达式，代入程序执行。 

例如，图表上有五根线 `y = C`，C 的取值及其大小关系为：`USL > UCL > Target > LCL > lSL`，五根线中可能有不存在的线。设一系列点 P 中满足 `P > USL || P < LSL` 的点数称为 `OOS(out of specfication)`，满足 `P > UCL || P < LCL`的点数称为 `OOC(out of control)`。求一系列点 arrP 的 OOS，OOC。

本题麻烦的点在于五根线的存在性不确定，判断起来很繁琐。以求 OOS 为例

1. 若 USL 不存在，则过滤的表达式为 `P < LSL`
2. 若 LSL 不存在，则过滤的表达式为 `P > USL`
3. 若 USL、LSL 同时不存在，则不需计算 OOS
4. 若 USL、LSL 同时存在，则过滤的表达式为`P > USL || P < LSL`

若是基于 `if else` ，那肯定写出一大堆，还不好理解。若是转换下思路，将表达式看做是字符串，则基于以上逻辑，拼接字符串还是会稍微简单一些的。

```js
const conditionExpression = (up,low) => {
    let cond = [];
    // 设 d 为 filter 函数中的迭代变量名，直接写死会简便一些，也可通过参数传入
    up && cond.push(`d > ${up}`); 
    low && cond.push(`d < ${low}`);
    (!up && !low) && cond.push('false')
    return cond.join(' || ')
}

// call
let oos = arrP.filter(d => eval(conditionExpression(usl,lsl))).length
let ooc = arrP.filter(d => eval(conditionExpression(ucl,lcl))).length
```

## 属性计算

统一即简洁，统一即可计算。 

# js clean code

https://github.com/ryanmcdermott/clean-code-javascript

## 变量

### Use meaningful and pronounceable variable names

**Bad:**

```js
const yyyymmdstr = moment().format("YYYY/MM/DD");
```

**Good:**

```js
const currentDate = moment().format("YYYY/MM/DD");
```

### 使用可自解释的变量名

**Bad:**

```js
const address = "One Infinite Loop, Cupertino 95014";
const cityZipCodeRegex = /^[^,\\]+[,\\\s]+(.+?)\s*(\d{5})?$/;
saveCityZipCode(
  address.match(cityZipCodeRegex)[1],
  address.match(cityZipCodeRegex)[2]
);
```

**Good:**

```js
const address = "One Infinite Loop, Cupertino 95014";
const cityZipCodeRegex = /^[^,\\]+[,\\\s]+(.+?)\s*(\d{5})?$/;
const [_, city, zipCode] = address.match(cityZipCodeRegex) || [];
saveCityZipCode(city, zipCode);
```

### Explicit is better than implicit.

**Bad:**

```js
const locations = ["Austin", "New York", "San Francisco"];
locations.forEach(l => {
  doStuff();
  doSomeOtherStuff();
  // ...
  // ...
  // ...
  // Wait, what is `l` for again?
  dispatch(l);
});
```

**Good:**

```js
const locations = ["Austin", "New York", "San Francisco"];
locations.forEach(location => {
  doStuff();
  doSomeOtherStuff();
  // ...
  // ...
  // ...
  dispatch(location);
});
```

### Don't add unneeded context

If your class/object name tells you something, don't repeat that in your variable name.

**Bad:**

```js
const Car = {
  carMake: "Honda",
  carModel: "Accord",
  carColor: "Blue"
};

function paintCar(car) {
  car.carColor = "Red";
}
```

**Good:**

```js
const Car = {
  make: "Honda",
  model: "Accord",
  color: "Blue"
};

function paintCar(car) {
  car.color = "Red";
}
```

### Use default arguments instead of short circuiting or conditionals

Default arguments are often cleaner than short circuiting. Be aware that if you use them, your function will only provide default values for `undefined` arguments. Other "falsy" values such as `''`, `""`, `false`, `null`, `0`, and `NaN`, **will not be replaced by a default value**.

**Bad:**

```js
function createMicrobrewery(name) {
  const breweryName = name || "Hipster Brew Co.";
  // ...
}
```

**Good:**

```JS
function createMicrobrewery(name = "Hipster Brew Co.") {
  // ...
}
```

## 函数

### Function arguments (2 or fewer ideally)

**Bad:**

```js
function createMenu(title, body, buttonText, cancellable) {
  // ...
}

createMenu("Foo", "Bar", "Baz", true);
```

**Good:**

```js
function createMenu({ title, body, buttonText, cancellable }) {
  // ...
}
// 多了一个解释参数意义的机会
createMenu({
  title: "Foo",
  body: "Bar",
  buttonText: "Baz",
  cancellable: true
});
```

### Function names should say what they do

**Bad:**

```js
function addToDate(date, month) {
  // ...
}

const date = new Date();

// It's hard to tell from the function name what is added
addToDate(date, 1);
```

**Good:**

```js
function addMonthToDate(month, date) {
  // ...
}

const date = new Date();
addMonthToDate(1, date);
```

### Avoid Side Effects

避免函数带有副作用。因为 JavaScript 函数传参时，基本类型传值，对象类型传引用，故最好是只操作传入数据的副本，函数返回副本，对参数无影响，这样其他地方再次使用该参数时，不会出现意想不到的情况。

**Bad:**

```js
// Global variable referenced by following function.
// If we had another function that used this name, now it'd be an array and it could break it.
let name = "Ryan McDermott";

function splitIntoFirstAndLastName() {
  name = name.split(" ");
}

splitIntoFirstAndLastName();

console.log(name); // ['Ryan', 'McDermott'];
```

**Good:**

```js
function splitIntoFirstAndLastName(name) {
  return name.split(" ");
}

const name = "Ryan McDermott";
const newName = splitIntoFirstAndLastName(name);

console.log(name); // 'Ryan McDermott';
console.log(newName); // ['Ryan', 'McDermott'];
```

**Except:**基于 Vuex 的类责任链模式——就是靠副作用才使得代码简洁、流程清晰。

### Don't write to global functions

<span name='prototype'>不要给已有内置对象添加属性（方法），这样会影响到全局，造成污染。</span>

正是因为 JavaScript 的便利，才不可能预防他人（库）定义和你自定义函数同名的函数，一旦这种情况发生，前者的自定义函数将失效，整个系统将发生不可预估的问题。

Polluting globals is a bad practice in JavaScript because you could clash with another library and the user of your API would be none-the-wiser until they get an exception in production. Let's think about an example: what if you wanted to extend JavaScript's native Array method to have a `diff` method that could show the difference between two arrays? You could write your new function to the `Array.prototype`, but it could clash with another library that tried to do the same thing. What if that other library was just using `diff` to find the difference between the first and last elements of an array? This is why it would be much better to just use ES2015/ES6 classes and simply extend the `Array` global.

**Bad:**

```js
Array.prototype.diff = function diff(comparisonArray) {
  const hash = new Set(comparisonArray);
  return this.filter(elem => !hash.has(elem));
};
```

**Good:**

```js
class SuperArray extends Array {
  diff(comparisonArray) {
    const hash = new Set(comparisonArray);
    return this.filter(elem => !hash.has(elem));
  }
}
```

### 尽量使用函数式而不是命令式

**Bad:**

```js
const programmerOutput = [{
    name: "Uncle Bobby",
    linesOfCode: 500
  },{
    name: "Suzie Q",
    linesOfCode: 1500
  }];

let totalOutput = 0;

for (let i = 0; i < programmerOutput.length; i++) {
  totalOutput += programmerOutput[i].linesOfCode;
}
```

**Good:**

```js
const programmerOutput = [{
    name: "Uncle Bobby",
    linesOfCode: 500
  },{
    name: "Suzie Q",
    linesOfCode: 1500
  }];

const totalOutput = programmerOutput.reduce((totalLines, output) => totalLines + output.linesOfCode, 0);
```

### 具名条件表达式

**Bad:**

```js
if (fsm.state === "fetching" && isEmpty(listNode)) {
  // ...
}
```

**Good:**

```js
const shouldShowSpinner = (fsm, listNode) => fsm.state === "fetching" && isEmpty(listNode)

if (shouldShowSpinner(fsmInstance, listNodeInstance)) {
  // ...
}
```

### 尽量避免‘非’条件

这只会让人陷入逻辑黑洞

**Bad:**

```js
function isDOMNodeNotPresent(node) {
  // ...
}

if (!isDOMNodeNotPresent(node)) {
  // ...
}
```

**Good:**

```js
function isDOMNodePresent(node) {
  // ...
}

if (isDOMNodePresent(node)) {
  // ...
}
```

## 格式

### Function callers and callees should be close

If a function calls another, keep those functions vertically close in the source file. Ideally, keep the caller right above the callee. We tend to read code from top-to-bottom, like a newspaper. Because of this, make your code read that way.

**Bad:**

```js
class PerformanceReview {
  constructor(employee) {
    this.employee = employee;
  }

  lookupPeers() {
    return db.lookup(this.employee, "peers");
  }

  lookupManager() {
    return db.lookup(this.employee, "manager");
  }

  getPeerReviews() {
    const peers = this.lookupPeers();
    // ...
  }

  perfReview() {
    this.getPeerReviews();
    this.getManagerReview();
    this.getSelfReview();
  }

  getManagerReview() {
    const manager = this.lookupManager();
  }

  getSelfReview() {
    // ...
  }
}

const review = new PerformanceReview(employee);
review.perfReview();
```

**Good:**

```js
class PerformanceReview {
  constructor(employee) {
    this.employee = employee;
  }

  // 总-分 结构
  perfReview() {
    this.getPeerReviews();
    this.getManagerReview();
    this.getSelfReview();
  }

  getPeerReviews() {
    const peers = this.lookupPeers();
    // ...
  }

  lookupPeers() {
    return db.lookup(this.employee, "peers");
  }

  getManagerReview() {
    const manager = this.lookupManager();
  }

  lookupManager() {
    return db.lookup(this.employee, "manager");
  }

  getSelfReview() {
    // ...
  }
}

const review = new PerformanceReview(employee);
review.perfReview();
```

## 注释

Comments are an apology, not a requirement. Good code *mostly* documents itself.

**Bad:**

```js
function hashIt(data) {
  // The hash
  let hash = 0;

  // Length of string
  const length = data.length;

  // Loop through every character in data
  for (let i = 0; i < length; i++) {
    // Get character code.
    const char = data.charCodeAt(i);
    // Make the hash
    hash = (hash << 5) - hash + char;
    // Convert to 32-bit integer
    hash &= hash;
  }
}
```

**Good:**

```js
function hashIt(data) {
  let hash = 0;
  const length = data.length;

  for (let i = 0; i < length; i++) {
    const char = data.charCodeAt(i);
    hash = (hash << 5) - hash + char;

    // Convert to 32-bit integer
    hash &= hash;
  }
}
```

# SQL

`row_number() over([partion by xx] order by xx) as rowNum` 该列是 SQL 执行后对结果集的每一行进行排序后生成的编号，可以用来做分页的依据。若加 `partion by xx` 则会依据指定字段分组，每组单独生成行号。

`case when a=1 and b=2 then avg(age) when a=2 and b=3 then ... end as avg_age`

`group by` 按字段分组，相同的在一组，并对各组执行聚合操作。透视表就是基于分组实现的，在分组的基础上，选择聚合操作。统计中的分组也是一样的原理。

# js 并发模型与事件循环

[ref](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/EventLoop)

JavaScript有一个基于**事件循环**的并发模型，事件循环负责执行代码、收集和处理事件以及执行队列中的子任务。这个模型与其它语言中的模型截然不同，比如 C 和 Java。

js 引擎是单线程模型，故一个函数在执行时，不会被抢占，只有在运行完之后才会运行其他任务（**执行至完成**）。该模型的缺点就是当一个任务运行时间过长，则会影响到其他程序的执行，如 Web 应用程序就无法响应用户的交互，会出现页面卡死的现象。

在浏览器中，每当有一个事件发生，且有一个事件监听器绑定在该事件上，该事件就会被加入消息队列。函数 `setTimeout` 可以将一个函数推迟一段时间执行，原理是当调用 `setTimeout` 时，传入的第一个参数（函数）将被加入消息队列等待执行，理想情况下，队列为空，则到了指定时间后，加入队列的消息会在指定的时间间隔后执行。非理想情况下，可能在消息入队之前，消息队列已经排有耗时远超指定的时间间隔，则该消息不会在指定的之间后执行，而是会在队列执行到该消息时执行。也就是：**`setTimeout` 的第二个参数仅仅表示消息延迟执行的最小时间间隔。** 同样的，`setTimeout(fn, 0)` 并不能立即执行`fn`。

也正是因为 js 引擎采用事件循环模型和消息队列，故可以实现“**永不阻塞**”。如一个 Web 应用在等待 XHR 的返回时，依然可以处理其他如用户输入的事情，因为这类 I/O 事务通常通过事件和回调来处理。

# Spark SQL 自定义函数

## 普通函数

[doc-Scalar User Defined Functions (UDFs)](https://spark.apache.org/docs/latest/sql-ref-functions-udf-scalar.html),非聚合的函数，若放置在 SQL 语句中，且将列作为参数传入，则该函数会被列的每个值调用，即该函数会产生多行的结果。

```scala
val random = udf(() => Math.random()) // 使用 udf 函数包装普通函数（返回值为 UserDefinedFunction）即可注册成为 UDF
spark.udf.register("random", random) // 注册
spark.sql("SELECT random()").show() // SparkSession

def normSDist(): UserDefinedFunction = {
    val fn = (x: Double) => normalSDist(x, 0, 1) // 函数包装，适合写很多函数，可对函数加文档注释说明
    udf(fn)
}

spark.udf.register("norms_dist", normSDist())
// SELECT norms_dist(height) FROM student;
```



## 聚合函数

[doc-User Defined Aggregate Functions (UDAFs)](https://spark.apache.org/docs/latest/sql-ref-functions-udf-aggregate.html),聚合函数，会将多行汇集成一行输出。内部通过类似 `reduce` 的方法迭代生成。

```scala
import org.apache.spark.sql.Row
import org.apache.spark.sql.expressions.{MutableAggregationBuffer, UserDefinedAggregateFunction}
import org.apache.spark.sql.types.{DataType, IntegerType, LongType, StructField, StructType}

class CountFail extends UserDefinedAggregateFunction {
  /**
    * 设置输入数据的类型，指定输入数据的字段与类型，它与在生成表时创建字段时的方法相同
    * @return
    */
  override def inputSchema: StructType = StructType(StructField("inputColumn", IntegerType) :: Nil)

  /**
    * 指定缓冲数据的字段与类型，相当于中间变量
    * @return
    */
  override def bufferSchema: StructType = {
    StructType(StructField("count", LongType) :: Nil)
  }

  // 返回值的数据类型
  override def dataType: DataType = LongType

  /**
    * 设置该函数是否为幂等函数
    * 幂等函数:即只要输入的数据相同，结果一定相同
    * true表示是幂等函数，false表示不是
    * @return
    */
  override def deterministic: Boolean = true

  /**
    * initialize用于初始化缓存变量的值，也就是初始化 bufferSchema 函数中定义的变量值
    * 其中buffer(i)就表示第 i 个参数 （i = 0， i++）
    * @param buffer
    */
  override def initialize(buffer: MutableAggregationBuffer): Unit = {
    buffer(0) = 0L
  }

  /**
    * 当有一行数据进来时就会调用 update 一次，有多少行就会调用多少次，input 就表示在调用自定义函数中有多少个参数，最终会将
    * 这些参数生成一个Row对象，在使用时可以通过input.getString或input.getLong等方式获得对应的值
    * 缓冲中的变量sum,count使用buffer(0)或buffer.getDouble(0)的方式获取到
    * @param buffer 已有 buffer
    * @param input  新加入的行
    */
  override def update(buffer: MutableAggregationBuffer, input: Row): Unit = {
    if (!input.isNullAt(0) && input.getInt(0) == 0) {
      buffer(1) = buffer.getLong(0) + 1
    }
  }

  /**
    * 将更新的缓存变量进行合并，有可能每个缓存变量的值都不在一个节点上，最终是要将所有节点的值进行合并才行
    * 其中 buffer1 是本节点上的缓存变量，而 buffer2 是从其他节点上过来的缓存变量然后转换为一个 Row 对象,然后将 buffer2
    * 中的数据合并到 buffer1 中去即可
    * @param buffer1 本节点上的缓存变量
    * @param buffer2 其他节点上过来的缓存变量
    */
  override def merge(buffer1: MutableAggregationBuffer, buffer2: Row): Unit = {
    buffer1(0) = buffer1.getLong(0) + buffer2.getLong(0)
  }

  /**
    * 一个计算方法，用于计算最终结果,也就相当于返回值
    * @param buffer
    * @return
    */
  override def evaluate(buffer: Row): Long = buffer.getLong(0)
}
```

复杂的例子（spark 3.0.0 版本更简便的写法）:

```scala
import org.apache.spark.sql.{Encoder, Encoders}
import org.apache.spark.sql.expressions.Aggregator

/**
  * 给定一组点回归线的斜率
  * @author jianxinliu
  * @date 2020/8/27
  */
case class Params(colY: Double, colX: Double)
case class Buffer(var countX: Int, var sumX: Double, var sumY: Double, var sumXX: Double, var sumXY: Double)

object Slope extends Aggregator[Params, Buffer, Double] {
  override def zero: Buffer = Buffer(0, 0.0, 0.0, 0.0, 0.0)

  override def reduce(buffer: Buffer, params: Params): Buffer = {
    if (params != null) {
      buffer.countX += 1
      buffer.sumX += params.colX
      buffer.sumY += params.colY
      buffer.sumXX += params.colX * params.colX
      buffer.sumXY += params.colX * params.colY
    }
    buffer
  }

  override def merge(buffer: Buffer, b2: Buffer): Buffer = {
    buffer.countX += b2.countX
    buffer.sumX += b2.sumX
    buffer.sumY += b2.sumY
    buffer.sumXX += b2.sumXX
    buffer.sumXY += b2.sumXY
    buffer
  }

  override def finish(b: Buffer): Double = {
    (b.countX * b.sumXY - b.sumX * b.sumY) / (b.countX * b.sumXX - b.sumX * b.sumX)
  }

  // An encoder for Scala's product type (tuples, case classes, etc).
  override def bufferEncoder: Encoder[Buffer] = Encoders.product

  override def outputEncoder: Encoder[Double] = Encoders.scalaDouble
}

// 注册方法
sparkSession.udf.register("slope", functions.udaf(Slope))
```

聚合中有聚合（求一列数的标准差$\sigma = \sqrt{\frac{\Sigma_{i=1}^n(x_i - \bar{x})^2}{n-1}}$,每一次聚合中都有 $\bar{x}$ 这个固定且需要预先求取的值，按以往的聚合方式，肯定无法实现。则可以转换为聚合时不做任何 reduce 操作，只是把所有元素收集起来，最终再做运算）**稍微转变思路即可将原先不可能实现的事情做成**

```scala
import org.apache.spark.sql.{Encoder, Encoders}
import org.apache.spark.sql.expressions.Aggregator
import scala.collection.mutable.ListBuffer

case class KurtParam(col: Double)
case class KurtBuffer(list: ListBuffer[Double])

/**
  * 峰度 Kurtosis 计算，使用 Excel 公式
  * https://baike.baidu.com/item/%E5%B3%B0%E5%BA%A6
  *
  * kurtosis = { [n(n+1) / (n -1)(n - 2)(n-3)] sum[(x_i - mean)^4] / std^4 } - [3(n-1)^2 / (n-2)(n-3)]
  */
object Kurt extends Aggregator[KurtParam, KurtBuffer, Double] {
  override def zero: KurtBuffer = KurtBuffer(ListBuffer[Double]())

  override def reduce(b: KurtBuffer, a: KurtParam): KurtBuffer = {
    b.list.append(a.col)
    b
  }

  override def merge(b1: KurtBuffer, b2: KurtBuffer): KurtBuffer = {
    b1.list.appendAll(b2.list)
    b1
  }

  override def finish(r: KurtBuffer): Double = {
    val n = r.list.size.toFloat
    if (n < 4) {
      return Double.NaN
    }
    val v1 = (n * (n + 1)) / ((n - 1) * (n - 2) * (n - 3))
    val v3 = 3 * math.pow(n - 1, 2) / ((n - 2) * (n - 3))
    val mean = r.list.sum / n
    // 可使用 reduce 简化
    var stdDiffSum = 0.0
    r.list.foreach(li => {
      stdDiffSum += math.pow(li - mean, 2)
    })
    val std = math.sqrt(stdDiffSum / (n - 1))
    var tmpPow = 0.0
    r.list.foreach(li => {
      tmpPow += math.pow((li - mean) / std, 4)
    })
    v1 * tmpPow - v3
  }
  override def bufferEncoder: Encoder[KurtBuffer] = Encoders.product
  override def outputEncoder: Encoder[Double] = Encoders.scalaDouble
}
```



# HTML 表格任意选中

HTML 元素（主要是文本）能否被选中，是由 `user-select` css 属性控制的，若设置为 `none` 则不可选中，更多属性值参考 [MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/user-select).

HTML 页面的默认选中方式是行选择模式，即鼠标从按下到释放中间经过的所有行都会被选中。若要实现列选中模式或是任意选中模式，基本思路是：**将表格所有单元格设置为不可选中，在鼠标经过时，将对应的单元格设置可选中，即可实现任意选择的模式。** 以上思路有几点需要注意的：

1. 浏览器适配：完整的设置不可选中的样式为: `-webkit-user-select: none; -moz-user-select: none; -ms-user-select: none; user-select: none;`
2. 不可选中的元素：不一定是给单元格 `td` 设置不可选中，而应该给直接包裹文字的元素设置（如下例中是 `td` 中 class 为 `cell`的 `div`）。
3. 框选模式：该思路只能直线涂抹选中，即鼠标经过的 cell 会被选中。若想实现画对角线进行框选，还需要添加逻辑。
4. 事件：会涉及的事件：`mousedown`,`mousemove`,`mouseup`。若使用 jquery 则可以很方便的进行事件注册和 DOM操作，若使用 vue 则可以通过自定义指令 `directives` 得到需要操作的 DOM元素。

示例代码(Vue + elementUI)：

```js
const selectDisableStyle = `-webkit-user-select:none; -moz-user-select: none; -ms-user-select: none; user-select: none;`
...
directives: {
    areaSelect: { // 在需要自定义选择的元素上添加 v-areaSelect
        inserted: (el, binding, vnode) => {
            let randIds = new Map()
            let mouseDownFlag = false
            let mouseUpFlag = false
            let cells = []
            el.addEventListener('mousedown', function (event) {
                mouseDownFlag = true
                mouseUpFlag = false
                cells = []
                el.querySelectorAll('tr').forEach(tr => {
                    let row = tr.querySelectorAll('td div.cell')
                    row.length > 0 && cells.push(row)
                })
                cells.forEach((tdRow, idy) => {
                    tdRow.forEach((tdCol, idx) => {
                        const style = tdCol.getAttribute('style')
                        if (style.indexOf(selectDisableStyle) < 0) {
                            tdCol.setAttribute('style', style + selectDisableStyle)
                        }
                        // 若表格有 rowIndex ,cellIndex 则可不设 id
                        tdCol.setAttribute('id', `${idy + 1}_${idx + 1}`)
                    })
                })
                // 选中点击的 cell
                removeStyle(event)
            })

            function mouseMove(evt) {
                if (mouseUpFlag || !mouseDownFlag) {
                    return
                }
                // 缓存经过的 cell id
                randIds.set(evt.target.id, evt.target.id)
                // 选中
                removeStyle(evt)
            }

            el.addEventListener('mousemove', mouseMove)
            el.addEventListener('mouseup', function (evt) {
                mouseUpFlag = true
                mouseDownFlag = false
                // 框选逻辑
                let posList = Array.from(randIds).filter(v => v[0]).map(v => v[0]).map(v => v.split('_'))
                let posYList = posList.map(v => v[0])
                let posXList = posList.map(v => v[1])
                let minX = Math.min(...posXList), minY = Math.min(...posYList)
                let maxX = Math.max(...posXList), maxY = Math.max(...posYList)
                cells.forEach(cellRow => {
                    cellRow.forEach(cell => {
                        let [idy, idx] = cell.id.split('_').map(v => Number(v))
                        if (idx >= minX && idx <= maxX && idy >= minY && idy <= maxY) {
                            removeStyle(cell)
                        }
                    })
                })
                // 重置
                randIds = new Map()
                cells = []
            })
        }
    }
}

// 清除禁止选中的样式，同时选中
function removeStyle(evt) {
    let target = evt.target || evt
    let style = target.getAttribute('style') || selectDisableStyle
    let reg = new RegExp(selectDisableStyle, 'g')
    target.setAttribute('style', style.replace(reg, ''))
}
```

该方法虽然可实现任意区域框选，但复制的操作仍然不理想，复制到 Excel 中仍然会复制整行（可能是 ElementUI 的行为），复制到文本编辑器，多行多列的内容也会被合并为一列（单元格内容被换行或制表符分割）。

**第二种思路** 不去依赖浏览器的默认复制操作，而是自动将被复制内容写入剪切板。依然可以借鉴上一方法中对各种事件的监听，以及区域框选算法。只是对鼠标经过的单元格，不是设置 `user-select：none` 之类的样式，而是将单元格添加边框，以示选中。执行区域选中之后，程序是可以知道哪些单元格被选中的，此时可以将这些单元格的内容以想要的格式写入剪切板。

写入剪切板的思路：利用一个不可见 input 元素（若需要多行内容可以使用 textarea），将要复制的文本写入，再执行 setSelectionRange 选中，然后执行 `document.execCommand('copy')`，将 value 写入系统剪切板。

操作方式：按住 <kbd>Ctrl</kbd> 再使用鼠标选择，鼠标释放时自动框选，并将内容复制到剪切板。若不按 <kbd>Ctrl</kbd> 则仍旧可以使用浏览器自身的行选择模式。

```js
const selectStyle = 'border: 1px solid rgb(51,144,255); box-shadow: 0px 0px 5px 1px rgb(51,144,255);'
...
mounted() {
    // 按下 control 键
    // isCtrlPressed => this.ctrlPress > 0
    document.onkeydown = (e) => {
        if (e.keyCode === 17) {
            this.ctrlPress += 1
        }
    }
    document.onkeyup = (e) => {
        if (e.keyCode === 17) {
            this.ctrlPress = 0
        }
    }
}，
directives: {
    areaSelect: {
        inserted: (el, binding, vnode) => {
            let randIds = new Map()
            let mouseDownFlag = false
            let mouseUpFlag = false
            const vm = vnode.context // 获取当前组件的 Vue 实例
            let cells = []    // 表格中所有 cell
            let selectedCells = [] // 最终选中的 cell

            // 复制之后清除选中样式，单击会与现有事件冲突，改为双击
            document.addEventListener('dblclick', function () {
                if (!el) { // 该事件不好注销，故加此判断
                    return
                }
                el.querySelectorAll('tr').forEach(tr => {
                    let row = tr.querySelectorAll('td div.cell')
                    row.forEach(tdCol => {
                        tdCol.setAttribute('style', "")
                    })
                })
            })

            el.addEventListener('mousedown', function (event) {
                if (!vm.isCtrlPressed) {
                    return
                }
                mouseDownFlag = true
                mouseUpFlag = false
                cells = []
                el.querySelectorAll('tr').forEach(tr => {
                    let row = tr.querySelectorAll('td div.cell')
                    row.length > 0 && cells.push(row)
                })
                cells.forEach((tdRow, idy) => {
                    tdRow.forEach((tdCol, idx) => {
                        const style = tdCol.getAttribute('style')
                        // 为了界面简洁明了，选择过程中仍然禁止浏览器自身选中行为
                        if (style.indexOf(selectDisableStyle) < 0) {
                            tdCol.setAttribute('style', style + selectDisableStyle)
                        }
                        tdCol.setAttribute('id', `${idy + 1}_${idx + 1}`)
                    })
                })
                // 选中点击的 cell
                selectCell(event)
            })

            el.addEventListener('mousemove', function mouseMove(evt) {
                if (!vm.isCtrlPressed) {
                    return
                }
                if (mouseUpFlag || !mouseDownFlag) {
                    return
                }
                // 缓存经过的 cell id
                randIds.set(evt.target.id, evt.target.id)
                selectCell(evt)
            })

            el.addEventListener('mouseup', function (evt) {
                if (!vm.isCtrlPressed) {
                    return
                }
                mouseUpFlag = true
                mouseDownFlag = false
                let posList = Array.from(randIds).filter(v => v[0]).map(v => v[0]).map(v => v.split('_'))
                let posYList = posList.map(v => v[0])
                let posXList = posList.map(v => v[1])
                let minX = Math.min(...posXList), minY = Math.min(...posYList)
                let maxX = Math.max(...posXList), maxY = Math.max(...posYList)
                cells.forEach(cellRow => {
                    let selectedRow = []
                    cellRow.forEach(cell => {
                        let [idy, idx] = cell.id.split('_').map(v => Number(v))
                        if (idx >= minX && idx <= maxX && idy >= minY && idy <= maxY) {
                            selectCell(cell)
                            selectedRow.push(cell)
                        }
                        // 去除禁止选择的样式，仍然支持浏览器自身的行选择模式
                        removeStyle(cell)
                    })
                    selectedRow.length > 0 && selectedCells.push(selectedRow)
                })
                // WPS 默认单元格以 \t 分割，行以 \n 分割
                const done = copyToClipboard(selectedCells.map(v => v).map(row => row.map(cell => cell.innerText).join("\t")).join("\n"))
                if (done) {
                    vm.$message.success("内容已复制到剪切板！")
                } else {
                    vm.$message.error("该浏览器不支持复制！")
                }
                selectedCells = []
                randIds = new Map()
                cells = []
            })
        }
    }
}

function removeStyle(evt) {
    let target = evt.target || evt
    let style = target.getAttribute('style') || selectDisableStyle
    let reg = new RegExp(selectDisableStyle, 'g')
    target.setAttribute('style', style.replace(reg, ''))
}

function selectCell(evt) {
    let target = evt.target || evt
    // 可能会有其他元素进入，导致样式不美观
    if (target.getAttribute('class').indexOf('cell') < 0) {
        return
    }
    const style = target.getAttribute('style')
    if (style.indexOf(selectStyle) < 0) {
        target.setAttribute('style', style + ';' + selectStyle)
    }
}

function copyToClipboard(text) {
    const input = document.createElement('TEXTAREA');
    input.style.opacity = 0;
    input.style.position = 'absolute';
    input.style.left = '-100000px';
    document.body.appendChild(input);

    input.value = text;
    input.select();
    input.setSelectionRange(0, text.length);
    const done = document.execCommand('copy');
    document.body.removeChild(input);
    return done;
}
```

效果：

![tableColumnSelectMode](tableColumnSelectMode.png)



# 统计学

基本概率分布及其实现：Excel方式 -> [openoffice](https://github.com/apache/openoffice),实现代码主要在 [interpr](https://github.com/apache/openoffice/blob/trunk/main/sc/source/core/tool/interpr3.cxx)

js 中有 [jStat](https://github.com/jstat/jstat),[doc](https://jstat.github.io/all.html)

java 中有 [apache commons math3](https://commons.apache.org/proper/commons-math/)

c# 中有 [NMath]()

工程统计[概念、原理、公式](https://www.itl.nist.gov/div898/handbook/index.htm)



# 一个登录并保持的例子

前后端分离的登录例子。登录之后，使用 Session 保持一段时间有效，若这这件有请求，则每次请求更新过期时间。实际过期是因为一段时间没和服务端交互。

**后端做法**   在登录之后，后端返回 sessionId 和 token，前端每次请求都带上。后端再拦截每个请求，进行 sessionId 和 token 校验，并校验时间有无超出设定。

**注意**  在前端资源集成在后端发布时，会访问 `/` ,`/index.html` ,`/static/*` ，`/error` 等路径，需要排除。否则拦截后，基本页面都打不开。**惊魂一小时！！！系统宕机一小时，还好这个小系统没什么用户，而且是晚上，客户都下班了。否则宕机一小时可是一个很大的生产事故。**

**发布上线规范**  应该再正规一点，至少有回滚机制，发现错误后能回滚至上一个能运行的版本。发布时不要先删除已有 jar，否则新版发布不上去就无法回滚了。

```java
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.handler.HandlerInterceptorAdapter;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
import java.time.Instant;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.Arrays;
import java.util.Date;
import java.util.List;

@Slf4j
@Component
// 配合 ResponseBody 使用时会出现不能设置响应头的情况
public class LoginInterceptor extends HandlerInterceptorAdapter {

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        // 排除前端资源，不拦截
        List<String> excludePath = Arrays.asList("/static", "/error", "/user");
        if (excludePath.stream().anyMatch(v -> request.getRequestURI().startsWith(v))) {
            return true;
        }
        boolean passed;
        try {
            HttpSession session = request.getSession();
            String sessionId = Stream.of(request.getCookies()).filter(v -> "JSESSIONID".equals(v.getName())).limit(1).collect(Collectors.toList()).get(0).getValue();
            String reqToken = Optional.ofNullable(request.getHeader("access_token")).orElse("");
            boolean legalToken = StringUtils.isNotEmpty(reqToken) && reqToken.equals(session.getAttribute("access_token"));
            Instant lastTime = new Date(session.getLastAccessedTime()).toInstant();
            boolean notExpired = LocalDateTime.ofInstant(lastTime, ZoneId.systemDefault()).plusMinutes(30).isAfter(LocalDateTime.now());
            if (!notExpired) {
                session.invalidate();
            }
            passed = legalToken && StringUtils.isNotEmpty(sessionId) && notExpired;
        } catch (Exception e) {
            log.info("[Exception {}] Auth failure: 登录超时！", e.getClass());
            passed = false;
        }
        if (!passed) {
            response.setStatus(HttpStatus.UNAUTHORIZED.value());
            log.info("Auth failure !");
        }
        return passed;
    }
}
```

给每个请求返回 `refresh_token` 刷新。注意：使用 `@ResponseBody` 时，后置拦截器`postHandler` 中不能添加响应头，[原因](https://stackoverflow.com/questions/48823794/spring-interceptor-doesnt-add-header-to-restcontroller-services)

```java
import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

import java.util.Arrays;
import java.util.List;
import java.util.UUID;

@ControllerAdvice
public class ResponseBodyAdvisor implements ResponseBodyAdvice<Object> {
    @Override
    public boolean supports(MethodParameter returnType, Class<? extends HttpMessageConverter<?>> converterType) {
        return true;
    }

    @Override
    public Object beforeBodyWrite(Object body, MethodParameter returnType, MediaType selectedContentType, Class<? extends HttpMessageConverter<?>> selectedConverterType, ServerHttpRequest request, ServerHttpResponse response) {
        List<String> allowedHeaders = Arrays.asList("refresh_time", "refresh_token");
        // 需要添加前端允许接收的响应头，否则前端取不到
        response.getHeaders().setAccessControlAllowHeaders(allowedHeaders);
        response.getHeaders().setAccessControlExposeHeaders(allowedHeaders);
        response.getHeaders().add("refresh_time", String.valueOf(System.currentTimeMillis()));
        response.getHeaders().add("refresh_token", UUID.randomUUID().toString());
        return body;
    }
}
```

注册拦截器：

```java
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Component
public class WebConfig implements WebMvcConfigurer {

    private LoginInterceptor loginInterceptor;

    @Autowired
    public WebConfig(LoginInterceptor loginInterceptor){
        this.loginInterceptor = loginInterceptor;
    }

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // 前端资源挂载在后端的 static 中，第一次访问会先访问 / 和 /index.html ,需要排除掉
        registry.addInterceptor(loginInterceptor).excludePathPatterns("/", "/index.html");
    }
}
```

控制器中需要做的事：登录成功之后记录本次会话的 sessionId

```java
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.Errors;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;
import javax.validation.Valid;
import java.util.*;

@Slf4j
@RestController
@RequestMapping("/user")
public class UserController {
    
    private UserService userService;
    
    @Autowired
    public UserController(UserService userService) {
        this.userService = userService;
    }
    
    @PostMapping("/login")
    public BaseResp login(@RequestBody @Valid LoginParams params, HttpServletRequest req, Errors errors) {
        if (errors.hasErrors()) {
            StringJoiner sj = new StringJoiner(",");
            errors.getAllErrors().forEach(err -> sj.add(err.getDefaultMessage()));
            return new BaseResp(VALIDATED_PARAMETER_EX, sj.toString(), "");
        }
        log.info("login prams:{}", params.getUserId());
        Boolean passed = this.userService.login(params.getUserId(), Utils.pwdEncoder(params.getPwd()));
        String accessToken = "";
        String message = "failure";
        String code = VALIDATED_PARAMETER_EX;
        HttpSession session = req.getSession(true);
        if (passed) {
            accessToken = UUID.randomUUID().toString();
            message = "success";
            code = RETURN_CODE_OK;
            session.setAttribute("access_token", accessToken);
            // 将 session 设为永不过期，拦截器中判断是否过期，手动过期
            session.setMaxInactiveInterval(-1);
        }
        String ret = passed ? accessToken : "";
        return new BaseResp(code, message, ret);
    }
}
```

**前端做法**   登录成功之后，保存 sessionId 和 token，以后每次请求都带上

```js
let ret = await userService.login(params)
if (ret.data.result) {
    let token = res.data.result
    window.localStorage.setItem('access_token', token);
    let success_time = new Date().toString();
    window.localStorage.setItem('token_time', success_time);
}
```

axios 拦截器：设置每次请求携带相关请求头，自我检查是否登录超时，超时取消所发请求，并做相关清理工作。

```js
axios.defaults.withCredentials = true

const service = axios.create({
  baseURL: ''
});

service.interceptors.request.use(request => {
  // 登录相关请求忽略
  if (request.url.indexOf('/user') !== -1) {
    return request
  }
  //所有请求加入 token
  let access_token = window.localStorage.getItem('access_token');
  request.headers['access_token'] = access_token;

  // 校验 token_time 是否超时
  let token_time = window.localStorage.getItem('token_time');
  let date = new Date()
  let second = date.getTime() - new Date(token_time).getTime();   //时间差的毫秒数
  if (second > 1800000) {
    console.log('超时，请重新登录！')
    window.location.href = '/';
    throw new axios.Cancel("cancel: 超时，请重新登录！")
  }
  return request;
}, error => {
  return Promise.reject(error);
});

service.interceptors.response.use(response => {
    // 每次请求后，更新 token 和最后操作时间
    window.localStorage.setItem('token_time', response.headers['refresh_time']);
    window.localStorage.setItem('refresh_token', response.headers['refresh_token']);
    return response;
  },
  err => {
    // request canceled
    if (err.message && err.message.startsWith('cancel:')) {
      NProgress.done()
      Message.warning('超时，请重新登录！')
      console.log('request canceled')
      removeLocalStorageItem();
      throw new axios.Cancel(err.message)
    }
    if (err.response.status === 403 || err.response.status === 401) { // 没有权限
      let html = '<div><p> 即将跳转到登录页！</p></div>';
      MessageBox.alert(html, '登录超时！', {dangerouslyUseHTMLString: true})
      removeLocalStorageItem()
      window.location.href = '/';
    } else {
      let html = '<div>' + '<p> 错误提示: ' + JSON.stringify(err) + '</p>' + '</div>';
      MessageBox.alert(html, '服务器出错了', {dangerouslyUseHTMLString: true})
    }
  }
);

function removeLocalStorageItem(){
  window.localStorage.removeItem('access_token');
  window.localStorage.removeItem('refresh_token');
  window.localStorage.removeItem('token_time');
}
```

Session 工作原理：浏览器第一次请求服务时，服务器不识别该客户端，便生成新的 Session，并且在这次请求的响应头里添加类似 `Set-Cookies:JSESSIONID=XXXXX;HttpOnly` 的信息，告诉浏览器把本次的 SessionId 缓存起来，以后每次请求都在 Cookies 里带上，这样服务器就能识别是哪个客户端发的请求了。也就是不需要手动将 SessionId 返回给前端，也不需要前端每次请求在请求头携带 SessionId，不需要这样暴露。

**改进**   

1. 无需 refresh_token ，refresh_time.
2. 前端无需知道，无需维护超时时间，全由后端把控，超时返回 401 或 403 即可。



# 一个 CRUD 例子

使用 Jpa & Hibernate 进行数据库操作。

存在 Entity Student & Teacher，表结构 & Entity 如下：

```sql
create table student(
	sid varchar(10) primary key,
    name varchar（10） not null,
    birth timestamp,
    age int
)

create table teacher(
	tid varchar(10) primary key,
    name varchar(10),
    dept varchar(20)
)
```



```java
import javax.persistence.*;
import java.sql.Timestamp;

@Entity
@Table(name = "student", schema = "school", catalog = "")
public class Student extends UUIDObject {
    private String sid;
    private String name;
    private Timestamp birth;
    private Integer age;

    @Id
    @Column(name = "sid")
    public String getSid() {
        return colTyp;
    }

    public void setSid(String colTyp) {
        this.colTyp = colTyp;
    }

    @Basic
    @Column(name = "birth")
    public Timestamp getBirth() {
        return birth;
    }

    public void setBirth(Timestamp birth) {
        this.birth = birth;
    }
    
    // 其余字段 getter & setter 省略
    
    // equals & hashcode 方法省略
}

// teacher Entity 同理
```



## 数据源设定

```java
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;

import javax.sql.DataSource;

@Configuration
public class DataSourceConfiguration {

    /**
     *  第一个数据连接，默认优先级最高
     * @return
     */
    @Bean(name = "dataSourceFirst")
    @Primary
    @ConfigurationProperties(prefix = "spring.datasource.first")
    public DataSource dataSourceFirst() {
        //这种方式的配置默认只满足spring的配置方式，如果使用其他数据连接（druid）,需要自己独立获取配置
        return DataSourceBuilder.create().build();
    }

    /**
     * 第二个数据源
     * @return
     */
    @Bean(name = "dataSourceSecond")
    @ConfigurationProperties(prefix = "spring.datasource.second")
    public DataSource dataSourceSecond() {
        return DataSourceBuilder.create().build();
    }
}

```



```java
import com.xx.common.impl.BaseRepositoryImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;
import org.springframework.orm.jpa.JpaTransactionManager;
import org.springframework.orm.jpa.LocalContainerEntityManagerFactoryBean;
import org.springframework.orm.jpa.vendor.HibernateJpaVendorAdapter;
import org.springframework.transaction.PlatformTransactionManager;
import org.springframework.transaction.annotation.EnableTransactionManagement;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.sql.DataSource;

@Configuration
@EnableJpaRepositories(
        basePackages = "com.xx.repository.first",
        entityManagerFactoryRef = "firstEntityManagerFactoryBean",
        transactionManagerRef = "firstTransactionManager",
        repositoryBaseClass = BaseRepositoryImpl.class)
@EnableTransactionManagement
public class JpaFirstConfiguration {

    @Autowired
    @Qualifier("dataSourceFirst")
    private DataSource dataSource;

    @Bean(name = "firstEntityManagerFactoryBean")
    @Primary
    public LocalContainerEntityManagerFactoryBean entityManagerFactoryBean() {
        HibernateJpaVendorAdapter vendorAdapter = new HibernateJpaVendorAdapter();
        vendorAdapter.setGenerateDdl(false); // 第一次可开启，使用hibernate生成表，之后需关闭
        LocalContainerEntityManagerFactoryBean factory = new LocalContainerEntityManagerFactoryBean();
        factory.setJpaVendorAdapter(vendorAdapter);
        factory.setPackagesToScan("com.xx.entity.first");
        factory.setDataSource(dataSource);
        return factory;
    }

    @Bean(name = "firstEntityManager")
    @Primary
    public EntityManager entityManager() {
        return entityManagerFactoryBean().getObject().createEntityManager();
    }

    @Bean(name = "firstTransactionManager")
    @Primary
    public PlatformTransactionManager transactionManager(EntityManagerFactory entityManagerFactory) {
        JpaTransactionManager txManager = new JpaTransactionManager();
        txManager.setEntityManagerFactory(entityManagerFactory);
        return txManager;
    }
}

```



```java
import org.springframework.data.domain.Sort;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.NoRepositoryBean;

import java.io.Serializable;
import java.util.List;

@NoRepositoryBean
public interface BaseRepository<T extends UUIDObject, ID extends Serializable> extends JpaRepository<T, ID>, JpaSpecificationExecutor<T> {
}

```



```java
import com.navi.spc.common.core.dao.BaseRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.jpa.repository.support.JpaEntityInformation;
import org.springframework.data.jpa.repository.support.SimpleJpaRepository;

@Slf4j
public class BaseRepositoryImpl<T extends UUIDObject, ID extends Serializable> extends SimpleJpaRepository<T, ID> implements BaseRepository<T, ID> {

    private final JpaEntityInformation<T, ?> entityInformation;

    private final EntityManager em;

    BaseRepositoryImpl(JpaEntityInformation entityInformation,
                       EntityManager entityManager) {
        super(entityInformation, entityManager);
        // Keep the EntityManager around to used from the newly introduced methods.
        this.em = entityManager;
        this.entityInformation = entityInformation;
    }
}

```



```java
package com.xx.repository.first;

import java.util.List;
import java.util.Optional;

public interface StudentRepository extends BaseRepository<Student, String> {
    
    // jpa 默认提供部分 crud 实现，可直接调用
    // 另外，jpa 可根据定义的方法名生成 sql,而不需要实现，如：
    
    /**
     * 获取指定条件下最大的 sid
     *
     * @param name
     * @param age
     * @return
     */
    Optional<Student> findFirstByNameAndAgeOrderBySidNoDesc(String name, Integer age);

    /**
     * 重名检查,也可调用内置的 findById，支持联合主键
     *
     * @param sid
     * @param name
     * @return
     */
    List<Student> findBySidAndName(String sid, String name);

    /**
     *
     * @param name
     * @param age
     */
    void deleteByNameAndSid(String Name, String sid);

    // ....
}

```



总结： 对每个表建立一个 Entity 对应。通过操作该 Entity 进行数据库操作。

# Java 序列化

**什么是序列化**  将内存中的 Java 对象持久化成某种格式保存至磁盘或用于网络传输，可以简单的理解为 JavaScript 的 `JSON.stringify` 函数的功能。与之对应的还有**反序列化**，将保存成某种格式的对象恢复到内存中可运行的对象，可以简单的理解为 JavaScript 的 `JSON.parse` 函数。但与 JavaScript 相关函数不同的是，`stringify` 和 `parse` 函数会将对象原原本本的序列化、反序列化（js 的对象本身几乎和其字面量等价，也是简单的对象，甚至可以理解为是弱类型限制的 Map）,而 Java 的序列化则考虑更多如：访问性、是否静态、构造器，反序列化的对象比较……

**SerialVersionUID**  （也叫流标识符（Stream UniqueIdentifier），即类的版本定义的）在序列化反序列化时用于唯一标识一个对象。在反序列化时，JVM 通过该 ID 确认对象是否一致，是否被修改过，不同则序列化失败，`InvalidClassException`。

**反序列化时不执行构造函数，也不执行 getter,setter**  若在构造函数中修改 final  变量，则反序列化时赋值不会被应用。

**序列化的结果构成**  序列化后的文件构成：

1. 类描述信息。包括包路径、继承关系、访问权限、变量描述、变量访问权限、方法签名、返回值，以及变量的关联类信息。要注意的一点是，它**并不是class文件的翻版，它不记录方法、构造函数、static变量等的具体实现**。之所以类描述会被保存，很简单，是因为能去也能回嘛，这保证反序列化的健壮运行。
2. 非瞬态（transient关键字）和非静态（static关键字）的实例变量值。注意，这里的值如果是一个基本类型，好说，就是一个简单值保存下来；如果是复杂对象，也简单，连该对象和关联类信息一起保存，并且持续递归下去（关联类也必须实现Serializable接口，否则会出现序列化异常），也就是说递归到最后，其实还是基本数据类型的保存。

正是因为这两点原因，一个持久化后的对象文件会比一个class类文件大很多

**重写序列化反序列化方法控制序列化逻辑**  一个类能被序列化的前提是实现 Serializable接口，因为可以实现两个私有方法 `writeObject` 和 `readObject`，用来影响和控制序列化过程。如：控制部分属性不参与序列化（也可使用 `transient` 关键字）

```java
@Data
@AllArgsConstructor
class Student implements Serializable {
    private String name;
    private Integer age;
    private transient Double salary;
}

public class Test {
     public static void main(String[] args) {
        Student stu = new Student("jianxin", 23, 100.0);
        writeObject(stu);
        Student student = (Student) readObject(); // 不执行构造函数，不执行 getter\setter
        System.out.println(student.getName());
        System.out.println(student.getAge());
        System.out.println(student.getSalary());
    }

    private static String FIEL_PATH = "./out.bin";

    private static void writeObject(Serializable serializable) {
        try (ObjectOutputStream oot = new ObjectOutputStream(new FileOutputStream(FIEL_PATH))) {
            oot.writeObject(serializable);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static Object readObject() {
        Object ret = null;
        try (ObjectInputStream objectInputStream = new ObjectInputStream(new FileInputStream(FIEL_PATH))){
            ret = objectInputStream.readObject();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        return ret;
    }
}
```



# 项目总结









































