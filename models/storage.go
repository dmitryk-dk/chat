package models

import (
	"database/sql"
)

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
