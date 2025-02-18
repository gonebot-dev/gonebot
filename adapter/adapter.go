package adapter

import "github.com/gonebot-dev/gonebot/message"

type GoneAdapter struct {
	Name              string
	Description       string
	Version           string
	SupportedPlatform string // "qq", "wechat", etc.

	Connector func(incomingChan chan message.Message, resultChan chan message.Message) // Main Threading
}

var adapter GoneAdapter

func SetAdapter(a GoneAdapter) {
	adapter = a
}
