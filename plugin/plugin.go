package plugin

import (
	"container/list"
	"log"

	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/rule"
)

// GoneHandler discribes a handler for a plugin.
type GoneHandler struct {
	// What type of message should trigger the handler?
	//
	// The filter results are ORed together.
	Rules []rule.FilterBundle
	// The handler function of the Handler. Will be triggerd by []Command
	//
	// The handlers will be triggered according to the loading order(plugin first, then the handler)
	//
	// Return true if you want to block the propagation, false if you want other plugins to handle it too.
	Handler func(a *adapter.Adapter, msg message.Message) bool
}

// Use this to create your own plugin.
type GonePlugin struct {
	// The name of the plugin.
	Name string
	// The description of the plugin.
	Description string
	// The version of the plugin
	Version string
	// Handlers of the plugin.
	Handlers []GoneHandler
}

var PluginList *list.List = list.New()

// Load plugin to PluginList for usage.
func LoadPlugin(plugin *GonePlugin) {
	log.Printf("[GONEBOT] | Loading Plugin: %s", plugin.Name)
	PluginList.PushBack(plugin)
}

// Get how many plugins have been loaded.
func GetPluginCount() int {
	return PluginList.Len()
}
