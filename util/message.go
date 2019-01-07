package util

import (
	"github.com/bwmarrin/discordgo"
)

// SendMessage sends a message to the current discord session
func SendMessage(s *discordgo.Session, m string) {
	channelID := FetchPrimaryTextChannelID(s)
	LogInfo("SENDING MESSAGE:", m)
	RetryOnBadGateway(func() error {
		_, err := s.ChannelMessageSend(channelID, m)
		return err
	})
}
