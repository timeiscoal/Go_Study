package Network

import (
	"encoding/json"
	"fmt"
	"golang_server/Struct"
	"net"
	"os"
)

func SeverCheck() int {

	// Resolve the string address to a TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:2023")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Start listening for TCP connections on the given address
	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Accept new connections
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}
	// Handle new connections in a Goroutine for concurrency

	method := severCheckHandler(conn)

	if method == 1 {
		for {
			// Accept new connections
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
			}
			// Handle new connections in a Goroutine for concurrency

			go connectionHandler(conn)
		}
	} else {
		listener.Close()
		return method

	}

}

func severCheckHandler(conn net.Conn) int {

	defer conn.Close()

	data := make([]byte, 1024)

	// 클라이언트가 보낸 데이터.
	n, err := conn.Read(data)

	if err != nil {
		panic(err)
	}
	// 구조체 선언 및 초기화
	datas := Struct.MethodType{}

	// 읽은 데이터 선언 및 초기화
	read := data[:n]

	// 읽은 데이터를 구조체에 맞게 넣어주기
	errs := json.Unmarshal(read, &datas)
	if err != nil {
		panic(errs)
	}
	return datas.Method

}
