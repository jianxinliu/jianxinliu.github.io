## Angular.js

angularjs 是一种单页应用的前端框架。

angular 通过指令扩展 HTML，通过表达式绑定数据到HTML	。

最简单的一段代码：

```html
<div ng-app="">
    <input type="text" ng-model="name">
    <h1>{{name}}</h1>
</div>
```

- `ng-app`指令定义angular 的作用域
- `ng-model`指令将数据绑定到程序变量

## 指令

### ng-repeat

可以重复 HTML 元素

```html
<ul>
    <li ng-repeat='x in arr'>
        {{x}}
    </li>
</ul>
```

遍历对象属性：

```html
<ul>
    <li ng-repeat='x in obj'>
        {{x}}
    </li>
</ul>
```

### ng-model

将输入域的值与变量绑定

##### 校验用户输入

```html
<form ng-app="" name="myForm">
    Email:
    <input type="email" name="myAddress" ng-model="text">
    <span ng-show="myForm.myAddress.$error.email">不是一个合法的邮箱地址</span>
</form>
```

##### `ng-model` 指令可以为应用数据提供状态值(invalid, dirty, touched, error): 

```html
<form ng-app="" name="myForm" ng-init="myText = 'test@runoob.com'">
    Email:
    <input type="email" name="myAddress" ng-model="myText" required></p>
    <h1>状态</h1>
    {{myForm.myAddress.$valid}}
    {{myForm.myAddress.$dirty}}
    {{myForm.myAddress.$touched}}
</form>
```

`ng-model` 指令根据表单域的状态添加/移除以下类：

- ng-empty
- ng-not-empty
- ng-touched
- ng-untouched
- ng-valid
- ng-invalid
- ng-dirty
- ng-pending
- ng-pristine

所以可以指定对以上类指定自己的样式

### 自定义指令

```js
app.directive('jx', function () {
    return {
        restrict: "ECMA",
        template: `自定义指令`
    }
})
```

## scope

Scope(作用域) 是应用在 HTML (视图) 和 JavaScript (控制器)之间的纽带。

Scope 是一个对象，有可用的方法和属性。

Scope 可应用在视图和控制器上。



当在控制器中添加 **$scope** 对象时，视图 (HTML) 可以获取了这些属性。

视图中，你不需要添加 **$scope** 前缀, 只需要添加属性名即可，如： **{{carname}}**。



AngularJS 应用组成如下：

- View(视图), 即 HTML。
- Model(模型), 当前视图中可用的数据。
- Controller(控制器), 即 JavaScript 函数，可以添加或修改属性。

scope 是模型。

#### scope 的作用域

所有的应用都有一个 **$rootScope**，它可以作用在 **ng-app** 指令包含的所有 HTML 元素中。

**$rootScope** 可作用于整个应用中。是各个 controller 中 scope 的桥梁。用 rootscope 定义的值，可以在各个 controller 中使用。

### `$scope` 继承

每一个 `ng-controller` 指令都会创建一个新的 `$scope`,但都属于 `$rootScope` 。

若创建一个嵌套的controller，如下：

```html
 <div ng-controller="c1">
     <p>{{name}},{{age}}</p>
     <div ng-controller="c2">
         <p>{{name}},{{age}}</p>
         <div ng-controller="c3">
             <p>{{name}},{{age}}</p>
         </div>
     </div>
</div>
```

```js
app.controller('c1',function($scope){
    $scope.name = 'c1ppp'
    $scope.age = 1
})
app.controller('c2',function($scope){
    // 覆盖 c1 scope 中的两个变量
    $scope.name = 'c1ppp'
    $scope.age = 2
})
app.controller('c3',function($scope){
    // 页面显示 c2 scope 中的值
    // $scope.name = 'c1ppp'
    // $scope.age = 3
})
```



Controller c3 会继承 c2 ,若controller c3 的Scope中没有的变量，angular会到 c2 的Scope中去找，一直往上，直到找到。和 java 的类加载机制一样。



