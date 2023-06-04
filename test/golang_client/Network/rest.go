package Network

import (
	"fmt"
	"golang_client/Random"
	"golang_client/Reader"

	"github.com/go-resty/resty/v2"
)

func RestClient() {

	// 변수 선언
	var ini Reader.IniInfo
	var cars Random.Car

	i := ini.Reader()
	restApi := "post"

	// tcp port , host , 주소 정의
	address := fmt.Sprintf("http://%s:%s/%s/", i["netHost"], i["netPort"], restApi)

	// post 요청
	client := resty.New()

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

	_, err := client.R().SetBody(data).Post(address)
	if err != nil {
		panic(err)
	}
}
