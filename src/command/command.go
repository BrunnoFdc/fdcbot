package command

import (
	"fdcteam-bot/src/config"
	"github.com/bwmarrin/discordgo"
	discordgo_scm "github.com/ethanent/discordgo-scm"
)

type Command struct {
	Name        string
	Description string
	Version     string
	Options     []*discordgo.ApplicationCommandOption
	Handler     func(session *discordgo.Session, interaction *discordgo.InteractionCreate)
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
		Type:               discordgo.InteractionApplicationCommand,
		Handler:            command.Handler,
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

func CurrentCommands() CommandList {
	return CommandList{
		GroupCommand(),
	}
}
