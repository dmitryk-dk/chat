package models

import (
	"database/sql"

	"github.com/google/uuid"
)

var (
	createMessage = `
		INSERT INTO message (userId, text, date) 
		VALUES(?,?,?)
	`
)

type Message struct {
	UserID uuid.UUID
	Text   string
	Date   string
}

func (msg Message) Create(db *sql.DB) error {
	row, err := db.Query(
		createMessage,
		msg.UserID, msg.Text, msg.Date,
	)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
