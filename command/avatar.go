package command

import (
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
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

	e.SetImage(ctx.message.Author.AvatarURL())
	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(embed.Green)

	ctx.client.Channel.SendMessage(ctx.message.ChannelId, e.Embed())

	return true
}
