package plugins

import "github.com/gonebot-dev/gonebot/messages"

type GoneHandler struct {
	Command []string
	Handler func(msg messages.MessageStruct) messages.ResultStruct
}

type GonePlugin struct {
	Name        string
	Description string
	Handlers    []GoneHandler
}
