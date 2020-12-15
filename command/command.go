package command

import (
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

type Context struct {
	config *config.Config
	cmdMgr *Manager
	client *gateway.Session
	message *discord.Message
	args []string
}

type Command interface {
	GetName() string
	GetDescription() string
	GetCategory() string
	Execute(ctx *Context) bool
}