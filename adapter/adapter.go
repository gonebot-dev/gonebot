package adapter

import (
	"container/list"

	"github.com/gonebot-dev/gonebot/message"
)

// Implement this to create an adapter
type Adapter struct {
	// Name of the adapter
	Name string
	// Version of the adapter
	Version string
	// Description of the adapter
	Description string
	// Start the adapter, will be run as a goroutine
	Start func()
	// Finalize the adapter, will be run after everything
	Finalize func()

	// Will automatically initialize when LoadAdapter is called
	ReceiveChannel message.MessageChannel
	SendChannel    message.MessageChannel
	ActionChannel  message.ActionChannel
}

// CallAction will call the adapter action and wait for the result, the result should be nil if error or ignored
//
// # Your adapter must push result to the ResultChannel!
//
// # Otherwise the plugin handler would stuck!
//
// # I warn you!
func (a *Adapter) CallAction(action *message.ActionCall) (result any) {
	action.ResultChannel = make(chan any, 1)
	a.ActionChannel.Push(action)
	return <-action.ResultChannel
}

var AdapterList *list.List = list.New()
