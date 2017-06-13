package actions

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Usage struct {
	Keywords []string
}

var UsageKeywords = []string{
	"Usage",
}

func NewUsageAction() *Usage {
	return &Usage{
		Keywords: UsageKeywords,
	}
}

func (a *Usage) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	var kb = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Monetization"),
			tgbotapi.NewKeyboardButton("Prefer to ask"),
		),
	)
	msg.ReplyMarkup = kb
	msg.Text = `You can train your bot in wit.ai`
	bot.Send(msg)
	return nil
}

func (a *Usage) GetKeywords() []string {
	return a.Keywords
}
