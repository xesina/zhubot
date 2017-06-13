package main

import (
	"fmt"
	"github.com/botanio/sdk/go"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xesina/zhubot/actions"
	"github.com/xesina/zhubot/commands"
	"github.com/xesina/zhubot/db"
	"log"
)

type actionInterface interface {
	Handle(*tgbotapi.BotAPI, tgbotapi.Update) error
	GetKeywords() []string
}

var actionsRegistry []actionInterface
var askAction *actions.Ask
var Botan botan.Botan

func init() {
	//yandex appmetrica token
	Botan = botan.New("YOUR_BOTAN_TOKEN")

	actionsRegistry = append(actionsRegistry, actions.NewAboutAction())
	actionsRegistry = append(actionsRegistry, actions.NewUsecaseAction())
	actionsRegistry = append(actionsRegistry, actions.NewUsageAction())
	actionsRegistry = append(actionsRegistry, actions.NewMonetizeAction())

	askAction = actions.NewAskAction()
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	db.InitDB()
	defer db.DB.Close()

	bot, err := tgbotapi.NewBotAPI("YOU_TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		//no input
		if update.Message == nil {
			continue
		}

		proccessAnalytics(update)

		//handle commands
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				start := commands.Start{}
				err = start.Handle(bot, update)
				if err != nil {
					log.Println(err)
				}
			case "help":
				help := commands.Help{}
				err = help.Handle(bot, update)
				if err != nil {
					log.Println(err)
				}
			case "faq":
				faq := commands.Faq{}
				err = faq.Handle(bot, update)
				if err != nil {
					log.Println(err)
				}
			default:
				unknown := commands.Unknown{}
				err = unknown.Handle(bot, update)
				if err != nil {
					log.Println(err)
				}
			}

			continue
		}

		//hanlde actions aka phrases
		actionProccessed := false
		for _, action := range actionsRegistry {
			keywords := action.GetKeywords()
			if stringInSlice(update.Message.Text, keywords) {
				err = action.Handle(bot, update)
				actionProccessed = true
				if err != nil {
					log.Println(err)
				}

				//need break??
			}
		}

		//default action ask (ask from wit)
		if !actionProccessed {
			err = askAction.Handle(bot, update)
			if err != nil {
				log.Println(err)
			}
			log.Println("Ask Action Triggered.")
		}

	}
}

func proccessAnalytics(update tgbotapi.Update) {
	// Synchronous track example
	rsp, _ := Botan.Track(update.Message.From.ID, *update.Message, update.Message.Text)
	fmt.Printf("Synchronous: %+v\n", rsp)
}
