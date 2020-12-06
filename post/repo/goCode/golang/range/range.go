package main

import "fmt"


/**
实验结论：
使用range遍历数组时，会复制数组的值用于遍历，如果想在遍历时改变数组的值，并且根据改变做相应的操作，那么此时
range 出来的值不是最新的
此时应该使用切片，切片会实时遍历数组的值，而不是其复制的值
*/
func main()  {
	var arri = [3]int{10,20,30}

	/**
	range 数组时，赋给 x 的值一直是原始数组的数据，因为range复制了数组的值
	*/
	for i,x := range arri {
		if i == 0 {
			arri[0] += 100
			arri[1] += 200
			arri[2] += 300
		}
		fmt.Println(x,arri[i],arri)
	}

	fmt.Println(arri)
	fmt.Println()

	/**
	range不会赋值切片的值，所以会赋给 x 的值动态变化
	*/
	for i,x := range arri[:] {
		if i == 0 {
			arri[0] += 100
			arri[1] += 200
			arri[2] += 300
		}
		fmt.Println(x,arri[i],arri)
	}
	fmt.Println(arri)

	/** 10 110 [110 220 330]
		20 220 [110 220 330]
		30 330 [110 220 330]
		[110 220 330]

		110 210 [210 420 630] // 此时是因为 x 取值时，切片还未被改变
		420 420 [210 420 630]
		630 630 [210 420 630]
		[210 420 630]
	*/
}