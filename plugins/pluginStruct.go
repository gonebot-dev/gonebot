package plugins

import "gonebot/messages"

type GoneHandler struct {
	Command string
	Handler func(msg messages.MessageStruct) messages.MessageStruct
}

type GonePlugin struct {
	Name     string
	Handlers []GoneHandler
}