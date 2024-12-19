package commands

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Edit(inputMessage *tgbotapi.Message) {
	arrText := strings.Split(inputMessage.Text, " ")

	if len(arrText) < 3 {
		fmt.Printf("/edit has small arguments: %v \n", arrText)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Для редактирование введите команду в формате \n /edit индекс_элемента новое_имя")
		c.bot.Send(msg)
		return
	}

	idx := arrText[1]
	newName := arrText[2]

	num, err := strconv.ParseInt(idx, 10, 64)
	if err != nil {
		fmt.Println("parse idx edit commands:  %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Передайте число в качестве индекса элемента")
		c.bot.Send(msg)
		return
	}

	elem, err := c.productService.Edit(int(num), newName)

	if err != nil{
		fmt.Printf("error edit product %v /n", err)
	}

	outputMsg := "Продукт был успешно изменен: " + elem.Title
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	c.bot.Send(msg)
	
}