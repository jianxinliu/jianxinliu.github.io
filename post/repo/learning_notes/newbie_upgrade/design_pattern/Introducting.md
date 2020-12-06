[TOC]

# 设计模式概论

参考：

- [如何通俗理解设计模式及其思想?](https://juejin.im/post/5b3cddb6f265da0f8145c049)
- [设计模式总结之模式分类](https://blog.csdn.net/cooldragon/article/details/52164380)
- [JAVA设计模式总结之23种设计模式](https://www.cnblogs.com/pony1223/p/7608955.html)
- https://java-design-patterns.com/patterns/
- https://design-patterns.readthedocs.io/zh_CN/latest/
- https://so.csdn.net/so/search/s.do?p=1&q=%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F&t=blog&domain=&o=&u=lmj623565791&s=&l=&f=false&rbg=0   常用设计模式参考此处

## 设计模式的分类

| 范围 |                            创建型                            |                            结构型                            |                            行为型                            |
| :--: | :----------------------------------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
|  类  |                  Factory Method（工厂方法）                  |                    Adapter(类) （适配器）                    |      Interpreter（解释器） Template Method（模版方法）       |
| 对象 | Abstract Factory（抽象工厂） Builder（建造者）   Prototype（原型） Singleton（单例） | Bridge（桥接） Composite（组合） Decorator（装饰者） Façade（外观） Flyweight（享元） Proxy（代理） | Chain of Responsibility（职责链） Command（命令） Iterator（迭代器） Mediator（中介者） Memento（备忘录） Observer（观察者） State（状体） Strategy（策略） Visitor（访问者） |

### 细分

|      范围      |                            创建型                            |                    结构型                     |                       行为型                       |
| :------------: | :----------------------------------------------------------: | :-------------------------------------------: | :------------------------------------------------: |
|    对象创建    | Singleton（单例）Prototype（原型）Factory Method（工厂方法）Abstract Factory（抽象工厂）Builder（建造者） |                                               |                                                    |
|    接口适配    |                                                              | Adapter（适配器）Bridge（桥接）Facade（外观） |                                                    |
|    对象去耦    |                                                              |                                               |        Mediator（中介者）Observer（观察者）        |
|    抽象集合    |                                                              |               Composite（组合）               |                 Iterator（迭代器）                 |
|    行为扩展    |                                                              |               Decorator（装饰）               | Visitor（访问者）Chain of Responsibility（职责链） |
|    算法封装    |                                                              |                                               | Template Method（模板方法）Strategy（策略）Command |
| 性能与对象访问 |                                                              |        Flyweight（享元）Proxy（代理）         |                                                    |
|    对象状态    |                                                              |                                               |           Memento（备忘录）State（状态）           |
|      其他      |                                                              |                                               |               Interpreter（解释器）                |

# 6个基本原则

- 单一职责原则（Single Responsibility Principle）

- 里氏代换原则（Liskov Substitution Principle）

- 依赖倒转原则（Dependence Inversion Principle）

- 接口隔离原则（Interface Segregation Principle）

- 迪米特法则，又称最少知道原则（Demeter Principle）

- 开闭原则（Open Close Principle）

