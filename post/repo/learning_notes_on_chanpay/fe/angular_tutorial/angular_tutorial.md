# Angular学习

## 一。 环境搭建

```
1.  nodejs 和 npm 包管理工具安装

        下载地址 [https://nodejs.org/zh-cn/download/](https://nodejs.org/zh-cn/download/)

        下载完成，一直下一步，就可以安装完成

        X:\Users\sun>node --version
        v10.13.0

        X:\Users\sun>npm --version
        6.4.1
        以上，就表示安装成功
        （安装的其中有一步会自动的在环境变量path下添加 nodejs 的安装目录,若敲上面的命令不成功，可以自己去添加path下添加 nodejs的安装目录）

2. 因为 npm 管理的有些包在墙外，下包的时候会比较慢，或下不下来最好挂vpn，或者使用淘宝镜像

    挂npm 淘宝镜像 的方法 ：[https://blog.csdn.net/quuqu/article/details/64121812](https://blog.csdn.net/quuqu/article/details/64121812)

3. 安装 angular 命令行工具 Angular CLI 

        npm install -g angular-cli 这是安装最新的cli 工具
        npm install -g @angular/cli@6.2.1  这是安装指定版本的cli 工具

        要是想更换cli版本，卸载angular cli的方法：
        npm uninstall -g angular-cli  卸载 angular-cli
        npm uninstall -g @angular/cli 卸载 @angular/cli
        npm cache clean   清除缓存
        若重新打开命令行工具 ng 命令还有效
        删除 X:\Users\sun\AppData\Roaming\npm\node_modules 下的angular-cli 文件夹
        在重新安装
4. 开发工具选择： webStrom 或者 VS code
```

## 二。 Angular简单介绍

**Angular：是一个基于typescript的开源前端web应用程序平台，由Google的Angular团队开发和领导**

typescript：是由微软开发和开源的编程语言，是JavaScript的超集,

```
angular最开始 是叫 AngularJs 是基于JavaScript的开源前端框架，在升级到2.0时，全部使用TypeScript重写后，就改Angular了。

Angular的优点：

    1. 组件模块化
    2. MVC分层
    3. 依赖注入
    4. 自由灵活的路由模块
    5. 强大的命令行工具，能帮助我们快速，构建，添加组件，测试和部署等
    ...
```

```
angular中文文档： [https://angular.cn/docs](https://angular.cn/docs)
angular官方文档： [https://angular.io/](https://angular.io/)
```

## 三。 TypeScript简单介绍：

1. 基础变量类型：

```
布尔型 : bolean  
eg: let isDone: booelan = fasle;

数字型： number  Typescript里所有数字都是浮点数，并且支持二进制，八进制，十六进制
eg: let  id: number = 6;

字符串： string     用' 或者" 表示 （``号是模板字符串能通过 ${} 嵌入值）
eg : let name: string = 'John';
eg : let desc: string = `my name  is ${name} => "my name is john"

数组： [] 或者 Array ,支持泛型
eg： let list: number[] = [1,2,3];
eg:  let list2: Array<number> = [1,2,3];

枚举： enum
eg： enum Color{ Red , Green , Blue }
    let c: Color = Color.Green;

Any: 在编程阶段尚不清楚类型的，可以编译阶段移除类型检查(Any 和 Object 区别：Object类型的变量允许给它任意类型的值，但是不能调用该类型的方法)
eg: let nouSure: any = 4;
    notSure.ifItExits(); //正确

    let obj : Object =4;
    obj.ifitExists();   // 错误

...（null,undefind,never,Object）
```

2.变量的声明

```
let  和 const  ： 

    const 声明的变量被赋值后不能更改

和var 的区别：https://www.tslang.cn/docs/handbook/variable-declarations.html
```

3.装饰器

```
装饰器用来给类添加标注
类装饰器应用于类构造函数，可以用来监视，修改或替换类定义

@sealed  //应用装饰器
class Greeter {
    greeting: string;
    constructor(message: string) {
        this.greeting = message;
    }
    greet() {
        return "Hello, " + this.greeting;
    }
}

