# Java 中 static 修饰的方法不能调用非 static 修饰的方法的原因分析

## 方法调用

对于 java 来说，调用成员变量，方法时，主调是必不可少的，即便代码中省略了主调，实际上主调依然存在。省略主调调用static修饰的方法，此时的主调默认为当前类。省略主调调用非static修饰的方法，默认主调为this。

所以在static修饰的方法中省略主调得调用方法

- 调用static方法：默认用当前类作为主调
- 调用非static方法：默认用this作为主调

显然第一种情况可以成功。第二种情况有这么个问题，在static修饰的方法中，`this`指向谁？`this`不能指向合适的对象，因为实例化这个类的对象不止一个。

```Java
public class Test{
    public void info(){
        //do something
    }
    publuc static void info1(){
        //do something
    }
    public static void main(String[] args){
        //省略的this无法指向有效的对象
        info();//省略this这个主调，直接调用方法
    }
}
```

## 内存角度分析

static修饰的方法不能调用非static修饰的方法，是因为这两个方法在内存中不同的位置，且调用时不能指明主调。

在类进行加载时，static修饰的方法先载入内存，如上例中的 `info1()`方法，当对这个类进行实例化时，`info()`方法随实例分配到了另一块内存。此时在 `info1()`方法中省略主调，直接调用 `info()`方法，自然是找不到的。正确的做法是实例化这个对象（确认主调）再调用(`new Test().info()`)。
