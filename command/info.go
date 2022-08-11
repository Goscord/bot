package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type ServerinfoCommand struct{}

func (c *ServerinfoCommand) Name() string {
	return "serverinfo"
}

func (c *ServerinfoCommand) Description() string {
	return "Display some infos about the server!"
}

func (c *ServerinfoCommand) Category() string {
	return "general"
}

func (c *ServerinfoCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *ServerinfoCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	guild, err := ctx.client.State().Guild(ctx.message.GuildId)
	if err != nil {
		e.SetDescription("Could not fetch server informations!")
		ctx.client.Channel.SendMessage(ctx.message.ChannelId, e)

		return false
	}

	e.SetColor(embed.Green)
	e.SetTitle(":books: | Server infos")
	e.AddField("Server name", guild.Name, false)
	e.AddField("Server ID", guild.Id, false)
	e.AddField("Members count", fmt.Sprintf("%d", guild.MemberCount), false)

	ctx.client.Channel.SendMessage(ctx.message.ChannelId, e)

	return true
}
