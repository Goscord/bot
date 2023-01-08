package main

import (
	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/gateway"
	"github.com/joho/godotenv"
	"os"

	"github.com/Goscord/Bot/command"
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/Bot/event"
)

var (
	client *gateway.Session
	Config *config.Config
	cmdMgr *command.CommandManager
)

func main() {
	// Load envionment variables :
	godotenv.Load()

	Config, _ = config.GetConfig()

	// Create client instance :
	client = goscord.New(&gateway.Options{
		Token:   os.Getenv("BOT_TOKEN"),
		Intents: gateway.IntentGuilds | gateway.IntentGuildMessages | gateway.IntentGuildMembers,
	})

	// Load command manager :
	cmdMgr = command.NewCommandManager(client, Config)

	// Load events :
	_ = client.On("ready", event.OnReady(client, Config, cmdMgr))
	_ = client.On("interactionCreate", cmdMgr.Handler(client, Config))
	_ = client.On("guildMemberAdd", event.OnGuildMemberAdd(client, Config))

	// Login client :
	if err := client.Login(); err != nil {
		panic(err)
	}

	// Keep bot running :
	select {}
}
