package builtin

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
	Echo.Command = "echo"
	Echo.Handler = handler
}
