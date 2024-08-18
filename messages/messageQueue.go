package messages

import (
	"container/list"
	"encoding/json"
	"log"
)

var bufferSize int = 32
var MessageQueue *list.List = list.New()

// Push messgage into a fifo queue with <bufferSize> limit.
func PushMessage(newMsg MessageStruct) {
	MessageQueue.PushBack(newMsg)

	dNewMsg, _ := json.Marshal(newMsg)
	log.Printf("Receive message: %s\n", dNewMsg)
}
