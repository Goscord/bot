package command

import (
	"github.com/Goscord/goscord/discord/embed"
)

type AvatarCommand struct{}

func (c *AvatarCommand) GetName() string {
	return "avatar"
}

func (c *AvatarCommand) GetDescription() string {
	return "Display your profile picture!"
}

func (c *AvatarCommand) GetCategory() string {
	return "general"
}

func (c *AvatarCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	e.SetImage(ctx.message.Author.AvatarURL())
	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(embed.Blurple)

	_, _ = ctx.client.Channel.Send(ctx.message.ChannelId, e.Embed())

	return true
}
