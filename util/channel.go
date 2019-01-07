package util

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

// FetchPrimaryTextChannelID retrieves the primary text channel ID of the given session
func FetchPrimaryTextChannelID(sess *discordgo.Session) string {
	var channelID string
	RetryOnBadGateway(func() error {
		guilds, err := sess.UserGuilds(1, "", "")
		if err != nil {
			return err
		}
		guild, err := sess.Guild(guilds[0].ID)
		if err != nil {
			return err
		}
		channels, err := sess.GuildChannels(guild.ID)
		if err != nil {
			return err
		}
		for _, channel := range channels {
			channel, err = sess.Channel(channel.ID)
			if err != nil {
				return err
			}
			if channel.Type == discordgo.ChannelTypeGuildText {
				channelID = channel.ID
				return nil
			}
		}
		return errors.New("No primary channel found")
	})
	return channelID
}
