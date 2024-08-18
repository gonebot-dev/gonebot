package backend

import "gonebot/backend/onebot11"

// TODO support multi backend.
var backend string = "onebot11"

// Currently supported backend: "onebot11"
func SetBackend(bknd string) {
	backend = bknd
}

// Start backend. If not set, use "onebot11" backend
func Initialization() {
	switch backend {
	case "onebot11":
		onebot11.Initialization()
	}

}
