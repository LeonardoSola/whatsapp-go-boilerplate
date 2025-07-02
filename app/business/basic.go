package business

import (
	"github.com/LeonardoSola/whatsapp-go-boilerplate/app/models"
)

func Ping(platform models.Platform, message models.Message) {
	platform.Reply(message, "Pong")
}
