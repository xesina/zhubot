package db

import "github.com/Sirupsen/logrus"

type Operator struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	ChatID string `json:"chat_id"`
	Active bool   `json:"active"`
}

func FindAllOperators() []*Operator {
	rows, err := DB.Query("SELECT * FROM zhubot.operators WHERE active = TRUE")
	if err != nil {
		logrus.Fatal("Error fetching operators")
	}
	defer rows.Close()

	operators := make([]*Operator, 0)
	for rows.Next() {
		op := new(Operator)
		err = rows.Scan(&op.ID, &op.Name, &op.ChatID, &op.Active)
		if err != nil {
			logrus.Fatal("Error scanning operators")
		}
		operators = append(operators, op)
	}

	return operators
}
