package adaptor

import "github.com/gonebot-dev/gonebot/adaptor/onebot11"

// Start backend. If not set, use "onebot11" backend
func StartBackend(backend string) {
	switch backend {
	case "onebot11":
		onebot11.StartBackend()
	default:
		onebot11.StartBackend()
	}
}
