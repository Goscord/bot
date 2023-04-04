package command

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/builder"
)

type AvatarCommand struct{}

func (c *AvatarCommand) Name() string {
	return "avatar"
}

func (c *AvatarCommand) Description() string {
	return "Display your profile picture!"
}

func (c *AvatarCommand) Category() string {
	return "general"
}

func (c *AvatarCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *AvatarCommand) Execute(ctx *Context) bool {
	e := builder.NewEmbedBuilder()

	e.SetTitle("Your avatar")
	e.SetImage(ctx.Interaction.Member.User.AvatarURL())
	e.SetFooter(ctx.Client.Me().Username, ctx.Client.Me().AvatarURL())
	e.SetColor(discord.EmbedGreen)

	_, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

	return true
}
