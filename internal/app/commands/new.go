package commands

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)



func (c *Commander) New(inputMessage *tgbotapi.Message) {

	arrInputMessageText := strings.Split(inputMessage.Text, " ")

	if len(arrInputMessageText) < 2 {
		fmt.Printf("arrInputMessageText new: %v", arrInputMessageText)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Введите имя элемента, который желаете добавить")
		c.bot.Send(msg)
		return
	}

	newElemName := strings.Join(arrInputMessageText[1:], " ")

	product := c.productService.New(newElemName)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title + " был успешно добавлен")
	c.bot.Send(msg)
}		