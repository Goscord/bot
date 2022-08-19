package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord"
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
	ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, &discord.InteractionCallbackMessage{Content: fmt.Sprintf("Pong! 🏓 (%dms)", ctx.client.Latency().Milliseconds()), Flags: discord.MessageFlagEphemeral})

	return true
}
