package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/Sirupsen/logrus"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("postgres", "postgres://zhubot:zhubot@127.0.0.1/zhubot?sslmode=disable")
	if err != nil {
		logrus.Fatal("Error: invalid dsn")
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatal("Error connecting to database")
	}

	DB = db
}