## Controller

AngularJS 应用程序被控制器控制。

**ng-controller** 指令定义了应用程序控制器。

控制器是 **JavaScript 对象**，由标准的 JavaScript **对象的构造函数** 创建。

用Controller来做：

- 给 `$scope` 设置初始化对象
- 给 `$scope` 设置行为

**不用Controller 来做**：

- 操作DOM，应该使用数据绑定来操作DOM。Controller只应该包含业务逻辑，将视图逻辑显示的放在Controller中会影响Controller的可测试性。
- 格式化输入。使用AngularJs form controls代替。
- 过滤输出。使用 AngularJs filters代替。
- 跨controller访问代码或状态，使用Angularjs Service 代替。
- 管理其他 组件的生命周期，如创建Service。

通常，**控制器只应该包含一个视图的必要的业务逻辑**。保持Controller简洁的办法常常是将不应该属于Controller的工作封装成一个service ,再通过依赖注入来使用。



## Filter

管道，转换数据

| 过滤器    | 描述                     |
| --------- | ------------------------ |
| currency  | 格式化数字为货币格式。   |
| filter    | 从数组项中选择一个子集。 |
| lowercase | 格式化字符串为小写。     |
| orderBy   | 根据某个表达式排列数组。 |
| uppercase | 格式化字符串为大写。     |

出了可以在表达式中添加管道，也可以在指令中添加：

```html
<ul>
    <li ng-repeat="x in names | orderBy:'country'">
        {{ x.name + ', ' + x.country }}
    </li>
</ul>
```

#### 过滤输入 

```html
 <ul>
  <li ng-repeat="x in names | filter:test | orderBy:'country'">
    {{ (x.name | uppercase) + ', ' + x.country }}
  </li>
</ul>
```

#### 自定义过滤器

```js
app.filter('reverse', function() { //可以注入依赖
    return function(text) {
        return text.split("").reverse().join("");
    }
});
```

## Service

在 AngularJS 中，服务是一个函数或对象，可在你的 AngularJS 应用中使用。一般通过依赖注入将服务注入到Controller中使用。服务可以在组件甚至 app 间分享代码。

服务的特性：

- 懒加载：服务只在被依赖时才会被实例化。
- 单例：服务是单例的，多个组件依赖的也是同一个服务对象。Service是通过工厂获得的。

AngularJS 内建了30 多个服务。

- 有个 **$location** 服务，它可以返回当前页面的 URL 地址。

注意 **$location** 服务是作为一个参数传递到 controller 中。如果要使用它，需要在 controller 中定义。 

- **$http** 是 AngularJS 应用中最常用的服务。 服务向服务器发送请求，应用响应服务器传送过来的数据。 

```js
app.controller('myCtrl', function($scope, $http) {
    $http.get("welcome.htm").then(function (response) {
        $scope.myWelcome = response.data;
    });
});
```

- AngularJS **$timeout** 服务对应了 JS **window.setTimeout** 函数。 
- AngularJS **$interval** 服务对应了 JS **window.setInterval** 函数。 

#### 自定义服务

```js
app.service('hexafy', function() {
    this.myFunc = function (x) {
        return x.toString(16);
    }
});
// use
app.controller('myCtrl', function($scope, hexafy) {
    $scope.hex = hexafy.myFunc(255);
});
```

## HTTP

```js
// 简单的 GET 请求，可以改为 POST
$http({
    method: 'GET',
    url: '/someUrl'
}).then(function successCallback(response) {
        // 请求成功执行代码
    	//response.data
    }, function errorCallback(response) {
        // 请求失败执行代码
});
```

此外还有以下简写方法：

- $http.get
- $http.head
- $http.post
- $http.put
- $http.delete
- $http.jsonp
- $http.patch

## Select

**ng-repeat** 指令是通过数组来循环 HTML 代码来创建下拉列表，但 **ng-options** 指令更适合创建下拉列表，它有以下优势：

