package bot

import (
	msg "TgBotGo/bot/messages"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var regularMsgs = map[string]interface{}{
	"hi": msg.Greet}

func messageHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {

	message := update.Message.Text

	var responseQ []interface{} // contain the message received

	// Store all the possible response that can be peformed based in the regex
	for keyRegex, valueRegex := range regularMsgs {
		if ok, _ := regexp.MatchString(keyRegex, message); ok {
			responseQ = append(responseQ, valueRegex)
		}
	}

	// Take a decition about which actions will be performed.
	// For this example I will only executed the last action found.
	if resQLen := len(responseQ); resQLen > 0 {
		responseQ[resQLen-1].(func(int64, *tgbotapi.BotAPI, string))(update.Message.Chat.ID, bot, message)
	}

	return nil
}
