package command

import (
	"github.com/Goscord/Bot/utils"
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
	// Check permission :
	// HACK/TODO : wait the permission on Goscord
	authorIds := []string{
		"233351173665456129", // Bluzzi
		"810596177857871913", // szeroki
	}

	e := embed.NewEmbedBuilder()

	if !utils.ArrayContains(authorIds, ctx.interaction.Member.User.Id) {
		e.SetDescription("You do not have permission to run this command")
		e.SetColor(embed.Red)
	} else {
		title := ctx.interaction.Data.Options[0].String()
		description := ctx.interaction.Data.Options[1].String()

		e.SetTitle(title)
		e.SetDescription(description)
		e.SetColor(embed.Green)

		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, e.Embed())
	}

	return true
}
