package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

// 데이터 베이스 테이블
type carCountingInfo struct {
	AccessSequence int8 //접근로 시퀀스
	Speed          int8 //속도
	DetectLine     int8 //검지된 차선
	Direction      int8 //이동방향
	CarInfo        int8 //차량 정보(종류)
	// Created_at     time.Time //데이터 생성시간
}

func main() {
	newServer().Run()

}

func newServer() *gin.Engine {
	r := gin.Default()
	r.GET("", helloHandler)
	r.GET("/:name", helloUserHandler)
	r.POST("", helloAccountHandler)
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
	fmt.Println(car.CarInfo, car.Direction)
	fmt.Println(car)

	// c.JSON(http.StatusOK, gin.H{
	// 	"dataRecevied": user,
	// })

}

// func testPostHandler(c *gin.Context) {

// }
