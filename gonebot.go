package gonebot

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/configurations"
	"github.com/gonebot-dev/gonebot/logging"
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
	"github.com/rs/zerolog"
)

func parseMessage(a *adapter.Adapter, msg message.Message) {
	shouldBlock := false
	for pluginElement := plugin.PluginList.Front(); pluginElement != nil; pluginElement = pluginElement.Next() {
		plg, _ := pluginElement.Value.(*plugin.GonePlugin)
		for _, handler := range plg.Handlers {
			if !handler.Rules.Filter(msg) {
				continue
			}
			shouldBlock := handler.Handler(a, msg)
			if shouldBlock {
				break
			}
		}
		if shouldBlock {
			break
		}
	}
}

func messageListener(a *adapter.Adapter) {
	for {
		msg := a.ReceiveChannel.Pull()
		go parseMessage(a, msg)
	}
}

const banner = `
=========================================
   ______                 __          __
  / ____/___  ____  ___  / /_  ____  / /_
 / / __/ __ \/ __ \/ _ \/ __ \/ __ \/ __/
/ /_/ / /_/ / / / /  __/ /_/ / /_/ / /_
\____/\____/_/ /_/\___/_.___/\____/\__/
=========================================

`

// LoadPlugin helps you with loading any plugin for gonebot
//
// Just a protocol for plugin.LoadPlugin
func LoadPlugin(plug *plugin.GonePlugin) {
	plugin.LoadPlugin(plug)
}

// Load Adapter helps you with loading any adapter for gonebot
func LoadAdapter(a *adapter.Adapter) {
	a.ReceiveChannel = *message.NewMessageChannel()
	a.SendChannel = *message.NewMessageChannel()
	a.ActionChannel = *message.NewActionChannel()
	adapter.AdapterList.PushBack(a)
}

// Run gonebot, this will start all adapters and wait for them to end.
func Run() {
	var waitGroup sync.WaitGroup
	for adapterInstance := adapter.AdapterList.Front(); adapterInstance != nil; adapterInstance = adapterInstance.Next() {
		a, _ := adapterInstance.Value.(*adapter.Adapter)
		waitGroup.Add(2)
		go a.Start()
		go messageListener(a)
	}
	waitGroup.Wait()
	for adapterInstance := adapter.AdapterList.Front(); adapterInstance != nil; adapterInstance = adapterInstance.Next() {
		a, _ := adapterInstance.Value.(*adapter.Adapter)
		a.Finalize()
	}
}

func init() {
	fmt.Print(banner)
	logging.Init()
	configurations.Init()
	level, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		level = "INFO"
	}
	level = strings.ToUpper(level)
	switch level {
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logging.Log(zerolog.InfoLevel, "GoneBot", "GoneBot initilization complete!")
}
