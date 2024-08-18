package main

import (
	"gonebot/backend/onebot11"
	"gonebot/plugins"
)

func LoadPlugin(plugin plugins.GonePlugin) {
	plugins.LoadPlugin(plugin)
}

func Start() {
	onebot11.WebsocketServerInit()
}

func main() {
	Start()
}
