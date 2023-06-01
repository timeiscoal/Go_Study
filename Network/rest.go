package Network

import (
	"fmt"
	"golang_server/DataBase"
	"golang_server/DataCheck"
	"net/http"

	"github.com/gin-gonic/gin"
)

type clientData struct {
	DataCheck.CarInfo
	DataBase.DbInfo
}

// restAPI
func RestServer() *gin.Engine {

	fmt.Println("Rest Server")

	r := gin.Default()
	r.POST("post/", requestRestHandler)

	return r
}

func requestRestHandler(c *gin.Context) {
	// 사용자에게 입력받을 데이터를 담을 변수 선언
	var data clientData

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	// 입력 받은 데이터 선언 및 초기화
	// DB정보
	dbHost := data.Host
	dbUser := data.User
	dbPassword := data.Password
	dbDatabase := data.Database
	dbPort := data.Port
	// 자동차 정보
	carAccessSequence := byte(data.CarInfo.AccessSequence)
	carCategory := byte(data.CarInfo.Category)
	carDetectLine := byte(data.CarInfo.DetectLine)
	carDirection := byte(data.CarInfo.Direction)
	carSpeed := byte(data.CarInfo.Speed)

	// 데이터 전달을 위해 슬라이서 생성
	var dbData []string
	var carData []byte
	dbData = append(dbData, dbHost, dbUser, dbPassword, dbDatabase, dbPort)
	carData = append(carData, carAccessSequence, carSpeed, carDetectLine, carDirection, carCategory)

	// 입력 받은 데이터검증
	if DataCheck.DataHandler(carData) {
		// 검증에 성공하면 데이터베이스 커밋
		DataBase.AccessDB(dbData, carData)
		// response
		c.JSON(http.StatusCreated, gin.H{
			"data": data,
		})
	}

}
