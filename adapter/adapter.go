package adapter

type GoneAdapter struct {
	Name              string
	Description       string
	Version           string
	SupportedPlatform string // "qq", "wechat", etc.

	Connector func(incomingChan chan message, resultChan chan message) // Main Threading
}

var adapter GoneAdapter

func SetAdapter(a GoneAdapter) {
	adapter = a
}
