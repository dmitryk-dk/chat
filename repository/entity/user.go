package repository

import (
	"database/sql"

	"github.com/dmitryk-dk/chat/model"
	"github.com/google/uuid"
)

var (
	createUsr = `INSERT INTO users (uuid, nickname, regDate) VALUES(?,?,?)	`
	readUsr   = `SELECT uuid, nickname, regDate FROM users WHERE uuid = ?`
	updateUsr = `UPDATE users SET uuid=$1, nickname=$2, regDate=$3 WHERE userId=$1, ID, Nickname, RegDate`
	deleteUsr = `DELETE FROM users WHERE userId=$1, ID`
)

type UserRepository struct {
	DB *sql.DB
}

// CreateUser make request to database and set new user
// to table users
func (usr *UserRepository) CreateUser(user model.User) error {
	rows, err := usr.DB.Query(
		createUsr,
		user.ID, user.Nickname, user.RegDate,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

// ReadUser return user from database
func (usr *UserRepository) ReadUser(uuid uuid.UUID) (model.User, error) {
	var user model.User
	rows, err := usr.DB.Query(readUsr, uuid)
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

// DeleteUser delete user from database
func (usr *UserRepository) DeleteUser(uuid uuid.UUID) error {
	res, err := usr.DB.Exec(deleteUsr, uuid)
	if err == nil {
		_, err := res.RowsAffected()
		if err == nil {
			return err
		}
	}

	return nil
}

// UpdateUser update user in database
func (usr *UserRepository) UpdateUser(user model.User) error {
	res, err := usr.DB.Exec(
		updateUsr,
		user.ID,
		user.Nickname,
		user.RegDate,
	)
	if err == nil {
		_, err := res.RowsAffected()
		if err == nil {
			return err
		}
	}

	return nil
}
