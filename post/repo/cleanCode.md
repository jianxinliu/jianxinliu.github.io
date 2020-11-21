# Clean Code for JavaScript

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



## 工程

1. 使用全局通用的常量，定义共同术语
2. 使用全局通用的类型，规范类型定义（class）
3. 尽量组件化，而不是复制。便于更改。

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



# for Echarts

责任链模式对 Option 进行设置



# for Vue

使用 Vuex mapState 函数，给 store 中的变量添加 namespace，对一堆变量进行分类、分层。

把逻辑交给 Vue（or ElementUI）。制定特殊的数据格式，以此来代替 `if else` 或循环。

Vue option API : options 大致按 `name -> data(computed) -> created(mounted) -> methods -> afterXXX -> watchs ` 的顺序写。其中 `methods` 各方法之间至少有一个空行。