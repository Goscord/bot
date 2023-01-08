package command

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
)

type CommandManager struct {
	client   *gateway.Session
	commands map[string]Command
}

func NewCommandManager(client *gateway.Session) *CommandManager {
	return &CommandManager{
		client:   client,
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

func (mgr *CommandManager) Handler(client *gateway.Session) func(*discord.Interaction) {
	return func(interaction *discord.Interaction) {
		if interaction.Type != discord.InteractionTypeApplicationCommand {
			return
		}

		if interaction.Member == nil {
			return
		}

		if interaction.Member.User.Bot {
			return
		}

		cmd := mgr.Get(interaction.ApplicationCommandData().Name)

		if cmd != nil {
			client.Interaction.CreateResponse(interaction.Id, interaction.Token, nil) // defer interaction

			_ = cmd.Execute(&Context{client: client, interaction: interaction, cmdMgr: mgr})
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
