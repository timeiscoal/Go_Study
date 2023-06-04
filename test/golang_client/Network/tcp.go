package Network

import (
	"encoding/json"
	"fmt"
	"golang_client/Random"
	"golang_client/Reader"
	"net"
)

/// TCP , UDP 구현 할 예정.
// TCP 통신 우선 구현을 목표로 하자
// 일요일까지 Rest 구현 완료

func TcpClient() {

	// 변수 선언
	var ini Reader.IniInfo
	var cars Random.Car

	i := ini.Reader()

	// tcp port , host , 주소 정의
	address := fmt.Sprintf("%s:%s", i["netHost"], i["netPort"])

	// Dial : 프로토콜, ip주소,포트번호를 설정하여 서버에 연결. TCP프로토콜
	// 127.0.0.1 : 2023
	client, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	// main 함수가 끝나기 직전에 TCP 연결 닫기
	defer client.Close()

	//ini에서 읽은 데이터와 자동차 데이터를 서버로 전송한다.

	done := make(chan bool)
	go func() {
		for count := 0; ; count++ {

			// 검지된 차량 데이터 변수 선언 및 초기화
			car := cars.Detect()

			// 서버로 보낼 데이터 형 변환
			data := map[string]interface{}{
				"accessSequence": car["accessSequence"],
				"category":       car["category"],
				"detectline":     car["detectline"],
				"direction":      car["direction"],
				"speed":          car["speed"],
				"dbHost":         i["dbHost"],
				"dbUser":         i["dbUser"],
				"dbPassword":     i["dbPassword"],
				"dbDatabase":     i["dbDatabase"],
				"dbPort":         i["dbPort"],
				"netHost":        i["netHost"],
				"netMethod":      i["netMethod"],
				"netPort":        i["netPort"],
			}

			jsonBytes, _ := json.Marshal(data)
			// 서버로 데이터 전송
			_, err = client.Write([]byte(jsonBytes))
			if err != nil {
				fmt.Println(err)
				return
			}

			// 채널을 활용해서 고루틴 제어
			if count > 10 {
				done <- true
			}

		}

	}()
	<-done

}
