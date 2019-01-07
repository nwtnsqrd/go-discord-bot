package util

import (
	"github.com/bwmarrin/discordgo"
)

// FetchUser retrieves a user for the given userID from the discord session
func FetchUser(s *discordgo.Session, userID string) *discordgo.User {

	var res *discordgo.User

	RetryOnBadGateway(func() error {

		var err error
		res, err = s.User(userID)
		if err != nil {
			return err
		}
		return nil
	})

	return res
}
