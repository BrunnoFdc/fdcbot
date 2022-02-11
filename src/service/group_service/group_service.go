package group_service

import (
	"fdcteam-bot/src/bot"
	"fdcteam-bot/src/config"
	"fdcteam-bot/src/domain/errors"
	"fdcteam-bot/src/util/member_util"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	allowedRoles = [...]string{
		"940638311993716777", // Bukkit
	}
)

func AddGroupToUser(user *discordgo.User, groupRole *discordgo.Role) *errors.BotError {
	if !isGroupRole(*groupRole) {
		return &errors.BotError{
			ErrorType: errors.InvalidGroupRole,
		}
	}

	err := bot.Session().GuildMemberRoleAdd(config.GuildId, user.ID, groupRole.ID)

	if err != nil {
		log.Errorf("Erro ao adicionar cargo para o membro %s: %s", member_util.GetFullDisplayName(user), err)

		return &errors.BotError{
			ErrorType:   errors.UnexpectedError,
			SourceError: err,
		}
	}

	log.Debugf("Adicionado cargo %s para o membro %s", groupRole.Name, member_util.GetFullDisplayName(user))

	return nil
}

func isGroupRole(role discordgo.Role) bool {
	for _, allowedRole := range allowedRoles {
		if role.ID == allowedRole {
			return true
		}
	}

	return false
}
