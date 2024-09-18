# gonebot
A Golang chatbot, currently support onebotv11.
[中文文档](./README_CN.md)
## Catalog
- [gonebot](#gonebot)
	- [Catalog](#catalog)
	- [Why gonebot?](#why-gonebot)
	- [How to create a bot](#how-to-create-a-bot)
	- [How to create a plugin](#how-to-create-a-plugin)
	- [How to create an adapter](#how-to-create-an-adapter)
	- [The logic of gonebot](#the-logic-of-gonebot)
	- [TODO](#todo)
## Why gonebot?
- Easy to use. You can easily [load plugins](#how-to-create-a-bot) or [create them](#how-to-create-a-plugin) with less than 10 lines of code!
- High Performance. Powered by golang, we use minimum CPU and memory resource.
- Portable. Static compilation means you can compile the bot into an excutable file.
## How to create a bot
You can refer to [gonedemo](https://github.com/gonebot-dev/gonedemo)
```
package main
import (
	"github.com/gonebot-dev/gonebot"
	"github.com/gonebot-dev/goneplugin-echo"
	"github.com/gonebot-dev/goneadapter-onebotv11"
)
func main() {
	gonebot.LoadPlugin(&echo.Echo)
	gonebot.StartAdapter(&onebotv11.OneBotV11)
	gonebot.Run()
}
```
Done. Three lines, one bot.
## How to create a plugin
To create a plugin, you need to implement a [`GonePlugin` struct](./plugins/plugin.go).
```go
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
```
For example, the echo plugin, which detects message with `/echo` as prefix and with Message.IsToMe as true, and replies with the same message
By the way, we are using our builtin rules to filter messages, [`Details`](./rule/builtins.go):
```go
package echo

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
	"github.com/gonebot-dev/gonebot/rule"
)

var Echo plugin.GonePlugin

func init() {
	Echo.Name = "Echo"
	Echo.Version = "v0.0.1"
	Echo.Description = "Reply the same message of what you have sent"

	Echo.Handlers = append(Echo.Handlers, plugin.GoneHandler{
		Rules: []rule.FilterBundle{{Filters: []rule.FilterRule{rule.ToMe(), rule.Command([]string{"echo"})}}},
		Handler: func(a *adapter.Adapter, msg message.Message) bool {
			reply := msg
			reply.Sender = reply.Self
			reply.Receiver = msg.Sender
			a.SendChannel.Push(reply, false)
			return true
		},
	})
}
```
For the `GoneHandler.Rules` part, you shall use `FilterBundle` to get the results of the filters `AND`ed together. For example, a message must both satisfies `rule.ToMe()` and `rule.Command([]string{"echo"})` to be parsed in this handler above.
And the `Rules` can contain a bunch of `FilterBundle`s, which results are `OR`ed together.
For the `GoneHandler.Handler` part, you will receive an [Adapter](./adapter/adapter.go) and a [Message](./message/message.go). You can parse the message and send one or multiple reply by calling the `Adapter.SendChannel.Push()` method.
You may notice that there is also an `Adapter.ActionChannel`, that is the channel to call the custom actions provided by the adapters, refer to [action.go](./message/action.go) for how to create an action call. After calling the actions, the result will be pushed into the `ActionCall.ResultChannel`.
We only support `text`, `image`, `voice`, `video` and `file` type of message currently, but don't be worry, the `Adapter`s may provide some custom message types!
After the `Handler` is done, it should return a boolean value to indicate whether the deliver process should continue. That means, if the handler returns `false`, the message will not be delivered to other handlers.
And the order of the handlers depends on the order which you call `gonebot.LoadPlugin` to load the plugins
## How to create an adapter
Creating an adapter is more than complicated, before that, you may need to read all the codes of gonebot to get a better understanding of how it works. And you can refer to [OneBotV11](https://github.com/gonebot-dev/goneadapter-onebotv11) adapter to get an idea of how the adapters are written.
## The logic of gonebot
How does gonebot work?
1. gonebot initialize: loading [configuations](./configuations/) and [builtin message types](./message/message.go).
2. First, the `Adaptor`s will start their own goroutine in the background and push messages into their `ReceiveChannel`s
3. Then gonebot will use a goroutine for each of the `Adapter`s to automatically pull messages from `ReceiveChannel` and deliver them to plugins for processing.
4. `Plugins` will then process those messages and push some results into the `SendChannel`s and `ActionChannel`s
5. The `Adaptor`s will fetch results from their own `SendChannel` and `ActionChannel` and send the results

## TODO
- [x] Docs about how to create a plugin.
- [x] Docs about how to create an adapter.
- [x] Dotenv configuation file support.
- [x] Images message support.
- [x] README_CN.
- [x] Multiple adaptor support.
- [x] Advanced context support.
- [ ] Plugin & Adapter repository.
- [ ] Project creater
- [ ] Make it perform better
