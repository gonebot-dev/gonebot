package onebot11

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

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
	BackendHostAddress := os.Getenv("ONEBOT11_HOST")
	if BackendHostAddress == "" {
		BackendHostAddress = "127.0.0.1:21390"
	}

	log.Println("Trying to establish connection with onebot11.")
	http.HandleFunc("/onebot/v11/ws", socketHandler)
	log.Fatal(http.ListenAndServe(BackendHostAddress, nil))
}
