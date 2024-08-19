package processor

import (
	"encoding/json"
	"gonebot/messages"
	"gonebot/plugins"
	"log"
)

// The message processor thread.
func MessageProcessor() {
	for {
		msg, succ := messages.PopMessage()
		if !succ {
			continue
		}

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
