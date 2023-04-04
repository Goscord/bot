package command

import (
	"fmt"
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
	mgr.Register(new(PlayCommand))
	mgr.Register(new(StopCommand))
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
			_ = client.Interaction.DeferResponse(interaction.Id, interaction.Token, true)

			_ = cmd.Execute(&Context{Client: client, Interaction: interaction, CmdMgr: mgr})
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

	if _, err := mgr.client.Application.RegisterCommand(mgr.client.Me().Id, "", appCmd); err != nil {
		fmt.Println(err)
	}

	mgr.commands[cmd.Name()] = cmd
}

// ToDo : Unregister commands
