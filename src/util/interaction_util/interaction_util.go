package interaction_util

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func AnswerInteraction(session *discordgo.Session, interaction *discordgo.Interaction, responseContent string) {
	err := session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: responseContent,
		},
	})

	if err != nil {
		log.Error("Erro ao responder ao usuário: ", err)
	}
}
