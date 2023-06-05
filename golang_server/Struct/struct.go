package Struct

import "time"

type ClientData struct {
	AvenueSeq    int // 접근로 시퀀스
	Speed        int // 속도
	LaneNum      int // 감지된 차선
	ModelType    int // 차량(종류)
	MovementType int // 이동방향
	DbHost       string
	DbUser       string
	DbPassword   string
	DbDatabase   string
	DbPort       string
	NetHost      string
	NetMethod    string
	NetPort      string
	Created      time.Time
}

type MethodType struct {
	Method int
}
