package interaction_util

import (
	"fdcteam-bot/src/bot"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func AnswerInteraction(interaction *discordgo.Interaction, responseContent string) {
	err := bot.Session().InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: responseContent,
		},
	})

	if err != nil {
		log.Error("Erro ao responder ao usuário: ", err)
	}
}
