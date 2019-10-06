package bot

import (
	"log"
	"regexp"

	msg "github.com/xsami/TgBotGo/bot/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var regularMsgs = map[string]interface{}{
	"hi": msg.Greet}

func messageHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {

	message := update.Message.Text

	var responseQ []interface{} // contain the message received

	// Store all the possible response that can be performed based in the regex
	for keyRegex, valueRegex := range regularMsgs {
		if ok, _ := regexp.MatchString(keyRegex, message); ok {
			responseQ = append(responseQ, valueRegex)
		}
	}

	// Take a decition about which actions will be performed.
	// For this example I will only executed the last action found.
	if resQLen := len(responseQ); resQLen > 0 {
		err := responseQ[resQLen-1].(func(int64, *tgbotapi.BotAPI, string) error)(update.Message.Chat.ID, bot, message)

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
