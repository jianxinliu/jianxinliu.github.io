# hibernate 映射知识点

## 1. 一对一关联

一对一关联关系在hibernate中有两种实现方式：
主键关联和外键关联：
1. 主键关联是两张表共用一个主键值，但是这张方式依赖性太强，不推荐使用。
2. 外键关联是两张表有各自的主键，但其中一个表有一个外键引用另一张表的主键。
   此处以外键关联为主，进行介绍。
   例：客户（client）和地址（address）是一对一的关系，Client类的映射文件中的address为外键，
   引用Address类的对应表中的主键。

**表结构**client表里有一个address属性列，其余均为一般属性列
**PO.Address.java**
```JAVA
public class Address{
//省略若干一般属性
private ...;
private Client client;

//省略setter getter 方法
}
```

**PO.Client.java**

```JAVA
public class Client{

//省略若干一般属性

private ...;

private Address address;

//省略setter getter 方法

}

```



**Address.hbm.xml文件部分示例如下：**

```XML
<class name = "Address" table = "address">

<id column = "ID" name = "id" type = "Integer">

<generator....../>

</id>

<property...../>

..

..

<!-- 映射Clent 与Address 的一对一外键关联,双向关联是需要写以下部分，单向关联则不用 -->

<!-- one-to-one标签不维护关联关系，所以需要其他标签来维护关系 -->

<one-to-one name = "clint" class = "PO.Client" property-ref = "address"/>

<!-- 其中的 property-ref属性表示对方哪个属性引用我，此处为Client类中的address属性为外键引用我方的主键 -->

</class>
```



**Client.hbm.xml部分文件示例如下：**

```XML
<class name = "Client" table = "client">

<id column = "ID" name = "id" type = "Integer">

<generator....../>

</id>

<property...../>

..

..

<!-- 映射Clent 与Address 的一对一外键关联,唯一的多对一(unique = "true")，实际上变成了一对一-->

<many-to-one name = "address" class = "PO.Address" column = "address"

cascade = "save-update" unique = "true"/>

</class>
```



**总结**：

1. 一对一关联关系，一般使用外键关联。若是双向关联，则需要在一方使用<one-to-one>标签，

并用**property-ref**属性标明对方的哪个属性引用我方的主键。在另一方使用<many-to-one>标签，并使用**unique="true"**来将外键唯一。

2. 若是单向关联，则不需要使用<one-to-one>标签。
3. 不论是<many-to-one>还是<one-to-one>或是<many-to-many>，可以将to 后面的词换成 who ，然后向 hibernate回答这个问题。必须要回答的是从哪里到哪里。比如此处<many-to-one>，从 Client类中的类型为PO.Address的属性address 到 PO.Address对应表中列名为 address的列。

## 2. 一对多关联关系

以客户（Customer）和订单（order）为例。
**表结构**：在Orders表里有一个CUSTOMER_ID列，其他均为一般属性列

一：单向关联：
单向关联只需要在一的一方进行映射，所以配置Customer
**PO.Customer.java**

```JAVA
public class Customer{

//省略若干一般属性

private ...;

private Set<Orders> 【orders】 = new HashSet<>();

//省略setter getter 方法

}
```

**Customer.hbm.xml**

```XML
如果还需要设置一对多双向关联，那还需要在多的一方使用**<many-to-one>**
```

二：双向关联：
如果还需要设置一对多双向关联，那还需要在多的一方使用**<many-to-one>**

**PO.Orders.java**

```JAVA
public class Customer{

//省略若干一般属性

private ...;

private Customer 【customer】;

//省略setter getter 方法

}
```

**Orders.hbm.xml**

```XML
<class name = "Orders" table = "orders">

<id column = "ID" name = "id" type = "Integer">

<generator....../>

</id>

<property...../>

..

..

<!-- 一对多中多的一方的配置 -->

<many-to-one name = "customer" class = "PO.Customer" 

column = "CUSTOMER_ID"/>

<!--many-to-one means many orders to one customer.who is customer? PO.Customer is!where can I find her? from my froeign key: CUSTOMER_ID-->

</class>
```

## 3. 多对多关联：

在多对多关系中，需要一个第三方表来维护关系，其包含两个字段，分别是两张表的主键。
以 商品（Items）和订单（Orders）为例：
**表结构**两表均为一般属性，都没有外键。第三方表为selected_items(Order_id,Items_id)
**PO.Items.java**

```JAVA
public class Items{

//省略若干一般属性

private ...;

private Set<Orders> 【orders】 = new HashSet<>();

//省略setter getter 方法

}
```

**Item.hbm.xml**

```XML

```

**PO.Orders.java**

```JAVA
public class Orders{

//省略若干一般属性

private ...;

private Set<Items> items = new HashSet<>();

//省略setter getter 方法

}
```

**Orders.hbm.xml**

```XML
<class name = "Orders" table = "orders">

<id column = "ID" name = "id" type = "Integer">

<generator....../>

</id>

<property...../>

..

..

<set name = "items" table = "selected_items">

<key column = "Order_id"/><!-- key标签的column 属性指定第三方表中自己对应的字段列 -->

<many-to-many class = "PO.Orders" column = "Items_id">

<!-- key 中 column 属性的值和 many-to-many 中的 column 属性值刚好相反 -->

</set>

</class>
```

