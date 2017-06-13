package commands

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Unknown struct{}

func (s *Unknown) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)


	msg.Text = `Unknown command, need help?`

	bot.Send(msg)

	return nil
}
