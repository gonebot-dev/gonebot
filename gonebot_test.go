package gonebot

import (
	"testing"

	"github.com/gonebot-dev/gonebot/plugins/builtinplugins"
)

func TestMain(m *testing.M) {
	LoadPlugin(builtinplugins.Echo)
	StartBackend("onebot11")
}