使用 **ng-options** 的选项是一个对象， **ng-repeat** 是一个字符串。

```jsx
<select ng-model="selectedSite" ng-options="x.site for x in sites">
</select>


$scope.cars = {
    car01: { brand: "Ford", model: "Mustang", color: "red" },
    car02: { brand: "Fiat", model: "500", color: "white" },
    car03: { brand: "Volvo", model: "XC90", color: "black" }
};
<!--(x,y)是遍历对象是，x 为 属性名 , y 为 属性值-->
<select ng-model="oobj" ng-options="y.brand for (x,y) in cars"></select>
<span>{{oobj}}</span><!--是一个对象-->

// ng-option 指令的值：XX for WW in PP
// XX 为 option 的 label
// WW 为 遍历的项
// PP 为数据源
// ng-model 对应的是一个对象
```

## 表格

```html
<table class="table">
    <thead>
        <tr>
            <th>No.</th>
            <th>brand</th>
            <th>model</th>
            <th>color</th>
        </tr>
    </thead>
    <tbody>
        <tr ng-repeat="x in cars | orderBy : 'color'">
            <td>{{$index + 1}}</td>
            <td ng-if="$odd" style="background-color:#f1f1f1">{{x.brand | uppercase}}</td>
            <td ng-if="$even">{{x.brand | uppercase}}</td>
            <td ng-if="$odd" style="background-color:#f1f1f1">{{x.model}}</td>
            <td ng-if="$even">{{x.model}}</td>
            <td>{{x.color}}</td>
        </tr>
    </tbody>
</table>
```

## DOM

`ng-disabled`

`ng-show`

`ng-hide`

## 事件

`ng-click`

## 模块

模块定义了一个应用程序。

模块是应用程序中不同部分的容器。

模块是应用控制器的容器。

控制器通常属于一个模块。

```js
var app = angular.module("myApp", []); 
```

JavaScript 中应避免使用全局函数。因为他们很容易被其他脚本文件覆盖。

AngularJS 模块让所有函数的作用域在该模块下，避免了该问题。

## 表单

`ng-model`来实现数据绑定

## 包含其他HTML文件

`ng-include`

## DI

AngularJS 提供很好的依赖注入机制。以下5个核心组件用来作为依赖注入：

- value
- factory
- service
- provider
- constant

#### value

Value 是一个简单的 javascript 对象，用于向控制器传递值（配置阶段）：

```js
// 定义一个模块
var mainApp = angular.module("mainApp", []);

// 创建 value 对象 "defaultInput" 并传递数据
mainApp.value("defaultInput", 5);
...

// 将 "defaultInput" 注入到控制器
mainApp.controller('CalcController', function($scope, CalcService, defaultInput) {
   $scope.number = defaultInput;
   $scope.result = CalcService.square($scope.number);
   
   $scope.square = function() {
      $scope.result = CalcService.square($scope.number);
   }
});
```

## factory

factory 是一个函数用于返回值。在 service 和 controller 需要时创建。

通常我们使用 factory 函数来计算或返回值。

```js
// 定义一个模块
var mainApp = angular.module("mainApp", []);

// 创建 factory "MathService" 用于两数的乘积 provides a method multiply to return multiplication of two numbers
mainApp.factory('MathService', function() {
   var factory = {};
   
   factory.multiply = function(a, b) {
      return a * b
   }
   return factory;
}); 

// 在 service 中注入 factory "MathService"
mainApp.service('CalcService', function(MathService){
   this.square = function(a) {
      return MathService.multiply(a,a);
   }
});
...
```

 .......

## 路由

 \#! 号之后的内容在向服务端请求时会被浏览器忽略掉。 所以我们就需要在客户端实现 #! 号后面内容的功能实现。 AngularJS 路由就通过 **#! + 标记** 帮助我们区分不同的逻辑页面并将不同的页面绑定到对应的控制器上。 





# 结合项目学习

chanpay-web-boss 项目使用的是 angularjs 1.4.5版本。

