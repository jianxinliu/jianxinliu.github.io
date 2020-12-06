package main

import (
	"fmt"
	"./list"
)

func main()  {
	fmt.Println("...")
	var arr = []int{2,4,6,8,9,5,3}
	l := list.Create(arr)
	l.Print()
	list := l.Reverse1()
	list.Print()
	ll := l.ReverseInstant()
	n := ll
	for n != nil{
		fmt.Println(n.Value)
		n = n.Next
	}
	l.Print()
}