//自定义的装饰器（Object.seal 方法是将对象或方法进行封闭）
function sealed(constructor: Function) {
    Object.seal(constructor);
    Object.seal(constructor.prototype);
}
```

Typescript 官网： <https://www.typescriptlang.org/docs/home.html>

typescript 中文网：[ https://www.tslang.cn/](https://www.tslang.cn/)

## 四。进入Angular

### angular架构

简单版：![img](https://i.imgur.com/L0AxgyG.png)

![img](https://i.imgur.com/73Lxvtb.png)

angular主要构成：

**模块：**ngModule 为组件提供编译的上下文环境，@NgModule装饰器函数修饰的typescript类，提供 声明declarations本模块所包含的模块，导入imports，导出exports，本模块向全局提供的服务providers，和启动引导根模块bootstrap的信息，一个Angualr应用由多个ngModule组成，应用至少会有一个用于引导应用的根模块AppComponent, 跟模块提供了用来启动应用的引导机制，通过模块化，可以实现可复用，并且可以实现惰性加载，也就是按需加载的优点

**组件：**component 组件定义模板和样式，以及部分页面时间的方法。每个应用都应该有个根组件，@Component装饰器函数修饰的typescript类 ，并提供包含模板信息的元数据信息

**服务：**service 服务用来处理数据逻辑，服务可以通过依赖注入注入到组件中，服务之间也可以依赖注入 用@Injectable装饰器函数修饰的typescript类，并提供包含依赖注入信息的元数据（）

**模板：**显示页面，模板通过数据绑定来和组件进行数据交互

**路由：**定义在应用的各个不同状态和视图层次结构之间导航时要使用的路径

**指令：** 用来拓展HTML

**管道:** 常用来格式化数据

### 新建一个angular 项目

```
    ng new  app  // 新建一个Angular 应用
    ng new app --routing // 会在新的app上添加一个默认的路由文件

新建的项目的文件目录

--e2e 端到端测试目录（测试用）
--node_modules 第三方依赖包
--src 源代码目录
    --app
    --assert          用来存放静态资源
    --environment     用来配置多环境的目录
    --browserlist     执行的浏览器
    --favicon.ico     图标
    --index.html      根html 系统默认进入的页面
    --kama.conf.js    执行自动化测试
    --main.ts            脚本执行的入口文件
    --polyfills.ts       导入必要的库，是angulaar能运行在老的环境下
    --styles.css         全局的样式
    -- test.ts           执行自动化测试
    --tscnnfig.app.json  typescript编译的配置
    --tscnnfig.spec.json  typescript编译的配置
    --tslint.json         定义typescript代码检查
--.editorconfig          idea生成的配置文件
--.gitignore             git的配置文件
--angular.json           angular cli 工具的配置文件
-- package.json          npm工具配置文件 应用依赖的包
-- package-lock.json     npm工具配置文件 应用依赖的包
-- readme.md            
--tsconfig.json          typescript编译的配置
-- tslint.json           定义typescript代码检查
```

### angular的数据绑定：

![img](https://i.imgur.com/m1Tc2kC.png)

{{ }} --> 是从组件到页面的 数据值 的单向绑定

[ ] --> 是从组件到页面的 属性值 的单向绑定

（）--> 是从页面到组件的 事件 的单向绑定

[( ng-model)] –> 是页面到组件的 属性 的双向绑定 双向绑定一般用于表单中

**属性绑定：[] 事件绑定：()**

双向绑定： 当数据发生变化时，页面数据也会发生改变 eg： {{ hero.name }} <input [(ngModule)] = ‘hero.name’>

单向绑定： 当页面数据发生改变时，后台数据会发生改变，但是不会影响前台其他相同属性的值 eg： {{ hero.name }} <input [ngModule] = ‘hero.name’>

**父子组件间通过单向绑定传值：**

父组件：	<app-hero-detail [hero]="selectedHero"></app-hero-detail> 通过单向属性绑定 传入所选的值

子组件：通过 @Input() hero: Hero; 接受外部组件传入的值

### 服务注入：

当angular发现某个组件依赖某个服务时，会检查注入器中是否已经有了那个服务的任何实例，如果所请求的服务尚不存在，注入器就会使用注册provider 来新建一个，并把它加入到注入器，然后返回实例

实现：

```
通过 ng 命令生成的service 

