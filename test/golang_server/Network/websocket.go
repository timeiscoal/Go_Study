package Network

import (
	"encoding/json"
	"fmt"
	"golang_server/CheckData"
	"golang_server/Postgres"
	"golang_server/Struct"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	//연결 업그레이드
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	for {
		// 메세지 읽기
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		var data Struct.ClientData

		errs := json.Unmarshal(p, &data)
		if err != nil {
			fmt.Println(errs)
		}

		check := CheckData.DataCheck(&data)
		if check {
			fmt.Println("체크 성공")

			Postgres.AccessDb(&data)
		}
	}
}

func WebSocketFunc() {
	fmt.Println("WebSocket On")

	// 모든 요청을 처리하기 위해서, localhost에 핸핸핸 핸들러를 마운트 한다.
	log.Panic(http.ListenAndServe("localhost:2023", http.HandlerFunc(wsHandler)))

}
