package main

import (
	"fdcteam-bot/bot"
	"fdcteam-bot/config"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	setupDotEnv()

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

func setupDotEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Error("Erro ao carregar as variáveis de ambiente", err)
		panic(err)
	}
}
