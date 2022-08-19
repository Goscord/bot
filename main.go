package main

import (
	"github.com/Goscord/Bot/command"
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/Bot/event"
	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/gateway"
)

var (
	Client *gateway.Session
	Config *config.Config
	CmdMgr *command.CommandManager
)

func main() {
	Config, _ = config.GetConfig()
	Client = goscord.New(&gateway.Options{
		Token:   Config.Token,
		Intents: gateway.IntentGuilds + gateway.IntentGuildMessages + gateway.IntentGuildMembers,
	})
	CmdMgr = command.NewCommandManager(Client)

	_ = Client.On("ready", event.OnReady(Client, Config, CmdMgr))
	_ = Client.On("interactionCreate", CmdMgr.Handler(Client, Config))
	_ = Client.On("guildMemberAdd", event.OnGuildMemberAdd(Client, Config))

	if err := Client.Login(); err != nil {
		panic(err)
	}

	select {}
}
