package messagehandler

import (
	"container/list"
	"log"

	"github.com/tidwall/gjson"
)

var bufferSize int = 16
var messageQueue *list.List = list.New()

func PushMessage(msg string) {
	log.Printf("Receive message: %s\n", msg)
	var newMsg messageStruct
	newMsg.messageType = gjson.Get(msg, "message_type").String()
	if newMsg.messageType == "private" {
		newMsg.isToMe = true
	}
}
