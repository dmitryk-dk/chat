package messages

import (
	"database/sql"

	"github.com/google/uuid"
)

var (
	create = `
		INSERT INTO message (id, text) 
		VALUES($1,$2)
	`
)

type Message struct {
	ID   uuid.UUID
	Text string
	db   *sql.DB
}

func Create(msg *Message) error {
	row, err := msg.db.Query(
		create,
		msg.ID, msg.Text,
	)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
