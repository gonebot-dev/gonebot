package message

import (
	"github.com/gonebot-dev/gonebot/utils"
)

var IncomingChan chan Message
var ResultChan chan Message

func init() {
	IncomingChan = make(chan Message, 128)
	ResultChan = make(chan Message, 128)
}

// Push a new message to plugins.
func PushIncomingMsg(newMsg Message) {
	// Drop when full
	if cap(IncomingChan) == len(IncomingChan) {
		<-IncomingChan
	}

	// Push
	IncomingChan <- newMsg

	utils.AddIncomingCount()
}

// Pop a new message for plugins to handle.
func GetIncomingMsg() Message {
	msg := <-IncomingChan
	return msg
}

// Push a result message to adapter.
func PushResultMsg(resultMsg Message) {
	if cap(ResultChan) == len(ResultChan) {
		<-ResultChan
	}

	ResultChan <- resultMsg
}

// Pop a result message to send.
func GetResultMsg() Message {
	msg := <-ResultChan
	return msg
}
