package handler

import "github.com/gonebot-dev/gonebot/message"

type GoneHandler struct {
	Matcher func(msg message.Message) bool
	Handler func(incomingMsg message.Message, resultMsg *message.Message) bool
}

type GoneActiveHandler struct {
	Handler func() message.Message
}
