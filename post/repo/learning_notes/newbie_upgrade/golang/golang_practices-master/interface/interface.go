package main

import . "fmt"

// 难点

func main(){
	var p Printer = &User{54,"jak"}
	p.Print()

	out(3)
	out("hello")
	out(User{53,"kkl"})
}

// 接口，只要实现了接口定义中的所有方法，就实现了这个接口
// 接口是一个方法或多个方法签名的集合


// 空接口 interface{},按照上面的定义，控接口不包含任何方法签名，也就是说所有类型均实现了空接口，相当于 Object 对象

func out(in interface{}){
	Println(in)
}


type Lake interface {
	String() string
}

type Printer interface {
	Lake                  // 接口嵌入
	Print()
}

type User struct {
	age int
	name string
}

func (u * User) String() string{
	return u.name
}

func (u *User) Print() {
	Println(u.age)
}