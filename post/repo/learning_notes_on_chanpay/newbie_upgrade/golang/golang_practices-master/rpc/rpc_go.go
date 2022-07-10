package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {}


// Hello 必须满足Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。
func (p *HelloService) Hello(req string,reply *string) error {
	*reply = "hello: " + req
	return nil
}

func main(){
	// 将 HelloService 注册为一个 RPC 服务
	rpc.RegisterName("HelloService",new(HelloService))

	listener , err := net.Listen("tcp",":1234")   // 在这个端口监听连接
	if err != nil {
		log.Fatal("ListenTCP err :",err)
	}

	conn ,err := listener.Accept()
	if err != nil {
		log.Fatal("Accpter err:",err)
	}
	rpc.ServeConn(conn)
}