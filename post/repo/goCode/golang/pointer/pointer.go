package main

import "fmt"
func main()  {
	var a = 3
	var ap = &a // ap 是 a 的地址值
	var app *int = &a // app 是保存 a 的地址的指针变量
	// 值是一样的，就是说对地址值取值还是对地址变量取值，得到的都是目标对象的值
	// 只不过使用指针变量取值多了一个取值的步骤，先要得到指针变量的值（就是地址）再将这个地址拿去取值。
	fmt.Println(ap,*ap) // 0xc042010080 3   
	fmt.Println(app,*app) // 0xc042010080 3

	*app+=20 // 间接更新对象值

	fmt.Println(app,*app) // 0xc042010080 23
}