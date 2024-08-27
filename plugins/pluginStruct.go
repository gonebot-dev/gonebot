package plugins

import "github.com/gonebot-dev/gonebot/messages"

type GoneHandler struct {
	Command []string                                                //Commands to trigger the handler function.
	Handler func(msg messages.IncomingStruct) messages.ResultStruct //The handler function of the Handler. Will be triggerd by []Command
}

type GonePlugin struct {
	Name        string        //The name of the plugin.
	Description string        //The description of the plugin.
	Handlers    []GoneHandler //Handlers of the plugin.
}
