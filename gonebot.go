package gonebot

import (
	"log/slog"

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
	slog.Info("Starting Gonebot")
	controller.Start()
}
