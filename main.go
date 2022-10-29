package main

import (
	"os"

	"github.com/Goscord/Bot/command"
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/Bot/event"
	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/gateway"
	"github.com/joho/godotenv"
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
