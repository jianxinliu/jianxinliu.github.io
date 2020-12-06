package main

import "fmt"

func main(){

	// ============================= pointer =========================== //
	var a = 12
	fmt.Println("pointer:",&a,*(&a))

	var stu = new(Student)

	fmt.Println(stu == nil)
	
	stu.name = "jack"

	fmt.Println("student:",stu.name,&stu,*(&stu),*(&(stu.name)))

}

type Student struct{
	name string
	age int
}