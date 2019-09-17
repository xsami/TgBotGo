package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Help display the help for the bot
func Help(chatID int64, bot *tgbotapi.BotAPI, arguments string) error {

	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = "this is the help"

	_, err := bot.Send(msg)

	return err
}
