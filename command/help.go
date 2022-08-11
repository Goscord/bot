package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type HelpCommand struct{}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "Display help page!"
}

func (c *HelpCommand) Category() string {
	return "general"
}

func (c *HelpCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *HelpCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	e.SetTitle(":books: | Help page")

	for _, cmd := range ctx.cmdMgr.Commands {
		e.AddField(fmt.Sprintf("%s%s", ctx.config.Prefix, cmd.Name()), cmd.Description(), false)
	}

	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(embed.Green)

	ctx.client.Channel.SendMessage(ctx.message.ChannelId, e.Embed())

	return true
}
