package handlers

import (
	"github.com/LeonardoSola/whatsapp-go-boilerplate/adapters/cmd/whastapp"
	"github.com/LeonardoSola/whatsapp-go-boilerplate/app/constants"
	"github.com/LeonardoSola/whatsapp-go-boilerplate/app/handlers"
	"github.com/LeonardoSola/whatsapp-go-boilerplate/app/models"

	"go.mau.fi/whatsmeow/types/events"
)

func MessageHandler(evt *events.Message) {
	message := models.Message{
		ID:        evt.Info.ID,
		Text:      evt.Message.GetConversation(),
		From:      evt.Info.Sender.String(),
		Recipient: evt.Info.Chat.String(),
		Timestamp: evt.Info.Timestamp,
		Platform:  constants.Whatsapp.Name,
	}

	platform := whastapp.NewWhatsappPlatform(*evt)
	if platform == nil {
		return
	}

	handlers.MessageHandler(platform, message)
}
