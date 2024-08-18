package plugins

import "gonebot/messages"

type GonePlugin struct {
	Name    string
	Command string
	Handler func(msg messages.MessageStruct) messages.MessageStruct
}
