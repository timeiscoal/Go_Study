package CheckData

import (
	"fmt"
	"golang_server/Struct"
)

func DataCheck(data *Struct.ClientData) bool {
	fmt.Println("1. 데이터 체크")
	switch true {
	case data.AvenueSeq < 0 || data.AvenueSeq > 100:
		return false
	case data.Speed < 20 || data.Speed > 61:
		return false
	case data.LaneNum <= 0 || data.LaneNum > 6:
		return false
	case data.ModelType < 2 || data.ModelType > 8:
		return false
	case data.MovementType <= 0 || data.MovementType > 5:
		return false
	default:
		return true
	}

}
