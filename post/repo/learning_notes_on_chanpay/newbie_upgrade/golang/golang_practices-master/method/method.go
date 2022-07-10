package main

import . "fmt"

type Student struct{
	name string
	age int
}

// recevicer
func (stu *Student) introduction() {
	Println(stu.name)
}

func (stu *Student) Hello(){
	Println("hello")
}

func main(){
	stu := Student{"jack",34}
	stu.introduction()   // 隐式传递 receiver

	hello := (*Student).Hello
	hello(&stu)          // 显示传递 receiver
}