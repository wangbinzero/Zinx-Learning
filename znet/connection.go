package znet

import (
	"Zinx-Learning/ziface"
	"fmt"
	"net"
)

//定义连接模块
type Connection struct {
	Conn     *net.TCPConn   //当前连接conn
	ConnID   uint32         //连接ID
	isClosed bool           //当前连接状态
	ExitChan chan bool      //通知chan
	Router   ziface.IRouter //该连接处理的方法router
}

//初始化连接对象
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
}

//启动连接：让当前连接开始工作
func (c *Connection) Start() {
	fmt.Println("conn start , connID is : ", c.ConnID)

	//启动从当前连接的读数据业务
	go c.StartRead()
}

//停止连接：结束当前连接工作
func (c *Connection) Stop() {
	fmt.Println("conn stop , connID is : ", c.ConnID)

	//如果当前连接已经关闭
	if c.isClosed {
		return
	}

	c.isClosed = true

	//关闭socket
	c.Conn.Close()
}

//获取当前连接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接模块ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端的状态
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//将数据发送给远程客户端
func (c *Connection) Send([]byte) error {

	return nil
}

//读取套接字方法
func (c *Connection) StartRead() {
	fmt.Println("read goRoutine is running...")

	for {
		//读取客户端的数据到buf中，最大512字节
		buf := make([]byte, 512)

		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive message error: ", err)
			continue
		}

		//TODO 改为调用路由方法
		//调用当前连接锁绑定的handlerFunc
		//err = c.handleAPI(c.Conn, buf, n)
		//if err != nil {
		//	fmt.Println("connID: ", c.ConnID, "handle is error: ", err)
		//	break
		//}
		//获取request数据
		req := Request{conn: c, data: buf}

		//执行注册路由方法
		go func(request ziface.IRequest) {
			//从路由中找到注册绑定的connection 对应的router
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
	defer c.Stop()
	defer fmt.Println("connID is : ", c.ConnID)

}
