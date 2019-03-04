package storage

import (
	"database/sql"

	repository "github.com/dmitryk-dk/chat/repository/entity"
	"github.com/google/uuid"
)

// StorageInteface describe method for working with store
type StorageInteface interface {
	CreateUser(user repository.User) error
	GetUser(userID string) (repository.User, error)

	CreateMessage(msg repository.Message) error
	GetMessages(userID uuid.UUID) ([]repository.Message, error)
}

// DB describe a store struct
type DB struct {
	*sql.DB
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
	return &DB{db}, nil
}
