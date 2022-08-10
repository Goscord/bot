package command

import (
	"strings"

	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

type Manager struct {
	Commands map[string]Command
}

func Init() *Manager {
	mgr := &Manager{Commands: make(map[string]Command)}

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
	cmd, _ := mgr.Commands[name]
	return cmd
}

func (mgr *Manager) Register(cmd Command) {
	mgr.Commands[cmd.GetName()] = cmd
}
