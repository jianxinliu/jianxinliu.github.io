# typescript learning

什么是 typescript？是 js  的一个超集，为 js 添加类型等静态语言的属性，增加编译环节用于检测语法，增加更易用的面向对象的属性，使得 ts 能被用于更加大型的工程，最终会编译成 js 代码。为 js 加上类型，其实就是在编译时对增加了类型注解的变量进行类型检查。

https://www.typescriptlang.org/docs/handbook/intro.html

## 基础类型

typescript几乎拥有相同的数据结构

``` ts
let isDone: boolean = false
let a: number = 6
let b: number = 4.3
let c: number = 0xfff
let name: string = "jack";// 支持es6的模板字符串

let arr: number[] = [1, 2, 3] // or
let arr: Array<number> = [1, 2, 3]

let x: [string, number] = ['hello', 10]// tuple

enum Color {Red, Green, Blue}

let color: Color = Color.Green

// 'maybe' could be a string, object, boolean, undefined, or other types
let maybe: unknown = ''
maybe = 4
maybe = false

if (typeof maybe === "boolean") {
    // TypeScript knows that maybe is a boolean now
    let aBoolean: boolean = maybe
    // Type 'boolean' is not assignable to type 'string'.
    let bString: string = maybe
}

// Unlike unknown, variables of type any allow you to access arbitrary properties, even ones that don’t exist. These properties include functions and TypeScript will not check their existence or type
let notSure: any = 4
notSure = "maybe s string"
notSure = false

let looselyTyped: any = 4;
// OK, ifItExists might exist at runtime
looselyTyped.ifItExists();
// OK, toFixed exists (but the compiler doesn't check)
looselyTyped.toFixed();

let strictlyTyped: unknown = 4;
strictlyTyped.toFixed();

// union type
type Season = 'Spring' | 'Summer' | 'Autumn' | 'Winter'
let day : Season = 'aa' // TS2322: Type '"as"' is not assignable to type 'Season'.
let day : Season = 'Spring' // right assignment
type LockState = 'Locked' | 'unLocked'
type oddNumberUnderTen = 1 | 3 | 5 | 7 | 9
function f(s: string | string[]) {
    return s.length
}

// Generic
type StringArray = Array<string>
let sArr: StringArray = ['1,2']
class Student {
    name: string
    age: number
}
type NameArray = Array<{ name: string }>
type StudentArray = Array<Student>
let student: NameArray = [{name: ''}]
// TS2322: Type '{ name: string; age: number; }' is not assignable to type '{ name: string; }'.
let student: NameArray = [{name: '', age: 1}] 
let stus: StudentArray = [{name: '', age: 2}]
let stu2 :StudentArray = [new Student()]


// recommended
function reverse(s: string): string {
    return s.split('').reverse().join()
}
function reverse2(s: String): String {
    return s.split('').reverse().join()
}
```

### 类型断言

一种是：

```typescript
let strlen : number = (<string>someStr).length
```
另一种是：
```typescript
let strlen : number = (someStr as string).length
```

## 变量声明

- 也能使用 `let`，`const`声明块作用域变量。
- 支持es6的解构声明

```js
let input = [1,2]
let [first,second] = input;
```

- 支持ES6的展开

```js
let first = [1,2,3]
let second = [4,5,6]
let both = [...first,...second,5]
```



## Interface

typescript 的接口并不像 java 一样有着严格的要求，而是聚焦在接口所要求的值上（type checking focuses on the *shape* that values have）。一个简单的例子理解 ts 中的接口：

```ts
// 也可以使用 interface 关键字写成单独的接口定义（只相当于给接口一个名称，同时具有可移植性）
function printLabel(labelObject:{label:string}) {
    console.log(labelObject.label)
}
// 正常工作。ts 会检查传入函数的参数是否包含 label 的字符串属性，有则认为传入的对象符合接口定义，而不会管是否传入的其他值
// 也不必显式声明该对象实现了接口
let obj = {label:'hello', value:12}
printLabel(obj) 

// 接口声明
interface LabelValue {
    label: string
}
function printLabel(labelObject: LabelValue) {
    console.log(labelObject.label)
}
```



### 可选属性

```ts
// 定义一个正方形的接口
interface Square {
    width: number
    color: string
}

function drawSeuare(s: Square) {
    draw(s.width, s.widht, s.color)
}

// 此处 Square 接口中包含了正方形的颜色属性，若是仅仅需要宽高，而不需要颜色，则该接口不能满足需求
// Optional Properties
interface Square {
    width: number
    color?: string
}

function calculateArea(s: Square) {
    return s.width * s.height
}
calculateArea({width: 100})
```



### readonly 只读属性

```ts
interface Point {
    readonly x:number
    readonly y:number
}
let p1:Point = {x:10,y:20}
p1.x = 1 // Cannot assign to 'x' because it is a read-only property.

// ReadonlyArray, ReadonlyMap, ReadonlySet  在原始类型上移除了所有改变元素的方法
let arr: ReadonlyArray<string> = ['1', '2']
arr = arr.push('4') // Property 'push' does not exist on type 'readonly string[]'
arr = ['4'] // 可以被成功赋值 ？

// readonly vs const
// readonly -> proerties; const -> variable
```



### 函数类型

```ts
interface searchFunc {
    (source: string, subStr: string, start: number): number
}
let findFrom: searchFunc = (source, subStr, start) => {
    return source.substring(start).indexOf(subStr)
}
console.log(findFrom('hello', 'l', 4))
```



## 类

ts 的类和 java 差不多，ts 类的 一般结构如下：

```
// 将此类导出，这样可以被 import
export class Person{
    private _id:number;
    public name:string;
    age:number;
    
    constroctor(id:number,name:string,age:number = 0){
        this.id = id;
        this.name = name;
        this.age = age
    }
    
    // getter and setter
    get id():number{
        return this._id
    }
    
    set id(id:number){
        this.id = id
    }
    
    // method
    private think(){
        console.log('I am thinking....')
    }   
}
```



## 函数





## 泛型



