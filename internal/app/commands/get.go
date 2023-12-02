package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("wrong args: %s", args))
		c.bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("successfuly parsed argument: %v", arg))
	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*Commander).Get
}
