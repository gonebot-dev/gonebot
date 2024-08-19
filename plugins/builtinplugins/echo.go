package builtinplugins

import (
	"gonebot/messages"
	"gonebot/plugins"
)

var Echo plugins.GonePlugin

func handler(msg messages.MessageStruct) messages.MessageStruct {
	return msg
}

func init() {
	Echo.Name = "echo"
	echoHandler := plugins.GoneHandler{Command: "echo", Handler: handler}
	Echo.Handlers = append(Echo.Handlers, echoHandler)
}
