package commands

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Start struct{}

func (s *Start) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	var kb = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("About Zhubot"),
			tgbotapi.NewKeyboardButton("Prefe to ask"),
		),
	)

	msg.ReplyMarkup = kb
	msg.Text = `Hey, lets discover zhubot!`

	bot.Send(msg)

	return nil
}
