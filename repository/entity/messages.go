package repository

import (
	"database/sql"

	"github.com/dmitryk-dk/chat/model"
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

type MessageRepository struct {
	DB *sql.DB
}

// Create make request to db and make row with new message
func (msg *MessageRepository) Create(message model.Message) error {
	row, err := msg.DB.Query(
		createMessage,
		message.UserID, message.Text, message.Date,
	)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}

// GetMessages make request to db and return messages from userId
func (msg *MessageRepository) GetMessages(userID uuid.UUID) ([]model.Message, error) {
	var message model.Message
	var messages []model.Message
	rows, err := msg.DB.Query(
		getMessages,
		userID,
	)
	if err != nil {
		return []model.Message{}, err
	}

	for rows.Next() {
		err := rows.Scan(&message.UserID, &message.Text, &message.Date)
		if err != nil {
			return []model.Message{}, err
		}
	}

	messages = append(messages, message)

	return messages, nil
}
