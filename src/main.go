package main

import (
	"fdcteam-bot/src/bot"
	"fdcteam-bot/src/database"
)

func main() {
	dbErr := database.Connect()

	if dbErr != nil {
		panic(dbErr)
	}

	_, botErr := bot.StartBot()

	if botErr != nil {
		panic(botErr)
	}

	<-make(chan struct{})
}
