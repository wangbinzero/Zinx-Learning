package ziface

//封装客户端的请求信息，和请求连接
type IRequest interface {

	//获取当前连接
	GetConnection() IConnection
	//获取请求数据
	GetData() []byte

	//得到当前请求消息ID
	GetMsgId() uint32
}
