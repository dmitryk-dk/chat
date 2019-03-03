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
	getUser = `
		SELECT uuid, nickname, regDate FROM users WHERE uuid = ?
	`
)

type User struct {
	ID       uuid.UUID
	Nickname string
	RegDate  string
}

// Create make request to database and set new user
// to table users
func (user *User) Create(db *sql.DB) error {
	rows, err := db.Query(
		createUser,
		user.ID, user.Nickname, user.RegDate,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (db *DB) GetUser(uuid string) (User, error) {
	var user User
	rows, err := db.Query(
		getUser,
		uuid,
	)
	if err != nil {
		return User{}, err
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Nickname, &user.RegDate)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}
