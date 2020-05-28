package znet

import "Zinx-Learning/ziface"

//实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写就好了
type BaseRouter struct{}

//这里之所以BaseRouter的方法都为空，目的是因为Router不希望有pre或者post这两个方法业务
//所以router全部继承BaseRouter的好处就是不需要实现pre,post等等
func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

func (br *BaseRouter) Handle(request ziface.IRequest) {

}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}
