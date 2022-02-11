package command

import (
	"fdcteam-bot/config"
	"fdcteam-bot/providers"
	"github.com/bwmarrin/discordgo"
	discordgo_scm "github.com/ethanent/discordgo-scm"
	log "github.com/sirupsen/logrus"
)

var currentCommands = CommandList{
	GroupCommand(),
}

type Command struct {
	Name        string
	Description string
	Version     string
	Options     []*discordgo.ApplicationCommandOption
	Handler     func(interaction *discordgo.InteractionCreate)
}

type CommandList []Command

func (command Command) ToFeature() *discordgo_scm.Feature {
	commandInfo := &discordgo.ApplicationCommand{
		ApplicationID: config.DiscordApplicationId,
		Type:          discordgo.ChatApplicationCommand,
		Name:          command.Name,
		Description:   command.Description,
		Version:       command.Version,
		Options:       command.Options,
	}

	return &discordgo_scm.Feature{
		Type: discordgo.InteractionApplicationCommand,
		Handler: func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			command.Handler(interaction)
		},
		ApplicationCommand: commandInfo,
	}
}

func (commandList CommandList) ToFeature() []*discordgo_scm.Feature {
	var commandsAsFeatures []*discordgo_scm.Feature

	for _, command := range commandList {
		commandsAsFeatures = append(commandsAsFeatures, command.ToFeature())
	}

	return commandsAsFeatures
}

func RegisterCommands() {

	bot := providers.BotSession()
	manager := discordgo_scm.NewSCM()

	for _, commandFeature := range currentCommands.ToFeature() {
		manager.AddFeature(commandFeature)
	}

	err := manager.CreateCommands(bot, config.GuildId)

	if err != nil {
		log.Error("Não foi possível registrar comandos no servidor", err)
		return
	}

	bot.AddHandler(manager.HandleInteraction)
}
