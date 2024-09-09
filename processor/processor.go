package processor

import (
	"encoding/json"
	"log"

	"github.com/gonebot-dev/gonebot/messages"
	"github.com/gonebot-dev/gonebot/plugins"
)

func process(msg messages.IncomingStruct) {
	result, succ := plugins.TraversePlugins(msg)
	if !succ {
		return

	}
	//Default reply.
	if result.To == "" {
		result.To = msg.SenderID
	}
	if result.MessageType == "" {
		result.MessageType = msg.MessageType
	}

	messages.PushResult(result)

	dResult, _ := json.Marshal(result)
	log.Printf("Finish Processing: %s\n", dResult)
}

// The message processor thread.
func MessageProcessor() {
	for {
		msg := messages.PopIncoming()
		go process(msg)
	}
}
