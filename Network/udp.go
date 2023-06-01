package Network

import (
	"fmt"
	"golang_server/DataBase"
	"golang_server/DataCheck"
	"log"
	"net"
	"sync"
)

var udpWg sync.WaitGroup

func requestUdpHandler(conn *net.UDPConn) {

	// 버퍼 생성 (슬라이스 생성)
	buffer := make([]byte, 1024)

	// 슬라이스 생성 (DB 정보 슬라이스 , 자동차 정보 슬라이스)
	dbSlice := []string{}
	carSlice := []byte{}

	// DB 정보 데이터 받기.
	for {
		// 클라이언트에게 데이터 받기.
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		// 클라이언트에게 받은 데이터를 선언 및 슬라이스에 저장.
		dbdata := (string(buffer[:n]))
		dbSlice = append(dbSlice, dbdata)

		// 클라이언트에서 보내는 데이터의 마지막 데이터 확인.
		endpoint := "END"
		if endpoint == dbdata {
			break
		}

		// 클라이언트에게 메세지 전달.
		msg := []byte("Data Check")
		_, err = conn.WriteToUDP(msg, addr)
		if err != nil {
			log.Fatal(err)
		}

	}

	// 자동차 데이터 받기.
	for {
		// 클라이언트에게 데이터 받기.
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}

		// 클라이언트에게 받은 데이터를 선언 및 슬라이스에 저장.
		carData := buffer[:n]
		carSlice = append(carSlice, carData...)

		// 클라이언트에게 받은 데이터를 모두 받았다면 데이터 체크 진행.
		if len(carSlice) >= 5 {
			dbSlice = dbSlice[:len(dbSlice)-1]
			check := DataCheck.DataHandler(carSlice)
			// 차량 정보가 검증을 통과했다면 DB에 커밋한다.
			if check {
				DataBase.AccessDB(dbSlice, carSlice)
			} else {
				break
			}

		}

	}
	udpWg.Done()

}

func UdpConn() {

	fmt.Println("UDP ON")

	udpWg.Add(1)

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

	udpWg.Wait()

}
