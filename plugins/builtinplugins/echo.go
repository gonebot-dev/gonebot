package builtinplugins

import (
	"encoding/json"
	"log"

	"github.com/gonebot-dev/gonebot/messages"
	"github.com/gonebot-dev/gonebot/plugins"
)

var Echo plugins.GonePlugin

func handler(msg messages.MessageStruct) messages.ResultStruct {
	var result messages.ResultStruct
	result.Text = msg.Text

	dResult, _ := json.Marshal(result)
	log.Printf("Echo: %s\n", dResult)

	return result
}

func init() {
	Echo.Name = "echo"
	echoHandler := plugins.GoneHandler{Command: "echo", Handler: handler}
	Echo.Handlers = append(Echo.Handlers, echoHandler)
}
