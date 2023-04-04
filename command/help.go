package command

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord/builder"

	"github.com/Goscord/goscord/goscord/discord"
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
	e := builder.NewEmbedBuilder()

	e.SetTitle(":books: | Help page")

	for _, cmd := range ctx.CmdMgr.commands {
		e.AddField(fmt.Sprintf("/%s", cmd.Name()), cmd.Description(), false)
	}

	e.SetFooter(ctx.Client.Me().Username, ctx.Client.Me().AvatarURL())
	e.SetColor(discord.EmbedGreen)

	_, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

	return true
}
