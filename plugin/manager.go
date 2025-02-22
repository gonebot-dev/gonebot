package plugin

import (
	"container/list"
	"fmt"
	"log/slog"
	"time"

	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin/handler"
)

var pluginList *list.List = list.New()

// Load a plugin.
func LoadPlugin(plugin GonePlugin) {
	slog.Info(fmt.Sprintf("Loading Plugin: %s", plugin.Name))
	pluginList.PushBack(plugin)
}

// Process Message.
func ProcessMsg(rawMsg message.Message) (resultMsg message.Message) {
	// Init resultMsg's metadata.
	resultMsg.SenderID = rawMsg.ReceiverID
	resultMsg.ReceiverID = rawMsg.SenderID

	for pluginElement := pluginList.Front(); pluginElement != nil; pluginElement = pluginElement.Next() {
		plg, _ := pluginElement.Value.(GonePlugin)
		for _, handler := range plg.Handlers {
			if handler.Matcher(rawMsg) {
				if handler.Handler(rawMsg, &resultMsg) {
					break
				}
			}
		}
	}
	return resultMsg
}

func connWrapper(rawMsg message.Message) {
	resultMsg := ProcessMsg(rawMsg)
	if len(resultMsg.Segments) > 0 {
		message.PushResultMsg(resultMsg)
	}
}

func activeHandlerWrapper(handler handler.GoneActiveHandler) {
	for {
		message.PushResultMsg(handler.Handler())
		// Sleep 1s to avoid bad plugin's performance issue
		time.Sleep(time.Second)
	}
}

func activeHandler() {
	for pluginElement := pluginList.Front(); pluginElement != nil; pluginElement = pluginElement.Next() {
		plg, _ := pluginElement.Value.(GonePlugin)
		for _, handler := range plg.ActiveHandlers {
			go activeHandlerWrapper(handler)
		}
	}
}

// Connecting Plugin to MessageChannel
func Connector() {
	// Start Active Handler
	go activeHandler()

	// Message Handler
	for {
		rawMsg := message.GetIncomingMsg()
		go connWrapper(rawMsg)
	}
}
