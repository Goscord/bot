package command

import (
	"log"
	"strings"

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

	if !utils.ArrayContains(authorIds, ctx.message.Author.Id) {
		e.SetDescription("You do not have permission to run this command")
		e.SetColor(embed.Red)
	} else {
		if len(ctx.args) > 1 {
			if ctx.args[0] != "nil" {
				e.SetTitle(ctx.args[0])
			}

			e.SetDescription(strings.Join(ctx.args[1:], " "))
			e.SetColor(embed.Green)
		} else {
			e.SetDescription("You must do /embed <title | nil> <content>")
			e.SetColor(embed.Red)
		}
	}

	if m, err := ctx.client.Channel.SendMessage(ctx.message.ChannelId, e); err == nil {
		channel, err := ctx.client.State().Channel(ctx.message.ChannelId)

		if err != nil {
			log.Println("Cannot find channel")
		}

		if channel.Type == discord.ChannelTypeNews {
			ctx.client.Channel.CrosspostMessage(ctx.message.ChannelId, m.Id) // Crosspost the message to the news channels
		}
	}

	return true
}
