package main

import (
	"fmt"
	"golang_server/Network"
)

func main() {

	fmt.Println("server start")

	// 어떤 통신으로 진행할지 결정
	n := Network.CheckServer()

	switch n {
	case "0":
		fmt.Println("None")
	case "1":
		Network.TcpConn()
	case "2":
		Network.UdpConn()
	case "3":
		Network.WebSocketFunc()
	case "4":
		Network.RestServer().Run("localhost:2023")
	default:
		fmt.Println(n)
	}

}
