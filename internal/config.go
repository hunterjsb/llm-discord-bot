package internal

import (
	"fmt"
	"os"
)

// Config holds application configuration settings
type Config struct {
	DiscordToken   string
	OpenAIAPIKey   string
	GuildID        string
	RemoveCommands bool
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		return nil, fmt.Errorf("DISCORD_TOKEN environment variable is not set")
	}

	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	if openAIAPIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	// Guild ID is optional - if not set, commands will be registered globally
	guildID := os.Getenv("GUILD_ID")

	// By default, remove commands when shutting down
	removeCommandsStr := os.Getenv("REMOVE_COMMANDS")
	removeCommands := true
	if removeCommandsStr == "false" {
		removeCommands = false
	}

	return &Config{
		DiscordToken:   discordToken,
		OpenAIAPIKey:   openAIAPIKey,
		GuildID:        guildID,
		RemoveCommands: removeCommands,
	}, nil
}