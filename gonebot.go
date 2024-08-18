package main

import (
	"gonebot/backend"
	"gonebot/plugins"
	"gonebot/plugins/builtinplugins"
)

func LoadPlugin(plugin plugins.GonePlugin) {
	plugins.LoadPlugin(plugin)
}

func Start() {
	backend.Initialization()
}

func main() {
	Start()
	LoadPlugin(builtinplugins.Echo)
}
