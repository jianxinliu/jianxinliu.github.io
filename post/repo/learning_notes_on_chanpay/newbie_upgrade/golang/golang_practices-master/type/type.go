package main

import "fmt"

func main(){

	// ================================= array ============================ //
	arr := [3]int{3,5,7}

	arr1 := make([]int,4) // 返回对象。编译器翻译为具体的创建函数，分配内存、初始化成员结构，返回创建的对象

	arr2 := new([3]int)   // 返回指针。计算类型大小，分配零值内存，返回指针
	arr2[1] = 10

	fmt.Println("arr:",arr)   // arr: [3 5 7]
	fmt.Println("arr1:",arr1) // arr1: [0 0 0 0]
	fmt.Println("arr2:",arr2) // arr2: &[0 0 0]
	fmt.Println("arr2:",arr2[1]) // arr2: &[0 0 0]

	str5 := "jack"
	fmt.Println("slice:",str5[:2],str5[:],str5[2:],str5[1:3])  // slice 作为动态数组存在

	// ==================================== string ============================= //
	str := "hello"
	// 可以shiy9ong索引访问字符串内部字符
	// 只能显式类型转换
	fmt.Println(str,str[2],string(str[2]),str[2] == 'l')

	str2 := "hello " +  // 拼接字符串时，+ 必须在上一行末尾。否则编译出错。go 语言对部分编程风格进行编译级别的限制
		"world"

	fmt.Println(str2)


	type myType string

	var abc myType = "jxl"

	fmt.Println(abc)

	type student struct{
		name string
		age int
	}

	// ======================================= map ============================================ //
	kv := make(map[int]string,2)

	kv[1] = "qw"
	kv[3] = "3r"

	fmt.Println(kv)
	fmt.Println(kv[3])

	kvmap := map[int]struct{
		name string
		age int
	}{
		1:{name:"jack",age:23},
		2:{name:"robin",age:45},
	}

	fmt.Println(kvmap)
	fmt.Println(kvmap[2])

	for _,v := range kvmap {
		fmt.Println(v)
	}
}