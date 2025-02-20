package adapter

import (
	"log"
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
		log.Fatal("No Adapter Loaded!")
	}
	log.Printf("Starting Adapter %s", adapter.Name)
	adapter.Connector()
}
