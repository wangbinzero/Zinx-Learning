package main

import "Zinx-Learning/znet"

func main() {

	s := znet.NewServer("zinx v0.1")
	s.Serve()
}
