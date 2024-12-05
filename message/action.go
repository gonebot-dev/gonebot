package message

// If you want to call adapter action, you should use this struct
type ActionCall struct {
	// Which action
	Action any
	// Which channel to push the result, will initialize automatically
	ResultChannel *chan any
}

const ACTION_CHAN_CAPACITY = 64

type ActionChannel struct {
	channel chan *ActionCall
}

// Create a new ActionChannel
func NewActionChannel() *ActionChannel {
	return &ActionChannel{
		channel: make(chan *ActionCall, ACTION_CHAN_CAPACITY),
	}
}

// Push a message to the channel
//
// # If the channel is full, the oldest message will be dropped
func (ac *ActionChannel) Push(action *ActionCall) {
	// channel full, drop
	if cap((*ac).channel) == len((*ac).channel) {
		<-(*ac).channel
	}
	// push
	(*ac).channel <- action
}

// Pull a message from the channel
func (ac *ActionChannel) Pull() *ActionCall {
	return <-(*ac).channel
}
