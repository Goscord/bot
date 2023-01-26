package command

import (
	"strings"

	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/embed"
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
			Description: "Description of the embed (use -br for break line)",
			Required:    true,
		},
	}
}

func (c *EmbedCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	if !ctx.Interaction.Member.Permissions.Has(discord.BitwisePermissionFlagManageMessages) {
		e.SetDescription("You do not have permission to run this command")
		e.SetColor(embed.Red)

		ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, &discord.Message{Embeds: []*embed.Embed{e.Embed()}})
	} else {
		title := ctx.Interaction.ApplicationCommandData().Options[0].String()
		description := ctx.Interaction.ApplicationCommandData().Options[1].String()

		e.AddField(title, strings.ReplaceAll(description, "-br", "\n"), false)
		e.SetColor(embed.Green)

		ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, &discord.Message{Embeds: []*embed.Embed{e.Embed()}})
	}

	return true
}
