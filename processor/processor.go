package processor

import (
	"encoding/json"
	"log"

	"github.com/gonebot-dev/gonebot/messages"
	"github.com/gonebot-dev/gonebot/plugins"
)

// The message processor thread.
func MessageProcessor() {
	for {
		msg := messages.PopIncoming()

		result, succ := plugins.TraversePlugins(msg)
		if !succ {
			continue

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
}
