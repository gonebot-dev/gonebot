package gonebot

import (
	"gonebot/adaptor"
	"gonebot/plugins"
	"gonebot/processor"
)

func LoadPlugin(plugin plugins.GonePlugin) {
	plugins.LoadPlugin(plugin)
}
func StartBackend(backend string) {
	go processor.MessageProcessor()
	adaptor.StartBackend(backend)
}
