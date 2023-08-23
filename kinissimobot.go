package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Retrieve the Telegram bot token from an environment variable.
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN not set")
	}
	// Create a new bot instance.
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Set bot options.
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up an update configuration.
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Get updates from the bot.
	updates, err := bot.GetUpdatesChan(u)
	// Process incoming updates.
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Echo the received message.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		bot.Send(msg)
	}
}
