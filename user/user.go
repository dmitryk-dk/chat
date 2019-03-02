package user

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	create = `
		INSERT INTO users (id, name, regDate, regTime) 
		VALUES($1,$2,$3,$4)	
	`
)

type User struct {
	ID       int
	Nickname string
	RegDate  time.Time
	RegTime  time.Time
	db       *sql.DB
}

// Create make request to database and set new user
// to table users
func (user *User) Create() error {
	bs := []byte(strconv.Itoa(user.ID))
	uuid, err := uuid.FromBytes(bs)
	if err != nil {
		return err
	}

	row, err := user.db.Query(
		create,
		uuid, user.Nickname, user.RegDate, user.RegTime,
	)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
