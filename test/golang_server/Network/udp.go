package Network

import (
	"encoding/json"
	"fmt"
	"golang_server/CheckData"
	"golang_server/Postgres"
	"golang_server/Struct"
	"log"
	"net"
)

func UdpServer() {

	fmt.Println("udp server on")

	// upd addr 생성
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2023")
	if err != nil {
		log.Fatal(err)
	}

	// 지정된 addr로 요청을 대기
	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}

	// main 함수가 종료 되기 직전에 종료
	defer listener.Close()

	requestUdpHandler(listener)

}

func requestUdpHandler(conn *net.UDPConn) {

	fmt.Println("================")
	// 버퍼 생성 (슬라이스 생성)
	buffer := make([]byte, 1024)

	for {
		// 클라이언트에게 데이터 받기.
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		// 구조체 선언 및 초기화
		datas := Struct.ClientData{}

		// 읽은 데이터 선언 및 초기화
		read := buffer[:n]

		// 읽은 데이터를 구조체에 맞게 넣어주기
		errs := json.Unmarshal(read, &datas)
		if err != nil {
			fmt.Println(errs)
		}

		check := CheckData.DataCheck(&datas)
		fmt.Println(check)
		if check {
			fmt.Println("체크 성공")

			Postgres.AccessDb(&datas)
		}
	}

}
