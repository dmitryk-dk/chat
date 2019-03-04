package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

var (
	createMessage = `
		INSERT INTO message (userId, text, date) 
		VALUES(?,?,?)
	`
	getMessages = `
		SELECT userId, text, date FROM message WHERE userId = ?
	`
)

type Message struct {
	UserID uuid.UUID
	Text   string
	Date   string
}

// Create make request to db and make row with new message
func (msg *Message) Create(db *sql.DB) error {
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

// GetMessages make request to db and return messages from userId
func (msg *Message) GetMessages(db *sql.DB, userId string) ([]Message, error) {
	messages := make([]Message, 1)
	rows, err := db.Query(
		getMessages,
		userId,
	)
	if err != nil {
		return []Message{}, err
	}

	for rows.Next() {
		err := rows.Scan(&msg.UserID, &msg.Text, &msg.Date)
		if err != nil {
			return []Message{}, err
		}
	}

	messages = append(messages, *msg)

	return messages, nil
}
