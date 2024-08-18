package messages

import (
	"container/list"
	"encoding/json"
	"log"
)

var bufferSize int = 16
var messageQueue *list.List = list.New()

func PushMessage(newMsg MessageStruct) {
	//Push message into queue.

	dNewMsg, _ := json.Marshal(newMsg)
	log.Printf("Receive message: %s\n", dNewMsg)
}
