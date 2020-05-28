package znet

import (
	"Zinx-Learning/utils"
	"Zinx-Learning/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct {
}

func (dp *DataPack) GetHeadLen() uint32 {

	//Id uint32  4byte
	//HeadLen uint32 4byte
	return 8
}

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

	buf := bytes.NewBuffer([]byte{})

	//将dataLen写进buf中  关注点：大小端
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	//将id写进buf中
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//将data数据写进buf中
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

func (dp *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	buf := bytes.NewReader(data)

	//只解压head信息，得到dataLen
	msg := &Message{}

	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断dataLen是否已经超出了我们允许的最大长度
	if utils.BaseConfig.MaxPackageSize > 0 && msg.DataLen > utils.BaseConfig.MaxPackageSize {
		return nil, errors.New("too large msg data receive")
	}

	return msg, nil

}

//初始化
func NewDataPack() *DataPack {
	return &DataPack{}
}
