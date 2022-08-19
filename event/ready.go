package event

import (
	"log"

	"github.com/Goscord/Bot/command"
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

func OnReady(client *gateway.Session, config *config.Config, cmdMgr *command.CommandManager) func() {
	return func() {
		log.Printf("Logged in as %s\n", client.Me().Tag())

		cmdMgr.Init()

		_ = client.SetActivity(&discord.Activity{Name: "/help", Type: discord.ActivityListening})
		_ = client.SetStatus(discord.StatusTypeDoNotDisturb)
	}
}
