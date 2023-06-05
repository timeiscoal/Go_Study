package Network

import (
	"encoding/json"
	"fmt"
	"golang_server/CheckData"
	"golang_server/Postgres"
	"golang_server/Struct"
	"net"
	"os"
)

func TcpServer() {

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

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}

	connectionHandler(conn)

}

func connectionHandler(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 1024)

	for {
		// 클라이언트가 보낸 데이터.
		n, err := conn.Read(data)

		if err != nil {
			fmt.Println("", err)
			return
		}
		// 구조체 선언 및 초기화
		datas := Struct.ClientData{}

		// 읽은 데이터 선언 및 초기화
		read := data[:n]

		// 읽은 데이터를 구조체에 맞게 넣어주기
		errs := json.Unmarshal(read, &datas)
		if err != nil {
			fmt.Println(errs)
		}

		check := CheckData.DataCheck(&datas)
		if check {
			fmt.Println("2. 체크 성공")

			Postgres.AccessDb(&datas)
		}

	}

}
