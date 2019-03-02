package messages

import (
	"github.com/dmitryk-dk/chat/storage"
)

var (
	create = `
		INSERT INTO message (userId, text, date) 
		VALUES(?,?,?)
	`
)

type Message struct {
	UserID int
	Text   string
	Date   string
	DB     *storage.DbStore
}

func (msg *Message) Create() error {
	row, err := msg.DB.DB.Query(
		create,
		msg.UserID, msg.Text, msg.Date,
	)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
