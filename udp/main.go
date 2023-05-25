package main

import (
	"fmt"
	"log"
	"net"
)

func handleUDPConnection(conn *net.UDPConn) {
	// 버퍼 : 슬라이스 생성
	buffer := make([]byte, 5)
	// 데이터 받아오기
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal(err)
	}
	// 받아온 데이터
	fmt.Println(addr, buffer, n)

	//클라이언트에게 수신
	msg := []byte("HI")
	_, err = conn.WriteToUDP(msg, addr)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// udp addr 생성
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6000")
	if err != nil {
		log.Fatal(err)
	}

	// 지정된 addr(주소)로 요청을 대기
	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	// 접속 종료(지연 호출)
	defer listener.Close()

	// UDP 수신 시작
	fmt.Println("UDP 6000")
	// 종료 시점 지정 플래그 변수로 지정해볼까?

	count := 0
	for {
		handleUDPConnection(listener)
		count++
		if count > 10 {
			break
		}
	}

}
