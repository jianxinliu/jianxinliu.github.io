package main

import (
	"fmt"
	"log"
)

type DivError struct{
	x,y int
}

// 实现 error 接口
func (DivError)Error() string {
	return "division by zero"
}

func div(x,y int) (int,error) {
	if y == 0 {
		return 0,DivError{x,y}
	}
	return x / y,nil
}
func main()  {
	z,err := div(5,0)

	if err != nil {
		// 按类型匹配时，注意case的顺序，应将自定义类型方在前面，有限匹配更具体的错误类型
		switch e:=err.(type) {
		case DivError:
			fmt.Println("e:",e,e.x,e.y)
		default:
			fmt.Println(e)
		}
		log.Fatalln("log:",err)
	}

	fmt.Println("z:",z)
}