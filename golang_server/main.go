package main

import (
	"fmt"
	"golang_server/Network"
	"time"
)

func Test() {
	main()
}

func main() {

	//Network.TcpServer()

	method := Network.SeverCheck()

	time.Sleep(1 * time.Second)

	switch method {
	case 0:
		fmt.Println("0")
	case 2:
		fmt.Println("2")
		Network.UdpServer()
	case 3:
		fmt.Println("3")
		Network.WebSocketFunc()
	case 4:
		fmt.Println("4")
		Network.RestServer().Run("localhost:2023")
	}

}
