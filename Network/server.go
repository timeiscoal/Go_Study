package Network

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 어떤 통신을 진행할 지 결정
func CheckServer() string {

	fmt.Println("Check Server")

	// 리스너 생성
	li, err := net.Listen("tcp", ":2023")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// 연결 승락
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		data := handle(conn)

		// 통신 방법 리턴
		return data
	}

}

// Reader
func handle(conn net.Conn) string {
	defer conn.Close()
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)

		}
		// 통신 방법 리턴
		return data

	}

}
