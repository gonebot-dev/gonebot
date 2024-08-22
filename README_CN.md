# gonebot
基于Golang实现的插件化聊天机器人。目前支持onebotv11协议。
## 目录
- [gonebot](#gonebot)
  - [目录](#目录)
  - [为什么选择gonebot？](#为什么选择gonebot)
  - [如何创建一个gonebot](#如何创建一个gonebot)
  - [如何创建一个插件](#如何创建一个插件)
  - [gonebot工作的逻辑](#gonebot工作的逻辑)
## 为什么选择gonebot？
- 简单易用。你可以用10行不到的代码，轻松[创建](#如何创建一个插件)或[加载](#如何创建一个gonebot)插件。
- 高性能。由于Golang提供的特性，我们只需要占用极小的CPU和内存资源。
- 可移植。静态编译，一次编译，到处运行。
## 如何创建一个gonebot
你可以参考[gonedemo](https://github.com/gonebot-dev/gonedemo)
```
package main
import (
	"github.com/gonebot-dev/gonebot"
	"github.com/gonebot-dev/gonebot/plugins/builtinplugins"
)
func main() {
	gonebot.LoadPlugin(builtinplugins.Echo)
	gonebot.StartBackend("onebot11")
}
```
就这么简单。两行完事。
## 如何创建一个插件
只需要实现[`GonePlugin`结构](./plugins/pluginStruct.go)
```go
type GoneHandler struct {
	Command []string
	Handler func(msg messages.MessageStruct) messages.ResultStruct
}
type GonePlugin struct {
	Name        string
	Description string
	Handlers    []GoneHandler
}
```
比如我们内置的echo插件：
```go
func handler(msg messages.MessageStruct) messages.ResultStruct {
	var result messages.ResultStruct
	result.Text = msg.Text
	return result
}
var Echo plugins.GonePlugin
Echo.Name = "echo"
echoHandler := plugins.GoneHandler{}
echoHandler.Command = []string{"echo"}
echoHandler.Handler = handler
Echo.Handlers = append(Echo.Handlers, echoHandler)
```
你的插件得到了一个`MessageStruct`变量，你可以读取它，然后创建并返回一个`ResultStruct`。好好利用你的IDE的提示插件！
## gonebot工作的逻辑
gonebot工作的流程：
1. gonebot初始化，加载[配置文件](./configuations/)和[插件](./plugins/pluginManager.go)
2. 适配器(adaptor)连接前端，它的接收器`receiver`并解码json为一个[`MessageStruct`](./messages/messageStruct.go)。该`MessageStruct`进入[`MessageQueue`](./messages/messageQueue.go)等待处理。这是第一个协程。
3. 处理器(processor)从`MessageQueue`取出消息并匹配命令、调用插件的handler函数进行处理，将结果[`resultStruct`](./messages/resultStruct.go)送入[`resultQueue`](./messages/resultQueue.go)中等待发送。这是第二个协程。
4. 适配器(adaptor)从[`resultQueue`]中取出结果并发送给前端。这是第三个协程。
2-4步将会反复执行，从而不断处理消息。
```mermaid
flowchart LR
    Thread1[Receiver]
    Queue1(MessageQueue)
    Thread2[Processor]
    Queue2(ResultQueue)
    Thread3[Sender]
    Thread1-->Queue1
    Queue1-->Thread2
    Thread2-->Queue2
    Queue2-->Thread3
    Thread3-->Queue2
    Queue2-->Thread2
    Thread2-->Queue1
    Queue1-->Thread1
```