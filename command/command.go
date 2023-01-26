package command

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
)

type Context struct {
	CmdMgr      *CommandManager
	Client      *gateway.Session
	Interaction *discord.Interaction
}

type Command interface {
	Name() string
	Description() string
	Category() string
	Options() []*discord.ApplicationCommandOption
	Execute(ctx *Context) bool
}
