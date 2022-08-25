package main

import (
	"github.com/Goscord/Bot/command"
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/Bot/event"
	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/gateway"
)

var (
	client *gateway.Session
	Config *config.Config
	cmdMgr *command.CommandManager
)

func main() {
	Config, _ = config.GetConfig()
	client = goscord.New(&gateway.Options{
		Token:   Config.Token,
		Intents: gateway.IntentGuilds | gateway.IntentGuildMessages | gateway.IntentGuildMembers,
	})
	cmdMgr = command.NewCommandManager(client, Config)

	_ = client.On("ready", event.OnReady(client, Config, cmdMgr))
	_ = client.On("interactionCreate", cmdMgr.Handler(client, Config))
	_ = client.On("guildMemberAdd", event.OnGuildMemberAdd(client, Config))

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}
