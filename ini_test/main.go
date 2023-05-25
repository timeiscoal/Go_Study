package main

import (
	"fmt"

	"gopkg.in/ini.v1"
	"github.com/lib/pq"
)

// 데이터 읽기 기능 / 필요 데이터는 어떤 메소드 (통신을 할 것인지를 결정하는 문제)

func main() {
	cfg, err := ini.Load("my.ini")

	if err != nil {
		fmt.Println(err)
	}

	dbHost := cfg.Section("DB").Key("Host").String()
	dbUser := cfg.Section("DB").Key("User").String()
	dbPassword := cfg.Section("DB").Key("Password").String()
	dbDatabase := cfg.Section("DB").Key("Database").String()
	dbPort := cfg.Section("DB").Key("Port").String()

	netMethod := cfg.Section("Network").Key("Method").String()
	netHost := cfg.Section("Network").Key("Host").String()
	netPort := cfg.Section("Network").Key("Port").String()

	// 데이터들을 어떤 자료 구조로 담는 것을 좋을까?

	fmt.Println(dbHost, dbUser, dbPassword, dbDatabase, dbPort, netMethod, netHost, netPort)
}
