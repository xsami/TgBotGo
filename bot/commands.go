package bot

import (
	command "TgBotGo/bot/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var commandList = map[string]interface{}{
	"help": command.Help}

func cmdHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {

	if update.Message.IsCommand() {

		commandReq := update.Message.Command()
		commandArgs := update.Message.CommandArguments()

		if _, ok := commandList[commandReq]; !ok {
			// Make a default action ?
			return nil
		}

		commandList[commandReq].(func(int64, *tgbotapi.BotAPI, string))(update.Message.Chat.ID, bot, commandArgs)

	}
	return nil
}

func commandHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	return cmdHandler(update, bot)
}
