package onebot11

import (
	"sync"

	"github.com/gorilla/websocket"
)

func StartBackend() {

	Initialization()
}

func BackendIO(ws *websocket.Conn) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go SendingMessage(ws)
	go ReadingMessage(ws)
	waitGroup.Wait()
}
