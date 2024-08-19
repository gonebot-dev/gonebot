package main

import (
	"gonebot/adaptor"
	"gonebot/plugins"
	"gonebot/plugins/builtinplugins"
	"gonebot/processor"
)

func LoadPlugin(plugin plugins.GonePlugin) {
	plugins.LoadPlugin(plugin)
}
func StartBackend(backend string) {
	go processor.MessageProcessor()
	adaptor.StartBackend(backend)
}

func main() {
	LoadPlugin(builtinplugins.Echo)

	StartBackend("onebot11")
}
