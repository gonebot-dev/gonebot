package gonebot

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/controller"
	"github.com/gonebot-dev/gonebot/plugin"
)

func LoadAdapter(a adapter.GoneAdapter) {
	adapter.SetAdapter(a)
}

func LoadPlugin(p plugin.GonePlugin) {
	plugin.LoadPlugin(p)
}

func Start() {
	controller.Start()
}
