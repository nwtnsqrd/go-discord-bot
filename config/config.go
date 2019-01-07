package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is the wrapper for the bot's configuration
type Config struct {
	Bot      *Bot
	Commands map[string]string
}

// Bot defines the generic bot structure
type Bot struct {
	Name  string `json:"bot_name"`
	Token string `json:"bot_token"`
	Info  string `json:"bot_info"`
}

// GetConfig populates the Bot structure
func GetConfig() *Config {

	bs, err := ioutil.ReadFile("./config/config.json")

	if err != nil {
		fmt.Println("Error loading the config file:", err)
		os.Exit(-1)
	}

	var config Config
	err = json.Unmarshal(bs, &config)

	if err != nil {
		fmt.Println("Error reading the config file:", err)
		os.Exit(-2)
	}

	return &config
}
