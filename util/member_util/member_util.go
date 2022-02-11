package member_util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func GetFullDisplayName(user *discordgo.User) string {
	return fmt.Sprintf("%s#%s", user.Username, user.Discriminator)
}
