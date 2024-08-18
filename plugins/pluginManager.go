package plugins

import (
	"container/list"
	"gonebot/configuations"
	"gonebot/messages"
	"log"
	"strings"
)

var pluginList *list.List = list.New()

func LoadPlugin(plugin GonePlugin) {
	log.Printf("Plugin: %s", plugin.Name)
	pluginList.PushBack(plugin)
}

func TraversePlugins(msg messages.MessageStruct) (messages.MessageStruct, bool) {
	for pluginElement := pluginList.Front(); pluginElement != nil; pluginElement = pluginElement.Next() {
		plg, _ := pluginElement.Value.(GonePlugin)
		if strings.HasPrefix(msg.Text, plg.Command) {
			//Cut prefix off.
			msg.Text = msg.Text[len(plg.Command)+len(configuations.GlobalPrefix):]
			//Invoke handler
			return plg.Handler(msg), true
		}
	}
	return messages.MessageStruct{}, false
}
