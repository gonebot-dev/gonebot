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
	defer MessageQueue.mutex.Unlock()
	if MessageQueue.queue.Len() == MessageQueue.bufferSize {
		// queue full
		MessageQueue.queue.Remove(MessageQueue.queue.Front())
	}
	MessageQueue.queue.PushBack(newMsg)

	dNewMsg, _ := json.Marshal(newMsg)
	log.Printf("Receive message: %s\n", dNewMsg)
}
func PopMessage() (MessageStruct, bool) {
	MessageQueue.mutex.Lock()
	defer MessageQueue.mutex.Unlock()
	if MessageQueue.queue.Len() > 0 {
		msg, _ := MessageQueue.queue.Front().Value.(MessageStruct)
		MessageQueue.queue.Remove(MessageQueue.queue.Front())
		return msg, true
	}
	return MessageStruct{}, false
}

func init() {
	MessageQueue.queue = list.New()
	MessageQueue.bufferSize = 32
}
