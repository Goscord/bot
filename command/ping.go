package command

import (
	"fmt"

	"github.com/Goscord/goscord/goscord/discord"
)

type PingCommand struct{}

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "Get the bot latency!"
}

func (c *PingCommand) Category() string {
	return "general"
}

func (c *PingCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *PingCommand) Execute(ctx *Context) bool {
	_, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, fmt.Sprintf("Pong! 🏓 (%dms)", ctx.Client.Latency().Milliseconds()))

	return true
}
