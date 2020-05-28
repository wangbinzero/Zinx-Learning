package ziface

type IDataPack interface {

	//获取包头的长度
	GetHeadLen() uint32

	//装包
	Pack(msg IMessage) ([]byte, error)

	//拆包
	UnPack([]byte) (IMessage, error)
}
