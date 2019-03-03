package models

import (
	"database/sql"

	"github.com/google/uuid"
)

var (
	createUser = `
		INSERT INTO users (uuid, nickname, regDate) 
		VALUES(?,?,?)	
	`
)

type User struct {
	UUID     uuid.UUID
	Nickname string
	RegDate  string
}

// Create make request to database and set new user
// to table users
func (user *User) Create(db *sql.DB) error {
	rows, err := db.Query(
		createUser,
		user.UUID, user.Nickname, user.RegDate,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
