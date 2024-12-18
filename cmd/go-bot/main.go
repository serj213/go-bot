package main

import (
	"go-bot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command(){
			case "help":
				helpMessage(bot, update.Message)
			case "list":
				listMessage(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpMessage(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text + "-" + inputMessage.Command())
	bot.Send(msg)
}

func listMessage(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, service *product.Service){
	outputMsgText := "The list your products: \n"
	for _, t := range service.List(){
		outputMsgText += t.Title
		outputMsgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	bot.Send(msg)
}