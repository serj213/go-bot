package commands

import (
	"go-bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (c *Commander) List(inputMessage *tgbotapi.Message, service *product.Service){
	outputMsgText := "The list your products: \n"
	for _, t := range service.List(){
		outputMsgText += t.Title
		outputMsgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}