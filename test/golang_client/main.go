package main

import (
	"fmt"
	"golang_client/Network"
	"golang_client/Reader"
)

func main() {

	var ini Reader.IniInfo

	i := ini.Reader()

	switch i["netMethod"] {
	case "0":
		fmt.Println("0")
	case "1":
		fmt.Println("1")
		Network.TcpClient()
	case "2":
		fmt.Println("2")
		Network.UdpClient()
	case "3":
		fmt.Println("3")
		Network.WebsocketHandler()
	case "4":
		fmt.Println("4")
		Network.RestClient()
	}

	// go Network.TcpClient()

	//위 예시에서 s1 객체는 생성자 함수로 data 필드의 맵을 초기화했기 때문에 바로 data 필드에 값을 저장할 수 있습니다.
	//하지만 s2 객체는 구조체만 생성했기 때문에 data 필드에 값을 저장하기 위해 선언이 필요한 맵은 따로 초기화해야 합니다.
	//이렇게 생성자 함수를 사용하면 구조체의 사용이 훨씬 수월해질 수 있습니다. 형태를 익히기 바랍니다.
}
