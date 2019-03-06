package repository

import (
	"database/sql"

	"github.com/dmitryk-dk/chat/model"
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

type UserRepository struct {
	DB *sql.DB
}

// Create make request to database and set new user
// to table users
func (usr *UserRepository) Create(user model.User) error {
	rows, err := usr.DB.Query(
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
func (usr *UserRepository) GetUser(uuid uuid.UUID) (model.User, error) {
	var user model.User
	rows, err := usr.DB.Query(
		getUser,
		uuid,
	)
	if err != nil {
		return model.User{}, err
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Nickname, &user.RegDate)
		if err != nil {
			return model.User{}, err
		}
	}
	return user, nil
}
