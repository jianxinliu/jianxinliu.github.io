package main

import "fmt"

// main 函数就是以 goroutine 的形式运行的   协程coroutine，运行在用户态的用户线程
// 配合 channel 来实现“以通讯来共享内存”的 CSP 模式（https://www.jianshu.com/p/36e246c6153d）
func main(){

	data := make(chan int,3)      // 数据队列，默认为阻塞，同步模式，当指定缓冲区大小之后，缓冲区满则写被阻塞，缓冲区空则读被阻塞，也就是异步模式
	flag := make(chan bool,1)     // 退出信号队列，阻塞


	// read := make(chan<- int)     // 只读
	// write := make(chan-> int)    // 只写

	go func() {			          // 以协程的方式运行这个函数
		fmt.Println("reciver start....")
		for d := range data{      // 协程阻塞在数据队列，队列中有数据就读取，直到队列关闭
			fmt.Println(d)
		}
		fmt.Println("recv over")
		flag <- true              // 数据读完，发送关闭信号，在此之前，需要 main 线程等待协程完成工作而不提前退出
	}()


	// 另一个协程进行推数据
	go func(){
		fmt.Println("sender start....")
		data <- 1
		data <- 2
		data <- 3
		data <- 4
		data <- 5

		/**
		* 预设队列长度为 3 ，向队列推 3 个数据时的输出
		sender start....
		send over                 // 队列满之前，读取被阻塞
		reciver start....
		1
		2
		3
		recv over

		* 预设队列长度为 3 ，向队列推 5 个数据时的输出
		sender start....
		reciver start....
		1
		2
		3
		4
		send over
		5
		recv over

		* 当不预设队列长度时的输出
		sender start....
		reciver start....                    // 默认队列长度为 1 ，发一个读一个
		1
		2
		3
		4
		5
		send over
		recv over
		*/

		close(data)            
		fmt.Println("send over")
	}()

	// main 开始向数据队列推数据，同时，协程开始读数据
	// data <- 1
	// data <- 2
	// data <- 3

	// close(data)                   // main 关闭数据队列，协程退出
	// fmt.Println("send over")

	<- flag                       // main 等待退出的信号，main 被阻塞

}


func init(){
	fmt.Println("init ....")
}