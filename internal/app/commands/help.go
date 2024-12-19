package commands

import (
	"encoding/json"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HelpFile struct {
	Title string `json:"title"`
	Commands []Command `json:"commands"`
}

type Command struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Command string `json:"command"`
}

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	
	file, err := os.ReadFile("./internal/app/commands/data/help.json")
	if err != nil{
		log.Fatalf("Error reads file with list commands: %v", err)
		return
	}

	var helpFile HelpFile
	
	if err := json.Unmarshal(file, &helpFile); err != nil{
		log.Fatalf("Error parse file of list commands: %v", err )
		return
	}

	
	var commansList []Command = helpFile.Commands

	outputMsg := helpFile.Title + "\n\n"

	for _, commandItem := range commansList{
		outputMsg += commandItem.Command +
			" - " + commandItem.Description + "\n"
	}
	

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	c.bot.Send(msg)
}