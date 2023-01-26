package player

var players = make(map[string]*Player)

func PlayerByGuild(guildId string) (*Player, bool) {
	player, ok := players[guildId]
	return player, ok
}

func AddPlayer(player *Player) {
	players[player.guildId] = player
}

func RemovePlayer(guildId string) {
	delete(players, guildId)
}
