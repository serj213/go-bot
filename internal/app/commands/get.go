package commands

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (c *Commander) Get(inputMessage *tgbotapi.Message) {

	idx := strings.Split(inputMessage.Text, " ")

	if len(idx) == 1{
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Введите id элемента, который желаете получить")
		c.bot.Send(msg)
		return
	}


	num, err := strconv.ParseInt(idx[1], 10, 64)
	if err != nil {
		fmt.Println("parse idx delete commands:  %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Передайте число")
		c.bot.Send(msg)
		return
	}

	product, err := c.productService.Get(int(num))

	if err != nil {
		fmt.Println("fail get product: %s", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Введите новое значение")
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}