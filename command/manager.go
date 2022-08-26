package command

import (
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

type CommandManager struct {
	client   *gateway.Session
	config   *config.Config
	commands map[string]Command
}

func NewCommandManager(client *gateway.Session, config *config.Config) *CommandManager {
	return &CommandManager{
		client:   client,
		config:   config,
		commands: make(map[string]Command),
	}
}

func (mgr *CommandManager) Init() {
	mgr.Register(new(HelpCommand))
	mgr.Register(new(AvatarCommand))
	mgr.Register(new(PingCommand))
	mgr.Register(new(EmbedCommand))
	mgr.Register(new(ServerInfoCommand))
}

func (mgr *CommandManager) Handler(client *gateway.Session, config *config.Config) func(*discord.Interaction) {
	return func(interaction *discord.Interaction) {
		if interaction.Member == nil {
			return
		}

		if interaction.Member.User.Bot {
			return
		}

		cmd := mgr.Get(interaction.Data.Name)

		if cmd != nil {
			_ = cmd.Execute(&Context{config: config, client: client, interaction: interaction, cmdMgr: mgr})
		}
	}
}

func (mgr *CommandManager) Get(name string) Command {
	if cmd, ok := mgr.commands[name]; ok {
		return cmd
	}

	return nil
}

func (mgr *CommandManager) Register(cmd Command) {
	appCmd := &discord.ApplicationCommand{
		Name:        cmd.Name(),
		Type:        discord.ApplicationCommandChat,
		Description: cmd.Description(),
		Options:     cmd.Options(),
	}

	mgr.client.Application.RegisterCommand(mgr.client.Me().Id, "", appCmd)

	mgr.commands[cmd.Name()] = cmd
}

// ToDo : Unregister commands
