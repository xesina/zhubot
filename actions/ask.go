package actions

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xesina/zhubot/db"
	"strconv"
)

type Ask struct {
	Keywords []string
}

func NewAskAction() *Ask {
	return &Ask{
		Keywords: nil,
	}
}

func (a *Ask) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	if update.Message.ReplyToMessage != nil {

		// TODO: Check if an operator was replied, not users, we should handle user replies properly

		// TODO: Check if ticket already answered and closed by same or other operators

		reply := db.FindReplyByMessageID(int64(update.Message.ReplyToMessage.MessageID))
		t := db.FindTicketByTicketID(reply.TicketID)

		reply.OperatorReply.String = update.Message.Text
		err := db.UpdateReply(reply)
		if err != nil {
			logrus.Fatal("Error updating reply")
		}

		err = db.UpdateTicket(t)
		if err != nil {
			logrus.Fatal("Error updating ticket")
		}

		msg.ChatID = t.ChatID
		msg.Text = update.Message.Text
		msg.ReplyToMessageID = int(t.MessageID)
		bot.Send(msg)
		return nil
	}

	thresholdMet := false
	message, err := Wit.Message(update.Message.Text, nil, "", "")
	if err != nil {
		return err
	}

	if _, ok := message.Entities["intent"]; !ok {
		msg.Text = "I can't answer your question let me forward it to an operator."
		_, err := bot.Send(msg)
		if err != nil {
			logrus.Fatal("Error sending msg")
		}

		t := db.Ticket{
			ChatID:    update.Message.Chat.ID,
			MessageID: int64(update.Message.MessageID),
			Text:      update.Message.Text,
		}

		ticketID := db.AddTicket(&t)

		reply := new(db.OperatorReply)
		operators := db.FindAllOperators()
		for _, op := range operators {
			chatID, err := strconv.ParseInt(op.ChatID, 10, 64)
			if err != nil {
				logrus.Fatal("Invalid chat_id")
			}

			msgForOperator := tgbotapi.NewForward(chatID, update.Message.Chat.ID, update.Message.MessageID)
			fwdRsp, err := bot.Send(msgForOperator)
			if err != nil {
				logrus.Fatal("Error Forwarding to operator")
			}

			reply.TicketID = ticketID
			reply.MessageID = int64(fwdRsp.MessageID)
			reply.OperatorID = op.ID
			db.AddReply(reply)
		}

		return nil
	}

	for _, i := range message.Entities["intent"] {
		c := i.(map[string]interface{})["confidence"].(float64)
		if c >= witConfidenceThreshold {
			thresholdMet = true
		}
	}

	if !thresholdMet {
		msg.Text = msg.Text = "I can't answer your question let me forward it to an operator."
		bot.Send(msg)
		return nil
	}

	ctype := "start"
	query := update.Message.Text
	for {
		if ctype != "start" {
			query = ""
		}
		converse, err := Wit.Converse(update.Message.Chat.UserName, query, nil)
		if err != nil {
			return err
		}
		ctype = *converse.Type

		if ctype == "stop" {
			break
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

			msg.Text = *converse.Message
			//check error
			bot.Send(msg)
		}
	}

	return nil
}

func (a *Ask) GetKeywords() []string {
	return a.Keywords
}
