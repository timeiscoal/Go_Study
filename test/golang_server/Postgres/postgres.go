package Postgres

import (
	"fmt"
	"golang_server/Struct"
)

func AccessDb(data *Struct.ClientData) {
	fmt.Println("데이터 베이스 접속")
	fmt.Println(data)
}
