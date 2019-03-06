package model

import (
	"github.com/google/uuid"
)

type Message struct {
	UserID uuid.UUID
	Text   string
	Date   string
}
