package models

type Platform interface {
	Reply(message Message, reply string)
}
