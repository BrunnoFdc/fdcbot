package command

import (
	"fdcteam-bot/src/bot"
	"fdcteam-bot/src/service/group_service"
	"fdcteam-bot/src/util/interaction_util"
	"github.com/bwmarrin/discordgo"
)

func GroupCommand() Command {
	return Command{
		Name:        "grupo",
		Description: "Entrar em um grupo e obter a permissão para acessar a sessão do servidor específica daquele assunto.",
		Version:     "1",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "cargo",
				Description: "Mencione o cargo do grupo que você deseja entrar.",
				Type:        discordgo.ApplicationCommandOptionRole,
				Required:    true,
			},
		},
		Handler: handleCommand,
	}
}

func handleCommand(commandInteraction *discordgo.InteractionCreate) {
	interaction := commandInteraction.Interaction
	user := commandInteraction.Member.User

	option := commandInteraction.ApplicationCommandData().Options[0]
	informedRole := option.RoleValue(bot.Session(), commandInteraction.GuildID)

	err := group_service.AddGroupToUser(user, informedRole)

	if err != nil {
		interaction_util.AnswerInteraction(interaction, err.Message())
	} else {
		interaction_util.AnswerInteraction(interaction, "Cargo adicionado com sucesso!")
	}
}
