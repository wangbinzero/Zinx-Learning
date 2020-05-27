package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("connect to server error: ", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello zinx"))
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}

		buf := make([]byte, 512)
		_, err = conn.Read(buf)

		if err != nil {
			fmt.Println("read buf error: ", err)
			return
		}

		fmt.Println("receive message from server: ", string(buf))

		time.Sleep(1 * time.Second)

	}

}
