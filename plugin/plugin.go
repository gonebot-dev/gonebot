package plugin

import "github.com/gonebot-dev/gonebot/message/handler"

type GonePlugin struct {
	Name        string
	Description string
	Version     string

	Handlers       []handler.GoneHandler
	ActiveHandlers []handler.GoneActiveHandler
}
