package onebot11

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var hostAddress string = "localhost:2048"

func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error creating connection:\n%s\n", err)
	}
	defer ws.Close()
	//Polling messages.
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Read message Error:\n%s\n", err)
		}
		msg := string(message)
		messageHandler(msg)
		//log.Printf("Received: %s\nType: %d\n", message, messageType)
	}
}

func WebsocketServerInit() {
	http.HandleFunc("/onebot/v11/ws", socketHandler)
	log.Fatal(http.ListenAndServe(hostAddress, nil))
}
