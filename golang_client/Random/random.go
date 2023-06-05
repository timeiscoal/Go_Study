package Random

import (
	"math/rand"
	"time"
)

type Car struct {
	AvenueSeq    int // 접근로 시퀀스
	Speed        int // 속도
	LaneNum      int // 감지된 차선
	ModelType    int // 차량(종류)
	MovementType int // 이동방향

}

// 구조체 정의
type carStruct struct {
	data map[string]int
}

// 생성자 함수 정의
func newCarStruct() carStruct {
	d := carStruct{}
	d.data = map[string]int{}
	return d
}

// 랜덤 데이터 생성 메서드 정의
func (c Car) Detect() map[string]int {
	// 생성 간격
	time.Sleep(1 * time.Second)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	// 랜덤 데이터 생성
	c.AvenueSeq = rand.Intn(81) + 20
	c.Speed = rand.Intn(31) + 30
	c.LaneNum = rand.Intn(5) + 1
	c.ModelType = rand.Intn(6) + 2
	c.MovementType = rand.Intn(4) + 1

	car := newCarStruct() // 생성자 호출

	car.data["avenueSeq"] = c.AvenueSeq
	car.data["speed"] = c.Speed
	car.data["laneNum"] = c.LaneNum
	car.data["modelType"] = c.ModelType
	car.data["movementType"] = c.MovementType

	return car.data

}
