package main

import . "fmt"

func main(){
	// 不支持三元运算符

	// 流程控制
	var flag bool = false
	// ===================================== if =================================== //
	if ret := "Nice" ; flag {   // 支持初始化语句和局部变量
		Println(ret + " to")
	}else{
		Println(ret + " from")
	}

	// ======================================= 1. 常见 ============================ //
	for i := 0 ; i < 10 ;i++ {
		Println(i+1)
	}

	// ======================================= 2. 代替 while ========================= // 
	n := 10
	for n > 0 {             // 代替 while(n > 0)         代替 for(;n>0;)
		Println(n)
		n--
	}

	// ================================ 3.range ============================== //
	arr := []int{4,7,1}
	for k,v := range arr {
		Println(k,v)
	}

	str := "hello"
	for k,v := range(str) {
		Println(k,v,string(v))
	}

	str1 := "hello"
	for _,v := range(str1) {
		Println(v,string(v))
	}

	// ==================================== switch ======================== //
	x := []int{3,6,1}
	i := 6
	switch i {
	case x[1] :
		Println("qq")               // 省略 break 默认自动终止
	case 1,3:
		Println("ww")
	default:
		Println("ee")
	}

	na := 2
	switch {         // 当做 if-else if-else 使用
	case na == 1:
		Println("1")
	case na== 2:
		Println("2")
	}

	// goto break continue
}