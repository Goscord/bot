package event

import (
	"fmt"
	"log"

	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

func OnReady(client *gateway.Session, config *config.Config) func() {
	return func() {
		log.Printf("Logged in as %s\n", client.Me().Tag())

		_ = client.SetActivity(&discord.Activity{Name: fmt.Sprintf("%shelp", config.Prefix), Type: discord.ActivityListening})
		_ = client.SetStatus(discord.StatusTypeDoNotDisturb)
	}
}
