package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Nickname string
	RegDate  string
}
