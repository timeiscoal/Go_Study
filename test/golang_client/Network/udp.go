package Network

import (
	"encoding/json"
	"fmt"
	"golang_client/Random"
	"golang_client/Reader"
	"log"
	"net"
)

func UdpClient() {
	// 변수 선언
	var ini Reader.IniInfo
	var cars Random.Car

	i := ini.Reader()

	// tcp port , host , 주소 정의
	address := fmt.Sprintf("%s:%s", i["netHost"], i["netPort"])

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}

	// DialUDP 는 반환된 소켓(혹은 connection)이 지정된 주소로만 전송/수신이 가능하도록 강제하는 역할.
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
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

			_, err = conn.Write([]byte(jsonBytes))
			if err != nil {
				log.Fatal(err)
			}
			// 채널을 활용해서 고루틴 제어
			if count > 10 {
				done <- true
			}

		}
	}()
	<-done
}
