package storage

import (
	"database/sql"
	"strconv"

	"github.com/dmitryk-dk/chat/user"
	"github.com/google/uuid"
)

// StoreInterface describe methods which help work with
// data base
type StoreInterface interface {
	Create(user *user.User) error
	Read(id int) (user.User, error)
	Update(user *user.User) error
	Delete(id int) error
}

// DbStore describe a store struct
type DbStore struct {
	db *sql.DB
}

// Function Create make request to database and set new user
// to table users
func (dbStore *DbStore) Create(user *user.User) error {
	bs := []byte(strconv.Itoa(user.Id))
	uuid, err := uuid.FromBytes(bs)
	if err != nil {
		return err
	}

	row, err := dbStore.db.QueryRow(
		"INSERT INTO users (id, name, regDate, regTime) VALUES(?,?,?,?)",
		uuid, user.Nickname, user.RegDate, user.RegTime,
	)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
