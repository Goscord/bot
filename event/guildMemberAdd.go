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
			if channel, err := client.State().Channel(config.WelcomeChannelId); err == nil {
				_, _ = client.Channel.Send(channel.Id, fmt.Sprintf("Welcome <@%s> to the server !", member.User.Id))
			} else {
				fmt.Println("Cannot find welcome channel")
			}
		}

		if config.MemberRoleId != "" {
			// ToDo : check if the role is in the guild state

			fmt.Printf("Adding role %s to user %s\n", config.MemberRoleId, member.User.Tag())

			client.Guild.AddMemberRole(member.GuildId, member.User.Id, config.MemberRoleId)
		}
	}
}
