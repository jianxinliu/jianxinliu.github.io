# angular学习随笔记

## 1.\<ng-content\>

是外部组件在本组件中的占位符。当进行组件嵌套时很管用。

## 2.ngOnChange（changes:SimpleChange）	

是生命周期的钩子函数，当组件或指令的输入属性发生变化时调用。

```typescript
ngOnChanges(changes: SimpleChanges) {
  for (let propName in changes) {
    let chng = changes[propName];
    let cur  = JSON.stringify(chng.currentValue);
    let prev = JSON.stringify(chng.previousValue);
    this.changeLog.push(`${propName}: currentValue = ${cur}, previousValue = ${prev}`);
  }
}
```

## 3.模板引用变量

```typescript
<input #phone placeholder="phone number">

<!-- lots of other elements -->

<!-- phone refers to the input element; pass its `value` to an event handler -->
<button (click)="callPhone(phone.value)">Call</button>
```

## 4.输入和输出属性

组件间通信。

父组件传值给子组件，子组件通过 `@Input()`装饰变量进行接收，子组件传值给父组件，子组件通过`@Output()`装饰变量进行输出。

```typescript
// 子组件
@Component({
    selector:'son',
    template:`
    	<p>{{value}}</p>
    `
})
export class Son{
    @Input() value;
}

// 父组件调用
@Component({
    selector:'parent',
    template:`
    	<son [value]='kk'></son>
    `
})
```

## 5.管道操作符

json管道对调试有帮助

```typescript
<div>{{currentHero  | json}}</div>

// output like that:

{ "id": 0, "name": "Hercules", "emotion": "happy",
  "birthdate": "1970-02-25T08:00:00.000Z",
  "url": "http://www.imdb.com/title/tt0065832/",
  "rate": 325 }
```

## 6.@ViewChild(${Component})

是一个属性装饰器，用于配置一个视图查询，变更检测器会在DOM中查找一地个匹配该选择器的元素或指令。就是提供了在ts代码中引用组件的能力。

所支持的选择器包括：

- 任何带有 @Component @Directive装饰器的类
- 模板引用变量。@ViewChild('cmp') 引用\<cc #cmp\>\</cc\>
- 组件树种任何当前组件的子组件所定义的提供商。@ViewChild('someService') 引用 someService:SomeService
- 任何通过字符串令牌定义的提供商（比如 `@ViewChild('someToken') someTokenVal: any` ） 
- templateRef。@ViewChild('TemplateRef')template;引用\<ng-template\>\</ng-template\>

