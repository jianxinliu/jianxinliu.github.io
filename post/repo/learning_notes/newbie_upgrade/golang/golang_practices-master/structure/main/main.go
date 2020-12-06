package main

import (
	// . "go_code/start/structure"
	. "fmt"
)

type Student struct {
	name string
	age int
	addr string
	teacher struct {
		name string
		age int
		addr string
	}
}

func main(){
	stu := Student{
		name:"jack",
		age:34,
		addr:"China",
	}
	attr := struct {
		name string
		age int
		addr string
		}{"mike",34,"England"}
	stu.teacher = attr

	Println(stu)
}