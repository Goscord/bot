package event

import (
	"fmt"
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

func OnReady(client *gateway.Session, config *config.Config) func() {
	return func() {
		fmt.Println("Logged in as " + client.Me().Tag())

		_ = client.SetActivity(&discord.Activity{Name: fmt.Sprintf("%shelp", config.Prefix), Type: 1})
		_ = client.SetStatus("idle")
	}
}
