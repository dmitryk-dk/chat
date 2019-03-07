package repository

import (
	"database/sql"

	"github.com/dmitryk-dk/chat/model"
	"github.com/google/uuid"
)

var (
	createMsg = `INSERT INTO messages (userId, text, date) VALUES(?,?,?)`
	readMsg   = `SELECT userId, text, date FROM messages WHERE userId = ?`
	updateMsg = `UPDATE messages SET userId=$1, text=$2, date=$3 WHERE userId=$1, UserId, Text, Date`
	deleteMsg = `DELETE FROM messages WHERE userId=$1, UserId`
)

// MessageRepository describe the message instance
// for work with database
type MessageRepository struct {
	DB *sql.DB
}

// Create make request to db and make row with new message
func (msg *MessageRepository) Create(message model.Message) error {
	row, err := msg.DB.Query(
		createMsg,
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
	rows, err := msg.DB.Query(readMsg, userID)
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

// DeleteMessage delete message from database
func (msg *MessageRepository) DeleteMessage(userID uuid.UUID) error {
	res, err := msg.DB.Exec(deleteMsg, userID)
	if err == nil {
		_, err := res.RowsAffected()
		if err == nil {
			return err
		}
	}

	return nil
}

// UpdateMessage update message in database
func (msg *MessageRepository) UpdateMessage(message model.Message) error {
	res, err := msg.DB.Exec(
		updateMsg,
		message.UserID,
		message.Text,
		message.Date,
	)
	if err == nil {
		_, err := res.RowsAffected()
		if err == nil {
			return err
		}
	}

	return nil
}
