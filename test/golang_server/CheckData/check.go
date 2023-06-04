package CheckData

import (
	"fmt"
	"golang_server/Struct"
)

func DataCheck(data *Struct.ClientData) bool {
	fmt.Println("데이터 체크")
	switch true {
	case data.AccessSequence < 0 || data.AccessSequence > 100:
		return false
	case data.Speed < 20 || data.Speed > 61:
		return false
	case data.DetectLine <= 0 || data.DetectLine > 6:
		return false
	case data.Direction <= 0 || data.Direction > 5:
		return false
	case data.Category < 2 || data.Category > 8:
		return false
	default:
		return true
	}

}
