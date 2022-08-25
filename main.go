package main

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"log"
	"os"
)

type application struct {
	client *tbot.Client
}

var (
	app   application
	bot   *tbot.Server
	token string
)

func main() {
	launch()
}

func launch() {
	//var envs map[string]string
	//envs, err := godotenv.Read(".env")
	//
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	bot = tbot.New(os.Getenv("token"))
	app.client = bot.Client()
	handlers()
	log.Fatal(bot.Start())
}

func handlers() {
	bot.HandleMessage(".*yo.*", app.catchYOHandler)
	bot.HandleMessage("/start", app.startHandler)
}

func (a *application) startHandler(m *tbot.Message) {
	message := "This bot is >>>>"
	a.client.SendMessage(m.Chat.ID, message)
}

func (a *application) catchYOHandler(m *tbot.Message) {
	app.client.SendChatAction(m.Chat.ID, tbot.ActionTyping)
	app.client.SendMessage(m.Chat.ID, "hello!")
	fmt.Println(m.Chat)
}
