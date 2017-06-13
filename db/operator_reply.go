package db

import (
	"github.com/Sirupsen/logrus"
	"time"
)

type OperatorReply struct {
	ID            int64      `json:"id"`
	OperatorID    int64      `json:"operator_id"`
	OperatorReply NullString `json:"operator_reply"`
	MessageID     int64      `json:"message_id"`
	TicketID      int64      `json:"ticket_id"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func AddReply(reply *OperatorReply) {
	var replyID int64
	err := DB.QueryRow(`INSERT INTO zhubot.operator_replies(operator_id, ticket_id, message_id) VALUES($1, $2, $3) RETURNING id`,
		reply.OperatorID, reply.TicketID, reply.MessageID).Scan(&replyID)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("New reply inserted")
}

func FindReplyByMessageID(mid int64) *OperatorReply {
	rep := new(OperatorReply)
	err := DB.QueryRow("SELECT * FROM zhubot.operator_replies WHERE message_id = $1 AND status = 'open'", mid).Scan(
		&rep.ID, &rep.OperatorID, &rep.OperatorReply, &rep.MessageID, &rep.TicketID, &rep.Status, &rep.CreatedAt, &rep.UpdatedAt,
	)
	if err != nil {
		logrus.Fatal("Error fetching operator reply")
	}

	return rep
}

func UpdateReply(reply *OperatorReply) error {
	var replyID int64
	err := DB.QueryRow("Update zhubot.operator_replies SET operator_reply = $1, status = 'closed' WHERE message_id = $2 RETURNING id",
		reply.OperatorReply.String,
		reply.MessageID,
	).Scan(&replyID)

	if err != nil {
		return err
	}

	return nil
}
