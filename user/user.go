package user

import (
	"github.com/dmitryk-dk/chat/storage"
)

var (
	create = `
		INSERT INTO users (nickname, regDate) 
		VALUES(?,?)	
	`
)

type User struct {
	ID       int
	Nickname string
	RegDate  string
	DB       *storage.DbStore
}

// Create make request to database and set new user
// to table users
func (user *User) Create() error {
	err := user.DB.DB.Ping()
	if err != nil {
		return err
	}

	row, err := user.DB.DB.Query(
		create,
		user.Nickname, user.RegDate,
	)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
