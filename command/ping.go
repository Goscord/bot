package command

import (
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
	/*
		m, err := ctx.client.Channel.SendMessage(ctx.interaction.ChannelId, "Pinging...")

		if err != nil {
			return true
		}

		latency := m.Timestamp.Sub(ctx.interaction.Message.Timestamp)
		e := embed.NewEmbedBuilder()

		e.SetTitle("Pong!")
		e.SetDescription(fmt.Sprintf("Bot : %d ms\nWebsocket : %d ms", latency.Milliseconds(), ctx.client.Latency().Milliseconds()))
		e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
		e.SetColor(embed.Green)

		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, e.Embed())
	*/

	ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, &discord.InteractionCallbackMessage{Content: "Pong!", Flags: discord.MessageFlagEphemeral})

	return true
}
