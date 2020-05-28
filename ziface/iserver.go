package ziface

type IServer interface {
	//启动
	Start()
	//停止
	Stop()
	//运行
	Serve()
	//给当前服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(router IRouter)
}
