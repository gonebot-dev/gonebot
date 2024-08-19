package main

import (
	"gonebot/adaptor"
	"gonebot/plugins"
	"gonebot/plugins/builtinplugins"
)

func LoadPlugin(plugin plugins.GonePlugin) {
	plugins.LoadPlugin(plugin)
}

func main() {
	adaptor.UseBackend("onebot11")
	LoadPlugin(builtinplugins.Echo)
}
