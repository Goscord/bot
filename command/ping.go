package command

import (
	"fmt"
	"log"

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
	_, err := ctx.client.Interaction.CreateFollowupMessage(ctx.client.Me().Id, ctx.interaction.Token, &discord.InteractionCallbackMessage{Content: fmt.Sprintf("Pong! 🏓 (%dms)", ctx.client.Latency().Milliseconds()), Flags: discord.MessageFlagEphemeral})
	log.Println(err)
	return true
}
