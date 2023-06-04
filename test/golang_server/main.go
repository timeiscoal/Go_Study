package main

import (
	"fmt"
	"golang_server/Network"
)

func main() {

	fmt.Println("t")

	//Network.TcpServer()
	//Network.UdpServer()
	//Network.RestServer().Run("localhost:2023")
	Network.WebSocketFunc()
}
