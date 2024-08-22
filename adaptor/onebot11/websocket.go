package onebot11

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var hostAddress string = "localhost:2048"

// The main thread to receive messages.
func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	var err error
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error creating connection:\n%s\n", err)
	}
	log.Printf("Connection Established.\n")
	defer ws.Close()
	BackendIO(ws)
}

func Initialization() {
	http.HandleFunc("/onebot/v11/ws", socketHandler)
	log.Fatal(http.ListenAndServe(hostAddress, nil))
}
