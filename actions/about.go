package actions

import "github.com/go-telegram-bot-api/telegram-bot-api"

type About struct {
	Keywords []string
}

var keywords = []string{
	"About Zhubot",
}

func NewAboutAction() *About {
	return &About{
		Keywords: keywords,
	}
}

func (a *About) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	var kb = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Zhubot usecases"),
			tgbotapi.NewKeyboardButton("Prefer to Ask"),
		),
	)
	msg.ReplyMarkup = kb
	msg.Text = `This is an example of smart bot :)`
	bot.Send(msg)
	return nil
}

func (a *About) GetKeywords() []string {
	return a.Keywords
}
