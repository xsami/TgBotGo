package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartBot this function to start the bot
func StartBot(token string) error {

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			if err = commandHandler(update, bot); err != nil { // Handle the command message
				return err
			}
			continue
		}

		// Do something else with the message received
		if err = messageHandler(update, bot); err != nil { // Handle the single message
			return err
		}

	}

	return nil
}
