package command

import (
	"fmt"
	"github.com/Goscord/Bot/player"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/embed"
	"github.com/kkdai/youtube/v2"
)

type PlayCommand struct{}

func (c *PlayCommand) Name() string {
	return "play"
}

func (c *PlayCommand) Description() string {
	return "Play music!"
}

func (c *PlayCommand) Category() string {
	return "music"
}

func (c *PlayCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionString,
			Name:        "url",
			Description: "YouTube video url",
			Required:    true,
		},
	}
}

func (c *PlayCommand) Execute(ctx *Context) bool {
	var m *discord.Message
	e := embed.NewEmbedBuilder()

	vt, err := ctx.Client.State().VoiceState(ctx.Interaction.GuildId, ctx.Interaction.Member.User.Id)
	if err != nil {
		e.SetColor(embed.Red)
		e.SetDescription("⚠️ | You are not in a voice channel!")

		m, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

		return true
	}

	gPlayer, ok := player.PlayerByGuild(ctx.Interaction.GuildId)
	if !ok {
		gPlayer = player.NewPlayer(ctx.Client, ctx.Interaction.GuildId, vt.ChannelId, ctx.Interaction.ChannelId)
	}

	if gPlayer.ChannelId() != vt.ChannelId {
		e.SetColor(embed.Red)
		e.SetDescription("⚠️ | You are not in the same voice channel as the bot!")
		m, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())
		return true
	}

	e.SetColor(embed.Yellow)
	e.SetDescription("⏳ | Searching for your query...")
	m, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

	ytb := youtube.Client{}

	video, err := ytb.GetVideo(ctx.Interaction.ApplicationCommandData().Options[0].String())
	if err != nil {
		e.SetColor(embed.Red)
		e.SetDescription("❌ | Video not found!")

		ctx.Client.Interaction.EditFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, m.Id, e.Embed())

		return true
	}

	formats := video.Formats.WithAudioChannels()
	streamUrl, err := ytb.GetStreamURL(video, &formats[0])

	track := &player.Track{
		Title:     video.Title,
		Author:    video.Author,
		StreamUrl: streamUrl,
		Requester: ctx.Interaction.Member.User,
	}

	gPlayer.AddTrack(track)

	e.SetColor(embed.Green)
	e.SetDescription(fmt.Sprintf("Added **%s** by %s to the queue!", video.Title, video.Author))
	e.SetThumbnail(video.Thumbnails[0].URL)
	e.SetFooter(fmt.Sprintf("Requested by %s", ctx.Interaction.Member.User.Username), ctx.Interaction.Member.User.AvatarURL())

	ctx.Client.Interaction.EditFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, m.Id, e.Embed())

	if !gPlayer.IsPlaying() {
		go gPlayer.Play()
	}

	return true
}
