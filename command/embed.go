package command

import (
	"strings"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type EmbedCommand struct{}

func (c *EmbedCommand) Name() string {
	return "embed"
}

func (c *EmbedCommand) Description() string {
	return "Send a embed with your message"
}

func (c *EmbedCommand) Category() string {
	return "general"
}

func (c *EmbedCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Name:        "title",
			Type:        discord.ApplicationCommandOptionString,
			Description: "Title of the embed",
			Required:    true,
		},
		{
			Name:        "description",
			Type:        discord.ApplicationCommandOptionString,
			Description: "Description of the embed",
			Required:    true,
		},
	}
}

func (c *EmbedCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	if !ctx.interaction.Member.Permissions.Has(discord.BitwisePermissionFlagManageMessages) {
		e.SetDescription("You do not have permission to run this command")
		e.SetColor(embed.Red)

		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, &discord.InteractionCallbackMessage{Embeds: []*embed.Embed{e.Embed()}, Flags: discord.MessageFlagEphemeral})
	} else {
		title := ctx.interaction.Data.Options[0].String()
		description := ctx.interaction.Data.Options[1].String()

		e.AddField(title, strings.ReplaceAll(description, "{line}", "\n"), false)
		e.SetColor(embed.Green)

		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, e.Embed())
	}

	return true
}
