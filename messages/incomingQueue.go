package messages

import (
	"encoding/json"
	"log"
)

var MessageChannel chan IncomingStruct

// Push messgage into a fifo queue with <bufferSize> limit.
func PushIncoming(newMsg IncomingStruct) {
	//channel full, drop
	if cap(MessageChannel) == len(MessageChannel) {
		<-MessageChannel
	}
	//push
	MessageChannel <- newMsg

	dNewMsg, _ := json.Marshal(newMsg)
	log.Printf("Receive message: %s\n", dNewMsg)
}
func PopIncoming() IncomingStruct {
	msg := <-MessageChannel
	return msg

}

func init() {
	MessageChannel = make(chan IncomingStruct, 32)
}