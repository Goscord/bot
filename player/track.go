package player

import "github.com/Goscord/goscord/goscord/discord"

type Track struct {
	Title     string
	Author    string
	StreamUrl string
	Requester *discord.User
}
