package command

import (
	"fmt"

	"github.com/Goscord/Bot/config"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

type CommandManager struct {
	client   *gateway.Session
	Commands map[string]Command
}

func NewCommandManager(client *gateway.Session) *CommandManager {
	return &CommandManager{
		client:   client,
		Commands: make(map[string]Command),
	}
}

func (mgr *CommandManager) Init() {

	mgr.Register(new(HelpCommand))
	mgr.Register(new(AvatarCommand))
	mgr.Register(new(PingCommand))
	mgr.Register(new(EmbedCommand))
	mgr.Register(new(ServerinfoCommand))
}

func (mgr *CommandManager) Handler(client *gateway.Session, config *config.Config) func(*discord.Interaction) {
	return func(interaction *discord.Interaction) {
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
	fmt.Println(name)
	if cmd, ok := mgr.Commands[name]; ok {
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

	mgr.client.Interaction.RegisterCommand(mgr.client.Me().Id, "", appCmd)

	mgr.Commands[cmd.Name()] = cmd
}

// ToDo : Unregister commands
