package main

import "fmt"

/**
二级指针
	/ === 口袋的地址： 0xc042074018
	================= 开始寻宝 ==================
	/ === 线人说宝物的地址在： 0xc042054058
	/ === 宝物是： 100
	/ === 宝物的地址确实是： 0xc042054058
	/ === 口袋的地址还是： 0xc042074018
*/
func main()  {
	var p *int // 一个口袋。。。

	fmt.Println("/ === 口袋的地址：",&p)
	fmt.Println("================= 开始寻宝 ==================")

	test(&p) // 1.给你一个口袋的地址


	// *p 口袋里的宝物
	fmt.Println("/ === 宝物是：",*p) // 4. 很好，我自己根据口袋里的地址去找宝物吧
	// p 宝物的地址
	fmt.Println("/ === 宝物的地址确实是：",p)
	// &p 口袋的地址
	fmt.Println("/ === 口袋的地址还是：",&p)
}

// 二级指针，装指针变量的变量，应该从调用开始看
func test(p **int){
	// 2.到达这个口袋的地址
	x := 100
	fmt.Println("/ === 线人说宝物的地址在：",&x)
	// 3.往口袋里放了宝物的地址
	*p = &x
}