package commands

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Faq struct{}

func (s *Faq) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	var kb = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("What is zhubot?"),

		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("What is zhubot usecases?"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("How to use zhubot?"),
		),
	)

	msg.ReplyMarkup = kb

	bot.Send(msg)

	return nil
}
