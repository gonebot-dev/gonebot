package adapter

import (
	"log"

	"github.com/gonebot-dev/gonebot/message"
)

type GoneAdapter struct {
	Name              string
	Description       string
	Version           string
	SupportedPlatform string // "qq", "wechat", etc.

	Connector func(incomingChan chan message.Message, resultChan chan message.Message) // Main Thread
}

var adapter GoneAdapter
var adapterLoaded bool = false

func SetAdapter(a GoneAdapter) {
	adapter = a
	adapterLoaded = true
}

func StartAdapter(incomingChan chan message.Message, resultChan chan message.Message) {
	if !adapterLoaded {
		log.Fatal("No Adapter Loaded!")
	}
	adapter.Connector(incomingChan, resultChan)
}
