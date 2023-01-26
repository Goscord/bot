package command

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/embed"
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
	e := embed.NewEmbedBuilder()

	e.SetImage(ctx.Interaction.Member.User.AvatarURL())
	e.SetFooter(ctx.Client.Me().Username, ctx.Client.Me().AvatarURL())
	e.SetColor(embed.Green)

	ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, &discord.Message{Embeds: []*embed.Embed{e.Embed()}})

	return true
}
