package bot

import (
	"fdcteam-bot/src/command"
	"fdcteam-bot/src/config"
	"github.com/bwmarrin/discordgo"
	scm "github.com/ethanent/discordgo-scm"
	log "github.com/sirupsen/logrus"
)

func StartBot() (startedBot *discordgo.Session, err error) {
	log.Info("Iniciando FdcBot...")

	startedBot, err = discordgo.New("Bot " + config.DiscordToken)

	if err != nil {
		log.Error("Erro ao criar nova sessão na API do Discord: ", err)
		return
	}

	err = startedBot.Open()

	if err != nil {
		log.Error("Erro ao conectar com o Discord: ", err)
		return
	}

	log.Info("Bot iniciado com sucesso!")

	printBotUserInfo(startedBot)
	configureScm(startedBot)

	return
}

func printBotUserInfo(bot *discordgo.Session) {
	botUser, err := bot.User("@me")

	if err != nil {
		log.Error("Erro ao obter informações do usuário utilizado pelo Bot", err)
		return
	}

	log.Printf("ID do Bot: %s#%s\n", botUser.Username, botUser.Discriminator)
}

func configureScm(bot *discordgo.Session) {
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
