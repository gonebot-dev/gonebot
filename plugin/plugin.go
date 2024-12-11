package plugin

import (
	"container/list"

	"github.com/gonebot-dev/gonebot/logging"
	"github.com/rs/zerolog"
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
	logging.Logf(zerolog.InfoLevel, "GoneBot", "Loading Plugin: %s", plugin.Name)
	PluginList.PushBack(plugin)
}

// Get how many plugins have been loaded.
func GetPluginCount() int {
	return PluginList.Len()
}
