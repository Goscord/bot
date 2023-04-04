package command

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord/builder"
	"math/rand"
	"strings"

	"github.com/Goscord/goscord/goscord/discord"
)

type EmbedCommand struct{}

func (c *EmbedCommand) Name() string {
	return "embed"
}

func (c *EmbedCommand) Description() string {
	return "Send a embed with your message"
}

func (c *EmbedCommand) Category() string {
	return "general"
}

func (c *EmbedCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Name:        "title",
			Type:        discord.ApplicationCommandOptionString,
			Description: "Title of the embed",
			Required:    true,
		},
		{
			Name:        "description",
			Type:        discord.ApplicationCommandOptionString,
			Description: "Description of the embed (use -br for break line)",
			Required:    true,
		},
	}
}

func (c *EmbedCommand) Execute(ctx *Context) bool {
	e := builder.NewEmbedBuilder()

	if !ctx.Interaction.Member.Permissions.Has(discord.BitwisePermissionFlagManageMessages) {
		e.SetDescription("You do not have permission to run this command")
		e.SetColor(discord.EmbedRed)

		_, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, e.Embed())
	} else {
		title := ctx.Interaction.ApplicationCommandData().Options[0].String()
		description := ctx.Interaction.ApplicationCommandData().Options[1].String()

		e.AddField(title, strings.ReplaceAll(description, "-br", "\n"), false)
		e.SetColor(discord.EmbedGreen)

		emojis := []string{"🏳️‍🌈", "🏳️‍⚧️"} // bluzzi: we need to do like discord.js supporting lgbtq+ community
		emoji := emojis[rand.Intn(len(emojis))]

		_, _ = ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, fmt.Sprintf("Message sent! %s", emoji))

		_, _ = ctx.Client.Channel.SendMessage(ctx.Interaction.ChannelId, e.Embed())
	}

	return true
}
