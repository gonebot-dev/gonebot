package onebot11

import (
	"gonebot/messages"
	"log"

	"github.com/tidwall/gjson"
)

func messageHandler(msg string) {
	if !gjson.Valid(msg) {
		log.Printf("Receive invalid JSON.\n")
		return
	}
	//Not a message
	MetaEventType := gjson.Get(msg, "meta_event_type")
	if MetaEventType.Exists() {
		//heartbeat
		if MetaEventType.String() == "heartbeat" {
			log.Printf("Receive Heartbeat.\n")
		}
	}
	//Is a message
	messageType := gjson.Get(msg, "message_type")
	if messageType.Exists() {
		messageDecoder(msg)
	}
}

// Format onebot message json and push into fifo queue.
func messageDecoder(rawMessage string) {
	log.Printf("Receive raw message: %s\n", rawMessage)
	var newMsg messages.MessageStruct
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
	newMsg.Text = ""
	textMessages := gjson.Get(rawMessage, "message.#(type==\"text\")#.data.text")
	textMessages.ForEach(func(_, value gjson.Result) bool {
		newMsg.Text += value.String()
		return true // keep iterating, gjson
	})
	messages.PushMessage(newMsg)
}
