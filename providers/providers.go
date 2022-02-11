package providers

import "github.com/bwmarrin/discordgo"

var (
	botSession *discordgo.Session
)

func SetBotSession(newBotSession *discordgo.Session) {
	botSession = newBotSession
}

func BotSession() *discordgo.Session {
	return botSession
}
