package controller

import (
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
)

func Start() {
	incomingChan, resultChan := message.GetMsgChans()

	// Start Plugins
	go plugin.Connector(incomingChan, resultChan)
}
