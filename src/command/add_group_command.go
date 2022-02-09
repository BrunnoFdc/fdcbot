package command

import (
	"fdcteam-bot/src/util/interaction_util"
	"fdcteam-bot/src/util/member_util"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	allowedRoles = []string{
		"940638311993716777", // Bukkit
	}
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

func handleCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	rootInteraction := interaction.Interaction
	user := interaction.Member.User

	option := interaction.ApplicationCommandData().Options[0]
	informedRole := option.RoleValue(session, interaction.GuildID)

	if isRoleAllowed(*informedRole) {
		err := session.GuildMemberRoleAdd(interaction.GuildID, user.ID, informedRole.ID)

		if err != nil {
			interaction_util.AnswerInteraction(
				session, rootInteraction, "Erro inesperado: Não foi possível adicionar o cargo.")

			log.Errorf("Erro ao adicionar cargo para o membros %s: %s",
				member_util.GetFullDisplayName(user), err)
			return
		}

		log.Debugf("Adicionado cargo para o membro %s", member_util.GetFullDisplayName(user))

		interaction_util.AnswerInteraction(session, rootInteraction, "Cargo adicionado com sucesso!")

		return
	}

	interaction_util.AnswerInteraction(session, rootInteraction, "Não é possível entrar neste grupo.")
}

func isRoleAllowed(role discordgo.Role) bool {
	for _, allowedRole := range allowedRoles {
		if role.ID == allowedRole {
			return true
		}
	}

	return false
}
