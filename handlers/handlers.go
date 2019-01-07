package handlers

import (
	"discord-bot/config"
	"discord-bot/util"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var startTime = time.Now()
var usersOnline map[string]struct{}

func init() {
	usersOnline = make(map[string]struct{})
}

// InitHandlers sets the handlers for the discord session to execute commands
func InitHandlers(session *discordgo.Session, cfg *config.Config) {

	util.LogInfo("Setting up event handlers")

	// Uptime
	session.AddHandler(func(s *discordgo.Session, e *discordgo.MessageCreate) {

		if e.Author.ID == s.State.User.ID {
			return
		}

		if strings.ToLower(strings.TrimSpace(e.Content)) == cfg.Commands["uptime"] {
			s.ChannelTyping(e.ChannelID)
			duration := time.Now().Sub(startTime)
			util.SendMessage(s, fmt.Sprintf(
				"Uptime is %02d:%02d:%02d (since %s)",
				int(duration.Hours()),
				int(duration.Minutes())%60,
				int(duration.Seconds())%60,
				startTime.Format(time.Stamp)))
			util.LogInfo("User", e.Author.Username, "triggered the !uptime command")
		}

	})

	// Ping
	session.AddHandler(func(s *discordgo.Session, e *discordgo.MessageCreate) {

		if e.Author.ID == s.State.User.ID {
			return
		}

		if strings.ToLower(strings.TrimSpace(e.Message.Content)) == cfg.Commands["ping"] {
			s.ChannelTyping(e.ChannelID)
			util.SendMessage(s, "Pong!")
			util.LogInfo("User", e.Author.Username, "triggered the !ping command")
		}

	})

	// Info
	session.AddHandler(func(s *discordgo.Session, e *discordgo.MessageCreate) {

		if e.Author.ID == s.State.User.ID {
			return
		}

		if strings.ToLower(strings.TrimSpace(e.Message.Content)) == cfg.Commands["info"] {
			s.ChannelTyping(e.ChannelID)
			util.SendMessage(s, cfg.Bot.Info)
			util.LogInfo("User", e.Author.Username, "triggered the !info command")
		}
	})

	// @todo !commands
	// should send a message with a link to the available commands
}
