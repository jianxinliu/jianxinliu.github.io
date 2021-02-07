# Clean Code for JavaScript[^almost]

https://github.com/ryanmcdermott/clean-code-javascript	

https://github.com/airbnb/javascript



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



### 不要使用 `var` , 尽量使用 `const`

使用 `const` 不可保证变量不会被重新赋值，在弱类型语言中，变量被重新赋值不能保证类型不发生改变。而类型发生改变也会导致不可预期的问题和增加理解成本。也可使用 `let` 在必须重新赋值的场景。

`let`,`const` 声明的变量是块级作用域，而 `var` 是函数级作用域。 

### 使用可自解释的变量名

**Bad:**

```js
const address = "One Infinite Loop, Cupertino 95014";
const cityZipCodeRegex = /^[^,\\]+[,\\\s]+(.+?)\s*(\d{5})?$/;
saveCityZipCode(
  address.match(cityZipCodeRegex)[1],
  address.match(cityZipCodeRegex)[2]
);

// flag for what???
let flag = true
let obj = {
    flag: false
}
if (obj.flag) {
    // xxx
}
```

**Good:**

```js
const address = "One Infinite Loop, Cupertino 95014";
const cityZipCodeRegex = /^[^,\\]+[,\\\s]+(.+?)\s*(\d{5})?$/;
const [_, city, zipCode] = address.match(cityZipCodeRegex) || [];
saveCityZipCode(city, zipCode);

let isActive = true
let tableData = {
    notEmpty: false
}
if (tableData.notEmpty) {
    // xxx
}
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



## 对象

### 对象属性

```js
// bad 
let value = 1
let obj = {
    value: value,
    addValue: function(v) {
        return obj.value + v
    }
}

// good
let obj = {
    'a': 1,
    value,
    'min-width': 3,
    addValue(v) {
        return obj.value + v
    }
}

