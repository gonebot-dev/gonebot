package plugin

import (
	"container/list"
	"log"
)

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
