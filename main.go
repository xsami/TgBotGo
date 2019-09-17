package main

import (
	bot "TgBotGo/bot"
	helper "TgBotGo/bot/helpers"
	"log"
)

// main function start the application.
// First load the file .env with the enviroment variables
// and ask for the token to call the main bot initializer
func main() {

	if err := helper.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	token := helper.GetBotToken()

	log.Fatal(bot.StartBot(token))
}
