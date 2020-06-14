# Vue

## lifecycle

![](./vue-lifecycle.png)

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

# CSS

css 选择器上的`/deep/` 修饰：https://stackoverflow.com/questions/25609678/what-do-deep-and-shadow-mean-in-a-css-selector

HTML5 Web Components offer full encapsulation of CSS styles.

This means that:

- styles defined within a component cannot leak out and effect the rest of the page
- styles defined at the page level do not modify the component's own styles

 

- 修改第三方组件的样式时，应该注意影响范围，最好在第三方组件外围包裹元素，再使用 CSS 子元素选择器选择第三方组件。



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



# 文档生成工具

https://docsify.js.org/#/    https://github.com/docsifyjs/docsify

https://www.showdoc.cc/

https://github.com/phachon/mm-wiki



# JS Promise

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
    if(filds.length < 1){
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
      ret[p] = groupByMutil(ret[p], fields, i)
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