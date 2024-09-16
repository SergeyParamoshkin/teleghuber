package main

import (
	"github.com/amarnathcjd/gogram/telegram"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type config struct {
	AppID    int64  `env:"TELEGRAM_APP_ID"`
	AppHash  string `env:"TELEGRAM_APP_HASH"`
	BotToken string `env:"TELEGRAM_BOT_TOKEN"`
}

func main() {
	var config config
	if err := env.Parse(&config); err != nil {
		log.WithError(err).Fatal("Failed to parse environment variables")
	}

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:    int32(config.AppID),
		AppHash:  config.AppHash,
		LogLevel: telegram.LogInfo,
	})
	if err != nil {
		log.WithError(err).Fatal("Failed to create Telegram client")
	}

	client.LoginBot(config.BotToken)

	client.On(telegram.OnMessage, messageHandler,
		telegram.FilterPrivate)

	client.On("message:/start", startHandler)

	client.Idle()
}

func messageHandler(message *telegram.NewMessage) error {
	message.Respond(message.Text())
	log.Infof(message.Text())
	return nil
}

func startHandler(message *telegram.NewMessage) error {
	message.Reply("Hello, I am a bot!")
	return nil
}
