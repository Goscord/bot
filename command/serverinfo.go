package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type ServerInfoCommand struct{}

func (c *ServerInfoCommand) Name() string {
	return "serverinfo"
}

func (c *ServerInfoCommand) Description() string {
	return "Display some infos about the server!"
}

func (c *ServerInfoCommand) Category() string {
	return "general"
}

func (c *ServerInfoCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *ServerInfoCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	guild, err := ctx.client.State().Guild(ctx.interaction.GuildId)
	if err != nil {
		e.SetColor(embed.Red)
		e.SetDescription("Could not fetch server informations!")

		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, &discord.InteractionCallbackMessage{Embeds: []*embed.Embed{e.Embed()}, Flags: discord.MessageFlagEphemeral})

		return false
	}

	e.SetColor(embed.Green)
	e.SetTitle(":books: | Server infos")
	e.AddField("Server name", guild.Name, false)
	e.AddField("Server ID", guild.Id, false)
	e.AddField("Members count", fmt.Sprintf("%d", guild.MemberCount), false)

	ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, &discord.InteractionCallbackMessage{Embeds: []*embed.Embed{e.Embed()}, Flags: discord.MessageFlagEphemeral})

	return true
}
