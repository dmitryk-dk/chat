package repository

import (
	"database/sql"

	"github.com/dmitryk-dk/chat/model"
	"github.com/google/uuid"
)

// StorageInteface describe method for working with store
type StorageInteface interface {
	CreateUser(user model.User) error
	GetUser(userID uuid.UUID) (model.User, error)

	CreateMessage(msg model.Message) error
	GetMessages(userID uuid.UUID) ([]model.Message, error)
}

// DB describe a store struct
type DB struct {
	db *sql.DB
}

// NewDB return constructor of database connection
func NewDB(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

// Close database connection
func Close(db *DB) error {
	return db.db.Close()
}
