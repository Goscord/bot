package command

import (
	"strings"

	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

type Manager struct {
	client   *gateway.Session
	Commands map[string]Command
}

func Init(client *gateway.Session) *Manager {
	mgr := &Manager{
		client:   client,
		Commands: make(map[string]Command),
	}

	mgr.Register(new(HelpCommand))
	mgr.Register(new(AvatarCommand))
	mgr.Register(new(PingCommand))
	mgr.Register(new(EmbedCommand))
	mgr.Register(new(ServerinfoCommand))

	return mgr
}

func (mgr *Manager) Handler(client *gateway.Session, config *config.Config) func(*discord.Message) {
	return func(message *discord.Message) {
		if !strings.HasPrefix(strings.ToLower(message.Content), config.Prefix) {
			return
		}

		if message.Author.Bot {
			return
		}

		messageArray := strings.Split(message.Content, " ")
		cmdName := messageArray[0][len(config.Prefix):]
		args := messageArray[1:]
		cmd := mgr.Get(cmdName)

		if cmd != nil {
			_ = cmd.Execute(&Context{config: config, client: client, args: args, message: message, cmdMgr: mgr})
		}
	}
}

func (mgr *Manager) Get(name string) Command {
	if cmd, ok := mgr.Commands[name]; ok {
		return cmd
	}

	return nil
}

func (mgr *Manager) Register(cmd Command) {
	appCmd := &discord.ApplicationCommand{
		Name:        cmd.Name(),
		Type:        discord.ApplicationCommandChat,
		Description: cmd.Description(),
		Options:     cmd.Options(),
	}

	mgr.client.Interaction.RegisterCommand(mgr.client.Me().Id, "", appCmd)

	mgr.Commands[cmd.Name()] = cmd
}

// ToDo : Unregister commands
