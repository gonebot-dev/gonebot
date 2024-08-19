package adaptor

import "gonebot/adaptor/onebot11"

// Start backend. If not set, use "onebot11" backend
func UseBackend(backend string) {
	switch backend {
	case "onebot11":
		onebot11.Initialization()
	default:
		onebot11.Initialization()
	}
}
