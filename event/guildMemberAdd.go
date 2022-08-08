package event

import (
	"fmt"

	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

func OnGuildMemberAdd(client *gateway.Session, config *config.Config) func(*discord.GuildMember) {
	return func(member *discord.GuildMember) {
		if config.WelcomeChannelId != "" {
			channel, err := client.State().Channel(config.WelcomeChannelId)

			if err == nil {
				_, _ = client.Channel.Send(channel.Id, fmt.Sprintf("Welcome <@%s> to the server !", member.User.Id))
			}
		}

		if config.MemberRoleId != "" {
			// ToDo : check if the role is in the guild state

			fmt.Printf("Adding role %s to user %s\n", config.MemberRoleId, member.User.Tag())

			client.Guild.AddMemberRole(member.GuildId, member.User.Id, config.MemberRoleId)
		}
	}
}
