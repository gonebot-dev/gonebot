package onebot11

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var hostAddress string = "localhost:2048"
var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error creating connection:\n%s\n", err)
	}
	defer ws.Close()
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Read message Error:\n%s\n", err)
		}
		log.Printf("Received: %s\nType: %d\n", message, messageType)
	}
}

func WebsocketServerInit() {
	http.HandleFunc("/onebot/v11/ws", socketHandler)
	log.Fatal(http.ListenAndServe(hostAddress, nil))
}
