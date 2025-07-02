package handlers

import (
	"strings"

	"github.com/LeonardoSola/whatsapp-go-boilerplate/adapters/env"
	"github.com/LeonardoSola/whatsapp-go-boilerplate/app/business"
	"github.com/LeonardoSola/whatsapp-go-boilerplate/app/models"
)

func MessageHandler(platform models.Platform, message models.Message) {
	if !strings.HasPrefix(message.From, env.AdminPhoneNumber) {
		return
	}

	command := strings.ToUpper(message.Text)
	if _, ok := commands[command]; ok {
		commands[command](platform, message)
	}
}

var commands = map[string]func(platform models.Platform, message models.Message){
	"PING": business.Ping,
}
