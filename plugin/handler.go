package plugin

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin/rule"
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

func (g *GonePlugin) AddRule(filter []rule.FilterBundle, handler func(*adapter.Adapter, message.Message) bool) {
	g.Handlers = append(g.Handlers, GoneHandler{
		Rules:   filter,
		Handler: handler,
	})
}

func (g *GonePlugin) AddCommandRule(cmd string, handler func(a *adapter.Adapter, msg message.Message) bool) {
	fb := rule.FilterBundle{}
	fb.AddFilter(rule.Command([]string{cmd}))
	g.Handlers = append(g.Handlers, GoneHandler{
		Rules:   []rule.FilterBundle{fb},
		Handler: handler,
	})
}
