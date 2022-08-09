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
	CmdMgr *command.Manager
)

func main() {
	Config, _ = config.GetConfig()
	CmdMgr = command.Init()

	Client = goscord.New(&gateway.Options{
		Token:   Config.Token,
		Intents: gateway.IntentGuildMessages + gateway.IntentGuildMembers,
	})

	_ = Client.On("ready", event.OnReady(Client, Config))
	_ = Client.On("messageCreate", CmdMgr.Handler(Client, Config))
	_ = Client.On("guildMemberAdd", event.OnGuildMemberAdd(Client, Config))

	if err := Client.Login(); err != nil {
		panic(err)
	}

	select {}
}