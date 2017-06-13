package actions

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Monetize struct {
	Keywords []string
}

var MonetizeKeywords = []string{
	"Monetization",
}

func NewMonetizeAction() *Monetize {
	return &Monetize{
		Keywords: MonetizeKeywords,
	}
}

func (a *Monetize) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	msg.Text = `724 supporting your users`
	bot.Send(msg)

	msg.Text = `Ask about zhubot`
	bot.Send(msg)

	return nil
}

func (a *Monetize) GetKeywords() []string {
	return a.Keywords
}
