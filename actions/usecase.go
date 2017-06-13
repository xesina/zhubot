package actions

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Usecase struct {
	Keywords []string
}

var Usecasekeywords = []string{
	"Usecases",
}

func NewUsecaseAction() *Usecase {
	return &Usecase{
		Keywords: Usecasekeywords,
	}
}

func (a *Usecase) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	var kb = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Usage"),
			tgbotapi.NewKeyboardButton("Prefer to ask"),
		),
	)
	msg.ReplyMarkup = kb
	msg.Text = `Faq bots, customers service and many other usecases.`
	bot.Send(msg)
	return nil
}

func (a *Usecase) GetKeywords() []string {
	return a.Keywords
}
