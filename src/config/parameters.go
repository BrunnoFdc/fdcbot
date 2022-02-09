package config

import "os"

var (
	DiscordToken         = os.Getenv("FDCBOT_DISCORD_TOKEN")
	DiscordApplicationId = os.Getenv("FDCBOT_APPLICATION_ID")
	GuildId              = os.Getenv("FDCBOT_GUILD_ID")
	MongodbConnectionUrl = os.Getenv("FDCBOT_MONGODB_CONNECTION_URL")
)
