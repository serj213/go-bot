package commands

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message){

	idx := strings.Split(inputMessage.Text, " ")

	if len(idx) == 1{
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Введите id элемента, который желаете удалить")
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

	elem, err := c.productService.Delete(int(num))
	if err != nil {
		fmt.Println("error found product: %s", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Продукт не найден")
		c.bot.Send(msg)
		return
	}

	outputMsgText := "Продукт: " + elem.Title + " был успешно удален \n"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}