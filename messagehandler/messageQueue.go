package messagehandler

import (
	"container/list"
	"encoding/json"
	"log"

	"github.com/tidwall/gjson"
)

var bufferSize int = 16
var messageQueue *list.List = list.New()

func PushMessage(rawMessage string) {
	//Format onebot message json and push into fifo queue.

	//log.Printf("Receive raw message: %s\n", rawMessage)
	var newMsg messageStruct
	//Is private message?
	newMsg.MessageType = gjson.Get(rawMessage, "message_type").String()
	if newMsg.MessageType == "private" {
		newMsg.IsToMe = true
	}
	//Is to me?
	selfID := gjson.Get(rawMessage, "self_id").String()
	atUsers := gjson.GetMany(rawMessage, "message.#(type==\"at\")#.data.qq")
	for _, value := range atUsers {
		if value.String() == selfID {
			newMsg.IsToMe = true
		}
	}

	//Extract all text from message.
	textMessages := gjson.GetMany(rawMessage, "message.#(type==\"text\")#.data.text")
	newMsg.Text = ""
	for _, value := range textMessages {
		newMsg.Text += value.String()
	}

	//TODO: unicode convert

	dNewMsg, _ := json.Marshal(newMsg)
	log.Printf("Receive message: %s\n", dNewMsg)
}
