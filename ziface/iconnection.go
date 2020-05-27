package ziface

import "net"

//定义连接模块抽象接口
type IConnection interface {

	//启动连接：让当前连接开始工作
	Start()
	//停止连接：结束当前连接工作
	Stop()
	//获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	//获取当前连接模块ID
	GetConnID() uint32
	//获取远程客户端的状态
	RemoteAddr() net.Addr
	//将数据发送给远程客户端
	Send([]byte) error
}

//定义链式调用
type HandleFunc func(*net.TCPConn, []byte, int) error
