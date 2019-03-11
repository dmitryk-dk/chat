package repository

import (
	"database/sql"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate/source/file"

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

// MigrateUp migrate tabled to dataabase
func MigrateUp(db *DB) error {
	driver, err := postgres.WithInstance(db.db, &postgres.Config{})
	if err != nil {
		return err
	}

	fsrc, err := (&file.File{}).Open("file://repository/migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Steps(2)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	defer m.Close()
	return nil
}

// MigrateDown drop tables from database
func MigrateDown(db *DB) error {
	driver, err := postgres.WithInstance(db.db, &postgres.Config{})
	if err != nil {
		return err
	}

	fsrc, err := (&file.File{}).Open("file://repository/migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Steps(2)
	if err != nil {
		return err
	}

	err = m.Down()
	if err != nil {
		return err
	}

	defer m.Close()
	return nil
}
