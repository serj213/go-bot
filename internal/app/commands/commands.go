package commands

import (
	"go-bot/internal/service/product"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)



type Commander struct {
	bot *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, service *product.Service) *Commander{

	return &Commander{
		bot: bot,
		productService: service,
	}
}


func (c *Commander) HandleUpdate(updates tgbotapi.UpdatesChannel, products *product.Service){

	defer func(){
		if err := recover(); err != nil{
			log.Printf("Recovered application: %v", err)
		}
	}()

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command(){
			case "help":
				c.Help(update.Message)
			case "list":
				c.List(update.Message, products)
			case "get":
				c.Get(update.Message)

			case "delete": 
				c.Delete(update.Message)

			case "new":
				c.New(update.Message)	
			default:
				c.Default(update.Message)
			}
		}
	}
}