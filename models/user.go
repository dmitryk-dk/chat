package models

import (
	"database/sql"
)

var (
	createUser = `
		INSERT INTO users (nickname, regDate) 
		VALUES(?,?)	
	`
)

type User struct {
	ID       int
	Nickname string
	RegDate  string
}

// Create make request to database and set new user
// to table users
func (user *User) Create(db *sql.DB) error {
	rows, err := db.Query(
		createUser,
		user.Nickname, user.RegDate,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
