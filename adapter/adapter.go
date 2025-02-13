package adapter

type GoneAdapter struct {
	Name              string
	Description       string
	Version           string
	SupportedPlatform string // "qq", "wechat", etc.

	Init  func() // Run when starting
	Conn  func(incomingChan chan message, resultChan chan message)
	Final func() // Run when closing
}

var adapter GoneAdapter

func SetAdapter(a GoneAdapter) {
	adapter = a
}
