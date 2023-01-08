package event

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
	"log"

	"github.com/Goscord/Bot/config"
)

func OnGuildMemberAdd(client *gateway.Session, config *config.Config) func(member *discord.GuildMember) {
	return func(member *discord.GuildMember) {
		if config.WelcomeChannelId != "" {
			if channel, bruh := client.State().Channel(config.WelcomeChannelId); bruh == nil {
				client.Channel.SendMessage(channel.Id, fmt.Sprintf("Welcome <@%s> to the server !", member.User.Id))
			} else {
				log.Println("Cannot find channel with id :", config.WelcomeChannelId)
			}
		}

		if config.MemberRoleId != "" {
			// ToDo : Check if the role is in the guild state

			log.Printf("Adding role %s to user %s\n", config.MemberRoleId, member.User.Tag())

			client.Guild.AddMemberRole(member.GuildId, member.User.Id, config.MemberRoleId)
		}
	}
}