ng generate service testService
```

注入到root中 @Injectable({ providedIn: 'root' })

**service中Observable：** Observable支持在应用程序中的发布者和订阅者之间传递消息 Observable是声明性的 - 也就是说，定义了一个用于发布值的函数，但是在消费者订阅它之前它不会被执行。订阅的消费者然后接收通知，直到该功能完成，或直到他们取消订阅 Es6 的 => 箭头函数 回调函数：返回一个 Promise（承诺），或者 Observable（可观察）对象 Observable 是 RxJS 库 中的关键类

```
this.heroService.getHeroes().subscribe(heroes => this.heroes = heroes);
```

等价于：

```
var _this = this; 
this.heroService.getHeroes()
.subscribe(function(heroes) {
    _this.heroes = heroes;
});
```

对于第一个heroes就是getHeroes返回的值 ，然后又将 返回的值heroes在传递给this。Heroes，所以此处 heroes就是 一个参数名，取什么都可以，你要是修改为 this.heroService.getHeroes().subscribe(aaa=> this.heroes = aaa); 都行

### 路由：

页面配置：

Index.html 必须添加

```
<base href="/">
```

不然会报错：Error: No base href set. Please provide a value for the APP*BASE*HREF token or add a base element to the document.

```
<nav>
  <a routerLink="/heroes">Heroes</a>  //表示会跳转到 /hero目录下
</nav>
<router-outlet></router-outlet>   //用来显示路由component的占位符
<app-messages></app-messages>
```

路由Typescript类：

```
const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' }, //默认页面重定向到dashboard页面
  { path: 'dashboard', component: DashboardComponent },    //指定dashboard跳转到 dashboard组件页面
  { path: 'detail/:id', component: HeroDetailComponent },  //指定挑转到 detail中某个id 的页面
  { path: 'heroes', component: HeroesComponent }
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
```

可以在component中通过

```
router.navigateByUrl()方法来进行路由
```

### Http模块使用：

通过在service中，，引入 HttpClientModule 并在service 的构造函数中进行注入， get请求

```
getConfig() {
  return this.http.get<Config>(this.configUrl)
    .pipe(
      retry(3), // retry a failed request up to 3 times
      catchError(this.handleError) // then handle the error
    );
}
```

post请求

```
const httpOption = {
  headers: new HttpHeaders({
    'Content-Type':'application/json'
  })
};
addHero (hero: Hero): Observable<Hero> {
  return this.http.post<Hero>(this.heroesUrl, hero, httpOptions)
    .pipe(
      catchError(this.handleError('addHero', hero))
    );
}
```

通用的Http错误处理

```
private handleError(error: HttpErrorResponse) {
  if (error.error instanceof ErrorEvent) {
    // A client-side or network error occurred. Handle it accordingly.
    console.error('An error occurred:', error.error.message);
  } else {
    // The backend returned an unsuccessful response code.
    // The response body may contain clues as to what went wrong,
    console.error(
      `Backend returned code ${error.status}, ` +
      `body was: ${error.error}`);
  }
  // return an observable with a user-facing error message
  return throwError(
    'Something bad happened; please try again later.');
};
```

### 打包发布

```
ng build  自动构建
ng serve  启动应用
ng serve --port 0 --open // 在 支持 ng-zorro后的启动
```

与springboot+maven集成发布：

项目结构应该是：

```
parent---
      --view
        pom.xml
      --web
        pom.xml
      pom.xml
