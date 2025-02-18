package message

import (
	"github.com/gonebot-dev/gonebot/utils"
)

var incomingChan chan Message
var resultChan chan Message

func init() {
	incomingChan = make(chan Message, 128)
	resultChan = make(chan Message, 128)
}

// Push a new message to plugins.
func PushIncomingMsg(newMsg Message) {
	// Drop when full
	if cap(incomingChan) == len(incomingChan) {
		<-incomingChan
	}

	// Push
	incomingChan <- newMsg

	utils.AddIncomingCount()
}

// Pop a new message for plugins to handle.
func GetIncomingMsg() (msg Message) {
	msg := <-incomingChan
	return msg
}

// Push a result message to adapter.
func PushResultMsg(resultMsg Message) {
	if cap(resultChan) == len(resultChan) {
		<-resultChan
	}

	resultChan <- resultMsg
}

// Pop a result message to send.
func GetResultMsg() (msg Message) {
	msg := <-resultChan
	return msg
}

func GetMsgChans() (chan Message, chan Message) {
	return incomingChan, resultChan
}
