package gonebot

import (
	"github.com/gonebot-dev/gonebot/adaptor"
	"github.com/gonebot-dev/gonebot/plugins"
	"github.com/gonebot-dev/gonebot/processor"
)

func LoadPlugin(plugin plugins.GonePlugin) {
	plugins.LoadPlugin(plugin)
}
func StartBackend(backend string) {
	go processor.MessageProcessor()
	adaptor.StartBackend(backend)
}
