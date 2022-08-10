package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord/embed"
)

type HelpCommand struct{}

func (c *HelpCommand) GetName() string {
	return "help"
}

func (c *HelpCommand) GetDescription() string {
	return "Display help page!"
}

func (c *HelpCommand) GetCategory() string {
	return "general"
}

func (c *HelpCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	e.SetTitle(":books: | Help page")

	for _, cmd := range ctx.cmdMgr.Commands {
		e.AddField(fmt.Sprintf("%s%s", ctx.config.Prefix, cmd.GetName()), cmd.GetDescription(), false)
	}

	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(embed.Green)

	_, _ = ctx.client.Channel.SendMessage(ctx.message.ChannelId, e.Embed())

	return true
}
