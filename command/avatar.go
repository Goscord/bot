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
	return []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionUser,
			Name:        "user",
			Description: "The user to get the avatar from leave empty to get your own avatar",
			Required:    false,
		},
	}
}

func (c *AvatarCommand) Execute(ctx *Context) bool {
	e := builder.NewEmbedBuilder()
	appData := ctx.Interaction.ApplicationCommandData()

	e.SetColor(discord.EmbedGreen)
	e.SetFooter(ctx.Client.Me().Username, ctx.Client.Me().AvatarURL())

	if len(appData.Options) <= 0 {
		e.SetTitle("Your avatar")
		e.SetImage(ctx.Interaction.Member.User.AvatarURL())
	} else {
		user, err := ctx.Client.State().Member(ctx.Interaction.GuildId, appData.Options[0].UserId())

		if err != nil {
			e.SetColor(discord.EmbedRed)
			e.SetDescription("Could not fetch user informations!")

			_, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

			return false
		}

		e.SetTitle(user.User.Username + "'s avatar")
		e.SetImage(user.User.AvatarURL())
	}

	_, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())

	return true
}
