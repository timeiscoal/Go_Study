package Network

import (
	"fmt"
	"golang_client/Random"
	"golang_client/Reader"

	"github.com/go-resty/resty/v2"
)

func RestClient(cars *Random.Car, ini *Reader.IniInfo) {

	fmt.Println("Rest")

	i := ini.Reader()
	restApi := "post"

	// tcp port , host , 주소 정의
	address := fmt.Sprintf("http://%s:%s/%s/", i["netHost"], i["netPort"], restApi)

	client := resty.New()

	// post 요청

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

			_, err := client.R().SetBody(VehicleCount).Post(address)
			if err != nil {
				panic(err)
			}

			if count > 10 {
				done <- true
			}
		}
	}()
	<-done
}
