package main

import (
	"fmt"
	"llm-discord-bot/internal"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		// Try relative path from cmd/bot directory
		err = godotenv.Load("../../.env")
		if err != nil {
			fmt.Printf("Warning: Error loading .env file: %v\n", err)
			fmt.Println("Continuing with environment variables from system...")
		}
	}

	// Load configuration
	config, err := internal.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Create and start bot
	bot, err := internal.NewDiscordBot(config)
	if err != nil {
		fmt.Printf("Error creating bot: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting Discord bot...")
	err = bot.Start()
	if err != nil {
		fmt.Printf("Error starting bot: %v\n", err)
		os.Exit(1)
	}

	// Set up graceful shutdown
	internal.SetupCloseHandler(func() error {
		fmt.Println("Shutting down bot...")
		return bot.Stop()
	})

	// Block main goroutine indefinitely
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	select {}
}