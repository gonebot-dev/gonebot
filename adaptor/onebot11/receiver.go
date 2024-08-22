package onebot11

import (
	"log"

	"github.com/gonebot-dev/gonebot/messages"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

// This file converts onebot11 json payload into MessageStruct.

// Handle raw websocket payload
func messageHandler(msg string) {
	if !gjson.Valid(msg) {
		log.Printf("Receive invalid JSON.\n")
		return
	}

	postType := gjson.Get(msg, "post_type")
	if postType.Exists() {
		switch postType.String() {
		//Meta event
		case "meta_event":
			metaEventType := gjson.Get(msg, "meta_event_type")
			if metaEventType.Exists() {
				//heartbeat
				if metaEventType.String() == "heartbeat" {
					log.Printf("Receive Heartbeat.\n")
				}
			}
		//message
		case "message":
			messageDecoder(msg)
		}
	}
}

// Format onebot message json and push into fifo queue.
func messageDecoder(rawMessage string) {
	//log.Printf("Receive raw message: %s\n", rawMessage)
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

	//Who send the message?
	newMsg.SenderID = gjson.Get(rawMessage, "sender.user_id").String()
	newMsg.SenderName = gjson.Get(rawMessage, "sender.nickname").String()

	//Who am i?
	newMsg.SelfID = gjson.Get(rawMessage, "self_id").String()

	//Push message into messages queue.
	messages.PushMessage(newMsg)
}

func ReadingMessage(ws *websocket.Conn) {
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Read message Error:\n%s\n", err)
		}
		msg := string(message)
		messageHandler(msg)
	}
}
