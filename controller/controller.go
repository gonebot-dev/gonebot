package controller

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
)

func Start() {

	// Start Plugins
	go plugin.Connector(message.IncomingChan, message.ResultChan)

	// Start Adapter
	adapter.StartAdapter(message.IncomingChan, message.ResultChan)
}
