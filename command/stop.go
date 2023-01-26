package command

import (
	"github.com/Goscord/Bot/player"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/embed"
)

type StopCommand struct{}

func (c *StopCommand) Name() string {
	return "stop"
}

func (c *StopCommand) Description() string {
	return "Stop music!"
}

func (c *StopCommand) Category() string {
	return "music"
}

func (c *StopCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *StopCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	vt, err := ctx.Client.State().VoiceState(ctx.Interaction.GuildId, ctx.Interaction.Member.User.Id)
	if err != nil {
		e.SetColor(embed.Red)
		e.SetDescription("⚠️ | You are not in a voice channel!")
		ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

		return true
	}

	gPlayer, ok := player.PlayerByGuild(ctx.Interaction.GuildId)
	if !ok {
		e.SetColor(embed.Red)
		e.SetDescription("⚠️ | The bot is not playing music!")
		ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

		return true
	}

	if gPlayer.ChannelId() != vt.ChannelId {
		e.SetColor(embed.Red)
		e.SetDescription("⚠️ | You are not in the same voice channel as the bot!")
		ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

		return true
	}

	gPlayer.Stop()

	e.SetColor(embed.Green)
	e.SetDescription("The player has been stopped!")
	ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

	return true
}
