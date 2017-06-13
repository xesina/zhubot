package commands

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Help struct{}

func (h *Help) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())
	bot.Send(msg)

	return nil
}
