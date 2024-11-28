# gonebot
基于Golang实现的插件化聊天机器人。目前支持onebotv11协议。
## 目录
- [gonebot](#gonebot)
	- [目录](#目录)
	- [为什么选择gonebot？](#为什么选择gonebot)
	- [如何创建一个gonebot](#如何创建一个gonebot)
	- [如何创建一个插件](#如何创建一个插件)
	- [如何创建一个适配器](#如何创建一个适配器)
	- [gonebot工作的逻辑](#gonebot工作的逻辑)
## 为什么选择gonebot？
- 简单易用。你可以用10行不到的代码，轻松[创建](#如何创建一个插件)或[加载](#如何创建一个gonebot)插件。
- 高性能。由于Golang提供的特性，我们只需要占用极小的CPU和内存资源。
- 可移植。静态编译，一次编译，到处运行。~~（其实对于不同平台还是要编译不同的可执行文件）~~
## 如何创建一个gonebot
你可以参考[gonedemo](https://github.com/gonebot-dev/gonedemo)
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
就这么简单。三行完事。
## 如何创建一个插件
只需要实现[`GonePlugin`结构](./plugins/plugin.go)
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
比如我们编写的echo插件（查找`/echo`前缀并指向自己的消息，并返回相同的信息）：
```go
package echo

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/message"
	"github.com/gonebot-dev/gonebot/plugin"
	"github.com/gonebot-dev/gonebot/plugin/rule"
)

var Echo plugin.GonePlugin

func init() {
	Echo.Name = "Echo"
	Echo.Version = "v0.0.1"
	Echo.Description = "Reply the same message of what you have sent"

	Echo.Handlers = append(Echo.Handlers, plugin.GoneHandler{
		Rules: rule.NewRules(rule.ToMe()).And(rule.Command("echo")),
		Handler: func(a *adapter.Adapter, msg message.Message) bool {
			reply := message.NewReply(msg).Join(msg)
			a.SendMessage(reply)
			return true
		},
	})
}
```
对于 `GoneHandler.Rules`，你应该使用 `FilterBundle` 来得到所有其中过滤器的过滤结果`与`在一起，例如上面的echo插件，需要同时满足这两条规则的信息才能被传递到 `Handler` 中。
`Rules` 可以装下一坨 `FilterBundle` 来将他们的结果`或`在一起
对于 `GoneHandler.Handler`，你会得到一个 `adapter.Adapter` 和一个 `message.Message`，你可以解析消息并通过 `Adapter.SendChannel.Push()` 发送若干条消息。
你也许已经发现，适配器中还有一个 `Adapter.ActionChannel`，这是通过调用适配器自定义的动作来实现一些操作，比如说发送戳一戳之类的其它聊天软件不存在的操作，查看 [action.go](./message/action.go) 来了解如何创建一次动作调用，这些操作往往有返回值，返回值会被塞到 `ActionCall.ResultChannel` 里
我们现在只支持文本和文件（图片、音频、视频、文件）这几种类型的消息段，但是别担心！不同的适配器会提供更多的自定义消息！
在 `Handler` 结束以后，你可以选择返回 `false` 来阻止之后的插件继续处理这条消息。插件的顺序取决于你调用 `gonebot.LoadPlugin` 加载插件的顺序
## 如何创建一个适配器
创建一个适配器还挺复杂的，你可能需要通读一下项目代码才能理解，但是你可以参考 [OneBotV11](https://github.com/gonebot-dev/goneadapter-onebotv11)。
## gonebot工作的逻辑
gonebot工作的流程：
1. gonebot 初始化，加载[配置文件](./configuations/)和[内置消息类型](./message/message.go)
2. 每个适配器(adaptor)启动自己的协程并接受消息，塞到 `ReceiveChannel` 里面
3. gonebot 接下来会为每个适配器启动协程，并自动拿出收到的消息，分发给插件来处理
4. 插件(plugin)接下来会处理消息并发送消息到 `SendChannel` 以及发送动作调用到 `ActionChannel` 里面
5. 适配器(adaptor)会从 `SendChannel` 和 `ActionChannel` 里面取出消息和动作调用，并发送消息
