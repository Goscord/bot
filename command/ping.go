package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord/embed"
)

type PingCommand struct{}

func (c *PingCommand) GetName() string {
	return "ping"
}

func (c *PingCommand) GetDescription() string {
	return "Get the bot latency!"
}

func (c *PingCommand) GetCategory() string {
	return "general"
}

func (c *PingCommand) Execute(ctx *Context) bool {
	m, err := ctx.client.Channel.Send(ctx.message.ChannelId, "Pinging...")

	if err != nil {
		return true
	}

	latency := m.Timestamp.Sub(ctx.message.Timestamp)
	e := embed.NewEmbedBuilder()

	e.SetTitle("Pong!")
	e.SetDescription(fmt.Sprintf("Bot : %d ms\nWebsocket : %d ms", latency.Milliseconds(), ctx.client.Latency().Milliseconds()))
	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(embed.Green)

	_, _ = ctx.client.Channel.Edit(ctx.message.ChannelId, m.Id, e.Embed())

	return true
}
