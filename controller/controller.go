package controller

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
)

func Start() {
	incomingChan, resultChan := message.GetMsgChans()

	// Start Plugins
	go plugin.Connector(incomingChan, resultChan)

	// Start Adapter
	adapter.StartAdapter(incomingChan, resultChan)
}
