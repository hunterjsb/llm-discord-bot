# Discord LLM Chatbot

A lightweight Discord bot written in Go that uses OpenAI's API to respond to messages via slash commands.

## Features

- Uses modern Discord Slash Commands for easy interaction
- Integrates with OpenAI's GPT API to generate responses
- Handles long responses by splitting them into multiple messages
- Can be tested in a specific guild or deployed globally

## Setup

1. Clone this repository
2. Set up environment variables:
   ```
   export DISCORD_TOKEN="your_discord_bot_token"
   export OPENAI_API_KEY="your_openai_api_key"
   export GUILD_ID="your_test_guild_id"  # Optional: for testing in a specific server
   export REMOVE_COMMANDS="true"         # Optional: defaults to true
   ```
3. Build and run the bot:
   ```
   go build -o bot ./cmd/bot
   ./bot
   ```

## Discord Bot Setup

1. Create a Discord application at [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a bot for your application
3. Enable the following Privileged Gateway Intents:
   - Server Members Intent
   - Message Content Intent
4. Generate and copy your bot token for use with this application
5. Use the OAuth2 URL Generator with the following scopes:
   - `bot`
   - `applications.commands`
6. Add the following bot permissions:
   - Send Messages
   - Use Slash Commands
7. Use the generated URL to invite the bot to your server

## Usage

The bot registers a `/chat` slash command that you can use in any channel where the bot has access:

```
/chat prompt: Your message to the AI
```

## Requirements

- Go 1.16 or later
- Discord bot token from the [Discord Developer Portal](https://discord.com/developers/applications)
- OpenAI API key from the [OpenAI Dashboard](https://platform.openai.com/account/api-keys)

## License

MIT