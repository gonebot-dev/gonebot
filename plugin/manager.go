package plugin

import (
	"container/list"
	"log"

	"github.com/gonebot-dev/gonebot/message"
)

var pluginList *list.List = list.New()

func LoadPlugin(plugin GonePlugin) {
	log.Printf("Loading Plugin: %s", plugin.Name)
	pluginList.PushBack(plugin)
}

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
