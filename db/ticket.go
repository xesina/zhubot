package db

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"time"
)

type Ticket struct {
	ID        int64     `json:"id"`
	ChatID    int64     `json:"chat_id"`
	MessageID int64     `json:"message_id"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AddTicket(t *Ticket) int64 {
	var ticketID int64
	err := DB.QueryRow(`INSERT INTO zhubot.tickets(chat_id, message_id, text) VALUES($1, $2, $3) RETURNING id`,
		t.ChatID, t.MessageID, t.Text).Scan(&ticketID)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("New ticket inserted")
	return ticketID
}

func UpdateTicket(t *Ticket) error {
	var ticketID int64
	err := DB.QueryRow("Update zhubot.tickets SET status = 'closed' RETURNING id").Scan(&ticketID)

	if err != nil {
		return err
	}
	return nil
}

func FindTicketByTicketID(tid int64) *Ticket {
	fmt.Printf("finding ticket_id: %d", tid)
	t := new(Ticket)
	err := DB.QueryRow("SELECT * FROM zhubot.tickets WHERE id = $1", tid).Scan(
		&t.ID, &t.ChatID, &t.MessageID, &t.Text, &t.Status, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		logrus.Fatal(err)
	}
	return t
}
