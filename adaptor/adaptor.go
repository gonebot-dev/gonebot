package adaptor

import (
	"log"

	"github.com/gonebot-dev/gonebot/adaptor/onebot11"
)

// Start backend. If not set, use "onebot11" backend
func StartBackend(backend string) {
	switch backend {
	case "onebot11":
		log.Printf("Using Backend Onebot11\n")
		onebot11.StartBackend()
	default:
		log.Printf("Using Backend Onebot11\n")
		onebot11.StartBackend()
	}
}
