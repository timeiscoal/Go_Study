package Postgres

import (
	"database/sql"
	"fmt"
	"golang_server/Struct"
	"time"

	_ "github.com/lib/pq"
)

func AccessDb(data *Struct.ClientData) {
	fmt.Println("3. 데이터 베이스 접속")

	address := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		data.DbHost,
		data.DbPort,
		data.DbUser,
		data.DbPassword,
		data.DbDatabase,
	)

	database, err := sql.Open("postgres", address)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	data.Created = time.Now()

	insert := fmt.Sprintf(
		"INSERT INTO cardata (avenueseq, speed, lanenum, modeltype, movementtype, created)VALUES(%d,%d,%d,%d,%d,'%v');",
		data.AvenueSeq,
		data.Speed,
		data.LaneNum,
		data.ModelType,
		data.MovementType,
		data.Created.Format("2006-01-02 15:04:05"))

	result, err := database.Exec(insert)

	if err != nil {
		fmt.Println("err", err)
	}
	cntAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error", err)
	}

	commit := equal(database, data)
	if commit {
		fmt.Println("4. commit success")
	} else {
		fmt.Println("4. commit false")
	}

	fmt.Println("5. cntAffected", cntAffected)
	fmt.Println("----------------------------")
}

func equal(database *sql.DB, data *Struct.ClientData) bool {

	rows, err := database.Query(
		`select 
		laon.cardata.avenueseq, 
		laon.cardata.speed , 
		laon.cardata.lanenum ,
		laon.cardata.modeltype ,
		laon.cardata.movementtype ,
		laon.cardata.created  
		from laon.laon.cardata 
		where (id) in (select max(id) from laon.laon.cardata);`,
	)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&data.AvenueSeq, &data.Speed, &data.LaneNum, &data.ModelType, &data.MovementType, &data.Created)
		if err != nil {
			panic(err)
		}
	}

	return true
}
