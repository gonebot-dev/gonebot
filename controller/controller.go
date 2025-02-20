package controller

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/plugin"
)

func Start() {

	// Start Plugins
	go plugin.Connector()

	// Start Adapter
	adapter.StartAdapter()
}
