package main

import "fmt"

func main(){

	var x = 1.2     // 1.自动类型推断
	var y int = 3   // 2.定义类型
	z := 8			// 3.定义局部变量的简易方法

	var m,n = 3,"kk"  // 4.解构赋值

	var (
		q int = 3
		p string = "qwe"
	)					// 5. 组合赋值

	// 空变量（占位符）
	var data,_ = return2Var()

	// 6. 编译器检查未使用的变量（notUserVar declared and not used）
	// var notUserVar = 2     

	fmt.Println(x,y,z)
	fmt.Println(m,n)
	fmt.Println(p,q)
	fmt.Println("忽略占位符：",data)

	// ============================= const ======================================== //
		// 常量，编译期可确定值的量
	const X = "hello\n"

	const (
		Y = "world\n"
		Z
		W = len(Y)
	)

	fmt.Println("常量：",X,Y,Z,W)


	// =============================== enum ========================================= //
	const (
		SUNDAY = iota
		MONDAY
		TUESDAY
		WEDNESDAY
		THURDAY
		FRIDAY
		SATURDAY
	)

	const (
		_ = iota
		KB int64 = 1 << (10 * iota)
		MB
		GB
		TB
	)

	fmt.Println("星期：",
		SUNDAY,
		MONDAY,
		TUESDAY,
		WEDNESDAY,
		THURDAY,
		FRIDAY,
		SATURDAY)

	fmt.Println("单位：",
	KB,
	MB,
	GB,
	TB)
}

func return2Var() (int,string){
	var err = "wrong"
	return 2,err
}