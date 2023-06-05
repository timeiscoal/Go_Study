package Network

import (
	"encoding/json"
	"fmt"
	"golang_client/Random"
	"golang_client/Reader"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func WebsocketHandler(cars *Random.Car, ini *Reader.IniInfo) {

	fmt.Println("Websocket")

	// 변수 선언

	i := ini.Reader()

	// Websocket port , host 주소 정의
	address := fmt.Sprintf("ws://%s:%s/", i["netHost"], i["netPort"])

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := address
	log.Println(u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	process(c, i, cars)
}

func process(c *websocket.Conn, i map[string]string, cars *Random.Car) {

	done := make(chan bool)
	go func() {
		for count := 0; ; count++ {

			// 검지된 차량 데이터 변수 선언 및 초기화
			car := cars.Detect()

			// 서버로 보낼 데이터 형 변환
			VehicleCount := map[string]interface{}{
				"avenueSeq":    car["avenueSeq"],
				"speed":        car["speed"],
				"laneNum":      car["laneNum"],
				"modelType":    car["modelType"],
				"movementType": car["movementType"],
				"dbHost":       i["dbHost"],
				"dbUser":       i["dbUser"],
				"dbPassword":   i["dbPassword"],
				"dbDatabase":   i["dbDatabase"],
				"dbPort":       i["dbPort"],
				"netHost":      i["netHost"],
				"netMethod":    i["netMethod"],
				"netPort":      i["netPort"],
			}

			jsonBytes, _ := json.Marshal(VehicleCount)

			// 서버로 데이터 전송
			err := c.WriteMessage(websocket.TextMessage, []byte(jsonBytes))
			if err != nil {
				log.Println("Error", err)
				return
			}

			_, p, errs := c.ReadMessage()
			if errs != nil {
				log.Println(err)
				return
			}
			fmt.Println(string(p))
			// 채널을 활용해서 고루틴 제어
			if count > 10 {
				done <- true
			}

		}

	}()
	<-done

}
