package Network

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

func ServerCheck(m string) {
	fmt.Println("Check Server")

	client, err := net.Dial("tcp", "127.0.0.1:2023")
	if err != nil {
		fmt.Println(err)
	}

	defer client.Close()

	n, _ := strconv.Atoi(m)

	VehicleCount := map[string]interface{}{
		"method": n,
	}

	jsonBytes, _ := json.Marshal(VehicleCount)
	// 서버로 데이터 전송

	_, err = client.Write([]byte(jsonBytes))
	if err != nil {
		fmt.Println(err)
		return
	}

}
