package ziface

//定义路由抽象接口
//路由请求数据都是IRequest格式
type IRouter interface {

	//在处理conn业务之前的钩子方法 hook
	PreHandle(request IRequest)

	//在处理conn业务方法hook
	Handle(request IRequest)

	//在处理conn业务之后的方法hook
	PostHandle(request IRequest)
}
