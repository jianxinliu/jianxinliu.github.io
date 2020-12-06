package main

import (
	"time"
	. "fmt"
	"runtime"
	)

const number  = 800000
var cpus = runtime.NumCPU()
var times = number/int(cpus)
func main() {

	Println(cpus)
	runtime.GOMAXPROCS(cpus)

	now := time.Now()
	result := doAll()
	end := time.Now()
	Printf("get result concurrency:%d in ",result)
	Println(end.Sub(now))

	now = time.Now()
	result2:=countNormal()
	end = time.Now()
	Printf("get result normal:%d in ",result2)
	Print(end.Sub(now))
}

func countConcurrency(st,end int ,ch chan int){
	var ret = 0
	for i := st; i <= end ; i++ {
		ret += i
	}
	ch <- ret
}

func doAll() (result int) {
	chs := make([] chan int,int(cpus))
	for i := 0; i < int(cpus) ; i++ {
		chs[i] = make(chan int)
		go countConcurrency(1+times*i,times*(i+1),chs[i])
	}
	for _,ch := range chs{
		result += <- ch
	}
	return
}

func countNormal() (ret int) {
	for i := 1; i <= number ; i++ {
		ret += i
	}
	return
}