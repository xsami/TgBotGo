package messages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Greet this function send a random
// greet to the bot
func Greet(chatID int64, bot *tgbotapi.BotAPI, message string) {

	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = "Hiiiii!"

	bot.Send(msg)
}
