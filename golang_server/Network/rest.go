package Network

import (
	"fmt"
	"golang_server/CheckData"
	"golang_server/Postgres"
	"golang_server/Struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RestServer() *gin.Engine {


	fmt.Println("Rest Server")

	r := gin.Default()
	r.POST("post/", requestRestHandler)

	return r

}

func requestRestHandler(c *gin.Context) {
	

	var data Struct.ClientData
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	check := CheckData.DataCheck(&data)
	fmt.Println(check)
	if check {
		Postgres.AccessDb(&data)

		c.JSON(http.StatusCreated, gin.H{
			"data": data,
		})
	}
}
