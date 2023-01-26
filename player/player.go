package player

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord/embed"
	"github.com/Goscord/goscord/goscord/gateway"
	"go.uber.org/atomic"
	"sync"
)

type Player struct {
	sync.RWMutex

	client          *gateway.Session
	voiceConnection *gateway.VoiceConnection
	stop            chan bool
	playing         atomic.Bool
	paused          atomic.Bool

	guildId        string
	messageChannel string
	messageId      string
	channelId      string
	currentTrack   *Track
	queue          []*Track
}

func NewPlayer(client *gateway.Session, guildId, channelId, messageChannel string) *Player {
	player := &Player{
		client: client,
		stop:   make(chan bool),

		guildId:        guildId,
		channelId:      channelId,
		messageChannel: messageChannel,
		queue:          make([]*Track, 0),
	}

	AddPlayer(player)

	return player
}

func (p *Player) Play() error {
	p.RLock()
	client := p.client
	guildId := p.guildId
	channelId := p.channelId
	messageChannel := p.messageChannel
	p.RUnlock()

	voiceConnection, err := client.JoinVoiceChannel(guildId, channelId, false, true)
	if err != nil {
		return err
	}

	p.Lock()
	p.voiceConnection = voiceConnection
	p.Unlock()

	if !p.playing.Load() {
		for {
			p.RLock()
			queue := p.queue
			p.RUnlock()

			if len(queue) == 0 {
				break
			}

			p.RLock()
			track := p.queue[0]
			stop := p.stop
			messageId := p.messageId
			p.RUnlock()

			p.playing.Store(true)

			e := embed.NewEmbedBuilder()
			e.SetColor(embed.Green)
			e.SetTitle("💿 | Now playing")
			e.SetDescription(fmt.Sprintf("**%s** by %s", track.Title, track.Author))
			e.SetFooter(fmt.Sprintf("Requested by %s", track.Requester.Username), track.Requester.AvatarURL())

			if messageId == "" {
				m, _ := client.Channel.SendMessage(messageChannel, e.Embed())
				p.Lock()
				p.messageId = m.Id
				p.Unlock()
			} else {
				client.Channel.Edit(messageChannel, messageId, e.Embed())
			}

			PlayUrlOrFile(voiceConnection, track.StreamUrl, stop)

			p.playing.Store(false)

			p.NextTrack()
		}

		p.RLock()
		messageId := p.messageId
		p.RUnlock()

		e := embed.NewEmbedBuilder()
		e.SetColor(embed.Red)
		e.SetTitle("💿 | Queue finished!")
		e.SetDescription("The music queue has finished playing. You can add more tracks to the queue by using the `/play` command.")
		client.Channel.Edit(messageChannel, messageId, e.Embed())

		p.Stop()
	}

	return nil
}

func (p *Player) SkipTrack() {
	p.Lock()
	p.stop <- true
	p.Unlock()
}

func (p *Player) AddTrack(track *Track) {
	p.Lock()
	p.queue = append(p.queue, track)
	p.Unlock()
}

func (p *Player) RemoveTrack(index int) {
	p.Lock()
	p.queue = append(p.queue[:index], p.queue[index+1:]...)
	p.Unlock()
}

func (p *Player) NextTrack() {
	p.RemoveTrack(0)

	p.Lock()
	if len(p.queue) > 0 {
		p.currentTrack = p.queue[0]
	}
	p.Unlock()
}

func (p *Player) ClearQueue() {
	p.Lock()
	p.queue = make([]*Track, 0)
	p.Unlock()
}

func (p *Player) Stop() {
	p.RLock()
	guildId := p.guildId
	voiceConnection := p.voiceConnection
	stop := p.stop
	p.RUnlock()

	voiceConnection.Disconnect()
	stop <- true

	p.playing.Store(false)

	RemovePlayer(guildId)
}

func (p *Player) Pause() {
	p.paused.Store(true)
}

func (p *Player) IsPlaying() bool {
	return p.playing.Load()
}

func (p *Player) IsPaused() bool {
	return p.paused.Load()
}

func (p *Player) Queue() []*Track {
	p.RLock()
	queue := p.queue
	p.RUnlock()

	return queue
}

func (p *Player) CurrentTrack() *Track {
	p.RLock()
	track := p.currentTrack
	p.RUnlock()

	return track
}

func (p *Player) VoiceConnection() *gateway.VoiceConnection {
	p.RLock()
	voiceConnection := p.voiceConnection
	p.RUnlock()

	return voiceConnection
}

func (p *Player) GuildId() string {
	p.RLock()
	guildId := p.guildId
	p.RUnlock()

	return guildId
}

func (p *Player) ChannelId() string {
	p.RLock()
	channelId := p.channelId
	p.RUnlock()

	return channelId
}

func (p *Player) MessageChannel() string {
	p.RLock()
	messageChannel := p.messageChannel
	p.RUnlock()

	return messageChannel
}
