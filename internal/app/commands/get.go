package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	fmt.Println(idx)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("wrong args: %s", args))
		c.bot.Send(msg)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("successfuly parsed argument: %v", arg))
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*Commander).Get
}
