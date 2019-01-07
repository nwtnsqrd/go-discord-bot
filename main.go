package main

import (
	"discord-bot/config"
	"discord-bot/handlers"
	"nwtn-lib/errors"
	"nwtn-lib/logging"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var cfg = config.GetConfig()

func main() {

	discord, err := discordgo.New("Bot " + cfg.Bot.Token)
	errors.RaiseIfExists("Error creating new discord session", err, true)

	handlers.InitHandlers(discord, cfg)

	err = discord.Open()
	errors.RaiseIfExists("Error opening the discord session", err, true)

	logging.Info(cfg.Bot.Name, "is running")

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-channel

	discord.Close()
	logging.Info(cfg.Bot.Name, "is shutting down")
}
