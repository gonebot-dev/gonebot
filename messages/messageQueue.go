package messages

import (
	"container/list"
	"encoding/json"
	"log"
)

var MessageQueue MessageQueueStruct

// Push messgage into a fifo queue with <bufferSize> limit.
func PushMessage(newMsg MessageStruct) {
	MessageQueue.mutex.Lock()
	if MessageQueue.queue.Len() == MessageQueue.bufferSize {
		// queue full
		MessageQueue.queue.Remove(MessageQueue.queue.Front())
	}
	MessageQueue.queue.PushBack(newMsg)
	MessageQueue.mutex.Unlock()

	dNewMsg, _ := json.Marshal(newMsg)
	log.Printf("Receive message: %s\n", dNewMsg)
}
func init() {
	MessageQueue.queue = list.New()
	MessageQueue.bufferSize = 32
}
