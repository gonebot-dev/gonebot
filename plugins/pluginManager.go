package plugins

import (
	"container/list"
	"log"
	"strings"

	"github.com/gonebot-dev/gonebot/configuations"
	"github.com/gonebot-dev/gonebot/messages"
)

var pluginList *list.List = list.New()

func LoadPlugin(plugin GonePlugin) {
	log.Printf("Loading Plugin: %s", plugin.Name)
	pluginList.PushBack(plugin)
}

func TraversePlugins(msg messages.MessageStruct) (messages.ResultStruct, bool) {
	for pluginElement := pluginList.Front(); pluginElement != nil; pluginElement = pluginElement.Next() {
		plg, _ := pluginElement.Value.(GonePlugin)
		for _, handler := range plg.Handlers {
			if strings.HasPrefix(msg.Text, configuations.GlobalPrefix+handler.Command) {
				//Cut prefix off.
				msg.Text = msg.Text[len(handler.Command)+len(configuations.GlobalPrefix):]
				//Invoke handler
				return handler.Handler(msg), true
			}

		}
	}
	return messages.ResultStruct{}, false
}
