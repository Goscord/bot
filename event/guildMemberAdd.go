package event

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
	"log"
	"os"
)

func OnGuildMemberAdd(client *gateway.Session) func(member *discord.GuildMember) {
	return func(member *discord.GuildMember) {
		if os.Getenv("WELCOME_CHANNEL_ID") != "" {
			if channel, bruh := client.State().Channel(os.Getenv("WELCOME_CHANNEL_ID")); bruh == nil {
				client.Channel.SendMessage(channel.Id, fmt.Sprintf("Welcome <@%s> to the server !", member.User.Id))
			} else {
				log.Println("Cannot find channel with id :", os.Getenv("WELCOME_CHANNEL_ID"))
			}
		}

		if os.Getenv("MEMBER_ROLE_ID") != "" {
			// ToDo : Check if the role is in the guild state

			log.Printf("Adding role %s to user %s\n", os.Getenv("MEMBER_ROLE_ID"), member.User.Tag())

			client.Guild.AddMemberRole(member.GuildId, member.User.Id, os.Getenv("MEMBER_ROLE_ID"))
		}
	}
}
