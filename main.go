package main

import (
	"os"

	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/gateway"
	"github.com/joho/godotenv"

	"github.com/Goscord/Bot/command"
	"github.com/Goscord/Bot/event"
)

var (
	client *gateway.Session
	cmdMgr *command.CommandManager
)

func main() {
	// Load envionment variables :
	godotenv.Load()

	// Create client instance :
	client = goscord.New(&gateway.Options{
		Token:   os.Getenv("BOT_TOKEN"),
		Intents: gateway.IntentGuilds | gateway.IntentGuildMessages | gateway.IntentGuildMembers | gateway.IntentGuildVoiceStates,
	})

	// Load command manager :
	cmdMgr = command.NewCommandManager(client)

	// Load events :
	_ = client.On("ready", event.OnReady(client, cmdMgr))
	_ = client.On("interactionCreate", cmdMgr.Handler(client))
	_ = client.On("guildMemberAdd", event.OnGuildMemberAdd(client))
	// TODO: Check player status with voiceStateUpdate event

	// Login client :
	if err := client.Login(); err != nil {
		panic(err)
	}

	// Keep bot running :
	select {}
}
