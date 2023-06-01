package DataCheck

import (
	"fmt"
)

// 자동차 정보 구조체 선언
type CarInfo struct {
	AccessSequence int //접근로 시퀀스
	Speed          int //속도
	DetectLine     int //검지된 차선
	Direction      int //이동방향
	Category       int //차량 정보(종류)
	// Created_at     time.Time //데이터 생성시간
}

// 입력받은 데이터 검증
func DataHandler(data []byte) bool {

	fmt.Println("data checking")

	var c CarInfo

	c.AccessSequence = int(data[0])
	c.Speed = int(data[1])
	c.DetectLine = int(data[2])
	c.Direction = int(data[3])
	c.Category = int(data[4])

	switch true {
	case c.AccessSequence < 0 || c.AccessSequence > 100:
		return false
	case c.Speed < 20 || c.Speed > 61:
		return false
	case c.DetectLine <= 0 || c.DetectLine > 6:
		return false
	case c.Direction <= 0 || c.Direction > 5:
		return false
	case c.Category < 2 || c.Category > 8:
		return false
	default:
		return true
	}

}
