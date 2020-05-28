package main

import (
	"Zinx-Learning/ziface"
	"Zinx-Learning/znet"
	"fmt"
)

func main() {

	s := znet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}

//自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("custom preHandle router")
	request.GetConnection().GetTCPConnection().Write([]byte("before ping"))
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("custom handle router")
	request.GetConnection().GetTCPConnection().Write([]byte("handle ping"))
}

func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("custom postHandle router")
}
