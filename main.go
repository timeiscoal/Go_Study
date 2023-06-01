package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Method struct {
	x, y int
}

func (m Method) restApi() int {

	fmt.Println("restAPi 선택")

	return m.x + m.y

}

type User struct {
	Name string
	Age  int
}

// 데이터 베이스 테이블
type carCountingInfo struct {
	AccessSequence int //접근로 시퀀스
	Speed          int //속도
	DetectLine     int //검지된 차선
	Direction      int //이동방향
	CarInfo        int //차량 정보(종류)
	// Created_at     time.Time //데이터 생성시간
}

// 통신을 메소드를 통해서 묶어서 조건문을 통해서 실행되도록 구현하자.

func main() {
	newServer().Run()

	rest := Method{1, 2}
	rest.restApi()
}

func newServer() *gin.Engine {
	r := gin.Default()
	r.GET("", helloHandler)
	r.GET("/:name", helloUserHandler)
	r.POST("", testCarData)
	return r
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"responseData": "Hello World",
	})
}

func helloUserHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"greetings": fmt.Sprintf("hello %v", name),
	})
}

// 클라이언트가 넘겨주는 데이터 검증

func helloAccountHandler(c *gin.Context) {
	var car carCountingInfo

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("err: %v", err),
		})
	} else {

		fmt.Println(car)
		if car.Speed == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintln("자동차를 찾을 수 없습니다."),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": car,
			})
		}
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"dataRecevied": user,
	// })

}

// int8은 annot use n (variable of type int8) as int value in map index 오류 발생/ int로 변경하니 잘 됨

func checkData(n, k int) (string, string) {


	// 테이블을 각각 따로 생성
	car := map[int]string{
		2: "소형차",
		3: "소형 트럭",
		4: "대형 트럭",
		5: "소형 버스",
		6: "대형 버스",
		7: "오토바이",
	}

	direction := map[int]string{
		1: "직진",
		2: "좌회전",
		3: "우회전",
		4: "유턴",
	}

	return car[n], direction[k]

}

// 입력받은 데이터가 범주를 벗어나는가 아닌가를 체그해야한다.

func testCarData(c *gin.Context) {

	var car carCountingInfo

	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintln(err),
		})
	} else {
		// 자동차 속도 범위를 벗어났을 경우.
		if car.Speed < 20 || car.Speed > 60 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintln("자동차 속도가 너무 느리거나 빠릅니다. 감지할 수 없습니다."),
			})
		} else if car.CarInfo < 3 || car.CarInfo > 7 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintln("분류되지 않은 차량입니다."),
			})

		} else {
			c.JSON(http.StatusCreated, gin.H{
				"data": car,
			})
		}
	}

	// DB 적용 시간 데이터
	t := time.Now().Format(time.DateTime)

	// 차종 , 이동방향
	var carInfo, direction = checkData(car.CarInfo, car.Direction)

	// 접근로 시퀀스 , 검지된 차선 , 이동속도
	fmt.Println(car.AccessSequence, car.DetectLine, car.Speed)

	fmt.Println(t)
	fmt.Println(carInfo)
	fmt.Println("현재시각 : ", direction)
}
