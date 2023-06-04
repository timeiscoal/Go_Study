package Random

import (
	"math/rand"
	"time"
)

type Car struct {
	accessSequence int // 접근로 시퀀스
	speed          int // 속도
	detectline     int /// 감지된 차선
	direction      int // 이동방향
	category       int // 차량(종류)
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
	c.accessSequence = rand.Intn(81) + 20
	c.speed = rand.Intn(31) + 30
	c.detectline = rand.Intn(5) + 1
	c.direction = rand.Intn(4) + 1
	c.category = rand.Intn(6) + 2

	car := newCarStruct() // 생성자 호출

	car.data["accessSequence"] = c.accessSequence
	car.data["category"] = c.category
	car.data["detectline"] = c.detectline
	car.data["direction"] = c.direction
	car.data["speed"] = c.speed

	return car.data

}
