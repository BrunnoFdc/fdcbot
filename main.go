package main

import (
	"fdcteam-bot/bot"
	"fdcteam-bot/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	setupLogger()

	botErr := bot.StartBot()

	if botErr != nil {
		panic(botErr)
	}

	<-make(chan struct{})
}

func setupLogger() {
	parsedLevel, err := log.ParseLevel(config.LogLevel)

	if err == nil {
		log.Infof("Log level setado para %s", config.LogLevel)
		log.SetLevel(parsedLevel)
	}
}
