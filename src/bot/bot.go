package bot

import (
	"fdcteam-bot/src/command"
	"fdcteam-bot/src/config"
	"github.com/bwmarrin/discordgo"
	scm "github.com/ethanent/discordgo-scm"
	log "github.com/sirupsen/logrus"
)

var currentBotSession *discordgo.Session

func Session() *discordgo.Session {
	return currentBotSession
}

func StartBot() (err error) {
	log.Info("Iniciando FdcBot...")

	currentBotSession, err = discordgo.New("Bot " + config.DiscordToken)

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

	printBotUserInfo()
	configureScm()

	return
}

func printBotUserInfo() {

	botUser, err := Session().User("@me")

	if err != nil {
		log.Error("Erro ao obter informações do usuário utilizado pelo Bot", err)
		return
	}

	log.Printf("ID do Bot: %s#%s\n", botUser.Username, botUser.Discriminator)
}

func configureScm() {
	bot := Session()
	manager := scm.NewSCM()

	for _, commandFeature := range command.CurrentCommands().ToFeature() {
		manager.AddFeature(commandFeature)
	}

	err := manager.CreateCommands(bot, config.GuildId)

	if err != nil {
		log.Error("Não foi possível registrar comandos no servidor", err)
		return
	}

	bot.AddHandler(manager.HandleInteraction)
}
