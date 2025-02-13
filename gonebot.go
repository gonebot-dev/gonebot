package gonebot

import "github.com/gonebot-dev/gonebot/adapter"

func LoadAdapter(a GoneAdapter) {
	adapter.SetAdapter(a)
}
