package DataBase

import (
	"database/sql"
	"fmt"
	"golang_server/DataCheck"
	"time"

	_ "github.com/lib/pq"
)

type DbInfo struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
	created  time.Time
}

func DataTestI(info []string, data []byte) {
	fmt.Println(info, data)
}

func AccessDB(info []string, data []byte) {

	// db 정보 선언 및 초기화
	var db DbInfo

	db.Host = info[0]
	db.User = info[1]
	db.Password = info[2]
	db.Database = "collection_data" //info[3]
	db.Port = info[4]
	db.created = time.Now()

	// 데이터베이스 연결
	dbaddress := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.Database)

	database, err := sql.Open("postgres", dbaddress)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	// data 구조체 선언
	var c DataCheck.CarInfo
	// 데이터 선언
	c.AccessSequence = int(data[0])
	c.Speed = int(data[1])
	c.DetectLine = int(data[2])
	c.Direction = int(data[3])
	c.Category = int(data[4])

	fmt.Println(c)

	// 데이터 추가하기
	insertData := fmt.Sprintf("INSERT INTO cardata_table (AccessSequence,Speed,detectline,direction_fk,category_fk,created) VALUES(%d,%d,%d,%d,%d,'%v');",
		c.AccessSequence, c.Speed, c.DetectLine, c.Direction, c.Category, db.created.Format("2006-01-02 15:04:05"))

	result, ers := database.Exec(insertData)

	if ers != nil {
		fmt.Println("err:", ers)
	}

	cntAffected, erros := result.RowsAffected()
	if erros != nil {
		fmt.Println("Error", erros)
	}

	fmt.Println("Affected Rows:", cntAffected)

	// 생성한 데이터가 정상적으로 데이터베이스에 적용되었는지 체크
	commit := equal(database, c)

	if commit {
		fmt.Println("commit success")
	} else {
		fmt.Println("commit false")
	}
}

// 데이터가 잘 들어 갔는지 체크
func equal(database *sql.DB, data DataCheck.CarInfo) bool {

	// db 정보 선언 및 초기화
	var flag bool
	var c DataCheck.CarInfo
	// 마지막 인덱스의 데이터 조회
	rows, err := database.Query("select laon.cardata_table.accesssequence,laon.cardata_table.speed ,laon.cardata_table.detectline ,laon.cardata_table.direction_fk ,laon.cardata_table.category_fk from laon.cardata_table where (id) in (select  max(id) from laon.cardata_table);")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	// 조회한 데이터
	for rows.Next() {
		err := rows.Scan(&c.AccessSequence, &c.Speed, &c.DetectLine, &c.Direction, &c.Category)
		if err != nil {
			panic(err)
		}

	}
	if c.AccessSequence == data.AccessSequence && c.Category == data.Category && c.DetectLine == data.DetectLine && c.Direction == data.Direction && c.Speed == data.Speed {
		flag = true
	} else {
		flag = false
	}
	return flag
}