// better (简写和非简写形式不要混在一起。分组写，并且简写形式放在对象定义的开头)
let obj = {
    value,
    a: 1,
    'min-width': 3, // 只有包含非法字符的属性名才加引号，否则不加
    addValue(v) {
        return obj.value + v
    }
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

### 使用默认参数而不是改变参数值

```js
// really bad
function handleThings(opts) {
  // No! We shouldn’t mutate function arguments.
  // Double bad: if opts is falsy it'll be set to an object which may
  // be what you want but it can introduce subtle bugs.
  opts = opts || {};
  // ...
}

// still bad
function handleThings(opts) {
  if (opts === void 0) {
    opts = {};
  }
  // ...
}

// good
function handleThings(opts = {}) {
  // ...
}

// 且将带默认值的参数放在最后（因为可以不传，而不会占位）
```



### 不要使用 `eval`



### 不要改变参数的值，不要给参数重新赋值

```js
// bad
function f1(a) {
  a = 1;
  // ...
}

function f2(a) {
  if (!a) { a = 1; }
  // ...
}

// good
function f3(a) {
  const b = a || 1;
  // ...
}

function f4(a = 1) {
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
    
  //some comment一些中英文混杂的注释comments with xxx
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
  // some comment 一些中英文混杂的注释 comments with xxx
}
```



## 工程

1. 使用全局通用的常量，定义共同术语
2. 使用全局通用的类型，规范类型定义（class）
3. 尽量组件化，而不是复制。便于更改。
4. **项目伊始，就要做出规范，并严格执行，否则后期维护是无尽的深渊。规范包括代码规范，业务上的流程规范**

```js
// Constants.js
export default {
    PAGE_SIZE: 100,
    EXPIRE_SECONDS: 30 * 60 * 60
    ...
}

// Types.js
class SqlTransformError {
    constructor(sql) {
        this.sql = sql
    }
    getError(msg) {
        return new Error(JSON.stringify({sql:this.sql, msg}))
    }
}
export class SqlTransformError
```



## 模块

1. 总是使用 `import` ,`export`， 而不是其他的模块系统。这是 ECMAScript 标准，是未来。
2. 尽量按需导入，减少全量导入 `imort * as xxx from 'xxx'`
3. 合并相同来源的导入 `import foo, {bar, buzz} from 'xxx'`
4. 模块导出的变量需要是 `const`  的，防止引用变化导致的隐含问题
5. 导入语句全部写在非导入语句前面
6. bad: `import foo from './foo.js'`, good: `import foo from './foo'`



## 杂项

1. 使用高级函数代替 `for in` `for of` 循环。以此强化代码的不可变性（immutable）。
2. 使用 `.` 访问对象属性。特殊情况如： 属性含有非法字符，属性是变量，才使用 `[]` 访问。（容易与数组访问混淆）
3. 使用 `**` 计算幂。（ `Math.pow(2, 3)` -> `2 ** 3`）
4. 一个 `const` 或 `let` 一次只声明一个变量。（bad: `let a, b, c = 0`）
5. 将使用 `const` 和 `let` 声明的变量放置在一起。
6. 变量声明处尽量靠近其使用处。
7. 使用 `===` , `!==` 代替 `==`, `!=` 进行比较
8. **条件语句如 if 对其表达式求值有一定的隐含规则**。如对象解析为 true, 0 解析为 false。所以通过 `if(array){...}` 判断数组有值是行不通的，数组就算是空 `[]` 也是一个对象。在条件语句表达式中，尽量写明，而不是让编译器去解析。只有 boolean 类型的可以直接写如 `if(isValid)`，字符串和数字最好写明比较表达式 `if(name !== '')` ,`if(list.length > 0)`。
9. 避免嵌套三元运算符。`a ? b : c ? d : e`  -> `const maybe = c ? d : e; const may = a ? b : maybe` 。也避免不必要的三元运算符。 `foo = a ? a : b` -> `foo = a || b`, `foo = a ? true : false` -> `foo = !!a`, `foo = a ? false : true` -> `foo = !a`
10. 对运算符使用括号分组，防止因为优先级混淆导致的错误。`const foo = a && b < 0 || c > 0 || d + 1 === 0;` -> `const foo = (a && b < 0) || c > 0 || (d + 1 === 0);`, `const bar = a ** b - 5 % d;` -> `const bar = a ** b - (5 % d);`。
11. 不要在循环内包含大量的代码，如果有，请抽成函数，尽量让循环的逻辑在一屏之内。循环内应该是逻辑的总体描述，而不应该包含太多细节。
12. 同样的，分支代码也不要超过一屏，如果有，请抽成函数。
13. 不同分支的代码，请抽出公共的部分，而不是再复制一份。即节省代码量又减少看代码时的理解负担。

## 其他

1. 通过设定默认值或添加卫语句，减少 `if-else`
2. 格式化！格式化！格式化！
3. 使用 `async await` 代替 `then` 回调带来的“缩进地狱”

```js
// 1.
function confirm({name, sex}) {
    if (!name || !sex) {
        // alarm ...
        return
    }
    if (!/someReg/.test(name)) {
        // alarm...
        return
    }
    let handler = FemaleHandler()
    if (sex === 'male') {
        handler = MaleHandler()
    }
    // ...
}

// 3.
// Bad
api.getClassId(className).then(res => {
    if (res.data.result) {
        const classId = res.data.result
        api.listStudents(classId).then(result => {
            if (result.data.result.length > 0) {
----------------let students = result.data.result
----------------// .... balabala
            }
        })
    }
})

// Better
const classIdResp = await api.getClassId(className)
if (!classIdResp.data.result) return
const classId = classIdResp.data.result
const studentResp = await api.listStudents(classId)
if (studentResp.data.result.length < 1) return
let students = studentResp.data.result
// .... balabala
```



`Array.prototype.forEach` VS `Array.prototype.map`

使用 `map` 必须在回调函数中 `return` ，以此返回一个新的数组，否则对原数组修改，使用 `forEach`



```js
// bad
students.map(stu => {
    stu.name = `${stu.firstName}-${stu.lastName}`
})

// better
students.forEach(stu => {
    stu.name = `${stu.firstName}-${stu.lastName}`
})
const students = students.map(stu => {
    stu.name = `${stu.firstName}-${stu.lastName}`
    return stu
})
```



# for Echarts

责任链模式对 Option 进行设置



# for Vue

使用 Vuex mapState 函数，给 store 中的变量添加 namespace，对一堆变量进行分类、分层。（否则 store 中的变量全部都是绝对路径，一行代码全看路径了，看不到真正的代码逻辑）

把逻辑交给 Vue（or ElementUI）。制定特殊的数据格式，以此来代替 `if else` 或循环。

Vue option API : options 大致按 `name -> data(computed) -> created(mounted) -> methods -> afterXXX -> watchs ` 的顺序写。其中 `methods` 各方法之间至少有一个空行。

逻辑操作和数据操作剥离。 **store 中的数据操作，在 store 中完成，不要放在程序逻辑中**。获取 store 中的变量值，使用 `getters`, 设置 store 中的变量，使用 `mutation`。





# 一般规范 & 体验优化

1. 表格中每行须有一个 id , 便于删改
2. select 选择组件，需支持搜索
3. 输入框后表示输入结束的按钮，需添加按键触发
4. 删除不必要的注释，被注释的代码
5. Vue 等框架中，剥离和框架实例不相关的功能函数
6. 相关常量单文件存放
7. Vue 中，Vuex 相关变量操作在 mutation 中执行，代码中使用 mapMutation 引入 （代码中分离数据操作和逻辑操作）
8. 函数中，参数不符合函数执行条件，及时退出
9. 及时整理代码，发现坏味道及时清理
10. 严格控制单行代码长度（看代码时尽量不需要横向滚动）及函数行数。（缩进地狱）
11. 给条件表达式命名。如 ： `if (a && b || c.indexOf('xxx') > -1) then ...` -> `withXXX = c.indexOf('xxx') > -1; needUpdate = a && b || withXXX; if (needUpdate) then ...`
12. 整体观感：紧凑不挤，疏密得当，减少复制
13. 整体原则：不多写一行无用代码，不少写一行增强明确性的代码
14. 代码提交前再过一遍改动的地方，检查有无废弃变量，代码，注释，日志打印未删除，代码未格式化。
15. 和 DB 紧密接触的系统需要做 DB 抽象层，将 DB 相关的细节限制在这一层中，便于集中修改。可以使用 ts 等强类型手段限制 DB 层的输入输出。



# ESLint 配置

```js

```







# 工程 & 项目规范

## 代码

代码规范参考上文

### 代码提交规范

具体什么强制格式倒没必要。主要目的是一目了然，便于搜索，提交时写上以下信息即可：

1. 什么类型的提交。新功能？ debug？ 更新文档？……
2. 提交的代码属于的模块、页面、接口、服务……
3. 提交内容详情，最好分要点写。debug 需要写简明的原因

把信息持久化到文档里，而不是脑子里。

## 文档

文档类型：

1. 业务 & 需求需要固化为文档，可以由 PM 和 Dev 共同完善。

2. 代码 & 逻辑文档，如果模块足够复杂的话。主要有 Dev 编写。

3. 使用手册，可由 PM 、 Dev、Test 编写完善。

4. 版本记录 & 上线日志。由 Ops 负责编写维护。

5. 项目基础信息文档。如服务器、DB地址 & 账密等。

以上文档注意及时更新。同时：

1. 文档以文本或图片形式，加入版本控制（Markdown，Git）。
2. 文档应当和代码一样贯穿整个项目，是项目重要的一部分。
3. 如使用手册应当支持开放给客户。

不错的文档 wiki 产品：

1. http://dynalon.github.io/mdwiki/#





