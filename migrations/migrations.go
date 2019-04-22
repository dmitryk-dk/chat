package migrations

import (
	"database/sql"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate/source/file"
)

// MigrateUp migrate tabled to dataabase
func MigrateUp(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	fsrc, err := (&file.File{}).Open("file://migrations/tables")
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
func MigrateDown(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	fsrc, err := (&file.File{}).Open("file://maigrations/tables")
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
