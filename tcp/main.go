package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	// 슬라이스 길이를 들어오는 입력의 크기 만큼 생성할 수 있는 방법이 없을까?
	// 찾아보고는 있지만 아직까지 ...
	data := make([]byte, 100)

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}
		datas := (string(data[:n]))
		endpoint := "END"

		if endpoint == datas {
			fmt.Println("test")

			break
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		_, err = c.Write(data[:n]) // 클라이언트로 데이터를 보냄
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(data[:n])
		
	}

}

func main() {
	ln, err := net.Listen("tcp", ":2023") // TCP 프로토콜에 2023 포트로 연결을 받음
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() // main 함수가 끝나기 직전에 연결 대기를 닫음

	for {
		conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

		go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 실행
	}
}
