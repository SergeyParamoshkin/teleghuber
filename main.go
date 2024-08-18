package main

import (
	"github.com/amarnathcjd/gogram/telegram"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	appID    int64
	appHash  string
	botToken string
)

func GetToken(configPath string) error {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("LOGLEVEL", "debug")

	if configPath != "" {
		log.Infof("Parsing config: %s", configPath)
		viper.SetConfigFile(configPath)
		err := viper.ReadInConfig()
		if err != nil {
			log.Warnf("Unable to read config file: %s", err)
		}
	} else {
		log.Warnf("Config file is not specified.")
	}

	logLevelString := viper.GetString("loglevel")
	logLevel, err := log.ParseLevel(logLevelString)
	if err != nil {
		log.Errorf("Unable to parse loglevel: %s", logLevelString)
	}

	log.SetLevel(logLevel)

	botToken = viper.GetString("BOT_TOKEN")
	if botToken == "" {
		log.Errorf("Bot token error")
		os.Exit(2)
	}
	appID = viper.GetInt64("APP_ID")
	appHash = viper.GetString("APP_HASH")

	return nil
}

func main() {

	err := GetToken(".env")
	if err != nil {
		log.Panicf("GetEnv error")
	}

	// create a new client object
	client, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:    int32(appID),
		AppHash:  appHash,
		LogLevel: telegram.LogInfo,
	})

	client.LoginBot(botToken)

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

	// lock the main routine
	client.Idle()
}
