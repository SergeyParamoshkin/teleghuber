package main

import (
	"fmt"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type config struct {
	AppID    int64  `env:"APP_ID"`
	AppHash  string `env:"APP_HASH"`
	BotToken string `env:"BOT_TOKEN"`
}

func main() {

	var cfg config
	err := env.Parse(&cfg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)

	// create a new client object
	client, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:    int32(cfg.AppID),
		AppHash:  cfg.AppHash,
		LogLevel: telegram.LogInfo,
	})

	client.LoginBot(cfg.BotToken)

	client.On(telegram.OnMessage, func(message *telegram.NewMessage) error {
		message.Respond(message)
		log.Infof(message.Text())
		return nil
	},
		telegram.FilterPrivate)

	client.On("message:/start", func(message *telegram.NewMessage) error {
		message.Reply("Hello, I am a bot!")
		return nil
	})

	// // lock the main routine
	client.Idle()
}
