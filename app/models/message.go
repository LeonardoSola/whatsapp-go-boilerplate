package models

import (
	"time"
)

type Message struct {
	ID        string
	Text      string
	From      string
	Recipient string
	Platform  string
	Timestamp time.Time
	Event     any
}
