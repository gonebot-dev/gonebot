package adapter

import (
	"fmt"
	"log/slog"
)

type GoneAdapter struct {
	Name              string
	Description       string
	Version           string
	SupportedPlatform string // "qq", "wechat", etc.

	Connector func() // Main Thread
}

var adapter GoneAdapter
var adapterLoaded bool = false

func SetAdapter(a GoneAdapter) {
	adapter = a
	adapterLoaded = true
}

func StartAdapter() {
	if !adapterLoaded {
		slog.Error("No Adapter Loaded!")
	}
	slog.Info(fmt.Sprintf("Starting Adapter %s", adapter.Name))
	adapter.Connector()
}
