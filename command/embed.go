package command

import (
	"strings"

	"github.com/Goscord/Bot/utils"
	"github.com/Goscord/goscord/discord/embed"
)

type EmbedCommand struct{}

func (c *EmbedCommand) GetName() string {
	return "embed"
}

func (c *EmbedCommand) GetDescription() string {
	return "Send a embed with your message"
}

func (c *EmbedCommand) GetCategory() string {
	return "general"
}

func (c *EmbedCommand) Execute(ctx *Context) bool {
	// Check permission :
	// HACK/TODO : wait the permission on Goscord
	authorIds := []string{
		"233351173665456129", // Bluzzizi
		"810596177857871913", // szerki
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

	if m, err := ctx.client.Channel.Send(ctx.message.ChannelId, e); err == nil {
		ctx.client.Channel.CrosspostMessage(ctx.message.ChannelId, m.Id) // Crosspost the message to the news channels
	}

	return true
}
