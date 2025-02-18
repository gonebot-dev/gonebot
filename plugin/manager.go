package plugin

import (
	"container/list"
	"log"

	"github.com/gonebot-dev/gonebot/message"
)

var pluginList *list.List = list.New()

// Load a plugin.
func LoadPlugin(plugin GonePlugin) {
	log.Printf("Loading Plugin: %s", plugin.Name)
	pluginList.PushBack(plugin)
}

// Process Message.
func ProcessMsg(rawMsg message.Message) (resultMsg message.Message) {
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

func connWrapper(rawMsg message.Message, resultChan chan message.Message) {
	resultMsg := ProcessMsg(rawMsg)
	resultChan <- resultMsg
}

func activeHandler(resultChan chan message.Message) {
	handlerCount := 0
	for pluginElement := pluginList.Front(); pluginElement != nil; pluginElement = pluginElement.Next() {
		plg, _ := pluginElement.Value.(GonePlugin)
		handlerCount = handlerCount + len(plg.ActiveHandlers)
	}
	if handlerCount == 0 {
		return
	}
	for {
		for pluginElement := pluginList.Front(); pluginElement != nil; pluginElement = pluginElement.Next() {
			plg, _ := pluginElement.Value.(GonePlugin)
			for _, handler := range plg.ActiveHandlers {
				resultChan <- handler.Handler()
			}
		}
	}
}

// Plugin MsgChannel connector.
func Connector(incomingChan chan message.Message, resultChan chan message.Message) {
	for {
		rawMsg := <-incomingChan
		go connWrapper(rawMsg, resultChan)
	}
}
