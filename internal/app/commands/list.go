package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputmsg := "Here all the products:\n\n"
	products := c.productService.List()
	for _, p := range products {
		outputmsg += p.Title
		outputmsg += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputmsg)
	c.bot.Send(msg)
}
