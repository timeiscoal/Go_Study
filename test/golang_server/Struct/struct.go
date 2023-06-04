package Struct

type ClientData struct {
	AccessSequence int //접근로 시퀀스
	Speed          int //속도
	DetectLine     int //검지된 차선
	Direction      int //이동방향
	Category       int //차량 정보(종류)
	DbHost         string
	DbUser         string
	DbPassword     string
	DbDatabase     string
	DbPort         string
	NetHost        string
	NetMethod      string
	NetPort        string
}
