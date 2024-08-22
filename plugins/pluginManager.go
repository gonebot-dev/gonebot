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
			for _, prefix := range handler.Command {
				if strings.HasPrefix(msg.Text, configuations.GlobalPrefix+prefix) {
					//Cut prefix off.
					msg.Text = msg.Text[len(prefix)+len(configuations.GlobalPrefix):]
					log.Printf("Plugin %s Handling:\n", plg.Name)
					//Invoke handler
					return handler.Handler(msg), true
				}
			}

		}
	}
	return messages.ResultStruct{}, false
}
