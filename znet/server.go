package znet

import (
	"errors"
	"fmt"
	"net"
)

//定义Server服务模块
type Server struct {
	Name      string //服务器名称
	IPVersion string //服务器绑定的IP版本
	IP        string //服务器监听的IP
	Port      int    //服务器绑定端口

}

//定义当前客户端连接所绑定的handlerFunc,后面应该由客户端去自定义该方法
func CallBack(conn *net.TCPConn, data []byte, n int) error {
	//回显
	fmt.Println("[callback] message is : ", string(data))
	if _, err := conn.Write(data[:n]); err != nil {
		fmt.Println("write back message error: ", err)
		return errors.New("callback error")
	}
	return nil
}

//初始化
func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}

//-------------------------------------iserver接口方法实现------------------------------------------------

func (s *Server) Start() {
	fmt.Printf("[start server listener at IP: %s Port: %d is starting...]\n", s.IP, s.Port)

	//获取tcp addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error: ", err)
		return
	}

	//监听
	listen, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen: ", s.IPVersion, " error: ", err)
		return
	}

	//等待连接
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}

		//V1.0 读取套接字
		//go func() {
		//	for {
		//		//每次读取512字节
		//		buf := make([]byte, 512)
		//		n, err := conn.Read(buf)
		//		if err != nil {
		//			fmt.Println("receive buf error: ", err)
		//			continue
		//		}
		//
		//		//打印读取到的套接字信息
		//		fmt.Println("receive message from client: ", string(buf))
		//
		//		//回写套接字信息
		//		if _, err := conn.Write(buf[:n]); err != nil {
		//			fmt.Println("write back buf error: ", err)
		//			continue
		//		}
		//	}
		//}()
		var connID uint32
		connID = 0
		//将处理所连接的业务方法和conn进行绑定，得到我们的连接模块
		dealConn := NewConnection(conn, connID, CallBack)
		connID++

		//启动
		go dealConn.Start()
	}

}

func (s *Server) Stop() {

	//释放服务器资源
}

func (s *Server) Serve() {
	s.Start()
	//TODO 此处可以做一些服务启动后的额外操作

	//阻塞处理
	select {}
}
