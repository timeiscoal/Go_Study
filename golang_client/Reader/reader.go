package Reader

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// Ini 구조체 정의
type IniInfo struct {
	dbPort     string
	netMethod  string
	netPort    string
	netHost    string
	dbHost     string
	dbUser     string
	dbPassword string
	dbDatabase string
}

// 구조체 정의
type iniStruct struct {
	data map[string]string
}

// 생성자 함수 정의
func newIniStruct() iniStruct {
	d := iniStruct{}
	d.data = map[string]string{}
	return d
}

// ini 파일 읽기]

func (i IniInfo) Reader() map[string]string {

	fmt.Println("Reading")

	cfg, err := ini.Load("./static/laon.ini")
	if err != nil {
		fmt.Println(err)
	}

	// ini 데이터.
	i.dbHost = cfg.Section("DB").Key("Host").String()
	i.dbUser = cfg.Section("DB").Key("User").String()
	i.dbPassword = cfg.Section("DB").Key("Password").String()
	i.dbDatabase = cfg.Section("DB").Key("Database").String()
	i.dbPort = cfg.Section("DB").Key("Port").String()
	i.netHost = cfg.Section("Network").Key("Host").String()
	i.netMethod = cfg.Section("Network").Key("Method").String()
	i.netPort = cfg.Section("Network").Key("Port").String()

	// 생성자 호출
	ini := newIniStruct()

	// map
	ini.data["dbHost"] = i.dbHost
	ini.data["dbUser"] = i.dbUser
	ini.data["dbPassword"] = i.dbPassword
	ini.data["dbDatabase"] = i.dbDatabase
	ini.data["dbPort"] = i.dbPort
	ini.data["netHost"] = i.netHost
	ini.data["netMethod"] = i.netMethod
	ini.data["netPort"] = i.netPort

	return ini.data

}
