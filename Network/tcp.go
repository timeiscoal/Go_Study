package Network

import (
	"fmt"
	"golang_server/DataBase"
	"golang_server/DataCheck"
	"sync"

	"net"
)

var tcpWg sync.WaitGroup

func requestTcpHandler(c net.Conn) {
	// 슬라이스 길이를 들어오는 입력의 크기 만큼 생성할 수 있는 방법이 없을까?
	// 찾아보고는 있지만 아직까지 ...

	//클라이언트에게 받은 DB 데이터
	data := make([]byte, 100)
	dbSlice := []string{}

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}
		// 클라이언트에게 받은 데이터를 슬라이스에 담아준다.
		dbdata := (string(data[:n]))
		dbSlice = append(dbSlice, dbdata)
		// 클라이언트에서 마지막 "END" 메세지를 보내면 for문 정지
		endpoint := "END"

		if endpoint == dbdata {
			fmt.Println("db정보 체크")
			// "END"를 제외한 데이터를 db에 전송

			//클라이언트에게 받은 차량 데이터
			for {
				n, err := c.Read(data)
				if err != nil {
					fmt.Println(err)
					return
				}

				// 데이터 체크
				check := DataCheck.DataHandler(data[:n])
				// 데이터 체크가 true 라면
				if check {
					// dbslice의 마지막 요소 "END"를 뺀 후, AccessDB로 보낸다.
					dbSlice = dbSlice[:len(dbSlice)-1]
					DataBase.AccessDB(dbSlice, data)

				} else {
					break
				}
			}
			break
		}
	}
	tcpWg.Done()
}

func TcpConn() {

	fmt.Println("TCP Server")
	tcpWg.Add(1)

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

		requestTcpHandler(conn)
		break
	}
	tcpWg.Wait()
}
