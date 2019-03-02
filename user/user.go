package user

import "time"

type User struct {
	Id       int
	Nickname string
	RegDate  time.Time
	RegTime  time.Time
}
