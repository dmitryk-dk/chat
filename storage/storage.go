package storage

import (
	"database/sql"
)

// StoreInterface describe methods which help work with
// data base
type StoreInterface interface {
	Create(args ...interface{}) error
	Read(id int) (interface{}, error)
	Update(args ...interface{}) error
	Delete(id int) error
}

// DbStore describe a store struct
type DbStore struct {
	db *sql.DB
}

// New return constructor of database connection
func New(driverName, dataSourceName string) (*DbStore, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &DbStore{db: db}, nil
}
