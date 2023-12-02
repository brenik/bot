package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github/brenik/bot/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

//func (c *Commander) Help(inputMessage *tgbotapi.Message) {
//	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help\n /list - lust products")
//	c.bot.Send(msg)
//}

//func (c *Commander) List(inputMessage *tgbotapi.Message) {
//	outputmsg := "Here all the products:\n\n"
//	products := c.productService.List()
//	for _, p := range products {
//		outputmsg += p.Title
//		outputmsg += "\n"
//	}
//	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputmsg)
//	c.bot.Send(msg)
//}

//func (c *Commander) Default(inputMessage *tgbotapi.Message) {
//	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
//	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
//	c.bot.Send(msg)
//
//}