```

思路应该是：

```
1. 先在 view 的pom中 添加 clean插件 当执行 mvnclean时，清空view/dist 的文件夹和static
2. 再在 view 的pom中 添加 font插件 安装node 到view目录下，然后package时执行 build
3. web的 pom中 添加 从 view/dist 复制到 src/resource/static下
4. 打包时 将src/resource/static 也打进去
```

参考：[ http://bohr.me/frontend-backend/](http://bohr.me/frontend-backend/)

使用前端frontend-maven-plugin 插件打包插件：

```
<build>
    <plugins>
        <!-- 当执行mvn clean 时 清空 dist 的目录 -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-clean-plugin</artifactId>
            <version>3.0.0</version>
            <configuration>
                <filesets>
                    <fileset>
                        <!--要清理的目录位置-->
                        <directory>${basedir}/dist</directory>
                        <!--对这些文件进行清理-->
                        <includes>
                            <include>**/*</include>
                        </includes>
                    </fileset>
                </filesets>
            </configuration>
        </plugin>

        <plugin>
            <groupId>com.github.eirslett</groupId>
            <artifactId>frontend-maven-plugin</artifactId>
            <version>1.6</version>
            <executions>
                <execution>
                    <id>install node and npm</id>       <!-- 安装 node 和 npm -->
                    <goals>
                        <goal>install-node-and-npm</goal>
                    </goals>
                    <configuration>
                        <nodeVersion>v8.9.0</nodeVersion>
                        <npmVersion>5.5.1</npmVersion>
                        <nodeDownloadRoot>http://npm.taobao.org/mirrors/node/</nodeDownloadRoot>
                    </configuration>
                </execution>
                <execution>
                    <id>npm run build</id>
                    <goals>
                        <goal>npm</goal>
                    </goals>
                    <configuration>
                        <arguments>run build</arguments>
                    </configuration>
                </execution>
            </executions>
        </plugin>
    </plugins>
</build>
```

整体打包

```
<build>
    <finalName>spm</finalName>
    <plugins>

        <!-- 当执行mvn clean 时 清空 static 的目录 -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-clean-plugin</artifactId>
            <version>3.0.0</version>
            <configuration>
                <filesets>
                    <fileset>
                        <!--要清理的目录位置-->
                        <directory>${basedir}/src/main/resources/static</directory>
                        <!--对这些文件进行清理-->
                        <includes>
                            <include>**/*</include>
                        </includes>
                    </fileset>
                </filesets>
            </configuration>
        </plugin>

        <plugin>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-maven-plugin</artifactId>
        </plugin>

        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-resources-plugin</artifactId>
            <version>3.0.2</version>
            <executions>
                <execution>
                    <id>copy frontend content</id>
                    <phase>package</phase>
                    <goals>
                        <goal>copy-resources</goal>
                    </goals>
                    <configuration>
                        <!-- 复制到的目录 -->
                        <outputDirectory>${basedir}/src/main/resources/static</outputDirectory>
                        <overwrite>true</overwrite>
                        <resources>
                            <resource>
                                <!--  resources插件处理目录 -->
                                <directory>${project.parent.basedir}/spm-view/dist/serviceProviderManage</directory>
                                <includes>
                                    <include>**/*</include>
                                </includes>
                                <filtering>false</filtering>
                            </resource>
                        </resources>
                    </configuration>
                </execution>
            </executions>
        </plugin>
    </plugins>
</build>
```

### 集成JQuery：

```
npm install jquery --save  // --save 是将jquery的依赖添加到当前项目的package.json中
npm install @types/jquery   //  使Typescript能认识 jquery 方法

在 angular.json 中的 styles 和 scripts 添加jquery的css  和 js 的相对路径
```

### 常用的命令

```
ng new  app  // 新建一个Angular 应用
ng new app --routing // 会在新的app上添加一个默认的路由文件
ng generate component app //新建组件app
ng generate service testService  //新建一个服务
ng generate module app-routing --flat --module=app
//在app目录下新建一个route路由配置文件
ng serve 启动Angular
ng add ng-zorro-antd  //添加ng-zorro的ui支持
ng serve --port 0 --open // 在 支持 ng-zorro后的启动
```

### 外部资源：

```
搜索 Angular 相关时  最好 使用 浏览器的 排除查询   将 AngularJs 相关的排除掉    
    angular+springboot -AngularJs

Typescript语法：https://ts.xcatliu.com/introduction/hello-typescript.html  
angualr官方文档：https://angular.io/tutorial/toh-pt0
angualr汉化文档：https://angular.cn/tutorial/toh-pt1
Angular4 视频教程： https://pan.baidu.com/s/1GE0HqB0ITPRDK9_apbeb5w
       提取码：8lsu
Angular服务的注入：https://segmentfault.com/a/1190000015391334
Angular发布摇树优化：http://coyee.com/article/11724-optimize-your-angular-2-application-with-tree-shaking#tip1

ng-zorro ui : https://ng.ant.design/docs/getting-started/zh
```