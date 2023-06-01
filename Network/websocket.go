package Network

import (
	"fmt"
	"golang_server/DataBase"
	"golang_server/DataCheck"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// upgrader는 http 웹소켓 변환
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var wsWg sync.WaitGroup

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//연결 업그레이드
	conn, err := upgrader.Upgrade(w, r, nil)
	endpoint := "END"

	dbSlice := []string{}
	if err != nil {
		log.Println(err)
		return
	}
	for {
		// 메세지 읽기 DB 정보
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		data := string(p)
		dbSlice = append(dbSlice, data)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		// END 문자열이 넘어오면 종료
		if string(p) == endpoint {
			break
		}
	}
	count := 0

	for {
		// 차량 정보 읽기
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(p)
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		count++
		if count >= 5 {
			check := DataCheck.DataHandler(p)
			fmt.Println(check)
			if check {

				dbSlice = dbSlice[:len(dbSlice)-1]
				DataBase.AccessDB(dbSlice, p)
			} else {
				break
			}
		}

	}
	wsWg.Done()

}

func WebSocketFunc() {
	fmt.Println("WebSocket On")

	wsWg.Add(1)

	// 모든 요청을 처리하기 위해서, localhost에 핸핸핸 핸들러를 마운트 한다.
	log.Panic(http.ListenAndServe("localhost:2023", http.HandlerFunc(wsHandler)))
	wsWg.Wait()
}
