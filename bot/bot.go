package bot

import (
	"fdcteam-bot/command"
	"fdcteam-bot/config"
	"fdcteam-bot/providers"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func StartBot() (err error) {
	log.Info("Iniciando FdcBot...")

	currentBotSession, err := discordgo.New("Bot " + config.DiscordToken)

	if err != nil {
		log.Error("Erro ao criar nova sessão na API do Discord: ", err)
		return
	}

	err = currentBotSession.Open()

	if err != nil {
		log.Error("Erro ao conectar com o Discord: ", err)
		return
	}

	log.Info("Bot iniciado com sucesso!")

	providers.SetBotSession(currentBotSession)

	printBotUserInfo()

	command.RegisterCommands()

	return
}

func printBotUserInfo() {

	botUser, err := providers.BotSession().User("@me")

	if err != nil {
		log.Error("Erro ao obter informações do usuário utilizado pelo Bot", err)
		return
	}

	log.Printf("ID do Bot: %s#%s\n", botUser.Username, botUser.Discriminator)
}
