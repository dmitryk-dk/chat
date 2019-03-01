package storage

import (
	"database/sql"

	"github.com/dmitryk-dk/chat/user"
)

type StoreInterface interface {
	Create(user *user.User) error
	Read(user *user.User) error
	Update(id int) error
	Delete(id int) error
}

type DbStore struct {
	db *sql.DB
}

func (dbStore *DbStore) Add(user *user.User) error {
	row, err := dbStore.db.Query()
	if err != nil {
		return err
	}
}
