package command

import (
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

type Context struct {
	config  *config.Config
	cmdMgr  *Manager
	client  *gateway.Session
	message *discord.Message
	args    []string
}

type Command interface {
	Name() string
	Description() string
	Category() string
	Options() []*discord.ApplicationCommandOption
	Execute(ctx *Context) bool
}
