package plugin

import "github.com/gonebot-dev/gonebot/message/handler"

type GonePlugin struct {
	Name        string
	Description string
	Version     string

	Handler       []handler.GoneHandler
	ActiveHandler []handler.GoneActiveHandler
}
