package onebot11

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func StartBackend() {
	log.Printf("Using Backend Onebot11\n")
	Initialization()
}

func BackendIO(ws *websocket.Conn) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go SendingMessage(ws)
	go ReadingMessage(ws)
	waitGroup.Wait()
}
