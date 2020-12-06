# typescript learning

什么是 typescript？是 js  的一个超集，为 js 添加类型等静态语言的属性，增加编译环节用于检测语法，增加更易用的面向对象的属性，使得 ts 能被用于更加大型的工程，最终会编译成 js 代码。为 js 加上类型，其实就是在编译时对增加了类型注解的变量进行类型检查。

https://www.typescriptlang.org/docs/handbook/intro.html

## 基础类型

typescript几乎拥有相同的数据结构

```typescript
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



