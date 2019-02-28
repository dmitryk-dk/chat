package storage

import (
	"database/sql"

	"github.com/dmitryk-dk/chart/user"
)

type StoreInterface interface {
	Create(user *user.User) error
	Remove(user *user.User) error
	Update(id int) error
	Delete(id int) error
	Get(id int) error
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
