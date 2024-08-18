package processor

import (
	"gonebot/messages"
	"gonebot/plugins"
)

// The message processor thread.
func MessageProcessor() {
	for {
		msg, succ := messages.PopMessage()
		if !succ {
			continue
		}
		plugins.TraversePlugins(msg)
	}
}
