package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//单元测试
func TestDataPack(t *testing.T) {

	//创建模拟服务器
	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建go 负责从客户端接收消息
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println(err)
			}

			go func(conn net.Conn) {
				//拆包
				dp := NewDataPack()

				for {
					headData := make([]byte, dp.GetHeadLen())
					//一次读满 读满8字节
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println(err)
						break
					}
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println(err)
						return
					}

					if msgHead.GetMsgLen() > 0 {
						//表示 msg是有数据的，需要进行第二次读取
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println(err)
							return
						}

						fmt.Println("receive msgId: ", msg.Id, "dataLen: ", msg.DataLen, "data: ", string(msg.Data))
					}
				}
			}(conn)
		}

	}()

	//创建客户端

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}

	dp := NewDataPack()
	//模拟粘包过程，封装两个msg一同发送

	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte("zinx"),
	}
	data1, _ := dp.Pack(msg1)

	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte("hellozi"),
	}

	data2, _ := dp.Pack(msg2)

	data1 = append(data1, data2...)

	for {
		conn.Write(data1)
	}

	//客户端阻塞
	select {}
}
