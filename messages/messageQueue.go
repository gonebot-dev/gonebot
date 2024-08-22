package messages

import (
	"encoding/json"
	"log"
)

var MessageChannel chan MessageStruct

// Push messgage into a fifo queue with <bufferSize> limit.
func PushMessage(newMsg MessageStruct) {
	//channel full, drop
	if cap(MessageChannel) == len(MessageChannel) {
		<-MessageChannel
	}
	//push
	MessageChannel <- newMsg

	dNewMsg, _ := json.Marshal(newMsg)
	log.Printf("Receive message: %s\n", dNewMsg)
}
func PopMessage() MessageStruct {
	msg := <-MessageChannel
	return msg

}

func init() {
	MessageChannel = make(chan MessageStruct, 32)
}
