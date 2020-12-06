package main

import (
	"net/rpc"
	"log"
	"fmt"
)

func main()  {
	// 通过rpc.Dial拨号RPC服务，然后通过client.Call调用具体的RPC方法
	client,err := rpc.Dial("tcp","localhost:1234")     // rpc 调用
	if err != nil {
		log.Fatal("Dial err:",err)
	}
	var reply string
	err = client.Call("HelloService.Hello","jack",&reply)    // 命名空间
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}