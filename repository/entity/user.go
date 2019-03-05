package repository

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
func (user *User) Create() error {
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

// GetUser return user from database
func (user *User) GetUser(uuid string) (User, error) {
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
	return *user, nil
}
