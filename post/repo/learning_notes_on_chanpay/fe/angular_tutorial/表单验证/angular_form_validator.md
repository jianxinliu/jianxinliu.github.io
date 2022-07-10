# angular表单

- 响应式表单，如果表单是应用中的重要部分，应该使用响应式表单

- 模板驱动表单，在往应用中添加简单表单时非常有用

响应式表单中，表单模型是显式定义在组件类中的。接着，响应式表单指令（如`FormControl`）会把这个现有的表单控件实例通过数据访问器（`ControlvalueAccessor`）来指派给视图的表单元素。

```typescript
import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';
 
@Component({
  selector: 'app-reactive-favorite-color',
  template: `
    Favorite Color: <input type="text" [formControl]="favoriteColorControl">
  `
})
export class FavoriteColorComponent {
  favoriteColorControl = new FormControl('');
}
```



模板驱动表单的例子：

```typescript
import { Component } from '@angular/core';

@Component({
  selector: 'app-template-favorite-color',
  template: `
    Favorite Color: <input type="text" [(ngModel)]="favoriteColor">
  `
})
export class FavoriteColorComponent {
  favoriteColor = '';
}
```



模板驱动表单，是抽象了表单模型，`ngModule`指令负责创建和管理制定表单上的表单控件实例。它不那么明显，但你不必要再直接操作表单了。

|                  | 响应式                 | 模板驱动           |
| ---------------- | ---------------------- | ------------------ |
| 建立（表单模式） | 显式，在组件类中创建。 | 隐式，由组件创建。 |
| 数据模式         | 结构化                 | 非结构化           |
| 可预测性         | 同步                   | 异步               |
| 表单验证         | 函数                   | 指令               |
| 可变性           | 不可变                 | 可变               |
| 可伸缩性         | 访问底层 API           | 在 API 之上的抽象  |

## 共同基础

无论是响应式表单还是模板驱动表单，都有一些共同的基础

- `FormControl`实例用于追踪单个表单的值和验证状态，代表一个控件
- `FormGroup`实例用于追踪一个表单控件集的值和状态，代表一个控件组
- `FormArray`实例用于追踪一个表单控件数组的值和状态
- `ControlValueAccessor`用于在Angular的FormControl实例和原生DOM元素之间创建桥梁

## 表单验证

- 响应式表单把自定义验证器定义成函数，它以要验证的控件作为参数
- 模板驱动表单和模板指令相关，并且必须提供包装了验证函数的自定义验证器指令

## 可变性

如何跟踪变更，对于应用的运行效率起着重要作用。

- **响应式表单**通过将数据模型提供为**不可变数据结构**来**保持数据模型的纯粹性**。每当在数据模型上触发更改时，`FormControl` 实例都会返回一个新的数据模型，而不是直接修改原来的。这样能让你通过该控件的可观察对象来跟踪那些具有唯一性的变更。这可以让变更检测更高效，因为它只需要在发生了唯一性变更的时候进行更新。它还遵循与操作符相结合使用的 "响应式" 模式来转换数据。
- **模板驱动表单**依赖于可变性，它使用双向数据绑定，以便在模板中发生变更时修改数据模型。因为在使用双向数据绑定时无法在数据模型中跟踪具有唯一性的变更，因此变更检测机制在要确定何时需要更新时效率较低。

以 "喜欢的颜色" 输入框元素为例来看看两者有什么不同：

- 对于响应式表单，每当控件值变化时，**FormControl 实例**就会返回一个新的值。
- 对于模板驱动表单，**favoriteColor** 属性总是会修改成它的新值。



# 响应式表单

## `FormControl`,`FormGroup`,`FormBuilder`

前两个在上文已经粗略的涉及过，在此说明FormBuilder。

FormBuilder用来生成表单控件，省去了频繁`new FormControl()`操作。使用FormBuilder需要在构造器中注入即可。`constructor(private fb: FormBuilder) { } `。有三个方法,分别用于生成对应对象实例：

- control()：FormControl
- group():FormGroup
- array():FormArray

例如：

```typescript
import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-profile-editor',
  templateUrl: './profile-editor.component.html',
  styleUrls: ['./profile-editor.component.css']
})
export class ProfileEditorComponent {
  profileForm = this.fb.group({
    firstName: [''],
    lastName: [''],
    address: this.fb.group({
      street: [''],
      city: [''],
      state: [''],
      zip: ['']
    }),
  });

  constructor(private fb: FormBuilder) { }
}
```



上面的代码使用初始值 `''` 来创建FormControl，但如果控件需要同步或异步验证器，那就在这个数组中的第二项和第三项中提供同步和异步验证器。如：`name    : [ null, [ Validators.required ] ], `。当往表单控件上	添加了一个**必填字段**后，其**初始值是无效**（invalid），这种**无效状态会传到其父组件**FormGroup中，从而也让FormGroup的状态变为无效。可以使用FormGroup实例的status属性访问当前状态。



## 表单校验Demo1

```html
<form nz-form [formGroup]="jx" (ngSubmit)="submit()">
  <input nz-input formControlName='name'>
  <input nz-input formControlName='age'>
  <button type="submit" nz-button [disabled]="!jx.valid">up</button>
</form>
```

```typescript
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, FormBuilder, Validators } from '@angular/forms';

export class FormComponent implements OnInit {
  jx :FormGroup
  constructor(private fb:FormBuilder) { }

  ngOnInit() {
    // 创建组件并设置验证器
    this.jx = this.fb.group({
      name:['jx',[Validators.required,
                Validators.maxLength(4),
                Validators.minLength(2)]],
      age:[12,[Validators.required]]
    })
  }
  getFormValue(){
    let ret = {};
    let form = this.jx.controls;
    for (const c in form) {
      ret[c] = form[c].value
    }
    return ret;
  }
  submit(){
    console.log(this.jx.controls)// 这是一个对象
    for (const control in this.jx.controls) {// 遍历对象的属性
      console.log(this.jx.controls[control].value)
    }
    // 获取控件 name 的值
    console.log(this.jx.controls['name'].value)
  }
}

```

注意点：

1. form 标签的 `[formGroup]`属性用于指定一个`FormGroup`
2. input 标签中有一个 `formControlName` 属性，用于指定`FormGroup`中对应的控件
3. button 按钮中的 `[disabled]="!jx.valid"`，表示当`FormGroup`的状态不是已验证的时候后，按钮不可用。
4. 构造器中注入 `FormBuilder`
5. 页面初始化完成后实例化`FormBuilder`，同时指定验证器。
6. name 这个控件中默认值是 "jx" ，并指定了三个验证器：`required`,`maxlength=4`,`minlength=2`



## FormArray 控件

FormArray控件也是用来批量添加控件的，**优势是可以动态添加控件**，就像操作数组一样。使用场景如问卷调查设计网站支持动态添加问题和答案。





## 模板驱动表单

